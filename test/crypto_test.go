package test

import (
	"crypto/aes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSHA(t *testing.T) {
	data := []byte("hello")

	sha := sha256.New()
	sha.Write(data)
	sha1 := fmt.Sprintf("%x", sha.Sum(nil))
	println(sha1)

	sha2 := fmt.Sprintf("%x", sha256.Sum256(data))
	println(sha2)

	if sha1 != sha2 {
		t.Error()
	}
}

func TestAes(t *testing.T) {
	key := make([]byte, 16)
	_, _ = rand.Read(key)
	println("key\t\t", hex.EncodeToString(key))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		t.Error(err)
	}

	src := []byte("0123456789abcdef")

	encrypt := make([]byte, 16)
	cipher.Encrypt(encrypt, src)
	println("encrypt\t", hex.EncodeToString(encrypt))

	decrypt := make([]byte, 16)
	cipher.Decrypt(decrypt, encrypt)
	println("decrypt\t", hex.EncodeToString(decrypt))
	println("string\t", string(decrypt))
}

func TestPublicKey(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Error(err)
	}

	println("public key", base64.StdEncoding.EncodeToString(publicKey))
	println("private key", base64.StdEncoding.EncodeToString(privateKey))

	message := []byte("hello")
	sign := ed25519.Sign(privateKey, message)
	println("sign", hex.EncodeToString(sign))

	verify := ed25519.Verify(publicKey, message, sign)
	if !verify {
		t.Error()
	}
}
