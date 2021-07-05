package environment

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	gatewaypb "CS467_SU21/proto/service"

	gateway "CS467_SU21/src/service/gateway"

	product "CS467_SU21/src/service/product"

	fdbDriver "CS467_SU21/src/store/fdb"

	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var secret []byte

func RegisterAndServeEnvironment(tcpTarget string, httpTarget string) {
	secret = []byte(os.Getenv("SCM_APP_SECRET"))
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
	log.Fatalln(http.ListenAndServeTLS(httpTarget, "example.crt", "example.key", mux))
}

func Authenticate(gwmux runtime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, status := r.BasicAuth()
		if r.Method == "POST" && r.URL.Path == "/createUser" {
			if !fdbDriver.CreateUser(user, pass) || !status {
				log.Println("User could not be created")
			}
			return
		} else if r.URL.Path == "/getToken" {

			if !fdbDriver.CheckCredentials(user, pass) {
				log.Println("Unauthorized")
				return
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user,
				"exp": time.Now().Add(1 * time.Hour).Unix(),
			})
			tokenString, tokenErr := token.SignedString(secret)
			if tokenErr != nil {
				log.Println("error generating jwt: ", tokenErr)
				return
			}
			w.Write([]byte(tokenString))
			return
		} else {
			tokenString := r.Header.Get("Authorization")
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			parsedToken, parseErr := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
				if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", jwtToken.Header["alg"])
				}
				return secret, nil
			})
			if parseErr != nil {
				log.Println("Could not parse JWT: ", parseErr)
				return
			}
			if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
				log.Println("User: ", claims["sub"])
			}
		}
		gwmux.ServeHTTP(w, r)
	})
}
