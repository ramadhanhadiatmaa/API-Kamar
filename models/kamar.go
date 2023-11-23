package models

import "time"

type Kamar struct {
	Id          string    `gorm:"type:varchar(20); primaryKey" json:"id"`
	Markusnicu  int    	  `json:"markusnicu"`
	Markusvvip  int    	  `json:"markusvvip"`
	Markusvip   int    	  `json:"markusvip"`
	Lukas       int    	  `json:"lukas"`
	Maria       int    	  `json:"maria"`
	Fransiskus  int    	  `json:"fransiskus"`
	Matius      int    	  `json:"matius"`
	Teresia     int    	  `json:"teresia"`
	Teresiatiga int    	  `json:"teresiatiga"`
	Yosef       int    	  `json:"yosef"`
	Klara       int    	  `json:"klara"`
	Egidio      int    	  `json:"egidio"`
	Yohanes     int    	  `json:"yohanes"`
	Total     	int    	  `json:"total"`
	Tersedia    int    	  `json:"tersedia"`
	Terisi    int    	  `json:"terisi"`
	Updated   time.Time `json:"updated_at"`
}