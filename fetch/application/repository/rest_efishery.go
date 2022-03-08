package repository

import (
	"crypto/tls"
	"encoding/json"
	"fetch/application/model"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type (
	IRestEFisheryRepository interface {
		RestEFishery() (model []model.EFisheryDataResponse, err error)
	}

	RestEFisheryRepository struct {
		URL string
	}
)

func NewRestEFisheryRepository() *RestEFisheryRepository {
	return &RestEFisheryRepository{
		URL: os.Getenv("EFISHERY_URL"),
	}
}

func (repository *RestEFisheryRepository) RestEFishery() (model []model.EFisheryDataResponse, err error) {
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
