package controllers

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"net/http"
	"time"
	"userService/database"
	"userService/gokafka"
	helper "userService/helpers"
	"userService/logger"
	"userService/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func getUserCollection() *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.WithFields(logrus.Fields{"err": err.Error()}).Error("Failed to load .env file")
	}

	COLLECTION := os.Getenv("COLLECTION")
	userCollection := database.OpenCollection(database.MongoClient, COLLECTION)

	return userCollection
}

var log logrus.Logger = *logger.GetLogger()
var userCollection *mongo.Collection = getUserCollection()
var validate *validator.Validate

func HashPassword() {

}

func VerifyPassword() {

}

// ShowAccount godoc
// @Summary      Sign up the user
// @Description  user registration API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        FirstName  body 	string  true  "first name of user"
// @Param        LastName 	body	string  true  "last name of user"
// @Param        Password 	body	string  true  "password"
// @Param        Email 		body	string  true  "email id"
// @Success      200  {object} 	models.SignUpResponse
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/signup [post]
func Signup() gin.HandlerFunc {

	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		var user models.User

		if err := g.BindJSON(&user); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("Error to bind the input json")
			return
		}
		log.Info("input/body json bind done")

		validate = validator.New()
		validationErr := validate.Struct(user)
		if validationErr != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			log.WithFields(logrus.Fields{"error": validationErr.Error()}).
				Error("User validation error")
			return
		}

		log.Info("user validation done")

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		defer cancel()
		if err != nil {
			// log.Panic(err)
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Error! occured while checking for email"},
			)
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("Error! occured while checking for email")
			return
		}

		if count > 0 {
			// log.Panic("Email is already exists.")
			g.JSON(http.StatusBadRequest, gin.H{"error": "User already exists."})
			log.WithFields(logrus.Fields{"error": err.Error()}).
				Error("Error! occured while checking for email")

			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			// log.Panic("Password not hashed!")
			g.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("Password not hashed!")
			return
		}
		user.Password = string(hashedPassword)
		log.Info("password successfully hashed")

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()
		defer cancel()

		user.Token, err = helper.CreateToken(
			user.Email,
			user.FirstName,
			user.LastName,
			user.User_id,
		)

		defer cancel()
		if err != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("Token not created")
			return
		}

		log.Info("user token created successfully.")
		insertedNumber, insertErr := userCollection.InsertOne(ctx, user)
		defer cancel()
		if insertErr != nil {
			g.JSON(http.StatusBadGateway, gin.H{"error": insertErr.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("Insertion to mongoDB failed")
			return
		}

		g.JSON(
			http.StatusOK,
			gin.H{"insertedNumber": insertedNumber, "Token": user.Token},
		)
		go gokafka.WriteMsgToKafka("email", user)
		log.WithFields(logrus.Fields{"insertedNumber": insertedNumber}).
			Info("user successfully registered")
	}
}

// ShowAccount godoc
// @Summary      Login the user
// @Description  user log in API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        Password 	body	string  true  "password"
// @Param        Email 		body	string  true  "email id"
// @Success      200  {object} 	models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/login [post]
func Login() gin.HandlerFunc {

	return func(g *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		var user models.User

		if err := g.BindJSON(&user); err != nil {
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "gin JSON bind error", "details": err.Error()},
			)
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("gin JSON bind error")
			return
		}

		log.Debug("gin JSON binding is done")
		var userRecord models.User
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).
			Decode(&userRecord)
		defer cancel()
		if err != nil {
			g.JSON(
				http.StatusBadRequest,
				gin.H{"error": "User not found", "details": err.Error()},
			)
			log.WithFields(logrus.Fields{"error": err.Error(), "email": user.Email}).
				Error("user not found")
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(userRecord.Password),
			[]byte(user.Password),
		)
		defer cancel()
		if err != nil {
			g.JSON(
				http.StatusBadGateway,
				gin.H{"error": "wrong password", "details": err.Error()},
			)
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("user password is wrong")
			return
		}

		g.JSON(http.StatusOK, gin.H{"user": userRecord})
		log.WithFields(logrus.Fields{"email": user.Email}).Info("user successfully logged in")
	}
}

func GetUsers(gin *gin.Context) {

}

// ShowAccount 	 godoc
// @Summary      Get user details on ID
// @Description  get user details using ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        User_ID 	body	string  true  "unique user id"
// @Param        Token  	header	string  true  "user token"
// @Success      200  {object} 	models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/getUserDetails [post]
func GetUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		var body struct {
			UserID string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("gin JSON bind error")
			return
		}

		var UserDetails models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": body.UserID}).
			Decode(&UserDetails)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error(), "user_id": body.UserID}).
				Error("user id not found")
			return
		}

		response := struct {
			firstName string
			lastName  string
			Email     string
			Phone     string
			UserId    string
		}{
			UserDetails.FirstName,
			UserDetails.LastName,
			UserDetails.Email,
			UserDetails.Phone,
			UserDetails.User_id,
		}

		c.JSON(http.StatusOK, gin.H{"userDetails": response})
		log.WithFields(logrus.Fields{"user_id": body.UserID}).Info("user found in DB")
	}
}

// ShowAccount 	 godoc
// @Summary      Delete user details on ID
// @Description  delete user details using ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        User_ID 	body	string  true  "unique user id"
// @Param        Token  	header	string  true  "user token"
// @Success      200  {object} 	models.User
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /user/getUserDetails [post]
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
		defer cancel()

		var body struct {
			UserID string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("gin JSON bind error")
			return
		}

		_, err := userCollection.DeleteOne(ctx, bson.M{"user_id": body.UserID})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error(), "user_id": body.UserID}).
				Error("user id not found")
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "user deleted from DB", "userID": body.UserID})
		log.WithFields(logrus.Fields{"user_id": body.UserID}).Info("user deleted from DB")
	}
}
