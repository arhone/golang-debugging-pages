package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type (
	// Config - Настройки
	Config struct {
		Port  int    `json:"port"`
		Token string `json:"token"`
		Debug bool   `json:"debug"`
		Delay int64  `json:"Delay"`
	}

	// PageVariables - переменные страниц
	PageVariables struct {
		Title       string
		Description string
		Keywords    string
		Content     string
		Delay       int64
	}

	// ResponseJsonBody - Структура тела ответа
	ResponseJsonBody struct {
		Status struct {
			Code    int    `json:"code"`
			Message string `json:"message,omitempty"`
		} `json:"status"`
		Meta struct {
			Total int `json:"total,omitempty"`
		} `json:"meta,omitempty"`
		Object  interface{} `json:"object,omitempty"`
		Objects interface{} `json:"objects,omitempty"`
	}

	// ResponseHtmlBody - Структура тела ответа
	ResponseHtmlBody struct {
		Status struct {
			Code    int
			Message string
		}
		Html string
	}
)

// Хранилище настроек
var configStorage = Config{}

// main - Запуск сервиса
func main() {

	// Подключение логирования
	if err := os.MkdirAll("logs", os.FileMode(0755)); err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile("logs/main.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Запуск сервиса")

	// Загрузка настроек
	configFile, err := ioutil.ReadFile("config/main/config.json")
	if err != nil {
		log.Error(err.Error())
	}
	if err = json.Unmarshal(configFile, &configStorage); err != nil {
		log.Error(err.Error())
	}

	if configStorage.Debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("debug mod: " + strconv.FormatBool(configStorage.Debug))

	// Отдача статики
	http.Handle("/favicon.ico", http.FileServer(http.Dir("public")))
	http.Handle("/robots.txt", http.FileServer(http.Dir("public")))
	http.Handle("/assets/", http.FileServer(http.Dir("public")))

	// Отдача динамики
	http.HandleFunc("/", router)
	err = http.ListenAndServe(":"+strconv.Itoa(configStorage.Port), nil)
	if err != nil {
		log.Error(err)
		return
	}

}

// router - Маршрутизация
func router(response http.ResponseWriter, request *http.Request) {

	request.URL.Path = filepath.Clean(request.URL.Path)

	if request.Method == "GET" {

		if request.URL.Path == "/" {

			body := new(ResponseHtmlBody)
			body.Status.Code = http.StatusOK
			body.Html = getRenderString("templates/index.html", PageVariables{
				Title:       "Главная",
				Description: "Текущее время",
				Keywords:    "время",
				Content:     getCurrentTime(),
				Delay:       configStorage.Delay,
			})
			sendHtmlResponse(response, body)
			return

		} else if request.URL.Path == "/api/0/time/current.json" {

			time.Sleep(time.Duration(configStorage.Delay) * time.Second)

			body := new(ResponseJsonBody)
			body.Status.Code = http.StatusOK
			body.Status.Message = getCurrentTime()
			sendJsonResponse(response, body)
			return

		}

	}

	body := new(ResponseHtmlBody)
	body.Status.Code = http.StatusNotFound
	body.Html = getRenderString("templates/index.html", PageVariables{
		Title:       "Страница не найдена",
		Description: "Такой страница не существует",
		Keywords:    "404",
	})
	sendHtmlResponse(response, body)
	return

}

// getCurrentTime - Возвращает текущее время
func getCurrentTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d", t.Day()) + "." + fmt.Sprintf("%02d", t.Month()) + "." + fmt.Sprintf("%02d", t.Year()) +
		" " + fmt.Sprintf("%02d", t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute()) + ":" + fmt.Sprintf("%02d", t.Second())
}

// getRenderString - Возвращает рендер шаблона
func getRenderString(path string, pageVariables PageVariables) string {

	html, err := template.ParseFiles(path)
	if err != nil {
		log.Error()
	}

	var tpl bytes.Buffer
	if err := html.Execute(&tpl, pageVariables); err != nil {
		log.Error()
	}

	return tpl.String()

}

// sendHtmlResponse - Отвечает клиенту
func sendHtmlResponse(response http.ResponseWriter, body *ResponseHtmlBody) bool {

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.WriteHeader(body.Status.Code)

	if _, err := response.Write([]byte(body.Html)); err != nil {
		return false
	}

	return true

}

// sendTextResponse - Отвечает клиенту
func sendTextResponse(response http.ResponseWriter, text string) bool {

	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)

	if _, err := response.Write([]byte(text)); err != nil {
		return false
	}

	return true

}

// sendResponse - Отвечает клиенту
func sendJsonResponse(response http.ResponseWriter, body *ResponseJsonBody) bool {

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(http.StatusOK)

	data, err := json.Marshal(body)
	if err != nil {
		if _, err := response.Write(data); err != nil {
			return false
		}
	}

	if _, err := response.Write(data); err != nil {
		return false
	}

	return true

}
