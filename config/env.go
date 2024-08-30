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

// docs: global variable that holds the environment variables
var Envs = initConfig();

func initConfig() Config {
	// docs: load the environment variables into ENV for this process.
	godotenv.Load();

	// docs: get the env on the process
	return Config {
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBName: getEnv("DB_NAME", "ecom_tiago_db"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	}
}

// docs: get the environment variable or fallback to the default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value;
	}
	return fallback;
}