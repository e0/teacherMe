package main

import (
	"fmt"
    
	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
    "github.com/e0/teacherMe/go_app/back_end/model"
    "github.com/e0/teacherMe/go_app/back_end/helper"
)

var session *r.Session

func main() {
	connectToDB()
	// courseId := createCourse()
    // fmt.Println(courseId)

	r := gin.Default()

    r.Use(static.Serve("/", static.LocalFile("../front_end/", true)))

    api := r.Group("/api")
    {
        api.GET("/course/:courseId", func(c *gin.Context) {
            courseId := c.Param("courseId")
            course, err := fetchCourse(courseId)
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

func createCourse() string {
	var data = map[string]interface{}{
		"id":          "1",
		"title":       "Making your first web page",
		"description": "<p>There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.</p><p>It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).</p>",
		"discussions": []map[string]interface{}{
			{"id": 2, "name": "second discussion", "date": "2015-06-12"},
			{"id": 1, "name": "first discussion", "date": "2015-06-11"},
		},
		"downloads": []map[string]interface{}{
			{"id": 2, "name": "Lesson 2", "date": "2015-06-12"},
			{"id": 1, "name": "Lesson 1", "date": "2015-06-11"},
		},
		"assignments": []map[string]interface{}{
			{"id": 2, "name": "Lesson 2 assignment", "date": "2015-06-12"},
			{"id": 1, "name": "Lesson 1 assignment", "date": "2015-06-11"},
		},
		"teacherName":   "e0",
		"studentsCount": 42,
	}

	result, err := r.Table("courses").Insert(data).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result.GeneratedKeys[0]
}

func fetchCourse(courseId string) (model.Course, error) {
    var course model.Course

    cursor, err := r.Table("courses").Get(courseId).Run(session)

    if err != nil {
        fmt.Println(err)
        return course, err
    }

    cursor.One(&course)
    cursor.Close()

    return course, nil
}
