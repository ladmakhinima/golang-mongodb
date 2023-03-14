package posts

import (
	"fmt"
	"net/http"
)

func InitPostRouter() {
	http.HandleFunc("/api/posts/create", CreatePostHandler)
	http.HandleFunc("/api/posts/select", GetAllPostsHandler)
	http.HandleFunc("/api/posts/select/:id", GetPostByIdHandler)
	http.HandleFunc("/api/posts/update", UpdatePostHandler)
	http.HandleFunc("/api/posts/delete", DeletePostHandler)
	fmt.Println("Posts router init successfully ...")
}
