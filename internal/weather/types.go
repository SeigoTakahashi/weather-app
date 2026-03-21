package weather

type WeatherResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Id		  int    `json:"id"`
		Icon		string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}