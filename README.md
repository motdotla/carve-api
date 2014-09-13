# carve-api 

<img src="https://raw.githubusercontent.com/motdotla/carve-api/master/carve-api.gif" alt="carve-api" align="right" />

API for converting PDFs into an array of PNGs.

The carve-api is based on [CONTRA]() API design. [JSON](http://www.json.org) is returned in all responses from the API, including errors. 

I've tried to make it as easy to use as possible, but if you have any feedback please [let me know](mailto:mot@mot.la).

## Installation

### Production

```
git clone https://github.com/motdotla/carve-api.git
cd carve-api
heroku create -b https://github.com/kr/heroku-buildpack-go.git
heroku config:set IRON_TOKEN=YOUR_TOKEN 
heroku config:set IRON_PROJECT_ID=YOUR_PROJECT_ID 
heroku config:set QUEUE=carve 
git push heroku master
heroku open
```

### Development

```
git clone https://github.com/motdotla/carve-api.git
cd carve-api
go get 
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
http://carve.io/api/v0/documents/create.json?url=http://motdotla.com/assets/resume.pdf&webhook=http://requestb.in/nxwuxynx"
```

##### Example Response

```
{
  "document": {
    "pages": [ ]
    "status": "unprocessed",
    "url": "http://motdotla.com/assets/resume.pdf",
    "webhook": "http://requestb.in/nxwuxynx",
  },
  "success": true
}
```

##### Example Error

```
{
  "error": {
    "message": "This is the error message"
  }
  "success": false
}
```

## TODO

* validation on webhook and url
* url cleansing on webhook and url
* 500 response returns on errors
