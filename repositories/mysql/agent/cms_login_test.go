package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/test"
	"testing"
	"xorm.io/builder"
)

func TestResolve(t *testing.T) {
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	repo := NewCmsLoginRepository(option)
	res, err := repo.Find(builder.Like{"staff_code", "alor"})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Logf("TestResolve:GetAll %v", len(res))
}
