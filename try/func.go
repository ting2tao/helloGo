package main

import (
	"encoding/json"
	"fmt"
)

type Ino struct {
	RequestInfos  string `json:"request_infos"`
	RequestMethod string
}

func main() {
	fmt.Println(InnerFunc())
}

func InnerFunc() error {
	ino := Ino{
		RequestInfos:  "",
		RequestMethod: "AsyncKsbActivityOrder",
	}
	var err error
	check := func(p interface{}) error {
		err = json.Unmarshal([]byte(ino.RequestInfos), &p)
		return err
	}
	switch ino.RequestMethod {
	case "AsyncKsbActivityOrder":
		var params string
		err = check(params)
	}
	return err
}
