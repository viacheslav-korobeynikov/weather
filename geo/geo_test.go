package geo_test

import (
	"testing"

	"github.com/viacheslav-korobeynikov/weather/geo"
)

// Позитивный тест
func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка (ОР) данные для проверки
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата (ФР)
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

// Негативный тест
func TestGetMyLocationNoCity(t *testing.T) {
	// Arrange - подготовка (ОР) данные для проверки
	city := "Londonasdfad"
	// Act - выполняем функцию
	_, err := geo.GetMyLocation(city)
	// Assert - проверка результата (ФР)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrorNoCity, err)
	}
}
