package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/plummertr/shippy/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*consignment.Consignment, error) {
	var consignment *consignment.Consignment

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := consignment.NewShippingServiceClient(conn)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Create: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &consignment.GetRequest)
	if err != nil {
		log.Fatalf("could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}