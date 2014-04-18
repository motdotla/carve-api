# carve-api 

API for converting PDFs into an array of PNGs

The carve-api is based on [CONTRA]() API design. [JSON](http://www.json.org) is returned in all responses from the API, including errors. 

I've tried to make it as easy to use as possible, but if you have any feedback please [let me know](mailto:scott@scottmotte.com).

## Requirements

* Linux or Mac environment
* Account at [Iron.io](http://iron.io)

## Installation

```
git clone 
go get github.com/go-martini/martini
go get github.com/martini-contrib/render
go get github.com/joho/godotenv
go get github.com/iron-io/iron_go/mq
cp .env.example .env
go run app.go
```

Make sure you edit the contents of `.env`.

## Summary

### API Endpoint

* [http://carve.io/api/v0](http://carve.io/api/v0)

### Documents

#### /documents/create

Create a document to generate PNGs.

##### Definition

```
POST|GET http://carve.io/api/v0/documents/create.json?url=[url]&webhook=[webhook]
```

##### Parameters

* url*
* webhook*

##### Example Request

```
http://carve.io/api/v0/documents/create.json?url=http://scottmotte.com/assets/resume.pdf&webhook=http://requestb.in/nxwuxynx"
```

##### Example Response

```
{
  "document": {
    "pngs": [ ]
    "status": "unprocessed",
    "url": "http://scottmotte.com/assets/resume.pdf",
    "webhook": "http://requestb.in/nxwuxynx",
  },
  "success": true
}
```
