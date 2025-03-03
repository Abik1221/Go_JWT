package controllers

import (
	"context"
	"net/http"
	"time"

	database "github.com/NahomKeneni/go_jwt/databse"
	"github.com/NahomKeneni/go_jwt/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Validator = validator.New()
var UserCollection = database.OpenCollection(database.Client, "user")

func HashPassward(Passward string) (string, err) {
    bytes, err := bcrypt.GenerateFromPassward([]byte(Passward), 14)
	if err != nil {
		return "", err
	}
    return string(bytes), nil
}

func VerifyPassword(userPaassword, providedPassword string) (bool, string){
    err := bycrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPaassword))
	msg := ""
	check := true

	if err != nil{
		msg = "the pasword doesn't match please insert the correct password"
		check = false
		return check, msg
	}

	return check, msg
} 

func Signup() gin.HandlerFunc { 
      return func (c* gin.Context){
		var User = models.User
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validator.Struct(user)
		if validatonErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"error accured while cheecking for the validation"})
			return
		}
        

		count, err := database.UserCollection.CountDocuments(ctx, bsom.M{"email":user.Email})
		if err != nil {
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

		user.Passward , err := HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
        
		user.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refresh_token, err := helpers.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.Uid, user.User_type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Token = &token
		user.Refresh_token = &refresh_token

		rseultInsertionNumber, insertionErr := UserCollection.InsertOne(ctx, user)
		if insertionErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while inserting the uer detail in the database !"})
		}
		c.JSON(http.StatusOk, rseultInsertionNumber)
	  }
}

func Login() gin.HandlerFunc{
   return func(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()
	var user models.User

	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

   err = database.userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
   if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while cheecking for the user"})
	return
}

   msg, is_valid := VerifyPassword(user.Passward, foundUser.Passward)
   if msg != nil && is_valid == false {
	  c.JSON{http.StatusBadRequest, gin.H{"error": msg }}
	  return
   }

   token, refresh_token, err := helpers.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.Uid, foundUser.User_type)
   if err != nil {
	c.JSON{http.StatusBadRequest, gin.H{"error", err.Error()}}
   }
  
}
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
