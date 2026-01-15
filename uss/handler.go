package uss

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/labstack/echo/v4"
)

func HandleCreate(c echo.Context) error {
	urlInput := c.FormValue("url")
	parsed, e := url.Parse(urlInput)
	if e != nil {
		return c.String(http.StatusBadRequest, "invalid url!")
	}
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}
	shortId := Save(parsed.String())
	return c.String(http.StatusOK, shortId)
}

func HandlePreview(c echo.Context) error {
	shortId := c.Param("shortId")
	url, ok := Get(shortId)
	if ok {
		return c.String(http.StatusOK, url)
	} else {
		return c.String(http.StatusNotFound, fmt.Sprintf("%s not found!", shortId))
	}
}

func HandleResolve(c echo.Context) error {
	shortId := c.Param("shortId")
	url, ok := Get(shortId)
	if ok {
		return c.Redirect(http.StatusFound, url)
	} else {
		return c.String(http.StatusNotFound, fmt.Sprintf("%s not found!", shortId))
	}
}

func HandleDelete(c echo.Context) error {
	shortId := c.Param("shortId")
	Delete(shortId)
	return c.String(http.StatusOK, "")
}