package main

import (
	"context"
	"log"
	"time"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAnywhatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	tp, err := ptypes.TimestampProto(t)
	if err != nil {
		log.Fatalf(err.Error())
	}

	pfx := t.Format(time.RFC3339Nano)

	req1 := pb.CreateAnythingRequest{
		Anything: &pb.Anything{
			Name:        "name",
			Description: "description (" + pfx + ")",
			CreatedAt:   tp,
			UpdatedAt:   tp,
		},
	}
	res1, err := c.CreateAnything(ctx, &req1)
	if err != nil {
		log.Fatalf("CreateAnything failed: %v", err)
	}
	log.Printf("CreateAnything result: <%+v>\n\n", res1)

	id := res1.Id

	req2 := pb.GetAnythingRequest{
		Id: id,
	}
	res2, err := c.GetAnything(ctx, &req2)
	if err != nil {
		log.Fatalf("GetAnything failed: %v", err)
	}
	log.Printf("GetAnything result: <%+v>\n\n", res2)

	req3 := pb.UpdateAnythingRequest{
		Anything: &pb.Anything{
			Id:          res2.Anything.Id,
			Name:        res2.Anything.Name,
			Description: res2.Anything.Description + " + updated",
			UpdatedAt:   res2.Anything.UpdatedAt,
		},
	}
	res3, err := c.UpdateAnything(ctx, &req3)
	if err != nil {
		log.Fatalf("UpdateAnything failed: %v", err)
	}
	log.Printf("UpdateAnything result: <%+v>\n\n", res3)

	res4, err := c.ListAnything(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("ListAnything failed: %v", err)
	}
	log.Printf("ListAnything result: <%+v>\n\n", res4)

	req5 := pb.DeleteAnythingRequest{
		Id: id,
	}
	res5, err := c.DeleteAnything(ctx, &req5)
	if err != nil {
		log.Fatalf("DeleteAnything failed: %v", err)
	}
	log.Printf("DeleteAnything result: <%+v>\n\n", res5)

	conn2, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()

	c2 := pb.NewUserServiceClient(conn2)

	req6 := pb.LoginRequest{
		Username: "admin",
		Password: "admin",
	}
	res6, err := c2.Login(ctx, &req6)
	if err != nil {
		log.Fatalf("Login failed : %v", err)
	}
	log.Printf("Login result: <%+v>\n\n", res6)

	res7, err := c2.Me(ctx, new(empty.Empty))
	if err != nil {
		log.Fatalf("Me failed : %v", err)
	}
	log.Printf("Me result: <%+v>\n\n", res7)
}
