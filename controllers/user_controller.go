package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Validator = validator.New()

func HashPassward() {

}

func VerifyPassword() {

}

func Signup() gin.HandlerFunc {
      return func (c* gin.Context){
		var User = models.User
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := Validator.Struct(user)
		if validatonErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err":"error accured while cheecking for the validation"})
			return
		}
        

		count, err := database.UserCollection.CountDocuments(ctx, bsom.M{"email":user.Email})
		if err := nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": "error occured while cheecking for the count"}}
			return
		}

		if count > 0 {
			c.JSON{http.StatusInternalServerError, gin.H{"error":"the email already exists"}}
			return
		}

		count, err = database.UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
        if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occured while cheecking for the phone count",
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this user phone number already exists please change it"})
			return
		}
        
		user.Created_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refresh_token, err = helpers.GenerateAllTokens(user.User_id, user.First_name, user.Last_name, user.)
		user.Token = &token
		user.refresh_token = &refresh_token

		rseultInsertionNumber, insertionErr :=database.UserCollection.InsertOne(ctx, user)
		if insertionErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while inserting the uer detail in the database !"})
		}
		c.JSON(http.StatusOk, rseultInsertionNumber)
	  }
}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Param("id")

		err := helper.MatchUserIdToUid(c, user_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()

		var user models.User

		user, err := database.UserCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&User)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
