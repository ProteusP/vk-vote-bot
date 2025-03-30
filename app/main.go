package main

import (
	"log"
	"time"
	"vk-vote-bot/config"
	"vk-vote-bot/handlers"
	"vk-vote-bot/tarantooldb"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[INFO] Запуск бота...")

	cfg := config.Load()
	cfg.MattermostURL = "http://localhost:8065"
	cfg.TarantoolAddr = "127.0.0.1:3301"
	cfg.MattermostToken = "m9xuf91pi7bu7buzd7e13i8wjc"
	cfg.WebhookToken = "18q9sia4r3rbpxt5rirm6n8h7y"
	log.Printf("[DEBUG] Конфигурация: %+v", cfg)
	client := model.NewAPIv4Client(cfg.MattermostURL)
	client.SetToken(cfg.MattermostToken)

	u, _, _ := client.GetMe("")

	log.Printf("[INFO] Айди клиента: %s", u.Id)
	log.Printf("[INFO] Имя клиента: %s", u.Username)

	team, _, _ := client.GetTeamByName("testing", "")
	channels, _, _ := client.GetChannelsForTeamForUser(team.Id, u.Id, false, "")
	for _, ch := range channels {
		log.Printf("[INFO] Название канала, где я есть: %s, а ID: %s", ch.Name, ch.Id)

		mem, _, _ := client.GetChannelMember(ch.Id, u.Id, "")
		log.Printf("[INFO] И права в нем: %s", mem.Roles)
	}
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
	data, _ := msgpack.Marshal(tarantooldb.Vote{})
	log.Printf("[DEBUG] Msgpack marhs: %v", data)
	resp, _ := tarantoolConn.Select("votes", "primary", 0, 1, tarantool.IterEq, []interface{}{"b84869e4-ec7a-4284-87a3-1ecbf3b39382"})
	log.Printf("[DEBUG] Ответ от Tarantool: %v", resp)
	handler := handlers.NewHandler(client, tarantoolConn, cfg)
	handler.Start()
}
