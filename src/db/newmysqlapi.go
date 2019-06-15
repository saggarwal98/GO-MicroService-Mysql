package db

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql" //sql database
)

//Article contains the required fields whicwill be used for database
type Article struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
}

//Displayarticles will display every row from database
func Displayarticles() string {
	var db, err = sql.Open("mysql", "saggarwal98:shubham@tcp(127.0.0.1:3306)/NewMysqlApi")
	defer db.Close()
	results, err := db.Query("Select * from Articles")
	flag := false
	var str []string
	if results.Next() == true {
		flag = true
		var a Article
		err = results.Scan(&a.ID, &a.Title, &a.Description, &a.Price)
		if err != nil {
			log.Println(err.Error())
			return "got error"
		}
		str = append(str, "ID:`"+strconv.Itoa(a.ID)+"` Title:`"+a.Title+"` Description:`"+a.Description+"` Price:`"+strconv.Itoa(a.Price)+"`")
		for results.Next() {
			var b Article
			err = results.Scan(&b.ID, &b.Title, &b.Description, &b.Price)
			if err != nil {
				log.Println(err.Error())
				return "got error"
			}
			str = append(str, "ID:`"+strconv.Itoa(a.ID)+"` Title:`"+a.Title+"` Description:`"+a.Description+"` Price:`"+strconv.Itoa(a.Price)+"`")
		}

	}
	if flag == false {
		return "Nothing to show"
	}
	st := strings.Join(str, "\n")
	return st
}
//Displayarticlebyid will display row starting with that id
func Displayarticlebyid(id string) string {
	var db, err = sql.Open("mysql", "saggarwal98:shubham@tcp(127.0.0.1:3306)/NewMysqlApi")
	defer db.Close()
	flag := false
	var str string
	results, err := db.Query("Select * from Articles where ID=" + id)
	if results != nil {

		if results.Next() == true {
			flag = true
			var a Article
			err = results.Scan(&a.ID, &a.Title, &a.Description, &a.Price)
			if err != nil {
				log.Println(err.Error())
				return "got error"
			}
			str = "ID:`" + strconv.Itoa(a.ID) + "` Title:`" + a.Title + "` Description:`" + a.Description + "` Price:`" + strconv.Itoa(a.Price) + "`"
		}
	}

	if flag == false {
		return "Nothing found with that id"
	}
	return str
}

//CreateArticles will create a new row in database
func CreateArticles(id int, title string, description string, price int) string {
	key := strconv.Itoa(id)
	key3 := strconv.Itoa(price)
	var db, err = sql.Open("mysql", "saggarwal98:shubham@tcp(127.0.0.1:3306)/NewMysqlApi")
	defer db.Close()
	_, err = db.Query("INSERT INTO Articles VALUES('" + key + "','" + title + "','" + description + "','" + key3 + "')")
	if err != nil {
		log.Println(err.Error())
		return "Could not create Article"
	}
	return "created new article"
}

//Deletearticles will delete row on the basis of title
func Deletearticles(title string) string {
	var db, err = sql.Open("mysql", "saggarwal98:shubham@tcp(127.0.0.1:3306)/NewMysqlApi")
	defer db.Close()
	flag := false
	results, err := db.Query("Select * from Articles")
	for results.Next() {
		var a Article
		err = results.Scan(&a.ID, &a.Title, &a.Description, &a.Price)
		if err != nil {
			log.Print(err.Error())
		}
		if a.Title == title {
			flag = true
			_, err := db.Query("Delete from Articles where Title='" + a.Title + "'")
			if err != nil {
				log.Println(err.Error())
				return "error received"
			}
		}
	}
	if flag == false {
		return "No article found with that title"
	}
	return "Article Deleted"
}

//Updatearticles is used to update the row with that id
func Updatearticles(id int, description string) string {
	var db, err = sql.Open("mysql", "saggarwal98:shubham@tcp(127.0.0.1:3306)/NewMysqlApi")
	defer db.Close()
	flag := false
	results, err := db.Query("Select * from Articles")
	for results.Next() {
		var a Article
		err = results.Scan(&a.ID, &a.Title, &a.Description, &a.Price)
		if err != nil {
			log.Print(err.Error())
			return "got error"
		}
		if a.ID == id {
			flag = true
			_, err := db.Query("update Articles set Description='" + description + "' where ID='" + strconv.Itoa(id) + "'")
			if err != nil {
				log.Println(err.Error())
				return "got error"
			}
		}
	}
	if flag == false {
		return "No article found with that id"
	}
	return "Article Updated"
}
