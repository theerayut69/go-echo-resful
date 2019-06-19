package main

import (
	"log"

	"github.com/labstack/echo"
)

type CustomerHandler struct {
	DB *gorm.DB
}

func (h *CustomerHandler) Initialize() {
	db, err := gorm.Open("mysql", "webservice:P@ssw0rd@tcp(127.0.0.1:3306)/db_webservice?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})

	h.DB = db
}

func main() {
	e := echo.New()

	h := CustomerHandler{}
	h.Initialize()
	e.GET("/customers", h.GetAllCustomer)
	e.POST("/customers", h.SaveCustomer)
	e.GET("/customers/:id", h.GetCustomer)
	e.PUT("/customers/:id", h.UpdateCustomer)
	e.DELETE("/customers/:id", h.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
