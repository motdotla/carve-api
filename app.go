package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/iron-io/iron_go/mq"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/render"
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

	pages := []string{}
	document := map[string]interface{}{"url": _url, "webhook": webhook, "status": "processing", "pages": pages}
	payload := map[string]interface{}{"success": true, "document": document}

	_, err := addToQueue(document)
	if err != nil {
		error_object := map[string]interface{}{"message": err.Error()}
		err_payload := map[string]interface{}{"success": false, "error": error_object}
		r.JSON(500, err_payload)
	} else {
		r.JSON(200, payload)
	}
}

func addToQueue(document interface{}) (string, error) {
	marshaled_document, _ := json.Marshal(document)
	queue := mq.New(os.Getenv("QUEUE"))

	_, err := queue.PushString(string(marshaled_document))
	if err != nil {
		return "", err
	}

	return "", nil
}
