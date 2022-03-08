package repository

import (
	"crypto/tls"
	"encoding/json"
	"fetch/application/model"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type (
	ICacheCurrencyRepository interface {
		ConvertIDRToUSD() (model model.RestCurrencyResponse, err error)
	}

	CacheCurrencyRepository struct {
		URL      string
		Cache    *cache.Cache
		CacheKey string
	}
)

func NewCacheCurrencyRepository() ICacheCurrencyRepository {
	return &CacheCurrencyRepository{
		URL:      "https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=e418d7133e35a3a195e6",
		Cache:    cache.New(5*time.Minute, 10*time.Minute),
		CacheKey: "IDR_USD",
	}
}

func (repository *CacheCurrencyRepository) ConvertIDRToUSD() (model model.RestCurrencyResponse, err error) {
	// Get on cache
	dataCache, found := repository.Cache.Get(repository.CacheKey)
	if found {
		model.IDRToUSD = dataCache.(float64)
		return model, nil
	}

	// Rest Currency and Set Cache
	model, err = repository.rest()
	if err != nil {
		return model, err
	}
	log.Println("REST DATA CURRENCY")
	repository.Cache.Set(repository.CacheKey, model.IDRToUSD, cache.DefaultExpiration)
	return model, nil
}

func (repository *CacheCurrencyRepository) rest() (model model.RestCurrencyResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, repository.URL, nil)
	if err != nil {
		return model, err
	}

	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return model, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &model)
	if err != nil {
		return model, err
	}

	return model, err
}
