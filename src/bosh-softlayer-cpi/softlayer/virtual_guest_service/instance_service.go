package instance

import (
	"github.com/softlayer/softlayer-go/datatypes"
)

//go:generate counterfeiter -o fakes/fake_Instance_Service.go . Service
type Service interface {
	AttachDisk(id int, diskID int) (string, string, error)
	AttachedDisks(id int) ([]string, error)
	AttachEphemeralDisk(id int, diskSize int) error
	Create(virtualGuest datatypes.Virtual_Guest, networks Networks, registryEndpoint string) (int, error)
	ConfigureNetworks(id int, networks Networks) (Networks, error)
	CleanUp(id int)
	Delete(id int) error
	DetachDisk(id int, diskID int) error
	Edit(id int, instance datatypes.Virtual_Guest) (bool, error)
	Find(id int) (datatypes.Virtual_Guest, bool, error)
	FindByPrimaryBackendIp(ip string) (datatypes.Virtual_Guest, bool, error)
	FindByPrimaryIp(ip string) (datatypes.Virtual_Guest, bool, error)
	GetVlan(id int, mask string) (datatypes.Network_Vlan, error)
	ReAttachLeftDisk(id int, devicePath string, diskID int) error
	Reboot(id int) error
	ReloadOS(id int, stemcellID int, sshKeyIds []int, vmNamePrefix string, domain string) (string, error)
	SetMetadata(id int, vmMetadata Metadata) error
	CreateSshKey(label string, key string, fingerPrint string) (int, error)
	DeleteSshKey(id int) error
}

type Metadata map[string]interface{}

type DavConfig map[string]interface{}
