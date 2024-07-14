package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type InvoiceSalesWithItems struct {
	M *entities.CmsInvoiceSales
	D []*entities.CmsInvoiceDetails
}

type InvoiceSalesWithCustomer struct {
	I *entities.CmsInvoiceSales
	C *entities.CmsCustomer
}
