package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type ProductWithDetails struct {
	P *entities.CmsProduct
	A []*entities.CmsProductAtch
	I []*entities.CmsProductImage
	U []*entities.CmsProductUomPriceV2
	T []*entities.CmsProductPriceV2
	B []*entities.CmsProductBatch
	W []*entities.CmsWarehouseStock
}
