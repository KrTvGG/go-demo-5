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

	resutl := weather.GetWeather(geoData, format)

	if !strings.Contains(resutl, expected) {
		t.Errorf("Ожидаломь %v, полученом %v", expected, resutl)
	}
}