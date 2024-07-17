package helper

import (
	"context"
	"fmt"
	"log"

	database "github.com/Aman123at/gin-userauth/db"
	"github.com/Aman123at/gin-userauth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection()

func SignUpUser(userDetail model.User) {
	_, err := collection.InsertOne(context.Background(), userDetail)

	if err != nil {
		fmt.Println("Signup Error")
		log.Fatal(err)
	}
}

func LogoutUser(userId string) {
	uid, hexerr := primitive.ObjectIDFromHex(userId)

	if hexerr != nil {
		fmt.Println("Error: Unable to convert Id to hex")
		log.Fatal(hexerr.Error())
	}

	filter := bson.M{"_id": uid}

	result, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"token": ""}})

	fmt.Printf("Result count on Update %v", result.ModifiedCount)

	if err != nil {
		fmt.Println("Error: Unable to update token in DB")
		log.Fatal(err)
	}
}

func GetAllUsers() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		fmt.Println("Error: while executing find query")
		log.Fatal(err.Error())
	}

	var users []primitive.M

	for cursor.Next(context.Background()) {
		var user bson.M

		decodeErr := cursor.Decode(&user)

		if decodeErr != nil {
			fmt.Println("Error: Unable to decode user in get all users")
			log.Fatal(decodeErr.Error())
		}

		users = append(users, user)
	}
	return users
}

func UpdateTokenInDB(email string, token string) {
	filter := bson.M{"email": email}

	result, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"token": token}})

	fmt.Printf("Result count on Update %v", result.ModifiedCount)

	if err != nil {
		fmt.Println("Error: Unable to update token in DB")
		log.Fatal(err)
	}
}

func GetOneUser(userId string) model.User {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		fmt.Println("Error: Unable to convert Id to hex")
		log.Fatal(err)
	}

	var user model.User

	filter := bson.M{"_id": id}

	_ = collection.FindOne(context.Background(), filter).Decode(&user)

	return user
}

func GetUserByMail(email string) model.User {
	filter := bson.M{"email": email}

	var user model.User

	_ = collection.FindOne(context.Background(), filter).Decode(&user)

	return user
}
