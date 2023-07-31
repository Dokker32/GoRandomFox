package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Fox struct {
	ImageURL string `json:"image"`
	LinkURL  string `json:"link"`
}

func getRandomFox() (Fox, error) {
	response, err := http.Get("https://randomfox.ca/floof/")
	if err != nil {
		return Fox{}, err
	}
	defer response.Body.Close()

	var fox Fox
	err = json.NewDecoder(response.Body).Decode(&fox)
	if err != nil {
		return Fox{
			ImageURL: "",
			LinkURL:  "",
		}, err
	}

	return fox, nil
}

func foxHandler(w http.ResponseWriter, r *http.Request) {
	fox, err := getRandomFox()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := "<meta charset='utf-8'>"
	html += "<title>Генератор картинок лис</title>"
	html += "<h1>Генерация случайных картинок лис</h1>"
	html += "<h2>Пример работы с API в Golang </h2>"
	html += "<h3><a href='" + fox.LinkURL + "'>Оригинальное изображение</a></h3>"
	html += "<img src='" + fox.ImageURL + "' alt='Лиса' height = '50%' width = '50%'"
	html += "<br/>"

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	fmt.Print("Сервер по генерации картинок случайных лис запущен!")
	http.HandleFunc("/", foxHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
