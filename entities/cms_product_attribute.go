package entities

import (
	"time"
)

type CmsProductAttribute struct {
	ProductAttributeId int       `xorm:"not null pk autoincr INT"`
	ProductCode        string    `xorm:"unique(unq_key) VARCHAR(50)"`
	AttributeName      string    `xorm:"comment('the name can be repeat here (e.g Color)') unique(unq_key) VARCHAR(50)"`
	AttributeValue     string    `xorm:"comment('the value should not be repeat,for example, if the attribute name is Color, the value of color is Red, Brown, Blue, Black and etc. please refer to data sample in this table') unique(unq_key) VARCHAR(50)"`
	AvailableQty       int       `xorm:"INT"`
	SequenceNo         int       `xorm:"INT"`
	UpdatedAt          time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsProductAttribute) TableName() string {
	return "cms_product_attribute"
}
