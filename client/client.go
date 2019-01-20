package main

import (
	"context"
	"flag"
	"log"

	//"runtime/debug"

	"github.com/huyntsgs/grpc-bi-directional-stream/mathclient"
)

func main() {

	//debug.SetGCPercent(10)
	// Using two parameters -len is number of integers will be sent to server
	// and -max value of these integers.
	arrLen := flag.Int("len", 1000, "Max int number will be send by stream")
	maxValue := flag.Int("max", 2011888999, "Max value of numbers")
	flag.Parse()

	done := make(chan struct{})
	ret := make([]int32, 0)

	mathClient := mathclient.NewMathClient()
	mathClient.Connect()

	stream, err := mathClient.Client.FindMaxNumber(context.Background())
	if err != nil {
		log.Fatalf("Can not call FindMaxNumber rpc: %v", err)
	}

	//ctx := stream.Context()

	// Generate slice integers.
	req, res := mathclient.GenIntSlice(*arrLen, *maxValue)
	log.Println("Request stream", req)

	go mathClient.SendWorkerStream(stream, req)
	go func() {
		ret, _ = mathClient.ReceiveStream(stream)
		close(done)
	}()

	// go func() {
	// 	<-ctx.Done()
	// 	if err := ctx.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// 	close(done)
	// }()
	<-done
	log.Println("Received result:", ret)
	log.Println("Expected response:", res)
}
