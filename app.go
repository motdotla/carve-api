package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/handshakejs/handshakejserrors"
	"github.com/iron-io/iron_go/mq"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
	"os"
)

const (
	LOGIC_ERROR_CODE_UNKNOWN = "unknown"
)

var (
	QUEUE string
)

func CrossDomain() martini.Handler {
	return func(res http.ResponseWriter) {
		res.Header().Add("Access-Control-Allow-Origin", "*")
	}
}

type Document struct {
	Url     string `form:"url" json:"url"`
	Webhook string `form:"webhook" json:"webhook"`
}

func main() {
	loadEnvs()

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Any("/api/v0/documents/create.json", binding.Bind(Document{}), DocumentsCreate)

	m.Run()
}

func ErrorPayload(logic_error *handshakejserrors.LogicError) map[string]interface{} {
	error_object := map[string]interface{}{"code": logic_error.Code, "field": logic_error.Field, "message": logic_error.Message}
	errors := []interface{}{}
	errors = append(errors, error_object)
	payload := map[string]interface{}{"errors": errors}

	return payload
}

func DocumentsPayload(document map[string]interface{}) map[string]interface{} {
	documents := []interface{}{}
	documents = append(documents, document)
	payload := map[string]interface{}{"documents": documents}

	return payload
}

func DocumentsCreate(document Document, req *http.Request, r render.Render) {
	_url := document.Url
	webhook := document.Webhook

	pages := []string{}
	params := map[string]interface{}{"url": _url, "webhook": webhook, "status": "processing", "pages": pages}

	_, logic_error := addToQueue(params)
	if logic_error != nil {
		payload := ErrorPayload(logic_error)
		statuscode := determineStatusCodeFromLogicError(logic_error)
		r.JSON(statuscode, payload)
	} else {
		payload := DocumentsPayload(params)
		r.JSON(200, payload)
	}
}

func addToQueue(document interface{}) (string, *handshakejserrors.LogicError) {
	marshaled_document, _ := json.Marshal(document)
	queue := mq.New(QUEUE)

	_, err := queue.PushString(string(marshaled_document))
	if err != nil {
		logic_error := &handshakejserrors.LogicError{"unknown", "", err.Error()}
		return "", logic_error
	}

	return "", nil
}

func determineStatusCodeFromLogicError(logic_error *handshakejserrors.LogicError) int {
	code := 400
	if logic_error.Code == LOGIC_ERROR_CODE_UNKNOWN {
		code = 500
	}

	return code
}

func loadEnvs() {
	godotenv.Load()

	QUEUE = os.Getenv("QUEUE")
}
