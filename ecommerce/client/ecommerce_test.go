package main

import (
	"context"
	"log"
	"testing"

	pb "go-grpc.com/grpc-go-course/ecommerce/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

type updateUserPuchaseTest struct {
	name            string
	KeyCloakUser    string
	KeyCloakPasswd  string
	TicketId         int64	
	Seat             int32		
}

type deleteUserPurchaseTest struct {
	name            string
	KeyCloakUser    string
	KeyCloakPasswd  string
	TicketId         int64			
}

type submitTest struct {
	name            string
	KeyCloakUser    string
	KeyCloakPasswd  string	
	UserId           int64
	From            string
	To              string
	Price          float64
	Section         string
	Seat             int32
}

type getTicketDetailsByIdTest struct {
	name              string
	KeyCloakUser      string
	KeyCloakPasswd    string		
	TicketId           int64
}

type getUserSeatDetailsBySectionTest struct {
	name              string
	KeyCloakUser      string
	KeyCloakPasswd    string		
	Section           string	
}

var updateUserPuchaseTests = []updateUserPuchaseTest{
	{
		name:                           "updateUserPuchaseAsRegularUser",
		KeyCloakUser:                                            "user1",
		KeyCloakPasswd:                                          "user1",
		TicketId:                                                      3,	
		Seat:                                                          3,
	},
	{
		name:                           "updateUserPuchaseAsAdminUser",
		KeyCloakUser:                                            "admin1",
		KeyCloakPasswd:                                          "admin1",
		TicketId:                                                      3,	
		Seat:                                                          3,
	},			
}

var deleteUserPurchaseTests = []deleteUserPurchaseTest{
	{
		name:                           "deleteUserPuchaseAsRegularUser",
		KeyCloakUser:                                            "user1",
		KeyCloakPasswd:                                          "user1",
		TicketId:                                                      3,	

	},
	{
		name:                              "deleteUserPuchaseAsAdminUser",
		KeyCloakUser:                                            "admin1",
		KeyCloakPasswd:                                          "admin1",
		TicketId:                                                      3,	

	},	
}

var getUserSeatDetailsBySectionTests = []getUserSeatDetailsBySectionTest{
	{
		name: 					"getUserSeatDetailsBySectionAsRegularUser",
		KeyCloakUser:                                              "user1",
		KeyCloakPasswd:                                            "user1",
		Section: 								                       "A",
	},

	{
		name: 					"getUserSeatDetailsBySectionAsAdminUser",
		KeyCloakUser:                          "admin1",
		KeyCloakPasswd:                        "admin1",		
		Section: 									"A",
	},	
}

var getTicketDetailsByIdTests = []getTicketDetailsByIdTest{
	{
		name: 					"getTicketDetailsByIdAsRegularUser",
		KeyCloakUser:                          "user1",
		KeyCloakPasswd:                        "user1",		
		TicketId: 									 3,
	},
	{
		name: 					"getTicketDetailsByIdAsAdminUser",
		KeyCloakUser:                          "admin1",
		KeyCloakPasswd:                        "admin1",		
		TicketId: 									 4,
	},	
}

var submitTests = []submitTest{

	{

		name:            "submitPurchaseAsRegularUser", 
		KeyCloakUser:                          "user1",
		KeyCloakPasswd:                        "user1",
		UserId:                                      3,
		From:                                 "London",
		To:                                    "Paris",
		Price:                                    20.0,
		Section:                                   "A",
		Seat :                                       7,
	},

	{
		name:               "submitPurchaseAsAdminUser",
		KeyCloakUser:                          "admin1",
		KeyCloakPasswd:                        "admin1",	
		UserId:                                       4,
		From:                                  "London",
		To:                                     "Paris",
		Price:                                     20.0,
		Section:                                    "B",
		Seat :                                        8,
	},

}

func RunDeleteUserPuchase(t *testing.T, ctx context.Context, config oauth2.Config, c pb.EcommerceServiceClient, KeyCloakUser string, KeyCloakPasswd string, TicketId int64) {
	// keycloak client (defined in config) gets a token for a keycloak user
	token, err := config.PasswordCredentialsToken(ctx, KeyCloakUser, KeyCloakPasswd)
	if err != nil {
		panic(err)
	}

	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))
	if _, err := c.DeleteUserPurchase(ctx, &pb.DeleteUserPurchaseRequest{TicketId: TicketId,}); err != nil {
		if t.Name() == "TestGRPC/deleteUserPuchaseAsRegularUser"{
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	
}

func RunUpdateUserPuchase(t *testing.T, ctx context.Context, config oauth2.Config, c pb.EcommerceServiceClient, KeyCloakUser string, KeyCloakPasswd string, TicketId int64, Seat int32) {
	// keycloak client (defined in config) gets a token for a keycloak user
	token, err := config.PasswordCredentialsToken(ctx, KeyCloakUser, KeyCloakPasswd)
	if err != nil {
		panic(err)
	}

	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))
	if _, err := c.UpdateUserPuchase(ctx, &pb.UpdateUserPuchaseRequest{TicketId: TicketId, Seat: Seat,}); err != nil {
		if t.Name() == "TestGRPC/updateUserPuchaseAsRegularUser"{
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	
}

func RunGetUserSeatDetailsBySection(t *testing.T, ctx context.Context, config oauth2.Config, c pb.EcommerceServiceClient, KeyCloakUser string, KeyCloakPasswd string, Section string) {
	// keycloak client (defined in config) gets a token for a keycloak user
	token, err := config.PasswordCredentialsToken(ctx, KeyCloakUser, KeyCloakPasswd)
	if err != nil {
		panic(err)
	}

	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))
	if _, err := c.GetUserSeatDetailsBySection(ctx, &pb.GetUserSeatDetailsBySectionRequest{Section: Section}); err != nil {
		if t.Name() == "TestGRPC/getUserSeatDetailsBySectionAsRegularUser"{
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	
}


func RunGetTicketDetailsById(t *testing.T, ctx context.Context, config oauth2.Config, c pb.EcommerceServiceClient, KeyCloakUser string, KeyCloakPasswd string, TicketId int64) {
	// keycloak client (defined in config) gets a token for a keycloak user
	token, err := config.PasswordCredentialsToken(ctx, KeyCloakUser, KeyCloakPasswd)

	if err != nil {
		panic(err)
	}

	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))
	if _, err := c.GetTicketDetailsById(ctx, &pb.GetTicketDetailsRequest{TicketId: TicketId}); err != nil {
		t.Fatal(err)
	}
}

