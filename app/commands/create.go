package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

func CreateVote(client *model.Client4, conn *tarantool.Connection, creator, channelID string, args []string) {
	log.Printf("[DEBUG] Создание голосования: Creator=%s, Args=%v", creator, args)

	if len(args) < 2 {
		sendError(client, channelID, "Формат: vote create \"Вопрос?\" \"Опция1\" \"Опция2\"")
		return
	}

	options := map[string]int{}
	for _, opt := range args[1:] {
		options[opt] = 0
	}

	ID := uuid.New().String()
	_, err := conn.Insert("votes", []interface{}{
		ID,
		creator,
		strings.Trim(args[0], `"`),
		options,
		map[string]string{},
		"active",
	})

	if err != nil {
		log.Printf("[ERROR] Ошибка создания: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	}

	sendSuccess(client, channelID, fmt.Sprintf("✅ Голосование создано! ID: `%s`", ID))
}
