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

var testCases = []struct {
	name   string
	format int
}{
	{ name: "BigFormat", format: 147 },
	{ name: "0 format", format: 0 },
	{ name: "Minus format", format: -1 },
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "moscow"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geoData, tc.format)

			if err != weather.ErrorWrongFormat {
				t.Errorf("Ожидаломь %v, полученом %v", weather.ErrorWrongFormat, err)
			}
		})
	}
}
