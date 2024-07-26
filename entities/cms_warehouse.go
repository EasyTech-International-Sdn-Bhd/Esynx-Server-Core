package entities

type CmsWarehouse struct {
	Id        uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	WhName    string `xorm:"not null VARCHAR(200)" json:"whName,omitempty" xml:"whName"`
	WhCode    string `xorm:"not null unique VARCHAR(50)" json:"whCode,omitempty" xml:"whCode"`
	WhAddress string `xorm:"VARCHAR(1000)" json:"whAddress,omitempty" xml:"whAddress"`
	WhRemark  string `xorm:"VARCHAR(2000)" json:"whRemark,omitempty" xml:"whRemark"`
	WhStatus  int    `xorm:"not null default 1 INT" json:"whStatus,omitempty" xml:"whStatus"`
}

func (m *CmsWarehouse) TableName() string {
	return "cms_warehouse"
}
