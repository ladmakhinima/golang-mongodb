package users

import (
	"fmt"
	"net/http"
)

func InitUserRouter() {
	http.HandleFunc("/api/users/create", CreateUserHandler)
	http.HandleFunc("/api/users/select", GetAllUsersHandler)
	http.HandleFunc("/api/users/single", GetUserByIdHandler)
	http.HandleFunc("/api/users/update", UpdateUserHandler)
	http.HandleFunc("/api/users/delete", DeleteUserHandler)
	fmt.Println("Users router init successfully ...")
}
