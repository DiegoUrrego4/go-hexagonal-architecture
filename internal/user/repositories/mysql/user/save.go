package user

import (
	"log"
	"time-tracker/internal/user/domain"
)

func (r Repository) Save(user domain.User) error {
	_, err := r.Db.Exec("INSERT INTO users (id, name, last_name, age) VALUES (?,?,?,?)", user.ID, user.Name, user.LastName, user.Age)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}
