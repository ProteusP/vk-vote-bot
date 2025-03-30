package tarantooldb

import (
	"log"
	"os"
	"time"

	"github.com/tarantool/go-tarantool"
)

func Connect(addr string) (*tarantool.Connection, error) {
	conn, err := tarantool.Connect(addr, tarantool.Opts{
		User:          os.Getenv("DB_USER"),
		Pass:          os.Getenv("DB_PASSWORD"),
		Reconnect:     2 * time.Second,
		MaxReconnects: 3,
	})

	if err != nil {
		log.Fatalf("[ERROR] Ошибка подключения к Tarantool: %v", err)
	}

	return conn, err
}
