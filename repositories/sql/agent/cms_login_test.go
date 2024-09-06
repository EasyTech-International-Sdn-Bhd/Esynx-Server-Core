package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
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
	var bulkUpdate []*entities.CmsLogin
	for _, row := range res {
		if row.LoginId == 35 {
			err := repo.Delete(row)
			if err != nil {
				t.Error(err)
				t.Fail()
				return
			}
		}
		up := entities.CmsLogin{
			AgentCode: row.AgentCode,
			Password:  "********",
			Login:     row.AgentCode,
		}
		bulkUpdate = append(bulkUpdate, &up)
	}
	if len(bulkUpdate) > 0 {
		err := repo.UpdateMany(bulkUpdate)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
	}

	var otherStaffs []*entities.CmsLogin
	for _, row := range res {
		if row.LoginStatus == 1 {
			otherStaffs = append(otherStaffs, &entities.CmsLogin{AgentCode: row.AgentCode})
		}
	}
	err = repo.DeleteMany(otherStaffs)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	rows, err := repo.GetAll()
	for _, row := range rows {
		if row.LoginStatus == 1 {
			t.Fail()
			return
		}
	}
}
