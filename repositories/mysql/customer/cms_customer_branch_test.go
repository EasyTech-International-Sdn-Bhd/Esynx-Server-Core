package customer

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/core/entities"
	"math/rand/v2"
	"testing"
)

func TestCmsCustomerBranchRepository_InsertMany(t *testing.T) {
	db, err := dbConn()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerBranchRepository(db)
	var branches []*entities.CmsCustomerBranch
	for i := 0; i < 10; i++ {
		branches = append(branches, &entities.CmsCustomerBranch{
			BranchCode:   fmt.Sprintf("BR-%d", i),
			CustCode:     fmt.Sprintf("CS-%d", rand.IntN(100000)+1),
			CustId:       rand.IntN(100),
			BranchActive: rand.IntN(2),
		})
	}
	err = repo.InsertMany(branches)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
