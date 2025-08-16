package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	// Делаем запрос GET
	response, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	// Проверяем статус код
	if response.StatusCode != 200 {
		return nil, errors.New("NOT_200")
	}
	// Читаем тело ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	// Декодируем тело ответа
	json.Unmarshal(body, &geo)
	return &geo, nil
}
