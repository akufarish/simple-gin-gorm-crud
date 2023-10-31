package models

type Murid struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
	Kelas string `gorm:"type:varchar(255)" json:"kelas"`
	Jurusan string `gorm:"type:varchar(255)" json:"jurusan"`
}