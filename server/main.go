package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	petv1 "github.com/bufbuild/buf-tour/gen/pet/v1"
	"github.com/bufbuild/buf-tour/gen/pet/v1/petv1connect"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const address = "localhost:2080"

func main() {
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(&petStoreServiceServer{})
	mux.Handle(path, handler)

	log.Println("... Listening on", address)

	go func() {
		log.Fatalln(http.ListenAndServe(
			address,
			//Use h2c so we can serve HTTP/2 without TLS
			h2c.NewHandler(mux, &http2.Server{}),
		))
	}()

	conn, err := grpc.NewClient("0.0.0.0:2080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)

	}
	gwmux := runtime.NewServeMux()

	err = petv1.RegisterPetStoreServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    ":9085",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9085")
	log.Fatalln(gwServer.ListenAndServe())

}

// petStoreServiceServer implements the PetStoreService API
type petStoreServiceServer struct {
	petv1connect.UnimplementedPetStoreServiceHandler
}

func (s *petStoreServiceServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest]) (*connect.Response[petv1.PutPetResponse], error) {

	name := req.Msg.GetName()
	petType := req.Msg.GetPetType()

	if name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid argument"))
	}

	fmt.Println("Received request to add pet with name: ", req.Msg.Name)
	log.Printf("Got a request to create a %v named %s", petType, name)

	return connect.NewResponse(&petv1.PutPetResponse{
		Pet: &petv1.Pet{
			Name:    name,
			PetType: petType,
			PetId:   (uuid.New()).String(),
		},
	}), nil
}
