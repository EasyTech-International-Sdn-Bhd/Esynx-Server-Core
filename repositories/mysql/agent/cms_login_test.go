package agent

import (
	"fmt"
	"testing"
	"xorm.io/xorm"
)

func TestResolve(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsLoginRepository(db)
	res, err := repo.GetAll()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Logf("TestResolve:GetAll %v", len(res))
}

func dbConn() (db *xorm.Engine, err error) {
	conn := fmt.Sprintf("root:mysql@tcp(127.0.0.1:3306)/easysale_vit?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s")
	return xorm.NewEngine("mysql", conn)
}
