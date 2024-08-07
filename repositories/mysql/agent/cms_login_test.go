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
	cond := builder.Select("*")
	repo := NewCmsLoginRepository(option)
	res, err := repo.Find(cond)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Logf("TestResolve:GetAll %v", len(res))
}
