# GrpcUser
    A gRPC microservice.

---

## Project Structure

    1. client/ : Rest Client
    2. cmd/ : All the main.go.
    3. server/ : Grpc server
    4. user/ : User proto files and generated files.
    5. *.Dockerfile: Dockerfiles for client and server
    6. deploy.sh : Simple Script to deploy to docker. (It does not handle error right now, keep the ports :8080 and :50051 free)
    7. generate-proto.sh : Generate the user.proto. (Generated files already exist).
    8. CorpUser.postman_collection.json: JSON file of collection for exporting in Postman.

*Populated the list of users in the starting*

---

## REST ENDPOINTS:
    a. /users/{id} : GET USER BY ID
    b. /users/ : GET ALL USERS
    c. /users?id=1,2,4 : GET USERS BASED ON THE LIST OF IDS PROVIDED IN QUERY

*If any invalid id provided in (c) endpoints like /users?id=1,2,"hello" : hello will be ignored and remaining ids 1,2 result would be provided.*

---

## Code Coverage :
    ```
            go test -coverprofile cover.out ./server
            ok      corpuser/server 1.318s  coverage: 100.0% of statements
    ```
    
*Golang version: 1.17*
*Docker Desktop version:4.0*
*Operating System Used: MacOS*
*If not able to run on Windows, install Git Bash, and run below commands*


## Instructions
    1. Locally
        a. Build the grpc server: `go build -o bin/grpcServer cmd/server/main.go`

        b. Build the grpc client : `go build -o bin/grpcClient cmd/client/main.go`
    
        c. In one terminal run:  ./bin/grpcServer
        d. In another terminal run: ./bin/grpcClient

    2. Via Docker
        a. Build docker image for grpc-server: docker build -t grpc-server -f server.Dockerfile .
        b. Build docker image for client: docker build -t grpc-client -f client.Dockerfile .  
        c. Run the grpc container :docker run -d --name grpc-server -p 50051:50051 grpc-server
        d. Run the client container docker run -d --name grpc-client -p 8080:8080 -e USER_ENDPOINT=grpc-server:50051 --link=grpc-server grpc-client
    
    3. Via Docker with with single file
        a. chmod 775 ./deploy.sh
        b. ./deploy.sh


## Demo
    a. GET USER BY ID
```
        curl -i http://0.0.0.0:8080/users/1
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Sun, 13 Feb 2022 12:38:46 GMT
        Content-Length: 88

        {"id":1,"fname":"Hello1","city":"Dubai1","phone":123456789,"height":5.8,"married":true}

```

    b. GET USERS BASED ON LIST OF IDS.
    ```
            curl -i http://0.0.0.0:8080/users/\?id\=1,3
            HTTP/1.1 200 OK
            Content-Type: application/json
            Date: Sun, 13 Feb 2022 12:40:45 GMT
            Content-Length: 178

            [{"id":1,"fname":"Hello1","city":"Dubai1","phone":123456789,"height":5.8,"married":true},{"id":3,"fname":"Hello3","city":"Dubai3","phone":123456789,"height":5.8,"married":true}]
    ```

    c. GET USERS BASED ON LIST OF IDS. (if no id given)
    ```
            curl -i http://0.0.0.0:8080/users/    
            HTTP/1.1 200 OK
            Content-Type: application/json
            Date: Sun, 13 Feb 2022 12:42:10 GMT
            Content-Length: 885

            [{"id":1,"fname":"Hello1","city":"Dubai1","phone":123456789,"height":5.8,"married":true},{"id":2,"fname":"Hello2","city":"Dubai2","phone":123456789,"height":5.8,"married":true},{"id":3,"fname":"Hello3","city":"Dubai3","phone":123456789,"height":5.8,"married":true},{"id":4,"fname":"Hello4","city":"Dubai4","phone":123456789,"height":5.8,"married":true},{"id":5,"fname":"Hello5","city":"Dubai5","phone":123456789,"height":5.8,"married":true},{"id":6,"fname":"Hello6","city":"Dubai6","phone":123456789,"height":5.8,"married":true},{"id":7,"fname":"Hello7","city":"Dubai7","phone":123456789,"height":5.8,"married":true},{"id":8,"fname":"Hello8","city":"Dubai8","phone":123456789,"height":5.8,"married":true},{"id":9,"fname":"Hello9","city":"Dubai9","phone":123456789,"height":5.8,"married":true},{"id":10,"fname":"Hello10","city":"Dubai10","phone":123456789,"height":5.8,"married":true}]
    ```

    d. GET USERS BASED ON LIST OF IDS. (if no id given)
    ```
            curl -i http://0.0.0.0:8080/users/    
            HTTP/1.1 200 OK
            Content-Type: application/json
            Date: Sun, 13 Feb 2022 12:42:10 GMT
            Content-Length: 885

            [{"id":1,"fname":"Hello1","city":"Dubai1","phone":123456789,"height":5.8,"married":true},{"id":2,"fname":"Hello2","city":"Dubai2","phone":123456789,"height":5.8,"married":true},{"id":3,"fname":"Hello3","city":"Dubai3","phone":123456789,"height":5.8,"married":true},{"id":4,"fname":"Hello4","city":"Dubai4","phone":123456789,"height":5.8,"married":true},{"id":5,"fname":"Hello5","city":"Dubai5","phone":123456789,"height":5.8,"married":true},{"id":6,"fname":"Hello6","city":"Dubai6","phone":123456789,"height":5.8,"married":true},{"id":7,"fname":"Hello7","city":"Dubai7","phone":123456789,"height":5.8,"married":true},{"id":8,"fname":"Hello8","city":"Dubai8","phone":123456789,"height":5.8,"married":true},{"id":9,"fname":"Hello9","city":"Dubai9","phone":123456789,"height":5.8,"married":true},{"id":10,"fname":"Hello10","city":"Dubai10","phone":123456789,"height":5.8,"married":true}]
    ```

    e. INVALID ID PROVIDED IN THE LIST OF USERS
    ```
            curl -i http://0.0.0.0:8080/users\?id\="hello"
            HTTP/1.1 400 Bad Request
            Content-Type: application/json
            Date: Sun, 13 Feb 2022 16:02:54 GMT
            Content-Length: 56

            {"message":"No valid ids provided: hello","status":400}
    ```
