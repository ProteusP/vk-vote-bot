package commands

import (
	"log"
	"vk-vote-bot/tarantooldb"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

const StatusIndex = 5

func EndVote(client *model.Client4, conn *tarantool.Connection, userID, channelID string, args []string) {
	if len(args) == 0 {
		sendError(client, channelID, "Формат: vote end [ID]")
		return
	}

	voteID := args[0]
	var vote tarantooldb.Vote

	vote.Status = "ended"

	_, err := conn.Update("votes", "primary", []any{voteID},
		[]tarantool.Op{
			tarantool.Op{
				Field: StatusIndex,
				Op:    "=",
				Arg:   vote.Status,
			},
		},
	)

	if err != nil {
		log.Printf("[ERROR] Ошибка обновления: %v", err)
		sendError(client, channelID, "Ошибка при завершении голосования")
		return
	}

	//TODO: FORMAT
	sendSuccess(client, channelID, "✅ Голосование завершено!")
}
