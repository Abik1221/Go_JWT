package helpers

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	User_type  string
	Uid        string
	jwt.StandardClaims
}

var UserCollection = database.OprnCollection(database.Client, "user")

var SECRETE_KEY = os.Getenv("SECRETE_KEY")

func GenerateAllTokens(email string, firstname string, lastname string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstname,
		Last_name:  lastname,
		User_type:  userType,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		}
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETE_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodSH256, refreshClaims).SignedString([]byte(SECRETE_KEY))

	if err != nil {
		log.Panic(err)
		return err
	}
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"$set", bson.D{{"token", signedToken}, {"refresh_token", signedRefreshToken}}})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"$set", bson.D{{"updated_at", Updated_at}}})

	upsert := true
	filter := bson.M{"user_id": userId}
	_, err := UserCollection.UpdateOne(ctx, filter, updateObj, &options.UpdateOptions{Upsert: &upsert})
	defer cancel()
}

