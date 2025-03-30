package commands

import (
	"fmt"
	"log"
	"strings"
	"vk-vote-bot/tarantooldb"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

func ShowResults(client *model.Client4, conn *tarantool.Connection, channelID string, args []string) {
	if len(args) == 0 {
		sendError(client, channelID, "Формат: vote results [ID]")
		return
	}

	voteID := args[0]
	log.Print("[INFO] Айди из запроса:", voteID)

	resp, err := conn.Select("votes", "primary", 0, 1, tarantool.IterEq, []interface{}{voteID})

	switch {
	case err != nil:
		log.Printf("[ERROR] Ошибка поиска голосования: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	case len(resp.Data) == 0:
		sendError(client, channelID, "Голосование не найдено")
		return

	}

	log.Printf("[INFO] Структура ответа: %v", resp.Data[0])

	var vote tarantooldb.Vote
	err = vote.LoadFromResponse(resp.Data)
	if err != nil {
		log.Printf("[ERROR] Ошибка загрузки голосования: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	}
	log.Printf("[INFO] Новое голосование: %v", vote)

	var result strings.Builder
	result.WriteString("📊 Результаты:\n")

	for option := range vote.Options {
		count := vote.Options[option]

		result.WriteString(fmt.Sprintf("- %s: %d\n", option, count))
	}

	SendMessage(client, channelID, result.String())
}
