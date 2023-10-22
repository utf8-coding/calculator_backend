package db

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
)

var (
	RedisCli *redis.Client
)

const (
	defaultRedisAdd  = "www.sunsnasserver.top:6379"
	defaultRedisPswd = "hidden"
	defaultRedisDB   = 0
)

func init() {
	for {
		connSentinel := 0
		initRedis(&connSentinel)

		if connSentinel != 0 {
			fmt.Printf("End - 0; Retry - 1: \n")
			var command string
			_, _ = fmt.Scanln(&command)
			if command == "0" {
				os.Exit(1)
			}
		} else {
			break
		}
	}

}

func initRedis(connSentinel *int) {
	redisAddr := os.Getenv("DB_REDIS_ADDR")
	if redisAddr == "" {
		log.Printf("empty DB_REDIS_ADDR? Using defaultRedisAdd \n")
		redisAddr = defaultRedisAdd
	}
	redisPswd := os.Getenv("DB_REDIS_PSWD")
	redisDB, err := strconv.Atoi(os.Getenv("DB_REDIS_DB"))
	if err != nil {
		log.Printf("Redis conn err, err: %v, using empty \n", err)
		redisDB = defaultRedisDB
	}
	if redisPswd == "" {
		log.Printf("empty DB_REDIS_DB? Using defaultRedisDB \n")
		redisPswd = defaultRedisPswd
	}
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPswd,
		DB:       redisDB,
	})
	pong, err := RedisCli.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Err conn redsi! %v", err)
		*connSentinel++
	} else {
		log.Printf("Redis conn success! pong: %v", pong)
	}
}
