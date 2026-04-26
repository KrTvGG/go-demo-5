package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3

	resutl, err := weather.GetWeather(geoData, format)

	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(resutl, expected) {
		t.Errorf("Ожидаломь %v, полученом %v", expected, resutl)
	}
}

func TestGetWeatherWrongFormat(t *testing.T) {
	expected := "moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 125

	_, err := weather.GetWeather(geoData, format)

	if err != weather.ErrorWrongFormat {
		t.Errorf("Ожидаломь %v, полученом %v", weather.ErrorWrongFormat, err)
	}
}