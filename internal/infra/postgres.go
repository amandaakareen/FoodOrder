package infra

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connetc() *gorm.DB {
	var err error
	connStr := "user=admin password=foood dbname=food sslmode=disable"

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	return db
}
