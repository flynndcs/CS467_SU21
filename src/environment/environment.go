package environment

import (
	"context"
	"log"
	"net"
	"net/http"

	gatewaypb "flynndcs.com/flynndcs/grpc-gateway/proto/service"
	gateway "flynndcs.com/flynndcs/grpc-gateway/src/service/gateway"
	product "flynndcs.com/flynndcs/grpc-gateway/src/service/product"
	fdbDriver "flynndcs.com/flynndcs/grpc-gateway/src/store/fdb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterAndServeEnvironment(tcpTarget string, httpTarget string) {
	fdbDriver.InitFDB()
	lis := createTCPListener(tcpTarget)
	createGRPCServer(lis)

	//validate the creation of connection and clients for the gRPC services
	gateway.CreateGRPCConnAndClients()

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
	//create the gRPC server that manages gRPC service server instances
	s := grpc.NewServer()
	// Attach the service server instances to the grpc server
	gatewaypb.RegisterGatewayServer(s, gateway.NewGatewayServer())
	gatewaypb.RegisterProductServer(s, product.NewProductServer())
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
