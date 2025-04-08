package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	Userd string `json:"user"`
}

func GET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := r.FormValue("user")
	Vuser := &User{Userd: user}
	value, err := json.Marshal(Vuser)
	if err != nil {
		fmt.Println("err go:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(value)
}

func main() {
	router := httprouter.New()
	router.GET("/", GET)
	router.ServeFiles("/src/app/*filepath", http.Dir("./src/app"))
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
