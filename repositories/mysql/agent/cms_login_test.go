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
	var whatStaff uint64
	for _, row := range res {
		if row.LoginStatus == 1 {
			whatStaff = row.LoginId
			err := repo.Delete(row)
			if err != nil {
				t.Error(err)
				t.Fail()
				return
			}
			break
		}
	}
	staff, err := repo.Get(int64(whatStaff))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	if staff.LoginStatus == 1 {
		t.Fail()
		return
	}
	var otherStaffs []*entities.CmsLogin
	for _, row := range res {
		if row.LoginStatus == 1 {
			otherStaffs = append(otherStaffs, &entities.CmsLogin{StaffCode: row.StaffCode})
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
