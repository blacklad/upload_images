package main

import (
	"os/user"
	"testing"
)

func Test(t *testing.T) {
	//println(time.Now().Format("20060102150405"))
	//println(path.Join("aa", "bb", ".png"))

	//s := make([]string, 10)
	//s = append(s, "aaa")
	//fmt.Println(s)

	//config, err := conf.GetConf("~/.upload.conf")
	//if err != nil {
	//	println(err)
	//}
	//println(config)


	user, err := user.Current()
	if nil == err {
		println(user.HomeDir)
	}
}