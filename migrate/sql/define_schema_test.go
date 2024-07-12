package sql

import (
	"fmt"
	"testing"
	"xorm.io/xorm"
)

func TestDefineSchema(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer db.Close()
	err = DefineSchema(db)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func dbConn() (db *xorm.Engine, err error) {
	conn := fmt.Sprintf("root:mysql@tcp(127.0.0.1:3306)/test_define_schema?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s")
	return xorm.NewEngine("mysql", conn)
}
