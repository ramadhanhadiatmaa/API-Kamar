package models

import "time"

type Kamar struct {
	Id          string    `gorm:"type:varchar(20); primaryKey" json:"id"`
	MarkusNicu  int    `json:"markus_nicu"`
	MarkusVvip  int    `json:"markus_vvip"`
	MarkusVip   int    `json:"markus_vip"`
	Lukas       int    `json:"lukas"`
	Maria       int    `json:"maria"`
	Fransiskus  int    `json:"fransiskus"`
	Matius      int    `json:"matius"`
	Teresia     int    `json:"teresia"`
	TeresiaTiga int    `json:"teresia_tiga"`
	Yosef       int    `json:"yosef"`
	Klara       int    `json:"klara"`
	Egidio      int    `json:"egidio"`
	Yohanes     int    `json:"yohanes"`
	UpdatedAt   time.Time `json:"updated_at"`
}