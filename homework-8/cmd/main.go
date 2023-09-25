package main

import (
	"context"
	"fmt"
	"homework-8/internal/app/db"
	grpcserver "homework-8/internal/app/grpc_server"
	"homework-8/internal/app/grpc_server/pb"
	"homework-8/internal/app/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lsn, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "hinata34", "hinata34", "OZON")
	db, err := db.NewDatabase(ctx, psqlConn)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterGRPCServiceServer(server, grpcserver.NewImplementation(user.NewUsers(db)))

	server.Serve(lsn)
	// commandsCreated, err := commands.InitCommands()
	// if err != nil {
	// 	return
	// }

	// err = core.ProcessArgs(ctx, commandsCreated)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
