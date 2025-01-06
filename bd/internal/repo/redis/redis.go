package repo

import (
	"context"
	"log"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

// Создание нового клиента Redis
func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Адрес Redis
	})
	return rdb
}

// Функция для добавления задачи в очередь Redis
func addToQueue(rdb *redis.Client, task string) {
	err := rdb.LPush(ctx, "image_queue", task).Err()
	if err != nil {
		log.Println("Ошибка при добавлении задачи в очередь:", err)
	}
}
