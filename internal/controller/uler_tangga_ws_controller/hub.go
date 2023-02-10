package uler_tangga_ws_controller

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	world_service "github.com/umarkotak/go-uler-tangga-api/internal/service/world"
)

type Hub struct {
	clients       map[*Client]bool
	clientMap     map[string]*Client
	clientRoomMap map[string]map[string]*Client
	broadcast     chan []byte
	register      chan *Client
	unregister    chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		clientMap:     map[string]*Client{},
		clientRoomMap: map[string]map[string]*Client{},
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.clientMap[client.identiy.ID] = client
			_, ok := h.clientRoomMap[client.identiy.RoomID]
			if !ok {
				h.clientRoomMap[client.identiy.RoomID] = map[string]*Client{}
			}
			h.clientRoomMap[client.identiy.RoomID][client.identiy.ID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.justCloseClient(client)
			}
		case message := <-h.broadcast:
			h.messageRouter(message)
		}
	}
}

func (h *Hub) runTicker() {
}

func (h *Hub) justCloseClient(c *Client) {
	messageContract, _ := json.Marshal(model.MessageContract{
		Action: "player_leave_room",
		From:   c.identiy,
	})
	h.messageRouter(messageContract)
	close(c.send)
	delete(h.clients, c)
	delete(h.clientMap, c.identiy.ID)
}

func (h *Hub) messageRouter(rawMessage []byte) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error("Recovered. Error: ", err)
		}
	}()

	messageContract := model.MessageContract{}

	err := json.Unmarshal(rawMessage, &messageContract)
	if err != nil {
		logrus.Error(err)
		return
	}

	var responseContract model.ResponseContract

	switch messageContract.Action {
	case "player_join_room":
		responseContract, err = world_service.JoinRoom(messageContract)
	case "player_roll_number":
		responseContract, err = world_service.RollNumber(messageContract)
	case "player_move":
		responseContract, err = world_service.Move(messageContract)
	case "player_leave_room":
		responseContract, err = world_service.LeaveRoom(messageContract)
	case "player_end_turn":
		responseContract, err = world_service.EndTurn(messageContract)
	case "admin_move_player":
	case "admin_raw_broadcast":
	default:
		logrus.Error("Message contract action is invalid")
		return
	}
	if err != nil {
		logrus.Error(err)
		return
	}

	response := map[string]interface{}{
		"response_kind": responseContract.ResponseKind,
		"data":          responseContract.Data,
	}
	responseByte, err := json.Marshal(response)
	if err != nil {
		logrus.Error(err)
		return
	}

	switch responseContract.BroadcastMode {
	case model.BROADCAST_ALL:

	case model.BROADCAST_DIRECT_TO:
		h.clientRoomMap[responseContract.From.RoomID][responseContract.To.ID].send <- responseByte
	case model.BROADCAST_ROOM:
		for _, targetClient := range h.clientRoomMap[messageContract.From.RoomID] {
			targetClient.send <- responseByte
		}
	case model.BROADCAST_SELF:
		targetClient := h.clientMap[messageContract.From.ID]
		targetClient.send <- responseByte
	default:
		logrus.Error("Broadcast mode is invalid")
		return
	}

	return
}
