package models

type Murid struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(255)" json:"nama" form:"nama"`
	Kelas string `gorm:"type:varchar(255)" json:"kelas" form:"kelas"`
	Jurusan string `gorm:"type:varchar(255)" json:"jurusan" form:"jurusan"`
	Img string `gorm:"type:varchar(255)" json:"img" form:"image"`
}