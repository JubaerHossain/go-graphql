package resolver

import (
	"course-services/database"
	"course-services/graphql/types"
	"course-services/query"
	"course-services/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetCourses(params graphql.ResolveParams) (interface{}, error) {
	page, ok := params.Args["page"].(int)
	if !ok {
		page = 1
	}
	pageSize, ok := params.Args["pageSize"].(int)
	if !ok {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	fmt.Println(offset)
	rows, err := database.DB.Query("SELECT id, name, description FROM courses ORDER BY id DESC LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []types.Course
	for rows.Next() {
		var course types.Course
		var description sql.NullString
		err := rows.Scan(&course.ID, &course.Name, &description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func GetCourse(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var course types.Course
		row := database.DB.QueryRow("SELECT id, name, description, status FROM course WHERE id = ?", id)
		err := row.Scan(&course.ID, &course.Name, &course.Description, &course.Status)
		if err != nil {
			return nil, err
		}

		return course, nil
	}

	return nil, nil
}

func CreateCourse(params graphql.ResolveParams) (interface{}, error) {
	forms := map[string]interface{}{
		"name":        params.Args["name"],
		"description": params.Args["description"],
		"user_id":     params.Args["user_id"],
		"status":      params.Args["status"],
		"createdAt":   utils.GetTimeNow(),
	}
	forms["table"] = "courses" // fmt.Println(user)

	id, err := query.Insert(forms, database.DB)
	if err == nil {
		return nil, err
	}
	forms["id"] = id
	return forms, nil
}

func UpdateCourse(params graphql.ResolveParams) (interface{}, error) {
	forms := map[string]interface{}{
		"name":        params.Args["name"],
		"description": params.Args["description"],
		"status": params.Args["status"],
	}
	forms["table"] = "courses"
	forms["id"] = params.Args["id"].(int)
	err := query.Update(forms, database.DB)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("update course failed")
	}
	id, ok := params.Args["id"].(int)
	if ok {
		var course types.Course
		row := database.DB.QueryRow("SELECT id, name, description, status FROM courses WHERE id = ?", id)
		err := row.Scan(&course.ID, &course.Name, &course.Description, &course.Status)
		if err != nil {
			return nil, err
		}

		return course, nil
	}

	return nil, nil
}

func DeleteCourse(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		stmt, err := database.DB.Prepare("DELETE FROM courses WHERE id = ?")
		if err != nil {
			return nil, err
		}
		_, err = stmt.Exec(id)
		if err != nil {
			return nil, err
		}

		return id, nil
	}

	return nil, nil
}
