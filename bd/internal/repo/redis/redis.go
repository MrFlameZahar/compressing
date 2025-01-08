package repo

import (
	png "bd/bd/internal/services/png"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func InitRedis() {

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer func() {
		if err := RedisDB.Close(); err != nil {
			log.Fatalf("Failed to close Redis connection: %v", err)
		}
	}()

	for {
		// Извлекаем элемент из начала очереди
		task, err := RedisDB.LPop("image_queue").Result()
		if err == redis.Nil {
			fmt.Println("Queue is empty. Waiting...")
			time.Sleep(5 * time.Second)

		} else if err != nil {
			log.Fatalf("Failed to pop from queue: %v", err)
		}A
		var compressingService = png.NewPngServices()
		compressingService.Decode(task)

		// Обрабатываем задачу
		fmt.Printf("Processing task: %s\n", task)
	}
}

// Функция для добавления задачи в очередь Redis
func AddToQueue(imageID string) {
	err := RedisDB.LPush("image_queue", imageID).Err()
	if err != nil {
		log.Println("Ошибка при добавлении задачи в очередь:", err)
	}
}
