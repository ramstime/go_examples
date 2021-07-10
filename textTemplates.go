/*
Templates with Nested structs and slices
*/
package main

import (
	"log"
	"os"
	"text/template"
)

type email struct {
	Hotmail string
	Gmail   string
}
type Person struct {
	Name   string
	Country string
	Age int
	Emails []email
}

const tmpl = `  Person name is {{.Name}} and age is {{.Age }}. His Country is {{.Country}}
{{ range .Emails }}
    His hotemail id is "{{.Hotmail}}"
    His gmail id is {{.Gmail}}
{{end}}
`

func main() {
	person := Person{
		Name:   "ramstime",
		Country: "India",
		Age: 34,
		Emails: []email{{Hotmail: "rams.time@hotmail", Gmail: "rams.time@gmail.com"},},
	}

	t := template.New("Person template")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
