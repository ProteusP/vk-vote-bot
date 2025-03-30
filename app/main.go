package main

import (
	"log"
	"time"
	"vk-vote-bot/config"
	"vk-vote-bot/handlers"
	"vk-vote-bot/tarantooldb"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[INFO] Запуск бота...")

	cfg := config.Load()
	client := model.NewAPIv4Client(cfg.MattermostURL)
	client.SetToken(cfg.MattermostToken)

	u, _, _ := client.GetMe("")

	log.Printf("[INFO] Айди бота: %s", u.Id)
	log.Printf("[INFO] Имя бота: %s", u.Username)

	var tarantoolConn *tarantool.Connection
	for i := 0; i < 5; i++ {
		conn, err := tarantooldb.Connect(cfg.TarantoolAddr)
		if err == nil {
			tarantoolConn = conn
			log.Println("[INFO] Подключение к Tarantool успешно установлено")
			break
		}
		log.Printf("[WARN] Попытка подключения %d: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if tarantoolConn == nil {
		log.Fatal("Не удалось подключиться к Tarantool")
	}
	defer tarantoolConn.Close()

	handler := handlers.NewHandler(client, tarantoolConn, cfg)
	handler.Start()
}
