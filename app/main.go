package main

import (
	"fmt"
	"mongo-golang/app/loader"
	"mongo-golang/app/posts"
	"mongo-golang/app/users"
	"net/http"
)

func main() {
	loader.EnvLoader()
	loader.DBLoader()
	users.InitUserRouter()
	posts.InitPostRouter()
	err := http.ListenAndServe(":3000", nil)
	if err == nil {
		fmt.Println("Application started ...")
	} else {
		panic("Application not started ...")
	}
}
