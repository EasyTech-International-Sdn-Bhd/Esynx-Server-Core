package entities

type CmsWarehouse struct {
	Id        uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	WhName    string `xorm:"not null VARCHAR(200)"`
	WhCode    string `xorm:"not null unique VARCHAR(50)"`
	WhAddress string `xorm:"VARCHAR(1000)"`
	WhRemark  string `xorm:"VARCHAR(2000)"`
	WhStatus  int    `xorm:"not null default 1 INT"`
}

func (m *CmsWarehouse) TableName() string {
	return "cms_warehouse"
}
