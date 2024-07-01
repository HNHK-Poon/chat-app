package store

import (
	"chat-app/pkg/config"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewStore() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr + config.RedisPort,
	})
	return client
}

func AddUserToRoom(client *redis.Client, room, username string) error {
	return client.SAdd(ctx, "room:"+room+":members", username).Err()
}

func RemoveUserFromRoom(client *redis.Client, room, username string) error {
	return client.SRem(ctx, "room:"+room+":members", username).Err()
}

func SaveMessage(client *redis.Client, room string, message interface{}) error {
	err := client.LPush(ctx, room, message).Err()
	return err
}

func GetMessages(client *redis.Client, room string) ([]string, error) {
	messages, err := client.LRange(ctx, room, 0, -1).Result()
	return messages, err
}
