# gRPC Supply Chain Management
*David Chen, Daniel Flynn*

---

- gRPC
- gRPC-Gateway
- Go
- FoundationDB


## Installation / Run (Linux only)
- install Golang (1.16.5)
- run `./install`
    - downloads foundationdb, buf, and necessary dependencies

- run `./startService`
    - generates from protobuf files, sets logging settings and starts REST/GRPC services
    - If this script hangs on "Initializing FDB", run `sudo service foundationdb start`

## Usage (Gateway, Product API) 

- Using browser, Postman or cURL:
    - (cURL - make sure to enclose URLs in single or double quotes to recognize all query parameters)
    - **Create Account** HTTP POST to localhost:8090/createAccount
        - With basic auth headers for username, password
        - With body:
        ```
            {
                "accountName": <name>,
                "users": [<known users, may be empty>]
            }
        ```
        - Expected response:
        ```
            {
                "accountName": <name>,
            }
        ```
    - **Create User** HTTP POST to localhost:8090/createUser
        - With basic auth headers for username, password
        - With body:
        ```
            {
                "accountName": <name>,
            }
        ```
        - Expected response:
        ```
        ```
    - **Get JWT Token** HTTP GET to localhost:8090/getToken
        - With basic auth headers for same username, password as created previously
        - With body:
        ```
            {
                "accountName": <name>,
            }
        ```
        - Expected response:
        ```
        <token>
        ```
        - **This JWT token must be included in the Authorization header for each subsequent request.**
    - **Get Status**: HTTP GET to localhost:8090/api/status
        - Expected response:
            ```
            { "status": "GATEWAY STATUS: NORMAL, PRODUCT STATUS: NORMAL"}
            ```
    - **Create Product**: HTTP POST to localhost:8090/api/product
        - With body:
            ```
            {
                "productIdentifier": {"id": <id>}, 
                "name": "<name>", 
                "categories": [<ordered categories>],
                "tags": [<ordered tags>], 
                "origin": "<origin>", 
                "intermediateDestinations": [<destinations>],
                "endDestinations": [<destinations>], 
                "totalQuantity": <quantity>,
                "localProductFamily": {
                    "childProducts": [{"id": <id>},...]
                    "self": {"id": <id>}
                    "parentProducts": [{"id": <id>},...]
            }
            ```
        - Expected response: 
            - same as above with full product family, quantity by location, and quantity in transit added
    - **Get Single Product**: HTTP GET to localhost:8090/api/product?id=...
        - Expected response:
            - same as when product created
           
    - **Search Products By Category**: HTTP GET to localhost:8090/api/product/search?categories=...&categories=...
        - must supply a minimum of one category
        - Expected response
            - all matching products for this/these category(ies)
    - **Search Products By Origin**: HTTP GET to localhost:8090/api/product/search?origin=...
            - Expected response
                - all matching products for this origin
    - **Search Products By Tags**: HTTP GET to localhost:8090/api/product/search?tags=...&tags=...
            - must supply a minimum of one tag
            - Expected response
                - all matching products for this/these tag(s)

    - **Delete Product**: HTTP DELETE to localhost:8090/api/product?id=...&categories=...&tags=...&origin=...
        - must supply id, all categories, all tags, and origin to fully delete
        - Expected response (empty is success):
            ```
            {
                "id": <id>,
                "categories": [<categories>...],
                "tags": [<tags>],
                "origin": <origin>
            }
            ```



            


        

