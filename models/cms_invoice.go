package models

import "github.com/easytech-international-sdn-bhd/core/entities"

type InvoiceWithItems struct {
	M *entities.CmsInvoice
	D []*entities.CmsInvoiceDetails
}

type InvoiceWithCustomer struct {
	I *entities.CmsInvoice
	C *entities.CmsCustomer
}
