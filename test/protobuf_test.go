package test

import (
	"encoding/hex"
	"example/test/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProtobuf(t *testing.T) {
	user := pb.User{
		Id:       0x1234,
		Username: "john",
	}

	bytes, err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(bytes))

	var u pb.User
	err = proto.Unmarshal(bytes, &u)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.GetId(), u.GetUsername(), u.GetPassword())
}
