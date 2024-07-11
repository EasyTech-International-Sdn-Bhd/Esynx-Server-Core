package models

import "github.com/easytech-international-sdn-bhd/core/entities"

type CustomerWithBranches struct {
	C *entities.CmsCustomer
	B []*entities.CmsCustomerBranch
}

type CustomerWithAgents struct {
	C *entities.CmsCustomer
	A *entities.CmsLogin
}
