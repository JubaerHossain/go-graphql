package resolver

import (
	"course-services/database"
	"course-services/graphql/types"
	"course-services/query"
	"course-services/utils"
	"errors"
	"fmt"
	"strings"

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
	colmap := query.TableFields(params.Info.FieldASTs)
	funclabel := fmt.Sprint(params.Info.Path.Key)
	cols := colmap[funclabel].([]string) //
	selectColumn := strings.Join(cols, ",")
	sql := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT %d OFFSET %d;", selectColumn, "courses", pageSize, offset)
	fmt.Println(sql)
	rows, err := query.GetAllRowsByQuery(sql, database.DB)
	if err != nil {
		return nil, err
	}

	if len(rows) == 1 {
		return rows[0], nil
	}
	return nil, errors.New("No data found")
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
	fmt.Println(params.Source, params.Args, params.Info.VariableValues, params.Info.FieldASTs)
	funclabel := params.Info.Path.Key.(string)
	colmap := query.TableFields(params.Info.FieldASTs)
	cols := colmap[funclabel].([]string)
	id := params.Args["id"]

	selectColumn := strings.Join(cols, ",")
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id='%v';", selectColumn, "account", id)
	fmt.Println(sql)
	return nil, nil
	var course types.Course
	course.Name = params.Args["name"].(string)
	course.Description = params.Args["description"].(string)
	course.User_id = params.Args["user_id"].(int)
	course.Status = params.Args["status"].(string)
	course.CreatedAt = utils.GetTimeNow()
	fmt.Println(course)
	// forms["table"] = "courses"
	// fmt.Println(forms)
	id, err := query.Insert("courses", course, database.DB)
	if err != nil {
		return nil, err
	}
	// course.ID = int(id)
	return course, nil
}

func UpdateCourse(params graphql.ResolveParams) (interface{}, error) {
	forms := map[string]interface{}{
		"name":        params.Args["name"],
		"description": params.Args["description"],
		"status":      params.Args["status"],
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
