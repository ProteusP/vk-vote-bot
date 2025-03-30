package commands

import (
	"log"
	"vk-vote-bot/tarantooldb"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

func Vote(client *model.Client4, conn *tarantool.Connection, userID, channelID string, args []string) {
	log.Printf("[DEBUG] Голосование: User=%s, Args=%v", userID, args)

	if len(args) < 2 {
		sendError(client, channelID, "Формат: vote vote [ID] [вариант]")
		return
	}

	voteID := args[0]
	option := args[1]
	log.Printf("[DEBUG] Опция: %s", option)

	resp, err := conn.Select("votes", "primary", 0, 1, tarantool.IterEq, []interface{}{voteID})

	if err != nil {
		log.Printf("[ERROR] Ошибка поиска: %v", err)
		sendError(client, channelID, "Ошибка сервера")
		return
	}
	if len(resp.Data) == 0 {
		sendError(client, channelID, "Голосование не найдено")
		return
	}

	var vote tarantooldb.Vote
	vote.LoadFromResponse(resp.Data)

	if val, exists := vote.Options[option]; exists {
		vote.Options[option] = val + 1
	} else {
		log.Printf("[ERROR Опция %s не существует в голосовании: %v", option, vote)
		sendError(client, channelID, "Такой опции нет в этом голосовании!")
		return
	}

	_, err = conn.Replace("votes", []interface{}{
		vote.ID,
		vote.Creator,
		vote.Question,
		vote.Options,
		vote.Votes,
		vote.Status,
	})

	SendMessage(client, channelID, "✅ Ваш голос учтен!")
}
