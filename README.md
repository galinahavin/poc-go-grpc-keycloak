# poc-go-grpc-keycloak
1. generate certificates (scripts in ssl can be useful)
2. make server, client 
poc-go-grpc-keycloak> make ecommerce
(protoc -Iecommerce/proto --go_opt=module=go-grpc.com/grpc-go-course --go_out=. --go-grpc_opt=module=go-grpc.com/grpc-go-course --go-grpc_out=. ecommerce/proto/*.proto)
3. start keycloack in doker; configure admin on 8180 (the project configuration and the keycloack configuration are provided in .env file )
4. import realm (ecommerce/realm_import.json) or configure realm, client, user, roles manually
5. start server in terminal
poc-go-grpc-keycloak>  .\bin\ecommerce\server.exe
6. start client in the other terminal window
poc-go-grpc-keycloak> .\bin\ecommerce\client.exe
observe different output of APIs for different user roles

Example of output

2024/01/14 18:05:55 doSubmitPuchase was invoked
2024/01/14 18:05:55 Response: user_id:3 first_name:"name_user1" last_name:"lastname_user1" email_address:"user1@gmail.com" user_id:3 from:"London" to:"Paris" price:20 section:"A" seat:7
2024/01/14 18:05:55 doGetTicketDetailsById was invoked
2024/01/14 18:05:55 Response: ticket_id:3 user_id:3 from:"London" to:"Paris" price:20 section:"A" seat:7
2024/01/14 18:05:55 doGetUserSeatDetailsBySection was invoked
2024/01/14 18:05:55 Response: [user:{user_id:3 first_name:"name_user1" last_name:"lastname_user1" email_address:"user1@gmail.com"} ticket:{ticket_id:3 user_id:3 from:"London" to:"Paris" price:20 section:"A" seat:7}] 
2024/01/14 18:05:55 doUpdateUserPuchase was invoked
2024/01/14 18:05:55 Response: ticket_id:3 user_id:3 from:"London" to:"Paris" price:20 section:"A" seat:7
2024/01/14 18:05:55 doDeleteUserPurchase was invoked
2024/01/14 18:05:55 Response: 3