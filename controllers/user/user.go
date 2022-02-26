package user

import (
	"log"
	"net/http"
	"rail/database"
	models "rail/models/user"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "user")
var validate *validator.Validate

func HashPassword() {

}

func VerifyPassword() {

}

func Signup() gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		var user models.User

		if err := g.BindJSON(&user); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate = validator.New()
		validationErr := validate.Struct(user)
		if validationErr != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			g.JSON(http.StatusBadRequest, gin.H{"error": "Error! occured while checking for email"})
			return
		}

		if count > 0 {
			// log.Panic("Email is already exists.")
			g.JSON(http.StatusBadRequest, gin.H{"error": "User already exists."})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()
		user.Token = "hello"

		insertedNumber, insertErr := userCollection.InsertOne(ctx, user)
		defer cancel()
		if insertErr != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": insertErr.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"insertedNumber": insertedNumber})
	}
}

func Login(gin *gin.Context) {

}

func GetUsers(gin *gin.Context) {

}

func GetUser(gin *gin.Context) {

}
