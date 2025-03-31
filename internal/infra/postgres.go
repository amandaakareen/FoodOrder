package infra

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connetc() {
	var err error
	connStr := "user=admin password=foood dbname=food sslmode=disable"

	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}
}
