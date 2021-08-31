package model

type Kelime struct {
	KelimeId           int    `json:"kelime_id" gorm:"primaryKey;column:KelimeId"`
	Latince            string `json:"latince" gorm:"column:Latince"`
	Osmanlica          string `json:"osmanlica" gorm:"column:Osmanlica"`
	Aciklama           string `json:"aciklama" gorm:"column:Aciklama"`
	HataliLatince      string `json:"hatali_latince" gorm:"column:HataliLatince"`
	TemizHataliLatince string `json:"temiz_hatali_latince" gorm:"column:TemizHataliLatince"`
	TemizLatince       string `json:"temiz_latince" gorm:"column:TemizLatince"`
	TemizOsmanlica     string `json:"temiz_osmanlica" gorm:"column:TemizOsmanlica"`
	Seviye             int    `json:"seviye" gorm:"column:Seviye"`
	Sira               int    `json:"sira" gorm:"column:Sira"`
	Tur                int    `json:"tur" gorm:"column:Tur"`
	Kategori           string `json:"kategori" gorm:"column:Kategori"`
	SilindiMi          int    `json:"silindi_mi" gorm:"column:SilindiMi"`
	RastgeledeGetir    int    `json:"rastgelede_getir" gorm:"column:RastgeledeGetir"`
	OlusturmaTarihi    string `json:"olusturma_tarihi" gorm:"column:OlusturmaTarihi"`
	GuncellemeTarihi   string `json:"guncelleme_tarihi" gorm:"column:GuncellemeTarihi"`
}

func (Kelime) TableName() string {
	return "Kelime"
}
