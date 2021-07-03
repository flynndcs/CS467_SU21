package environment

import (
	"context"
	"log"
	"net"
	"net/http"

	gatewaypb "CS467_SU21/proto/service"

	gateway "CS467_SU21/src/service/gateway"

	product "CS467_SU21/src/service/product"

	fdbDriver "CS467_SU21/src/store/fdb"

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

	mux := http.NewServeMux()
	mux.Handle("/", Authenticate(*gwmux))

	log.Println("Serving gRPC-Gateway on " + httpTarget)
	log.Fatalln(http.ListenAndServe(httpTarget, mux))
}

func Authenticate(gwmux runtime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, status := r.BasicAuth()
		if !status {
			log.Println("could not get basic auth credentials")
		}
		log.Println(r.Header.Get("Authorization"))
		if r.Header.Get("Authorization") == "" {
			w.Write([]byte("User not authenticated"))
		} else {
			if r.Method == "POST" && r.URL.Path == "/auth" {
				if !status {
					log.Println("Could not get basic auth")
					return
				}
				if !fdbDriver.CreateUser(user, password) {
					log.Println("Could not create user")
					return
				}
				return
			} else {
				if !fdbDriver.CheckCredentials(user, password) {
					log.Println("User not authenticated")
					return
				}
			}
			gwmux.ServeHTTP(w, r)
		}
	})
}
