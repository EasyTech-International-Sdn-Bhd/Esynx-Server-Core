package customer

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/test"
	"math/rand/v2"
	"testing"
)

func TestCmsCustomerBranchRepository_InsertMany(t *testing.T) {
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	repo := NewCmsCustomerBranchRepository(option)
	var branches []*entities.CmsCustomerBranch
	for i := 0; i < 10; i++ {
		branches = append(branches, &entities.CmsCustomerBranch{
			BranchCode:   fmt.Sprintf("BR-%d", i),
			CustCode:     fmt.Sprintf("CS-%d", rand.IntN(100000)+1),
			BranchActive: rand.IntN(2),
		})
	}
	err = repo.InsertMany(branches)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
