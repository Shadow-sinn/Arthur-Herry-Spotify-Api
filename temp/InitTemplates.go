package templates

import (
    "fmt"
    "html/template"
    "os"
)

// Variable global que l'on appelle quand on veut touché aux templates(ex: ExecuteTemplate)
var Temp *template.Template

func InitTemp() {
    //Sert a initié les templates, si non, renvoie une erreur
    temp, errTemp := template.ParseGlob("./temp/*.html")
    if errTemp != nil {
        fmt.Printf("Oupss une erreur liée au Templates : %v", errTemp.Error())
        os.Exit(1)
    }
    Temp = temp
}