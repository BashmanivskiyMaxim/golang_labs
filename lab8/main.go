package main

import (
	"fmt"
	"html/template"
	"lab8/models"
	"net/http"
)

var news map[string]*models.News

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("mainPage.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "mainPage", nil)

	fmt.Println(news)
}

func saveNewsHandler(w http.ResponseWriter, r *http.Request) {
	id := GenerateId()
	title := r.FormValue("title")
	content := r.FormValue("content")
	allContent := r.FormValue("contentAll")

	newObj := models.NewNews(id, title, content, allContent)
	news[newObj.Id] = newObj

	http.Redirect(w, r, "/news", 302)

}

func writeHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/writeNews.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "writeNews", nil)

}

func NewsHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/newsPage.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "newsPage", news)

}

func NewsOneHandler(w http.ResponseWriter, r *http.Request) {

	NewsOneId := r.FormValue("id")

	NewsOneEl := news[NewsOneId]
	fmt.Println(NewsOneEl)

	t, err := template.ParseFiles("templates/newsOnePage.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "newsOnePage", NewsOneEl)

}

func AboutUsHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/aboutUsPage.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "aboutUs", nil)

}
func PhotoHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/photoPage.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "photo", nil)

}

func main() {

	news = make(map[string]*models.News, 0)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", NewsHandler)
	http.HandleFunc("/saveNews", saveNewsHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/newsOne/", NewsOneHandler)
	http.HandleFunc("/about", AboutUsHandler)
	http.HandleFunc("/photo", PhotoHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	http.ListenAndServe(":8000", nil)
}
