package test

import (
	"encoding/hex"
	"example/test/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProtobuf(t *testing.T) {
	user := pb.User{
		Id:       0x1234,
		Username: "john",
	}

	bytes, err := proto.Marshal(&user)
	assert.Nil(t, err)

	t.Log(hex.EncodeToString(bytes))

	var u pb.User
	err = proto.Unmarshal(bytes, &u)
	assert.Nil(t, err)

	assert.Equal(t, user.GetId(), u.GetId())
	assert.Equal(t, user.GetUsername(), u.GetUsername())
	assert.Equal(t, user.GetPassword(), u.GetPassword())
}
