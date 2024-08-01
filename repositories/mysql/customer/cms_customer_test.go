package customer

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/test"
	"math/rand/v2"
	"testing"
	"xorm.io/xorm"
)

func TestCmsCustomerRepository_InsertMany(t *testing.T) {
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(option)
	var newCustomers []*entities.CmsCustomer
	for i := 0; i < 10; i++ {
		newCustomers = append(newCustomers, &entities.CmsCustomer{
			CustCode:        fmt.Sprintf("CS-%d", rand.IntN(100000)+1),
			CustCompanyName: fmt.Sprintf("CS-Name-%d", rand.IntN(100)+1),
			CustEmail:       fmt.Sprintf("CS-Email-%d", rand.IntN(100)+1),
			CustomerStatus:  rand.IntN(2),
		})
	}
	err = repo.InsertMany(newCustomers)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCmsCustomerRepository_SearchByNameOrCode(t *testing.T) {
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(option)
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
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerRepository(option)
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
