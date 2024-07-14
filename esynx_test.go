package core

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/test"
	"testing"
)

func TestNewESynxCore(t *testing.T) {
	session := test.NewTestSession()
	provider, err := NewEsynxProvider(session)
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
