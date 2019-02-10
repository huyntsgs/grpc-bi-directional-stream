package main

import (
	"crypto/ecdsa"
	"io"
	"log"

	"github.com/huyntsgs/grpc-bi-directional-stream/api"
	"github.com/huyntsgs/grpc-bi-directional-stream/cryptoutil"
)

type MathService struct {
}

func (ms *MathService) FindMaxNumber(stream protobuf.Math_FindMaxNumberServer) error {
	var max int32 = -1
	reqs := make([]int32, 0)
	res := make([]int32, 0)
	notVerified := 0
	var pub *ecdsa.PublicKey

	ctx := stream.Context()
	for {

		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received request:", reqs)
			log.Println("Send client response:", res)
			log.Println("Unverified Numbers:", notVerified)
			return nil
		}
		if err != nil {
			return err
		}

		// Verify the signature in stream
		data, err := cryptoutil.IntToBytes(req.Number)
		if err != nil {
			log.Printf("Can not convert number %d to byte %v", req.Number, err)
			continue
		}
		// Parse public only at first time.
		if max == -1 {
			pub, err = cryptoutil.ParsePublicKey(req.PublicKey)
			if err != nil {
				log.Printf("Can not unmarshal public key: %v", err)
				return err
			}
		}

		reqs = append(reqs, req.Number)
		hash := cryptoutil.Hash256(data)
		isVerified := cryptoutil.Verify(pub, hash, req.R, req.S)

		if isVerified {
			if max == -1 {
				max = req.Number
				stream.Send(&protobuf.Response{max})
				res = append(res, req.Number)
			} else if max < req.Number {
				max = req.Number
				stream.Send(&protobuf.Response{max})
				res = append(res, req.Number)
			}
		} else {
			notVerified++
			log.Printf("Number is not signed: %d-%x\n", req.Number, data)
			log.Printf("\nSig %x. \nPub %x\n", req.Sig, req.PublicKey)
		}
	}
	return nil
}
