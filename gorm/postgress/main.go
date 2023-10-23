package main

import (
	"time"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"uniqueIndex"`
	LastName  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"not null"`
	Country   string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func main() {
	//Create a new Postgresql database connection
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// AutoMigrate will create the necessary tables based on the \
	// defined models/structs
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	// ... Define a new post instance ...
	newUser := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "janedoe@gmail.com",
		Country:   "USA",
		Role:      "Chef",
		Age:       30,
	}

	// create a new user record in the database
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create new user: " + result.Error.Error())
	}

	// ... Handle successful creation ...
	fmt.Printf("New user %s %s was created successfully!\n", newUser.FirstName, newUser.LastName)

	// Retrieve the first user from the database
	var user User
	result = db.First(&user)
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	// Use the user record
	fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID,
		user.FirstName, user.LastName, user.Email)
}
