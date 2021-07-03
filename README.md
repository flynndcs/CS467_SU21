# gRPC Supply Chain Management
*David Chen, Daniel Flynn*

---

- gRPC
- gRPC-Gateway
- Go
- FoundationDB


## Installation / Run (Linux only)
- run `./install`
    - downloads go, foundationdb, buf, and necessary dependencies

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
            {
                "name": "<name>",
                "scope": [<elements>]
                "expires": "<seconds since epoch>" //optional, default is 24 hours in advance
            }
            ```
        - Expected response: // TODO decide on minimum structure for naming here
            ```
            {"name": "<name>", "scope": [<elements>], "data": "<random UUID>", "expires": "<seconds since epoch>"}
            ```
    - **Get Single Product**: HTTP GET to localhost:8090/api/product?name=name&scope=element&scope=element` 
        - productName entry must have been previously created via POST and you must supply all elements as defined in scope and the name
        - Expected response:
            ```
            {"name": "<first element>", "scope": [<elements>], "data": "<random UUID>", "expires": "<seconds since epoch>"}
            ```
    - **Get Products In Scope**: HTTP GET to localhost:8090/api/product/range?scope=element&scope... 
        - must supply a minimum of one element for scoping, will match all records that were defined with the provided elements
        - Expected response
            - all matching values for these keys in this range
            ```
            {
                "products": [
                    {
                        "name": "<name>",
                        "scope": [<elements>],
                        "data": "<uuid>"
                        "expires": "<seconds since epoch>"
                    },
                    {
                        "name": "<name>",
                        "scope": [<elements>],
                        "data": "<uuid>"
                        "expires": "<seconds since epoch>"
                    }
                ]
            }
            ```
    - **Delete Product**: HTTP DELETE to localhost:8090/api/product?name=name&scope=...
        - must supply all elements for scope and name as defined when created
        - Expected response (empty is success):
            ```
            {"deletedName": "<name>", "scope": [<elements>]}
            ```



            


        

