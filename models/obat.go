package model

import "time"

type Obat struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	User_id   int       `gorm:"column:user_id"`
	Nama_Obat string    `gorm:"column:nama_obat"`
	Dosis     string    `gorm:"column:dosis"`
	Frekuensi string    `gorm:"column:frekuensi"`
	Catatan   string    `gorm:"column:catatan"`
	Tanggal   string    `gorm:"column:tanggal"`
	Waktu     string    `gorm:"column:waktu"`
	CreatedAt time.Time `gorm:"column:createdAt"`

	User User `gorm:"foreignKey:User_id"`
}

func (Obat) TableName() string {
	return "obat"
}
