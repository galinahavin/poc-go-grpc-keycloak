package main

import (
	"log"
	"net"
	"time"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

var userManager = NewUserManager()
type Server struct {
	pb.EcommerceServiceServer
}
const (
	secretKey     = "secret"
	tokenDuration = 150 * time.Minute
)

func accessibleRoles() map[string][]string {
	const ecommerceServicePath = "/ecommerce.EcommerceService/"

	return map[string][]string{
		ecommerceServicePath + "SubmitPuchase": {"admin_user, user"},
		ecommerceServicePath + "GetTicketDetailsById":  {"admin_user, user"},
		ecommerceServicePath + "GetUserSeatDetailsBySection":   {"admin_user"},
		ecommerceServicePath + "DeleteUserPurchase":   {"admin_user"},
		ecommerceServicePath + "UpdateUserPuchase":   {"admin_user"},
	}
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	jwtManager := NewJWTManager(secretKey, tokenDuration)
	keycloack := newKeycloak()
	userManager.CreateUsersInDB()
	interceptor := NewAuthInterceptor(jwtManager, keycloack, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading cetificates: %v\n", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	s := grpc.NewServer(serverOptions...)

	pb.RegisterEcommerceServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
