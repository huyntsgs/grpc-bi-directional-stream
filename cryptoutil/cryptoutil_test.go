package cryptoutil

import (
	"testing"
)

func TestSignVerifyInts(t *testing.T) {
	numbers := []int32{2070362201, 19934343, 0, 1986649646, 2147483647, -1767797977}
	priv, pub, _ := GenerateKeyPair()
	for _, i := range numbers {
		data, _ := IntToBytes(i)
		r, s, _ := Sign(priv, data)
		hash := Hash256(data)
		isVerifed := Verify(pub, hash, r, s)
		if !isVerifed {
			t.Fail()
		}
	}
}

func TestSignVerifyString(t *testing.T) {
	dataTest := []string{"jjdfqpeifnffndafa^989r**&^%$$$df", " ", "0", "557067bbf4c9991cd9206aee39de1a90c3329beb35f7749dde6fd0cd4ceb5abc19e589b4c104d053abe0a7639a0ab12bbbb970aad66a91d2e9e8878fb5f236"}
	priv, pub, _ := GenerateKeyPair()
	for _, s := range dataTest {
		data := []byte(s)
		r, s, _ := Sign(priv, data)

		hash := Hash256(data)
		isVerifed := Verify(pub, hash, r, s)
		if !isVerifed {
			t.Fail()
		}
	}
}

func TestWrongSig(t *testing.T) {
	dataTest := []string{"jjdfqpeifnffndafadf", " ", "0", "557067bbf4c9991cd9206aee39de1a90c3329beb35f7749dde6fd0cd4ceb5abc19e589b4c104d053abe0a7639a0ab12bbbb970aad66a91d2e9e8878fb5f236"}
	priv, pub, _ := GenerateKeyPair()
	for _, s := range dataTest {
		data := []byte(s)
		r, s, _ := Sign(priv, data)

		hash := Hash256(data)
		if hash[0] != 0x00 {
			hash[0] = 0x00
		} else {
			hash[0] = 0xff
		}
		if hash[5] != 0x00 {
			hash[5] = 0x00
		} else {
			hash[5] = 0xff
		}
		if hash[15] != 0x00 {
			hash[15] = 0x00
		} else {
			hash[15] = 0xff
		}
		isVerifed := Verify(pub, hash, r, s)
		if isVerifed {
			t.Fail()
		}
	}
}
