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
	ipapiClient := http.Client{}
	// Делаем запрос GET
	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipapiClient.Do(req)
	// Проверяем статус код
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT_200")
	}
	defer resp.Body.Close()
	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	// Декодируем тело ответа
	json.Unmarshal(body, &geo)
	return &geo, nil
}
