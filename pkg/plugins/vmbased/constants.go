// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package vmbased

const (
	ActionRunInstances       = "RunInstances"
	ActionStartInstances     = "StartInstances"
	ActionStopInstances      = "StopInstances"
	ActionTerminateInstances = "TerminateInstances"

	ActionCreateVolumes = "CreateVolumes"
	ActionAttachVolumes = "AttachVolumes"
	ActionDetachVolumes = "DetachVolumes"
	ActionDeleteVolumes = "DeleteVolumes"
	ActionResizeVolumes = "ResizeVolumes"

	ActionFormatAndMountVolume      = "FormatAndMountVolume"
	ActionWaitFrontgateAvailable    = "WaitFrontgateAvailable"
	ActionRegisterMetadata          = "RegisterMetadata"
	ActionDeregisterMetadata        = "DeregisterMetadata"
	ActionRegisterCmd               = "RegisterCmd"
	ActionDeregisterCmd             = "DeregisterCmd"
	ActionStartConfd                = "StartConfd"
	ActionStopConfd                 = "StopConfd"
	ActionSetFrontgateConfig        = "SetFrontgateConfig"
	ActionSetDroneConfig            = "SetDroneConfig"
	ActionPingDrone                 = "PingDrone"
	ActionPingFrontgate             = "PingFrontgate"
	ActionRunCommandOnDrone         = "RunCommandOnDrone"
	ActionRemoveContainerOnDrone    = "RemoveContainerOnDrone"
	ActionRunCommandOnFrontgateNode = "RunCommandOnFrontgateNode"
)

const (
	RegisterClustersRootPath         = "clusters"
	RegisterNodeHosts                = "hosts"
	RegisterNodeHost                 = "host"
	RegisterNodeCluster              = "cluster"
	RegisterNodeEnv                  = "env"
	RegisterNodeLoadbalancer         = "loadbalancer"
	RegisterNodeCmd                  = "cmd"
	RegisterNodeCmdId                = "id"
	RegisterNodeCmdTimeout           = "timeout"
	RegisterNodeEndpoint             = "endpoints"
	RegisterNodeAdding               = "adding-hosts"
	RegisterNodeDeleting             = "deleting-hosts"
	RegisterNodeVerticalScalingRoles = "vertical-scaling-roles"
)

const (
	// second
	TimeoutStartConfd           = 60
	TimeoutStopConfd            = 60
	TimeoutDeregister           = 60
	TimeoutRegister             = 60
	TimeoutFormatAndMountVolume = 600
	TimeoutUmountVolume         = 120
	TimeoutSshKeygen            = 120
	TimeoutRemoveContainer      = 120
)

const (
	OpenPitrixBasePath = "/opt/openpitrix/"
	OpenPitrixConfPath = OpenPitrixBasePath + "conf/"
	OpenPitrixSbinPath = OpenPitrixBasePath + "sbin/"
	OpenPitrixConfFile = "openpitrix.conf"
	DroneConfFile      = "drone.conf"
	FrontgateConfFile  = "frontgate.conf"
	UpdateFstabFile    = "update_fstab.sh"
	ConfdPath          = "/etc/confd/"
	MetadataLogLevel   = "debug"
	ConfdBackendType   = "libconfd-backend-etcdv3"
	ConfdCmdLogPath    = "/opt/openpitrix/log/cmd.log"
	HostCmdPrefix      = "nsenter -t 1 -m -u -n -i sh -c"
)
