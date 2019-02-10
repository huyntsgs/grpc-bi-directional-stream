## Golang GRPC bi-directional stream

Implement a FindMaxNumber RPC Bi-Directional Streaming Client and Server
system using Go GRPC framework as following.

- The function takes a stream of Request message that has one integer, and
returns a stream of Responses that represent the current maximum between
all these integers.

- Client will be having a cryptographic public key and client will be identified
using his private key. Client will sign every request message in the stream.

- Each requested message should be verified against the signature at the
server end. Only those numbers will be considered to be processed whose
sign is successfully verified.

Example: The client will send a stream of number (1,5,3,6,2,20) and each number
will be signed by the private key of the client and the server will respond with a
stream of numbers (1,5,6,20).

    =======================

Firstly needs to define the protobuf file. We also use ecdsa to generate keypair for client. For cryptography function, we create utility functions in cryptoutil package.

Protobuf file includes two messages Request, Response and one rpc FindMaxNumber takes input stream is Request and returns Response stream.

```
service Math {
        rpc FindMaxNumber(stream Request) returns (stream Response);
}
message Request {
        int32 number = 1;
        bytes sig = 2;
        bytes publicKey = 3;
}
message Response {
        int32 res = 1;
}
```

For inconvenient, we create function for generating slice of integers for sending.

Client starts connecting to server and get rpc stream. Using rpc stream to send Request message to server. Request message includes integer number, signature of number and public key of client in bytes.

After sending is complete, we need to call CloseSend of stream to terminate sending process. 
Client also having Receive stream function, which receiving response from server. The sending and receiving stream runs on two goroutines.

On server side, our grpc server listens on 8888 port. We define FindMaxNumber for listening incoming integers from clients. Server receives Request message, verifies the signature by public key with help of cryptoutil. Server maintains a max variable to check with verified integer. If the received integer is greater than max then updates max and sends the integer back to client.


# Source code

- api:
Contains protobuf file and command line for generate pb.go file.

- client:
The main client entry for starting send/receive stream.  

- cryptoutil:
Functions for generate keypair, sign, verify and hash.

- mathclient:
 Contains struct and function to execute send/receive stream. Each mathclient acts as one client instance.

- server:
Contains grpc server for listening clients connection, stream function for receive/send data.

- test:
Includes some integration and load test cases.

I just fixed bug related to sign and verify function. Actually, size of signature is not always 64 bytes. It depends on size of r and s. With the number signature, there are many cases are 63 bytes.


