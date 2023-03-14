package users

import (
	"context"
	"encoding/json"
	"fmt"
	"mongo-golang/app/base"
	"mongo-golang/app/loader"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	var usr UserModel
	err := json.NewDecoder(request.Body).Decode(&usr)
	if err != nil || usr.Firstname == "" || usr.Lastname == "" || usr.Age == 0 || usr.Tech == 0 {
		response.WriteHeader(http.StatusBadRequest)
		data := &base.BaseResponse{Message: "Invalid Request Body"}
		dataJSON, _ := json.Marshal(data)
		response.Write(dataJSON)
		return
	}
	usr.ID = primitive.NewObjectID()
	_, err = loader.MongoDB.Collection("users").InsertOne(context.Background(), &usr)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		data := &base.BaseResponse{Message: "Create User failed ..."}
		dataJSON, _ := json.Marshal(data)
		response.Write(dataJSON)
		return
	}
	data := &base.BaseResponse{Message: "Created Successfully ...", Data: usr}
	dataJSON, _ := json.Marshal(data)
	response.Write(dataJSON)
}

func GetAllUsersHandler(response http.ResponseWriter, request *http.Request) {
	dataMongo, err := loader.MongoDB.Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		data := &base.BaseResponse{Message: "Get All User failed ..."}
		dataJSON, _ := json.Marshal(data)
		response.Write(dataJSON)
		return
	}
	var users []UserModel
	for dataMongo.Next(context.Background()) {
		var usr UserModel
		dataMongo.Decode(&usr)
		users = append(users, usr)
	}
	data := &base.BaseResponse{Message: "Get All Lists ...", Data: users}
	dataJSON, _ := json.Marshal(data)
	response.Write(dataJSON)
}

func GetUserByIdHandler(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := string(query["id"][0])
	objectId, _ := primitive.ObjectIDFromHex(id)
	filterOption := bson.M{"_id": objectId}
	var usr UserModel
	fmt.Println(filterOption)
	err := loader.MongoDB.Collection("users").FindOne(context.Background(), filterOption).Decode(&usr)
	if err != nil {
		data := &base.BaseResponse{Message: "Error in Finding Document", Data: err.Error()}
		dataJSON, _ := json.Marshal(data)
		response.WriteHeader(http.StatusNotFound)
		response.Write(dataJSON)
		return
	}
	data := &base.BaseResponse{Message: "Get Single User ...", Data: usr}
	dataJSON, _ := json.Marshal(data)
	response.Write(dataJSON)
}

func UpdateUserHandler(response http.ResponseWriter, request *http.Request) {}

func DeleteUserHandler(response http.ResponseWriter, request *http.Request) {}
