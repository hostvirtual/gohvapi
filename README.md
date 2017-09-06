

# gohvapi
`import "github.com/hostvirtual/gohvapi"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package gohvapi provides a simple golang interface to the HostVirtual
Rest API at <a href="https://bapi.vr.org/">https://bapi.vr.org/</a>




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func GetKeyFromEnv() string](#GetKeyFromEnv)
* [type Client](#Client)
  * [func NewClient(apikey string) *Client](#NewClient)
  * [func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)](#Client.CreateSSHKey)
  * [func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error)](#Client.CreateServer)
  * [func (c *Client) DeleteSSHKey(id int) error](#Client.DeleteSSHKey)
  * [func (c *Client) DeleteServer(id int) error](#Client.DeleteServer)
  * [func (c *Client) GetLocations() ([]Location, error)](#Client.GetLocations)
  * [func (c *Client) GetOSs() ([]OS, error)](#Client.GetOSs)
  * [func (c *Client) GetPackage(id int) (pkg Package, err error)](#Client.GetPackage)
  * [func (c *Client) GetPackages() ([]Package, error)](#Client.GetPackages)
  * [func (c *Client) GetPlans() ([]Plan, error)](#Client.GetPlans)
  * [func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)](#Client.GetSSHKey)
  * [func (c *Client) GetSSHKeys() (keys []SSHKey, err error)](#Client.GetSSHKeys)
  * [func (c *Client) GetServer(id int) (server Server, err error)](#Client.GetServer)
  * [func (c *Client) GetServers() ([]Server, error)](#Client.GetServers)
  * [func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)](#Client.ProvisionServer)
  * [func (c *Client) RebootServer(id int) error](#Client.RebootServer)
  * [func (c *Client) StartServer(id int) error](#Client.StartServer)
  * [func (c *Client) StopServer(id int) error](#Client.StopServer)
  * [func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)](#Client.UpdateSSHKey)
* [type JobID](#JobID)
* [type Location](#Location)
* [type OS](#OS)
* [type Package](#Package)
* [type Plan](#Plan)
* [type SSHKey](#SSHKey)
* [type Server](#Server)
* [type ServerOptions](#ServerOptions)


#### <a name="pkg-files">Package files</a>
[client.go](/src/github.com/hostvirtual/gohvapi/client.go) [locations.go](/src/github.com/hostvirtual/gohvapi/locations.go) [os.go](/src/github.com/hostvirtual/gohvapi/os.go) [packages.go](/src/github.com/hostvirtual/gohvapi/packages.go) [plans.go](/src/github.com/hostvirtual/gohvapi/plans.go) [servers.go](/src/github.com/hostvirtual/gohvapi/servers.go) [sshkeys.go](/src/github.com/hostvirtual/gohvapi/sshkeys.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    Version      = "0.0.1"
    BaseEndpoint = "https://bapi.vr.org/"
    ContentType  = "application/json"
)
```



## <a name="GetKeyFromEnv">func</a> [GetKeyFromEnv](/src/target/client.go?s=727:754#L27)
``` go
func GetKeyFromEnv() string
```
GetKeyFromEnv is a simple function to try to yank the value for "VR_API_KEY" from
the environment




## <a name="Client">type</a> [Client](/src/target/client.go?s=521:623#L18)
``` go
type Client struct {
    // contains filtered or unexported fields
}
```
Client is the main object (struct) to which we attach most methods/functions.
It has the following fields: (client, userAgent, endPoint, apiKey)







### <a name="NewClient">func</a> [NewClient](/src/target/client.go?s=963:1000#L33)
``` go
func NewClient(apikey string) *Client
```
NewClient is the main entrypoint for instantiating a Client struct. It takes
your API Key as it's sole argument and returns the Client struct ready to talk to the API





### <a name="Client.CreateSSHKey">func</a> (\*Client) [CreateSSHKey](/src/target/sshkeys.go?s=621:695#L21)
``` go
func (c *Client) CreateSSHKey(name, key string) (sshkey SSHKey, err error)
```



### <a name="Client.CreateServer">func</a> (\*Client) [CreateServer](/src/target/servers.go?s=1698:1819#L77)
``` go
func (c *Client) CreateServer(name, plan string, locationID, osID int, options *ServerOptions) (server Server, err error)
```



### <a name="Client.DeleteSSHKey">func</a> (\*Client) [DeleteSSHKey](/src/target/sshkeys.go?s=1241:1284#L47)
``` go
func (c *Client) DeleteSSHKey(id int) error
```



### <a name="Client.DeleteServer">func</a> (\*Client) [DeleteServer](/src/target/servers.go?s=1535:1578#L68)
``` go
func (c *Client) DeleteServer(id int) error
```



