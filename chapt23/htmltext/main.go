package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	singleTemplate()

	fmt.Println("\nMultiple templates")
	multTemplates()

	fmt.Println("\nLoad multiple in one:")
	loadMultiplInOne()

	fmt.Println("\nEnumerate loaded templates:")
	enumLoadedTemplates()

	fmt.Println("\nLookup specific template:")
	lookupTemplate()

}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func lookupTemplate() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)
		os.Stdout.WriteString("\n")
	}
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}

func enumLoadedTemplates() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		for _, t := range allTemplates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v %v", err.Error())
	}
}

func loadMultiplInOne() {
	allTemplates, err1 := template.ParseFiles("templates/template.html",
		"templates/extras.html")
	if err1 == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error: %v %v", err1.Error())
	}
}

func multTemplates() {
	t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")
	if err1 == nil && err2 == nil {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t2.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
	} else {
		Printfln("Error: %v %v", err1.Error(), err2.Error())
	}
}

func singleTemplate() {
	t, err := template.ParseFiles("templates/template.html")
	if err == nil {
		t.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
