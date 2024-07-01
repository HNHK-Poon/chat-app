package chat

import (
	"chat-app/pkg/store"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type Message struct {
	Type     string `json:"type"`
	Room     string `json:"room,omitempty"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver,omitempty"`
	// Content is the message content, including the success or error message
	Content  string `json:"content"`
	Password string `json:"password,omitempty"`
}

type Chat struct {
	RedisClient *redis.Client
}

func NewChat() *Chat {
	return &Chat{
		RedisClient: store.NewStore(),
	}
}

func (c *Chat) IsUserInRoom(room, username string) (bool, error) {
	isMember, err := c.RedisClient.SIsMember(context.Background(), "room:"+room+":members", username).Result()
	if err != nil {
		return false, err
	}
	return isMember, nil
}

func (c *Chat) SaveMessage(room string, message string) error {
	return store.SaveMessage(c.RedisClient, room, message)
}

func (c *Chat) GetMessages(room string) ([]string, error) {
	return store.GetMessages(c.RedisClient, room)
}

func (c *Chat) JoinRoom(room, username string) error {
	return store.AddUserToRoom(c.RedisClient, room, username)
}

func (c *Chat) LeaveRoom(room, username string) error {
	return store.RemoveUserFromRoom(c.RedisClient, room, username)
}

func (c *Chat) SetRoomPassword(room, password string) error {
	// check if room exists
	_, err := c.RedisClient.Get(context.Background(), "room:"+room+":password").Result()
	if err == nil {
		return errors.New("room already exists")
	}
	return c.RedisClient.Set(context.Background(), "room:"+room+":password", password, 0).Err()
}

func (c *Chat) CheckRoomPassword(room, password string) error {
	// check if password matches
	storedPassword, err := c.RedisClient.Get(context.Background(), "room:"+room+":password").Result()
	if err != nil {
		return err
	}
	if storedPassword != password {
		return errors.New("incorrect password")
	}
	return nil
}

func (c *Chat) DeleteRoom(room, password string) error {
	// check if password matches
	err := c.CheckRoomPassword(room, password)
	if err != nil {
		return err
	}
	return c.RedisClient.Del(context.Background(), "room:"+room+":password").Err()
}
