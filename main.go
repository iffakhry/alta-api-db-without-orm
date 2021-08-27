package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string
	Name     string
	Username string
	Gender   string
}

func main() {
	db, err := sql.Open("mysql", "root:qwerty123@tcp(localhost:3306)/immersive")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Masukkan pilihan (lihat/tambah)")
	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "lihat":
		row, err := db.Query("SELECT id, name, username, gender FROM users")
		if err != nil {
			panic(err.Error())
		}
		for row.Next() {
			var user User
			err = row.Scan(&user.Id, &user.Name, &user.Username, &user.Gender)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("Name", user.Name)
			fmt.Println("Gender", user.Gender)
		}
	case "tambah":
		newUser := User{}
		fmt.Println("ID")
		fmt.Scanln(&newUser.Id)
		fmt.Println("Name")
		fmt.Scanln(&newUser.Name)
		fmt.Println("Username")
		fmt.Scanln(&newUser.Username)
		fmt.Println("Gender")
		fmt.Scanln(&newUser.Gender)
		_, err := db.Exec(fmt.Sprintf("INSERT INTO users(id, name, username, gender) VALUES ('%s', '%s', '%s', '%s')", newUser.Id, newUser.Name, newUser.Username, newUser.Gender))
		if err != nil {
			panic(err.Error())
		}
	}
}
