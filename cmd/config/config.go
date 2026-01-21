package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	sshIP        string
	sshPort      string
	sshLogin     string
	sshPassword  string
	sshDirectory string
	pcDirectory  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("не удалось прочитать env")
	}

	return &Config{
		sshIP:        os.Getenv("SSH_IP"),
		sshPort:      os.Getenv("SSH_PORT"),
		sshLogin:     os.Getenv("SSH_LOGIN"),
		sshPassword:  os.Getenv("SSH_PASSWORD"),
		sshDirectory: os.Getenv("SSH_DIRECTORY"),
		pcDirectory:  os.Getenv("PC_DIRECTORY"),
	}, nil
}
