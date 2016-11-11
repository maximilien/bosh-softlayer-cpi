package pool

import (
	"fmt"
	strfmt "github.com/go-openapi/strfmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	bslcommon "github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/common"

	sl "github.com/maximilien/softlayer-go/softlayer"
	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/client"
	operations "github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/client/vm"

	. "github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/vm"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/models"
)

const SOFTLAYER_POOL_DELETER_LOG_TAG = "SoftLayerPoolDeleter"

type softLayerPoolDeleter struct {
	softLayerClient        sl.Client
	softLayerVmPoolClient  *client.SoftLayerVMPool
	logger       boshlog.Logger
}

func NewSoftLayerPoolDeleter(softLayerVmPoolClient *client.SoftLayerVMPool, softLayerClient sl.Client, logger boshlog.Logger) VMDeleter {
	return &softLayerPoolDeleter{
		softLayerClient:        softLayerClient,
		softLayerVmPoolClient: softLayerVmPoolClient,
		logger: logger,
	}
}

func (c *softLayerPoolDeleter) Delete(cid int) error {
	getVmByCidResp, err := c.softLayerVmPoolClient.VM.GetVMByCid(operations.NewGetVMByCidParams().WithCid(int32(cid)))
	if getVmByCidResp.Payload.VM != nil {
		free := models.VMState{
			State: models.StateFree,
		}
		_, err = c.softLayerVmPoolClient.VM.UpdateVMWithState(operations.NewUpdateVMWithStateParams().WithBody(&free).WithCid(int32(cid)))
		if err != nil {
			return bosherr.WrapErrorf(err, "Updating state of vm %d in pool to free", cid)
		}
		return nil
	}

	virtualGuest, err := bslcommon.GetObjectDetailsOnVirtualGuest(c.softLayerClient, cid)
	if err != nil {
		return bosherr.WrapError(err, fmt.Sprintf("Getting virtual guest %d details from SoftLayer", cid))
	}

	slPoolVm := &models.VM{
		Cid: int32(cid),
		CPU: int32(virtualGuest.StartCpus),
		MemoryMb: int32(virtualGuest.MaxMemory),
		IP:  strfmt.IPv4(virtualGuest.PrimaryBackendIpAddress),
		Hostname: virtualGuest.FullyQualifiedDomainName,
		PrivateVlan: int32(virtualGuest.PrimaryBackendNetworkComponent.NetworkVlan.Id),
		PublicVlan: int32(virtualGuest.PrimaryNetworkComponent.NetworkVlan.Id),
		State: models.StateFree,
	}
	_, err = c.softLayerVmPoolClient.VM.AddVM(operations.NewAddVMParams().WithBody(slPoolVm))
	if err != nil {
		return bosherr.WrapError(err, fmt.Sprintf("Adding vm %d to pool", cid))
	}

	return nil
}