func RunSubmitPuchase(t *testing.T, ctx context.Context, config oauth2.Config, c pb.EcommerceServiceClient, KeyCloakUser string, KeyCloakPasswd string, UserId int64, From string, To string, Price float64, Section string, Seat int32) {
	ticket:= &pb.Ticket{

		UserId: 			UserId,
		From:                 From,
		To:  		            To,
		Price:				 Price,
		Section:	  	   Section,
		Seat:		  		  Seat,
	}

	// keycloak client (defined in config) gets a token for a keycloak user
	token, err := config.PasswordCredentialsToken(ctx, KeyCloakUser, KeyCloakPasswd)

	if err != nil {
		panic(err)
	}

	ctx = grpcMetadata.NewOutgoingContext(context.Background(), grpcMetadata.Pairs("Authorization", "Bearer "+token.AccessToken))

	t.Log("In SubmitPuchase test")
	if _, err := c.SubmitPuchase(ctx, &pb.SubmitPuchaseRequest{UserId: 3, Ticket: ticket,}); err != nil {
		t.Fatal(err)
	}

}
func TestGRPC(t *testing.T) {
	ctx := context.Background()
	config := oauth2.Config{
		ClientID: getConfig("KEYCLOAK_CLIENT_ID"),
		ClientSecret: getConfig("KEYCLOAK_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			TokenURL: getConfig("KEYCLOAK_TOKEN_URL"),
		},
		Scopes: []string{"openid"},
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
	for _, test := range submitTests {
		t.Run(test.name, func(t *testing.T) {
			RunSubmitPuchase(t,
				ctx,
				config,
				client, 
				test.KeyCloakUser, 
				test.KeyCloakPasswd, 
				test.UserId, 
				test.From, 
				test.To,
				test.Price,
				test.Section,
				test.Seat)
		})
	}
	for _, test := range getTicketDetailsByIdTests {
		t.Run(test.name, func(t *testing.T) {
			RunGetTicketDetailsById(t,
				ctx,
				config,
				client, 
				test.KeyCloakUser, 
				test.KeyCloakPasswd, 
				test.TicketId)
		})
	}
	for _, test := range getUserSeatDetailsBySectionTests {
		t.Run(test.name, func(t *testing.T) {
			RunGetUserSeatDetailsBySection(t,
				ctx,
				config,
				client, 
				test.KeyCloakUser, 
				test.KeyCloakPasswd, 
				test.Section)
		})
	}
	for _, test := range updateUserPuchaseTests {
		t.Run(test.name, func(t *testing.T) {
			RunUpdateUserPuchase(t,
				ctx,
				config,
				client, 
				test.KeyCloakUser, 
				test.KeyCloakPasswd, 
				test.TicketId,
				test.Seat)
		})
	}
	for _, test := range deleteUserPurchaseTests {
		t.Run(test.name, func(t *testing.T) {
			RunDeleteUserPuchase(t,
				ctx,
				config,
				client, 
				test.KeyCloakUser, 
				test.KeyCloakPasswd, 
				test.TicketId)
		})
	}			
}
