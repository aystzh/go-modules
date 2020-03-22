package main

import (
	"github.com/aystzh/go-modules/protobuf"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := &aystzh_protobuf_student.Student{
		Name:   "zhangsan",
		Male:   false,
		Scores: []int32{98, 100, 22},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	newTest := &aystzh_protobuf_student.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}

	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
