package entities

import (
	"time"
)

type CmsProductAttribute struct {
	ProductAttributeId uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"productAttributeId,omitempty" xml:"productAttributeId"`
	ProductCode        string    `xorm:"unique(unq_key) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	AttributeName      string    `xorm:"comment('the name can be repeat here (e.g Color)') unique(unq_key) VARCHAR(50)" json:"attributeName,omitempty" xml:"attributeName"`
	AttributeValue     string    `xorm:"comment('the value should not be repeat,for example, if the attribute name is Color, the value of color is Red, Brown, Blue, Black and etc. please refer to data sample in this table') unique(unq_key) VARCHAR(50)" json:"attributeValue,omitempty" xml:"attributeValue"`
	AvailableQty       int       `xorm:"INT" json:"availableQty,omitempty" xml:"availableQty"`
	SequenceNo         int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	UpdatedAt          time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsProductAttribute) TableName() string {
	return "cms_product_attribute"
}
