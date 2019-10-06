package main

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	type Banners struct {
	languageCode  string `json:"zoneId" xml:"zoneId" form:"languageCode" query:"languageCode"`
	pageId string `json:"id" xml:"id" form:"id" query:"id"`
	//others elements
}



	e.Logger.Fatal(e.Start(":1323"))
}
