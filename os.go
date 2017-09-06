package govrapi

type OS struct {
	ID      int    `json:"id,string"`
	Os      string `json:"os"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Size    string `json:"size"`
	Bits    string `json:"bits"`
	Tech    string `json:"tech"`
}

func (c *Client) GetOSs() ([]OS, error) {

	var osList []OS

	if err := c.get("cloud/images", &osList); err != nil {
		return nil, err
	}

	return osList, nil
}
