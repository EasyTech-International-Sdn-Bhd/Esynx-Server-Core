package entities

import (
	"time"
)

type CmsSetting struct {
	PoHeaderUrl           string    `xorm:"VARCHAR(500)"`
	PoFooterUrl           string    `xorm:"VARCHAR(500)"`
	LogoUrl               string    `xorm:"VARCHAR(200)"`
	ColorCode             string    `xorm:"VARCHAR(100)"`
	Currency              string    `xorm:"VARCHAR(3)"`
	ProductLastupdate     time.Time `xorm:"comment('when admin/officer update product table (except product iamge), the date will be updated.') DATETIME"`
	GeneralinfoLastupdate time.Time `xorm:"comment('when admin/officer update setting table, the date will be updated') DATETIME"`
	AboutusInfo           string    `xorm:"VARCHAR(500)"`
	TermsconditionsInfo   string    `xorm:"VARCHAR(500)"`
	UserguideInfo         string    `xorm:"VARCHAR(500)"`
	NewupdateInfo         string    `xorm:"VARCHAR(500)"`
	Faq                   string    `xorm:"VARCHAR(500)"`
	InvPrefix             string    `xorm:"default 'IV-1001' VARCHAR(22)"`
	CnPrefix              string    `xorm:"default 'CN-1001' VARCHAR(22)"`
	CoPrefix              string    `xorm:"default 'CO-1001' VARCHAR(22)"`
	CashPrefix            string    `xorm:"default 'CS-1001' VARCHAR(22)"`
	SoPrefix              string    `xorm:"default 'SO-1001' VARCHAR(22)"`
	QtPrefix              string    `xorm:"default 'QT-1001' VARCHAR(22)"`
	OrPrefix              string    `xorm:"default 'OR-1001' VARCHAR(22)"`
	DoPrefix              string    `xorm:"default 'DO-1001' VARCHAR(22)"`
	SelforderId           string    `xorm:"default 'SO-135243' VARCHAR(22)"`
	SvOrder               string    `xorm:"default '23.001' VARCHAR(22)"`
}

func (m *CmsSetting) TableName() string {
	return "cms_setting"
}
