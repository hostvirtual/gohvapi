package govrapi

import "strconv"

type Package struct {
	ID        int    `json:"mbpkgid,string"`
	Status    string `json:"package_status"`
	Locked    string `json:"locked"`
	PlanName  string `json:"name"`
	Installed int    `json:"installed,string"`
}

func (c *Client) GetPackages() ([]Package, error) {

	var packageList []Package

	if err := c.get("cloud/packages", &packageList); err != nil {
		return nil, err
	}

	return packageList, nil
}

func (c *Client) GetPackage(id int) (pkg Package, err error) {
	if err := c.get("/cloud/package/"+strconv.Itoa(id), &pkg); err != nil {
		return Package{Installed: 0}, err
	}
	return pkg, nil
}
