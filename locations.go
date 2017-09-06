//gohvapi package file for location handling
package gohvapi

type Location struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
}

func (c *Client) GetLocations() ([]Location, error) {

	var locationMap map[string]Location
	var locationList []Location

	if err := c.get("cloud/locations", &locationMap); err != nil {
		return nil, err
	}

	for _, loc := range locationMap {
		locationList = append(locationList, loc)
	}

	return locationList, nil
}
