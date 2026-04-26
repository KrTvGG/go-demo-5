package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - подготовка, expected  результат
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результат с expected
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидаломь %v, полученом %v", expected, got)
	}
}