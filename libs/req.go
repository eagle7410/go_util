package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func SetCorsMiddleware(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetCors(w, r)

		if r.Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func SetCors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(*Env.AllowedMethods, ","))
}

func SendJsonBadRequest(w http.ResponseWriter, mess string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("{\"message\" : \"" + mess + "\"}"))
}

func SendJsonServeError(w http.ResponseWriter, mess string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"message\" : \"" + mess + "\"}"))
}

func SendJsonInternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"message\" : \"Internal server error\"}"))
}

func SendJson(statusCode int, w http.ResponseWriter, data interface{}) {

	payload, err := json.Marshal(data)

	if err != nil {
		LogEF("Error pack json payload : %v", err)
		SendJsonInternalError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(payload)

}

func SendJson403(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("{\"message\" : \"Access denied\"}"))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG \n  IP: %v\n  Host: %v\n", ReadUserIP(r), r.Host)
}
