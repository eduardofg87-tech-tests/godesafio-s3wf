package entity

import (
	"time"
)

//User data
type User struct {
	UUID        UUID        `json:"uuid"`
	Name        string    	`json:"name"`
	LastName 	string    	`json:"last_name"`
	CPF        	string    	`json:"cpf"`
	Email       string  	`json:"email"`
	CreatedAt   time.Time 	`json:"created_at"`
}