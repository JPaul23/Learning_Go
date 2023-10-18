package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main (){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
}
// Migrate the schema
db.AutoMigrate(&Product{})

// Create
db.Create(&Product{Code: "D42", Price: 100})

// Read
var product Product
db.First(&product, 1)
db.First(&product, "code = ?", "D42")

//Update
db.Model(&product).Update("Price", 200);

// update multiple fields
db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)
}