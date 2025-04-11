package main

import (
	"errors"
	"fmt"
	"net/http"
	"projectgrom/internal/config"
	"projectgrom/internal/handlers"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Если запрос — это предварительный запрос (OPTIONS), сразу отправляем ответ
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	data, err := config.Config()
	if err != nil {
		if errors.Is(err, config.EmptyError) {
			fmt.Println("Данные не были переданы user password")
			return
		}
		fmt.Println("unkown error: ", err)
	}
	handler, err := handlers.NewHandler(data)
	if err != nil {
		fmt.Println("NewHandler error: ", err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/main", handler.Main)
	router := CORSMiddleware(mux)
	fmt.Println("Server is running at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
