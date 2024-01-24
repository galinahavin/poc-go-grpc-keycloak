# poc-go-grpc-keycloak

==============================================
Build and run the project
===============================================

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
7. run unit tests which demo API are protected (RBAC with KeyCloak)
==============================================
Unit tests and the of output:
===============================================
PS C:\Users\galin\poc-go-grpc-keycloak> go test ./... -v -count=1 
?       go-grpc.com/grpc-go-course/ecommerce/proto      [no test files]
?       go-grpc.com/grpc-go-course/ecommerce/server     [no test files]
=== RUN   TestGRPC
=== RUN   TestGRPC/submitPurchaseAsRegularUser
    ecommerce_test.go:239: In SubmitPuchase test
=== RUN   TestGRPC/submitPurchaseAsAdminUser
    ecommerce_test.go:239: In SubmitPuchase test
=== RUN   TestGRPC/getTicketDetailsByIdAsRegularUser
=== RUN   TestGRPC/getTicketDetailsByIdAsAdminUser
=== RUN   TestGRPC/getUserSeatDetailsBySectionAsRegularUser
    ecommerce_test.go:196: rpc error: code = PermissionDenied desc = User does not have permission to access this RPC
=== RUN   TestGRPC/getUserSeatDetailsBySectionAsAdminUser
=== RUN   TestGRPC/updateUserPuchaseAsRegularUser
    ecommerce_test.go:178: rpc error: code = PermissionDenied desc = User does not have permission to access this RPC
=== RUN   TestGRPC/updateUserPuchaseAsAdminUser
=== RUN   TestGRPC/deleteUserPuchaseAsRegularUser
    ecommerce_test.go:160: rpc error: code = PermissionDenied desc = User does not have permission to access this RPC
=== RUN   TestGRPC/deleteUserPuchaseAsAdminUser
--- PASS: TestGRPC (0.81s)
    --- PASS: TestGRPC/submitPurchaseAsRegularUser (0.09s)
    --- PASS: TestGRPC/submitPurchaseAsAdminUser (0.08s)
    --- PASS: TestGRPC/getTicketDetailsByIdAsRegularUser (0.08s)
    --- PASS: TestGRPC/getTicketDetailsByIdAsAdminUser (0.08s)
    --- PASS: TestGRPC/getUserSeatDetailsBySectionAsAdminUser (0.08s)
    --- PASS: TestGRPC/updateUserPuchaseAsRegularUser (0.08s)
    --- PASS: TestGRPC/updateUserPuchaseAsAdminUser (0.08s)
    --- PASS: TestGRPC/deleteUserPuchaseAsRegularUser (0.08s)
    --- PASS: TestGRPC/deleteUserPuchaseAsAdminUser (0.08s)
PASS
ok      go-grpc.com/grpc-go-course/ecommerce/client     0.951s

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