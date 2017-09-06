package gohvapi

import "encoding/json"
import "strconv"

type SSHKey struct {
	ID          int    `json:"id,string"`
	Name        string `json:"name"`
	Key         string `json:"ssh_key"`
	Fingerprint string `json:"fingerprint"`
}

func (c *Client) GetSSHKeys() (keys []SSHKey, err error) {

	var sshkeyList []SSHKey

	if err := c.get("/server/sshkeys", &sshkeyList); err != nil {
		return nil, err
	}

	return sshkeyList, nil
}

func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error) {
	if err := c.get("/server/sshkeys/"+strconv.Itoa(id), &sshkey); err != nil {
		return SSHKey{}, err
	}
	return sshkey, nil
}

func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error) {

	values := map[string]string{"ssh_key": key, "name": name}

	postData, _ := json.Marshal(values)

	if err := c.post("/server/sshkeys", postData, &sshkey); err != nil {
		return SSHKey{}, err
	}

	return sshkey, nil
}

func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error) {

	values := map[string]string{"ssh_key": key, "name": name}

	putData, _ := json.Marshal(values)

	var sshKey SSHKey
	if err := c.put("server/sshkeys"+strconv.Itoa(id), putData, &sshKey); err != nil {
		return SSHKey{}, err
	}
	return sshKey, nil
}

func (c *Client) DeleteSSHKey(id int) error {
	if err := c.delete("/server/sshkeys"+strconv.Itoa(id), nil, nil); err != nil {
		return err
	}
	return nil
}
