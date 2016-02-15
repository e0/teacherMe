package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/e0/teacherMe/go_app/back_end/controller"
	"github.com/e0/teacherMe/go_app/back_end/helper"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	controller.SetSession()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../front_end/", true)))

	publicAPI := router.Group("/api/public")
	privateAPI := router.Group("/api/private")
	privateAPI.Use(jwt.Auth(getDecodedSecret()))

	publicAPI.GET("/courses", func(c *gin.Context) {
		courses, err := controller.FetchAllCourses()

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, helper.GetJSONFormat(courses))
	})

	privateAPI.POST("/course_create", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		courseID := controller.CreateCourse(data)

		if courseID == "" {
			c.JSON(400, gin.H{"error": "Course creation failed."})
		} else {
			c.JSON(200, gin.H{"courseID": courseID})
		}
	})

	publicAPI.GET("/course/:courseID", func(c *gin.Context) {
		courseID := c.Param("courseID")
		course, err := controller.FetchCourse(courseID)

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, helper.GetJSONFormat(course))
	})

	router.Run(":8081")
}

func getDecodedSecret() string {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := map[string]string{}
	decoder.Decode(&configuration)
	decodedSecret, _ := base64.URLEncoding.DecodeString(configuration["Auth0Secret"])
	return string(decodedSecret)
}
