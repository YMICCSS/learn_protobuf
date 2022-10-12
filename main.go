package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protobuf/proto/userService"
)

func main(){

	fmt.Println("aaa")
	u :=&userService.Userinfo{
		Username: "張三",
		Age: 22,
		Hobby: []string{"吃飯","睡覺","寫程式"},
	}
	fmt.Println(u)
	fmt.Println(u.GetUsername())
	// proto的序列化
	data,_ :=proto.Marshal(u)
	fmt.Println(data)
	// proto的反序列化
	user := &userService.Userinfo{}
	proto.Unmarshal(data,user)
	fmt.Println(user)

}
