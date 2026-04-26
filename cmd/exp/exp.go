package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	//favouriteFoodList := []string{"KFC", "Sushi", "Kamboya"}
	//user := User{
	//	Name: "Darek Vu",
	//	Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	//	Age:  29,
	//}
	//age := 29
	//username := "Darek"

	favouriteFood := []string{"Pizza", "burger", "Pasta"}

	err = t.Execute(os.Stdout, favouriteFood)
	if err != nil {
		panic(err)

	}

}
