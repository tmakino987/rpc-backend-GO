package util

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

const path = "app/etc/env/.env"

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (env *Env) GetEnvValue(key string) string {
	value := os.Getenv(key)

	// .envから変数情報を取得
	if value == "" {
		err := godotenv.Load(path)
		if err != nil {
			log.Error("環境変数の読み込みに失敗しました")
			log.Error(err)
			return ""
		}
		value = os.Getenv(key)
	}

	return value
}

func (env *Env) GetCORSHosts() []string {
	return strings.Split(os.Getenv("CORS_ALLOW_HOSTS"), ",")
}
