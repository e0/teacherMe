package main

import (
	"encoding/json"
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/e0/teacherMe/go_app/back_end/helper"
	"github.com/e0/teacherMe/go_app/back_end/model"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var session *r.Session

func main() {
	connectToDB()

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("../front_end/", true)))

	api := r.Group("/api")
	{
		api.GET("/courses", func(c *gin.Context) {
			courses, err := fetchAllCourses()

			if err != nil {
				fmt.Println(err)
			}

			c.JSON(200, helper.GetJSONFormat(courses))
		})

		api.POST("/course_create", func(c *gin.Context) {
			data := c.Query("courseData")
			courseData := map[string]interface{}{}
			json.Unmarshal([]byte(data), &courseData)
			courseID := createCourse(courseData)

			if courseID == "" {
				c.JSON(400, gin.H{"error": "Course creation failed."})
			} else {
				c.JSON(200, gin.H{"courseID": courseID})
			}
		})

		api.GET("/course/:courseID", func(c *gin.Context) {
			courseID := c.Param("courseID")
			course, err := fetchCourse(courseID)

			if err != nil {
				fmt.Println(err)
			}

			c.JSON(200, helper.GetJSONFormat(course))
		})
	}

	r.Run(":8081")
}

func connectToDB() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "teacherMe",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func createCourse(courseData map[string]interface{}) string {
	result, err := r.Table("courses").Insert(courseData).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result.GeneratedKeys[0]
}

func fetchCourse(courseID string) (model.Course, error) {
	var course model.Course

	cursor, err := r.Table("courses").Get(courseID).Run(session)

	if err != nil {
		fmt.Println(err)
		return course, err
	}

	cursor.One(&course)
	cursor.Close()

	return course, nil
}

func fetchAllCourses() ([]model.Course, error) {
	var courses []model.Course

	rows, err := r.Table("courses").Run(session)

	if err != nil {
		fmt.Println(err)
		return courses, err
	}

	err2 := rows.All(&courses)
	if err2 != nil {
		fmt.Println(err2)
		return courses, err2
	}

	return courses, nil
}
