package main

import (
	"encoding/json"
	//"fmt"
	"github.com/go-martini/martini"
	"github.com/iron-io/iron_go/mq"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	m := martini.Classic()
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	m.Any("/api/v0/documents/create.json", DocumentsCreate)

	m.Run()
}

func DocumentsCreate(req *http.Request, r render.Render) {
	_url := req.URL.Query().Get("url")
	webhook := req.URL.Query().Get("webhook")

	pngs := []string{}
	document := map[string]interface{}{"url": _url, "webhook": webhook, "status": "unprocessed", "pngs": pngs}
	payload := map[string]interface{}{"success": true, "document": document}

	addToQueue(document)

	r.JSON(200, payload)
}

func addToQueue(document interface{}) {
	marshaled_document, _ := json.Marshal(document)
	queue := mq.New(os.Getenv("QUEUE"))

	_, err := queue.PushString(string(marshaled_document))
	if err != nil {
		log.Fatal(err)
	}
}
