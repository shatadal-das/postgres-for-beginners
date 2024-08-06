package main

import (
	"fmt"
	"os"
	"postgres-for-beginners/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func main() {
	godotenv.Load()

	DB_URI := os.Getenv("POSTGRES_URI")

	sql := services.NewSQLDb(DB_URI)
	defer sql.Close()

	sql.CreateTableFromFilename("sql/user/createTable.sql")

	user := User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	sql.InsertFromFilename("sql/user/insertData.sql", user.Name, user.Email, user.Password)

	var userData User
	sql.FindOne("SELECT name, email, password FROM public.user WHERE email = $1", &userData, "john@example.com")

	fmt.Println("Data from database:", userData)
}
