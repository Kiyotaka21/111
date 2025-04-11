package handlers

import (
	"fmt"
	"projectgrom/internal/cache"
	"projectgrom/internal/services/products"
)

type Handler struct {
	redis     *cache.RedisCache
	productDb *products.ProductsService
}

func NewHandler(data string) (*Handler, error) {
	redis, err := cache.InitRedis()
	if err != nil {
		return nil, err
	}
	product, err := products.InitProductsService(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Handler{redis: redis, productDb: product}, nil
}
