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
	Vuser := &User{Userd: "anton"}
	value, err := json.Marshal(Vuser)
	if err != nil {
		fmt.Println("err go:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(value)
}

func main() {
	router := httprouter.New()
	router.GET("/get", GET)
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
