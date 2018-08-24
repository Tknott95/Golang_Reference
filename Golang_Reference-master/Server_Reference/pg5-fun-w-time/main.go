package main

import ( 
    "log"
    "os"
    "text/template"
    "time"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func kitchenTime( t time.Time) string {
    return t.Format(time.Kitchen)
}

func monthDayYear( t time.Time) string {
    return t.Format("04-02-2010")
}

var fm = template.FuncMap{
    "fdateMDY": monthDayYear,
    "kitchenTimeByTrev": kitchenTime,
}

func main() {
    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
    if err != nil {
        log.Fatalln(err)
    }
}
