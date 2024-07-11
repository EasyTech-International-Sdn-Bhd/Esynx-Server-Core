package models

import "github.com/easytech-international-sdn-bhd/core/entities"

type InvoiceDetailsWithProduct struct {
	D *entities.CmsInvoiceDetails
	P *entities.CmsProduct
}
