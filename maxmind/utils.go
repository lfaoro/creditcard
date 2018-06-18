package maxmind

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	// Load variables from .env file in case the user has one.
	_ = godotenv.Load()

	envVars := []string{"MAXMIND_USER", "MAXMIND_PASSWORD"}
	for _, e := range envVars {
		_, yes := os.LookupEnv(e)
		if !yes {
			log.Fatal("Missing required ENV variable: ", e)
		}
	}
	_, yes := os.LookupEnv("MAXMIND_ENDPOINT")
	if yes {
		URL = os.Getenv("MAXMIND_ENDPOINT")
		return
	}
	URL = "https://minfraud.maxmind.com/minfraud/v2.0/insights"
}

func encodeAuth() string {
	user := os.Getenv("MAXMIND_USER")
	pass := os.Getenv("MAXMIND_PASSWORD")
	combined := fmt.Sprintf("%v:%v", user, pass)
	return base64.StdEncoding.EncodeToString([]byte(combined))
}
