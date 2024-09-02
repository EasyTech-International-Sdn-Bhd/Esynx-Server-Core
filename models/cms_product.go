package models

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ProductWithDetails struct {
	P *entities.CmsProduct
	A []*entities.CmsProductAtch
	I []*entities.CmsProductImage
	U []*entities.CmsProductUomPrice
	T []*entities.CmsProductPrice
	B []*entities.CmsProductBatch
	W []*entities.CmsWarehouseStock
}
