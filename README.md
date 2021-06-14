# Untitled CS467 Project Starter Repo

This repo is the beginnings of a back end microservice architecture written around gRPC, gRPC-Gateway, and Go. 

## Getting Started
- *Windows*
    - *[Install WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) and follow Linux installation for your distro*

- [Install Go](https://golang.org/doc/install)

- [Install buf](https://docs.buf.build/installation/) 

- Clone repository

- Generate gRPC client/server/proxy code 
    - `buf generate`

- Navigate to root directory and run:
    - `go run . :8080 :8090`
    - Expected output: 

        ```
        (timestamp) Serving gRPC on [::]:8080
        (timestamp) Serving gRPC-Gateway on :8090
        ```
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

    

