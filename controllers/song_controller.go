package controllers

import (
	"context"
	"fmt"
	"moodMusicSongsApi/configs"
	"moodMusicSongsApi/models"
	"moodMusicSongsApi/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var songCollection *mongo.Collection = configs.GetCollection("songs")

func GetManySongs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// var queryParamsList = make(map[string]string)
	// queryParamsList["mood"] = c.Params("emotion")
	// queryParamsList["movie"] = c.Query("movie")
	// queryParamsList["language"] = c.Query("language")
	// queryParamsList["genre"] = c.Query("genre")

	// fmt.Println(c.Params("emotion"))

	// filter := make(bson.M)
	// for key, element := range queryParamsList {
	// 	if element != "" {
	// 		filter[key] = element
	// 	}
	// }

	fmt.Println(songCollection)

	var songs []models.Song
	// results, err := songCollection.Find(ctx, filter)
	results, err := songCollection.Find(ctx, bson.M{"mood": "emotion"})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	fmt.Println(c.Params("emotion"))
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleSong models.Song
		if err := results.Decode(&singleSong); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		songs = append(songs, singleSong)
	}
	fmt.Println(c.Params("emotion"))
	return c.Status(http.StatusCreated).JSON(responses.Response{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": songs},
	})
}
