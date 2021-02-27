package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex = template.New("")
)

// เวลาที่ package อื่น import มา จะเรียกใช้ฟังก์ชั่น init
func init() {
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("template/root.html", "template/index.html")
	if err != nil {
		panic(err)
	}
	// ตั้ง template ที่ define root ให้เป็นตัวหลัก
	tpIndex = tpIndex.Lookup("root")
}

func render(t *template.Template, w http.ResponseWriter, date interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=uft-8")
	err := t.Execute(w, date)
	if err != nil {
		log.Println(err)
	}
}

// Index render index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
