package handler

import (
	"log"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)
type CustomerHandler struct {
	DB *gorm.DB
}

func (h *CustomerHandler) Initialize() {
	db, err := gorm.Open("mysql", "webservice:P@ssw0rd@tcp(127.0.0.1:3306)/db_webservice?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})

	h.DB = db
}
