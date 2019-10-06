package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"vo"

	"github.com/labstack/echo"
)

//http://localhost:8000/cats/json?name=arnold&type=fluffy
func GetBann(c echo.Context) error {
	zoneId := c.QueryParam("languageCode")
	id := c.QueryParam("pageId")
	//deviceTypeId := c.Param("deviceTypedId")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("%s\n", zoneId, id))
	} else if dataType == "json" {
		bann := vo.Bann{
			zoneId": languageCode,
    			width: "1920",
    			height": 400,
    			banners: ,
		}
		return c.JSON(http.StatusOK, bann)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON",
		})
	}

}

func AddBann(c echo.Context) error {
	cat := vo.Cat{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")
}
