package core

import (
	"fmt"
	"testing"
)

func TestNewESynxCore(t *testing.T) {
	conn := fmt.Sprintf("root:mysql@tcp(127.0.0.1:3306)/easysale_elk?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s")
	provider, err := NewEsynxProvider(MySQL, conn)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = provider.Destroy()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}
