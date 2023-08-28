package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ratings struct {
	Average float64 `json:"average"`
	Count   int     `json:"count"`
}

type reviews struct {
	User        string `json:"user"`
	Rating      int    `json:"rating"`
	Review_Text string `json:"review_text"`
}

type Book struct {
	Title            string    `json:"title"`
	Author           string    `json:"author"`
	Publication_Year int       `json:"publication_year"`
	ISBN             string    `json:"isbn"`
	Genre            string    `json:"genre"`
	Language         string    `json:"language"`
	Publisher        string    `json:"publisher"`
	Page_Count       int       `json:"page_count"`
	Summary          string    `json:"summary"`
	Cover_Image_URL  string    `json:"cover_image_url"`
	Ratings          ratings   `json:"ratings"`
	Reviews          []reviews `json:"reviews"`
}

func main() {
	r := Router()
	log.Fatal(r.Run(":9000"))
}

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/validate-book", ValidateBook)
	return router
}

func ValidateBook(c *gin.Context) {
	var receivedData map[string]interface{}
	/* map[string]interface{}: This is the type of the variable.
	It's a map where the keys are strings (string) and the values can be of any type (interface{})*/

	if err := json.NewDecoder(c.Request.Body).Decode(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON ",
		})
		return
	}
	/* json.NewDecoder(c.Request.Body):
		  This part creates a new JSON decoder (*json.Decoder) that reads from the Body of the incoming HTTP request.
		  The Body is a stream of data that represents the content of the HTTP request.

	     .Decode(&receivedData):
		 The Decode method of the JSON decoder reads the JSON data from the
		 input source (the HTTP request body) and attempts to decode it into the receivedData variable.
	*/

	/*
	 c.JSON(http.StatusBadRequest, gin.H{ ... }): If there's an error during decoding (invalid JSON data),
	  this line responds to the client with a 400 Bad Request status and an error message in JSON format.
	   The c.JSON function is part of the Gin framework and is used to send JSON responses.
	   In this case, it sends a JSON response with an "error" field containing the message "Invalid JSON".
	*/

	if len(receivedData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty JSON",
		})
		return
	}

	bookType := reflect.TypeOf(Book{})
	validField := make(map[string]bool)

	for i := 0; i < bookType.NumField(); i++ {
		field := bookType.Field(i)
		validField[field.Tag.Get("json")] = true
	}

	for key := range receivedData {
		if !validField[key] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid field: %s", key),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "JSON structure matches Book struct",
	})

}
