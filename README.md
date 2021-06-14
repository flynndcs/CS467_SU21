# Untitled CS467 Project Starter Repo

This repo is the beginnings of a back end microservice architecture written around gRPC, gRPC-Gateway, and Go. 

## Getting Started
- (*Windows*)
    - *[Install WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) and follow Linux installation for your distro*

- [Install Go](https://golang.org/doc/install)

- [Install buf](https://docs.buf.build/installation/) 

- Clone repository and navigate to root of repository

- Generate gRPC client/server/proxy code 
    - `buf generate`

- Navigate to root directory and run:
    - `go run . :8080 :8090`
    - Expected output: 
        ```
        (timestamp) Serving gRPC on [::]:8080
        (timestamp) Serving gRPC-Gateway on :8090
        ```

- Or... run `./startService`
    - generate from .proto and run gRPC server and REST proxy

- Using browser, Postman or cURL:
    - HTTP GET to localhost:8090/api/status
    - Expected response:
        ```
        { "status": "GATEWAY STATUS: NORMAL, PRODUCT STATUS: NORMAL"}
        ```



## Development

[Architecture Diagram (VERY WIP)](https://lucid.app/lucidchart/invitations/accept/inv_0a8665be-2794-4854-8e4a-c162c88fc41e?viewport_loc=-291%2C-20%2C2718%2C1354%2C0_0)

### Structure
- The repo currently implements minimal instances of 2 out of 4 services.
    - Gateway (gRPC with REST Proxy)
    - Product (gRPC)
    - *Network (gRPC) TBD* 
    - *Transport (gRPC) TBD*

- The **src/service** folder will hold directories for each folder housing the core logic for each service.
    - **/gateway/gatewayHandlers.go**
        - This file handles the response to each rpc action defined in the gateway.proto definition. 
    - **/gateway/gatewayServer.go**
        - This file instantiates the server for the service as specified by the generated code. 

- The **proto/service** folder contains the .proto files and the client and server code generated by the `buf` command. (The config for this is defined in the buf.yaml file)
    - ***.pb.go** files hold the generated information/structure for gRPC messages
    - **gateway_grpc.pb.go** holds the generated client and server code for the services.
    - **gateway.pb.gw.go** holds the generated code for the HTTP REST Proxy.

- The **proto/openapiv2** folder holds generated OpenAPI/Swagger documentation.
    - May be useful to serve somehow via the REST proxy.

- The **store** directory holds storage layer implementations
    - Currently in progress with FoundationDB (**store/fdb**)? Worth discussing and exploring alternatives.

    

