package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"vk-vote-bot/commands"
	"vk-vote-bot/config"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/tarantool/go-tarantool"
)

type Handler struct {
	client       *model.Client4
	conn         *tarantool.Connection
	webhookToken string
}

func NewHandler(client *model.Client4, conn *tarantool.Connection, cfg *config.Config) *Handler {
	return &Handler{
		client:       client,
		conn:         conn,
		webhookToken: cfg.WebhookToken,
	}
}

func (h *Handler) Start() {
	http.HandleFunc("/webhook", h.handleWebhook)
	log.Println("🚀 Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}

func (h *Handler) handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Print("[DEBUG] Пришел запрос")
	token := r.URL.Query().Get("token")
	if token != h.webhookToken {
		log.Printf("Неверный токен вебхука: %s", token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("[DEBUG] Токен верный, обработка запроса")

	var payload model.OutgoingWebhookPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Ошибка декодирования: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("[DEBUG] Запрос успешно декодирован")
	log.Println("[DEBUG] Message:", payload.Text)
	if !strings.HasPrefix(payload.Text, "vote") {
		log.Println("[DEBUG] Команда vote не обнаружена")
		return
	}

	log.Println("[DEBUG] Команда vote обнаружена")
	command := strings.TrimSpace(strings.TrimPrefix(payload.Text, "vote"))
	log.Printf("[DEBUG] Обработка команды: UserID=%s, Command=%s", payload.UserId, command)

	h.handleCommand(payload.UserId, payload.ChannelId, command)
	log.Println("[DEBUG] Обработка команды завершена")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleCommand(userID, channelID, command string) {
	args := strings.Fields(command)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "create":
		commands.CreateVote(h.client, h.conn, userID, channelID, args[1:])
	case "vote":
		commands.Vote(h.client, h.conn, userID, channelID, args[1:])
	case "results":
		commands.ShowResults(h.client, h.conn, channelID, args[1:])
	case "end":
		commands.EndVote(h.client, h.conn, userID, channelID, args[1:])
	case "delete":
		commands.DeleteVote(h.client, h.conn, userID, channelID, args[1:])
	default:
		log.Printf("Неизвестная команда: %s", args[0])
		commands.SendMessage(h.client, channelID, fmt.Sprintf("Неизвестная команда: %s, используйте create, vote, results, end или delete", args[0]))
	}
}
