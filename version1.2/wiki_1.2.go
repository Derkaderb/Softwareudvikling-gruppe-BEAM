
package main

import (
    "html/template"
    "io/ioutil"
    "net/http"
    "log"
    "os"
    "os/exec"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := "brugersvar" + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func fsharp(fsharpnavn string) {
  //  path, _ := os.Getwd()

    err := exec.Command("mono", fsharpnavn +".exe").Run()

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/quest/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func questHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/quest/"):]
    fsharp("surip")
    p, err := loadPage("question")
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "quest", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+"test", http.StatusFound)
}


func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/quest/", questHandler)
    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}