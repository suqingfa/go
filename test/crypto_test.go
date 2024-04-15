package test

import (
	"crypto/aes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSHA(t *testing.T) {
	data := []byte("hello")

	sha := sha256.New()
	sha.Write(data)
	sha1 := fmt.Sprintf("%x", sha.Sum(nil))

	sha2 := fmt.Sprintf("%x", sha256.Sum256(data))
	t.Log(sha2)

	assert.Equal(t, sha1, sha2)
}

func TestAes(t *testing.T) {
	key := make([]byte, 16)
	_, _ = rand.Read(key)
	t.Log("key\t\t", hex.EncodeToString(key))

	cipher, err := aes.NewCipher(key)
	assert.Nil(t, err)

	src := []byte("0123456789abcdef")

	encrypt := make([]byte, 16)
	cipher.Encrypt(encrypt, src)
	t.Log("encrypt\t", hex.EncodeToString(encrypt))

	decrypt := make([]byte, 16)
	cipher.Decrypt(decrypt, encrypt)
	t.Log("decrypt\t", hex.EncodeToString(decrypt))
	t.Log("string\t", string(decrypt))
}

func TestPublicKey(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	assert.Nil(t, err)

	t.Log("public key", base64.StdEncoding.EncodeToString(publicKey))
	t.Log("private key", base64.StdEncoding.EncodeToString(privateKey))

	message := []byte("hello")
	sign := ed25519.Sign(privateKey, message)
	t.Log("sign", hex.EncodeToString(sign))

	verify := ed25519.Verify(publicKey, message, sign)
	assert.True(t, verify)
}

func TestX509(t *testing.T) {
	bytes, err := os.ReadFile("/etc/ssl/certs/ca-certificates.crt")
	assert.Nil(t, err)

	for block, rest := pem.Decode(bytes); block != nil && block.Type == "CERTIFICATE"; block, rest = pem.Decode(rest) {
		certificate, err := x509.ParseCertificate(block.Bytes)
		assert.Nil(t, err)
		assert.NotNil(t, certificate)

		t.Log(certificate.Subject)
	}
}
