package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filterOptions := bson.M{"_id": objectId}
	ioBody, _ := ioutil.ReadAll(request.Body)
	var bodyJSON UserModel
	_ = json.Unmarshal(ioBody, &bodyJSON)
	var usr UserModel
	err := loader.MongoDB.Collection("users").FindOne(context.Background(), filterOptions).Decode(&usr)
	if err != nil {
		data := base.BaseResponse{Message: "User Not Found", Data: err.Error()}
		dataJSON, _ := json.Marshal(data)
		response.WriteHeader(http.StatusNotFound)
		response.Write(dataJSON)
	}
	updatedBody := bson.D{{"$set", bson.D{
		{"firstname", bodyJSON.Firstname},
		{"lastname", bodyJSON.Lastname},
		{"age", bodyJSON.Age},
		{"tech", bodyJSON.Tech},
	}}}
	updateResult, errUpdateResult := loader.MongoDB.Collection("users").UpdateOne(context.Background(), filterOptions, updatedBody)
	fmt.Println("------- user ---------", updateResult, errUpdateResult)

	if updateResult.ModifiedCount > 0 {
		data := base.BaseResponse{Message: "Update User Successfully ..."}
		dataJSON, _ := json.Marshal(data)
		response.WriteHeader(http.StatusOK)
		response.Write(dataJSON)
	} else {
		if errUpdateResult != nil {
			data := base.BaseResponse{Message: "Update Process Failed ...", Data: errUpdateResult.Error()}
			dataJSON, _ := json.Marshal(data)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(dataJSON)
		} else {
			data := base.BaseResponse{Message: "No Updating Happened"}
			dataJSON, _ := json.Marshal(data)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(dataJSON)
		}
	}
}

func DeleteUserHandler(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filterOptions := bson.M{"_id": objectId}
	var usr UserModel
	err := loader.MongoDB.Collection("users").FindOne(context.Background(), filterOptions).Decode(&usr)
	if err != nil {
		data := &base.BaseResponse{Message: "Error in Finding Document", Data: err.Error()}
		dataJSON, _ := json.Marshal(data)
		response.WriteHeader(http.StatusNotFound)
		response.Write(dataJSON)
		return
	}
	deletedResult, err := loader.MongoDB.Collection("users").DeleteOne(context.Background(), filterOptions)
	if deletedResult.DeletedCount > 0 {
		data := &base.BaseResponse{Message: "Delete User Successfully ..."}
		dataJSON, _ := json.Marshal(data)
		response.WriteHeader(http.StatusNotFound)
		response.Write(dataJSON)
		return
	} else {
		if err != nil {
			data := &base.BaseResponse{Message: "Error In Delete User", Data: err.Error()}
			dataJSON, _ := json.Marshal(data)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(dataJSON)
			return
		} else {
			data := &base.BaseResponse{Message: "Delete Operation Not Happen", Data: err.Error()}
			dataJSON, _ := json.Marshal(data)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(dataJSON)
			return
		}
	}
}
