package environment

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	gatewaypb "CS467_SU21/proto/service"

	gateway "CS467_SU21/src/service/gateway"

	product "CS467_SU21/src/service/product"

	fdbDriver "CS467_SU21/src/store/fdb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var secret []byte

func RegisterAndServeEnvironment(tcpTarget string, httpTarget string) {
	secret = []byte(os.Getenv("SCM_APP_SECRET"))
	fdbDriver.InitFDB()
	lis := createTCPListener(tcpTarget)
	createGRPCServer(lis)

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
	s := grpc.NewServer()
	gatewaypb.RegisterGatewayServer(s, gateway.NewGatewayServer())
	gatewaypb.RegisterProductServer(s, product.NewProductServer())
	log.Println("Serving gRPC on " + lis.Addr().String())
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func registerHTTPProxy(grpcTarget string, httpTarget string) {
	conn, err := grpc.DialContext(
		context.Background(),
		grpcTarget,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = gatewaypb.RegisterGatewayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", Authenticate(*gwmux))

	log.Println("Serving gRPC-Gateway on " + httpTarget)
	log.Fatalln(http.ListenAndServe(httpTarget, mux))
}
