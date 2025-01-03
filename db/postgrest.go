package db

import (
	"log"
	modelsdb "sms_server/models/models_db"
)

func (d *DB) PG_Get() {
	var u map[string]interface{}
	d.Pdb.Model(&modelsdb.User{}).First(&u)
	log.Println(u)
}
