package config

import "os"

func GetFrontendURL() string {
	url := os.Getenv("FRONTEND_URL")
	if url == "" {
		url = "http://localhost:5173"
	}
	return url
}