### <a name="Client.GetLocations">func</a> (\*Client) [GetLocations](/src/target/locations.go?s=233:284#L1)
``` go
func (c *Client) GetLocations() ([]Location, error)
```
GetLocations public method on Client to get a list of locations




### <a name="Client.GetOSs">func</a> (\*Client) [GetOSs](/src/target/os.go?s=253:292#L3)
``` go
func (c *Client) GetOSs() ([]OS, error)
```



### <a name="Client.GetPackage">func</a> (\*Client) [GetPackage](/src/target/packages.go?s=448:508#L14)
``` go
func (c *Client) GetPackage(id int) (pkg Package, err error)
```



### <a name="Client.GetPackages">func</a> (\*Client) [GetPackages](/src/target/packages.go?s=254:303#L3)
``` go
func (c *Client) GetPackages() ([]Package, error)
```



### <a name="Client.GetPlans">func</a> (\*Client) [GetPlans](/src/target/plans.go?s=282:325#L3)
``` go
func (c *Client) GetPlans() ([]Plan, error)
```



### <a name="Client.GetSSHKey">func</a> (\*Client) [GetSSHKey](/src/target/sshkeys.go?s=431:492#L14)
``` go
func (c *Client) GetSSHKey(id int) (sshkey SSHKey, err error)
```



### <a name="Client.GetSSHKeys">func</a> (\*Client) [GetSSHKeys](/src/target/sshkeys.go?s=233:289#L3)
``` go
func (c *Client) GetSSHKeys() (keys []SSHKey, err error)
```



### <a name="Client.GetServer">func</a> (\*Client) [GetServer](/src/target/servers.go?s=860:921#L34)
``` go
func (c *Client) GetServer(id int) (server Server, err error)
```



### <a name="Client.GetServers">func</a> (\*Client) [GetServers](/src/target/servers.go?s=673:720#L23)
``` go
func (c *Client) GetServers() ([]Server, error)
```



### <a name="Client.ProvisionServer">func</a> (\*Client) [ProvisionServer](/src/target/servers.go?s=2426:2536#L102)
``` go
func (c *Client) ProvisionServer(name string, id, locationID, osID int, options *ServerOptions) (JobID, error)
```



### <a name="Client.RebootServer">func</a> (\*Client) [RebootServer](/src/target/servers.go?s=1372:1415#L59)
``` go
func (c *Client) RebootServer(id int) error
```



### <a name="Client.StartServer">func</a> (\*Client) [StartServer](/src/target/servers.go?s=1048:1090#L41)
``` go
func (c *Client) StartServer(id int) error
```



### <a name="Client.StopServer">func</a> (\*Client) [StopServer](/src/target/servers.go?s=1209:1250#L50)
``` go
func (c *Client) StopServer(id int) error
```



### <a name="Client.UpdateSSHKey">func</a> (\*Client) [UpdateSSHKey](/src/target/sshkeys.go?s=917:988#L34)
``` go
func (c *Client) UpdateSSHKey(id int, name, key string) (SSHKey, error)
```



## <a name="JobID">type</a> [JobID](/src/target/servers.go?s=623:671#L19)
``` go
type JobID struct {
    ID int `json:"id,string"`
}
```









## <a name="Location">type</a> [Location](/src/target/locations.go?s=82:165#L1)
``` go
type Location struct {
    ID   int    `json:"id,string"`
    Name string `json:"name"`
}
```
Location is a struct for storing the id and name of a location










## <a name="OS">type</a> [OS](/src/target/os.go?s=17:251#L1)
``` go
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









## <a name="Package">type</a> [Package](/src/target/packages.go?s=35:252#L1)
``` go
type Package struct {
    ID        int    `json:"mbpkgid,string"`
    Status    string `json:"package_status"`
    Locked    string `json:"locked"`
    PlanName  string `json:"name"`
    Installed int    `json:"installed,string"`
}
```









## <a name="Plan">type</a> [Plan](/src/target/plans.go?s=17:280#L1)
``` go
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









## <a name="SSHKey">type</a> [SSHKey](/src/target/sshkeys.go?s=58:231#L1)
``` go
type SSHKey struct {
    ID          int    `json:"id,string"`
    Name        string `json:"name"`
    Key         string `json:"ssh_key"`
    Fingerprint string `json:"fingerprint"`
}
```









## <a name="Server">type</a> [Server](/src/target/servers.go?s=76:533#L1)
``` go
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









## <a name="ServerOptions">type</a> [ServerOptions](/src/target/servers.go?s=535:621#L13)
``` go
type ServerOptions struct {
    SSHKeyID    int
    Password    string
    CloudConfig string
}
```













- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
