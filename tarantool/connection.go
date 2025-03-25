package tarantool

import (
	"log"
	"github.com/tarantool/go-tarantool"
)

func Connect(addr string) *tarantool.Connection {
	conn, err := tarantool.Connect(addr, tarantool.Opts{
		User: "user",
		Pass: "password",
	})

	if err != nil {
		log.Fatal("Ошибка подключения к Tarantool: ", err)
	}

	return conn
}
