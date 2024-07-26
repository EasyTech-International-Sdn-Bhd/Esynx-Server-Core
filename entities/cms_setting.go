package entities

import (
	"time"
)

type CmsSetting struct {
	PoHeaderUrl           string    `xorm:"VARCHAR(500)" json:"poHeaderUrl,omitempty" xml:"poHeaderUrl"`
	PoFooterUrl           string    `xorm:"VARCHAR(500)" json:"poFooterUrl,omitempty" xml:"poFooterUrl"`
	LogoUrl               string    `xorm:"VARCHAR(200)" json:"logoUrl,omitempty" xml:"logoUrl"`
	ColorCode             string    `xorm:"VARCHAR(100)" json:"colorCode,omitempty" xml:"colorCode"`
	Currency              string    `xorm:"VARCHAR(3)" json:"currency,omitempty" xml:"currency"`
	ProductLastupdate     time.Time `xorm:"comment('when admin/officer update product table (except product iamge), the date will be updated.') DATETIME" json:"productLastupdate,omitempty" xml:"productLastupdate"`
	GeneralinfoLastupdate time.Time `xorm:"comment('when admin/officer update setting table, the date will be updated') DATETIME" json:"generalinfoLastupdate,omitempty" xml:"generalinfoLastupdate"`
	AboutusInfo           string    `xorm:"VARCHAR(500)" json:"aboutusInfo,omitempty" xml:"aboutusInfo"`
	TermsconditionsInfo   string    `xorm:"VARCHAR(500)" json:"termsconditionsInfo,omitempty" xml:"termsconditionsInfo"`
	UserguideInfo         string    `xorm:"VARCHAR(500)" json:"userguideInfo,omitempty" xml:"userguideInfo"`
	NewupdateInfo         string    `xorm:"VARCHAR(500)" json:"newupdateInfo,omitempty" xml:"newupdateInfo"`
	Faq                   string    `xorm:"VARCHAR(500)" json:"faq,omitempty" xml:"faq"`
	InvPrefix             string    `xorm:"default 'IV-1001' VARCHAR(22)" json:"invPrefix,omitempty" xml:"invPrefix"`
	CnPrefix              string    `xorm:"default 'CN-1001' VARCHAR(22)" json:"cnPrefix,omitempty" xml:"cnPrefix"`
	CoPrefix              string    `xorm:"default 'CO-1001' VARCHAR(22)" json:"coPrefix,omitempty" xml:"coPrefix"`
	CashPrefix            string    `xorm:"default 'CS-1001' VARCHAR(22)" json:"cashPrefix,omitempty" xml:"cashPrefix"`
	SoPrefix              string    `xorm:"default 'SO-1001' VARCHAR(22)" json:"soPrefix,omitempty" xml:"soPrefix"`
	QtPrefix              string    `xorm:"default 'QT-1001' VARCHAR(22)" json:"qtPrefix,omitempty" xml:"qtPrefix"`
	OrPrefix              string    `xorm:"default 'OR-1001' VARCHAR(22)" json:"orPrefix,omitempty" xml:"orPrefix"`
	DoPrefix              string    `xorm:"default 'DO-1001' VARCHAR(22)" json:"doPrefix,omitempty" xml:"doPrefix"`
	SelforderId           string    `xorm:"default 'SO-135243' VARCHAR(22)" json:"selforderId,omitempty" xml:"selforderId"`
	SvOrder               string    `xorm:"default '23.001' VARCHAR(22)" json:"svOrder,omitempty" xml:"svOrder"`
}

func (m *CmsSetting) TableName() string {
	return "cms_setting"
}
