package main

import (
	"fmt"
	pb "github.com/chenmingyong0423/blog/tutorial-code/go/protobuf/proto/user"
	"os"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// 写入消息
	user := pb.User{
		Id:   1,
		Name: "陈明勇",
		Age:  18,
		Phones: []*pb.User_PhoneNumber{
			{
				Number: "18888888888",
				Type:   pb.PhoneType_PHONE_TYPE_MOBILE,
			},
			{
				Number: "12345678901",
				Type:   pb.PhoneType_PHONE_TYPE_WORK,
			},
		},
		Birth: timestamppb.New(time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)),
	}
	out, err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("user.bin", out, 0644)
	if err != nil {
		panic(err)
	}

	// 读取消息
	in, err := os.ReadFile("user.bin")
	if err != nil {
		panic(err)
	}
	user2 := &pb.User{}
	err = proto.Unmarshal(in, user2)
	if err != nil {
		panic(err)
	}
	// id:1 name:"陈明勇" age:18 phones:{number:"18888888888"} phones:{number:"12345678901" type:PHONE_TYPE_WORK} birth:{seconds:915148800}
	fmt.Println(user2)
}
