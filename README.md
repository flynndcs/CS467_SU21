# gRPC Supply Chain Management
*David Chen, Daniel Flynn, Logan Kiser*

---

- gRPC
- gRPC-Gateway
- Go
- FoundationDB


## Getting Started
- (*Windows*)
    - *[Install WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) and follow subsequent Linux installations for your distro in a convenient folder on your Linux filesystem*

- [Install Go](https://golang.org/doc/install)

- [Install buf](https://docs.buf.build/installation/) 

- [Install FoundationDB client and server](https://apple.github.io/foundationdb/downloads.html)
    - This should start automatically. To verify, run `fdbcli` on your command line and check the status.
    - If not running, `sudo service foundationdb start`

- Clone this repository and navigate to root of repository

## Running the Services (Gateway, Product)

- Resolve dependency from buf.yaml
    - `buf beta mod update`

- run `./startService`
    - generates from protobuf files, sets logging settings and starts REST/GRPC services

## Usage (Gateway, Product API) 
- See Development below for explanations on implementation.
- Using browser, Postman or cURL:
    - (cURL - make sure to enclose URLs in single or double quotes to recognize all query parameters)
    - **Get Status**: HTTP GET to localhost:8090/api/status
        - Expected response:
            ```
            { "status": "GATEWAY STATUS: NORMAL, PRODUCT STATUS: NORMAL"}
            ```
    - **Put Product**: HTTP POST to localhost:8090/api/product
        - With body:
            ```
            {"scope": [<elements>]}
            ```
        - Expected response: // TODO decide on minimum structure for naming here
            ```
            {"productName": "<first element>", "productUUID": "<random UUID>"}
            ```
    - **Get Single Product**: HTTP GET to localhost:8090/api/product?scope=element&scope=element` 
        - productName entry must have been previously created via POST and you must supply all elements as defined in scope
        - Expected response:
            ```
            {"productName": "<first element>", "productUUID": "<random UUID>"}
            ```
    - **Get Products In Scope**: HTTP GET to localhost:8090/api/product/range?scope=element&scope... 
        - must supply a minimum of one element for scoping, will match all records that were defined with the provided elements
        - Expected response
            - all matching values for these keys in this range
            ```
            {
                "products": [
                    {
                        "productName": "<first element>",
                        "productUUID": "<uuid>"
                    },
                    {
                        "productName": "<first element",
                        "productUUID": "<uuid>"
                    }
                ]
            }
            ```
    - **Delete Product**: HTTP DELETE to localhost:8090/api/product?scope=...
        - must supply all elements for scopes as defined when created
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

- To clear the FoundationDB database of all key-value records, run the following commands:
    - `fdbcli`
    - `writemode on`
    - `clearrange "" \xFF`

- The **proto/service** folder contains the .proto files and the client and server code generated by the `buf` command. (The config for this is defined in the buf.yaml file)
    - ***.pb.go** files hold the generated information/structure for gRPC messages
    - **gateway_grpc.pb.go** holds the generated client and server code for the services.
    - **gateway.pb.gw.go** holds the generated code for the HTTP REST Proxy.

- The **proto/openapiv2** folder holds generated OpenAPI/Swagger documentation.
    - May be useful to serve somehow via the REST proxy.

- The **src/service** folder will hold directories for each service housing their core logic.
    - **/gateway/gatewayHandlers.go**
        - This file handles the response to each rpc action defined in the service.proto definition. 
    - **/gateway/gatewayServer.go**
        - This file instantiates the server for the service as specified by the generated code. 
    - **/gateway/productDelegators.go**
        - This file contains the handlers that delegate calls to product service.

- The **store** directory holds the FoundationDB implementation layer
    - FoundationDB is a key-value store meaning any create, read, update, delete operations operate on key-value pairs.
        - The current driver code uses strings for keys and encoded byte buffers for values which are deserialized from the protobuf messages and serialized back into protobuf messages upon retrieval.
    - Product Service
        - Each record is defined within a multi-element "scope" - each provided scope element provides more specificity about the categorization of a product when created or retrieved.
            - Example - a product could be defined as ["Coffee", "Mexico"] to represent a coffee product from Mexico. Similarly, ["Coffee", "Guatemala"]
            - Use the `product/range` endpoint with `?scope=Coffee` to get the records for both Mexican coffee and Guatemalan coffee
            - Use the `product` endpoint to get a single record which requires an exact and distinct scope. `?scope=Coffee&scope=Mexico` will return only the record that has a key defined with ["Coffee", "Mexico"]

        - Examples
            - Put
                - the gateway handler for PutProduct delegates to the product client's PutProduct handler which puts a value in FoundationDB using the product name as a key and a deserialized product name and random UUID for the value.
            - Get
                - the getProduct handler uses the product name as key to get the deserialized GetSingleProductResponse (byte representation of product name and random UUID) then serializes it into the GetSingleProductResponse format to be returned
            


        

