package vm_test

import (
	"encoding/json"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maximilien/bosh-softlayer-cpi/softlayer/vm"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	testhelpers "github.com/maximilien/bosh-softlayer-cpi/test_helpers"

	bslcommon "github.com/maximilien/bosh-softlayer-cpi/softlayer/common"
	bsldisk "github.com/maximilien/bosh-softlayer-cpi/softlayer/disk"
	bslvm "github.com/maximilien/bosh-softlayer-cpi/softlayer/vm"

	fakedisk "github.com/maximilien/bosh-softlayer-cpi/softlayer/disk/fakes"
	fakevm "github.com/maximilien/bosh-softlayer-cpi/softlayer/vm/fakes"
	fakesutil "github.com/maximilien/bosh-softlayer-cpi/util/fakes"
	fakeslclient "github.com/maximilien/softlayer-go/client/fakes"
)

var _ = Describe("SoftLayerVM", func() {
	var (
		softLayerClient *fakeslclient.FakeSoftLayerClient
		sshClient       *fakesutil.FakeSshClient
		agentEnvService *fakevm.FakeAgentEnvService
		logger          boshlog.Logger
		vm              SoftLayerVM
	)

	BeforeEach(func() {
		softLayerClient = fakeslclient.NewFakeSoftLayerClient("fake-username", "fake-api-key")
		sshClient = fakesutil.NewFakeSshClient()
		agentEnvService = &fakevm.FakeAgentEnvService{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		vm = NewSoftLayerVM(1234, softLayerClient, sshClient, agentEnvService, logger)
	})

	Describe("Delete", func() {
		Context("valid VM ID is used and averageDuration is normal", func() {
			BeforeEach(func() {
				fileNames := []string{
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
					"SoftLayer_Virtual_Guest_Service_deleteObject_true.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
					"SoftLayer_Virtual_Guest_Service_getObject.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransaction.json",
					"SoftLayer_Virtual_Guest_Service_getEmptyObject.json",
				}
				testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
			})

			It("deletes the VM successfully", func() {
				vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
				bslcommon.TIMEOUT = 2 * time.Second
				bslcommon.POLLING_INTERVAL = 1 * time.Second

				err := vm.Delete()
				Expect(err).ToNot(HaveOccurred())
			})

			It("postCheckActiveTransactionsForDeleteVM time out", func() {
				vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
				bslcommon.TIMEOUT = 1 * time.Second
				bslcommon.POLLING_INTERVAL = 1 * time.Second

				err := vm.Delete()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("valid VM ID is used and averageDuration is \"\"", func() {
			BeforeEach(func() {
				fileNames := []string{
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
					"SoftLayer_Virtual_Guest_Service_deleteObject_true.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
					"SoftLayer_Virtual_Guest_Service_getObject.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransaction_ADEmpty.json",
					"SoftLayer_Virtual_Guest_Service_getEmptyObject.json",
				}
				testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
			})

			It("deletes the VM successfully", func() {
				vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
				bslcommon.TIMEOUT = 2 * time.Second
				bslcommon.POLLING_INTERVAL = 1 * time.Second

				err := vm.Delete()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("valid VM ID is used and averageDuration is invalid", func() {
			BeforeEach(func() {
				fileNames := []string{
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
					"SoftLayer_Virtual_Guest_Service_deleteObject_true.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
					"SoftLayer_Virtual_Guest_Service_getObject.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransaction_ADInvalid.json",
					"SoftLayer_Virtual_Guest_Service_getEmptyObject.json",
				}
				testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
			})

			It("deletes the VM successfully", func() {
				vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
				bslcommon.TIMEOUT = 2 * time.Second
				bslcommon.POLLING_INTERVAL = 1 * time.Second

				err := vm.Delete()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("invalid VM ID is used", func() {
			BeforeEach(func() {
				fileNames := []string{
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
					"SoftLayer_Virtual_Guest_Service_deleteObject_false.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
					"SoftLayer_Virtual_Guest_Service_getObject.json",
					"SoftLayer_Virtual_Guest_Service_getActiveTransaction.json",
					"SoftLayer_Virtual_Guest_Service_getEmptyObject.json",
				}
				testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
				vm = NewSoftLayerVM(00000, softLayerClient, sshClient, agentEnvService, logger)
				bslcommon.TIMEOUT = 2 * time.Second
				bslcommon.POLLING_INTERVAL = 1 * time.Second
			})

			It("fails deleting the VM", func() {
				err := vm.Delete()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Reboot", func() {
		Context("valid VM ID is used", func() {
			BeforeEach(func() {
				softLayerClient.DoRawHttpRequestResponse = []byte("true")
				vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			})

			It("reboots the VM successfully", func() {
				err := vm.Reboot()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("invalid VM ID is used", func() {
			BeforeEach(func() {
				softLayerClient.DoRawHttpRequestResponse = []byte("false")
				vm = NewSoftLayerVM(00000, softLayerClient, sshClient, agentEnvService, logger)
			})

			It("fails rebooting the VM", func() {
				err := vm.Reboot()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("SetMetadata", func() {
		var (
			metadata VMMetadata
		)

		Context("no tags found in metadata", func() {
			BeforeEach(func() {
				metadataBytes := []byte(`{
				  "director": "fake-director-uuid",
				  "name": "fake-director"
				}`)

				metadata = bslvm.VMMetadata{}
				err := json.Unmarshal(metadataBytes, &metadata)
				Expect(err).ToNot(HaveOccurred())
			})

			It("does not set any tag values on the VM", func() {
				err := vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
				Expect(softLayerClient.DoRawHttpRequestResponseCount).To(Equal(0))
			})
		})

		Context("found tags in metadata", func() {
			BeforeEach(func() {
				softLayerClient.DoRawHttpRequestResponse = []byte("true")
			})

			It("the tags value is empty", func() {
				metadataBytes := []byte(`{
				  "director": "fake-director-uuid",
				  "fake1": "fake-deployment",
				  "fake2": "buildpack_python"
				}`)
				metadata = bslvm.VMMetadata{}
				err := json.Unmarshal(metadataBytes, &metadata)
				Expect(err).ToNot(HaveOccurred())

				err = vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
				Expect(softLayerClient.DoRawHttpRequestResponseCount).To(Equal(0))
			})

			It("at least one tag found", func() {
				metadataBytes := []byte(`{
				  "director": "fake-director-uuid",
				  "deployment": "fake-deployment",
				  "compiling": "buildpack_python"
				}`)

				metadata = bslvm.VMMetadata{}
				err := json.Unmarshal(metadataBytes, &metadata)
				Expect(err).ToNot(HaveOccurred())

				err = vm.SetMetadata(metadata)

				Expect(err).ToNot(HaveOccurred())
				Expect(softLayerClient.DoRawHttpRequestResponseCount).To(Equal(1))
			})

			Context("when SLVG.SetTags call fails", func() {
				BeforeEach(func() {
					metadataBytes := []byte(`{
				      "director": "fake-director-uuid",
				      "deployment": "fake-deployment",
				      "compiling": "buildpack_python"
				    }`)

					metadata = bslvm.VMMetadata{}
					err := json.Unmarshal(metadataBytes, &metadata)
					Expect(err).ToNot(HaveOccurred())

					softLayerClient.DoRawHttpRequestError = errors.New("fake-error")
				})

				It("fails with error", func() {
					err := vm.SetMetadata(metadata)

					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("ConfigureNetworks", func() {
		var (
			networks Networks
		)

		BeforeEach(func() {
			networks = map[string]Network{
				"fake-network0": Network{
					Type:    "fake-type",
					IP:      "fake-IP",
					Netmask: "fake-Netmask",
					Gateway: "fake-Gateway",
					DNS: []string{
						"fake-dns0",
						"fake-dns1",
					},
					Default:         []string{},
					CloudProperties: map[string]interface{}{},
				},
			}

			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			testhelpers.SetTestFixtureForFakeSoftLayerClient(softLayerClient, "SoftLayer_Virtual_Guest_Service_getObject.json")
		})

		It("returns the expected network", func() {
			err := vm.ConfigureNetworks(networks)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#AttachDisk", func() {
		var (
			disk bsldisk.Disk
		)

		const expectedDmSetupLs1 = `
36090a0c8600058baa8283574c302c0fc-part1	(252:1)
36090a0c8600058baa8283574c302c0fc	(252:0)
`
		const expectedDmSetupLs2 = `
36090a0c8600058baa8283574c302c0fc-part1	(252:1)
36090a0c8600058baa8283574c302c0fc	(252:0)
36090a0c8600058baa8283574c302c0fd-part1	(252:1)
36090a0c8600058baa8283574c302c0fd	(252:0)
`
		const expectedPartitions1 = `major minor  #blocks  name

   7        0     131072 loop0
 202       16    2097152 xvdb
 202       17    2096451 xvdb1
 202        0   26214400 xvda
 202        1     248832 xvda1
 202        2   25963520 xvda2
 202       32  314572800 xvdc
 202       33  314568733 xvdc1
`
		const expectedPartitions2 = `major minor  #blocks  name

   7        0     131072 loop0
 202       16    2097152 xvdb
 202       17    2096451 xvdb1
 202        0   26214400 xvda
 202        1     248832 xvda1
 202        2   25963520 xvda2
 202       32  314572800 xvdc
 202       33  314568733 xvdc1
 252        0  314572800 dm-0
 252        1  314572799 dm-1
   8       16  314572800 sdb
   8       17  314572799 sdb1
`

		BeforeEach(func() {
			disk = fakedisk.NewFakeDisk(1234)
			fileNames := []string{
				"SoftLayer_Virtual_Guest_Service_getObject.json",
				"SoftLayer_Network_Storage_Service_getIscsiVolume.json",
				"SoftLayer_Network_Storage_Service_getAllowedVirtualGuests_None.json",
				"SoftLayer_Network_Storage_Service_allowAccessFromVirtualGuest.json",
				"SoftLayer_Virtual_Guest_Service_getAllowedHost.json",
				"SoftLayer_Network_Storage_Allowed_Host_Service_getCredential.json",
				"SoftLayer_Virtual_Guest_Service_getUserData_Without_PersistentDisk.json",
				"SoftLayer_Virtual_Guest_Service_getPowerState.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
				"SoftLayer_Virtual_Guest_Service_setMetadata.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
				"SoftLayer_Virtual_Guest_Service_configureMetadataDisk.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
				"SoftLayer_Virtual_Guest_Service_isNotPingable.json",
				"SoftLayer_Virtual_Guest_Service_getPowerState.json",
			}
			testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
		})

		It("attaches the iSCSI volume successfully (multipath-tool installed)", func() {
			expectedCmdResults := []string{
				"/sbin/multipath",
				"No devices found",
				"",
				"",
				"",
				"",
				"",
				expectedDmSetupLs1,
			}
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, expectedCmdResults, nil)
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.AttachDisk(disk)
			Expect(err).ToNot(HaveOccurred())
		})

		It("attaches the iSCSI volume successfully (multipath-tool not installed)", func() {
			expectedCmdResults := []string{
				"",
				expectedPartitions1,
				"",
				"",
				"",
				"",
				"",
				expectedPartitions2,
			}
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, expectedCmdResults, nil)
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.AttachDisk(disk)
			Expect(err).ToNot(HaveOccurred())
		})

		It("attaches second iSCSI volume successfully (multipath-tool installed)", func() {
			expectedCmdResults := []string{
				"/sbin/multipath",
				expectedDmSetupLs1,
				"",
				"",
				"",
				"",
				"",
				expectedDmSetupLs2,
			}
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, expectedCmdResults, nil)
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.AttachDisk(disk)
			Expect(err).ToNot(HaveOccurred())
		})

		It("reports error when failed to attach the iSCSI volume", func() {
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, []string{"fake-result"}, errors.New("fake-error"))
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.AttachDisk(disk)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("#DetachDisk", func() {
		var (
			disk bsldisk.Disk
		)

		const expectMultipathInstalled = `/sbin/multipath
`

		const expectLogoutIscsi = `Logging out of session [sid: 1, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260]
Logging out of session [sid: 2, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260]
Logout of [sid: 1, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260] successful.
Logout of [sid: 2, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260] successful.
`
		const expectLoginIscsi = `Logging in to [iface: default, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260] (multiple)
Logging in to [iface: default, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260] (multiple)
Login to [iface: default, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260] successful.
Login to [iface: default, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260] successful.
`
		const expectStopOpenIscsi = `* Unmounting iscsi-backed filesystems                                                                                                                                                               [ OK ]
 * Disconnecting iSCSI targets                                                                                                                                                                              Logging out of session [sid: 3, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260]
Logging out of session [sid: 4, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260]
Logout of [sid: 3, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.67,3260] successful.
Logout of [sid: 4, target: iqn.1992-08.com.netapp:lon0201, portal: 10.1.222.52,3260] successful.
                                                                                                                                                                                                     [ OK ]
 * Stopping iSCSI initiator service
 `
		const expectStartOpenIscsi = `* Starting iSCSI initiator service iscsid                                                                                                                                                           [ OK ]
 * Setting up iSCSI targets
iscsiadm: No records found
 * Mounting network filesystems
 `
		const expectRestartMultipathd = `* Stopping multipath daemon multipathd                                                                                                                                                              [ OK ]
 * Starting multipath daemon multipathd
 `
		BeforeEach(func() {
			disk = fakedisk.NewFakeDisk(1234)
			fileNames := []string{
				"SoftLayer_Virtual_Guest_Service_getObject.json",
				"SoftLayer_Network_Storage_Service_getIscsiVolume.json",
				"SoftLayer_Network_Storage_Service_getAllowedVirtualGuests.json",
				"SoftLayer_Network_Storage_Service_removeAccessFromVirtualGuest.json",
				"SoftLayer_Virtual_Guest_Service_getUserData_With_PersistentDisk.json",
				"SoftLayer_Virtual_Guest_Service_getPowerState.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
				"SoftLayer_Virtual_Guest_Service_setMetadata.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions_None.json",
				"SoftLayer_Virtual_Guest_Service_configureMetadataDisk.json",
				"SoftLayer_Virtual_Guest_Service_getActiveTransactions.json",
				"SoftLayer_Virtual_Guest_Service_isNotPingable.json",
				"SoftLayer_Virtual_Guest_Service_getPowerState.json",
			}
			testhelpers.SetTestFixturesForFakeSoftLayerClient(softLayerClient, fileNames)
		})

		It("detaches iSCSI volume successfully without multipath-tools installed (one volume attached)", func() {
			expectedCmdResults := []string{
				"",
				expectStopOpenIscsi,
				"",
				"",
				expectStartOpenIscsi,
			}
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, expectedCmdResults, nil)
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.DetachDisk(disk)
			Expect(err).ToNot(HaveOccurred())
		})

		It("detaches iSCSI volume successfully with multipath-tools installed (one volume attached)", func() {
			expectedCmdResults := []string{
				expectMultipathInstalled,
				expectStopOpenIscsi,
				"",
				"",
				expectStartOpenIscsi,
				expectRestartMultipathd,
			}
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, expectedCmdResults, nil)
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.DetachDisk(disk)
			Expect(err).ToNot(HaveOccurred())
		})

		It("reports error when failed to detach iSCSI volume", func() {
			testhelpers.SetTestFixturesForFakeSSHClient(sshClient, []string{"fake-result"}, errors.New("fake-error"))
			vm = NewSoftLayerVM(1234567, softLayerClient, sshClient, agentEnvService, logger)
			bslcommon.TIMEOUT = 2 * time.Second
			bslcommon.POLLING_INTERVAL = 1 * time.Second

			err := vm.DetachDisk(disk)
			Expect(err).To(HaveOccurred())
		})
	})
})
