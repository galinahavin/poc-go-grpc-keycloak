package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

var addr string = getConfig("GRPC_SERVER_ADDR")

func getConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}



func main() {

	config := oauth2.Config{
		ClientID: getConfig("KEYCLOAK_CLIENT_ID"),
		ClientSecret: getConfig("KEYCLOAK_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			TokenURL: getConfig("KEYCLOAK_TOKEN_URL"),
		},
		Scopes: []string{"openid"},
	}
	ctx := context.Background()

	// TODO: on the server side, use the hashed password, use bcrypt to hash the password first
	//TODO: define a method on the user to check if a given password is correct or not. 
	// (call bcrypt.CompareHashAndPassword() function, 
	// pass in the user’s hashed password, and the given plaintext password. 
	// return true if error is nil)

	// get a token
	token, err := config.PasswordCredentialsToken(ctx, getConfig("KEYCLOAK_USERNAME"), getConfig("KEYCLOAK_PASSWORD"))

	if err != nil {
		panic(err)
	}

	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt" 
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewEcommerceServiceClient(conn)
	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))
	doSubmitPuchase(client, ctx)
	doGetTicketDetailsById(client, ctx)
	doGetUserSeatDetailsBySection(client, ctx)
	doUpdateUserPuchase(client, ctx)
	doDeleteUserPurchase(client, ctx)

}

