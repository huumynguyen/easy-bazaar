package services

import "github.com/go-redis/redis/v8"

type BazaarService struct {
	Client *redis.Client
}

type IBazaarService interface {
}
