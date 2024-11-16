package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fmt"
	"time"
)

func CreateUser(user models.User) (models.User, error) {
	fmt.Println(user)
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Id, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `
	  UPDATE users
	  SET name = $2, email = $3, password = $4, updated_at = $5
	  WHERE id = $1
	  RETURNING id`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = id
	return user, nil
}

func RetrunUsers(id int) []models.User {
	db := storage.GetDB()
	sqlStatement := `SELECT id ,name, email, password FROM users WHERE id = $1`
	var name string
	var email string
	var password string
	var ID int
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		err := rows.Scan(&ID, &name, &email, &password)
		if err != nil {
			panic(err)
		}
		users = append(users, models.User{Id: ID, Name: name, Email: email, Password: password})
	}
	return users
}
