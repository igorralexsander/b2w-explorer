package api_routes

import (
	"b2w-explorer/internal/app/model"
	"b2w-explorer/internal/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type productPageRoute struct {
	productPageService service.ProductPage
}

type ProductPageRoute interface {
}

func NewProductPageRoute(productPageService service.ProductPage) *productPageRoute {
	return &productPageRoute{productPageService}
}

func (r productPageRoute) RegisterEndpoints(e *echo.Echo) {
	group := e.Group("/product-page")
	group.POST("", r.FetchProductPage)
}

func (r productPageRoute) FetchProductPage(c echo.Context) error {
	var productPageRequest model.ProductPageRequest
	if err := c.Bind(&productPageRequest); err != nil {
		return err
	}
	result, err := r.productPageService.FetchProductPage(productPageRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	c.Response().Header().Set("Content-Type", "text/html")
	c.String(http.StatusOK, result)
	return nil
}
