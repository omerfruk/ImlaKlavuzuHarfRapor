package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main.go/model"
	"sort"
	"strings"
)

func main() {
	var KelimeAdedi int64
	db, err := gorm.Open(sqlite.Open("./imlakilavuzu3.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	hamVeri := make([]model.Kelime, 0)
	db.AutoMigrate(&model.Kelime{})
	if err = db.Model(&model.Kelime{}).Count(&KelimeAdedi).Find(&hamVeri).Error; err != nil {
		fmt.Println(err)
	}
	kelimeler := make(map[string]int, 0)
	toplam := make(map[string]int, 0)

	for _, kelime := range hamVeri {
		kelime.TemizOsmanlica = strings.ReplaceAll(kelime.TemizOsmanlica, "/", " ")
		kelime.TemizOsmanlica = strings.ReplaceAll(kelime.TemizOsmanlica, "-", " ")
		data := strings.Split(kelime.TemizOsmanlica, "")
		for harfIndex, harf := range data {
			kelimeler[fmt.Sprintf("%d-%s", harfIndex, harf)] = kelimeler[fmt.Sprintf("%d-%s", harfIndex, harf)] + 1
			toplam[harf] = toplam[harf] + 1
		}
	}
	var f *excelize.File
	f = excelize.NewFile()
	sheetIsmi := "Harf Raporu"
	f.NewSheet("Harf Raporu")

	err = f.SetCellValue(sheetIsmi, "A1", "Harf")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = f.SetCellValue(sheetIsmi, "B1", "Toplam Kaç Tane")
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 67; i < 86; i++ {
		err = f.SetCellValue(sheetIsmi, fmt.Sprintf("%s1", string(i)), fmt.Sprintf("%d%s", i-66, ".Sırada"))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	for key, _ := range kelimeler {
		harf := strings.Split(key, "-")[1]
		toplam[harf] = toplam[harf] + 1
	}
	keys := make([]string, 0, len(toplam))
	for k := range toplam {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	temp := 2
	for _, value := range keys {
		f.SetCellValue(sheetIsmi, fmt.Sprintf("A%d", temp), value)
		f.SetCellValue(sheetIsmi, fmt.Sprintf("B%d", temp), toplam[value])

		for i := 1; i < 20; i++ {
			f.SetCellValue(sheetIsmi, fmt.Sprintf("%s%d", string(67+i-1), temp), kelimeler[fmt.Sprintf("%d-%s", i, value)])
			fmt.Println(fmt.Sprintf("%s%d", string(67+i), temp))
		}
		temp++

	}

	f.DeleteSheet("Sheet1")

	if err = f.SaveAs("Rapor.xlsx"); err != nil {
		fmt.Println(err)
	}

}
