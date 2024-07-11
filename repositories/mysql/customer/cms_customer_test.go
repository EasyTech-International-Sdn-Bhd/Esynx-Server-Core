package customer

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/core/entities"
	"math/rand/v2"
	"testing"
	"xorm.io/xorm"
)

func TestCmsCustomerRepository_InsertBatch(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(db)
	var newCustomers []entities.CmsCustomer
	for i := 0; i < 10; i++ {
		newCustomers = append(newCustomers, entities.CmsCustomer{
			CustCode:        fmt.Sprintf("CS-%d", rand.IntN(100000)+1),
			CustCompanyName: fmt.Sprintf("CS-Name-%d", rand.IntN(100)+1),
			CustEmail:       fmt.Sprintf("CS-Email-%d", rand.IntN(100)+1),
			CustomerStatus:  rand.IntN(2),
		})
	}
	err = repo.InsertBatch(newCustomers)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCmsCustomerRepository_SearchByNameOrCode(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(db)
	cmsCustomers, err := repo.SearchByNameOrCode("CS-")
	if err != nil {
		t.Fail()
		t.Error(err)
		return
	}
	if len(cmsCustomers) == 0 {
		t.Fail()
	} else {
		t.Logf("%v", cmsCustomers)
	}
}

func TestCmsCustomerRepository_Update(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(db)
	err = repo.Update(&entities.CmsCustomer{
		CustCode:       fmt.Sprintf("CS-%d", 111),
		CustomerStatus: 2,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func dbConn() (db *xorm.Engine, err error) {
	conn := fmt.Sprintf("root:mysql@tcp(127.0.0.1:3306)/easysale_elk?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s")
	return xorm.NewEngine("mysql", conn)
}
