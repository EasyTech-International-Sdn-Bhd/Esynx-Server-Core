package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type CustomerWithBranches struct {
	C *entities.CmsCustomer
	B []*entities.CmsCustomerBranch
}

type CustomerWithAgent struct {
	C *entities.CmsCustomer
	A *entities.CmsLogin
}
