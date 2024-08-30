package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost	string;
	Port 				string;
	
	DBName			string;
	DBUser			string;
	DBPassword	string;
	DBAddress		string;
}

var Envs = initConfig();

func initConfig() Config {
	godotenv.Load();

	return Config {
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBName: getEnv("DB_NAME", "ecom_tiago_db"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value;
	}
	return fallback;
}