package cat

type Cat struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

func (c *Cat) GetCountry() string {
	return c.Country
}

func (c *Cat) GetBreed() string {
	return c.Breed
}
