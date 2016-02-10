package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/e0/teacherMe/go_app/back_end/controller"
	"github.com/e0/teacherMe/go_app/back_end/helper"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	controller.SetSession()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("../front_end/", true)))

	api := router.Group("/api")
	{
		api.GET("/courses", func(c *gin.Context) {
			courses, err := controller.FetchAllCourses()

			if err != nil {
				fmt.Println(err)
			}

			c.JSON(200, helper.GetJSONFormat(courses))
		})

		api.POST("/course_create", func(c *gin.Context) {
			data, _ := ioutil.ReadAll(c.Request.Body)
			courseData := map[string]interface{}{}
			json.Unmarshal(data, &courseData)
			courseID := controller.CreateCourse(courseData)

			if courseID == "" {
				c.JSON(400, gin.H{"error": "Course creation failed."})
			} else {
				c.JSON(200, gin.H{"courseID": courseID})
			}
		})

		api.GET("/course/:courseID", func(c *gin.Context) {
			courseID := c.Param("courseID")
			course, err := controller.FetchCourse(courseID)

			if err != nil {
				fmt.Println(err)
			}

			c.JSON(200, helper.GetJSONFormat(course))
		})
	}

	router.Run(":8081")
}
