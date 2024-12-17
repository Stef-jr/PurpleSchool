package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Key string
}

func NewCfg() *Config{

	err := godotenv.Load()
	if err != nil{
		panic("Не удалось Прочесть ENV файл")
	}
	key := os.Getenv("KEY")
	if key == ""{
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Config{Key: key}
}