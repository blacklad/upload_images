package conf

import (
	"fmt"
	"testing"
)

func TestGetConf(t *testing.T) {
	config, err := GetConf("")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(config)
}