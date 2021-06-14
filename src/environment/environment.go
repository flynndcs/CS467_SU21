package environment

import (
	"context"
	"log"
	"net"
	"net/http"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	gateway "flynndcs.com/flynndcs/grpc-gateway/src/service/gateway"
	product "flynndcs.com/flynndcs/grpc-gateway/src/service/product"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterAndServeEnvironment(tcpTarget string, httpTarget string) {
	lis := createTCPListener(tcpTarget)
	createGRPCServer(lis)
	registerHTTPProxy(tcpTarget, httpTarget)
}

func createTCPListener(target string) net.Listener {
	lis, err := net.Listen("tcp", target)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	return lis
}

func createGRPCServer(lis net.Listener) {
	//create the gRPC server that manages gRPC services
	s := grpc.NewServer()
	// Attach the services to the server
	gatewaypb.RegisterGatewayServer(s, &gateway.GatewayServer{})
	gatewaypb.RegisterProductServer(s, &product.ProductServer{})
	// Serve gRPC Server
	log.Println("Serving gRPC on " + lis.Addr().String())
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func registerHTTPProxy(grpcTarget string, httpTarget string) {
	//create a client connection for the HTTP proxy
	conn, err := grpc.DialContext(
		context.Background(),
		grpcTarget,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	//register the handler for REST calls
	gwmux := runtime.NewServeMux()
	err = gatewaypb.RegisterGatewayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    httpTarget,
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on " + gwServer.Addr)
	log.Fatalln(gwServer.ListenAndServe())
}
