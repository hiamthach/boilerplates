package middleware

import (
	"net/http"
	"regexp"

	"go-microservices/app/config"
)

func CORS(w http.ResponseWriter, r *http.Request) {
	if allowedOrigin(r.Header.Get("Origin")) {
		corsHeader := config.Get().Server.CorsHeader
		if len(corsHeader) == 0 {
			corsHeader = "origin, accept, mm-client-id, access-control-allow-methods, content-type, access-control-allow-origin, access-control-allow-credentials, access-control-allow-headers, shard, mm-user-id"
		}
		corsMethod := config.Get().Server.CorsMethod
		if len(corsMethod) == 0 {
			corsMethod = "GET, HEAD, OPTIONS, POST, PUT"
		}
		cors := config.Get().Server.Cors
		if len(cors) == 0 {
			cors = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", cors)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", corsMethod)
		w.Header().Set("Access-Control-Allow-Headers", corsHeader)
	}
	if r.Method == "OPTIONS" {
		return
	}
}

func allowedOrigin(origin string) bool {
	if config.Get().Server.Cors == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(config.Get().Server.Cors, origin); matched {
		return true
	}
	return false
}
