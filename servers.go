package gohvapi

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type Server struct {
	Name         string `json:"fqdn"`
	ID           int    `json:"mbpkgid,string"`
	OS           string `json:"os"`
	PrimaryIPv4  string `json:"ip"`
	PrimaryIPv6  string `json:"ipv6"`
	PlanID       int    `json:"plan_id,string"`
	PkgID        int    `json:"pkg_id,string"`
	LocationID   int    `json:"location_id,string"`
	OSID         int    `json:"os_id,string"`
	ServerStatus string `json:"status"`
	PowerStatus  string `json:"state"`
}

type ServerOptions struct {
	SSHKeyID    int
	Password    string
	CloudConfig string
}

type JobID struct {
	ID int `json:"id,string"`
}

func (c *Client) GetServers() ([]Server, error) {

	var serverList []Server

	if err := c.get("cloud/servers", &serverList); err != nil {
		return nil, err
	}

	return serverList, nil
}

func (c *Client) GetServer(id int) (server Server, err error) {
	if err := c.get("/cloud/server/"+strconv.Itoa(id), &server); err != nil {
		return Server{}, err
	}
	return server, nil
}

func (c *Client) StartServer(id int) error {

	if err := c.post("/cloud/server/start/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) StopServer(id int) error {

	if err := c.post("/cloud/server/shutdown/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) RebootServer(id int) error {

	if err := c.post("/cloud/server/reboot/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteServer(id int) error {

	if err := c.post("/cloud/server/delete/"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error) {

	values := map[string]string{"plan": plan, "fqdn": name, "location": strconv.Itoa(locationID), "image": strconv.Itoa(osID)}

	if options != nil {
		if options.SSHKeyID != 0 {
			values["ssh_key_id"] = strconv.Itoa(options.SSHKeyID)
		}
		if options.Password != "" {
			values["password"] = options.Password
		}
		if options.CloudConfig != "" {
			values["cloud_config"] = base64.StdEncoding.EncodeToString([]byte(options.CloudConfig))
		}
	}

	postData, _ := json.Marshal(values)

	if err := c.post("/cloud/buy_build/", postData, &server); err != nil {
		return Server{}, err
	}

	return server, nil
}

func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error) {

	var jobid JobID

	values := map[string]string{"fqdn": name, "location": strconv.Itoa(locationID), "image": strconv.Itoa(osID)}

	if options != nil {
		if options.SSHKeyID != 0 {
			values["ssh_key_id"] = strconv.Itoa(options.SSHKeyID)
		}
		if options.Password != "" {
			values["password"] = options.Password
		}
		if options.CloudConfig != "" {
			values["cloud_config"] = base64.StdEncoding.EncodeToString([]byte(options.CloudConfig))
		}

	}

	postData, _ := json.Marshal(values)

	if err := c.post("/cloud/server/"+strconv.Itoa(id), postData, &jobid); err != nil {
		return JobID{}, err
	}

	return jobid, nil
}
