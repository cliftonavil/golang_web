// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	db, err := gorm.Open("postgres", "user=postgres password=clifton dbname=gorm sslmode=disable")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	fmt.Println("Connected to Database!!")

// 	// err = db.DB().Ping()  //Checking for Error
// 	// fmt.Println(err)

// 	// db.SingularTable(false) // Table Name with 's

// 	db.Debug().DropTableIfExists(&Owner{}, &Book{}, &Author{})
// 	db.Debug().CreateTable(&Owner{}, &Book{}, &Author{})

// 	owner := Owner{
// 		FirstName: "CLifton",
// 		LastName:  "Avil",
// 	}
// 	db.Debug().Create(&owner)
// 	owner.LastName = "Dsouza" // Update useing current object
// 	db.Debug().Save(&owner)   // save

// 	// db.Debug().Delete(&owner) // Delete Record useing current  object

// 	var o Owner
// 	db.Debug().Find(&o)
// 	fmt.Println(o.FirstName, o.LastName)
// }

// type Owner struct {
// 	gorm.Model
// 	FirstName string
// 	LastName  string
// 	Books     []Book
// }

// type Book struct {
// 	gorm.Model
// 	Name        string
// 	PublishDate time.Time
// 	OwnerID     uint     `sql:"index"`
// 	Authors     []Author `gorm:"many2many:books_authors"`
// }

// type Author struct {
// 	gorm.Model
// 	FirstName string
// 	LastName  string
// }
