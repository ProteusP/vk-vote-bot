package commands

import (
	"log"
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
	log.Printf("[DEBUG] Айди из запроса: %s", voteID)

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

	log.Printf("[DEBUG] Структура ответа: %v", resp.Data[0])

	var vote tarantooldb.Vote
	err = vote.LoadFromResponse(resp.Data)
	if err != nil {
		log.Printf("[ERROR] Ошибка загрузки голосования: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	}

	SendMessage(client, channelID, vote.Results())
}
