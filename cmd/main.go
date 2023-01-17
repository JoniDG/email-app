package main

import (
	"context"
	"email-app/internal/controller"
	"email-app/internal/defines"
	"email-app/internal/repository"
	"email-app/internal/service"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/go-redis/redis/v9"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(defines.EnvRedisHost),
	})
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error Ping Redis: %+v\n", err)
	}

	from := os.Getenv("EMAIL_SENDER_USER")
	password := os.Getenv("EMAIL_SENDER_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	auth := smtp.PlainAuth("", from, password, host)
	emailRepo := repository.NewEmailRepository(auth)
	svc := service.NewEmailService(emailRepo)
	ctrl := controller.NewEmailController(svc)

	queue := defines.QueueEmail
	fmt.Printf("Email running, reading from %s\n", queue)
	for {
		payload, err := redisClient.LPop(ctx, queue).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				continue
			}
			log.Printf("Error receiving payload: %+v\n", err)
		}
		log.Printf("Processing: %+v\n", payload)
		ctrl.Handle(&payload)
	}
}
