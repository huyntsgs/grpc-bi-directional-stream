package integeration

import (
	"context"
	"log"

	"sync"
	"sync/atomic"
	"testing"

	"github.com/huyntsgs/grpc-bi-directional-stream/mathclient"
)

func Test_SendWithWorkerLen1000(t *testing.T) {
	arrLen, maxValue := 1000, 2100777888
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
		}
	}

}

func Test_SendLen1000(t *testing.T) {
	arrLen, maxValue := 1000, 100000
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
	go mathClient.SendStream(stream, req)
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
		}
	}
}
func Test_SendLen500(t *testing.T) {
	arrLen, maxValue := 500, 100000
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
	go mathClient.SendStream(stream, req)
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
		}
	}
}

func Test_MultiClients(t *testing.T) {
	var wg sync.WaitGroup
	var cnt uint32
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			arrLen, maxValue := 1000, 100000
			done := make(chan struct{})
			ret := make([]int32, 0)

			mathClient := mathclient.NewMathClient()
			mathClient.Connect()

			stream, err := mathClient.Client.FindMaxNumber(context.Background())
			if err != nil {
				log.Fatalf("Can not call FindMaxNumber service: %v", err)
			}

			// Will do send and receive stream
			req, res := mathclient.GenIntSlice(arrLen, maxValue)
			log.Println("Request stream", req)

			go mathClient.SendWorkerStream(stream, req)
			go func() {
				ret, _ = mathClient.ReceiveStream(stream)
				close(done)
			}()
			<-done
			log.Println("Received stream ", ret)
			log.Println("Expected response ", res)

			if len(ret) != len(res) {
				atomic.AddUint32(&cnt, 1)
			} else {
			LOOP:
				for i, _ := range ret {
					if ret[i] != res[i] {
						atomic.AddUint32(&cnt, 1)
						break LOOP
					}
				}
			}

		}()
	}
	wg.Wait()
	if cnt > 0 {
		log.Println("Number fails ", cnt)
		t.Fail()
	}
}
