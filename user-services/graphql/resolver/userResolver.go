package resolver

import (
	"github.com/graphql-go/graphql"

	"user-services/database"
	"user-services/graphql/types"
)

func GetUsers(params graphql.ResolveParams) (interface{}, error) {
	rows, err := database.DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var user types.User
		row := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, nil
}

func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	var user types.User
	user.Name = params.Args["name"].(string)
	user.Email = params.Args["email"].(string)
	user.Password = params.Args["password"].(string)

	stmt, err := database.DB.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(lastID)

	return user, nil
}

func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	var user types.User
	user.ID = params.Args["id"].(int)
	user.Name = params.Args["name"].(string)
	user.Email = params.Args["email"].(string)
	user.Password = params.Args["password"].(string)

	stmt, err := database.DB.Prepare("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		stmt, err := database.DB.Prepare("DELETE FROM users WHERE id = ?")
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
