package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type InvoiceDetailsWithProduct struct {
	D *entities.CmsInvoiceDetails
	P *entities.CmsProduct
}
