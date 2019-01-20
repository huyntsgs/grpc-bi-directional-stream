package mathclient

import (
	"crypto/ecdsa"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/huyntsgs/grpc-bi-directional-stream/api"
	"github.com/huyntsgs/grpc-bi-directional-stream/cryptoutil"
	"google.golang.org/grpc"
)

type (
	MathClient struct {
		PrivKey *ecdsa.PrivateKey
		PubKey  *ecdsa.PublicKey
		Client  protobuf.MathClient
	}
	Task struct {
		Id     int
		Number int32
	}
)

func NewMathClient() *MathClient {
	priv, pub, err := cryptoutil.GenerateKeyPair()
	if err != nil {
		return nil
	}
	privByte := cryptoutil.SerializePrivateKey(priv)
	pubByte := cryptoutil.SerializePublicKey(pub)

	log.Printf("priv key: %x", privByte)
	log.Printf("pub key: %x", pubByte)
	return &MathClient{PrivKey: priv, PubKey: pub}
}
func (mathClient *MathClient) Connect() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can not connect grpc server: %v", err)
	}
	mathClient.Client = protobuf.NewMathClient(conn)
}

// SendStream signs the number in slice and sending request data including number,
// signature and public key to server.
func (mathClient *MathClient) SendStream(stream protobuf.Math_FindMaxNumberClient, arrs []int32) error {
	pub := cryptoutil.SerializePublicKey(mathClient.PubKey)
	sig := make([]byte, 0)
	data := make([]byte, 0)
	var err error
	for _, i := range arrs {
		data, err = cryptoutil.IntToBytes(i)
		if err != nil {
			log.Printf("Can not convert int to byte: %v", err)
			return err
		}

		sig, err = cryptoutil.Sign(mathClient.PrivKey, data)
		if err != nil {
			log.Printf("Can not sign data: %v", err)
			return err
		}
		time.Sleep(50 * time.Millisecond)
		err = stream.Send(&protobuf.Request{Number: i, Sig: sig, PublicKey: pub})

		if err != nil {
			log.Printf("Can not send stream %v", err)
			return err
		}
		//log.Printf("%d-%x \nSig: %x.\nPubkey: %x\n", i, data, sig, pub)

	}
	if err := stream.CloseSend(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ReceiveStream receives response integer from server.
func (mathClient *MathClient) ReceiveStream(stream protobuf.Math_FindMaxNumberClient) ([]int32, error) {
	ret := make([]int32, 0)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return ret, nil
		}
		if err != nil {
			log.Printf("Can not send stream %v", err)
			return nil, err
		}

		ret = append(ret, res.Res)
	}
	return ret, nil
}

// GenIntSliceRnd generates integer slice with slice length
// and integer value are randomized within arrLen and maxValue.
func GenIntSliceRnd(arrLen, maxValue int) ([]int32, []int32) {
	req := make([]int32, 0)
	res := make([]int32, 0)

	arrayLength := rand.Intn(arrLen)
	if arrayLength == 0 {
		arrayLength = rand.Intn(arrLen)
	}
	var max int32 = 0
	for i := 0; i < arrayLength; i++ {
		ai := int32(rand.Intn(maxValue))
		req = append(req, ai)
		if i == 0 {
			max = ai
			res = append(res, ai)
		} else if ai > max {
			max = ai
			res = append(res, ai)
		}
	}
	return req, res
}

// GenIntSlice generates integer slice with the given arrLen is slice length and
// integer number is randomized within maxValue.
func GenIntSlice(arrLen, maxValue int) ([]int32, []int32) {
	req := make([]int32, 0)
	res := make([]int32, 0)

	var max int32 = -1
	for i := 0; i < arrLen; i++ {
		ai := int32(rand.Intn(maxValue))
		req = append(req, ai)
		if max == -1 {
			max = ai
			res = append(res, ai)
		} else if ai > max {
			max = ai
			res = append(res, ai)
		}
	}
	return req, res
}

// SendWorkerStream signs and send stream integer to server using worker pool.
//
func (mathClient *MathClient) SendWorkerStream(stream protobuf.Math_FindMaxNumberClient, arr []int32) error {
	reqs, err := mathClient.SignSlice(arr)
	if err != nil {
		log.Printf("Can not sign array %v", err)
		return err
	}

	for _, req := range reqs {
		err = stream.Send(req)
		time.Sleep(4 * time.Millisecond)
		if err != nil {
			log.Printf("Can not send stream %v", err)
			return err
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SignSlice signs slice integers using worker pool.
func (mathClient *MathClient) SignSlice(numbers []int32) ([]*protobuf.Request, error) {
	reqs := make([]*protobuf.Request, len(numbers))
	pub := cryptoutil.SerializePublicKey(mathClient.PubKey)

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	task := make(chan Task)
	nworker := 10
	go mathClient.pool(task, &wg, nworker, reqs, pub)
	for i, number := range numbers {
		task <- Task{Id: i, Number: number}
	}
	close(task)
	wg.Wait()

	return reqs, nil
}

// worker receives task from pool and performs signing the numbers.
// Add signed result to requests slice.
func (mathClient *MathClient) worker(task chan Task, wg *sync.WaitGroup, reqs []*protobuf.Request, pub []byte) {

	for {
		task, ok := <-task
		if !ok {
			return
		}
		b, err := cryptoutil.IntToBytes(task.Number)
		if err != nil {
			return
		}

		sig, err := cryptoutil.Sign(mathClient.PrivKey, b)
		if err != nil {
			log.Printf("Can not sign number %d: %v", task.Number, err)
			return
		}
		req := &protobuf.Request{Number: task.Number, Sig: sig, PublicKey: pub}
		reqs[task.Id] = req
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}
}

// pool receives task and delivers to workers.
func (mathClient *MathClient) pool(task chan Task, wg *sync.WaitGroup, nworker int, reqs []*protobuf.Request, pub []byte) {
	for i := 0; i < nworker; i++ {
		go mathClient.worker(task, wg, reqs, pub)
	}
}
