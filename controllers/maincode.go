package controllers

import (
	"fmt"
	"net/http"
	"test/repo/Parameters"
	"test/repo/structs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(c *gin.Context) {
	var SignUpTemp structs.SignUpStruct
	c.ShouldBindJSON(&SignUpTemp)

	if SignUpTemp.Name == "" || SignUpTemp.Surname == "" || SignUpTemp.Login == "" || SignUpTemp.Password == "" {
		c.JSON(404, "Error")
	} else {

		client, ctx := parameters.DBConnection()

		DBConnect := client.Database("Music").Collection("Users")

		id := primitive.NewObjectID().Hex()
		Hashed, _ := parameters.HashPassword(SignUpTemp.Password)
		DBConnect.InsertOne(ctx, bson.M{
			"_id":      id,
			"name":     SignUpTemp.Name,
			"surname":  SignUpTemp.Surname,
			"login":    SignUpTemp.Login,
			"password": Hashed,
		})
		c.JSON(200, "Success")
	}

}
func Login(c *gin.Context) {
	var LoginTemp structs.SignUpStruct
	c.ShouldBindJSON(&LoginTemp)

	if LoginTemp.Login == "" || LoginTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := parameters.DBConnection()

		DBConnect := client.Database("Music").Collection("Users")

		result := DBConnect.FindOne(ctx, bson.M{
			"login": LoginTemp.Login,
		})

		var userdata structs.SignUpStruct
		result.Decode(&userdata)
		isValidPass := parameters.CompareHashPasswords(userdata.Password, LoginTemp.Password)
		fmt.Println(isValidPass)

		if isValidPass {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    "FakhriyaPlaylist",
				Value:   userdata.Login,
				Expires: time.Now().Add(160 * time.Second),
			})
			c.JSON(200, "success")
		} else {
			c.JSON(404, "Wrong login or password")
		}
	}
}
func Music(c *gin.Context) {
	var Cookiedata, CookieError = c.Request.Cookie("FakhriyaPlaylist")
	fmt.Printf("Cookiedata: %v\n", Cookiedata)
	if CookieError != nil {
		fmt.Printf("CookieError: %v\n", CookieError)
		c.JSON(404, "errorcookie")
	} else {
		var AddingMusicTemp structs.SearchMusic
		c.ShouldBindJSON(&AddingMusicTemp)

		if AddingMusicTemp.Name == "" || AddingMusicTemp.Author == "" || AddingMusicTemp.Date == "" {
			c.JSON(404, "Error")
		} else {
			Client, context := parameters.DBConnection()
			id := primitive.NewObjectID().Hex()
			var createDB = Client.Database("Music").Collection("AddMusic")
			var InsertResult, InsertError = createDB.InsertOne(context, bson.M{
				"_id":    id,
				"name":   AddingMusicTemp.Name,
				"author": AddingMusicTemp.Author,
				"date":   AddingMusicTemp.Date,
			})

			if InsertError != nil {
				fmt.Printf("InsertError: %v\n", InsertError)
			} else {
				c.JSON(200, "Success")
				fmt.Printf("InsertResult: %v\n", InsertResult)
			}

		}

	}

}
func Search(c *gin.Context) {
	var SearchMusic = []structs.SearchMusic{}
	var SearchMusicTemp structs.SearchMusic
	c.ShouldBindJSON(&SearchMusicTemp)
	var Cookiedata, CookieError = c.Request.Cookie("FakhriyaPlaylist")
	fmt.Printf("Cookiedata: %v\n", Cookiedata)
	if CookieError != nil {
		c.JSON(401, "cookie do not exist")

	} else {
		Client, ctx := parameters.DBConnection()
		connectmusic := Client.Database("Music").Collection("AddMusic")
		result, _ := connectmusic.Find(ctx, bson.M{
			"name": SearchMusicTemp.Name,
		})
		fmt.Printf("result: %v\n", result)
		for result.Next(ctx) {
			var DBTemp structs.SearchMusic
			result.Decode(&DBTemp)

			SearchMusic = append(SearchMusic, DBTemp)
		}
		c.JSON(200, SearchMusic)
	}
}
func AllMusics(c *gin.Context) {
	var SearchMusic = []structs.SearchMusic{}
	var SearchMusicTemp structs.SearchMusic
	c.ShouldBindJSON(&SearchMusicTemp)
	var Cookiedata, CookieError = c.Request.Cookie("FakhriyaPlaylist")
	fmt.Printf("Cookiedata: %v\n", Cookiedata)
	if CookieError != nil {
		c.JSON(401, "cookie do not exist")

	} else {
		Client, ctx := parameters.DBConnection()
		connectmusic := Client.Database("Music").Collection("AddMusic")
		result, _ := connectmusic.Find(ctx, bson.M{})
		fmt.Printf("result: %v\n", result)
		for result.Next(ctx) {
			var DBTemp structs.SearchMusic
			result.Decode(&DBTemp)

			SearchMusic = append(SearchMusic, DBTemp)
		}
		c.JSON(200, SearchMusic)
	}
}
