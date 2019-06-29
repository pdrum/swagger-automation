package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// FooBarRequest represents body of FooBar request.
type FooBarRequest struct {
	Foo string `json:"foo"`
	Bar []int  `json:"bar"`
}

// FooBarResponse represents body of FooBar response.
type FooBarResponse struct {
	Baz struct {
		Prop string `json:"prop"`
	} `json:"baz"`
}

// FooBarHandler handles incoming foobar requests
func FooBarHandler(ctx echo.Context) error {
	req := FooBarRequest{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := doSthWithRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

func doSthWithRequest(req FooBarRequest) FooBarResponse {
	return FooBarResponse{}
}
