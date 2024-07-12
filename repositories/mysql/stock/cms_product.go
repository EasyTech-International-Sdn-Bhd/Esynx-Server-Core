package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"strings"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type CmsProductRepository struct {
	db *xorm.Engine
	a  *CmsProductAtchRepository
	b  *CmsProductBatchRepository
	i  *CmsProductImageRepository
	t  *CmsProductPriceTagRepository
	p  *CmsProductUomPriceRepository
	w  *CmsWarehouseStockRepository
}

func NewCmsProductRepository(db *xorm.Engine) *CmsProductRepository {
	return &CmsProductRepository{
		db: db,
		a:  NewCmsProductAtchRepository(db),
		b:  NewCmsProductBatchRepository(db),
		i:  NewCmsProductImageRepository(db),
		t:  NewCmsProductPriceTagRepository(db),
		p:  NewCmsProductUomPriceRepository(db),
		w:  NewCmsWarehouseStockRepository(db),
	}
}

func (r *CmsProductRepository) Get(productCode string) (*entities.CmsProduct, error) {
	var cmsProduct entities.CmsProduct
	has, err := r.db.Where("product_code = ?", productCode).Get(&cmsProduct)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsProduct, nil
}

func (r *CmsProductRepository) GetMany(productCodes []string) ([]*entities.CmsProduct, error) {
	var cmsProducts []*entities.CmsProduct
	err := r.db.In("product_code", productCodes).Find(&cmsProducts)
	if err != nil {
		return nil, err
	}
	return cmsProducts, nil
}

func (r *CmsProductRepository) Search(predicate string) ([]*entities.CmsProduct, error) {
	var records []*entities.CmsProduct
	tokens := strings.Split(predicate, " ")
	var where []builder.Cond
	for _, token := range tokens {
		where = append(where, builder.Like{"product_code", token})
		where = append(where, builder.Like{"product_name", token})
	}
	err := r.db.Where(builder.Or(where...)).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductRepository) InsertMany(records []*entities.CmsProduct) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProduct) *entities.CmsProduct {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductRepository) Update(record *entities.CmsProduct) error {
	_, err := r.db.Where("product_code = ?", record.ProductCode).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductRepository) UpdateMany(records []*entities.CmsProduct) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, product := range records {
		product.Validate()
		product.ToUpdate()
		_, err = session.Where("product_code = ?", product.ProductCode).Update(product)
		if err != nil {
			rollback = true
			sessionErr = err
			break
		}
	}
	if rollback {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return sessionErr
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductRepository) Delete(record *entities.CmsProduct) error {
	record.ProductStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductRepository) DeleteMany(records []*entities.CmsProduct) error {
	for _, record := range records {
		record.ProductStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}

func (r *CmsProductRepository) GetWithDetails(productCode string) (*models.ProductWithDetails, error) {
	p, err := r.Get(productCode)
	if err != nil {
		return nil, err
	}
	a, err := r.a.Get(productCode)
	if err != nil {
		return nil, err
	}
	b, err := r.b.Get(productCode)
	if err != nil {
		return nil, err
	}
	i, err := r.i.Get(strconv.Itoa(p.ProductId))
	if err != nil {
		return nil, err
	}
	t, err := r.t.Get(productCode)
	if err != nil {
		return nil, err
	}
	pr, err := r.p.Get(productCode)
	if err != nil {
		return nil, err
	}
	w, err := r.w.Get(productCode)
	return &models.ProductWithDetails{
		P: p,
		A: a,
		I: i,
		U: pr,
		T: t,
		B: b,
		W: w,
	}, nil
}
