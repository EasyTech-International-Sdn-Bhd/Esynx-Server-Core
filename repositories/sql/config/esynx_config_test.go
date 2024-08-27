package config

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/test"
	"testing"
	"time"
	"xorm.io/builder"
)

func TestEsynxConfig(t *testing.T) {
	option, err := test.TestOption()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	cond := builder.Select("*")
	repo := NewEsynxConfigRepository(option)

	err = repo.Insert(&entities.EsynxConfig{
		Name:      "---",
		ServiceId: "----",
		Config:    "-----",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	if err != nil {
		t.Fatal(err)
	}

	configs, err := repo.Find(cond)
	if err != nil {
		t.Fatal(err)
	}
	if len(configs) == 0 {
		t.Fail()
	}
	err = repo.Delete(configs[0].ServiceId)
	if err != nil {
		t.Fatal(err)
	}
}
