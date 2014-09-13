# carve-api 

<img src="https://raw.githubusercontent.com/motdotla/carve-api/master/carve-api.gif" alt="carve-api" align="right" width="320" />

API for converting PDFs into an array of PNGs. Works in tandem with [carve-worker](https://github.com/motdotla/carve-worker).

You can view the [documentation](http://docs.carveapi.apiary.io/).

## Installation

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Development

```
git clone https://github.com/motdotla/carve-api.git
cd carve-api
go get 
cp .env.example .env
go run app.go
```

Edit the contents of `.env`.

## TODO

* validation on webhook and url
* url cleansing on webhook and url

