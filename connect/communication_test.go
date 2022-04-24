package connect

import (
	"log"
	"testing"
)

var attr = Attributes{"192.168.23.130", "22", "root", "111111", true}

func TestConnect_test(t *testing.T) {
	err := Connect_test(attr)
	if err != nil {
		log.Println(err)
	}
}

func TestSsh_connect(t *testing.T) {
	_, err := Ssh_connect(attr)
	if err != nil {
		log.Println(err)
	}
}
