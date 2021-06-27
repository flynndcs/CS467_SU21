# Untitled CS467 Project Starter Repo

This repo is the beginnings of a microservice architecture using gRPC, gRPC-Gateway, and FoundationDB written in Go. 

## Getting Started
- (*Windows*)
    - *[Install WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) and follow subsequent Linux installations for your distro in a convenient folder on your Linux filesystem*

- [Install Go](https://golang.org/doc/install)

- [Install buf](https://docs.buf.build/installation/) 

- [Install FoundationDB client and server](https://apple.github.io/foundationdb/downloads.html)
    - This should start automatically. To verify, run `fdbcli` on your command line and check the status.

- Clone this repository and navigate to root of repository

## Running the Services (Gateway, Product)

- Resolve dependency from buf.yaml
    - `buf beta mod update`

- run `./startService`
    - generates from protobuf files, sets logging settings and starts REST/GRPC services

## Usage (Gateway, Product API)
- Using browser, Postman or cURL:
    - Get Status: HTTP GET to localhost:8090/api/status
        - Expected response:
            ```
            { "status": "GATEWAY STATUS: NORMAL, PRODUCT STATUS: NORMAL"}
            ```
    - Put Product: HTTP POST to localhost:8090/api/product/name
        - With body:
            ```
            {"productName": "<product name>"}
            ```
        - Expected response:
            ```
            {"productName": "<product name>", "productUUID": "<random UUID>"}
            ```
    - Get Product: HTTP GET to localhost:8090/api/product/name/{productName} *productName entry must have been previously created via POST*
        - Expected response:
            ```
            {"productName": "<product name>", "productUUID": "<random UUID>"}
            ```
    - Get Product Range: HTTP GET to localhost:8090/api/product/name/{beginName}/{endName}
        - the end of the range (endName) is inclusive - a matching value for this key will be included
        - Expected response
            - all matching values for these keys in this range
            ```
            {
                "products": [
                    {
                        "productName": "<first product matching beginName>",
                        "productUUID": "<uuid>"
                    },
                    {
                        "productName": "<last product matching endName",
                        "productUUID": "<uuid>"
                    }
                ]
            }
            ```
    - Delete Product: HTTP DELETE to localhost:8090/api/product/name/{productname} *productName entry must have been previously created via POST*
        - Expected response (empty is success):
            ```
            {}
            ```


## Development

[Architecture Diagram (VERY WIP)](https://lucid.app/lucidchart/invitations/accept/inv_0a8665be-2794-4854-8e4a-c162c88fc41e?viewport_loc=-291%2C-20%2C2718%2C1354%2C0_0)

### Structure Items Of Note
- The repo currently implements minimal instances of 2 out of 4 services.
    - Gateway (gRPC with REST Proxy)
    - Product (gRPC)
    - *Network (gRPC) TBD* 
    - *Transport (gRPC) TBD*

- The **proto/service** folder contains the .proto files and the client and server code generated by the `buf` command. (The config for this is defined in the buf.yaml file)
    - ***.pb.go** files hold the generated information/structure for gRPC messages
    - **gateway_grpc.pb.go** holds the generated client and server code for the services.
    - **gateway.pb.gw.go** holds the generated code for the HTTP REST Proxy.

- The **proto/openapiv2** folder holds generated OpenAPI/Swagger documentation.
    - May be useful to serve somehow via the REST proxy.

- The **src/service** folder will hold directories for each service housing their core logic.
    - **/gateway/gatewayHandlers.go**
        - This file handles the response to each rpc action defined in the gateway.proto definition. 
    - **/gateway/gatewayServer.go**
        - This file instantiates the server for the service as specified by the generated code. 

- The **store** directory holds storage layer implementations
    - Currently in progress with FoundationDB (**store/fdb**) Worth discussing and exploring alternatives.
        - this includes a basic driver file that can be called from handlers to connect and query the database.
            - Includes get, put, and clear methods
    - FoundationDB is a key-value store meaning any create, read, update, delete operations operate on key-value pairs.
        - The current driver code uses strings for keys and encoded byte buffers for values which are deserialized from the protobuf messages and serialized back into protobuf messages upon retrieval.
        - Examples
            - Put
                - the gateway handler for PutProduct delegates to the product client's PutProduct handler which puts a value in FoundationDB using the product name as a key and a deserialized product name and random UUID for the value.
            - Get
                - the getProduct handler uses the product name as key to get the deserialized GetProductResponse (byte representation of product name and random UUID) then serializes it into the GetProductResponse format to be returned
            


        

