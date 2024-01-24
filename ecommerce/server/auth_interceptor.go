package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)
type keycloak struct {
	gocloak      *gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func newKeycloak() *keycloak {
	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8180"),
		clientId:     "go-grpc-client",
		clientSecret: "secret",
		realm:        "ecommerce",
	}
}



// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	jwtManager      *JWTManager
	keycloak 		*keycloak	
	accessibleRoles map[string][]string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager, keycloak *keycloak, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, keycloak, accessibleRoles}
}

// function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("debug --> unary interceptor info: ", info)
		log.Println("--> unary interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)
		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
    accessibleRoles = strings.Split(accessibleRoles[0], ",")
	for i, v := range accessibleRoles {
        accessibleRoles[i] = strings.TrimSpace(v)
    }
	if !ok {
		return nil
	}
	var values []string
	var accessToken string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	
	if ok {
		values = md.Get("authorization")
	}

	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	if len(values) > 0 {
		accessToken = values[0]
	}
	accessToken = strings.Replace(accessToken, "Bearer ", "", 1)
	// call Keycloak API to verify the access token
	result, err := interceptor.keycloak.gocloak.RetrospectToken(context.Background(), accessToken, interceptor.keycloak.clientId, interceptor.keycloak.clientSecret, interceptor.keycloak.realm)
	if err != nil {
		log.Fatalf("Invalid or malformed token: %s", err.Error())
	}

	jwt, _, err := interceptor.keycloak.gocloak.DecodeAccessToken(context.Background(), accessToken, interceptor.keycloak.realm)
	if err != nil {
		log.Fatalf("Invalid or malformed token: %s", err.Error())
	}

	jwtj, _ := json.Marshal(jwt)

	log.Println("--> json.Marshal got jwtj token: \n", string(jwtj))

	// check if the token isn't expired and valid
	if !*result.Active {
		log.Fatalf("Invalid or expired Token %v", http.StatusUnauthorized)
	}
	roles, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	for _, accessibleRole := range accessibleRoles {
		for _, role := range roles {
			if (accessibleRole == role) {
				return nil
			}
		}
	}
	return status.Error(codes.PermissionDenied, "User does not have permission to access this RPC")

}