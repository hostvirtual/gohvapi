# gohvapi
--
    import "github.com/hostvirtual/gohvapi"


## Usage

```go
const (
	Version      = "0.0.1"
	BaseEndpoint = "https://bapi.vr.org/"
	ContentType  = "application/json"
)
```

#### func  GetKeyFromEnv

```go
func GetKeyFromEnv() string
```

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(apikey string) *Client
```

#### func (*Client) CreateSSHKey

```go
func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)
```

#### func (*Client) CreateServer

```go
func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error)
```

#### func (*Client) DeleteSSHKey

```go
func (c *Client) DeleteSSHKey(id int) error
```

#### func (*Client) DeleteServer

```go
func (c *Client) DeleteServer(id int) error
```

#### func (*Client) GetLocations

```go
func (c *Client) GetLocations() ([]Location, error)
```

#### func (*Client) GetOSs

```go
func (c *Client) GetOSs() ([]OS, error)
```

#### func (*Client) GetPackage

```go
func (c *Client) GetPackage(id int) (pkg Package, err error)
```

#### func (*Client) GetPackages

```go
func (c *Client) GetPackages() ([]Package, error)
```

#### func (*Client) GetPlans

```go
func (c *Client) GetPlans() ([]Plan, error)
```

#### func (*Client) GetSSHKey

```go
func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)
```

#### func (*Client) GetSSHKeys

```go
func (c *Client) GetSSHKeys() (keys []SSHKey, err error)
```

#### func (*Client) GetServer

```go
func (c *Client) GetServer(id int) (server Server, err error)
```

#### func (*Client) GetServers

```go
func (c *Client) GetServers() ([]Server, error)
```

#### func (*Client) ProvisionServer

```go
func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)
```

#### func (*Client) RebootServer

```go
func (c *Client) RebootServer(id int) error
```

#### func (*Client) StartServer

```go
func (c *Client) StartServer(id int) error
```

#### func (*Client) StopServer

```go
func (c *Client) StopServer(id int) error
```

#### func (*Client) UpdateSSHKey

```go
func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)
```

#### type JobID

```go
type JobID struct {
	ID int `json:"id,string"`
}
```


#### type Location

```go
type Location struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
}
```


#### type OS

```go
type OS struct {
	ID      int    `json:"id,string"`
	Os      string `json:"os"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Size    string `json:"size"`
	Bits    string `json:"bits"`
	Tech    string `json:"tech"`
}
```


#### type Package

```go
type Package struct {
	ID        int    `json:"mbpkgid,string"`
	Status    string `json:"package_status"`
	Locked    string `json:"locked"`
	PlanName  string `json:"name"`
	Installed int    `json:"installed,string"`
}
```


#### type Plan

```go
type Plan struct {
	ID        int    `json:"plan_id,string"`
	Name      string `json:"plan"`
	RAM       string `json:"ram"`
	Disk      string `json:"disk"`
	Transfer  string `json:"transfer"`
	Price     string `json:"price"`
	Available string `json:"available"`
}
```


#### type SSHKey

```go
type SSHKey struct {
	ID          int    `json:"id,string"`
	Name        string `json:"name"`
	Key         string `json:"ssh_key"`
	Fingerprint string `json:"fingerprint"`
}
```


#### type Server

```go
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
```


#### type ServerOptions

```go
type ServerOptions struct {
	SSHKeyID    int
	Password    string
	CloudConfig string
}
```
