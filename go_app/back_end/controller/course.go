package controller

import (
	"encoding/json"
	"fmt"
	"strings"

	r "github.com/dancannon/gorethink"
	"github.com/e0/teacherMe/go_app/back_end/model"
)

// CreateCourse inserts a new item in the courses table
func CreateCourse(data []byte) string {
	courseData := map[string]interface{}{}
	json.Unmarshal(data, &courseData)
	courseDescription, _ := courseData["description"].(string)
	courseData["description"] = strings.Replace(courseDescription, "\n", "<br />", -1)

	result, err := r.Table("courses").Insert(courseData).RunWrite(session)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result.GeneratedKeys[0]
}

// FetchCourse returns an item from the courses table
func FetchCourse(courseID string) (model.Course, error) {
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

// FetchAllCourses returns all items from the courses table
func FetchAllCourses() ([]model.Course, error) {
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
