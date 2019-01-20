package cryptoutil

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"log"
	"math/big"
)

// GenerateKeyPair generates private and public ecdsa keypair.
func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	ec := elliptic.P256()
	priv, err := ecdsa.GenerateKey(ec, rand.Reader)
	if err != nil {
		log.Printf("Error generate key %v", err)
		return nil, nil, err
	}
	pub := &priv.PublicKey
	return priv, pub, nil
}

// Sign hashes data and sign. Returns byte slice of r and s.
func Sign(priv *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	hash := Hash256(data)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	sig := append(r.Bytes(), s.Bytes()...)
	return sig, err
}

// Verify verifies data base on signature. The hashData is already hashed.
func Verify(pub *ecdsa.PublicKey, hashData []byte, signature []byte) bool {
	r := new(big.Int).SetBytes(signature[0 : len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])
	return ecdsa.Verify(pub, hashData, r, s)
}

// ParsePublicKey returns new instance of public key from bytes slice
func ParsePublicKey(data []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), data)
	if x == nil {
		return nil, errors.New("Invalid public key")
	}
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, nil
}

// SerializePublicKey converts public key to bytes slice
func SerializePublicKey(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(elliptic.P256(), pub.X, pub.Y)
}

// paddingZero appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddingZero(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}

// NewPrivateKey instantiates a new private key from a scalar encoded as a
// big integer.
func NewPrivateKey(d *big.Int) *ecdsa.PrivateKey {
	b := make([]byte, 0, 32)
	dB := paddingZero(32, b, d.Bytes())
	priv, _ := PrivKeyFromBytes(dB)
	return priv
}

// PrivKeyFromBytes returns a private/public key for `curve' based on the
// D field of private key in form bytes slice.
func PrivKeyFromBytes(pk []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult(pk)
	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(pk),
	}
	return priv, &priv.PublicKey
}

// SerializePrivateKey converts private key to bytes slice
func SerializePrivateKey(p *ecdsa.PrivateKey) []byte {
	return p.D.Bytes()
}

// IntToBytes converts integer number to bytes slice
func IntToBytes(i int32) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, i)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Hash256 performs 256 bits hash on bytes slice.
func Hash256(data []byte) []byte {
	checksum := sha256.Sum256(data)
	return checksum[:32]
}
