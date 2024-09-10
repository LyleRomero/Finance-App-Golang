package postgres

import (
	"fmt"
	"sync"
	"time"

	"github.com/lenovoz460/GoMenu/Pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var onceDBLoad sync.Once
var tables = []interface{}{
	&entity.User{},
}

func init() {
	connect()
}

func connect() *gorm.DB {
	onceDBLoad.Do(func() {
		source := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			"localhost",
			"root",
			"root",
			"gomenu",
			"5432",
		)
		var i int
		for {
			var err error
			if i >= 10 {
				panic("Failed to connect: " + source)
			}
			time.Sleep(3 * time.Second)
			db, err = gorm.Open(postgres.Open(source), &gorm.Config{})
			if err != nil {
				log.Info("Retrying connect....", err)
				i++
				continue
			}
			break
		}
		migrate()
		log.Info("Connect to db!")
	})
	return db
}

func migrate() {
	for _, table := range tables {
		db.AutoMigrate(table)
	}
}
