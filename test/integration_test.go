package integeration

import (
	"context"
	"log"
	"testing"

	"github.com/huyntsgs/grpc-bi-directional-stream/mathclient"
)

func Test_SendSimpleSlice(t *testing.T) {

	testTable := [][]int32{
		[]int32{1, 5, 3, 6, 2, 20}, []int32{1, 5, 6, 20},
		[]int32{10, 9, 8, 7, 2, 1}, []int32{10},
		[]int32{1, 1, 2, 2, 2, 30}, []int32{1, 2, 30},
		[]int32{100, 5, 9, 111, 202, 2020}, []int32{100, 111, 202, 2020},
	}
	mathClient := mathclient.NewMathClient()
	mathClient.Connect()

	for i := 0; i < len(testTable)/2; i++ {
		req := testTable[2*i]
		res := testTable[2*i+1]
		ret := make([]int32, 0)
		done := make(chan struct{})

		stream, err := mathClient.Client.FindMaxNumber(context.Background())
		if err != nil {
			log.Fatalf("Can not call FindMaxNumber service: %v", err)
		}

		// Call goroutine for send and receive
		log.Println("Request stream", req)
		log.Println("Expected response ", res)
		go mathClient.SendWorkerStream(stream, req)
		go func() {
			ret, _ = mathClient.ReceiveStream(stream)
			close(done)
		}()
		<-done

		log.Println("Received response ", ret)
		if len(ret) != len(res) {
			t.Fail()
		}
		for i, _ := range ret {
			if ret[i] != res[i] {
				t.Fail()
				break
			}
		}
	}
}

func Test_SendWithWorkerLen100(t *testing.T) {
	arrLen, maxValue := 100, 2100777888
	done := make(chan struct{})
	ret := make([]int32, 0)

	mathClient := mathclient.NewMathClient()
	mathClient.Connect()

	stream, err := mathClient.Client.FindMaxNumber(context.Background())
	if err != nil {
		log.Fatalf("Can not call FindMaxNumber service: %v", err)
	}

	// Generate slice of integer
	req, res := mathclient.GenIntSlice(arrLen, maxValue)
	log.Println("Request stream", req)

	// Call goroutine for send and receive
	go mathClient.SendWorkerStream(stream, req)
	go func() {
		ret, _ = mathClient.ReceiveStream(stream)
		close(done)
	}()
	<-done
	log.Println("Received response ", ret)
	log.Println("Expected response ", res)

	if len(ret) != len(res) {
		t.Fail()
	}
	for i, _ := range ret {
		if ret[i] != res[i] {
			t.Fail()
			break
		}
	}
}
