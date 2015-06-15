package main

import(
    "github.com/astaxie/beego"
    "time"
    "fmt"
)

func init(){
    beego.AddFuncMap("justdate",Date2HTML)
}

func Date2HTML(t time.Time)string{
    y,m,d := t.Date()
    return fmt.Sprintf("%d-%d-%d",y,m,d)
}
