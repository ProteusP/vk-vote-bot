package commands

import (
	"log"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

func DeleteVote(client *model.Client4, conn *tarantool.Connection, userID, channelID string, args []string) {
	if len(args) == 0 {
		sendError(client, channelID, "Формат: vote delete [ID]")
		return
	}

	voteID := args[0]

	_, err := conn.Delete("votes", "primary", []any{voteID})
	if err != nil {
		log.Printf("[ERROR] Ошибка удаления: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	}

	sendSuccess(client, channelID, "✅ Голосование удалено!")
}
