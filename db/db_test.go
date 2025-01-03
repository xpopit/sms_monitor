package db

import (
	"sms_server/cls"
	"testing"
)

func TestRDB(t *testing.T) {
	cls.LoadENV()
	db := Init()
	db.PG_Get()

	// fmt.Println(rdb.Get("d"))
	// rdb.Pub()
	// rdb.G
	// rdb.Get()
	// rdb.Pub()
}
