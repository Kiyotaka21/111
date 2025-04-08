package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	User string
}

func GET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := r.FormValue("user")
	va := &User{}
	err := json.Unmarshal([]byte(user), va)
	if err != nil {
		fmt.Println(err)
		return
	}
	value, err := json.Marshal(va)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(value)
}

func main() {
	router := httprouter.New()
	router.GET("/", GET)
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
