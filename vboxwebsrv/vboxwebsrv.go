package vboxwebsrv

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SettingsVersion string

const (
	SettingsVersionNull SettingsVersion = "Null"

	SettingsVersionV10 SettingsVersion = "v10"

	SettingsVersionV11 SettingsVersion = "v11"

	SettingsVersionV12 SettingsVersion = "v12"

	SettingsVersionV13pre SettingsVersion = "v13pre"

	SettingsVersionV13 SettingsVersion = "v13"

	SettingsVersionV14 SettingsVersion = "v14"

	SettingsVersionV15 SettingsVersion = "v15"

	SettingsVersionV16 SettingsVersion = "v16"

	SettingsVersionV17 SettingsVersion = "v17"

	SettingsVersionV18 SettingsVersion = "v18"

	SettingsVersionV19 SettingsVersion = "v19"

	SettingsVersionV110 SettingsVersion = "v110"

	SettingsVersionV111 SettingsVersion = "v111"

	SettingsVersionV112 SettingsVersion = "v112"

	SettingsVersionV113 SettingsVersion = "v113"

	SettingsVersionV114 SettingsVersion = "v114"

	SettingsVersionFuture SettingsVersion = "Future"
)

type AccessMode string

const (
	AccessModeReadOnly AccessMode = "ReadOnly"

	AccessModeReadWrite AccessMode = "ReadWrite"
)

type MachineState string

const (
	MachineStateNull MachineState = "Null"

	MachineStatePoweredOff MachineState = "PoweredOff"

	MachineStateSaved MachineState = "Saved"

	MachineStateTeleported MachineState = "Teleported"

	MachineStateAborted MachineState = "Aborted"

	MachineStateRunning MachineState = "Running"

	MachineStatePaused MachineState = "Paused"

	MachineStateStuck MachineState = "Stuck"

	MachineStateTeleporting MachineState = "Teleporting"

	MachineStateLiveSnapshotting MachineState = "LiveSnapshotting"

	MachineStateStarting MachineState = "Starting"

	MachineStateStopping MachineState = "Stopping"

	MachineStateSaving MachineState = "Saving"

	MachineStateRestoring MachineState = "Restoring"

	MachineStateTeleportingPausedVM MachineState = "TeleportingPausedVM"

	MachineStateTeleportingIn MachineState = "TeleportingIn"

	MachineStateFaultTolerantSyncing MachineState = "FaultTolerantSyncing"

	MachineStateDeletingSnapshotOnline MachineState = "DeletingSnapshotOnline"

	MachineStateDeletingSnapshotPaused MachineState = "DeletingSnapshotPaused"

	MachineStateRestoringSnapshot MachineState = "RestoringSnapshot"

	MachineStateDeletingSnapshot MachineState = "DeletingSnapshot"

	MachineStateSettingUp MachineState = "SettingUp"

	MachineStateFirstOnline MachineState = "FirstOnline"

	MachineStateLastOnline MachineState = "LastOnline"

	MachineStateFirstTransient MachineState = "FirstTransient"

	MachineStateLastTransient MachineState = "LastTransient"
)

type SessionState string

const (
	SessionStateNull SessionState = "Null"

	SessionStateUnlocked SessionState = "Unlocked"

	SessionStateLocked SessionState = "Locked"

	SessionStateSpawning SessionState = "Spawning"

	SessionStateUnlocking SessionState = "Unlocking"
)

type CPUPropertyType string

const (
	CPUPropertyTypeNull CPUPropertyType = "Null"

	CPUPropertyTypePAE CPUPropertyType = "PAE"

	CPUPropertyTypeSynthetic CPUPropertyType = "Synthetic"

	CPUPropertyTypeLongMode CPUPropertyType = "LongMode"

	CPUPropertyTypeTripleFaultReset CPUPropertyType = "TripleFaultReset"
)

type HWVirtExPropertyType string

const (
	HWVirtExPropertyTypeNull HWVirtExPropertyType = "Null"

	HWVirtExPropertyTypeEnabled HWVirtExPropertyType = "Enabled"

	HWVirtExPropertyTypeVPID HWVirtExPropertyType = "VPID"

	HWVirtExPropertyTypeNestedPaging HWVirtExPropertyType = "NestedPaging"

	HWVirtExPropertyTypeUnrestrictedExecution HWVirtExPropertyType = "UnrestrictedExecution"

	HWVirtExPropertyTypeLargePages HWVirtExPropertyType = "LargePages"

	HWVirtExPropertyTypeForce HWVirtExPropertyType = "Force"
)

type FaultToleranceState string

const (
	FaultToleranceStateInactive FaultToleranceState = "Inactive"

	FaultToleranceStateMaster FaultToleranceState = "Master"

	FaultToleranceStateStandby FaultToleranceState = "Standby"
)

type LockType string

const (
	LockTypeWrite LockType = "Write"

	LockTypeShared LockType = "Shared"

	LockTypeVM LockType = "VM"
)

type SessionType string

const (
	SessionTypeNull SessionType = "Null"

	SessionTypeWriteLock SessionType = "WriteLock"

	SessionTypeRemote SessionType = "Remote"

	SessionTypeShared SessionType = "Shared"
)

type DeviceType string

const (
	DeviceTypeNull DeviceType = "Null"

	DeviceTypeFloppy DeviceType = "Floppy"

	DeviceTypeDVD DeviceType = "DVD"

	DeviceTypeHardDisk DeviceType = "HardDisk"

	DeviceTypeNetwork DeviceType = "Network"

	DeviceTypeUSB DeviceType = "USB"

	DeviceTypeSharedFolder DeviceType = "SharedFolder"
)

type DeviceActivity string

const (
	DeviceActivityNull DeviceActivity = "Null"

	DeviceActivityIdle DeviceActivity = "Idle"

	DeviceActivityReading DeviceActivity = "Reading"

	DeviceActivityWriting DeviceActivity = "Writing"
)

type ClipboardMode string

const (
	ClipboardModeDisabled ClipboardMode = "Disabled"

	ClipboardModeHostToGuest ClipboardMode = "HostToGuest"

	ClipboardModeGuestToHost ClipboardMode = "GuestToHost"

	ClipboardModeBidirectional ClipboardMode = "Bidirectional"
)

type DragAndDropMode string

const (
	DragAndDropModeDisabled DragAndDropMode = "Disabled"

	DragAndDropModeHostToGuest DragAndDropMode = "HostToGuest"

	DragAndDropModeGuestToHost DragAndDropMode = "GuestToHost"

	DragAndDropModeBidirectional DragAndDropMode = "Bidirectional"
)

type Scope string

const (
	ScopeGlobal Scope = "Global"

	ScopeMachine Scope = "Machine"

	ScopeSession Scope = "Session"
)

type BIOSBootMenuMode string

const (
	BIOSBootMenuModeDisabled BIOSBootMenuMode = "Disabled"

	BIOSBootMenuModeMenuOnly BIOSBootMenuMode = "MenuOnly"

	BIOSBootMenuModeMessageAndMenu BIOSBootMenuMode = "MessageAndMenu"
)

type ProcessorFeature string

const (
	ProcessorFeatureHWVirtEx ProcessorFeature = "HWVirtEx"

	ProcessorFeaturePAE ProcessorFeature = "PAE"

	ProcessorFeatureLongMode ProcessorFeature = "LongMode"

	ProcessorFeatureNestedPaging ProcessorFeature = "NestedPaging"
)

type FirmwareType string

const (
	FirmwareTypeBIOS FirmwareType = "BIOS"

	FirmwareTypeEFI FirmwareType = "EFI"

	FirmwareTypeEFI32 FirmwareType = "EFI32"

	FirmwareTypeEFI64 FirmwareType = "EFI64"

	FirmwareTypeEFIDUAL FirmwareType = "EFIDUAL"
)

type PointingHIDType string

const (
	PointingHIDTypeNone PointingHIDType = "None"

	PointingHIDTypePS2Mouse PointingHIDType = "PS2Mouse"

	PointingHIDTypeUSBMouse PointingHIDType = "USBMouse"

	PointingHIDTypeUSBTablet PointingHIDType = "USBTablet"

	PointingHIDTypeComboMouse PointingHIDType = "ComboMouse"

	PointingHIDTypeUSBMultiTouch PointingHIDType = "USBMultiTouch"
)

type KeyboardHIDType string

const (
	KeyboardHIDTypeNone KeyboardHIDType = "None"

	KeyboardHIDTypePS2Keyboard KeyboardHIDType = "PS2Keyboard"

	KeyboardHIDTypeUSBKeyboard KeyboardHIDType = "USBKeyboard"

	KeyboardHIDTypeComboKeyboard KeyboardHIDType = "ComboKeyboard"
)

type DhcpOpt string

const (
	DhcpOptSubnetMask DhcpOpt = "SubnetMask"

	DhcpOptTimeOffset DhcpOpt = "TimeOffset"

	DhcpOptRouter DhcpOpt = "Router"

	DhcpOptTimeServer DhcpOpt = "TimeServer"

	DhcpOptNameServer DhcpOpt = "NameServer"

	DhcpOptDomainNameServer DhcpOpt = "DomainNameServer"

	DhcpOptLogServer DhcpOpt = "LogServer"

	DhcpOptCookie DhcpOpt = "Cookie"

	DhcpOptLPRServer DhcpOpt = "LPRServer"

	DhcpOptImpressServer DhcpOpt = "ImpressServer"

	DhcpOptResourseLocationServer DhcpOpt = "ResourseLocationServer"

	DhcpOptHostName DhcpOpt = "HostName"

	DhcpOptBootFileSize DhcpOpt = "BootFileSize"

	DhcpOptMeritDumpFile DhcpOpt = "MeritDumpFile"

	DhcpOptDomainName DhcpOpt = "DomainName"

	DhcpOptSwapServer DhcpOpt = "SwapServer"

	DhcpOptRootPath DhcpOpt = "RootPath"

	DhcpOptExtensionPath DhcpOpt = "ExtensionPath"

	DhcpOptIPForwardingEnableDisable DhcpOpt = "IPForwardingEnableDisable"

	DhcpOptNonLocalSourceRoutingEnableDisable DhcpOpt = "NonLocalSourceRoutingEnableDisable"

	DhcpOptPolicyFilter DhcpOpt = "PolicyFilter"

	DhcpOptMaximumDatagramReassemblySize DhcpOpt = "MaximumDatagramReassemblySize"

	DhcpOptDefaultIPTime2Live DhcpOpt = "DefaultIPTime2Live"

	DhcpOptPathMTUAgingTimeout DhcpOpt = "PathMTUAgingTimeout"

	DhcpOptIPLayerParametersPerInterface DhcpOpt = "IPLayerParametersPerInterface"

	DhcpOptInterfaceMTU DhcpOpt = "InterfaceMTU"

	DhcpOptAllSubnetsAreLocal DhcpOpt = "AllSubnetsAreLocal"

	DhcpOptBroadcastAddress DhcpOpt = "BroadcastAddress"

	DhcpOptPerformMaskDiscovery DhcpOpt = "PerformMaskDiscovery"

	DhcpOptMaskSupplier DhcpOpt = "MaskSupplier"

	DhcpOptPerformRouteDiscovery DhcpOpt = "PerformRouteDiscovery"

	DhcpOptRouterSolicitationAddress DhcpOpt = "RouterSolicitationAddress"

	DhcpOptStaticRoute DhcpOpt = "StaticRoute"

	DhcpOptTrailerEncapsulation DhcpOpt = "TrailerEncapsulation"

	DhcpOptARPCacheTimeout DhcpOpt = "ARPCacheTimeout"

	DhcpOptEthernetEncapsulation DhcpOpt = "EthernetEncapsulation"

	DhcpOptTCPDefaultTTL DhcpOpt = "TCPDefaultTTL"

	DhcpOptTCPKeepAliveInterval DhcpOpt = "TCPKeepAliveInterval"

	DhcpOptTCPKeepAliveGarbage DhcpOpt = "TCPKeepAliveGarbage"

	DhcpOptNetworkInformationServiceDomain DhcpOpt = "NetworkInformationServiceDomain"

	DhcpOptNetworkInformationServiceServers DhcpOpt = "NetworkInformationServiceServers"

	DhcpOptNetworkTimeProtocolServers DhcpOpt = "NetworkTimeProtocolServers"

	DhcpOptVendorSpecificInformation DhcpOpt = "VendorSpecificInformation"

	DhcpOptOption44 DhcpOpt = "Option44"

	DhcpOptOption45 DhcpOpt = "Option45"

	DhcpOptOption46 DhcpOpt = "Option46"

	DhcpOptOption47 DhcpOpt = "Option47"

	DhcpOptOption48 DhcpOpt = "Option48"

	DhcpOptOption49 DhcpOpt = "Option49"

	DhcpOptIPAddressLeaseTime DhcpOpt = "IPAddressLeaseTime"

	DhcpOptOption64 DhcpOpt = "Option64"

	DhcpOptOption65 DhcpOpt = "Option65"

	DhcpOptTFTPServerName DhcpOpt = "TFTPServerName"

	DhcpOptBootfileName DhcpOpt = "BootfileName"

	DhcpOptOption68 DhcpOpt = "Option68"

	DhcpOptOption69 DhcpOpt = "Option69"

	DhcpOptOption70 DhcpOpt = "Option70"

	DhcpOptOption71 DhcpOpt = "Option71"

	DhcpOptOption72 DhcpOpt = "Option72"

	DhcpOptOption73 DhcpOpt = "Option73"

	DhcpOptOption74 DhcpOpt = "Option74"

	DhcpOptOption75 DhcpOpt = "Option75"

	DhcpOptOption119 DhcpOpt = "Option119"
)

type VFSType string

const (
	VFSTypeFile VFSType = "File"

	VFSTypeCloud VFSType = "Cloud"

	VFSTypeS3 VFSType = "S3"

	VFSTypeWebDav VFSType = "WebDav"
)

type VFSFileType string

const (
	VFSFileTypeUnknown VFSFileType = "Unknown"

	VFSFileTypeFifo VFSFileType = "Fifo"

	VFSFileTypeDevChar VFSFileType = "DevChar"

	VFSFileTypeDirectory VFSFileType = "Directory"

	VFSFileTypeDevBlock VFSFileType = "DevBlock"

	VFSFileTypeFile VFSFileType = "File"

	VFSFileTypeSymLink VFSFileType = "SymLink"

	VFSFileTypeSocket VFSFileType = "Socket"

	VFSFileTypeWhiteOut VFSFileType = "WhiteOut"
)

type ImportOptions string

const (
	ImportOptionsKeepAllMACs ImportOptions = "KeepAllMACs"

	ImportOptionsKeepNATMACs ImportOptions = "KeepNATMACs"
)

type ExportOptions string

const (
	ExportOptionsCreateManifest ExportOptions = "CreateManifest"

	ExportOptionsExportDVDImages ExportOptions = "ExportDVDImages"

	ExportOptionsStripAllMACs ExportOptions = "StripAllMACs"

	ExportOptionsStripAllNonNATMACs ExportOptions = "StripAllNonNATMACs"
)

type VirtualSystemDescriptionType string

const (
	VirtualSystemDescriptionTypeIgnore VirtualSystemDescriptionType = "Ignore"

	VirtualSystemDescriptionTypeOS VirtualSystemDescriptionType = "OS"

	VirtualSystemDescriptionTypeName VirtualSystemDescriptionType = "Name"

	VirtualSystemDescriptionTypeProduct VirtualSystemDescriptionType = "Product"

	VirtualSystemDescriptionTypeVendor VirtualSystemDescriptionType = "Vendor"

	VirtualSystemDescriptionTypeVersion VirtualSystemDescriptionType = "Version"

	VirtualSystemDescriptionTypeProductUrl VirtualSystemDescriptionType = "ProductUrl"

	VirtualSystemDescriptionTypeVendorUrl VirtualSystemDescriptionType = "VendorUrl"

	VirtualSystemDescriptionTypeDescription VirtualSystemDescriptionType = "Description"

	VirtualSystemDescriptionTypeLicense VirtualSystemDescriptionType = "License"

	VirtualSystemDescriptionTypeMiscellaneous VirtualSystemDescriptionType = "Miscellaneous"

	VirtualSystemDescriptionTypeCPU VirtualSystemDescriptionType = "CPU"

	VirtualSystemDescriptionTypeMemory VirtualSystemDescriptionType = "Memory"

	VirtualSystemDescriptionTypeHardDiskControllerIDE VirtualSystemDescriptionType = "HardDiskControllerIDE"

	VirtualSystemDescriptionTypeHardDiskControllerSATA VirtualSystemDescriptionType = "HardDiskControllerSATA"

	VirtualSystemDescriptionTypeHardDiskControllerSCSI VirtualSystemDescriptionType = "HardDiskControllerSCSI"

	VirtualSystemDescriptionTypeHardDiskControllerSAS VirtualSystemDescriptionType = "HardDiskControllerSAS"

	VirtualSystemDescriptionTypeHardDiskImage VirtualSystemDescriptionType = "HardDiskImage"

	VirtualSystemDescriptionTypeFloppy VirtualSystemDescriptionType = "Floppy"

	VirtualSystemDescriptionTypeCDROM VirtualSystemDescriptionType = "CDROM"

	VirtualSystemDescriptionTypeNetworkAdapter VirtualSystemDescriptionType = "NetworkAdapter"

	VirtualSystemDescriptionTypeUSBController VirtualSystemDescriptionType = "USBController"

	VirtualSystemDescriptionTypeSoundCard VirtualSystemDescriptionType = "SoundCard"

	VirtualSystemDescriptionTypeSettingsFile VirtualSystemDescriptionType = "SettingsFile"
)

type VirtualSystemDescriptionValueType string

const (
	VirtualSystemDescriptionValueTypeReference VirtualSystemDescriptionValueType = "Reference"

	VirtualSystemDescriptionValueTypeOriginal VirtualSystemDescriptionValueType = "Original"

	VirtualSystemDescriptionValueTypeAuto VirtualSystemDescriptionValueType = "Auto"

	VirtualSystemDescriptionValueTypeExtraConfig VirtualSystemDescriptionValueType = "ExtraConfig"
)

type GraphicsControllerType string

const (
	GraphicsControllerTypeNull GraphicsControllerType = "Null"

	GraphicsControllerTypeVBoxVGA GraphicsControllerType = "VBoxVGA"

	GraphicsControllerTypeVMSVGA GraphicsControllerType = "VMSVGA"
)

type CleanupMode string

const (
	CleanupModeUnregisterOnly CleanupMode = "UnregisterOnly"

	CleanupModeDetachAllReturnNone CleanupMode = "DetachAllReturnNone"

	CleanupModeDetachAllReturnHardDisksOnly CleanupMode = "DetachAllReturnHardDisksOnly"

	CleanupModeFull CleanupMode = "Full"
)

type CloneMode string

const (
	CloneModeMachineState CloneMode = "MachineState"

	CloneModeMachineAndChildStates CloneMode = "MachineAndChildStates"

	CloneModeAllStates CloneMode = "AllStates"
)

type CloneOptions string

const (
	CloneOptionsLink CloneOptions = "Link"

	CloneOptionsKeepAllMACs CloneOptions = "KeepAllMACs"

	CloneOptionsKeepNATMACs CloneOptions = "KeepNATMACs"

	CloneOptionsKeepDiskNames CloneOptions = "KeepDiskNames"
)

type AutostopType string

const (
	AutostopTypeDisabled AutostopType = "Disabled"

	AutostopTypeSaveState AutostopType = "SaveState"

	AutostopTypePowerOff AutostopType = "PowerOff"

	AutostopTypeAcpiShutdown AutostopType = "AcpiShutdown"
)

type HostNetworkInterfaceMediumType string

const (
	HostNetworkInterfaceMediumTypeUnknown HostNetworkInterfaceMediumType = "Unknown"

	HostNetworkInterfaceMediumTypeEthernet HostNetworkInterfaceMediumType = "Ethernet"

	HostNetworkInterfaceMediumTypePPP HostNetworkInterfaceMediumType = "PPP"

	HostNetworkInterfaceMediumTypeSLIP HostNetworkInterfaceMediumType = "SLIP"
)

type HostNetworkInterfaceStatus string

const (
	HostNetworkInterfaceStatusUnknown HostNetworkInterfaceStatus = "Unknown"

	HostNetworkInterfaceStatusUp HostNetworkInterfaceStatus = "Up"

	HostNetworkInterfaceStatusDown HostNetworkInterfaceStatus = "Down"
)

type HostNetworkInterfaceType string

const (
	HostNetworkInterfaceTypeBridged HostNetworkInterfaceType = "Bridged"

	HostNetworkInterfaceTypeHostOnly HostNetworkInterfaceType = "HostOnly"
)

type AdditionsFacilityType string

const (
	AdditionsFacilityTypeNone AdditionsFacilityType = "None"

	AdditionsFacilityTypeVBoxGuestDriver AdditionsFacilityType = "VBoxGuestDriver"

	AdditionsFacilityTypeAutoLogon AdditionsFacilityType = "AutoLogon"

	AdditionsFacilityTypeVBoxService AdditionsFacilityType = "VBoxService"

	AdditionsFacilityTypeVBoxTrayClient AdditionsFacilityType = "VBoxTrayClient"

	AdditionsFacilityTypeSeamless AdditionsFacilityType = "Seamless"

	AdditionsFacilityTypeGraphics AdditionsFacilityType = "Graphics"

	AdditionsFacilityTypeAll AdditionsFacilityType = "All"
)

type AdditionsFacilityClass string

const (
	AdditionsFacilityClassNone AdditionsFacilityClass = "None"

	AdditionsFacilityClassDriver AdditionsFacilityClass = "Driver"

	AdditionsFacilityClassService AdditionsFacilityClass = "Service"

	AdditionsFacilityClassProgram AdditionsFacilityClass = "Program"

	AdditionsFacilityClassFeature AdditionsFacilityClass = "Feature"

	AdditionsFacilityClassThirdParty AdditionsFacilityClass = "ThirdParty"

	AdditionsFacilityClassAll AdditionsFacilityClass = "All"
)

type AdditionsFacilityStatus string

const (
	AdditionsFacilityStatusInactive AdditionsFacilityStatus = "Inactive"

	AdditionsFacilityStatusPaused AdditionsFacilityStatus = "Paused"

	AdditionsFacilityStatusPreInit AdditionsFacilityStatus = "PreInit"

	AdditionsFacilityStatusInit AdditionsFacilityStatus = "Init"

	AdditionsFacilityStatusActive AdditionsFacilityStatus = "Active"

	AdditionsFacilityStatusTerminating AdditionsFacilityStatus = "Terminating"

	AdditionsFacilityStatusTerminated AdditionsFacilityStatus = "Terminated"

	AdditionsFacilityStatusFailed AdditionsFacilityStatus = "Failed"

	AdditionsFacilityStatusUnknown AdditionsFacilityStatus = "Unknown"
)

type AdditionsRunLevelType string

const (
	AdditionsRunLevelTypeNone AdditionsRunLevelType = "None"

	AdditionsRunLevelTypeSystem AdditionsRunLevelType = "System"

	AdditionsRunLevelTypeUserland AdditionsRunLevelType = "Userland"

	AdditionsRunLevelTypeDesktop AdditionsRunLevelType = "Desktop"
)

type AdditionsUpdateFlag string

const (
	AdditionsUpdateFlagNone AdditionsUpdateFlag = "None"

	AdditionsUpdateFlagWaitForUpdateStartOnly AdditionsUpdateFlag = "WaitForUpdateStartOnly"
)

type GuestSessionStatus string

const (
	GuestSessionStatusUndefined GuestSessionStatus = "Undefined"

	GuestSessionStatusStarting GuestSessionStatus = "Starting"

	GuestSessionStatusStarted GuestSessionStatus = "Started"

	GuestSessionStatusTerminating GuestSessionStatus = "Terminating"

	GuestSessionStatusTerminated GuestSessionStatus = "Terminated"

	GuestSessionStatusTimedOutKilled GuestSessionStatus = "TimedOutKilled"

	GuestSessionStatusTimedOutAbnormally GuestSessionStatus = "TimedOutAbnormally"

	GuestSessionStatusDown GuestSessionStatus = "Down"

	GuestSessionStatusError GuestSessionStatus = "Error"
)

type GuestSessionWaitForFlag string

const (
	GuestSessionWaitForFlagNone GuestSessionWaitForFlag = "None"

	GuestSessionWaitForFlagStart GuestSessionWaitForFlag = "Start"

	GuestSessionWaitForFlagTerminate GuestSessionWaitForFlag = "Terminate"

	GuestSessionWaitForFlagStatus GuestSessionWaitForFlag = "Status"
)

type GuestSessionWaitResult string

const (
	GuestSessionWaitResultNone GuestSessionWaitResult = "None"

	GuestSessionWaitResultStart GuestSessionWaitResult = "Start"

	GuestSessionWaitResultTerminate GuestSessionWaitResult = "Terminate"

	GuestSessionWaitResultStatus GuestSessionWaitResult = "Status"

	GuestSessionWaitResultError GuestSessionWaitResult = "Error"

	GuestSessionWaitResultTimeout GuestSessionWaitResult = "Timeout"

	GuestSessionWaitResultWaitFlagNotSupported GuestSessionWaitResult = "WaitFlagNotSupported"
)

type GuestUserState string

const (
	GuestUserStateUnknown GuestUserState = "Unknown"

	GuestUserStateLoggedIn GuestUserState = "LoggedIn"

	GuestUserStateLoggedOut GuestUserState = "LoggedOut"

	GuestUserStateLocked GuestUserState = "Locked"

	GuestUserStateUnlocked GuestUserState = "Unlocked"

	GuestUserStateDisabled GuestUserState = "Disabled"

	GuestUserStateIdle GuestUserState = "Idle"

	GuestUserStateInUse GuestUserState = "InUse"

	GuestUserStateCreated GuestUserState = "Created"

	GuestUserStateDeleted GuestUserState = "Deleted"

	GuestUserStateSessionChanged GuestUserState = "SessionChanged"

	GuestUserStateCredentialsChanged GuestUserState = "CredentialsChanged"

	GuestUserStateRoleChanged GuestUserState = "RoleChanged"

	GuestUserStateGroupAdded GuestUserState = "GroupAdded"

	GuestUserStateGroupRemoved GuestUserState = "GroupRemoved"

	GuestUserStateElevated GuestUserState = "Elevated"
)

type FileSeekType string

const (
	FileSeekTypeSet FileSeekType = "Set"

	FileSeekTypeCurrent FileSeekType = "Current"
)

type ProcessInputFlag string

const (
	ProcessInputFlagNone ProcessInputFlag = "None"

	ProcessInputFlagEndOfFile ProcessInputFlag = "EndOfFile"
)

type ProcessOutputFlag string

const (
	ProcessOutputFlagNone ProcessOutputFlag = "None"

	ProcessOutputFlagStdErr ProcessOutputFlag = "StdErr"
)

type ProcessWaitForFlag string

const (
	ProcessWaitForFlagNone ProcessWaitForFlag = "None"

	ProcessWaitForFlagStart ProcessWaitForFlag = "Start"

	ProcessWaitForFlagTerminate ProcessWaitForFlag = "Terminate"

	ProcessWaitForFlagStdIn ProcessWaitForFlag = "StdIn"

	ProcessWaitForFlagStdOut ProcessWaitForFlag = "StdOut"

	ProcessWaitForFlagStdErr ProcessWaitForFlag = "StdErr"
)

type ProcessWaitResult string

const (
	ProcessWaitResultNone ProcessWaitResult = "None"

	ProcessWaitResultStart ProcessWaitResult = "Start"

	ProcessWaitResultTerminate ProcessWaitResult = "Terminate"

	ProcessWaitResultStatus ProcessWaitResult = "Status"

	ProcessWaitResultError ProcessWaitResult = "Error"

	ProcessWaitResultTimeout ProcessWaitResult = "Timeout"

	ProcessWaitResultStdIn ProcessWaitResult = "StdIn"

	ProcessWaitResultStdOut ProcessWaitResult = "StdOut"

	ProcessWaitResultStdErr ProcessWaitResult = "StdErr"

	ProcessWaitResultWaitFlagNotSupported ProcessWaitResult = "WaitFlagNotSupported"
)

type CopyFileFlag string

const (
	CopyFileFlagNone CopyFileFlag = "None"

	CopyFileFlagRecursive CopyFileFlag = "Recursive"

	CopyFileFlagUpdate CopyFileFlag = "Update"

	CopyFileFlagFollowLinks CopyFileFlag = "FollowLinks"
)

type DirectoryCreateFlag string

const (
	DirectoryCreateFlagNone DirectoryCreateFlag = "None"

	DirectoryCreateFlagParents DirectoryCreateFlag = "Parents"
)

type DirectoryRemoveRecFlag string

const (
	DirectoryRemoveRecFlagNone DirectoryRemoveRecFlag = "None"

	DirectoryRemoveRecFlagContentAndDir DirectoryRemoveRecFlag = "ContentAndDir"

	DirectoryRemoveRecFlagContentOnly DirectoryRemoveRecFlag = "ContentOnly"
)

type PathRenameFlag string

const (
	PathRenameFlagNone PathRenameFlag = "None"

	PathRenameFlagNoReplace PathRenameFlag = "NoReplace"

	PathRenameFlagReplace PathRenameFlag = "Replace"

	PathRenameFlagNoSymlinks PathRenameFlag = "NoSymlinks"
)

type ProcessCreateFlag string

const (
	ProcessCreateFlagNone ProcessCreateFlag = "None"

	ProcessCreateFlagWaitForProcessStartOnly ProcessCreateFlag = "WaitForProcessStartOnly"

	ProcessCreateFlagIgnoreOrphanedProcesses ProcessCreateFlag = "IgnoreOrphanedProcesses"

	ProcessCreateFlagHidden ProcessCreateFlag = "Hidden"

	ProcessCreateFlagNoProfile ProcessCreateFlag = "NoProfile"

	ProcessCreateFlagWaitForStdOut ProcessCreateFlag = "WaitForStdOut"

	ProcessCreateFlagWaitForStdErr ProcessCreateFlag = "WaitForStdErr"

	ProcessCreateFlagExpandArguments ProcessCreateFlag = "ExpandArguments"

	ProcessCreateFlagUnquotedArguments ProcessCreateFlag = "UnquotedArguments"
)

type ProcessPriority string

const (
	ProcessPriorityInvalid ProcessPriority = "Invalid"

	ProcessPriorityDefault ProcessPriority = "Default"
)

type SymlinkType string

const (
	SymlinkTypeUnknown SymlinkType = "Unknown"

	SymlinkTypeDirectory SymlinkType = "Directory"

	SymlinkTypeFile SymlinkType = "File"
)

type SymlinkReadFlag string

const (
	SymlinkReadFlagNone SymlinkReadFlag = "None"

	SymlinkReadFlagNoSymlinks SymlinkReadFlag = "NoSymlinks"
)

type ProcessStatus string

const (
	ProcessStatusUndefined ProcessStatus = "Undefined"

	ProcessStatusStarting ProcessStatus = "Starting"

	ProcessStatusStarted ProcessStatus = "Started"

	ProcessStatusPaused ProcessStatus = "Paused"

	ProcessStatusTerminating ProcessStatus = "Terminating"

	ProcessStatusTerminatedNormally ProcessStatus = "TerminatedNormally"

	ProcessStatusTerminatedSignal ProcessStatus = "TerminatedSignal"

	ProcessStatusTerminatedAbnormally ProcessStatus = "TerminatedAbnormally"

	ProcessStatusTimedOutKilled ProcessStatus = "TimedOutKilled"

	ProcessStatusTimedOutAbnormally ProcessStatus = "TimedOutAbnormally"

	ProcessStatusDown ProcessStatus = "Down"

	ProcessStatusError ProcessStatus = "Error"
)

type ProcessInputStatus string

const (
	ProcessInputStatusUndefined ProcessInputStatus = "Undefined"

	ProcessInputStatusBroken ProcessInputStatus = "Broken"

	ProcessInputStatusAvailable ProcessInputStatus = "Available"

	ProcessInputStatusWritten ProcessInputStatus = "Written"

	ProcessInputStatusOverflow ProcessInputStatus = "Overflow"
)

type FileStatus string

const (
	FileStatusUndefined FileStatus = "Undefined"

	FileStatusOpening FileStatus = "Opening"

	FileStatusOpen FileStatus = "Open"

	FileStatusClosing FileStatus = "Closing"

	FileStatusClosed FileStatus = "Closed"

	FileStatusDown FileStatus = "Down"

	FileStatusError FileStatus = "Error"
)

type FsObjType string

const (
	FsObjTypeUndefined FsObjType = "Undefined"

	FsObjTypeFIFO FsObjType = "FIFO"

	FsObjTypeDevChar FsObjType = "DevChar"

	FsObjTypeDevBlock FsObjType = "DevBlock"

	FsObjTypeDirectory FsObjType = "Directory"

	FsObjTypeFile FsObjType = "File"

	FsObjTypeSymlink FsObjType = "Symlink"

	FsObjTypeSocket FsObjType = "Socket"

	FsObjTypeWhiteout FsObjType = "Whiteout"
)

type DragAndDropAction string

const (
	DragAndDropActionIgnore DragAndDropAction = "Ignore"

	DragAndDropActionCopy DragAndDropAction = "Copy"

	DragAndDropActionMove DragAndDropAction = "Move"

	DragAndDropActionLink DragAndDropAction = "Link"
)

type DirectoryOpenFlag string

const (
	DirectoryOpenFlagNone DirectoryOpenFlag = "None"

	DirectoryOpenFlagNoSymlinks DirectoryOpenFlag = "NoSymlinks"
)

type MediumState string

const (
	MediumStateNotCreated MediumState = "NotCreated"

	MediumStateCreated MediumState = "Created"

	MediumStateLockedRead MediumState = "LockedRead"

	MediumStateLockedWrite MediumState = "LockedWrite"

	MediumStateInaccessible MediumState = "Inaccessible"

	MediumStateCreating MediumState = "Creating"

	MediumStateDeleting MediumState = "Deleting"
)

type MediumType string

const (
	MediumTypeNormal MediumType = "Normal"

	MediumTypeImmutable MediumType = "Immutable"

	MediumTypeWritethrough MediumType = "Writethrough"

	MediumTypeShareable MediumType = "Shareable"

	MediumTypeReadonly MediumType = "Readonly"

	MediumTypeMultiAttach MediumType = "MultiAttach"
)

type MediumVariant string

const (
	MediumVariantStandard MediumVariant = "Standard"

	MediumVariantVmdkSplit2G MediumVariant = "VmdkSplit2G"

	MediumVariantVmdkRawDisk MediumVariant = "VmdkRawDisk"

	MediumVariantVmdkStreamOptimized MediumVariant = "VmdkStreamOptimized"

	MediumVariantVmdkESX MediumVariant = "VmdkESX"

	MediumVariantFixed MediumVariant = "Fixed"

	MediumVariantDiff MediumVariant = "Diff"

	MediumVariantNoCreateDir MediumVariant = "NoCreateDir"
)

type DataType string

const (
	DataTypeInt32 DataType = "Int32"

	DataTypeInt8 DataType = "Int8"

	DataTypeString DataType = "String"
)

type DataFlags string

const (
	DataFlagsNone DataFlags = "None"

	DataFlagsMandatory DataFlags = "Mandatory"

	DataFlagsExpert DataFlags = "Expert"

	DataFlagsArray DataFlags = "Array"

	DataFlagsFlagMask DataFlags = "FlagMask"
)

type MediumFormatCapabilities string

const (
	MediumFormatCapabilitiesUuid MediumFormatCapabilities = "Uuid"

	MediumFormatCapabilitiesCreateFixed MediumFormatCapabilities = "CreateFixed"

	MediumFormatCapabilitiesCreateDynamic MediumFormatCapabilities = "CreateDynamic"

	MediumFormatCapabilitiesCreateSplit2G MediumFormatCapabilities = "CreateSplit2G"

	MediumFormatCapabilitiesDifferencing MediumFormatCapabilities = "Differencing"

	MediumFormatCapabilitiesAsynchronous MediumFormatCapabilities = "Asynchronous"

	MediumFormatCapabilitiesFile MediumFormatCapabilities = "File"

	MediumFormatCapabilitiesProperties MediumFormatCapabilities = "Properties"

	MediumFormatCapabilitiesTcpNetworking MediumFormatCapabilities = "TcpNetworking"

	MediumFormatCapabilitiesVFS MediumFormatCapabilities = "VFS"

	MediumFormatCapabilitiesCapabilityMask MediumFormatCapabilities = "CapabilityMask"
)

type MouseButtonState string

const (
	MouseButtonStateLeftButton MouseButtonState = "LeftButton"

	MouseButtonStateRightButton MouseButtonState = "RightButton"

	MouseButtonStateMiddleButton MouseButtonState = "MiddleButton"

	MouseButtonStateWheelUp MouseButtonState = "WheelUp"

	MouseButtonStateWheelDown MouseButtonState = "WheelDown"

	MouseButtonStateXButton1 MouseButtonState = "XButton1"

	MouseButtonStateXButton2 MouseButtonState = "XButton2"

	MouseButtonStateMouseStateMask MouseButtonState = "MouseStateMask"
)

type TouchContactState string

const (
	TouchContactStateNone TouchContactState = "None"

	TouchContactStateInContact TouchContactState = "InContact"

	TouchContactStateInRange TouchContactState = "InRange"

	TouchContactStateContactStateMask TouchContactState = "ContactStateMask"
)

type FramebufferPixelFormat string

const (
	FramebufferPixelFormatOpaque FramebufferPixelFormat = "Opaque"

	FramebufferPixelFormatFOURCCRGB FramebufferPixelFormat = "FOURCCRGB"
)

type NetworkAttachmentType string

const (
	NetworkAttachmentTypeNull NetworkAttachmentType = "Null"

	NetworkAttachmentTypeNAT NetworkAttachmentType = "NAT"

	NetworkAttachmentTypeBridged NetworkAttachmentType = "Bridged"

	NetworkAttachmentTypeInternal NetworkAttachmentType = "Internal"

	NetworkAttachmentTypeHostOnly NetworkAttachmentType = "HostOnly"

	NetworkAttachmentTypeGeneric NetworkAttachmentType = "Generic"

	NetworkAttachmentTypeNATNetwork NetworkAttachmentType = "NATNetwork"
)

type NetworkAdapterType string

const (
	NetworkAdapterTypeNull NetworkAdapterType = "Null"

	NetworkAdapterTypeAm79C970A NetworkAdapterType = "Am79C970A"

	NetworkAdapterTypeAm79C973 NetworkAdapterType = "Am79C973"

	NetworkAdapterTypeI82540EM NetworkAdapterType = "I82540EM"

	NetworkAdapterTypeI82543GC NetworkAdapterType = "I82543GC"

	NetworkAdapterTypeI82545EM NetworkAdapterType = "I82545EM"

	NetworkAdapterTypeVirtio NetworkAdapterType = "Virtio"
)

type NetworkAdapterPromiscModePolicy string

const (
	NetworkAdapterPromiscModePolicyDeny NetworkAdapterPromiscModePolicy = "Deny"

	NetworkAdapterPromiscModePolicyAllowNetwork NetworkAdapterPromiscModePolicy = "AllowNetwork"

	NetworkAdapterPromiscModePolicyAllowAll NetworkAdapterPromiscModePolicy = "AllowAll"
)

type PortMode string

const (
	PortModeDisconnected PortMode = "Disconnected"

	PortModeHostPipe PortMode = "HostPipe"

	PortModeHostDevice PortMode = "HostDevice"

	PortModeRawFile PortMode = "RawFile"
)

type USBControllerType string

const (
	USBControllerTypeNull USBControllerType = "Null"

	USBControllerTypeOHCI USBControllerType = "OHCI"

	USBControllerTypeEHCI USBControllerType = "EHCI"

	USBControllerTypeLast USBControllerType = "Last"
)

type USBDeviceState string

const (
	USBDeviceStateNotSupported USBDeviceState = "NotSupported"

	USBDeviceStateUnavailable USBDeviceState = "Unavailable"

	USBDeviceStateBusy USBDeviceState = "Busy"

	USBDeviceStateAvailable USBDeviceState = "Available"

	USBDeviceStateHeld USBDeviceState = "Held"

	USBDeviceStateCaptured USBDeviceState = "Captured"
)

type USBDeviceFilterAction string

const (
	USBDeviceFilterActionNull USBDeviceFilterAction = "Null"

	USBDeviceFilterActionIgnore USBDeviceFilterAction = "Ignore"

	USBDeviceFilterActionHold USBDeviceFilterAction = "Hold"
)

type AudioDriverType string

const (
	AudioDriverTypeNull AudioDriverType = "Null"

	AudioDriverTypeWinMM AudioDriverType = "WinMM"

	AudioDriverTypeOSS AudioDriverType = "OSS"

	AudioDriverTypeALSA AudioDriverType = "ALSA"

	AudioDriverTypeDirectSound AudioDriverType = "DirectSound"

	AudioDriverTypeCoreAudio AudioDriverType = "CoreAudio"

	AudioDriverTypeMMPM AudioDriverType = "MMPM"

	AudioDriverTypePulse AudioDriverType = "Pulse"

	AudioDriverTypeSolAudio AudioDriverType = "SolAudio"
)

type AudioControllerType string

const (
	AudioControllerTypeAC97 AudioControllerType = "AC97"

	AudioControllerTypeSB16 AudioControllerType = "SB16"

	AudioControllerTypeHDA AudioControllerType = "HDA"
)

type AuthType string

const (
	AuthTypeNull AuthType = "Null"

	AuthTypeExternal AuthType = "External"

	AuthTypeGuest AuthType = "Guest"
)

type Reason string

const (
	ReasonUnspecified Reason = "Unspecified"

	ReasonHostSuspend Reason = "HostSuspend"

	ReasonHostResume Reason = "HostResume"

	ReasonHostBatteryLow Reason = "HostBatteryLow"
)

type StorageBus string

const (
	StorageBusNull StorageBus = "Null"

	StorageBusIDE StorageBus = "IDE"

	StorageBusSATA StorageBus = "SATA"

	StorageBusSCSI StorageBus = "SCSI"

	StorageBusFloppy StorageBus = "Floppy"

	StorageBusSAS StorageBus = "SAS"
)

type StorageControllerType string

const (
	StorageControllerTypeNull StorageControllerType = "Null"

	StorageControllerTypeLsiLogic StorageControllerType = "LsiLogic"

	StorageControllerTypeBusLogic StorageControllerType = "BusLogic"

	StorageControllerTypeIntelAhci StorageControllerType = "IntelAhci"

	StorageControllerTypePIIX3 StorageControllerType = "PIIX3"

	StorageControllerTypePIIX4 StorageControllerType = "PIIX4"

	StorageControllerTypeICH6 StorageControllerType = "ICH6"

	StorageControllerTypeI82078 StorageControllerType = "I82078"

	StorageControllerTypeLsiLogicSas StorageControllerType = "LsiLogicSas"
)

type ChipsetType string

const (
	ChipsetTypeNull ChipsetType = "Null"

	ChipsetTypePIIX3 ChipsetType = "PIIX3"

	ChipsetTypeICH9 ChipsetType = "ICH9"
)

type NATAliasMode string

const (
	NATAliasModeAliasLog NATAliasMode = "AliasLog"

	NATAliasModeAliasProxyOnly NATAliasMode = "AliasProxyOnly"

	NATAliasModeAliasUseSamePorts NATAliasMode = "AliasUseSamePorts"
)

type NATProtocol string

const (
	NATProtocolUDP NATProtocol = "UDP"

	NATProtocolTCP NATProtocol = "TCP"
)

type BandwidthGroupType string

const (
	BandwidthGroupTypeNull BandwidthGroupType = "Null"

	BandwidthGroupTypeDisk BandwidthGroupType = "Disk"

	BandwidthGroupTypeNetwork BandwidthGroupType = "Network"
)

type VBoxEventType string

const (
	VBoxEventTypeInvalid VBoxEventType = "Invalid"

	VBoxEventTypeAny VBoxEventType = "Any"

	VBoxEventTypeVetoable VBoxEventType = "Vetoable"

	VBoxEventTypeMachineEvent VBoxEventType = "MachineEvent"

	VBoxEventTypeSnapshotEvent VBoxEventType = "SnapshotEvent"

	VBoxEventTypeInputEvent VBoxEventType = "InputEvent"

	VBoxEventTypeLastWildcard VBoxEventType = "LastWildcard"

	VBoxEventTypeOnMachineStateChanged VBoxEventType = "OnMachineStateChanged"

	VBoxEventTypeOnMachineDataChanged VBoxEventType = "OnMachineDataChanged"

	VBoxEventTypeOnExtraDataChanged VBoxEventType = "OnExtraDataChanged"

	VBoxEventTypeOnExtraDataCanChange VBoxEventType = "OnExtraDataCanChange"

	VBoxEventTypeOnMediumRegistered VBoxEventType = "OnMediumRegistered"

	VBoxEventTypeOnMachineRegistered VBoxEventType = "OnMachineRegistered"

	VBoxEventTypeOnSessionStateChanged VBoxEventType = "OnSessionStateChanged"

	VBoxEventTypeOnSnapshotTaken VBoxEventType = "OnSnapshotTaken"

	VBoxEventTypeOnSnapshotDeleted VBoxEventType = "OnSnapshotDeleted"

	VBoxEventTypeOnSnapshotChanged VBoxEventType = "OnSnapshotChanged"

	VBoxEventTypeOnGuestPropertyChanged VBoxEventType = "OnGuestPropertyChanged"

	VBoxEventTypeOnMousePointerShapeChanged VBoxEventType = "OnMousePointerShapeChanged"

	VBoxEventTypeOnMouseCapabilityChanged VBoxEventType = "OnMouseCapabilityChanged"

	VBoxEventTypeOnKeyboardLedsChanged VBoxEventType = "OnKeyboardLedsChanged"

	VBoxEventTypeOnStateChanged VBoxEventType = "OnStateChanged"

	VBoxEventTypeOnAdditionsStateChanged VBoxEventType = "OnAdditionsStateChanged"

	VBoxEventTypeOnNetworkAdapterChanged VBoxEventType = "OnNetworkAdapterChanged"

	VBoxEventTypeOnSerialPortChanged VBoxEventType = "OnSerialPortChanged"

	VBoxEventTypeOnParallelPortChanged VBoxEventType = "OnParallelPortChanged"

	VBoxEventTypeOnStorageControllerChanged VBoxEventType = "OnStorageControllerChanged"

	VBoxEventTypeOnMediumChanged VBoxEventType = "OnMediumChanged"

	VBoxEventTypeOnVRDEServerChanged VBoxEventType = "OnVRDEServerChanged"

	VBoxEventTypeOnUSBControllerChanged VBoxEventType = "OnUSBControllerChanged"

	VBoxEventTypeOnUSBDeviceStateChanged VBoxEventType = "OnUSBDeviceStateChanged"

	VBoxEventTypeOnSharedFolderChanged VBoxEventType = "OnSharedFolderChanged"

	VBoxEventTypeOnRuntimeError VBoxEventType = "OnRuntimeError"

	VBoxEventTypeOnCanShowWindow VBoxEventType = "OnCanShowWindow"

	VBoxEventTypeOnShowWindow VBoxEventType = "OnShowWindow"

	VBoxEventTypeOnCPUChanged VBoxEventType = "OnCPUChanged"

	VBoxEventTypeOnVRDEServerInfoChanged VBoxEventType = "OnVRDEServerInfoChanged"

	VBoxEventTypeOnEventSourceChanged VBoxEventType = "OnEventSourceChanged"

	VBoxEventTypeOnCPUExecutionCapChanged VBoxEventType = "OnCPUExecutionCapChanged"

	VBoxEventTypeOnGuestKeyboard VBoxEventType = "OnGuestKeyboard"

	VBoxEventTypeOnGuestMouse VBoxEventType = "OnGuestMouse"

	VBoxEventTypeOnNATRedirect VBoxEventType = "OnNATRedirect"

	VBoxEventTypeOnHostPCIDevicePlug VBoxEventType = "OnHostPCIDevicePlug"

	VBoxEventTypeOnVBoxSVCAvailabilityChanged VBoxEventType = "OnVBoxSVCAvailabilityChanged"

	VBoxEventTypeOnBandwidthGroupChanged VBoxEventType = "OnBandwidthGroupChanged"

	VBoxEventTypeOnGuestMonitorChanged VBoxEventType = "OnGuestMonitorChanged"

	VBoxEventTypeOnStorageDeviceChanged VBoxEventType = "OnStorageDeviceChanged"

	VBoxEventTypeOnClipboardModeChanged VBoxEventType = "OnClipboardModeChanged"

	VBoxEventTypeOnDragAndDropModeChanged VBoxEventType = "OnDragAndDropModeChanged"

	VBoxEventTypeOnNATNetworkChanged VBoxEventType = "OnNATNetworkChanged"

	VBoxEventTypeOnNATNetworkStartStop VBoxEventType = "OnNATNetworkStartStop"

	VBoxEventTypeOnNATNetworkAlter VBoxEventType = "OnNATNetworkAlter"

	VBoxEventTypeOnNATNetworkCreationDeletion VBoxEventType = "OnNATNetworkCreationDeletion"

	VBoxEventTypeOnNATNetworkSetting VBoxEventType = "OnNATNetworkSetting"

	VBoxEventTypeOnNATNetworkPortForward VBoxEventType = "OnNATNetworkPortForward"

	VBoxEventTypeOnGuestSessionStateChanged VBoxEventType = "OnGuestSessionStateChanged"

	VBoxEventTypeOnGuestSessionRegistered VBoxEventType = "OnGuestSessionRegistered"

	VBoxEventTypeOnGuestProcessRegistered VBoxEventType = "OnGuestProcessRegistered"

	VBoxEventTypeOnGuestProcessStateChanged VBoxEventType = "OnGuestProcessStateChanged"

	VBoxEventTypeOnGuestProcessInputNotify VBoxEventType = "OnGuestProcessInputNotify"

	VBoxEventTypeOnGuestProcessOutput VBoxEventType = "OnGuestProcessOutput"

	VBoxEventTypeOnGuestFileRegistered VBoxEventType = "OnGuestFileRegistered"

	VBoxEventTypeOnGuestFileStateChanged VBoxEventType = "OnGuestFileStateChanged"

	VBoxEventTypeOnGuestFileOffsetChanged VBoxEventType = "OnGuestFileOffsetChanged"

	VBoxEventTypeOnGuestFileRead VBoxEventType = "OnGuestFileRead"

	VBoxEventTypeOnGuestFileWrite VBoxEventType = "OnGuestFileWrite"

	VBoxEventTypeOnVideoCaptureChanged VBoxEventType = "OnVideoCaptureChanged"

	VBoxEventTypeOnGuestUserStateChanged VBoxEventType = "OnGuestUserStateChanged"

	VBoxEventTypeOnGuestMultiTouch VBoxEventType = "OnGuestMultiTouch"

	VBoxEventTypeOnHostNameResolutionConfigurationChange VBoxEventType = "OnHostNameResolutionConfigurationChange"

	VBoxEventTypeLast VBoxEventType = "Last"
)

type GuestMouseEventMode string

const (
	GuestMouseEventModeRelative GuestMouseEventMode = "Relative"

	GuestMouseEventModeAbsolute GuestMouseEventMode = "Absolute"
)

type GuestMonitorChangedEventType string

const (
	GuestMonitorChangedEventTypeEnabled GuestMonitorChangedEventType = "Enabled"

	GuestMonitorChangedEventTypeDisabled GuestMonitorChangedEventType = "Disabled"

	GuestMonitorChangedEventTypeNewOrigin GuestMonitorChangedEventType = "NewOrigin"
)

type IVirtualBoxErrorInfogetResultCode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultCode"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetResultCodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultCodeResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetResultDetail struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultDetail"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetResultDetailResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getResultDetailResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetInterfaceID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getInterfaceID"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetInterfaceIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getInterfaceIDResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetComponent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getComponent"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetComponentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getComponentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetText struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getText"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetTextResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getTextResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxErrorInfogetNext struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getNext"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxErrorInfogetNextResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBoxErrorInfo_getNextResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkgetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworksetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNetworkName"`

	This        string `xml:"_this,omitempty"`
	NetworkName string `xml:"networkName,omitempty"`
}

type INATNetworksetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNetworkNameResponse"`
}

type INATNetworkgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworksetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type INATNetworksetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setEnabledResponse"`
}

type INATNetworkgetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNetwork"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworksetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNetwork"`

	This    string `xml:"_this,omitempty"`
	Network string `xml:"network,omitempty"`
}

type INATNetworksetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNetworkResponse"`
}

type INATNetworkgetGateway struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getGateway"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetGatewayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getGatewayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkgetIPv6Enabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getIPv6Enabled"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetIPv6EnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getIPv6EnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworksetIPv6Enabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setIPv6Enabled"`

	This        string `xml:"_this,omitempty"`
	IPv6Enabled bool   `xml:"IPv6Enabled,omitempty"`
}

type INATNetworksetIPv6EnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setIPv6EnabledResponse"`
}

type INATNetworkgetIPv6Prefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getIPv6Prefix"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetIPv6PrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getIPv6PrefixResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworksetIPv6Prefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setIPv6Prefix"`

	This       string `xml:"_this,omitempty"`
	IPv6Prefix string `xml:"IPv6Prefix,omitempty"`
}

type INATNetworksetIPv6PrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setIPv6PrefixResponse"`
}

type INATNetworkgetAdvertiseDefaultIPv6RouteEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getAdvertiseDefaultIPv6RouteEnabled"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetAdvertiseDefaultIPv6RouteEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getAdvertiseDefaultIPv6RouteEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworksetAdvertiseDefaultIPv6RouteEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setAdvertiseDefaultIPv6RouteEnabled"`

	This                             string `xml:"_this,omitempty"`
	AdvertiseDefaultIPv6RouteEnabled bool   `xml:"advertiseDefaultIPv6RouteEnabled,omitempty"`
}

type INATNetworksetAdvertiseDefaultIPv6RouteEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setAdvertiseDefaultIPv6RouteEnabledResponse"`
}

type INATNetworkgetNeedDhcpServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNeedDhcpServer"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetNeedDhcpServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getNeedDhcpServerResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworksetNeedDhcpServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNeedDhcpServer"`

	This           string `xml:"_this,omitempty"`
	NeedDhcpServer bool   `xml:"needDhcpServer,omitempty"`
}

type INATNetworksetNeedDhcpServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setNeedDhcpServerResponse"`
}

type INATNetworkgetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkgetPortForwardRules4 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getPortForwardRules4"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetPortForwardRules4Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getPortForwardRules4Response"`

	Returnval []string `xml:"returnval,omitempty"`
}

type INATNetworkgetLocalMappings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getLocalMappings"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetLocalMappingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getLocalMappingsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type INATNetworkgetLoopbackIp6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getLoopbackIp6"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetLoopbackIp6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getLoopbackIp6Response"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type INATNetworksetLoopbackIp6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setLoopbackIp6"`

	This        string `xml:"_this,omitempty"`
	LoopbackIp6 int32  `xml:"loopbackIp6,omitempty"`
}

type INATNetworksetLoopbackIp6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_setLoopbackIp6Response"`
}

type INATNetworkgetPortForwardRules6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getPortForwardRules6"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkgetPortForwardRules6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_getPortForwardRules6Response"`

	Returnval []string `xml:"returnval,omitempty"`
}

type INATNetworkaddLocalMapping struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_addLocalMapping"`

	This   string `xml:"_this,omitempty"`
	Hostid string `xml:"hostid,omitempty"`
	Offset int32  `xml:"offset,omitempty"`
}

type INATNetworkaddLocalMappingResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_addLocalMappingResponse"`
}

type INATNetworkaddPortForwardRule struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_addPortForwardRule"`

	This      string       `xml:"_this,omitempty"`
	IsIpv6    bool         `xml:"isIpv6,omitempty"`
	RuleName  string       `xml:"ruleName,omitempty"`
	Proto     *NATProtocol `xml:"proto,omitempty"`
	HostIP    string       `xml:"hostIP,omitempty"`
	HostPort  uint16       `xml:"hostPort,omitempty"`
	GuestIP   string       `xml:"guestIP,omitempty"`
	GuestPort uint16       `xml:"guestPort,omitempty"`
}

type INATNetworkaddPortForwardRuleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_addPortForwardRuleResponse"`
}

type INATNetworkremovePortForwardRule struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_removePortForwardRule"`

	This     string `xml:"_this,omitempty"`
	ISipv6   bool   `xml:"iSipv6,omitempty"`
	RuleName string `xml:"ruleName,omitempty"`
}

type INATNetworkremovePortForwardRuleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_removePortForwardRuleResponse"`
}

type INATNetworkstart struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_start"`

	This      string `xml:"_this,omitempty"`
	TrunkType string `xml:"trunkType,omitempty"`
}

type INATNetworkstartResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_startResponse"`
}

type INATNetworkstop struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_stop"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkstopResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetwork_stopResponse"`
}

type IDHCPServergetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IDHCPServersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IDHCPServersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setEnabledResponse"`
}

type IDHCPServergetIPAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getIPAddress"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetIPAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getIPAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetNetworkMask struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkMask"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetNetworkMaskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkMaskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetLowerIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getLowerIP"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetLowerIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getLowerIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetUpperIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getUpperIP"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetUpperIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getUpperIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDHCPServergetGlobalOptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getGlobalOptions"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetGlobalOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getGlobalOptionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IDHCPServergetVmConfigs struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getVmConfigs"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServergetVmConfigsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getVmConfigsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IDHCPServeraddGlobalOption struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_addGlobalOption"`

	This   string   `xml:"_this,omitempty"`
	Option *DhcpOpt `xml:"option,omitempty"`
	Value  string   `xml:"value,omitempty"`
}

type IDHCPServeraddGlobalOptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_addGlobalOptionResponse"`
}

type IDHCPServeraddVmSlotOption struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_addVmSlotOption"`

	This   string   `xml:"_this,omitempty"`
	Vmname string   `xml:"vmname,omitempty"`
	Slot   int32    `xml:"slot,omitempty"`
	Option *DhcpOpt `xml:"option,omitempty"`
	Value  string   `xml:"value,omitempty"`
}

type IDHCPServeraddVmSlotOptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_addVmSlotOptionResponse"`
}

type IDHCPServerremoveVmSlotOptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_removeVmSlotOptions"`

	This   string `xml:"_this,omitempty"`
	Vmname string `xml:"vmname,omitempty"`
	Slot   int32  `xml:"slot,omitempty"`
}

type IDHCPServerremoveVmSlotOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_removeVmSlotOptionsResponse"`
}

type IDHCPServergetVmSlotOptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getVmSlotOptions"`

	This   string `xml:"_this,omitempty"`
	Vmname string `xml:"vmname,omitempty"`
	Slot   int32  `xml:"slot,omitempty"`
}

type IDHCPServergetVmSlotOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getVmSlotOptionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IDHCPServergetMacOptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getMacOptions"`

	This string `xml:"_this,omitempty"`
	Mac  string `xml:"mac,omitempty"`
}

type IDHCPServergetMacOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_getMacOptionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IDHCPServersetConfiguration struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setConfiguration"`

	This          string `xml:"_this,omitempty"`
	IPAddress     string `xml:"IPAddress,omitempty"`
	NetworkMask   string `xml:"networkMask,omitempty"`
	FromIPAddress string `xml:"FromIPAddress,omitempty"`
	ToIPAddress   string `xml:"ToIPAddress,omitempty"`
}

type IDHCPServersetConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_setConfigurationResponse"`
}

type IDHCPServerstart struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_start"`

	This        string `xml:"_this,omitempty"`
	NetworkName string `xml:"networkName,omitempty"`
	TrunkName   string `xml:"trunkName,omitempty"`
	TrunkType   string `xml:"trunkType,omitempty"`
}

type IDHCPServerstartResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_startResponse"`
}

type IDHCPServerstop struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_stop"`

	This string `xml:"_this,omitempty"`
}

type IDHCPServerstopResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDHCPServer_stopResponse"`
}

type IVirtualBoxgetVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersion"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetVersionNormalized struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersionNormalized"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetVersionNormalizedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getVersionNormalizedResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getRevisionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVirtualBoxgetPackageType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPackageType"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetPackageTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPackageTypeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetAPIVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getAPIVersion"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetAPIVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getAPIVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHomeFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHomeFolder"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHomeFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHomeFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSettingsFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSettingsFilePath"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSettingsFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSettingsFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHost struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHost"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHostResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHostResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSystemProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSystemProperties"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSystemPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSystemPropertiesResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetMachines struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachines"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetMachinesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachinesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetMachineGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachineGroups"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetMachineGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachineGroupsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetHardDisks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDisks"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetHardDisksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getHardDisksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetDVDImages struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImages"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetDVDImagesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDVDImagesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetFloppyImages struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImages"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetFloppyImagesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getFloppyImagesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetProgressOperations struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getProgressOperations"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetProgressOperationsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getProgressOperationsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetGuestOSTypes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypes"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetGuestOSTypesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypesResponse"`

	Returnval []*IGuestOSType `xml:"returnval,omitempty"`
}

type IVirtualBoxgetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IVirtualBoxgetPerformanceCollector struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPerformanceCollector"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetPerformanceCollectorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getPerformanceCollectorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetDHCPServers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDHCPServers"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetDHCPServersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getDHCPServersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetNATNetworks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getNATNetworks"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetNATNetworksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getNATNetworksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetInternalNetworks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getInternalNetworks"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetInternalNetworksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getInternalNetworksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetGenericNetworkDrivers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGenericNetworkDrivers"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetGenericNetworkDriversResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGenericNetworkDriversResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxcomposeMachineFilename struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_composeMachineFilename"`

	This        string `xml:"_this,omitempty"`
	Name        string `xml:"name,omitempty"`
	Group       string `xml:"group,omitempty"`
	CreateFlags string `xml:"createFlags,omitempty"`
	BaseFolder  string `xml:"baseFolder,omitempty"`
}

type IVirtualBoxcomposeMachineFilenameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_composeMachineFilenameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createMachine"`

	This         string   `xml:"_this,omitempty"`
	SettingsFile string   `xml:"settingsFile,omitempty"`
	Name         string   `xml:"name,omitempty"`
	Groups       []string `xml:"groups,omitempty"`
	OsTypeId     string   `xml:"osTypeId,omitempty"`
	Flags        string   `xml:"flags,omitempty"`
}

type IVirtualBoxcreateMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMachine"`

	This         string `xml:"_this,omitempty"`
	SettingsFile string `xml:"settingsFile,omitempty"`
}

type IVirtualBoxopenMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxregisterMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_registerMachine"`

	This    string `xml:"_this,omitempty"`
	Machine string `xml:"machine,omitempty"`
}

type IVirtualBoxregisterMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_registerMachineResponse"`
}

type IVirtualBoxfindMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findMachine"`

	This     string `xml:"_this,omitempty"`
	NameOrId string `xml:"nameOrId,omitempty"`
}

type IVirtualBoxfindMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetMachinesByGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachinesByGroups"`

	This   string   `xml:"_this,omitempty"`
	Groups []string `xml:"groups,omitempty"`
}

type IVirtualBoxgetMachinesByGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachinesByGroupsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetMachineStates struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachineStates"`

	This     string   `xml:"_this,omitempty"`
	Machines []string `xml:"machines,omitempty"`
}

type IVirtualBoxgetMachineStatesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getMachineStatesResponse"`

	Returnval []*MachineState `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateAppliance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createAppliance"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxcreateApplianceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createApplianceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateHardDisk struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createHardDisk"`

	This     string `xml:"_this,omitempty"`
	Format   string `xml:"format,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IVirtualBoxcreateHardDiskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createHardDiskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxopenMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMedium"`

	This         string      `xml:"_this,omitempty"`
	Location     string      `xml:"location,omitempty"`
	DeviceType   *DeviceType `xml:"deviceType,omitempty"`
	AccessMode   *AccessMode `xml:"accessMode,omitempty"`
	ForceNewUuid bool        `xml:"forceNewUuid,omitempty"`
}

type IVirtualBoxopenMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_openMediumResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetGuestOSType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSType"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IVirtualBoxgetGuestOSTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getGuestOSTypeResponse"`

	Returnval *IGuestOSType `xml:"returnval,omitempty"`
}

type IVirtualBoxcreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createSharedFolder"`

	This      string `xml:"_this,omitempty"`
	Name      string `xml:"name,omitempty"`
	HostPath  string `xml:"hostPath,omitempty"`
	Writable  bool   `xml:"writable,omitempty"`
	Automount bool   `xml:"automount,omitempty"`
}

type IVirtualBoxcreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createSharedFolderResponse"`
}

type IVirtualBoxremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeSharedFolderResponse"`
}

type IVirtualBoxgetExtraDataKeys struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataKeys"`

	This string `xml:"_this,omitempty"`
}

type IVirtualBoxgetExtraDataKeysResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataKeysResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualBoxgetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraData"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type IVirtualBoxgetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_getExtraDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxsetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setExtraData"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IVirtualBoxsetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setExtraDataResponse"`
}

type IVirtualBoxsetSettingsSecret struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setSettingsSecret"`

	This     string `xml:"_this,omitempty"`
	Password string `xml:"password,omitempty"`
}

type IVirtualBoxsetSettingsSecretResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_setSettingsSecretResponse"`
}

type IVirtualBoxcreateDHCPServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createDHCPServer"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxcreateDHCPServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createDHCPServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindDHCPServerByNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDHCPServerByNetworkName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IVirtualBoxfindDHCPServerByNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findDHCPServerByNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxremoveDHCPServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeDHCPServer"`

	This   string `xml:"_this,omitempty"`
	Server string `xml:"server,omitempty"`
}

type IVirtualBoxremoveDHCPServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeDHCPServerResponse"`
}

type IVirtualBoxcreateNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createNATNetwork"`

	This        string `xml:"_this,omitempty"`
	NetworkName string `xml:"networkName,omitempty"`
}

type IVirtualBoxcreateNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_createNATNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxfindNATNetworkByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findNATNetworkByName"`

	This        string `xml:"_this,omitempty"`
	NetworkName string `xml:"networkName,omitempty"`
}

type IVirtualBoxfindNATNetworkByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_findNATNetworkByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVirtualBoxremoveNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeNATNetwork"`

	This    string `xml:"_this,omitempty"`
	Network string `xml:"network,omitempty"`
}

type IVirtualBoxremoveNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_removeNATNetworkResponse"`
}

type IVirtualBoxcheckFirmwarePresent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_checkFirmwarePresent"`

	This         string        `xml:"_this,omitempty"`
	FirmwareType *FirmwareType `xml:"firmwareType,omitempty"`
	Version      string        `xml:"version,omitempty"`
}

type IVirtualBoxcheckFirmwarePresentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualBox_checkFirmwarePresentResponse"`

	Url       string `xml:"url,omitempty"`
	File      string `xml:"file,omitempty"`
	Returnval bool   `xml:"returnval,omitempty"`
}

type IVFSExplorergetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getPath"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorergetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorergetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getType"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorergetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_getTypeResponse"`

	Returnval *VFSType `xml:"returnval,omitempty"`
}

type IVFSExplorerupdate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_update"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorerupdateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_updateResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorercd struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cd"`

	This string `xml:"_this,omitempty"`
	Dir  string `xml:"dir,omitempty"`
}

type IVFSExplorercdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorercdUp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdUp"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorercdUpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_cdUpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVFSExplorerentryList struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_entryList"`

	This string `xml:"_this,omitempty"`
}

type IVFSExplorerentryListResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_entryListResponse"`

	Names []string `xml:"names,omitempty"`
	Types []uint32 `xml:"types,omitempty"`
	Sizes []int64  `xml:"sizes,omitempty"`
	Modes []uint32 `xml:"modes,omitempty"`
}

type IVFSExplorerexists struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_exists"`

	This  string   `xml:"_this,omitempty"`
	Names []string `xml:"names,omitempty"`
}

type IVFSExplorerexistsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_existsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVFSExplorerremove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_remove"`

	This  string   `xml:"_this,omitempty"`
	Names []string `xml:"names,omitempty"`
}

type IVFSExplorerremoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVFSExplorer_removeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getPath"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetDisks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getDisks"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetDisksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getDisksResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IAppliancegetVirtualSystemDescriptions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getVirtualSystemDescriptions"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetVirtualSystemDescriptionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getVirtualSystemDescriptionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IAppliancegetMachines struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getMachines"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetMachinesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getMachinesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IApplianceread struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_read"`

	This string `xml:"_this,omitempty"`
	File string `xml:"file,omitempty"`
}

type IAppliancereadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_readResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IApplianceinterpret struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_interpret"`

	This string `xml:"_this,omitempty"`
}

type IApplianceinterpretResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_interpretResponse"`
}

type IApplianceimportMachines struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_importMachines"`

	This    string           `xml:"_this,omitempty"`
	Options []*ImportOptions `xml:"options,omitempty"`
}

type IApplianceimportMachinesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_importMachinesResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancecreateVFSExplorer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_createVFSExplorer"`

	This string `xml:"_this,omitempty"`
	URI  string `xml:"URI,omitempty"`
}

type IAppliancecreateVFSExplorerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_createVFSExplorerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancewrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_write"`

	This    string           `xml:"_this,omitempty"`
	Format  string           `xml:"format,omitempty"`
	Options []*ExportOptions `xml:"options,omitempty"`
	Path    string           `xml:"path,omitempty"`
}

type IAppliancewriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_writeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IAppliancegetWarnings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getWarnings"`

	This string `xml:"_this,omitempty"`
}

type IAppliancegetWarningsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAppliance_getWarningsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptiongetCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getCount"`

	This string `xml:"_this,omitempty"`
}

type IVirtualSystemDescriptiongetCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptiongetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionResponse"`

	Types             []*VirtualSystemDescriptionType `xml:"types,omitempty"`
	Refs              []string                        `xml:"refs,omitempty"`
	OVFValues         []string                        `xml:"OVFValues,omitempty"`
	VBoxValues        []string                        `xml:"VBoxValues,omitempty"`
	ExtraConfigValues []string                        `xml:"extraConfigValues,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionByType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionByType"`

	This  string                        `xml:"_this,omitempty"`
	Type_ *VirtualSystemDescriptionType `xml:"type,omitempty"`
}

type IVirtualSystemDescriptiongetDescriptionByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getDescriptionByTypeResponse"`

	Types             []*VirtualSystemDescriptionType `xml:"types,omitempty"`
	Refs              []string                        `xml:"refs,omitempty"`
	OVFValues         []string                        `xml:"OVFValues,omitempty"`
	VBoxValues        []string                        `xml:"VBoxValues,omitempty"`
	ExtraConfigValues []string                        `xml:"extraConfigValues,omitempty"`
}

type IVirtualSystemDescriptiongetValuesByType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getValuesByType"`

	This  string                             `xml:"_this,omitempty"`
	Type_ *VirtualSystemDescriptionType      `xml:"type,omitempty"`
	Which *VirtualSystemDescriptionValueType `xml:"which,omitempty"`
}

type IVirtualSystemDescriptiongetValuesByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_getValuesByTypeResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVirtualSystemDescriptionsetFinalValues struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_setFinalValues"`

	This              string   `xml:"_this,omitempty"`
	Enabled           []bool   `xml:"enabled,omitempty"`
	VBoxValues        []string `xml:"VBoxValues,omitempty"`
	ExtraConfigValues []string `xml:"extraConfigValues,omitempty"`
}

type IVirtualSystemDescriptionsetFinalValuesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_setFinalValuesResponse"`
}

type IVirtualSystemDescriptionaddDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_addDescription"`

	This             string                        `xml:"_this,omitempty"`
	Type_            *VirtualSystemDescriptionType `xml:"type,omitempty"`
	VBoxValue        string                        `xml:"VBoxValue,omitempty"`
	ExtraConfigValue string                        `xml:"extraConfigValue,omitempty"`
}

type IVirtualSystemDescriptionaddDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVirtualSystemDescription_addDescriptionResponse"`
}

type IBIOSSettingsgetLogoFadeIn struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeIn"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoFadeInResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeInResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoFadeIn struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeIn"`

	This       string `xml:"_this,omitempty"`
	LogoFadeIn bool   `xml:"logoFadeIn,omitempty"`
}

type IBIOSSettingssetLogoFadeInResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeInResponse"`
}

type IBIOSSettingsgetLogoFadeOut struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeOut"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoFadeOutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoFadeOutResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoFadeOut struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeOut"`

	This        string `xml:"_this,omitempty"`
	LogoFadeOut bool   `xml:"logoFadeOut,omitempty"`
}

type IBIOSSettingssetLogoFadeOutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoFadeOutResponse"`
}

type IBIOSSettingsgetLogoDisplayTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoDisplayTime"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoDisplayTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoDisplayTimeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoDisplayTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoDisplayTime"`

	This            string `xml:"_this,omitempty"`
	LogoDisplayTime uint32 `xml:"logoDisplayTime,omitempty"`
}

type IBIOSSettingssetLogoDisplayTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoDisplayTimeResponse"`
}

type IBIOSSettingsgetLogoImagePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoImagePath"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetLogoImagePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getLogoImagePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IBIOSSettingssetLogoImagePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoImagePath"`

	This          string `xml:"_this,omitempty"`
	LogoImagePath string `xml:"logoImagePath,omitempty"`
}

type IBIOSSettingssetLogoImagePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setLogoImagePathResponse"`
}

type IBIOSSettingsgetBootMenuMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getBootMenuMode"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetBootMenuModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getBootMenuModeResponse"`

	Returnval *BIOSBootMenuMode `xml:"returnval,omitempty"`
}

type IBIOSSettingssetBootMenuMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setBootMenuMode"`

	This         string            `xml:"_this,omitempty"`
	BootMenuMode *BIOSBootMenuMode `xml:"bootMenuMode,omitempty"`
}

type IBIOSSettingssetBootMenuModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setBootMenuModeResponse"`
}

type IBIOSSettingsgetACPIEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getACPIEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetACPIEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getACPIEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetACPIEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setACPIEnabled"`

	This        string `xml:"_this,omitempty"`
	ACPIEnabled bool   `xml:"ACPIEnabled,omitempty"`
}

type IBIOSSettingssetACPIEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setACPIEnabledResponse"`
}

type IBIOSSettingsgetIOAPICEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getIOAPICEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetIOAPICEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getIOAPICEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetIOAPICEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setIOAPICEnabled"`

	This          string `xml:"_this,omitempty"`
	IOAPICEnabled bool   `xml:"IOAPICEnabled,omitempty"`
}

type IBIOSSettingssetIOAPICEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setIOAPICEnabledResponse"`
}

type IBIOSSettingsgetTimeOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getTimeOffset"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetTimeOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getTimeOffsetResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IBIOSSettingssetTimeOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setTimeOffset"`

	This       string `xml:"_this,omitempty"`
	TimeOffset int64  `xml:"timeOffset,omitempty"`
}

type IBIOSSettingssetTimeOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setTimeOffsetResponse"`
}

type IBIOSSettingsgetPXEDebugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getPXEDebugEnabled"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetPXEDebugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getPXEDebugEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBIOSSettingssetPXEDebugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setPXEDebugEnabled"`

	This            string `xml:"_this,omitempty"`
	PXEDebugEnabled bool   `xml:"PXEDebugEnabled,omitempty"`
}

type IBIOSSettingssetPXEDebugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_setPXEDebugEnabledResponse"`
}

type IBIOSSettingsgetNonVolatileStorageFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getNonVolatileStorageFile"`

	This string `xml:"_this,omitempty"`
}

type IBIOSSettingsgetNonVolatileStorageFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBIOSSettings_getNonVolatileStorageFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPCIAddressgetBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getBus"`

	This string `xml:"_this,omitempty"`
}

type IPCIAddressgetBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getBusResponse"`

	Returnval int16 `xml:"returnval,omitempty"`
}

type IPCIAddresssetBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setBus"`

	This string `xml:"_this,omitempty"`
	Bus  int16  `xml:"bus,omitempty"`
}

type IPCIAddresssetBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setBusResponse"`
}

type IPCIAddressgetDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getDevice"`

	This string `xml:"_this,omitempty"`
}

type IPCIAddressgetDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getDeviceResponse"`

	Returnval int16 `xml:"returnval,omitempty"`
}

type IPCIAddresssetDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setDevice"`

	This   string `xml:"_this,omitempty"`
	Device int16  `xml:"device,omitempty"`
}

type IPCIAddresssetDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setDeviceResponse"`
}

type IPCIAddressgetDevFunction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getDevFunction"`

	This string `xml:"_this,omitempty"`
}

type IPCIAddressgetDevFunctionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_getDevFunctionResponse"`

	Returnval int16 `xml:"returnval,omitempty"`
}

type IPCIAddresssetDevFunction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setDevFunction"`

	This        string `xml:"_this,omitempty"`
	DevFunction int16  `xml:"devFunction,omitempty"`
}

type IPCIAddresssetDevFunctionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_setDevFunctionResponse"`
}

type IPCIAddressasLong struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_asLong"`

	This string `xml:"_this,omitempty"`
}

type IPCIAddressasLongResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_asLongResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IPCIAddressfromLong struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_fromLong"`

	This   string `xml:"_this,omitempty"`
	Number int32  `xml:"number,omitempty"`
}

type IPCIAddressfromLongResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIAddress_fromLongResponse"`
}

type IMachinegetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParent"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetIcon struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIcon"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIconResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIconResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetIcon struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIcon"`

	This string `xml:"_this,omitempty"`
	Icon string `xml:"icon,omitempty"`
}

type IMachinesetIconResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIconResponse"`
}

type IMachinegetAccessible struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessible"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccessibleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessibleResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetAccessError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessError"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccessErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccessErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getName"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinesetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setNameResponse"`
}

type IMachinegetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IMachinesetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDescriptionResponse"`
}

type IMachinegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getId"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGroups"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGroupsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinesetGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGroups"`

	This   string   `xml:"_this,omitempty"`
	Groups []string `xml:"groups,omitempty"`
}

type IMachinesetGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGroupsResponse"`
}

type IMachinegetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getOSTypeId"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getOSTypeIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setOSTypeId"`

	This     string `xml:"_this,omitempty"`
	OSTypeId string `xml:"OSTypeId,omitempty"`
}

type IMachinesetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setOSTypeIdResponse"`
}

type IMachinegetHardwareVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareVersion"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHardwareVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetHardwareVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareVersion"`

	This            string `xml:"_this,omitempty"`
	HardwareVersion string `xml:"hardwareVersion,omitempty"`
}

type IMachinesetHardwareVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareVersionResponse"`
}

type IMachinegetHardwareUUID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareUUID"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHardwareUUIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHardwareUUIDResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetHardwareUUID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareUUID"`

	This         string `xml:"_this,omitempty"`
	HardwareUUID string `xml:"hardwareUUID,omitempty"`
}

type IMachinesetHardwareUUIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHardwareUUIDResponse"`
}

type IMachinegetCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUCount"`

	This     string `xml:"_this,omitempty"`
	CPUCount uint32 `xml:"CPUCount,omitempty"`
}

type IMachinesetCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUCountResponse"`
}

type IMachinegetCPUHotPlugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUHotPlugEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCPUHotPlugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUHotPlugEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetCPUHotPlugEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUHotPlugEnabled"`

	This              string `xml:"_this,omitempty"`
	CPUHotPlugEnabled bool   `xml:"CPUHotPlugEnabled,omitempty"`
}

type IMachinesetCPUHotPlugEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUHotPlugEnabledResponse"`
}

type IMachinegetCPUExecutionCap struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUExecutionCap"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCPUExecutionCapResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUExecutionCapResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetCPUExecutionCap struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUExecutionCap"`

	This            string `xml:"_this,omitempty"`
	CPUExecutionCap uint32 `xml:"CPUExecutionCap,omitempty"`
}

type IMachinesetCPUExecutionCapResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUExecutionCapResponse"`
}

type IMachinegetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemorySize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemorySizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemorySize"`

	This       string `xml:"_this,omitempty"`
	MemorySize uint32 `xml:"memorySize,omitempty"`
}

type IMachinesetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemorySizeResponse"`
}

type IMachinegetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemoryBalloonSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMemoryBalloonSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemoryBalloonSize"`

	This              string `xml:"_this,omitempty"`
	MemoryBalloonSize uint32 `xml:"memoryBalloonSize,omitempty"`
}

type IMachinesetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMemoryBalloonSizeResponse"`
}

type IMachinegetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPageFusionEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPageFusionEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetPageFusionEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPageFusionEnabled"`

	This              string `xml:"_this,omitempty"`
	PageFusionEnabled bool   `xml:"pageFusionEnabled,omitempty"`
}

type IMachinesetPageFusionEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPageFusionEnabledResponse"`
}

type IMachinegetGraphicsControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGraphicsControllerType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetGraphicsControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGraphicsControllerTypeResponse"`

	Returnval *GraphicsControllerType `xml:"returnval,omitempty"`
}

type IMachinesetGraphicsControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGraphicsControllerType"`

	This                   string                  `xml:"_this,omitempty"`
	GraphicsControllerType *GraphicsControllerType `xml:"graphicsControllerType,omitempty"`
}

type IMachinesetGraphicsControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGraphicsControllerTypeResponse"`
}

type IMachinegetVRAMSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRAMSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVRAMSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRAMSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVRAMSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVRAMSize"`

	This     string `xml:"_this,omitempty"`
	VRAMSize uint32 `xml:"VRAMSize,omitempty"`
}

type IMachinesetVRAMSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVRAMSizeResponse"`
}

type IMachinegetAccelerate3DEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate3DEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccelerate3DEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate3DEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAccelerate3DEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate3DEnabled"`

	This                string `xml:"_this,omitempty"`
	Accelerate3DEnabled bool   `xml:"accelerate3DEnabled,omitempty"`
}

type IMachinesetAccelerate3DEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate3DEnabledResponse"`
}

type IMachinegetAccelerate2DVideoEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate2DVideoEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAccelerate2DVideoEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAccelerate2DVideoEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAccelerate2DVideoEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate2DVideoEnabled"`

	This                     string `xml:"_this,omitempty"`
	Accelerate2DVideoEnabled bool   `xml:"accelerate2DVideoEnabled,omitempty"`
}

type IMachinesetAccelerate2DVideoEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAccelerate2DVideoEnabledResponse"`
}

type IMachinegetMonitorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMonitorCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMonitorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMonitorCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetMonitorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMonitorCount"`

	This         string `xml:"_this,omitempty"`
	MonitorCount uint32 `xml:"monitorCount,omitempty"`
}

type IMachinesetMonitorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setMonitorCountResponse"`
}

type IMachinegetVideoCaptureEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureEnabled"`

	This                string `xml:"_this,omitempty"`
	VideoCaptureEnabled bool   `xml:"videoCaptureEnabled,omitempty"`
}

type IMachinesetVideoCaptureEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureEnabledResponse"`
}

type IMachinegetVideoCaptureScreens struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureScreens"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureScreensResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureScreensResponse"`

	Returnval []bool `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureScreens struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureScreens"`

	This                string `xml:"_this,omitempty"`
	VideoCaptureScreens []bool `xml:"videoCaptureScreens,omitempty"`
}

type IMachinesetVideoCaptureScreensResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureScreensResponse"`
}

type IMachinegetVideoCaptureFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureFile"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureFile"`

	This             string `xml:"_this,omitempty"`
	VideoCaptureFile string `xml:"videoCaptureFile,omitempty"`
}

type IMachinesetVideoCaptureFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureFileResponse"`
}

type IMachinegetVideoCaptureWidth struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureWidth"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureWidthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureWidthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureWidth struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureWidth"`

	This              string `xml:"_this,omitempty"`
	VideoCaptureWidth uint32 `xml:"videoCaptureWidth,omitempty"`
}

type IMachinesetVideoCaptureWidthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureWidthResponse"`
}

type IMachinegetVideoCaptureHeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureHeight"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureHeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureHeightResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureHeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureHeight"`

	This               string `xml:"_this,omitempty"`
	VideoCaptureHeight uint32 `xml:"videoCaptureHeight,omitempty"`
}

type IMachinesetVideoCaptureHeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureHeightResponse"`
}

type IMachinegetVideoCaptureRate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureRate"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureRateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureRateResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureRate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureRate"`

	This             string `xml:"_this,omitempty"`
	VideoCaptureRate uint32 `xml:"videoCaptureRate,omitempty"`
}

type IMachinesetVideoCaptureRateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureRateResponse"`
}

type IMachinegetVideoCaptureFPS struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureFPS"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVideoCaptureFPSResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVideoCaptureFPSResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetVideoCaptureFPS struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureFPS"`

	This            string `xml:"_this,omitempty"`
	VideoCaptureFPS uint32 `xml:"videoCaptureFPS,omitempty"`
}

type IMachinesetVideoCaptureFPSResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setVideoCaptureFPSResponse"`
}

type IMachinegetBIOSSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBIOSSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetBIOSSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBIOSSettingsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetFirmwareType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFirmwareType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFirmwareTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFirmwareTypeResponse"`

	Returnval *FirmwareType `xml:"returnval,omitempty"`
}

type IMachinesetFirmwareType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFirmwareType"`

	This         string        `xml:"_this,omitempty"`
	FirmwareType *FirmwareType `xml:"firmwareType,omitempty"`
}

type IMachinesetFirmwareTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFirmwareTypeResponse"`
}

type IMachinegetPointingHIDType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPointingHIDType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetPointingHIDTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPointingHIDTypeResponse"`

	Returnval *PointingHIDType `xml:"returnval,omitempty"`
}

type IMachinesetPointingHIDType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPointingHIDType"`

	This            string           `xml:"_this,omitempty"`
	PointingHIDType *PointingHIDType `xml:"pointingHIDType,omitempty"`
}

type IMachinesetPointingHIDTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setPointingHIDTypeResponse"`
}

type IMachinegetKeyboardHIDType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getKeyboardHIDType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetKeyboardHIDTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getKeyboardHIDTypeResponse"`

	Returnval *KeyboardHIDType `xml:"returnval,omitempty"`
}

type IMachinesetKeyboardHIDType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setKeyboardHIDType"`

	This            string           `xml:"_this,omitempty"`
	KeyboardHIDType *KeyboardHIDType `xml:"keyboardHIDType,omitempty"`
}

type IMachinesetKeyboardHIDTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setKeyboardHIDTypeResponse"`
}

type IMachinegetHPETEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHPETEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetHPETEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHPETEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetHPETEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHPETEnabled"`

	This        string `xml:"_this,omitempty"`
	HPETEnabled bool   `xml:"HPETEnabled,omitempty"`
}

type IMachinesetHPETEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHPETEnabledResponse"`
}

type IMachinegetChipsetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getChipsetType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetChipsetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getChipsetTypeResponse"`

	Returnval *ChipsetType `xml:"returnval,omitempty"`
}

type IMachinesetChipsetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setChipsetType"`

	This        string       `xml:"_this,omitempty"`
	ChipsetType *ChipsetType `xml:"chipsetType,omitempty"`
}

type IMachinesetChipsetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setChipsetTypeResponse"`
}

type IMachinegetSnapshotFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotFolder"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSnapshotFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetSnapshotFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSnapshotFolder"`

	This           string `xml:"_this,omitempty"`
	SnapshotFolder string `xml:"snapshotFolder,omitempty"`
}

type IMachinesetSnapshotFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSnapshotFolderResponse"`
}

type IMachinegetVRDEServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRDEServer"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetVRDEServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getVRDEServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetEmulatedUSBCardReaderEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getEmulatedUSBCardReaderEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetEmulatedUSBCardReaderEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getEmulatedUSBCardReaderEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetEmulatedUSBCardReaderEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setEmulatedUSBCardReaderEnabled"`

	This                         string `xml:"_this,omitempty"`
	EmulatedUSBCardReaderEnabled bool   `xml:"emulatedUSBCardReaderEnabled,omitempty"`
}

type IMachinesetEmulatedUSBCardReaderEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setEmulatedUSBCardReaderEnabledResponse"`
}

type IMachinegetMediumAttachments struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachments"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetMediumAttachmentsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsResponse"`

	Returnval []*IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachinegetUSBControllers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllers"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetUSBControllersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinegetUSBDeviceFilters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBDeviceFilters"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetUSBDeviceFiltersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBDeviceFiltersResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetAudioAdapter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAudioAdapter"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAudioAdapterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAudioAdapterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllers"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStorageControllersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinegetSettingsFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsFilePath"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSettingsFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSettingsModified struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsModified"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSettingsModifiedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSettingsModifiedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetSessionState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionState"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionStateResponse"`

	Returnval *SessionState `xml:"returnval,omitempty"`
}

type IMachinegetSessionType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionTypeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSessionPID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionPID"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSessionPIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSessionPIDResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getState"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type IMachinegetLastStateChange struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLastStateChange"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetLastStateChangeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLastStateChangeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMachinegetStateFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateFilePath"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetStateFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStateFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetLogFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLogFolder"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetLogFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getLogFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetCurrentSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentSnapshot"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCurrentSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetSnapshotCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotCount"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSnapshotCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSnapshotCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinegetCurrentStateModified struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentStateModified"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetCurrentStateModifiedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCurrentStateModifiedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinegetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IMachinegetClipboardMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getClipboardMode"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetClipboardModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getClipboardModeResponse"`

	Returnval *ClipboardMode `xml:"returnval,omitempty"`
}

type IMachinesetClipboardMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setClipboardMode"`

	This          string         `xml:"_this,omitempty"`
	ClipboardMode *ClipboardMode `xml:"clipboardMode,omitempty"`
}

type IMachinesetClipboardModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setClipboardModeResponse"`
}

type IMachinegetDragAndDropMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDragAndDropMode"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetDragAndDropModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDragAndDropModeResponse"`

	Returnval *DragAndDropMode `xml:"returnval,omitempty"`
}

type IMachinesetDragAndDropMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDragAndDropMode"`

	This            string           `xml:"_this,omitempty"`
	DragAndDropMode *DragAndDropMode `xml:"dragAndDropMode,omitempty"`
}

type IMachinesetDragAndDropModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDragAndDropModeResponse"`
}

type IMachinegetGuestPropertyNotificationPatterns struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyNotificationPatterns"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetGuestPropertyNotificationPatternsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyNotificationPatternsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetGuestPropertyNotificationPatterns struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyNotificationPatterns"`

	This                              string `xml:"_this,omitempty"`
	GuestPropertyNotificationPatterns string `xml:"guestPropertyNotificationPatterns,omitempty"`
}

type IMachinesetGuestPropertyNotificationPatternsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyNotificationPatternsResponse"`
}

type IMachinegetTeleporterEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterEnabled"`

	This              string `xml:"_this,omitempty"`
	TeleporterEnabled bool   `xml:"teleporterEnabled,omitempty"`
}

type IMachinesetTeleporterEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterEnabledResponse"`
}

type IMachinegetTeleporterPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPort"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPortResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPort"`

	This           string `xml:"_this,omitempty"`
	TeleporterPort uint32 `xml:"teleporterPort,omitempty"`
}

type IMachinesetTeleporterPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPortResponse"`
}

type IMachinegetTeleporterAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterAddress"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterAddress"`

	This              string `xml:"_this,omitempty"`
	TeleporterAddress string `xml:"teleporterAddress,omitempty"`
}

type IMachinesetTeleporterAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterAddressResponse"`
}

type IMachinegetTeleporterPassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPassword"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTeleporterPasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTeleporterPasswordResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetTeleporterPassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPassword"`

	This               string `xml:"_this,omitempty"`
	TeleporterPassword string `xml:"teleporterPassword,omitempty"`
}

type IMachinesetTeleporterPasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTeleporterPasswordResponse"`
}

type IMachinegetFaultToleranceState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceState"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFaultToleranceStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceStateResponse"`

	Returnval *FaultToleranceState `xml:"returnval,omitempty"`
}

type IMachinesetFaultToleranceState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceState"`

	This                string               `xml:"_this,omitempty"`
	FaultToleranceState *FaultToleranceState `xml:"faultToleranceState,omitempty"`
}

type IMachinesetFaultToleranceStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceStateResponse"`
}

type IMachinegetFaultTolerancePort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultTolerancePort"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFaultTolerancePortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultTolerancePortResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetFaultTolerancePort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultTolerancePort"`

	This               string `xml:"_this,omitempty"`
	FaultTolerancePort uint32 `xml:"faultTolerancePort,omitempty"`
}

type IMachinesetFaultTolerancePortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultTolerancePortResponse"`
}

type IMachinegetFaultToleranceAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceAddress"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFaultToleranceAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetFaultToleranceAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceAddress"`

	This                  string `xml:"_this,omitempty"`
	FaultToleranceAddress string `xml:"faultToleranceAddress,omitempty"`
}

type IMachinesetFaultToleranceAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceAddressResponse"`
}

type IMachinegetFaultTolerancePassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultTolerancePassword"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFaultTolerancePasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultTolerancePasswordResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetFaultTolerancePassword struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultTolerancePassword"`

	This                   string `xml:"_this,omitempty"`
	FaultTolerancePassword string `xml:"faultTolerancePassword,omitempty"`
}

type IMachinesetFaultTolerancePasswordResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultTolerancePasswordResponse"`
}

type IMachinegetFaultToleranceSyncInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceSyncInterval"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetFaultToleranceSyncIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getFaultToleranceSyncIntervalResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetFaultToleranceSyncInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceSyncInterval"`

	This                       string `xml:"_this,omitempty"`
	FaultToleranceSyncInterval uint32 `xml:"faultToleranceSyncInterval,omitempty"`
}

type IMachinesetFaultToleranceSyncIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setFaultToleranceSyncIntervalResponse"`
}

type IMachinegetRTCUseUTC struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getRTCUseUTC"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetRTCUseUTCResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getRTCUseUTCResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetRTCUseUTC struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setRTCUseUTC"`

	This      string `xml:"_this,omitempty"`
	RTCUseUTC bool   `xml:"RTCUseUTC,omitempty"`
}

type IMachinesetRTCUseUTCResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setRTCUseUTCResponse"`
}

type IMachinegetIOCacheEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIOCacheEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIOCacheEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIOCacheEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetIOCacheEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIOCacheEnabled"`

	This           string `xml:"_this,omitempty"`
	IOCacheEnabled bool   `xml:"IOCacheEnabled,omitempty"`
}

type IMachinesetIOCacheEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIOCacheEnabledResponse"`
}

type IMachinegetIOCacheSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIOCacheSize"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetIOCacheSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getIOCacheSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetIOCacheSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIOCacheSize"`

	This        string `xml:"_this,omitempty"`
	IOCacheSize uint32 `xml:"IOCacheSize,omitempty"`
}

type IMachinesetIOCacheSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setIOCacheSizeResponse"`
}

type IMachinegetPCIDeviceAssignments struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPCIDeviceAssignments"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetPCIDeviceAssignmentsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getPCIDeviceAssignmentsResponse"`

	Returnval []*IPCIDeviceAttachment `xml:"returnval,omitempty"`
}

type IMachinegetBandwidthControl struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBandwidthControl"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetBandwidthControlResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBandwidthControlResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetTracingEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTracingEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTracingEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTracingEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetTracingEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTracingEnabled"`

	This           string `xml:"_this,omitempty"`
	TracingEnabled bool   `xml:"tracingEnabled,omitempty"`
}

type IMachinesetTracingEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTracingEnabledResponse"`
}

type IMachinegetTracingConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTracingConfig"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetTracingConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getTracingConfigResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetTracingConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTracingConfig"`

	This          string `xml:"_this,omitempty"`
	TracingConfig string `xml:"tracingConfig,omitempty"`
}

type IMachinesetTracingConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setTracingConfigResponse"`
}

type IMachinegetAllowTracingToAccessVM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAllowTracingToAccessVM"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAllowTracingToAccessVMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAllowTracingToAccessVMResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAllowTracingToAccessVM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAllowTracingToAccessVM"`

	This                   string `xml:"_this,omitempty"`
	AllowTracingToAccessVM bool   `xml:"allowTracingToAccessVM,omitempty"`
}

type IMachinesetAllowTracingToAccessVMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAllowTracingToAccessVMResponse"`
}

type IMachinegetAutostartEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostartEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAutostartEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostartEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetAutostartEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostartEnabled"`

	This             string `xml:"_this,omitempty"`
	AutostartEnabled bool   `xml:"autostartEnabled,omitempty"`
}

type IMachinesetAutostartEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostartEnabledResponse"`
}

type IMachinegetAutostartDelay struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostartDelay"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAutostartDelayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostartDelayResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinesetAutostartDelay struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostartDelay"`

	This           string `xml:"_this,omitempty"`
	AutostartDelay uint32 `xml:"autostartDelay,omitempty"`
}

type IMachinesetAutostartDelayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostartDelayResponse"`
}

type IMachinegetAutostopType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostopType"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetAutostopTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getAutostopTypeResponse"`

	Returnval *AutostopType `xml:"returnval,omitempty"`
}

type IMachinesetAutostopType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostopType"`

	This         string        `xml:"_this,omitempty"`
	AutostopType *AutostopType `xml:"autostopType,omitempty"`
}

type IMachinesetAutostopTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutostopTypeResponse"`
}

type IMachinegetDefaultFrontend struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDefaultFrontend"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetDefaultFrontendResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getDefaultFrontendResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetDefaultFrontend struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDefaultFrontend"`

	This            string `xml:"_this,omitempty"`
	DefaultFrontend string `xml:"defaultFrontend,omitempty"`
}

type IMachinesetDefaultFrontendResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setDefaultFrontendResponse"`
}

type IMachinegetUSBProxyAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBProxyAvailable"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetUSBProxyAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBProxyAvailableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinelockMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_lockMachine"`

	This     string    `xml:"_this,omitempty"`
	Session  string    `xml:"session,omitempty"`
	LockType *LockType `xml:"lockType,omitempty"`
}

type IMachinelockMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_lockMachineResponse"`
}

type IMachinelaunchVMProcess struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_launchVMProcess"`

	This        string `xml:"_this,omitempty"`
	Session     string `xml:"session,omitempty"`
	Type_       string `xml:"type,omitempty"`
	Environment string `xml:"environment,omitempty"`
}

type IMachinelaunchVMProcessResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_launchVMProcessResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetBootOrder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBootOrder"`

	This     string      `xml:"_this,omitempty"`
	Position uint32      `xml:"position,omitempty"`
	Device   *DeviceType `xml:"device,omitempty"`
}

type IMachinesetBootOrderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBootOrderResponse"`
}

type IMachinegetBootOrder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBootOrder"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IMachinegetBootOrderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getBootOrderResponse"`

	Returnval *DeviceType `xml:"returnval,omitempty"`
}

type IMachineattachDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDevice"`

	This           string      `xml:"_this,omitempty"`
	Name           string      `xml:"name,omitempty"`
	ControllerPort int32       `xml:"controllerPort,omitempty"`
	Device         int32       `xml:"device,omitempty"`
	Type_          *DeviceType `xml:"type,omitempty"`
	Medium         string      `xml:"medium,omitempty"`
}

type IMachineattachDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDeviceResponse"`
}

type IMachineattachDeviceWithoutMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDeviceWithoutMedium"`

	This           string      `xml:"_this,omitempty"`
	Name           string      `xml:"name,omitempty"`
	ControllerPort int32       `xml:"controllerPort,omitempty"`
	Device         int32       `xml:"device,omitempty"`
	Type_          *DeviceType `xml:"type,omitempty"`
}

type IMachineattachDeviceWithoutMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachDeviceWithoutMediumResponse"`
}

type IMachinedetachDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinedetachDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachDeviceResponse"`
}

type IMachinepassthroughDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_passthroughDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Passthrough    bool   `xml:"passthrough,omitempty"`
}

type IMachinepassthroughDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_passthroughDeviceResponse"`
}

type IMachinetemporaryEjectDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_temporaryEjectDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	TemporaryEject bool   `xml:"temporaryEject,omitempty"`
}

type IMachinetemporaryEjectDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_temporaryEjectDeviceResponse"`
}

type IMachinenonRotationalDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_nonRotationalDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	NonRotational  bool   `xml:"nonRotational,omitempty"`
}

type IMachinenonRotationalDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_nonRotationalDeviceResponse"`
}

type IMachinesetAutoDiscardForDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutoDiscardForDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Discard        bool   `xml:"discard,omitempty"`
}

type IMachinesetAutoDiscardForDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setAutoDiscardForDeviceResponse"`
}

type IMachinesetHotPluggableForDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHotPluggableForDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	HotPluggable   bool   `xml:"hotPluggable,omitempty"`
}

type IMachinesetHotPluggableForDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHotPluggableForDeviceResponse"`
}

type IMachinesetBandwidthGroupForDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBandwidthGroupForDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	BandwidthGroup string `xml:"bandwidthGroup,omitempty"`
}

type IMachinesetBandwidthGroupForDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setBandwidthGroupForDeviceResponse"`
}

type IMachinesetNoBandwidthGroupForDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setNoBandwidthGroupForDevice"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinesetNoBandwidthGroupForDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setNoBandwidthGroupForDeviceResponse"`
}

type IMachineunmountMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_unmountMedium"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Force          bool   `xml:"force,omitempty"`
}

type IMachineunmountMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_unmountMediumResponse"`
}

type IMachinemountMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_mountMedium"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
	Medium         string `xml:"medium,omitempty"`
	Force          bool   `xml:"force,omitempty"`
}

type IMachinemountMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_mountMediumResponse"`
}

type IMachinegetMedium struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMedium"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinegetMediumResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetMediumAttachmentsOfController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsOfController"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetMediumAttachmentsOfControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentsOfControllerResponse"`

	Returnval []*IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachinegetMediumAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachment"`

	This           string `xml:"_this,omitempty"`
	Name           string `xml:"name,omitempty"`
	ControllerPort int32  `xml:"controllerPort,omitempty"`
	Device         int32  `xml:"device,omitempty"`
}

type IMachinegetMediumAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getMediumAttachmentResponse"`

	Returnval *IMediumAttachment `xml:"returnval,omitempty"`
}

type IMachineattachHostPCIDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachHostPCIDevice"`

	This                string `xml:"_this,omitempty"`
	HostAddress         int32  `xml:"hostAddress,omitempty"`
	DesiredGuestAddress int32  `xml:"desiredGuestAddress,omitempty"`
	TryToUnbind         bool   `xml:"tryToUnbind,omitempty"`
}

type IMachineattachHostPCIDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_attachHostPCIDeviceResponse"`
}

type IMachinedetachHostPCIDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachHostPCIDevice"`

	This        string `xml:"_this,omitempty"`
	HostAddress int32  `xml:"hostAddress,omitempty"`
}

type IMachinedetachHostPCIDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_detachHostPCIDeviceResponse"`
}

type IMachinegetNetworkAdapter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNetworkAdapter"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetNetworkAdapterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getNetworkAdapterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineaddStorageController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addStorageController"`

	This           string      `xml:"_this,omitempty"`
	Name           string      `xml:"name,omitempty"`
	ConnectionType *StorageBus `xml:"connectionType,omitempty"`
}

type IMachineaddStorageControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addStorageControllerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllerByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetStorageControllerByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetStorageControllerByInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByInstance"`

	This     string `xml:"_this,omitempty"`
	Instance uint32 `xml:"instance,omitempty"`
}

type IMachinegetStorageControllerByInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getStorageControllerByInstanceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineremoveStorageController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeStorageController"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachineremoveStorageControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeStorageControllerResponse"`
}

type IMachinesetStorageControllerBootable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setStorageControllerBootable"`

	This     string `xml:"_this,omitempty"`
	Name     string `xml:"name,omitempty"`
	Bootable bool   `xml:"bootable,omitempty"`
}

type IMachinesetStorageControllerBootableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setStorageControllerBootableResponse"`
}

type IMachineaddUSBController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addUSBController"`

	This  string             `xml:"_this,omitempty"`
	Name  string             `xml:"name,omitempty"`
	Type_ *USBControllerType `xml:"type,omitempty"`
}

type IMachineaddUSBControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_addUSBControllerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineremoveUSBController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeUSBController"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachineremoveUSBControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeUSBControllerResponse"`
}

type IMachinegetUSBControllerByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllerByName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetUSBControllerByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllerByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetUSBControllerCountByType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllerCountByType"`

	This  string             `xml:"_this,omitempty"`
	Type_ *USBControllerType `xml:"type,omitempty"`
}

type IMachinegetUSBControllerCountByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getUSBControllerCountByTypeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachinegetSerialPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSerialPort"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetSerialPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getSerialPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetParallelPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParallelPort"`

	This string `xml:"_this,omitempty"`
	Slot uint32 `xml:"slot,omitempty"`
}

type IMachinegetParallelPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getParallelPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetExtraDataKeys struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataKeys"`

	This string `xml:"_this,omitempty"`
}

type IMachinegetExtraDataKeysResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataKeysResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinegetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraData"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type IMachinegetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getExtraDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesetExtraData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setExtraData"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IMachinesetExtraDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setExtraDataResponse"`
}

type IMachinegetCPUProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUProperty"`

	This     string           `xml:"_this,omitempty"`
	Property *CPUPropertyType `xml:"property,omitempty"`
}

type IMachinegetCPUPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUPropertyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetCPUProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUProperty"`

	This     string           `xml:"_this,omitempty"`
	Property *CPUPropertyType `xml:"property,omitempty"`
	Value    bool             `xml:"value,omitempty"`
}

type IMachinesetCPUPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUPropertyResponse"`
}

type IMachinegetCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUIDLeaf"`

	This string `xml:"_this,omitempty"`
	Id   uint32 `xml:"id,omitempty"`
}

type IMachinegetCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUIDLeafResponse"`

	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IMachinesetCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUIDLeaf"`

	This   string `xml:"_this,omitempty"`
	Id     uint32 `xml:"id,omitempty"`
	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IMachinesetCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setCPUIDLeafResponse"`
}

type IMachineremoveCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeCPUIDLeaf"`

	This string `xml:"_this,omitempty"`
	Id   uint32 `xml:"id,omitempty"`
}

type IMachineremoveCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeCPUIDLeafResponse"`
}

type IMachineremoveAllCPUIDLeaves struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeAllCPUIDLeaves"`

	This string `xml:"_this,omitempty"`
}

type IMachineremoveAllCPUIDLeavesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeAllCPUIDLeavesResponse"`
}

type IMachinegetHWVirtExProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHWVirtExProperty"`

	This     string                `xml:"_this,omitempty"`
	Property *HWVirtExPropertyType `xml:"property,omitempty"`
}

type IMachinegetHWVirtExPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getHWVirtExPropertyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinesetHWVirtExProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHWVirtExProperty"`

	This     string                `xml:"_this,omitempty"`
	Property *HWVirtExPropertyType `xml:"property,omitempty"`
	Value    bool                  `xml:"value,omitempty"`
}

type IMachinesetHWVirtExPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setHWVirtExPropertyResponse"`
}

type IMachinesetSettingsFilePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSettingsFilePath"`

	This             string `xml:"_this,omitempty"`
	SettingsFilePath string `xml:"settingsFilePath,omitempty"`
}

type IMachinesetSettingsFilePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setSettingsFilePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinesaveSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_saveSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinesaveSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_saveSettingsResponse"`
}

type IMachinediscardSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_discardSettings"`

	This string `xml:"_this,omitempty"`
}

type IMachinediscardSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_discardSettingsResponse"`
}

type IMachineunregister struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_unregister"`

	This        string       `xml:"_this,omitempty"`
	CleanupMode *CleanupMode `xml:"cleanupMode,omitempty"`
}

type IMachineunregisterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_unregisterResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMachinedeleteConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteConfig"`

	This  string   `xml:"_this,omitempty"`
	Media []string `xml:"media,omitempty"`
}

type IMachinedeleteConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteConfigResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineexportTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_exportTo"`

	This      string `xml:"_this,omitempty"`
	Appliance string `xml:"appliance,omitempty"`
	Location  string `xml:"location,omitempty"`
}

type IMachineexportToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_exportToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinefindSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_findSnapshot"`

	This     string `xml:"_this,omitempty"`
	NameOrId string `xml:"nameOrId,omitempty"`
}

type IMachinefindSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_findSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinecreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_createSharedFolder"`

	This      string `xml:"_this,omitempty"`
	Name      string `xml:"name,omitempty"`
	HostPath  string `xml:"hostPath,omitempty"`
	Writable  bool   `xml:"writable,omitempty"`
	Automount bool   `xml:"automount,omitempty"`
}

type IMachinecreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_createSharedFolderResponse"`
}

type IMachineremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachineremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_removeSharedFolderResponse"`
}

type IMachinecanShowConsoleWindow struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_canShowConsoleWindow"`

	This string `xml:"_this,omitempty"`
}

type IMachinecanShowConsoleWindowResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_canShowConsoleWindowResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineshowConsoleWindow struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_showConsoleWindow"`

	This string `xml:"_this,omitempty"`
}

type IMachineshowConsoleWindowResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_showConsoleWindowResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMachinegetGuestProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestProperty"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinegetGuestPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyResponse"`

	Value     string `xml:"value,omitempty"`
	Timestamp int64  `xml:"timestamp,omitempty"`
	Flags     string `xml:"flags,omitempty"`
}

type IMachinegetGuestPropertyValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyValue"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
}

type IMachinegetGuestPropertyValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyValueResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinegetGuestPropertyTimestamp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyTimestamp"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
}

type IMachinegetGuestPropertyTimestampResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getGuestPropertyTimestampResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMachinesetGuestProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestProperty"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
	Value    string `xml:"value,omitempty"`
	Flags    string `xml:"flags,omitempty"`
}

type IMachinesetGuestPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyResponse"`
}

type IMachinesetGuestPropertyValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyValue"`

	This     string `xml:"_this,omitempty"`
	Property string `xml:"property,omitempty"`
	Value    string `xml:"value,omitempty"`
}

type IMachinesetGuestPropertyValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_setGuestPropertyValueResponse"`
}

type IMachinedeleteGuestProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteGuestProperty"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMachinedeleteGuestPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_deleteGuestPropertyResponse"`
}

type IMachineenumerateGuestProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_enumerateGuestProperties"`

	This     string `xml:"_this,omitempty"`
	Patterns string `xml:"patterns,omitempty"`
}

type IMachineenumerateGuestPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_enumerateGuestPropertiesResponse"`

	Names      []string `xml:"names,omitempty"`
	Values     []string `xml:"values,omitempty"`
	Timestamps []int64  `xml:"timestamps,omitempty"`
	Flags      []string `xml:"flags,omitempty"`
}

type IMachinequerySavedGuestScreenInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedGuestScreenInfo"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinequerySavedGuestScreenInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedGuestScreenInfoResponse"`

	OriginX uint32 `xml:"originX,omitempty"`
	OriginY uint32 `xml:"originY,omitempty"`
	Width   uint32 `xml:"width,omitempty"`
	Height  uint32 `xml:"height,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IMachinequerySavedThumbnailSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedThumbnailSize"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinequerySavedThumbnailSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedThumbnailSizeResponse"`

	Size   uint32 `xml:"size,omitempty"`
	Width  uint32 `xml:"width,omitempty"`
	Height uint32 `xml:"height,omitempty"`
}

type IMachinereadSavedThumbnailToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	BGR      bool   `xml:"BGR,omitempty"`
}

type IMachinereadSavedThumbnailToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailToArrayResponse"`

	Width     uint32 `xml:"width,omitempty"`
	Height    uint32 `xml:"height,omitempty"`
	Returnval string `xml:"returnval,omitempty"`
}

type IMachinereadSavedThumbnailPNGToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailPNGToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinereadSavedThumbnailPNGToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedThumbnailPNGToArrayResponse"`

	Width     uint32 `xml:"width,omitempty"`
	Height    uint32 `xml:"height,omitempty"`
	Returnval string `xml:"returnval,omitempty"`
}

type IMachinequerySavedScreenshotPNGSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedScreenshotPNGSize"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinequerySavedScreenshotPNGSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_querySavedScreenshotPNGSizeResponse"`

	Size   uint32 `xml:"size,omitempty"`
	Width  uint32 `xml:"width,omitempty"`
	Height uint32 `xml:"height,omitempty"`
}

type IMachinereadSavedScreenshotPNGToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedScreenshotPNGToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IMachinereadSavedScreenshotPNGToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readSavedScreenshotPNGToArrayResponse"`

	Width     uint32 `xml:"width,omitempty"`
	Height    uint32 `xml:"height,omitempty"`
	Returnval string `xml:"returnval,omitempty"`
}

type IMachinehotPlugCPU struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotPlugCPU"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinehotPlugCPUResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotPlugCPUResponse"`
}

type IMachinehotUnplugCPU struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotUnplugCPU"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinehotUnplugCPUResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_hotUnplugCPUResponse"`
}

type IMachinegetCPUStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUStatus"`

	This string `xml:"_this,omitempty"`
	Cpu  uint32 `xml:"cpu,omitempty"`
}

type IMachinegetCPUStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_getCPUStatusResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachinequeryLogFilename struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_queryLogFilename"`

	This string `xml:"_this,omitempty"`
	Idx  uint32 `xml:"idx,omitempty"`
}

type IMachinequeryLogFilenameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_queryLogFilenameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinereadLog struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readLog"`

	This   string `xml:"_this,omitempty"`
	Idx    uint32 `xml:"idx,omitempty"`
	Offset int64  `xml:"offset,omitempty"`
	Size   int64  `xml:"size,omitempty"`
}

type IMachinereadLogResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_readLogResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachinecloneTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_cloneTo"`

	This    string          `xml:"_this,omitempty"`
	Target  string          `xml:"target,omitempty"`
	Mode    *CloneMode      `xml:"mode,omitempty"`
	Options []*CloneOptions `xml:"options,omitempty"`
}

type IMachinecloneToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachine_cloneToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEmulatedUSBgetWebcams struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_getWebcams"`

	This string `xml:"_this,omitempty"`
}

type IEmulatedUSBgetWebcamsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_getWebcamsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IEmulatedUSBwebcamAttach struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_webcamAttach"`

	This     string `xml:"_this,omitempty"`
	Path     string `xml:"path,omitempty"`
	Settings string `xml:"settings,omitempty"`
}

type IEmulatedUSBwebcamAttachResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_webcamAttachResponse"`
}

type IEmulatedUSBwebcamDetach struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_webcamDetach"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IEmulatedUSBwebcamDetachResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEmulatedUSB_webcamDetachResponse"`
}

type IConsolegetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMachine"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getState"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type IConsolegetGuest struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuest"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetGuestResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetKeyboard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getKeyboard"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetKeyboardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getKeyboardResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetMouse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMouse"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetMouseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getMouseResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetDisplay struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDisplay"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetDisplayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDisplayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetDebugger struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDebugger"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetDebuggerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDebuggerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IConsolegetRemoteUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetRemoteUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getRemoteUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IConsolegetSharedFolders struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getSharedFolders"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetSharedFoldersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getSharedFoldersResponse"`

	Returnval []*ISharedFolder `xml:"returnval,omitempty"`
}

type IConsolegetVRDEServerInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getVRDEServerInfo"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetVRDEServerInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getVRDEServerInfoResponse"`

	Returnval *IVRDEServerInfo `xml:"returnval,omitempty"`
}

type IConsolegetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolegetAttachedPCIDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getAttachedPCIDevices"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetAttachedPCIDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getAttachedPCIDevicesResponse"`

	Returnval []*IPCIDeviceAttachment `xml:"returnval,omitempty"`
}

type IConsolegetUseHostClipboard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUseHostClipboard"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetUseHostClipboardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getUseHostClipboardResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IConsolesetUseHostClipboard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_setUseHostClipboard"`

	This             string `xml:"_this,omitempty"`
	UseHostClipboard bool   `xml:"useHostClipboard,omitempty"`
}

type IConsolesetUseHostClipboardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_setUseHostClipboardResponse"`
}

type IConsolegetEmulatedUSB struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getEmulatedUSB"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetEmulatedUSBResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getEmulatedUSBResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolepowerUp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUp"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerUpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolepowerUpPaused struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpPaused"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerUpPausedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerUpPausedResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolepowerDown struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerDown"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerDownResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerDownResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolereset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_reset"`

	This string `xml:"_this,omitempty"`
}

type IConsoleresetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resetResponse"`
}

type IConsolepause struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_pause"`

	This string `xml:"_this,omitempty"`
}

type IConsolepauseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_pauseResponse"`
}

type IConsoleresume struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resume"`

	This string `xml:"_this,omitempty"`
}

type IConsoleresumeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_resumeResponse"`
}

type IConsolepowerButton struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerButton"`

	This string `xml:"_this,omitempty"`
}

type IConsolepowerButtonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_powerButtonResponse"`
}

type IConsolesleepButton struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_sleepButton"`

	This string `xml:"_this,omitempty"`
}

type IConsolesleepButtonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_sleepButtonResponse"`
}

type IConsolegetPowerButtonHandled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getPowerButtonHandled"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetPowerButtonHandledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getPowerButtonHandledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IConsolegetGuestEnteredACPIMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestEnteredACPIMode"`

	This string `xml:"_this,omitempty"`
}

type IConsolegetGuestEnteredACPIModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getGuestEnteredACPIModeResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IConsolesaveState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_saveState"`

	This string `xml:"_this,omitempty"`
}

type IConsolesaveStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_saveStateResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoleadoptSavedState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_adoptSavedState"`

	This           string `xml:"_this,omitempty"`
	SavedStateFile string `xml:"savedStateFile,omitempty"`
}

type IConsoleadoptSavedStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_adoptSavedStateResponse"`
}

type IConsolediscardSavedState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_discardSavedState"`

	This        string `xml:"_this,omitempty"`
	FRemoveFile bool   `xml:"fRemoveFile,omitempty"`
}

type IConsolediscardSavedStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_discardSavedStateResponse"`
}

type IConsolegetDeviceActivity struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDeviceActivity"`

	This  string      `xml:"_this,omitempty"`
	Type_ *DeviceType `xml:"type,omitempty"`
}

type IConsolegetDeviceActivityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_getDeviceActivityResponse"`

	Returnval *DeviceActivity `xml:"returnval,omitempty"`
}

type IConsoleattachUSBDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_attachUSBDevice"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoleattachUSBDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_attachUSBDeviceResponse"`
}

type IConsoledetachUSBDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_detachUSBDevice"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoledetachUSBDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_detachUSBDeviceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolefindUSBDeviceByAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByAddress"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IConsolefindUSBDeviceByAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolefindUSBDeviceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsolefindUSBDeviceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_findUSBDeviceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolecreateSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_createSharedFolder"`

	This      string `xml:"_this,omitempty"`
	Name      string `xml:"name,omitempty"`
	HostPath  string `xml:"hostPath,omitempty"`
	Writable  bool   `xml:"writable,omitempty"`
	Automount bool   `xml:"automount,omitempty"`
}

type IConsolecreateSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_createSharedFolderResponse"`
}

type IConsoleremoveSharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_removeSharedFolder"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IConsoleremoveSharedFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_removeSharedFolderResponse"`
}

type IConsoletakeSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_takeSnapshot"`

	This        string `xml:"_this,omitempty"`
	Name        string `xml:"name,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IConsoletakeSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_takeSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoledeleteSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshot"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoledeleteSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoledeleteSnapshotAndAllChildren struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotAndAllChildren"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IConsoledeleteSnapshotAndAllChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotAndAllChildrenResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoledeleteSnapshotRange struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotRange"`

	This    string `xml:"_this,omitempty"`
	StartId string `xml:"startId,omitempty"`
	EndId   string `xml:"endId,omitempty"`
}

type IConsoledeleteSnapshotRangeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_deleteSnapshotRangeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsolerestoreSnapshot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_restoreSnapshot"`

	This     string `xml:"_this,omitempty"`
	Snapshot string `xml:"snapshot,omitempty"`
}

type IConsolerestoreSnapshotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_restoreSnapshotResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IConsoleteleport struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_teleport"`

	This        string `xml:"_this,omitempty"`
	Hostname    string `xml:"hostname,omitempty"`
	Tcpport     uint32 `xml:"tcpport,omitempty"`
	Password    string `xml:"password,omitempty"`
	MaxDowntime uint32 `xml:"maxDowntime,omitempty"`
}

type IConsoleteleportResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IConsole_teleportResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getName"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetShortName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getShortName"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetShortNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getShortNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getId"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetDHCPEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getDHCPEnabled"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetDHCPEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getDHCPEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPAddress"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetNetworkMask struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkMask"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetNetworkMaskResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getNetworkMaskResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6Supported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6Supported"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6SupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6SupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6Address struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6Address"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6AddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6AddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6NetworkMaskPrefixLength"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getIPV6NetworkMaskPrefixLengthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetHardwareAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getHardwareAddress"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetHardwareAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getHardwareAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetMediumType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getMediumType"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetMediumTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getMediumTypeResponse"`

	Returnval *HostNetworkInterfaceMediumType `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getStatusResponse"`

	Returnval *HostNetworkInterfaceStatus `xml:"returnval,omitempty"`
}

type IHostNetworkInterfacegetInterfaceType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getInterfaceType"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfacegetInterfaceTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_getInterfaceTypeResponse"`

	Returnval *HostNetworkInterfaceType `xml:"returnval,omitempty"`
}

type IHostNetworkInterfaceenableStaticIPConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIPConfig"`

	This        string `xml:"_this,omitempty"`
	IPAddress   string `xml:"IPAddress,omitempty"`
	NetworkMask string `xml:"networkMask,omitempty"`
}

type IHostNetworkInterfaceenableStaticIPConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIPConfigResponse"`
}

type IHostNetworkInterfaceenableStaticIPConfigV6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIPConfigV6"`

	This                        string `xml:"_this,omitempty"`
	IPV6Address                 string `xml:"IPV6Address,omitempty"`
	IPV6NetworkMaskPrefixLength uint32 `xml:"IPV6NetworkMaskPrefixLength,omitempty"`
}

type IHostNetworkInterfaceenableStaticIPConfigV6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableStaticIPConfigV6Response"`
}

type IHostNetworkInterfaceenableDynamicIPConfig struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableDynamicIPConfig"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfaceenableDynamicIPConfigResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_enableDynamicIPConfigResponse"`
}

type IHostNetworkInterfaceDHCPRediscover struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_DHCPRediscover"`

	This string `xml:"_this,omitempty"`
}

type IHostNetworkInterfaceDHCPRediscoverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostNetworkInterface_DHCPRediscoverResponse"`
}

type IHostVideoInputDevicegetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getName"`

	This string `xml:"_this,omitempty"`
}

type IHostVideoInputDevicegetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostVideoInputDevicegetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getPath"`

	This string `xml:"_this,omitempty"`
}

type IHostVideoInputDevicegetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostVideoInputDevicegetAlias struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getAlias"`

	This string `xml:"_this,omitempty"`
}

type IHostVideoInputDevicegetAliasResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostVideoInputDevice_getAliasResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetDVDDrives struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDVDDrives"`

	This string `xml:"_this,omitempty"`
}

type IHostgetDVDDrivesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDVDDrivesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetFloppyDrives struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getFloppyDrives"`

	This string `xml:"_this,omitempty"`
}

type IHostgetFloppyDrivesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getFloppyDrivesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetUSBDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDevices"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUSBDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetUSBDeviceFilters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDeviceFilters"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUSBDeviceFiltersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUSBDeviceFiltersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetNetworkInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNetworkInterfaces"`

	This string `xml:"_this,omitempty"`
}

type IHostgetNetworkInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNetworkInterfacesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetNameServers struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNameServers"`

	This string `xml:"_this,omitempty"`
}

type IHostgetNameServersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getNameServersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetDomainName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDomainName"`

	This string `xml:"_this,omitempty"`
}

type IHostgetDomainNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getDomainNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetSearchStrings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getSearchStrings"`

	This string `xml:"_this,omitempty"`
}

type IHostgetSearchStringsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getSearchStringsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetProcessorCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorOnlineCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorOnlineCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorCoreCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCoreCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorCoreCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCoreCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorOnlineCoreCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCoreCount"`

	This string `xml:"_this,omitempty"`
}

type IHostgetProcessorOnlineCoreCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorOnlineCoreCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetMemorySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemorySize"`

	This string `xml:"_this,omitempty"`
}

type IHostgetMemorySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemorySizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetMemoryAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemoryAvailable"`

	This string `xml:"_this,omitempty"`
}

type IHostgetMemoryAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getMemoryAvailableResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetOperatingSystem struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOperatingSystem"`

	This string `xml:"_this,omitempty"`
}

type IHostgetOperatingSystemResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOperatingSystemResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetOSVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOSVersion"`

	This string `xml:"_this,omitempty"`
}

type IHostgetOSVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getOSVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetUTCTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUTCTime"`

	This string `xml:"_this,omitempty"`
}

type IHostgetUTCTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getUTCTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IHostgetAcceleration3DAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getAcceleration3DAvailable"`

	This string `xml:"_this,omitempty"`
}

type IHostgetAcceleration3DAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getAcceleration3DAvailableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostgetVideoInputDevices struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getVideoInputDevices"`

	This string `xml:"_this,omitempty"`
}

type IHostgetVideoInputDevicesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getVideoInputDevicesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostgetProcessorSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorSpeed"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IHostgetProcessorSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorSpeedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IHostgetProcessorFeature struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorFeature"`

	This    string            `xml:"_this,omitempty"`
	Feature *ProcessorFeature `xml:"feature,omitempty"`
}

type IHostgetProcessorFeatureResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorFeatureResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostgetProcessorDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorDescription"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IHostgetProcessorDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgetProcessorCPUIDLeaf struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCPUIDLeaf"`

	This    string `xml:"_this,omitempty"`
	CpuId   uint32 `xml:"cpuId,omitempty"`
	Leaf    uint32 `xml:"leaf,omitempty"`
	SubLeaf uint32 `xml:"subLeaf,omitempty"`
}

type IHostgetProcessorCPUIDLeafResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_getProcessorCPUIDLeafResponse"`

	ValEax uint32 `xml:"valEax,omitempty"`
	ValEbx uint32 `xml:"valEbx,omitempty"`
	ValEcx uint32 `xml:"valEcx,omitempty"`
	ValEdx uint32 `xml:"valEdx,omitempty"`
}

type IHostcreateHostOnlyNetworkInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createHostOnlyNetworkInterface"`

	This string `xml:"_this,omitempty"`
}

type IHostcreateHostOnlyNetworkInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createHostOnlyNetworkInterfaceResponse"`

	HostInterface string `xml:"hostInterface,omitempty"`
	Returnval     string `xml:"returnval,omitempty"`
}

type IHostremoveHostOnlyNetworkInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeHostOnlyNetworkInterface"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostremoveHostOnlyNetworkInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeHostOnlyNetworkInterfaceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostcreateUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createUSBDeviceFilter"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostcreateUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_createUSBDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostinsertUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_insertUSBDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
	Filter   string `xml:"filter,omitempty"`
}

type IHostinsertUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_insertUSBDeviceFilterResponse"`
}

type IHostremoveUSBDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeUSBDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IHostremoveUSBDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_removeUSBDeviceFilterResponse"`
}

type IHostfindHostDVDDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostDVDDrive"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostDVDDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostDVDDriveResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostFloppyDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostFloppyDrive"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostFloppyDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostFloppyDriveResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfaceByName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindHostNetworkInterfaceByNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfaceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostfindHostNetworkInterfaceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfaceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindHostNetworkInterfacesOfType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfacesOfType"`

	This  string                    `xml:"_this,omitempty"`
	Type_ *HostNetworkInterfaceType `xml:"type,omitempty"`
}

type IHostfindHostNetworkInterfacesOfTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findHostNetworkInterfacesOfTypeResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IHostfindUSBDeviceById struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceById"`

	This string `xml:"_this,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type IHostfindUSBDeviceByIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostfindUSBDeviceByAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByAddress"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IHostfindUSBDeviceByAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_findUSBDeviceByAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IHostgenerateMACAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_generateMACAddress"`

	This string `xml:"_this,omitempty"`
}

type IHostgenerateMACAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHost_generateMACAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestVRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestVRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestVRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestVRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestVRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestVRAM"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestVRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestVRAMResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinGuestCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestCPUCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMinGuestCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinGuestCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestCPUCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestCPUCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestCPUCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestCPUCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxGuestMonitors struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestMonitors"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxGuestMonitorsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxGuestMonitorsResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetInfoVDSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getInfoVDSize"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetInfoVDSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getInfoVDSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetSerialPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getSerialPortCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetSerialPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getSerialPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetParallelPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getParallelPortCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetParallelPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getParallelPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxBootPosition struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxBootPosition"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMaxBootPositionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxBootPositionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetExclusiveHwVirt struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getExclusiveHwVirt"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetExclusiveHwVirtResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getExclusiveHwVirtResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISystemPropertiessetExclusiveHwVirt struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setExclusiveHwVirt"`

	This            string `xml:"_this,omitempty"`
	ExclusiveHwVirt bool   `xml:"exclusiveHwVirt,omitempty"`
}

type ISystemPropertiessetExclusiveHwVirtResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setExclusiveHwVirtResponse"`
}

type ISystemPropertiesgetDefaultMachineFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultMachineFolder"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultMachineFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultMachineFolderResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultMachineFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultMachineFolder"`

	This                 string `xml:"_this,omitempty"`
	DefaultMachineFolder string `xml:"defaultMachineFolder,omitempty"`
}

type ISystemPropertiessetDefaultMachineFolderResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultMachineFolderResponse"`
}

type ISystemPropertiesgetLoggingLevel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLoggingLevel"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetLoggingLevelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLoggingLevelResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetLoggingLevel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLoggingLevel"`

	This         string `xml:"_this,omitempty"`
	LoggingLevel string `xml:"loggingLevel,omitempty"`
}

type ISystemPropertiessetLoggingLevelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLoggingLevelResponse"`
}

type ISystemPropertiesgetMediumFormats struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMediumFormats"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetMediumFormatsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMediumFormatsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDefaultHardDiskFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFormat"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultHardDiskFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultHardDiskFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFormat"`

	This                  string `xml:"_this,omitempty"`
	DefaultHardDiskFormat string `xml:"defaultHardDiskFormat,omitempty"`
}

type ISystemPropertiessetDefaultHardDiskFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultHardDiskFormatResponse"`
}

type ISystemPropertiesgetFreeDiskSpaceWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceWarning"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpaceWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceWarningResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceWarning"`

	This                 string `xml:"_this,omitempty"`
	FreeDiskSpaceWarning int64  `xml:"freeDiskSpaceWarning,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceWarningResponse"`
}

type ISystemPropertiesgetFreeDiskSpacePercentWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentWarning"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpacePercentWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentWarningResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentWarning struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentWarning"`

	This                        string `xml:"_this,omitempty"`
	FreeDiskSpacePercentWarning uint32 `xml:"freeDiskSpacePercentWarning,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentWarningResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentWarningResponse"`
}

type ISystemPropertiesgetFreeDiskSpaceError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceError"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpaceErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpaceErrorResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceError"`

	This               string `xml:"_this,omitempty"`
	FreeDiskSpaceError int64  `xml:"freeDiskSpaceError,omitempty"`
}

type ISystemPropertiessetFreeDiskSpaceErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpaceErrorResponse"`
}

type ISystemPropertiesgetFreeDiskSpacePercentError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentError"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetFreeDiskSpacePercentErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getFreeDiskSpacePercentErrorResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentError"`

	This                      string `xml:"_this,omitempty"`
	FreeDiskSpacePercentError uint32 `xml:"freeDiskSpacePercentError,omitempty"`
}

type ISystemPropertiessetFreeDiskSpacePercentErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setFreeDiskSpacePercentErrorResponse"`
}

type ISystemPropertiesgetVRDEAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getVRDEAuthLibrary"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetVRDEAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getVRDEAuthLibraryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetVRDEAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setVRDEAuthLibrary"`

	This            string `xml:"_this,omitempty"`
	VRDEAuthLibrary string `xml:"VRDEAuthLibrary,omitempty"`
}

type ISystemPropertiessetVRDEAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setVRDEAuthLibraryResponse"`
}

type ISystemPropertiesgetWebServiceAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getWebServiceAuthLibrary"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetWebServiceAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getWebServiceAuthLibraryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetWebServiceAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setWebServiceAuthLibrary"`

	This                  string `xml:"_this,omitempty"`
	WebServiceAuthLibrary string `xml:"webServiceAuthLibrary,omitempty"`
}

type ISystemPropertiessetWebServiceAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setWebServiceAuthLibraryResponse"`
}

type ISystemPropertiesgetDefaultVRDEExtPack struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultVRDEExtPack"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultVRDEExtPackResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultVRDEExtPackResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultVRDEExtPack struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultVRDEExtPack"`

	This               string `xml:"_this,omitempty"`
	DefaultVRDEExtPack string `xml:"defaultVRDEExtPack,omitempty"`
}

type ISystemPropertiessetDefaultVRDEExtPackResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultVRDEExtPackResponse"`
}

type ISystemPropertiesgetLogHistoryCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLogHistoryCount"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetLogHistoryCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getLogHistoryCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiessetLogHistoryCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLogHistoryCount"`

	This            string `xml:"_this,omitempty"`
	LogHistoryCount uint32 `xml:"logHistoryCount,omitempty"`
}

type ISystemPropertiessetLogHistoryCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setLogHistoryCountResponse"`
}

type ISystemPropertiesgetDefaultAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAudioDriver"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAudioDriverResponse"`

	Returnval *AudioDriverType `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetAutostartDatabasePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getAutostartDatabasePath"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetAutostartDatabasePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getAutostartDatabasePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetAutostartDatabasePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setAutostartDatabasePath"`

	This                  string `xml:"_this,omitempty"`
	AutostartDatabasePath string `xml:"autostartDatabasePath,omitempty"`
}

type ISystemPropertiessetAutostartDatabasePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setAutostartDatabasePathResponse"`
}

type ISystemPropertiesgetDefaultAdditionsISO struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAdditionsISO"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultAdditionsISOResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultAdditionsISOResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultAdditionsISO struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultAdditionsISO"`

	This                string `xml:"_this,omitempty"`
	DefaultAdditionsISO string `xml:"defaultAdditionsISO,omitempty"`
}

type ISystemPropertiessetDefaultAdditionsISOResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultAdditionsISOResponse"`
}

type ISystemPropertiesgetDefaultFrontend struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultFrontend"`

	This string `xml:"_this,omitempty"`
}

type ISystemPropertiesgetDefaultFrontendResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultFrontendResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISystemPropertiessetDefaultFrontend struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultFrontend"`

	This            string `xml:"_this,omitempty"`
	DefaultFrontend string `xml:"defaultFrontend,omitempty"`
}

type ISystemPropertiessetDefaultFrontendResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_setDefaultFrontendResponse"`
}

type ISystemPropertiesgetMaxNetworkAdapters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxNetworkAdapters"`

	This    string       `xml:"_this,omitempty"`
	Chipset *ChipsetType `xml:"chipset,omitempty"`
}

type ISystemPropertiesgetMaxNetworkAdaptersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxNetworkAdaptersResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxNetworkAdaptersOfType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxNetworkAdaptersOfType"`

	This    string                 `xml:"_this,omitempty"`
	Chipset *ChipsetType           `xml:"chipset,omitempty"`
	Type_   *NetworkAttachmentType `xml:"type,omitempty"`
}

type ISystemPropertiesgetMaxNetworkAdaptersOfTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxNetworkAdaptersOfTypeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxDevicesPerPortForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxDevicesPerPortForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxDevicesPerPortForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMinPortCountForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinPortCountForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMinPortCountForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMinPortCountForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxPortCountForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxPortCountForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxPortCountForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxPortCountForStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfStorageBus"`

	This    string       `xml:"_this,omitempty"`
	Chipset *ChipsetType `xml:"chipset,omitempty"`
	Bus     *StorageBus  `xml:"bus,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfStorageBusResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDeviceTypesForStorageBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDeviceTypesForStorageBus"`

	This string      `xml:"_this,omitempty"`
	Bus  *StorageBus `xml:"bus,omitempty"`
}

type ISystemPropertiesgetDeviceTypesForStorageBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDeviceTypesForStorageBusResponse"`

	Returnval []*DeviceType `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetDefaultIoCacheSettingForStorageController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultIoCacheSettingForStorageController"`

	This           string                 `xml:"_this,omitempty"`
	ControllerType *StorageControllerType `xml:"controllerType,omitempty"`
}

type ISystemPropertiesgetDefaultIoCacheSettingForStorageControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getDefaultIoCacheSettingForStorageControllerResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfUSBControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfUSBControllerType"`

	This    string             `xml:"_this,omitempty"`
	Chipset *ChipsetType       `xml:"chipset,omitempty"`
	Type_   *USBControllerType `xml:"type,omitempty"`
}

type ISystemPropertiesgetMaxInstancesOfUSBControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISystemProperties_getMaxInstancesOfUSBControllerTypeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessiongetUser struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getUser"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetUserResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getUserResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiongetDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getDomain"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getDomainResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiongetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getName"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiongetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getId"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getIdResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessiongetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getTimeout"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getTimeoutResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessionsetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_setTimeout"`

	This    string `xml:"_this,omitempty"`
	Timeout uint32 `xml:"timeout,omitempty"`
}

type IGuestSessionsetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_setTimeoutResponse"`
}

type IGuestSessiongetProtocolVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getProtocolVersion"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetProtocolVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getProtocolVersionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessiongetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getStatusResponse"`

	Returnval *GuestSessionStatus `xml:"returnval,omitempty"`
}

type IGuestSessiongetEnvironment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getEnvironment"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetEnvironmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getEnvironmentResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestSessionsetEnvironment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_setEnvironment"`

	This        string   `xml:"_this,omitempty"`
	Environment []string `xml:"environment,omitempty"`
}

type IGuestSessionsetEnvironmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_setEnvironmentResponse"`
}

type IGuestSessiongetProcesses struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getProcesses"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetProcessesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getProcessesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestSessiongetDirectories struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getDirectories"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetDirectoriesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getDirectoriesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestSessiongetFiles struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getFiles"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetFilesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getFilesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestSessiongetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessiongetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_close"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessioncloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_closeResponse"`
}

type IGuestSessioncopyFrom struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_copyFrom"`

	This   string          `xml:"_this,omitempty"`
	Source string          `xml:"source,omitempty"`
	Dest   string          `xml:"dest,omitempty"`
	Flags  []*CopyFileFlag `xml:"flags,omitempty"`
}

type IGuestSessioncopyFromResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_copyFromResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessioncopyTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_copyTo"`

	This   string          `xml:"_this,omitempty"`
	Source string          `xml:"source,omitempty"`
	Dest   string          `xml:"dest,omitempty"`
	Flags  []*CopyFileFlag `xml:"flags,omitempty"`
}

type IGuestSessioncopyToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_copyToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryCreate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryCreate"`

	This  string                 `xml:"_this,omitempty"`
	Path  string                 `xml:"path,omitempty"`
	Mode  uint32                 `xml:"mode,omitempty"`
	Flags []*DirectoryCreateFlag `xml:"flags,omitempty"`
}

type IGuestSessiondirectoryCreateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryCreateResponse"`
}

type IGuestSessiondirectoryCreateTemp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryCreateTemp"`

	This         string `xml:"_this,omitempty"`
	TemplateName string `xml:"templateName,omitempty"`
	Mode         uint32 `xml:"mode,omitempty"`
	Path         string `xml:"path,omitempty"`
	Secure       bool   `xml:"secure,omitempty"`
}

type IGuestSessiondirectoryCreateTempResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryCreateTempResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryExists struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryExists"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessiondirectoryExistsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryExistsResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryOpen struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryOpen"`

	This   string               `xml:"_this,omitempty"`
	Path   string               `xml:"path,omitempty"`
	Filter string               `xml:"filter,omitempty"`
	Flags  []*DirectoryOpenFlag `xml:"flags,omitempty"`
}

type IGuestSessiondirectoryOpenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryOpenResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryQueryInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryQueryInfo"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessiondirectoryQueryInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryQueryInfoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryRemove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRemove"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessiondirectoryRemoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRemoveResponse"`
}

type IGuestSessiondirectoryRemoveRecursive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRemoveRecursive"`

	This  string                    `xml:"_this,omitempty"`
	Path  string                    `xml:"path,omitempty"`
	Flags []*DirectoryRemoveRecFlag `xml:"flags,omitempty"`
}

type IGuestSessiondirectoryRemoveRecursiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRemoveRecursiveResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessiondirectoryRename struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRename"`

	This   string            `xml:"_this,omitempty"`
	Source string            `xml:"source,omitempty"`
	Dest   string            `xml:"dest,omitempty"`
	Flags  []*PathRenameFlag `xml:"flags,omitempty"`
}

type IGuestSessiondirectoryRenameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directoryRenameResponse"`
}

type IGuestSessiondirectorySetACL struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directorySetACL"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
	Acl  string `xml:"acl,omitempty"`
}

type IGuestSessiondirectorySetACLResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_directorySetACLResponse"`
}

type IGuestSessionenvironmentClear struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentClear"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionenvironmentClearResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentClearResponse"`
}

type IGuestSessionenvironmentGet struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentGet"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IGuestSessionenvironmentGetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentGetResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionenvironmentSet struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentSet"`

	This  string `xml:"_this,omitempty"`
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IGuestSessionenvironmentSetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentSetResponse"`
}

type IGuestSessionenvironmentUnset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentUnset"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IGuestSessionenvironmentUnsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_environmentUnsetResponse"`
}

type IGuestSessionfileCreateTemp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileCreateTemp"`

	This         string `xml:"_this,omitempty"`
	TemplateName string `xml:"templateName,omitempty"`
	Mode         uint32 `xml:"mode,omitempty"`
	Path         string `xml:"path,omitempty"`
	Secure       bool   `xml:"secure,omitempty"`
}

type IGuestSessionfileCreateTempResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileCreateTempResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionfileExists struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileExists"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessionfileExistsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileExistsResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestSessionfileRemove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileRemove"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessionfileRemoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileRemoveResponse"`
}

type IGuestSessionfileOpen struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileOpen"`

	This         string `xml:"_this,omitempty"`
	Path         string `xml:"path,omitempty"`
	OpenMode     string `xml:"openMode,omitempty"`
	Disposition  string `xml:"disposition,omitempty"`
	CreationMode uint32 `xml:"creationMode,omitempty"`
}

type IGuestSessionfileOpenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileOpenResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionfileOpenEx struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileOpenEx"`

	This         string `xml:"_this,omitempty"`
	Path         string `xml:"path,omitempty"`
	OpenMode     string `xml:"openMode,omitempty"`
	Disposition  string `xml:"disposition,omitempty"`
	SharingMode  string `xml:"sharingMode,omitempty"`
	CreationMode uint32 `xml:"creationMode,omitempty"`
	Offset       int64  `xml:"offset,omitempty"`
}

type IGuestSessionfileOpenExResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileOpenExResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionfileQueryInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileQueryInfo"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessionfileQueryInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileQueryInfoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionfileQuerySize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileQuerySize"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessionfileQuerySizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileQuerySizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IGuestSessionfileRename struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileRename"`

	This   string            `xml:"_this,omitempty"`
	Source string            `xml:"source,omitempty"`
	Dest   string            `xml:"dest,omitempty"`
	Flags  []*PathRenameFlag `xml:"flags,omitempty"`
}

type IGuestSessionfileRenameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileRenameResponse"`
}

type IGuestSessionfileSetACL struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileSetACL"`

	This string `xml:"_this,omitempty"`
	File string `xml:"file,omitempty"`
	Acl  string `xml:"acl,omitempty"`
}

type IGuestSessionfileSetACLResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_fileSetACLResponse"`
}

type IGuestSessionprocessCreate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processCreate"`

	This        string               `xml:"_this,omitempty"`
	Command     string               `xml:"command,omitempty"`
	Arguments   []string             `xml:"arguments,omitempty"`
	Environment []string             `xml:"environment,omitempty"`
	Flags       []*ProcessCreateFlag `xml:"flags,omitempty"`
	TimeoutMS   uint32               `xml:"timeoutMS,omitempty"`
}

type IGuestSessionprocessCreateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processCreateResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionprocessCreateEx struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processCreateEx"`

	This        string               `xml:"_this,omitempty"`
	Command     string               `xml:"command,omitempty"`
	Arguments   []string             `xml:"arguments,omitempty"`
	Environment []string             `xml:"environment,omitempty"`
	Flags       []*ProcessCreateFlag `xml:"flags,omitempty"`
	TimeoutMS   uint32               `xml:"timeoutMS,omitempty"`
	Priority    *ProcessPriority     `xml:"priority,omitempty"`
	Affinity    []int32              `xml:"affinity,omitempty"`
}

type IGuestSessionprocessCreateExResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processCreateExResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionprocessGet struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processGet"`

	This string `xml:"_this,omitempty"`
	Pid  uint32 `xml:"pid,omitempty"`
}

type IGuestSessionprocessGetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_processGetResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionsymlinkCreate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkCreate"`

	This   string       `xml:"_this,omitempty"`
	Source string       `xml:"source,omitempty"`
	Target string       `xml:"target,omitempty"`
	Type_  *SymlinkType `xml:"type,omitempty"`
}

type IGuestSessionsymlinkCreateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkCreateResponse"`
}

type IGuestSessionsymlinkExists struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkExists"`

	This    string `xml:"_this,omitempty"`
	Symlink string `xml:"symlink,omitempty"`
}

type IGuestSessionsymlinkExistsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkExistsResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestSessionsymlinkRead struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkRead"`

	This    string             `xml:"_this,omitempty"`
	Symlink string             `xml:"symlink,omitempty"`
	Flags   []*SymlinkReadFlag `xml:"flags,omitempty"`
}

type IGuestSessionsymlinkReadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkReadResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionsymlinkRemoveDirectory struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkRemoveDirectory"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IGuestSessionsymlinkRemoveDirectoryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkRemoveDirectoryResponse"`
}

type IGuestSessionsymlinkRemoveFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkRemoveFile"`

	This string `xml:"_this,omitempty"`
	File string `xml:"file,omitempty"`
}

type IGuestSessionsymlinkRemoveFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_symlinkRemoveFileResponse"`
}

type IGuestSessionwaitFor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_waitFor"`

	This      string `xml:"_this,omitempty"`
	WaitFor   uint32 `xml:"waitFor,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IGuestSessionwaitForResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_waitForResponse"`

	Returnval *GuestSessionWaitResult `xml:"returnval,omitempty"`
}

type IGuestSessionwaitForArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_waitForArray"`

	This      string                     `xml:"_this,omitempty"`
	WaitFor   []*GuestSessionWaitForFlag `xml:"waitFor,omitempty"`
	TimeoutMS uint32                     `xml:"timeoutMS,omitempty"`
}

type IGuestSessionwaitForArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSession_waitForArrayResponse"`

	Returnval *GuestSessionWaitResult `xml:"returnval,omitempty"`
}

type IProcessgetArguments struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getArguments"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetArgumentsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getArgumentsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IProcessgetEnvironment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getEnvironment"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetEnvironmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getEnvironmentResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IProcessgetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProcessgetExecutablePath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getExecutablePath"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetExecutablePathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getExecutablePathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProcessgetExitCode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getExitCode"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetExitCodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getExitCodeResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IProcessgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getName"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProcessgetPID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getPID"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetPIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getPIDResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProcessgetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IProcessgetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_getStatusResponse"`

	Returnval *ProcessStatus `xml:"returnval,omitempty"`
}

type IProcesswaitFor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_waitFor"`

	This      string `xml:"_this,omitempty"`
	WaitFor   uint32 `xml:"waitFor,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IProcesswaitForResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_waitForResponse"`

	Returnval *ProcessWaitResult `xml:"returnval,omitempty"`
}

type IProcesswaitForArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_waitForArray"`

	This      string                `xml:"_this,omitempty"`
	WaitFor   []*ProcessWaitForFlag `xml:"waitFor,omitempty"`
	TimeoutMS uint32                `xml:"timeoutMS,omitempty"`
}

type IProcesswaitForArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_waitForArrayResponse"`

	Returnval *ProcessWaitResult `xml:"returnval,omitempty"`
}

type IProcessread struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_read"`

	This      string `xml:"_this,omitempty"`
	Handle    uint32 `xml:"handle,omitempty"`
	ToRead    uint32 `xml:"toRead,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IProcessreadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_readResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProcesswrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_write"`

	This      string `xml:"_this,omitempty"`
	Handle    uint32 `xml:"handle,omitempty"`
	Flags     uint32 `xml:"flags,omitempty"`
	Data      string `xml:"data,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IProcesswriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_writeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProcesswriteArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_writeArray"`

	This      string              `xml:"_this,omitempty"`
	Handle    uint32              `xml:"handle,omitempty"`
	Flags     []*ProcessInputFlag `xml:"flags,omitempty"`
	Data      string              `xml:"data,omitempty"`
	TimeoutMS uint32              `xml:"timeoutMS,omitempty"`
}

type IProcesswriteArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_writeArrayResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProcessterminate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_terminate"`

	This string `xml:"_this,omitempty"`
}

type IProcessterminateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProcess_terminateResponse"`
}

type IDirectorygetDirectoryName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_getDirectoryName"`

	This string `xml:"_this,omitempty"`
}

type IDirectorygetDirectoryNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_getDirectoryNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDirectorygetFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_getFilter"`

	This string `xml:"_this,omitempty"`
}

type IDirectorygetFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_getFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDirectoryclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_close"`

	This string `xml:"_this,omitempty"`
}

type IDirectorycloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_closeResponse"`
}

type IDirectoryread struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_read"`

	This string `xml:"_this,omitempty"`
}

type IDirectoryreadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDirectory_readResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilegetCreationMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getCreationMode"`

	This string `xml:"_this,omitempty"`
}

type IFilegetCreationModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getCreationModeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFilegetDisposition struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getDisposition"`

	This string `xml:"_this,omitempty"`
}

type IFilegetDispositionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getDispositionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilegetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IFilegetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilegetFileName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getFileName"`

	This string `xml:"_this,omitempty"`
}

type IFilegetFileNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getFileNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getId"`

	This string `xml:"_this,omitempty"`
}

type IFilegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getIdResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFilegetInitialSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getInitialSize"`

	This string `xml:"_this,omitempty"`
}

type IFilegetInitialSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getInitialSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFilegetOpenMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getOpenMode"`

	This string `xml:"_this,omitempty"`
}

type IFilegetOpenModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getOpenModeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilegetOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getOffset"`

	This string `xml:"_this,omitempty"`
}

type IFilegetOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getOffsetResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFilegetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IFilegetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_getStatusResponse"`

	Returnval *FileStatus `xml:"returnval,omitempty"`
}

type IFileclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_close"`

	This string `xml:"_this,omitempty"`
}

type IFilecloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_closeResponse"`
}

type IFilequeryInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_queryInfo"`

	This string `xml:"_this,omitempty"`
}

type IFilequeryInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_queryInfoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFileread struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_read"`

	This      string `xml:"_this,omitempty"`
	ToRead    uint32 `xml:"toRead,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IFilereadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_readResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFilereadAt struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_readAt"`

	This      string `xml:"_this,omitempty"`
	Offset    int64  `xml:"offset,omitempty"`
	ToRead    uint32 `xml:"toRead,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IFilereadAtResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_readAtResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFileseek struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_seek"`

	This   string        `xml:"_this,omitempty"`
	Offset int64         `xml:"offset,omitempty"`
	Whence *FileSeekType `xml:"whence,omitempty"`
}

type IFileseekResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_seekResponse"`
}

type IFilesetACL struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_setACL"`

	This string `xml:"_this,omitempty"`
	Acl  string `xml:"acl,omitempty"`
}

type IFilesetACLResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_setACLResponse"`
}

type IFilewrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_write"`

	This      string `xml:"_this,omitempty"`
	Data      string `xml:"data,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IFilewriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_writeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFilewriteAt struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_writeAt"`

	This      string `xml:"_this,omitempty"`
	Offset    int64  `xml:"offset,omitempty"`
	Data      string `xml:"data,omitempty"`
	TimeoutMS uint32 `xml:"timeoutMS,omitempty"`
}

type IFilewriteAtResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFile_writeAtResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetAccessTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getAccessTime"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetAccessTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getAccessTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetAllocatedSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getAllocatedSize"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetAllocatedSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getAllocatedSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetBirthTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getBirthTime"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetBirthTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getBirthTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetChangeTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getChangeTime"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetChangeTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getChangeTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetDeviceNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getDeviceNumber"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetDeviceNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getDeviceNumberResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetFileAttributes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getFileAttributes"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetFileAttributesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getFileAttributesResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFsObjInfogetGenerationId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGenerationId"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetGenerationIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGenerationIdResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetGID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGID"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetGIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGIDResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetGroupName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGroupName"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetGroupNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getGroupNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFsObjInfogetHardLinks struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getHardLinks"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetHardLinksResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getHardLinksResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetModificationTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getModificationTime"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetModificationTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getModificationTimeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getName"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFsObjInfogetNodeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getNodeId"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetNodeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getNodeIdResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetNodeIdDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getNodeIdDevice"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetNodeIdDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getNodeIdDeviceResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetObjectSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getObjectSize"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetObjectSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getObjectSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IFsObjInfogetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getType"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getTypeResponse"`

	Returnval *FsObjType `xml:"returnval,omitempty"`
}

type IFsObjInfogetUID struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUID"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetUIDResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUIDResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetUserFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUserFlags"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetUserFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUserFlagsResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFsObjInfogetUserName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUserName"`

	This string `xml:"_this,omitempty"`
}

type IFsObjInfogetUserNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFsObjInfo_getUserNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetOSTypeId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getOSTypeId"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetOSTypeIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getOSTypeIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsRunLevel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsRunLevel"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetAdditionsRunLevelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsRunLevelResponse"`

	Returnval *AdditionsRunLevelType `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsVersion"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetAdditionsVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsRevision"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetAdditionsRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsRevisionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestgetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestgetFacilities struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getFacilities"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetFacilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getFacilitiesResponse"`

	Returnval []*IAdditionsFacility `xml:"returnval,omitempty"`
}

type IGuestgetSessions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSessions"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetSessionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getSessionsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestgetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getMemoryBalloonSize"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getMemoryBalloonSizeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestsetMemoryBalloonSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setMemoryBalloonSize"`

	This              string `xml:"_this,omitempty"`
	MemoryBalloonSize uint32 `xml:"memoryBalloonSize,omitempty"`
}

type IGuestsetMemoryBalloonSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setMemoryBalloonSizeResponse"`
}

type IGuestgetStatisticsUpdateInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getStatisticsUpdateInterval"`

	This string `xml:"_this,omitempty"`
}

type IGuestgetStatisticsUpdateIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getStatisticsUpdateIntervalResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestsetStatisticsUpdateInterval struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setStatisticsUpdateInterval"`

	This                     string `xml:"_this,omitempty"`
	StatisticsUpdateInterval uint32 `xml:"statisticsUpdateInterval,omitempty"`
}

type IGuestsetStatisticsUpdateIntervalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setStatisticsUpdateIntervalResponse"`
}

type IGuestinternalGetStatistics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_internalGetStatistics"`

	This string `xml:"_this,omitempty"`
}

type IGuestinternalGetStatisticsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_internalGetStatisticsResponse"`

	CpuUser         uint32 `xml:"cpuUser,omitempty"`
	CpuKernel       uint32 `xml:"cpuKernel,omitempty"`
	CpuIdle         uint32 `xml:"cpuIdle,omitempty"`
	MemTotal        uint32 `xml:"memTotal,omitempty"`
	MemFree         uint32 `xml:"memFree,omitempty"`
	MemBalloon      uint32 `xml:"memBalloon,omitempty"`
	MemShared       uint32 `xml:"memShared,omitempty"`
	MemCache        uint32 `xml:"memCache,omitempty"`
	PagedTotal      uint32 `xml:"pagedTotal,omitempty"`
	MemAllocTotal   uint32 `xml:"memAllocTotal,omitempty"`
	MemFreeTotal    uint32 `xml:"memFreeTotal,omitempty"`
	MemBalloonTotal uint32 `xml:"memBalloonTotal,omitempty"`
	MemSharedTotal  uint32 `xml:"memSharedTotal,omitempty"`
}

type IGuestgetFacilityStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getFacilityStatus"`

	This     string                 `xml:"_this,omitempty"`
	Facility *AdditionsFacilityType `xml:"facility,omitempty"`
}

type IGuestgetFacilityStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getFacilityStatusResponse"`

	Timestamp int64                    `xml:"timestamp,omitempty"`
	Returnval *AdditionsFacilityStatus `xml:"returnval,omitempty"`
}

type IGuestgetAdditionsStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsStatus"`

	This  string                 `xml:"_this,omitempty"`
	Level *AdditionsRunLevelType `xml:"level,omitempty"`
}

type IGuestgetAdditionsStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_getAdditionsStatusResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestsetCredentials struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setCredentials"`

	This                  string `xml:"_this,omitempty"`
	UserName              string `xml:"userName,omitempty"`
	Password              string `xml:"password,omitempty"`
	Domain                string `xml:"domain,omitempty"`
	AllowInteractiveLogon bool   `xml:"allowInteractiveLogon,omitempty"`
}

type IGuestsetCredentialsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_setCredentialsResponse"`
}

type IGuestdragHGEnter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGEnter"`

	This           string               `xml:"_this,omitempty"`
	ScreenId       uint32               `xml:"screenId,omitempty"`
	Y              uint32               `xml:"y,omitempty"`
	X              uint32               `xml:"x,omitempty"`
	DefaultAction  *DragAndDropAction   `xml:"defaultAction,omitempty"`
	AllowedActions []*DragAndDropAction `xml:"allowedActions,omitempty"`
	Formats        []string             `xml:"formats,omitempty"`
}

type IGuestdragHGEnterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGEnterResponse"`

	Returnval *DragAndDropAction `xml:"returnval,omitempty"`
}

type IGuestdragHGMove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGMove"`

	This           string               `xml:"_this,omitempty"`
	ScreenId       uint32               `xml:"screenId,omitempty"`
	X              uint32               `xml:"x,omitempty"`
	Y              uint32               `xml:"y,omitempty"`
	DefaultAction  *DragAndDropAction   `xml:"defaultAction,omitempty"`
	AllowedActions []*DragAndDropAction `xml:"allowedActions,omitempty"`
	Formats        []string             `xml:"formats,omitempty"`
}

type IGuestdragHGMoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGMoveResponse"`

	Returnval *DragAndDropAction `xml:"returnval,omitempty"`
}

type IGuestdragHGLeave struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGLeave"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IGuestdragHGLeaveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGLeaveResponse"`
}

type IGuestdragHGDrop struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGDrop"`

	This           string               `xml:"_this,omitempty"`
	ScreenId       uint32               `xml:"screenId,omitempty"`
	X              uint32               `xml:"x,omitempty"`
	Y              uint32               `xml:"y,omitempty"`
	DefaultAction  *DragAndDropAction   `xml:"defaultAction,omitempty"`
	AllowedActions []*DragAndDropAction `xml:"allowedActions,omitempty"`
	Formats        []string             `xml:"formats,omitempty"`
}

type IGuestdragHGDropResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGDropResponse"`

	Format    string             `xml:"format,omitempty"`
	Returnval *DragAndDropAction `xml:"returnval,omitempty"`
}

type IGuestdragHGPutData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGPutData"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	Format   string `xml:"format,omitempty"`
	Data     string `xml:"data,omitempty"`
}

type IGuestdragHGPutDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragHGPutDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestdragGHPending struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHPending"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IGuestdragGHPendingResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHPendingResponse"`

	Formats        []string             `xml:"formats,omitempty"`
	AllowedActions []*DragAndDropAction `xml:"allowedActions,omitempty"`
	Returnval      *DragAndDropAction   `xml:"returnval,omitempty"`
}

type IGuestdragGHDropped struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHDropped"`

	This   string             `xml:"_this,omitempty"`
	Format string             `xml:"format,omitempty"`
	Action *DragAndDropAction `xml:"action,omitempty"`
}

type IGuestdragGHDroppedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHDroppedResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestdragGHGetData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHGetData"`

	This string `xml:"_this,omitempty"`
}

type IGuestdragGHGetDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_dragGHGetDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestcreateSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_createSession"`

	This        string `xml:"_this,omitempty"`
	User        string `xml:"user,omitempty"`
	Password    string `xml:"password,omitempty"`
	Domain      string `xml:"domain,omitempty"`
	SessionName string `xml:"sessionName,omitempty"`
}

type IGuestcreateSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_createSessionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestfindSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_findSession"`

	This        string `xml:"_this,omitempty"`
	SessionName string `xml:"sessionName,omitempty"`
}

type IGuestfindSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_findSessionResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IGuestupdateGuestAdditions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_updateGuestAdditions"`

	This      string                 `xml:"_this,omitempty"`
	Source    string                 `xml:"source,omitempty"`
	Arguments []string               `xml:"arguments,omitempty"`
	Flags     []*AdditionsUpdateFlag `xml:"flags,omitempty"`
}

type IGuestupdateGuestAdditionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuest_updateGuestAdditionsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getId"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetInitiator struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getInitiator"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetInitiatorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getInitiatorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetCancelable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCancelable"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCancelableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCancelableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetPercent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getPercent"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetPercentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getPercentResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetTimeRemaining struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeRemaining"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetTimeRemainingResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeRemainingResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IProgressgetCompleted struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCompleted"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCompletedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCompletedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetCanceled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCanceled"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetCanceledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getCanceledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IProgressgetResultCode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getResultCode"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetResultCodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getResultCodeResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IProgressgetErrorInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getErrorInfo"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetErrorInfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getErrorInfoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetOperationCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationCount"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetOperation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperation"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetOperationDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationDescription"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IProgressgetOperationPercent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationPercent"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationPercentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationPercentResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetOperationWeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationWeight"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetOperationWeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getOperationWeightResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgressgetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeout"`

	This string `xml:"_this,omitempty"`
}

type IProgressgetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_getTimeoutResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IProgresssetTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setTimeout"`

	This    string `xml:"_this,omitempty"`
	Timeout uint32 `xml:"timeout,omitempty"`
}

type IProgresssetTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setTimeoutResponse"`
}

type IProgresssetCurrentOperationProgress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setCurrentOperationProgress"`

	This    string `xml:"_this,omitempty"`
	Percent uint32 `xml:"percent,omitempty"`
}

type IProgresssetCurrentOperationProgressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setCurrentOperationProgressResponse"`
}

type IProgresssetNextOperation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setNextOperation"`

	This                     string `xml:"_this,omitempty"`
	NextOperationDescription string `xml:"nextOperationDescription,omitempty"`
	NextOperationsWeight     uint32 `xml:"nextOperationsWeight,omitempty"`
}

type IProgresssetNextOperationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_setNextOperationResponse"`
}

type IProgresswaitForCompletion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForCompletion"`

	This    string `xml:"_this,omitempty"`
	Timeout int32  `xml:"timeout,omitempty"`
}

type IProgresswaitForCompletionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForCompletionResponse"`
}

type IProgresswaitForOperationCompletion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForOperationCompletion"`

	This      string `xml:"_this,omitempty"`
	Operation uint32 `xml:"operation,omitempty"`
	Timeout   int32  `xml:"timeout,omitempty"`
}

type IProgresswaitForOperationCompletionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForOperationCompletionResponse"`
}

type IProgresswaitForAsyncProgressCompletion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForAsyncProgressCompletion"`

	This           string `xml:"_this,omitempty"`
	PProgressAsync string `xml:"pProgressAsync,omitempty"`
}

type IProgresswaitForAsyncProgressCompletionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_waitForAsyncProgressCompletionResponse"`
}

type IProgresscancel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_cancel"`

	This string `xml:"_this,omitempty"`
}

type IProgresscancelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IProgress_cancelResponse"`
}

type ISnapshotgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getId"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getName"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotsetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ISnapshotsetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setNameResponse"`
}

type ISnapshotgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getDescription"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotsetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type ISnapshotsetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_setDescriptionResponse"`
}

type ISnapshotgetTimeStamp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getTimeStamp"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetTimeStampResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getTimeStampResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type ISnapshotgetOnline struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getOnline"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetOnlineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getOnlineResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISnapshotgetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getMachine"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getParent"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotgetChildren struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildren"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildrenResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type ISnapshotgetChildrenCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildrenCount"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotgetChildrenCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshot_getChildrenCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMediumgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getId"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setDescription"`

	This        string `xml:"_this,omitempty"`
	Description string `xml:"description,omitempty"`
}

type IMediumsetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setDescriptionResponse"`
}

type IMediumgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getState"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getStateResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumgetVariant struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getVariant"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetVariantResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getVariantResponse"`

	Returnval []*MediumVariant `xml:"returnval,omitempty"`
}

type IMediumgetLocation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLocation"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLocationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLocationResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getName"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetDeviceType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDeviceType"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetDeviceTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getDeviceTypeResponse"`

	Returnval *DeviceType `xml:"returnval,omitempty"`
}

type IMediumgetHostDrive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getHostDrive"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetHostDriveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getHostDriveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumgetSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSize"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMediumgetFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getFormat"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetMediumFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMediumFormat"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetMediumFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMediumFormatResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getType"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getTypeResponse"`

	Returnval *MediumType `xml:"returnval,omitempty"`
}

type IMediumsetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setType"`

	This  string      `xml:"_this,omitempty"`
	Type_ *MediumType `xml:"type,omitempty"`
}

type IMediumsetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setTypeResponse"`
}

type IMediumgetAllowedTypes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAllowedTypes"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetAllowedTypesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAllowedTypesResponse"`

	Returnval []*MediumType `xml:"returnval,omitempty"`
}

type IMediumgetParent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getParent"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetParentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getParentResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetChildren struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getChildren"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getChildrenResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumgetBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getBase"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getBaseResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetReadOnly struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getReadOnly"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetReadOnlyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getReadOnlyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumgetLogicalSize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLogicalSize"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLogicalSizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLogicalSizeResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IMediumgetAutoReset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAutoReset"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetAutoResetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getAutoResetResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumsetAutoReset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setAutoReset"`

	This      string `xml:"_this,omitempty"`
	AutoReset bool   `xml:"autoReset,omitempty"`
}

type IMediumsetAutoResetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setAutoResetResponse"`
}

type IMediumgetLastAccessError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLastAccessError"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetLastAccessErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getLastAccessErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumgetMachineIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMachineIds"`

	This string `xml:"_this,omitempty"`
}

type IMediumgetMachineIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getMachineIdsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumsetIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setIds"`

	This        string `xml:"_this,omitempty"`
	SetImageId  bool   `xml:"setImageId,omitempty"`
	ImageId     string `xml:"imageId,omitempty"`
	SetParentId bool   `xml:"setParentId,omitempty"`
	ParentId    string `xml:"parentId,omitempty"`
}

type IMediumsetIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setIdsResponse"`
}

type IMediumrefreshState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_refreshState"`

	This string `xml:"_this,omitempty"`
}

type IMediumrefreshStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_refreshStateResponse"`

	Returnval *MediumState `xml:"returnval,omitempty"`
}

type IMediumgetSnapshotIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSnapshotIds"`

	This      string `xml:"_this,omitempty"`
	MachineId string `xml:"machineId,omitempty"`
}

type IMediumgetSnapshotIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getSnapshotIdsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IMediumlockRead struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockRead"`

	This string `xml:"_this,omitempty"`
}

type IMediumlockReadResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockReadResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumlockWrite struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockWrite"`

	This string `xml:"_this,omitempty"`
}

type IMediumlockWriteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_lockWriteResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumclose struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_close"`

	This string `xml:"_this,omitempty"`
}

type IMediumcloseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_closeResponse"`
}

type IMediumgetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getProperty"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IMediumgetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getPropertyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setProperty"`

	This  string `xml:"_this,omitempty"`
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IMediumsetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setPropertyResponse"`
}

type IMediumgetProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getProperties"`

	This  string `xml:"_this,omitempty"`
	Names string `xml:"names,omitempty"`
}

type IMediumgetPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_getPropertiesResponse"`

	ReturnNames []string `xml:"returnNames,omitempty"`
	Returnval   []string `xml:"returnval,omitempty"`
}

type IMediumsetProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setProperties"`

	This   string   `xml:"_this,omitempty"`
	Names  []string `xml:"names,omitempty"`
	Values []string `xml:"values,omitempty"`
}

type IMediumsetPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setPropertiesResponse"`
}

type IMediumcreateBaseStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createBaseStorage"`

	This        string           `xml:"_this,omitempty"`
	LogicalSize int64            `xml:"logicalSize,omitempty"`
	Variant     []*MediumVariant `xml:"variant,omitempty"`
}

type IMediumcreateBaseStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createBaseStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumdeleteStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_deleteStorage"`

	This string `xml:"_this,omitempty"`
}

type IMediumdeleteStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_deleteStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcreateDiffStorage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createDiffStorage"`

	This    string           `xml:"_this,omitempty"`
	Target  string           `xml:"target,omitempty"`
	Variant []*MediumVariant `xml:"variant,omitempty"`
}

type IMediumcreateDiffStorageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_createDiffStorageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediummergeTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_mergeTo"`

	This   string `xml:"_this,omitempty"`
	Target string `xml:"target,omitempty"`
}

type IMediummergeToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_mergeToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcloneTo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneTo"`

	This    string           `xml:"_this,omitempty"`
	Target  string           `xml:"target,omitempty"`
	Variant []*MediumVariant `xml:"variant,omitempty"`
	Parent  string           `xml:"parent,omitempty"`
}

type IMediumcloneToResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneToResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcloneToBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneToBase"`

	This    string           `xml:"_this,omitempty"`
	Target  string           `xml:"target,omitempty"`
	Variant []*MediumVariant `xml:"variant,omitempty"`
}

type IMediumcloneToBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_cloneToBaseResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumsetLocation struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setLocation"`

	This     string `xml:"_this,omitempty"`
	Location string `xml:"location,omitempty"`
}

type IMediumsetLocationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_setLocationResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumcompact struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_compact"`

	This string `xml:"_this,omitempty"`
}

type IMediumcompactResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_compactResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumresize struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resize"`

	This        string `xml:"_this,omitempty"`
	LogicalSize int64  `xml:"logicalSize,omitempty"`
}

type IMediumresizeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resizeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumreset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_reset"`

	This string `xml:"_this,omitempty"`
}

type IMediumresetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMedium_resetResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getId"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getName"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumFormatgetCapabilities struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getCapabilities"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatgetCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_getCapabilitiesResponse"`

	Returnval []*MediumFormatCapabilities `xml:"returnval,omitempty"`
}

type IMediumFormatdescribeFileExtensions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describeFileExtensions"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatdescribeFileExtensionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describeFileExtensionsResponse"`

	Extensions []string      `xml:"extensions,omitempty"`
	Types      []*DeviceType `xml:"types,omitempty"`
}

type IMediumFormatdescribeProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describeProperties"`

	This string `xml:"_this,omitempty"`
}

type IMediumFormatdescribePropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumFormat_describePropertiesResponse"`

	Names        []string    `xml:"names,omitempty"`
	Descriptions []string    `xml:"descriptions,omitempty"`
	Types        []*DataType `xml:"types,omitempty"`
	Flags        []uint32    `xml:"flags,omitempty"`
	Defaults     []string    `xml:"defaults,omitempty"`
}

type ITokenabandon struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IToken_abandon"`

	This string `xml:"_this,omitempty"`
}

type ITokenabandonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IToken_abandonResponse"`
}

type ITokendummy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IToken_dummy"`

	This string `xml:"_this,omitempty"`
}

type ITokendummyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IToken_dummyResponse"`
}

type IKeyboardgetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardgetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IKeyboardputScancode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancode"`

	This     string `xml:"_this,omitempty"`
	Scancode int32  `xml:"scancode,omitempty"`
}

type IKeyboardputScancodeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodeResponse"`
}

type IKeyboardputScancodes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodes"`

	This      string  `xml:"_this,omitempty"`
	Scancodes []int32 `xml:"scancodes,omitempty"`
}

type IKeyboardputScancodesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putScancodesResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IKeyboardputCAD struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putCAD"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardputCADResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboard_putCADResponse"`
}

type IMousegetAbsoluteSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getAbsoluteSupported"`

	This string `xml:"_this,omitempty"`
}

type IMousegetAbsoluteSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getAbsoluteSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetRelativeSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getRelativeSupported"`

	This string `xml:"_this,omitempty"`
}

type IMousegetRelativeSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getRelativeSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetMultiTouchSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getMultiTouchSupported"`

	This string `xml:"_this,omitempty"`
}

type IMousegetMultiTouchSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getMultiTouchSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetNeedsHostCursor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getNeedsHostCursor"`

	This string `xml:"_this,omitempty"`
}

type IMousegetNeedsHostCursorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getNeedsHostCursorResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousegetEventSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getEventSource"`

	This string `xml:"_this,omitempty"`
}

type IMousegetEventSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_getEventSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMouseputMouseEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEvent"`

	This        string `xml:"_this,omitempty"`
	Dx          int32  `xml:"dx,omitempty"`
	Dy          int32  `xml:"dy,omitempty"`
	Dz          int32  `xml:"dz,omitempty"`
	Dw          int32  `xml:"dw,omitempty"`
	ButtonState int32  `xml:"buttonState,omitempty"`
}

type IMouseputMouseEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventResponse"`
}

type IMouseputMouseEventAbsolute struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventAbsolute"`

	This        string `xml:"_this,omitempty"`
	X           int32  `xml:"x,omitempty"`
	Y           int32  `xml:"y,omitempty"`
	Dz          int32  `xml:"dz,omitempty"`
	Dw          int32  `xml:"dw,omitempty"`
	ButtonState int32  `xml:"buttonState,omitempty"`
}

type IMouseputMouseEventAbsoluteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putMouseEventAbsoluteResponse"`
}

type IMouseputEventMultiTouch struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putEventMultiTouch"`

	This     string  `xml:"_this,omitempty"`
	Count    int32   `xml:"count,omitempty"`
	Contacts []int64 `xml:"contacts,omitempty"`
	ScanTime uint32  `xml:"scanTime,omitempty"`
}

type IMouseputEventMultiTouchResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putEventMultiTouchResponse"`
}

type IMouseputEventMultiTouchString struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putEventMultiTouchString"`

	This     string `xml:"_this,omitempty"`
	Count    int32  `xml:"count,omitempty"`
	Contacts string `xml:"contacts,omitempty"`
	ScanTime uint32 `xml:"scanTime,omitempty"`
}

type IMouseputEventMultiTouchStringResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouse_putEventMultiTouchStringResponse"`
}

type IFramebuffergetWidth struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getWidth"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetWidthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getWidthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetHeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getHeight"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetHeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getHeightResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetBitsPerPixel struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getBitsPerPixel"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetBitsPerPixelResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getBitsPerPixelResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetBytesPerLine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getBytesPerLine"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetBytesPerLineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getBytesPerLineResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetPixelFormat struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getPixelFormat"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetPixelFormatResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getPixelFormatResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetUsesGuestVRAM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getUsesGuestVRAM"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetUsesGuestVRAMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getUsesGuestVRAMResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IFramebuffergetHeightReduction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getHeightReduction"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetHeightReductionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getHeightReductionResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebuffergetOverlay struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getOverlay"`

	This string `xml:"_this,omitempty"`
}

type IFramebuffergetOverlayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_getOverlayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IFramebuffervideoModeSupported struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_videoModeSupported"`

	This   string `xml:"_this,omitempty"`
	Width  uint32 `xml:"width,omitempty"`
	Height uint32 `xml:"height,omitempty"`
	Bpp    uint32 `xml:"bpp,omitempty"`
}

type IFramebuffervideoModeSupportedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebuffer_videoModeSupportedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IFramebufferOverlaygetX struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getX"`

	This string `xml:"_this,omitempty"`
}

type IFramebufferOverlaygetXResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getXResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebufferOverlaygetY struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getY"`

	This string `xml:"_this,omitempty"`
}

type IFramebufferOverlaygetYResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getYResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebufferOverlaygetVisible struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getVisible"`

	This string `xml:"_this,omitempty"`
}

type IFramebufferOverlaygetVisibleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getVisibleResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IFramebufferOverlaysetVisible struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_setVisible"`

	This    string `xml:"_this,omitempty"`
	Visible bool   `xml:"visible,omitempty"`
}

type IFramebufferOverlaysetVisibleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_setVisibleResponse"`
}

type IFramebufferOverlaygetAlpha struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getAlpha"`

	This string `xml:"_this,omitempty"`
}

type IFramebufferOverlaygetAlphaResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_getAlphaResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IFramebufferOverlaysetAlpha struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_setAlpha"`

	This  string `xml:"_this,omitempty"`
	Alpha uint32 `xml:"alpha,omitempty"`
}

type IFramebufferOverlaysetAlphaResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_setAlphaResponse"`
}

type IFramebufferOverlaymove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_move"`

	This string `xml:"_this,omitempty"`
	X    uint32 `xml:"x,omitempty"`
	Y    uint32 `xml:"y,omitempty"`
}

type IFramebufferOverlaymoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IFramebufferOverlay_moveResponse"`
}

type IDisplaygetScreenResolution struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getScreenResolution"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IDisplaygetScreenResolutionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getScreenResolutionResponse"`

	Width        uint32 `xml:"width,omitempty"`
	Height       uint32 `xml:"height,omitempty"`
	BitsPerPixel uint32 `xml:"bitsPerPixel,omitempty"`
	XOrigin      int32  `xml:"xOrigin,omitempty"`
	YOrigin      int32  `xml:"yOrigin,omitempty"`
}

type IDisplaysetFramebuffer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setFramebuffer"`

	This        string `xml:"_this,omitempty"`
	ScreenId    uint32 `xml:"screenId,omitempty"`
	Framebuffer string `xml:"framebuffer,omitempty"`
}

type IDisplaysetFramebufferResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setFramebufferResponse"`
}

type IDisplaygetFramebuffer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getFramebuffer"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IDisplaygetFramebufferResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_getFramebufferResponse"`

	Framebuffer string `xml:"framebuffer,omitempty"`
	XOrigin     int32  `xml:"xOrigin,omitempty"`
	YOrigin     int32  `xml:"yOrigin,omitempty"`
}

type IDisplaysetVideoModeHint struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setVideoModeHint"`

	This         string `xml:"_this,omitempty"`
	Display      uint32 `xml:"display,omitempty"`
	Enabled      bool   `xml:"enabled,omitempty"`
	ChangeOrigin bool   `xml:"changeOrigin,omitempty"`
	OriginX      int32  `xml:"originX,omitempty"`
	OriginY      int32  `xml:"originY,omitempty"`
	Width        uint32 `xml:"width,omitempty"`
	Height       uint32 `xml:"height,omitempty"`
	BitsPerPixel uint32 `xml:"bitsPerPixel,omitempty"`
}

type IDisplaysetVideoModeHintResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setVideoModeHintResponse"`
}

type IDisplaysetSeamlessMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setSeamlessMode"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IDisplaysetSeamlessModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_setSeamlessModeResponse"`
}

type IDisplaytakeScreenShotToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	Width    uint32 `xml:"width,omitempty"`
	Height   uint32 `xml:"height,omitempty"`
}

type IDisplaytakeScreenShotToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotToArrayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDisplaytakeScreenShotPNGToArray struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotPNGToArray"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	Width    uint32 `xml:"width,omitempty"`
	Height   uint32 `xml:"height,omitempty"`
}

type IDisplaytakeScreenShotPNGToArrayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_takeScreenShotPNGToArrayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IDisplayinvalidateAndUpdate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_invalidateAndUpdate"`

	This string `xml:"_this,omitempty"`
}

type IDisplayinvalidateAndUpdateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_invalidateAndUpdateResponse"`
}

type IDisplayresizeCompleted struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_resizeCompleted"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
}

type IDisplayresizeCompletedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_resizeCompletedResponse"`
}

type IDisplayviewportChanged struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_viewportChanged"`

	This     string `xml:"_this,omitempty"`
	ScreenId uint32 `xml:"screenId,omitempty"`
	X        uint32 `xml:"x,omitempty"`
	Y        uint32 `xml:"y,omitempty"`
	Width    uint32 `xml:"width,omitempty"`
	Height   uint32 `xml:"height,omitempty"`
}

type IDisplayviewportChangedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDisplay_viewportChangedResponse"`
}

type INetworkAdaptergetAdapterType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAdapterType"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetAdapterTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAdapterTypeResponse"`

	Returnval *NetworkAdapterType `xml:"returnval,omitempty"`
}

type INetworkAdaptersetAdapterType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAdapterType"`

	This        string              `xml:"_this,omitempty"`
	AdapterType *NetworkAdapterType `xml:"adapterType,omitempty"`
}

type INetworkAdaptersetAdapterTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAdapterTypeResponse"`
}

type INetworkAdaptergetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getSlot"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type INetworkAdaptersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setEnabledResponse"`
}

type INetworkAdaptergetMACAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getMACAddress"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetMACAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getMACAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetMACAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setMACAddress"`

	This       string `xml:"_this,omitempty"`
	MACAddress string `xml:"MACAddress,omitempty"`
}

type INetworkAdaptersetMACAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setMACAddressResponse"`
}

type INetworkAdaptergetAttachmentType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAttachmentType"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetAttachmentTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getAttachmentTypeResponse"`

	Returnval *NetworkAttachmentType `xml:"returnval,omitempty"`
}

type INetworkAdaptersetAttachmentType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAttachmentType"`

	This           string                 `xml:"_this,omitempty"`
	AttachmentType *NetworkAttachmentType `xml:"attachmentType,omitempty"`
}

type INetworkAdaptersetAttachmentTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setAttachmentTypeResponse"`
}

type INetworkAdaptergetBridgedInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBridgedInterface"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetBridgedInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBridgedInterfaceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetBridgedInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBridgedInterface"`

	This             string `xml:"_this,omitempty"`
	BridgedInterface string `xml:"bridgedInterface,omitempty"`
}

type INetworkAdaptersetBridgedInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBridgedInterfaceResponse"`
}

type INetworkAdaptergetHostOnlyInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getHostOnlyInterface"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetHostOnlyInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getHostOnlyInterfaceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetHostOnlyInterface struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setHostOnlyInterface"`

	This              string `xml:"_this,omitempty"`
	HostOnlyInterface string `xml:"hostOnlyInterface,omitempty"`
}

type INetworkAdaptersetHostOnlyInterfaceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setHostOnlyInterfaceResponse"`
}

type INetworkAdaptergetInternalNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getInternalNetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetInternalNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getInternalNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetInternalNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setInternalNetwork"`

	This            string `xml:"_this,omitempty"`
	InternalNetwork string `xml:"internalNetwork,omitempty"`
}

type INetworkAdaptersetInternalNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setInternalNetworkResponse"`
}

type INetworkAdaptergetNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATNetwork"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetNATNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setNATNetwork"`

	This       string `xml:"_this,omitempty"`
	NATNetwork string `xml:"NATNetwork,omitempty"`
}

type INetworkAdaptersetNATNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setNATNetworkResponse"`
}

type INetworkAdaptergetGenericDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getGenericDriver"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetGenericDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getGenericDriverResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetGenericDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setGenericDriver"`

	This          string `xml:"_this,omitempty"`
	GenericDriver string `xml:"genericDriver,omitempty"`
}

type INetworkAdaptersetGenericDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setGenericDriverResponse"`
}

type INetworkAdaptergetCableConnected struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getCableConnected"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetCableConnectedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getCableConnectedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetCableConnected struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setCableConnected"`

	This           string `xml:"_this,omitempty"`
	CableConnected bool   `xml:"cableConnected,omitempty"`
}

type INetworkAdaptersetCableConnectedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setCableConnectedResponse"`
}

type INetworkAdaptergetLineSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getLineSpeed"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetLineSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getLineSpeedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptersetLineSpeed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setLineSpeed"`

	This      string `xml:"_this,omitempty"`
	LineSpeed uint32 `xml:"lineSpeed,omitempty"`
}

type INetworkAdaptersetLineSpeedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setLineSpeedResponse"`
}

type INetworkAdaptergetPromiscModePolicy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getPromiscModePolicy"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetPromiscModePolicyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getPromiscModePolicyResponse"`

	Returnval *NetworkAdapterPromiscModePolicy `xml:"returnval,omitempty"`
}

type INetworkAdaptersetPromiscModePolicy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setPromiscModePolicy"`

	This              string                           `xml:"_this,omitempty"`
	PromiscModePolicy *NetworkAdapterPromiscModePolicy `xml:"promiscModePolicy,omitempty"`
}

type INetworkAdaptersetPromiscModePolicyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setPromiscModePolicyResponse"`
}

type INetworkAdaptergetTraceEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceEnabled"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetTraceEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INetworkAdaptersetTraceEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceEnabled"`

	This         string `xml:"_this,omitempty"`
	TraceEnabled bool   `xml:"traceEnabled,omitempty"`
}

type INetworkAdaptersetTraceEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceEnabledResponse"`
}

type INetworkAdaptergetTraceFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceFile"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetTraceFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getTraceFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetTraceFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceFile"`

	This      string `xml:"_this,omitempty"`
	TraceFile string `xml:"traceFile,omitempty"`
}

type INetworkAdaptersetTraceFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setTraceFileResponse"`
}

type INetworkAdaptergetNATEngine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATEngine"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetNATEngineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getNATEngineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptergetBootPriority struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBootPriority"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetBootPriorityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBootPriorityResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INetworkAdaptersetBootPriority struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBootPriority"`

	This         string `xml:"_this,omitempty"`
	BootPriority uint32 `xml:"bootPriority,omitempty"`
}

type INetworkAdaptersetBootPriorityResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBootPriorityResponse"`
}

type INetworkAdaptergetBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBandwidthGroup"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdaptergetBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getBandwidthGroupResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBandwidthGroup"`

	This           string `xml:"_this,omitempty"`
	BandwidthGroup string `xml:"bandwidthGroup,omitempty"`
}

type INetworkAdaptersetBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setBandwidthGroupResponse"`
}

type INetworkAdaptergetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getProperty"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type INetworkAdaptergetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getPropertyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INetworkAdaptersetProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setProperty"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type INetworkAdaptersetPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_setPropertyResponse"`
}

type INetworkAdaptergetProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getProperties"`

	This  string `xml:"_this,omitempty"`
	Names string `xml:"names,omitempty"`
}

type INetworkAdaptergetPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapter_getPropertiesResponse"`

	ReturnNames []string `xml:"returnNames,omitempty"`
	Returnval   []string `xml:"returnval,omitempty"`
}

type ISerialPortgetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getSlot"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISerialPortsetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type ISerialPortsetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setEnabledResponse"`
}

type ISerialPortgetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIOBase"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIOBaseResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortsetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIOBase"`

	This   string `xml:"_this,omitempty"`
	IOBase uint32 `xml:"IOBase,omitempty"`
}

type ISerialPortsetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIOBaseResponse"`
}

type ISerialPortgetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIRQ"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getIRQResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ISerialPortsetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIRQ"`

	This string `xml:"_this,omitempty"`
	IRQ  uint32 `xml:"IRQ,omitempty"`
}

type ISerialPortsetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setIRQResponse"`
}

type ISerialPortgetHostMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getHostMode"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetHostModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getHostModeResponse"`

	Returnval *PortMode `xml:"returnval,omitempty"`
}

type ISerialPortsetHostMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setHostMode"`

	This     string    `xml:"_this,omitempty"`
	HostMode *PortMode `xml:"hostMode,omitempty"`
}

type ISerialPortsetHostModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setHostModeResponse"`
}

type ISerialPortgetServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getServer"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getServerResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISerialPortsetServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setServer"`

	This   string `xml:"_this,omitempty"`
	Server bool   `xml:"server,omitempty"`
}

type ISerialPortsetServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setServerResponse"`
}

type ISerialPortgetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getPath"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortgetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISerialPortsetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setPath"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type ISerialPortsetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPort_setPathResponse"`
}

type IParallelPortgetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getSlot"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IParallelPortsetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IParallelPortsetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setEnabledResponse"`
}

type IParallelPortgetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIOBase"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIOBaseResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortsetIOBase struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIOBase"`

	This   string `xml:"_this,omitempty"`
	IOBase uint32 `xml:"IOBase,omitempty"`
}

type IParallelPortsetIOBaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIOBaseResponse"`
}

type IParallelPortgetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIRQ"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getIRQResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IParallelPortsetIRQ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIRQ"`

	This string `xml:"_this,omitempty"`
	IRQ  uint32 `xml:"IRQ,omitempty"`
}

type IParallelPortsetIRQResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setIRQResponse"`
}

type IParallelPortgetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getPath"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortgetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_getPathResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IParallelPortsetPath struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setPath"`

	This string `xml:"_this,omitempty"`
	Path string `xml:"path,omitempty"`
}

type IParallelPortsetPathResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPort_setPathResponse"`
}

type IMachineDebuggergetSingleStep struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getSingleStep"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetSingleStepResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getSingleStepResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetSingleStep struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setSingleStep"`

	This       string `xml:"_this,omitempty"`
	SingleStep bool   `xml:"singleStep,omitempty"`
}

type IMachineDebuggersetSingleStepResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setSingleStepResponse"`
}

type IMachineDebuggergetRecompileUser struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRecompileUser"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetRecompileUserResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRecompileUserResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetRecompileUser struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRecompileUser"`

	This          string `xml:"_this,omitempty"`
	RecompileUser bool   `xml:"recompileUser,omitempty"`
}

type IMachineDebuggersetRecompileUserResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRecompileUserResponse"`
}

type IMachineDebuggergetRecompileSupervisor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRecompileSupervisor"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetRecompileSupervisorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRecompileSupervisorResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetRecompileSupervisor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRecompileSupervisor"`

	This                string `xml:"_this,omitempty"`
	RecompileSupervisor bool   `xml:"recompileSupervisor,omitempty"`
}

type IMachineDebuggersetRecompileSupervisorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRecompileSupervisorResponse"`
}

type IMachineDebuggergetExecuteAllInIEM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getExecuteAllInIEM"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetExecuteAllInIEMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getExecuteAllInIEMResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetExecuteAllInIEM struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setExecuteAllInIEM"`

	This            string `xml:"_this,omitempty"`
	ExecuteAllInIEM bool   `xml:"executeAllInIEM,omitempty"`
}

type IMachineDebuggersetExecuteAllInIEMResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setExecuteAllInIEMResponse"`
}

type IMachineDebuggergetPATMEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getPATMEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetPATMEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getPATMEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetPATMEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setPATMEnabled"`

	This        string `xml:"_this,omitempty"`
	PATMEnabled bool   `xml:"PATMEnabled,omitempty"`
}

type IMachineDebuggersetPATMEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setPATMEnabledResponse"`
}

type IMachineDebuggergetCSAMEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getCSAMEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetCSAMEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getCSAMEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetCSAMEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setCSAMEnabled"`

	This        string `xml:"_this,omitempty"`
	CSAMEnabled bool   `xml:"CSAMEnabled,omitempty"`
}

type IMachineDebuggersetCSAMEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setCSAMEnabledResponse"`
}

type IMachineDebuggergetLogEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggersetLogEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setLogEnabled"`

	This       string `xml:"_this,omitempty"`
	LogEnabled bool   `xml:"logEnabled,omitempty"`
}

type IMachineDebuggersetLogEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setLogEnabledResponse"`
}

type IMachineDebuggergetLogDbgFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgFlags"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogDbgFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgFlagsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetLogDbgGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgGroups"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogDbgGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgGroupsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetLogDbgDestinations struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgDestinations"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogDbgDestinationsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogDbgDestinationsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetLogRelFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelFlags"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogRelFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelFlagsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetLogRelGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelGroups"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogRelGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelGroupsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetLogRelDestinations struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelDestinations"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetLogRelDestinationsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getLogRelDestinationsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetHWVirtExEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetHWVirtExEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggergetHWVirtExNestedPagingEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExNestedPagingEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetHWVirtExNestedPagingEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExNestedPagingEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggergetHWVirtExVPIDEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExVPIDEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetHWVirtExVPIDEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExVPIDEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggergetHWVirtExUXEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExUXEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetHWVirtExUXEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getHWVirtExUXEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggergetOSName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getOSName"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetOSNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getOSNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetOSVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getOSVersion"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetOSVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getOSVersionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetPAEEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getPAEEnabled"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetPAEEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getPAEEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineDebuggergetVirtualTimeRate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getVirtualTimeRate"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggergetVirtualTimeRateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getVirtualTimeRateResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMachineDebuggersetVirtualTimeRate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setVirtualTimeRate"`

	This            string `xml:"_this,omitempty"`
	VirtualTimeRate uint32 `xml:"virtualTimeRate,omitempty"`
}

type IMachineDebuggersetVirtualTimeRateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setVirtualTimeRateResponse"`
}

type IMachineDebuggerdumpGuestCore struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpGuestCore"`

	This        string `xml:"_this,omitempty"`
	Filename    string `xml:"filename,omitempty"`
	Compression string `xml:"compression,omitempty"`
}

type IMachineDebuggerdumpGuestCoreResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpGuestCoreResponse"`
}

type IMachineDebuggerdumpHostProcessCore struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpHostProcessCore"`

	This        string `xml:"_this,omitempty"`
	Filename    string `xml:"filename,omitempty"`
	Compression string `xml:"compression,omitempty"`
}

type IMachineDebuggerdumpHostProcessCoreResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpHostProcessCoreResponse"`
}

type IMachineDebuggerinfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_info"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
	Args string `xml:"args,omitempty"`
}

type IMachineDebuggerinfoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_infoResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggerinjectNMI struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_injectNMI"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggerinjectNMIResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_injectNMIResponse"`
}

type IMachineDebuggermodifyLogGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogGroups"`

	This     string `xml:"_this,omitempty"`
	Settings string `xml:"settings,omitempty"`
}

type IMachineDebuggermodifyLogGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogGroupsResponse"`
}

type IMachineDebuggermodifyLogFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogFlags"`

	This     string `xml:"_this,omitempty"`
	Settings string `xml:"settings,omitempty"`
}

type IMachineDebuggermodifyLogFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogFlagsResponse"`
}

type IMachineDebuggermodifyLogDestinations struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogDestinations"`

	This     string `xml:"_this,omitempty"`
	Settings string `xml:"settings,omitempty"`
}

type IMachineDebuggermodifyLogDestinationsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_modifyLogDestinationsResponse"`
}

type IMachineDebuggerreadPhysicalMemory struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_readPhysicalMemory"`

	This    string `xml:"_this,omitempty"`
	Address int64  `xml:"address,omitempty"`
	Size    uint32 `xml:"size,omitempty"`
}

type IMachineDebuggerreadPhysicalMemoryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_readPhysicalMemoryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggerwritePhysicalMemory struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_writePhysicalMemory"`

	This    string `xml:"_this,omitempty"`
	Address int64  `xml:"address,omitempty"`
	Size    uint32 `xml:"size,omitempty"`
	Bytes   string `xml:"bytes,omitempty"`
}

type IMachineDebuggerwritePhysicalMemoryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_writePhysicalMemoryResponse"`
}

type IMachineDebuggerreadVirtualMemory struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_readVirtualMemory"`

	This    string `xml:"_this,omitempty"`
	CpuId   uint32 `xml:"cpuId,omitempty"`
	Address int64  `xml:"address,omitempty"`
	Size    uint32 `xml:"size,omitempty"`
}

type IMachineDebuggerreadVirtualMemoryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_readVirtualMemoryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggerwriteVirtualMemory struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_writeVirtualMemory"`

	This    string `xml:"_this,omitempty"`
	CpuId   uint32 `xml:"cpuId,omitempty"`
	Address int64  `xml:"address,omitempty"`
	Size    uint32 `xml:"size,omitempty"`
	Bytes   string `xml:"bytes,omitempty"`
}

type IMachineDebuggerwriteVirtualMemoryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_writeVirtualMemoryResponse"`
}

type IMachineDebuggerdetectOS struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_detectOS"`

	This string `xml:"_this,omitempty"`
}

type IMachineDebuggerdetectOSResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_detectOSResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetRegister struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRegister"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
	Name  string `xml:"name,omitempty"`
}

type IMachineDebuggergetRegisterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRegisterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggergetRegisters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRegisters"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IMachineDebuggergetRegistersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getRegistersResponse"`

	Names  []string `xml:"names,omitempty"`
	Values []string `xml:"values,omitempty"`
}

type IMachineDebuggersetRegister struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRegister"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IMachineDebuggersetRegisterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRegisterResponse"`
}

type IMachineDebuggersetRegisters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRegisters"`

	This   string   `xml:"_this,omitempty"`
	CpuId  uint32   `xml:"cpuId,omitempty"`
	Names  []string `xml:"names,omitempty"`
	Values []string `xml:"values,omitempty"`
}

type IMachineDebuggersetRegistersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_setRegistersResponse"`
}

type IMachineDebuggerdumpGuestStack struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpGuestStack"`

	This  string `xml:"_this,omitempty"`
	CpuId uint32 `xml:"cpuId,omitempty"`
}

type IMachineDebuggerdumpGuestStackResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpGuestStackResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineDebuggerresetStats struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_resetStats"`

	This    string `xml:"_this,omitempty"`
	Pattern string `xml:"pattern,omitempty"`
}

type IMachineDebuggerresetStatsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_resetStatsResponse"`
}

type IMachineDebuggerdumpStats struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpStats"`

	This    string `xml:"_this,omitempty"`
	Pattern string `xml:"pattern,omitempty"`
}

type IMachineDebuggerdumpStatsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_dumpStatsResponse"`
}

type IMachineDebuggergetStats struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getStats"`

	This             string `xml:"_this,omitempty"`
	Pattern          string `xml:"pattern,omitempty"`
	WithDescriptions bool   `xml:"withDescriptions,omitempty"`
}

type IMachineDebuggergetStatsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDebugger_getStatsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersgetDeviceFilters struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_getDeviceFilters"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltersgetDeviceFiltersResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_getDeviceFiltersResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IUSBDeviceFilterscreateDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_createDeviceFilter"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IUSBDeviceFilterscreateDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_createDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersinsertDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_insertDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
	Filter   string `xml:"filter,omitempty"`
}

type IUSBDeviceFiltersinsertDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_insertDeviceFilterResponse"`
}

type IUSBDeviceFiltersremoveDeviceFilter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_removeDeviceFilter"`

	This     string `xml:"_this,omitempty"`
	Position uint32 `xml:"position,omitempty"`
}

type IUSBDeviceFiltersremoveDeviceFilterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilters_removeDeviceFilterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBControllergetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getName"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBControllergetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getType"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getTypeResponse"`

	Returnval *USBControllerType `xml:"returnval,omitempty"`
}

type IUSBControllergetUSBStandard struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getUSBStandard"`

	This string `xml:"_this,omitempty"`
}

type IUSBControllergetUSBStandardResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBController_getUSBStandardResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVendorId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVendorIdResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductIdResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRevisionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getManufacturer"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getManufacturerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProduct"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getProductResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getSerialNumber"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getSerialNumberResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetAddress struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getAddress"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetAddressResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getAddressResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDevicegetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPort"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVersion"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getVersionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetPortVersion struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortVersion"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetPortVersionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getPortVersionResponse"`

	Returnval uint16 `xml:"returnval,omitempty"`
}

type IUSBDevicegetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRemote"`

	This string `xml:"_this,omitempty"`
}

type IUSBDevicegetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDevice_getRemoteResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltergetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getName"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setName"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IUSBDeviceFiltersetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setNameResponse"`
}

type IUSBDeviceFiltergetActive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getActive"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetActiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getActiveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetActive struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setActive"`

	This   string `xml:"_this,omitempty"`
	Active bool   `xml:"active,omitempty"`
}

type IUSBDeviceFiltersetActiveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setActiveResponse"`
}

type IUSBDeviceFiltergetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getVendorId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getVendorIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetVendorId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setVendorId"`

	This     string `xml:"_this,omitempty"`
	VendorId string `xml:"vendorId,omitempty"`
}

type IUSBDeviceFiltersetVendorIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setVendorIdResponse"`
}

type IUSBDeviceFiltergetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductId"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetProductId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductId"`

	This      string `xml:"_this,omitempty"`
	ProductId string `xml:"productId,omitempty"`
}

type IUSBDeviceFiltersetProductIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductIdResponse"`
}

type IUSBDeviceFiltergetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRevision"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRevisionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetRevision struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRevision"`

	This     string `xml:"_this,omitempty"`
	Revision string `xml:"revision,omitempty"`
}

type IUSBDeviceFiltersetRevisionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRevisionResponse"`
}

type IUSBDeviceFiltergetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getManufacturer"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getManufacturerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetManufacturer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setManufacturer"`

	This         string `xml:"_this,omitempty"`
	Manufacturer string `xml:"manufacturer,omitempty"`
}

type IUSBDeviceFiltersetManufacturerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setManufacturerResponse"`
}

type IUSBDeviceFiltergetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProduct"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getProductResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetProduct struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProduct"`

	This    string `xml:"_this,omitempty"`
	Product string `xml:"product,omitempty"`
}

type IUSBDeviceFiltersetProductResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setProductResponse"`
}

type IUSBDeviceFiltergetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getSerialNumber"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getSerialNumberResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetSerialNumber struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setSerialNumber"`

	This         string `xml:"_this,omitempty"`
	SerialNumber string `xml:"serialNumber,omitempty"`
}

type IUSBDeviceFiltersetSerialNumberResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setSerialNumberResponse"`
}

type IUSBDeviceFiltergetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getPort"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setPort"`

	This string `xml:"_this,omitempty"`
	Port string `xml:"port,omitempty"`
}

type IUSBDeviceFiltersetPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setPortResponse"`
}

type IUSBDeviceFiltergetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRemote"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getRemoteResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetRemote struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRemote"`

	This   string `xml:"_this,omitempty"`
	Remote string `xml:"remote,omitempty"`
}

type IUSBDeviceFiltersetRemoteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setRemoteResponse"`
}

type IUSBDeviceFiltergetMaskedInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getMaskedInterfaces"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceFiltergetMaskedInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_getMaskedInterfacesResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IUSBDeviceFiltersetMaskedInterfaces struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setMaskedInterfaces"`

	This             string `xml:"_this,omitempty"`
	MaskedInterfaces uint32 `xml:"maskedInterfaces,omitempty"`
}

type IUSBDeviceFiltersetMaskedInterfacesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceFilter_setMaskedInterfacesResponse"`
}

type IHostUSBDevicegetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDevice_getState"`

	This string `xml:"_this,omitempty"`
}

type IHostUSBDevicegetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDevice_getStateResponse"`

	Returnval *USBDeviceState `xml:"returnval,omitempty"`
}

type IHostUSBDeviceFiltergetAction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_getAction"`

	This string `xml:"_this,omitempty"`
}

type IHostUSBDeviceFiltergetActionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_getActionResponse"`

	Returnval *USBDeviceFilterAction `xml:"returnval,omitempty"`
}

type IHostUSBDeviceFiltersetAction struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_setAction"`

	This   string                 `xml:"_this,omitempty"`
	Action *USBDeviceFilterAction `xml:"action,omitempty"`
}

type IHostUSBDeviceFiltersetActionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostUSBDeviceFilter_setActionResponse"`
}

type IAudioAdaptergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IAudioAdaptersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IAudioAdaptersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setEnabledResponse"`
}

type IAudioAdaptergetAudioController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioController"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetAudioControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioControllerResponse"`

	Returnval *AudioControllerType `xml:"returnval,omitempty"`
}

type IAudioAdaptersetAudioController struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioController"`

	This            string               `xml:"_this,omitempty"`
	AudioController *AudioControllerType `xml:"audioController,omitempty"`
}

type IAudioAdaptersetAudioControllerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioControllerResponse"`
}

type IAudioAdaptergetAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioDriver"`

	This string `xml:"_this,omitempty"`
}

type IAudioAdaptergetAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_getAudioDriverResponse"`

	Returnval *AudioDriverType `xml:"returnval,omitempty"`
}

type IAudioAdaptersetAudioDriver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioDriver"`

	This        string           `xml:"_this,omitempty"`
	AudioDriver *AudioDriverType `xml:"audioDriver,omitempty"`
}

type IAudioAdaptersetAudioDriverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAudioAdapter_setAudioDriverResponse"`
}

type IVRDEServergetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDEServersetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setEnabled"`

	This    string `xml:"_this,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type IVRDEServersetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setEnabledResponse"`
}

type IVRDEServergetAuthType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthType"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetAuthTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthTypeResponse"`

	Returnval *AuthType `xml:"returnval,omitempty"`
}

type IVRDEServersetAuthType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthType"`

	This     string    `xml:"_this,omitempty"`
	AuthType *AuthType `xml:"authType,omitempty"`
}

type IVRDEServersetAuthTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthTypeResponse"`
}

type IVRDEServergetAuthTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthTimeout"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetAuthTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthTimeoutResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IVRDEServersetAuthTimeout struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthTimeout"`

	This        string `xml:"_this,omitempty"`
	AuthTimeout uint32 `xml:"authTimeout,omitempty"`
}

type IVRDEServersetAuthTimeoutResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthTimeoutResponse"`
}

type IVRDEServergetAllowMultiConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAllowMultiConnection"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetAllowMultiConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAllowMultiConnectionResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDEServersetAllowMultiConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAllowMultiConnection"`

	This                 string `xml:"_this,omitempty"`
	AllowMultiConnection bool   `xml:"allowMultiConnection,omitempty"`
}

type IVRDEServersetAllowMultiConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAllowMultiConnectionResponse"`
}

type IVRDEServergetReuseSingleConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getReuseSingleConnection"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetReuseSingleConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getReuseSingleConnectionResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVRDEServersetReuseSingleConnection struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setReuseSingleConnection"`

	This                  string `xml:"_this,omitempty"`
	ReuseSingleConnection bool   `xml:"reuseSingleConnection,omitempty"`
}

type IVRDEServersetReuseSingleConnectionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setReuseSingleConnectionResponse"`
}

type IVRDEServergetVRDEExtPack struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEExtPack"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetVRDEExtPackResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEExtPackResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVRDEServersetVRDEExtPack struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setVRDEExtPack"`

	This        string `xml:"_this,omitempty"`
	VRDEExtPack string `xml:"VRDEExtPack,omitempty"`
}

type IVRDEServersetVRDEExtPackResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setVRDEExtPackResponse"`
}

type IVRDEServergetAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthLibrary"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getAuthLibraryResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVRDEServersetAuthLibrary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthLibrary"`

	This        string `xml:"_this,omitempty"`
	AuthLibrary string `xml:"authLibrary,omitempty"`
}

type IVRDEServersetAuthLibraryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setAuthLibraryResponse"`
}

type IVRDEServergetVRDEProperties struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEProperties"`

	This string `xml:"_this,omitempty"`
}

type IVRDEServergetVRDEPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEPropertiesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IVRDEServersetVRDEProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setVRDEProperty"`

	This  string `xml:"_this,omitempty"`
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type IVRDEServersetVRDEPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_setVRDEPropertyResponse"`
}

type IVRDEServergetVRDEProperty struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEProperty"`

	This string `xml:"_this,omitempty"`
	Key  string `xml:"key,omitempty"`
}

type IVRDEServergetVRDEPropertyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServer_getVRDEPropertyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISessiongetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getState"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getStateResponse"`

	Returnval *SessionState `xml:"returnval,omitempty"`
}

type ISessiongetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getType"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getTypeResponse"`

	Returnval *SessionType `xml:"returnval,omitempty"`
}

type ISessiongetMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getMachine"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getMachineResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISessiongetConsole struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getConsole"`

	This string `xml:"_this,omitempty"`
}

type ISessiongetConsoleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_getConsoleResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISessionunlockMachine struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_unlockMachine"`

	This string `xml:"_this,omitempty"`
}

type ISessionunlockMachineResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISession_unlockMachineResponse"`
}

type IStorageControllergetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getName"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IStorageControllergetMaxDevicesPerPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxDevicesPerPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMaxDevicesPerPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxDevicesPerPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetMinPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMinPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMinPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMinPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetMaxPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetMaxPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getMaxPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllergetInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getInstance"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getInstanceResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllersetInstance struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setInstance"`

	This     string `xml:"_this,omitempty"`
	Instance uint32 `xml:"instance,omitempty"`
}

type IStorageControllersetInstanceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setInstanceResponse"`
}

type IStorageControllergetPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getPortCount"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getPortCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IStorageControllersetPortCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setPortCount"`

	This      string `xml:"_this,omitempty"`
	PortCount uint32 `xml:"portCount,omitempty"`
}

type IStorageControllersetPortCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setPortCountResponse"`
}

type IStorageControllergetBus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBus"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetBusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBusResponse"`

	Returnval *StorageBus `xml:"returnval,omitempty"`
}

type IStorageControllergetControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getControllerType"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getControllerTypeResponse"`

	Returnval *StorageControllerType `xml:"returnval,omitempty"`
}

type IStorageControllersetControllerType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setControllerType"`

	This           string                 `xml:"_this,omitempty"`
	ControllerType *StorageControllerType `xml:"controllerType,omitempty"`
}

type IStorageControllersetControllerTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setControllerTypeResponse"`
}

type IStorageControllergetUseHostIOCache struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getUseHostIOCache"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetUseHostIOCacheResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getUseHostIOCacheResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IStorageControllersetUseHostIOCache struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setUseHostIOCache"`

	This           string `xml:"_this,omitempty"`
	UseHostIOCache bool   `xml:"useHostIOCache,omitempty"`
}

type IStorageControllersetUseHostIOCacheResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_setUseHostIOCacheResponse"`
}

type IStorageControllergetBootable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBootable"`

	This string `xml:"_this,omitempty"`
}

type IStorageControllergetBootableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageController_getBootableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IManagedObjectRefgetInterfaceName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_getInterfaceName"`

	This string `xml:"_this,omitempty"`
}

type IManagedObjectRefgetInterfaceNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_getInterfaceNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IManagedObjectRefrelease struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_release"`

	This string `xml:"_this,omitempty"`
}

type IManagedObjectRefreleaseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IManagedObjectRef_releaseResponse"`
}

type IWebsessionManagerlogon struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logon"`

	Username string `xml:"username,omitempty"`
	Password string `xml:"password,omitempty"`
}

type IWebsessionManagerlogonResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logonResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IWebsessionManagergetSessionObject struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_getSessionObject"`

	RefIVirtualBox string `xml:"refIVirtualBox,omitempty"`
}

type IWebsessionManagergetSessionObjectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_getSessionObjectResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IWebsessionManagerlogoff struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logoff"`

	RefIVirtualBox string `xml:"refIVirtualBox,omitempty"`
}

type IWebsessionManagerlogoffResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IWebsessionManager_logoffResponse"`
}

type IPerformanceMetricgetMetricName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMetricName"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMetricNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMetricNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetObject struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getObject"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetObjectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getObjectResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetDescription struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getDescription"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetDescriptionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getDescriptionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetPeriod struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getPeriod"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetPeriodResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getPeriodResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getCount"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getCountResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetUnit struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getUnit"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetUnitResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getUnitResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetMinimumValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMinimumValue"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMinimumValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMinimumValueResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IPerformanceMetricgetMaximumValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMaximumValue"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceMetricgetMaximumValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceMetric_getMaximumValueResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IPerformanceCollectorgetMetricNames struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricNames"`

	This string `xml:"_this,omitempty"`
}

type IPerformanceCollectorgetMetricNamesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricNamesResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorgetMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorgetMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_getMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorsetupMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_setupMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
	Period      uint32   `xml:"period,omitempty"`
	Count       uint32   `xml:"count,omitempty"`
}

type IPerformanceCollectorsetupMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_setupMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorenableMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_enableMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorenableMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_enableMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectordisableMetrics struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_disableMetrics"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectordisableMetricsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_disableMetricsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IPerformanceCollectorqueryMetricsData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_queryMetricsData"`

	This        string   `xml:"_this,omitempty"`
	MetricNames []string `xml:"metricNames,omitempty"`
	Objects     []string `xml:"objects,omitempty"`
}

type IPerformanceCollectorqueryMetricsDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPerformanceCollector_queryMetricsDataResponse"`

	ReturnMetricNames     []string `xml:"returnMetricNames,omitempty"`
	ReturnObjects         []string `xml:"returnObjects,omitempty"`
	ReturnUnits           []string `xml:"returnUnits,omitempty"`
	ReturnScales          []uint32 `xml:"returnScales,omitempty"`
	ReturnSequenceNumbers []uint32 `xml:"returnSequenceNumbers,omitempty"`
	ReturnDataIndices     []uint32 `xml:"returnDataIndices,omitempty"`
	ReturnDataLengths     []uint32 `xml:"returnDataLengths,omitempty"`
	Returnval             []int32  `xml:"returnval,omitempty"`
}

type INATEnginegetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetwork"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetwork"`

	This    string `xml:"_this,omitempty"`
	Network string `xml:"network,omitempty"`
}

type INATEnginesetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkResponse"`
}

type INATEnginegetHostIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getHostIP"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetHostIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getHostIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetHostIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setHostIP"`

	This   string `xml:"_this,omitempty"`
	HostIP string `xml:"hostIP,omitempty"`
}

type INATEnginesetHostIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setHostIPResponse"`
}

type INATEnginegetTFTPPrefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPPrefix"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTFTPPrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPPrefixResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTFTPPrefix struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPPrefix"`

	This       string `xml:"_this,omitempty"`
	TFTPPrefix string `xml:"TFTPPrefix,omitempty"`
}

type INATEnginesetTFTPPrefixResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPPrefixResponse"`
}

type INATEnginegetTFTPBootFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPBootFile"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTFTPBootFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPBootFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTFTPBootFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPBootFile"`

	This         string `xml:"_this,omitempty"`
	TFTPBootFile string `xml:"TFTPBootFile,omitempty"`
}

type INATEnginesetTFTPBootFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPBootFileResponse"`
}

type INATEnginegetTFTPNextServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPNextServer"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetTFTPNextServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getTFTPNextServerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATEnginesetTFTPNextServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPNextServer"`

	This           string `xml:"_this,omitempty"`
	TFTPNextServer string `xml:"TFTPNextServer,omitempty"`
}

type INATEnginesetTFTPNextServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setTFTPNextServerResponse"`
}

type INATEnginegetAliasMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getAliasMode"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetAliasModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getAliasModeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INATEnginesetAliasMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setAliasMode"`

	This      string `xml:"_this,omitempty"`
	AliasMode uint32 `xml:"aliasMode,omitempty"`
}

type INATEnginesetAliasModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setAliasModeResponse"`
}

type INATEnginegetDNSPassDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSPassDomain"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDNSPassDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSPassDomainResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDNSPassDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSPassDomain"`

	This          string `xml:"_this,omitempty"`
	DNSPassDomain bool   `xml:"DNSPassDomain,omitempty"`
}

type INATEnginesetDNSPassDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSPassDomainResponse"`
}

type INATEnginegetDNSProxy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSProxy"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDNSProxyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSProxyResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDNSProxy struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSProxy"`

	This     string `xml:"_this,omitempty"`
	DNSProxy bool   `xml:"DNSProxy,omitempty"`
}

type INATEnginesetDNSProxyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSProxyResponse"`
}

type INATEnginegetDNSUseHostResolver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSUseHostResolver"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetDNSUseHostResolverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getDNSUseHostResolverResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATEnginesetDNSUseHostResolver struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSUseHostResolver"`

	This               string `xml:"_this,omitempty"`
	DNSUseHostResolver bool   `xml:"DNSUseHostResolver,omitempty"`
}

type INATEnginesetDNSUseHostResolverResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setDNSUseHostResolverResponse"`
}

type INATEnginegetRedirects struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getRedirects"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetRedirectsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getRedirectsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type INATEnginesetNetworkSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkSettings"`

	This      string `xml:"_this,omitempty"`
	Mtu       uint32 `xml:"mtu,omitempty"`
	SockSnd   uint32 `xml:"sockSnd,omitempty"`
	SockRcv   uint32 `xml:"sockRcv,omitempty"`
	TcpWndSnd uint32 `xml:"TcpWndSnd,omitempty"`
	TcpWndRcv uint32 `xml:"TcpWndRcv,omitempty"`
}

type INATEnginesetNetworkSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_setNetworkSettingsResponse"`
}

type INATEnginegetNetworkSettings struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkSettings"`

	This string `xml:"_this,omitempty"`
}

type INATEnginegetNetworkSettingsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_getNetworkSettingsResponse"`

	Mtu       uint32 `xml:"mtu,omitempty"`
	SockSnd   uint32 `xml:"sockSnd,omitempty"`
	SockRcv   uint32 `xml:"sockRcv,omitempty"`
	TcpWndSnd uint32 `xml:"TcpWndSnd,omitempty"`
	TcpWndRcv uint32 `xml:"TcpWndRcv,omitempty"`
}

type INATEngineaddRedirect struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_addRedirect"`

	This      string       `xml:"_this,omitempty"`
	Name      string       `xml:"name,omitempty"`
	Proto     *NATProtocol `xml:"proto,omitempty"`
	HostIP    string       `xml:"hostIP,omitempty"`
	HostPort  uint16       `xml:"hostPort,omitempty"`
	GuestIP   string       `xml:"guestIP,omitempty"`
	GuestPort uint16       `xml:"guestPort,omitempty"`
}

type INATEngineaddRedirectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_addRedirectResponse"`
}

type INATEngineremoveRedirect struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_removeRedirect"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type INATEngineremoveRedirectResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATEngine_removeRedirectResponse"`
}

type IBandwidthGroupgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getName"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthGroupgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IBandwidthGroupgetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getType"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthGroupgetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getTypeResponse"`

	Returnval *BandwidthGroupType `xml:"returnval,omitempty"`
}

type IBandwidthGroupgetReference struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getReference"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthGroupgetReferenceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getReferenceResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IBandwidthGroupgetMaxBytesPerSec struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getMaxBytesPerSec"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthGroupgetMaxBytesPerSecResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_getMaxBytesPerSecResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IBandwidthGroupsetMaxBytesPerSec struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_setMaxBytesPerSec"`

	This           string `xml:"_this,omitempty"`
	MaxBytesPerSec int64  `xml:"maxBytesPerSec,omitempty"`
}

type IBandwidthGroupsetMaxBytesPerSecResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroup_setMaxBytesPerSecResponse"`
}

type IBandwidthControlgetNumGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getNumGroups"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthControlgetNumGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getNumGroupsResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IBandwidthControlcreateBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_createBandwidthGroup"`

	This           string              `xml:"_this,omitempty"`
	Name           string              `xml:"name,omitempty"`
	Type_          *BandwidthGroupType `xml:"type,omitempty"`
	MaxBytesPerSec int64               `xml:"maxBytesPerSec,omitempty"`
}

type IBandwidthControlcreateBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_createBandwidthGroupResponse"`
}

type IBandwidthControldeleteBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_deleteBandwidthGroup"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IBandwidthControldeleteBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_deleteBandwidthGroupResponse"`
}

type IBandwidthControlgetBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getBandwidthGroup"`

	This string `xml:"_this,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IBandwidthControlgetBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getBandwidthGroupResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IBandwidthControlgetAllBandwidthGroups struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getAllBandwidthGroups"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthControlgetAllBandwidthGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthControl_getAllBandwidthGroupsResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IEventSourcecreateListener struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_createListener"`

	This string `xml:"_this,omitempty"`
}

type IEventSourcecreateListenerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_createListenerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventSourcecreateAggregator struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_createAggregator"`

	This         string   `xml:"_this,omitempty"`
	Subordinates []string `xml:"subordinates,omitempty"`
}

type IEventSourcecreateAggregatorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_createAggregatorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventSourceregisterListener struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_registerListener"`

	This        string           `xml:"_this,omitempty"`
	Listener    string           `xml:"listener,omitempty"`
	Interesting []*VBoxEventType `xml:"interesting,omitempty"`
	Active      bool             `xml:"active,omitempty"`
}

type IEventSourceregisterListenerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_registerListenerResponse"`
}

type IEventSourceunregisterListener struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_unregisterListener"`

	This     string `xml:"_this,omitempty"`
	Listener string `xml:"listener,omitempty"`
}

type IEventSourceunregisterListenerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_unregisterListenerResponse"`
}

type IEventSourcefireEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_fireEvent"`

	This    string `xml:"_this,omitempty"`
	Event   string `xml:"event,omitempty"`
	Timeout int32  `xml:"timeout,omitempty"`
}

type IEventSourcefireEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_fireEventResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IEventSourcegetEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_getEvent"`

	This     string `xml:"_this,omitempty"`
	Listener string `xml:"listener,omitempty"`
	Timeout  int32  `xml:"timeout,omitempty"`
}

type IEventSourcegetEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_getEventResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventSourceeventProcessed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_eventProcessed"`

	This     string `xml:"_this,omitempty"`
	Listener string `xml:"listener,omitempty"`
	Event    string `xml:"event,omitempty"`
}

type IEventSourceeventProcessedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSource_eventProcessedResponse"`
}

type IEventListenerhandleEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventListener_handleEvent"`

	This  string `xml:"_this,omitempty"`
	Event string `xml:"event,omitempty"`
}

type IEventListenerhandleEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventListener_handleEventResponse"`
}

type IEventgetType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getType"`

	This string `xml:"_this,omitempty"`
}

type IEventgetTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getTypeResponse"`

	Returnval *VBoxEventType `xml:"returnval,omitempty"`
}

type IEventgetSource struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getSource"`

	This string `xml:"_this,omitempty"`
}

type IEventgetSourceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getSourceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventgetWaitable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getWaitable"`

	This string `xml:"_this,omitempty"`
}

type IEventgetWaitableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_getWaitableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IEventsetProcessed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_setProcessed"`

	This string `xml:"_this,omitempty"`
}

type IEventsetProcessedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_setProcessedResponse"`
}

type IEventwaitProcessed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_waitProcessed"`

	This    string `xml:"_this,omitempty"`
	Timeout int32  `xml:"timeout,omitempty"`
}

type IEventwaitProcessedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEvent_waitProcessedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IReusableEventgetGeneration struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IReusableEvent_getGeneration"`

	This string `xml:"_this,omitempty"`
}

type IReusableEventgetGenerationResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IReusableEvent_getGenerationResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IReusableEventreuse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IReusableEvent_reuse"`

	This string `xml:"_this,omitempty"`
}

type IReusableEventreuseResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IReusableEvent_reuseResponse"`
}

type IMachineEventgetMachineId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineEvent_getMachineId"`

	This string `xml:"_this,omitempty"`
}

type IMachineEventgetMachineIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineEvent_getMachineIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMachineStateChangedEventgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineStateChangedEvent_getState"`

	This string `xml:"_this,omitempty"`
}

type IMachineStateChangedEventgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineStateChangedEvent_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type IMachineDataChangedEventgetTemporary struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDataChangedEvent_getTemporary"`

	This string `xml:"_this,omitempty"`
}

type IMachineDataChangedEventgetTemporaryResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineDataChangedEvent_getTemporaryResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMediumRegisteredEventgetMediumId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getMediumId"`

	This string `xml:"_this,omitempty"`
}

type IMediumRegisteredEventgetMediumIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getMediumIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumRegisteredEventgetMediumType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getMediumType"`

	This string `xml:"_this,omitempty"`
}

type IMediumRegisteredEventgetMediumTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getMediumTypeResponse"`

	Returnval *DeviceType `xml:"returnval,omitempty"`
}

type IMediumRegisteredEventgetRegistered struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getRegistered"`

	This string `xml:"_this,omitempty"`
}

type IMediumRegisteredEventgetRegisteredResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumRegisteredEvent_getRegisteredResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMachineRegisteredEventgetRegistered struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineRegisteredEvent_getRegistered"`

	This string `xml:"_this,omitempty"`
}

type IMachineRegisteredEventgetRegisteredResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMachineRegisteredEvent_getRegisteredResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ISessionStateChangedEventgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISessionStateChangedEvent_getState"`

	This string `xml:"_this,omitempty"`
}

type ISessionStateChangedEventgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISessionStateChangedEvent_getStateResponse"`

	Returnval *SessionState `xml:"returnval,omitempty"`
}

type IGuestPropertyChangedEventgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getName"`

	This string `xml:"_this,omitempty"`
}

type IGuestPropertyChangedEventgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestPropertyChangedEventgetValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getValue"`

	This string `xml:"_this,omitempty"`
}

type IGuestPropertyChangedEventgetValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getValueResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestPropertyChangedEventgetFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getFlags"`

	This string `xml:"_this,omitempty"`
}

type IGuestPropertyChangedEventgetFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestPropertyChangedEvent_getFlagsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISnapshotEventgetSnapshotId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshotEvent_getSnapshotId"`

	This string `xml:"_this,omitempty"`
}

type ISnapshotEventgetSnapshotIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISnapshotEvent_getSnapshotIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetVisible struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getVisible"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetVisibleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getVisibleResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetAlpha struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getAlpha"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetAlphaResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getAlphaResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetXhot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getXhot"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetXhotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getXhotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetYhot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getYhot"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetYhotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getYhotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetWidth struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getWidth"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetWidthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getWidthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetHeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getHeight"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetHeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getHeightResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IMousePointerShapeChangedEventgetShape struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getShape"`

	This string `xml:"_this,omitempty"`
}

type IMousePointerShapeChangedEventgetShapeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMousePointerShapeChangedEvent_getShapeResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsAbsolute struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsAbsolute"`

	This string `xml:"_this,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsAbsoluteResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsAbsoluteResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsRelative struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsRelative"`

	This string `xml:"_this,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsRelativeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsRelativeResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsMultiTouch struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsMultiTouch"`

	This string `xml:"_this,omitempty"`
}

type IMouseCapabilityChangedEventgetSupportsMultiTouchResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getSupportsMultiTouchResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IMouseCapabilityChangedEventgetNeedsHostCursor struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getNeedsHostCursor"`

	This string `xml:"_this,omitempty"`
}

type IMouseCapabilityChangedEventgetNeedsHostCursorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMouseCapabilityChangedEvent_getNeedsHostCursorResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IKeyboardLedsChangedEventgetNumLock struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getNumLock"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardLedsChangedEventgetNumLockResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getNumLockResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IKeyboardLedsChangedEventgetCapsLock struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getCapsLock"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardLedsChangedEventgetCapsLockResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getCapsLockResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IKeyboardLedsChangedEventgetScrollLock struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getScrollLock"`

	This string `xml:"_this,omitempty"`
}

type IKeyboardLedsChangedEventgetScrollLockResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IKeyboardLedsChangedEvent_getScrollLockResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IStateChangedEventgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStateChangedEvent_getState"`

	This string `xml:"_this,omitempty"`
}

type IStateChangedEventgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStateChangedEvent_getStateResponse"`

	Returnval *MachineState `xml:"returnval,omitempty"`
}

type INetworkAdapterChangedEventgetNetworkAdapter struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapterChangedEvent_getNetworkAdapter"`

	This string `xml:"_this,omitempty"`
}

type INetworkAdapterChangedEventgetNetworkAdapterResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INetworkAdapterChangedEvent_getNetworkAdapterResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISerialPortChangedEventgetSerialPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPortChangedEvent_getSerialPort"`

	This string `xml:"_this,omitempty"`
}

type ISerialPortChangedEventgetSerialPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISerialPortChangedEvent_getSerialPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IParallelPortChangedEventgetParallelPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPortChangedEvent_getParallelPort"`

	This string `xml:"_this,omitempty"`
}

type IParallelPortChangedEventgetParallelPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IParallelPortChangedEvent_getParallelPortResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IMediumChangedEventgetMediumAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumChangedEvent_getMediumAttachment"`

	This string `xml:"_this,omitempty"`
}

type IMediumChangedEventgetMediumAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumChangedEvent_getMediumAttachmentResponse"`

	Returnval *IMediumAttachment `xml:"returnval,omitempty"`
}

type IClipboardModeChangedEventgetClipboardMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IClipboardModeChangedEvent_getClipboardMode"`

	This string `xml:"_this,omitempty"`
}

type IClipboardModeChangedEventgetClipboardModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IClipboardModeChangedEvent_getClipboardModeResponse"`

	Returnval *ClipboardMode `xml:"returnval,omitempty"`
}

type IDragAndDropModeChangedEventgetDragAndDropMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDragAndDropModeChangedEvent_getDragAndDropMode"`

	This string `xml:"_this,omitempty"`
}

type IDragAndDropModeChangedEventgetDragAndDropModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IDragAndDropModeChangedEvent_getDragAndDropModeResponse"`

	Returnval *DragAndDropMode `xml:"returnval,omitempty"`
}

type ICPUChangedEventgetCPU struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUChangedEvent_getCPU"`

	This string `xml:"_this,omitempty"`
}

type ICPUChangedEventgetCPUResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUChangedEvent_getCPUResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type ICPUChangedEventgetAdd struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUChangedEvent_getAdd"`

	This string `xml:"_this,omitempty"`
}

type ICPUChangedEventgetAddResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUChangedEvent_getAddResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type ICPUExecutionCapChangedEventgetExecutionCap struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUExecutionCapChangedEvent_getExecutionCap"`

	This string `xml:"_this,omitempty"`
}

type ICPUExecutionCapChangedEventgetExecutionCapResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ICPUExecutionCapChangedEvent_getExecutionCapResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestKeyboardEventgetScancodes struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestKeyboardEvent_getScancodes"`

	This string `xml:"_this,omitempty"`
}

type IGuestKeyboardEventgetScancodesResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestKeyboardEvent_getScancodesResponse"`

	Returnval []int32 `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetMode struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getMode"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetModeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getModeResponse"`

	Returnval *GuestMouseEventMode `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetX struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getX"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetXResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getXResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetY struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getY"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetYResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getYResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetZ struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getZ"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetZResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getZResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetW struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getW"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetWResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getWResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMouseEventgetButtons struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getButtons"`

	This string `xml:"_this,omitempty"`
}

type IGuestMouseEventgetButtonsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMouseEvent_getButtonsResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetContactCount struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactCount"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetContactCountResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactCountResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetXPositions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getXPositions"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetXPositionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getXPositionsResponse"`

	Returnval []int16 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetYPositions struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getYPositions"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetYPositionsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getYPositionsResponse"`

	Returnval []int16 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetContactIds struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactIds"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetContactIdsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactIdsResponse"`

	Returnval []uint16 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetContactFlags struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactFlags"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetContactFlagsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getContactFlagsResponse"`

	Returnval []uint16 `xml:"returnval,omitempty"`
}

type IGuestMultiTouchEventgetScanTime struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getScanTime"`

	This string `xml:"_this,omitempty"`
}

type IGuestMultiTouchEventgetScanTimeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMultiTouchEvent_getScanTimeResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessionEventgetSession struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionEvent_getSession"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionEventgetSessionResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionEvent_getSessionResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionStateChangedEventgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getId"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionStateChangedEventgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getIdResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestSessionStateChangedEventgetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionStateChangedEventgetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getStatusResponse"`

	Returnval *GuestSessionStatus `xml:"returnval,omitempty"`
}

type IGuestSessionStateChangedEventgetError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getError"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionStateChangedEventgetErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionStateChangedEvent_getErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestSessionRegisteredEventgetRegistered struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionRegisteredEvent_getRegistered"`

	This string `xml:"_this,omitempty"`
}

type IGuestSessionRegisteredEventgetRegisteredResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestSessionRegisteredEvent_getRegisteredResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestProcessEventgetProcess struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessEvent_getProcess"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessEventgetProcessResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessEvent_getProcessResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestProcessEventgetPid struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessEvent_getPid"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessEventgetPidResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessEvent_getPidResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestProcessRegisteredEventgetRegistered struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessRegisteredEvent_getRegistered"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessRegisteredEventgetRegisteredResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessRegisteredEvent_getRegisteredResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestProcessStateChangedEventgetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessStateChangedEvent_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessStateChangedEventgetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessStateChangedEvent_getStatusResponse"`

	Returnval *ProcessStatus `xml:"returnval,omitempty"`
}

type IGuestProcessStateChangedEventgetError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessStateChangedEvent_getError"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessStateChangedEventgetErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessStateChangedEvent_getErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestProcessIOEventgetHandle struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessIOEvent_getHandle"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessIOEventgetHandleResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessIOEvent_getHandleResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestProcessIOEventgetProcessed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessIOEvent_getProcessed"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessIOEventgetProcessedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessIOEvent_getProcessedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestProcessInputNotifyEventgetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessInputNotifyEvent_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessInputNotifyEventgetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessInputNotifyEvent_getStatusResponse"`

	Returnval *ProcessInputStatus `xml:"returnval,omitempty"`
}

type IGuestProcessOutputEventgetData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessOutputEvent_getData"`

	This string `xml:"_this,omitempty"`
}

type IGuestProcessOutputEventgetDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestProcessOutputEvent_getDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestFileEventgetFile struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileEvent_getFile"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileEventgetFileResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileEvent_getFileResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestFileRegisteredEventgetRegistered struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileRegisteredEvent_getRegistered"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileRegisteredEventgetRegisteredResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileRegisteredEvent_getRegisteredResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IGuestFileStateChangedEventgetStatus struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileStateChangedEvent_getStatus"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileStateChangedEventgetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileStateChangedEvent_getStatusResponse"`

	Returnval *FileStatus `xml:"returnval,omitempty"`
}

type IGuestFileStateChangedEventgetError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileStateChangedEvent_getError"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileStateChangedEventgetErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileStateChangedEvent_getErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestFileIOEventgetOffset struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileIOEvent_getOffset"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileIOEventgetOffsetResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileIOEvent_getOffsetResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IGuestFileIOEventgetProcessed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileIOEvent_getProcessed"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileIOEventgetProcessedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileIOEvent_getProcessedResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestFileReadEventgetData struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileReadEvent_getData"`

	This string `xml:"_this,omitempty"`
}

type IGuestFileReadEventgetDataResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestFileReadEvent_getDataResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceStateChangedEventgetDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getDevice"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceStateChangedEventgetDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getDeviceResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IUSBDeviceStateChangedEventgetAttached struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getAttached"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceStateChangedEventgetAttachedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getAttachedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IUSBDeviceStateChangedEventgetError struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getError"`

	This string `xml:"_this,omitempty"`
}

type IUSBDeviceStateChangedEventgetErrorResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IUSBDeviceStateChangedEvent_getErrorResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type ISharedFolderChangedEventgetScope struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISharedFolderChangedEvent_getScope"`

	This string `xml:"_this,omitempty"`
}

type ISharedFolderChangedEventgetScopeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISharedFolderChangedEvent_getScopeResponse"`

	Returnval *Scope `xml:"returnval,omitempty"`
}

type IRuntimeErrorEventgetFatal struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getFatal"`

	This string `xml:"_this,omitempty"`
}

type IRuntimeErrorEventgetFatalResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getFatalResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IRuntimeErrorEventgetId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getId"`

	This string `xml:"_this,omitempty"`
}

type IRuntimeErrorEventgetIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IRuntimeErrorEventgetMessage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getMessage"`

	This string `xml:"_this,omitempty"`
}

type IRuntimeErrorEventgetMessageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IRuntimeErrorEvent_getMessageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventSourceChangedEventgetListener struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSourceChangedEvent_getListener"`

	This string `xml:"_this,omitempty"`
}

type IEventSourceChangedEventgetListenerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSourceChangedEvent_getListenerResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IEventSourceChangedEventgetAdd struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSourceChangedEvent_getAdd"`

	This string `xml:"_this,omitempty"`
}

type IEventSourceChangedEventgetAddResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IEventSourceChangedEvent_getAddResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IExtraDataChangedEventgetMachineId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getMachineId"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataChangedEventgetMachineIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getMachineIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IExtraDataChangedEventgetKey struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getKey"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataChangedEventgetKeyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getKeyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IExtraDataChangedEventgetValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getValue"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataChangedEventgetValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataChangedEvent_getValueResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVetoEventaddVeto struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_addVeto"`

	This   string `xml:"_this,omitempty"`
	Reason string `xml:"reason,omitempty"`
}

type IVetoEventaddVetoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_addVetoResponse"`
}

type IVetoEventisVetoed struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_isVetoed"`

	This string `xml:"_this,omitempty"`
}

type IVetoEventisVetoedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_isVetoedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IVetoEventgetVetos struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_getVetos"`

	This string `xml:"_this,omitempty"`
}

type IVetoEventgetVetosResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVetoEvent_getVetosResponse"`

	Returnval []string `xml:"returnval,omitempty"`
}

type IExtraDataCanChangeEventgetMachineId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getMachineId"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataCanChangeEventgetMachineIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getMachineIdResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IExtraDataCanChangeEventgetKey struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getKey"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataCanChangeEventgetKeyResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getKeyResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IExtraDataCanChangeEventgetValue struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getValue"`

	This string `xml:"_this,omitempty"`
}

type IExtraDataCanChangeEventgetValueResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IExtraDataCanChangeEvent_getValueResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IShowWindowEventgetWinId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IShowWindowEvent_getWinId"`

	This string `xml:"_this,omitempty"`
}

type IShowWindowEventgetWinIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IShowWindowEvent_getWinIdResponse"`

	Returnval int64 `xml:"returnval,omitempty"`
}

type IShowWindowEventsetWinId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IShowWindowEvent_setWinId"`

	This  string `xml:"_this,omitempty"`
	WinId int64  `xml:"winId,omitempty"`
}

type IShowWindowEventsetWinIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IShowWindowEvent_setWinIdResponse"`
}

type INATRedirectEventgetSlot struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getSlot"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetSlotResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getSlotResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type INATRedirectEventgetRemove struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getRemove"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetRemoveResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getRemoveResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATRedirectEventgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getName"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATRedirectEventgetProto struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getProto"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetProtoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getProtoResponse"`

	Returnval *NATProtocol `xml:"returnval,omitempty"`
}

type INATRedirectEventgetHostIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getHostIP"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetHostIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getHostIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATRedirectEventgetHostPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getHostPort"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetHostPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getHostPortResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type INATRedirectEventgetGuestIP struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getGuestIP"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetGuestIPResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getGuestIPResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATRedirectEventgetGuestPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getGuestPort"`

	This string `xml:"_this,omitempty"`
}

type INATRedirectEventgetGuestPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATRedirectEvent_getGuestPortResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type IHostPCIDevicePlugEventgetPlugged struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getPlugged"`

	This string `xml:"_this,omitempty"`
}

type IHostPCIDevicePlugEventgetPluggedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getPluggedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostPCIDevicePlugEventgetSuccess struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getSuccess"`

	This string `xml:"_this,omitempty"`
}

type IHostPCIDevicePlugEventgetSuccessResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getSuccessResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IHostPCIDevicePlugEventgetAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getAttachment"`

	This string `xml:"_this,omitempty"`
}

type IHostPCIDevicePlugEventgetAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getAttachmentResponse"`

	Returnval *IPCIDeviceAttachment `xml:"returnval,omitempty"`
}

type IHostPCIDevicePlugEventgetMessage struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getMessage"`

	This string `xml:"_this,omitempty"`
}

type IHostPCIDevicePlugEventgetMessageResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IHostPCIDevicePlugEvent_getMessageResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IVBoxSVCAvailabilityChangedEventgetAvailable struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVBoxSVCAvailabilityChangedEvent_getAvailable"`

	This string `xml:"_this,omitempty"`
}

type IVBoxSVCAvailabilityChangedEventgetAvailableResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVBoxSVCAvailabilityChangedEvent_getAvailableResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IBandwidthGroupChangedEventgetBandwidthGroup struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroupChangedEvent_getBandwidthGroup"`

	This string `xml:"_this,omitempty"`
}

type IBandwidthGroupChangedEventgetBandwidthGroupResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IBandwidthGroupChangedEvent_getBandwidthGroupResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetChangeType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getChangeType"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetChangeTypeResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getChangeTypeResponse"`

	Returnval *GuestMonitorChangedEventType `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetScreenId struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getScreenId"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetScreenIdResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getScreenIdResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetOriginX struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getOriginX"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetOriginXResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getOriginXResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetOriginY struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getOriginY"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetOriginYResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getOriginYResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetWidth struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getWidth"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetWidthResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getWidthResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestMonitorChangedEventgetHeight struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getHeight"`

	This string `xml:"_this,omitempty"`
}

type IGuestMonitorChangedEventgetHeightResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestMonitorChangedEvent_getHeightResponse"`

	Returnval uint32 `xml:"returnval,omitempty"`
}

type IGuestUserStateChangedEventgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getName"`

	This string `xml:"_this,omitempty"`
}

type IGuestUserStateChangedEventgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestUserStateChangedEventgetDomain struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getDomain"`

	This string `xml:"_this,omitempty"`
}

type IGuestUserStateChangedEventgetDomainResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getDomainResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IGuestUserStateChangedEventgetState struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getState"`

	This string `xml:"_this,omitempty"`
}

type IGuestUserStateChangedEventgetStateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getStateResponse"`

	Returnval *GuestUserState `xml:"returnval,omitempty"`
}

type IGuestUserStateChangedEventgetStateDetails struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getStateDetails"`

	This string `xml:"_this,omitempty"`
}

type IGuestUserStateChangedEventgetStateDetailsResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestUserStateChangedEvent_getStateDetailsResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type IStorageDeviceChangedEventgetStorageDevice struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getStorageDevice"`

	This string `xml:"_this,omitempty"`
}

type IStorageDeviceChangedEventgetStorageDeviceResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getStorageDeviceResponse"`

	Returnval *IMediumAttachment `xml:"returnval,omitempty"`
}

type IStorageDeviceChangedEventgetRemoved struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getRemoved"`

	This string `xml:"_this,omitempty"`
}

type IStorageDeviceChangedEventgetRemovedResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getRemovedResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type IStorageDeviceChangedEventgetSilent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getSilent"`

	This string `xml:"_this,omitempty"`
}

type IStorageDeviceChangedEventgetSilentResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IStorageDeviceChangedEvent_getSilentResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkChangedEventgetNetworkName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkChangedEvent_getNetworkName"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkChangedEventgetNetworkNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkChangedEvent_getNetworkNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkStartStopEventgetStartEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkStartStopEvent_getStartEvent"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkStartStopEventgetStartEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkStartStopEvent_getStartEventResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkCreationDeletionEventgetCreationEvent struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkCreationDeletionEvent_getCreationEvent"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkCreationDeletionEventgetCreationEventResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkCreationDeletionEvent_getCreationEventResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkSettingEventgetEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getEnabled"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkSettingEventgetEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkSettingEventgetNetwork struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getNetwork"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkSettingEventgetNetworkResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getNetworkResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkSettingEventgetGateway struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getGateway"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkSettingEventgetGatewayResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getGatewayResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabled struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getAdvertiseDefaultIPv6RouteEnabled"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabledResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getAdvertiseDefaultIPv6RouteEnabledResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkSettingEventgetNeedDhcpServer struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getNeedDhcpServer"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkSettingEventgetNeedDhcpServerResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkSettingEvent_getNeedDhcpServerResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetCreate struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getCreate"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetCreateResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getCreateResponse"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetIpv6 struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getIpv6"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetIpv6Response struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getIpv6Response"`

	Returnval bool `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetName struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getName"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetNameResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getNameResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetProto struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getProto"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetProtoResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getProtoResponse"`

	Returnval *NATProtocol `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetHostIp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getHostIp"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetHostIpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getHostIpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetHostPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getHostPort"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetHostPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getHostPortResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetGuestIp struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getGuestIp"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetGuestIpResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getGuestIpResponse"`

	Returnval string `xml:"returnval,omitempty"`
}

type INATNetworkPortForwardEventgetGuestPort struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getGuestPort"`

	This string `xml:"_this,omitempty"`
}

type INATNetworkPortForwardEventgetGuestPortResponse struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ INATNetworkPortForwardEvent_getGuestPortResponse"`

	Returnval int32 `xml:"returnval,omitempty"`
}

type InvalidObjectFault struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ InvalidObjectFault"`

	BadObjectID string `xml:"badObjectID,omitempty"`
}

type RuntimeFault struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ RuntimeFault"`

	ResultCode int32  `xml:"resultCode,omitempty"`
	Returnval  string `xml:"returnval,omitempty"`
}

type IPCIDeviceAttachment struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IPCIDeviceAttachment"`

	Name             string `xml:"name,omitempty"`
	IsPhysicalDevice bool   `xml:"isPhysicalDevice,omitempty"`
	HostAddress      int32  `xml:"hostAddress,omitempty"`
	GuestAddress     int32  `xml:"guestAddress,omitempty"`
}

type IVRDEServerInfo struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IVRDEServerInfo"`

	Active             bool   `xml:"active,omitempty"`
	Port               int32  `xml:"port,omitempty"`
	NumberOfClients    uint32 `xml:"numberOfClients,omitempty"`
	BeginTime          int64  `xml:"beginTime,omitempty"`
	EndTime            int64  `xml:"endTime,omitempty"`
	BytesSent          int64  `xml:"bytesSent,omitempty"`
	BytesSentTotal     int64  `xml:"bytesSentTotal,omitempty"`
	BytesReceived      int64  `xml:"bytesReceived,omitempty"`
	BytesReceivedTotal int64  `xml:"bytesReceivedTotal,omitempty"`
	User               string `xml:"user,omitempty"`
	Domain             string `xml:"domain,omitempty"`
	ClientName         string `xml:"clientName,omitempty"`
	ClientIP           string `xml:"clientIP,omitempty"`
	ClientVersion      uint32 `xml:"clientVersion,omitempty"`
	EncryptionStyle    uint32 `xml:"encryptionStyle,omitempty"`
}

type IGuestOSType struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IGuestOSType"`

	FamilyId                        string                 `xml:"familyId,omitempty"`
	FamilyDescription               string                 `xml:"familyDescription,omitempty"`
	Id                              string                 `xml:"id,omitempty"`
	Description                     string                 `xml:"description,omitempty"`
	Is64Bit                         bool                   `xml:"is64Bit,omitempty"`
	RecommendedIOAPIC               bool                   `xml:"recommendedIOAPIC,omitempty"`
	RecommendedVirtEx               bool                   `xml:"recommendedVirtEx,omitempty"`
	RecommendedRAM                  uint32                 `xml:"recommendedRAM,omitempty"`
	RecommendedVRAM                 uint32                 `xml:"recommendedVRAM,omitempty"`
	Recommended2DVideoAcceleration  bool                   `xml:"recommended2DVideoAcceleration,omitempty"`
	Recommended3DAcceleration       bool                   `xml:"recommended3DAcceleration,omitempty"`
	RecommendedHDD                  int64                  `xml:"recommendedHDD,omitempty"`
	AdapterType                     *NetworkAdapterType    `xml:"adapterType,omitempty"`
	RecommendedPAE                  bool                   `xml:"recommendedPAE,omitempty"`
	RecommendedDVDStorageController *StorageControllerType `xml:"recommendedDVDStorageController,omitempty"`
	RecommendedDVDStorageBus        *StorageBus            `xml:"recommendedDVDStorageBus,omitempty"`
	RecommendedHDStorageController  *StorageControllerType `xml:"recommendedHDStorageController,omitempty"`
	RecommendedHDStorageBus         *StorageBus            `xml:"recommendedHDStorageBus,omitempty"`
	RecommendedFirmware             *FirmwareType          `xml:"recommendedFirmware,omitempty"`
	RecommendedUSBHID               bool                   `xml:"recommendedUSBHID,omitempty"`
	RecommendedHPET                 bool                   `xml:"recommendedHPET,omitempty"`
	RecommendedUSBTablet            bool                   `xml:"recommendedUSBTablet,omitempty"`
	RecommendedRTCUseUTC            bool                   `xml:"recommendedRTCUseUTC,omitempty"`
	RecommendedChipset              *ChipsetType           `xml:"recommendedChipset,omitempty"`
	RecommendedAudioController      *AudioControllerType   `xml:"recommendedAudioController,omitempty"`
	RecommendedFloppy               bool                   `xml:"recommendedFloppy,omitempty"`
	RecommendedUSB                  bool                   `xml:"recommendedUSB,omitempty"`
}

type IAdditionsFacility struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ IAdditionsFacility"`

	ClassType   *AdditionsFacilityClass  `xml:"classType,omitempty"`
	LastUpdated int64                    `xml:"lastUpdated,omitempty"`
	Name        string                   `xml:"name,omitempty"`
	Status      *AdditionsFacilityStatus `xml:"status,omitempty"`
	Type_       *AdditionsFacilityType   `xml:"type,omitempty"`
}

type IMediumAttachment struct {
	// XMLName xml.Name `xml:"http://www.virtualbox.org/ IMediumAttachment"`

	Medium         string      `xml:"medium,omitempty"`
	Controller     string      `xml:"controller,omitempty"`
	Port           int32       `xml:"port,omitempty"`
	Device         int32       `xml:"device,omitempty"`
	Type_          *DeviceType `xml:"type,omitempty"`
	Passthrough    bool        `xml:"passthrough,omitempty"`
	TemporaryEject bool        `xml:"temporaryEject,omitempty"`
	IsEjected      bool        `xml:"isEjected,omitempty"`
	NonRotational  bool        `xml:"nonRotational,omitempty"`
	Discard        bool        `xml:"discard,omitempty"`
	HotPluggable   bool        `xml:"hotPluggable,omitempty"`
	BandwidthGroup string      `xml:"bandwidthGroup,omitempty"`
}

type ISharedFolder struct {
	XMLName xml.Name `xml:"http://www.virtualbox.org/ ISharedFolder"`

	Name            string `xml:"name,omitempty"`
	HostPath        string `xml:"hostPath,omitempty"`
	Accessible      bool   `xml:"accessible,omitempty"`
	Writable        bool   `xml:"writable,omitempty"`
	AutoMount       bool   `xml:"autoMount,omitempty"`
	LastAccessError string `xml:"lastAccessError,omitempty"`
}

type VboxPortType struct {
	client *SOAPClient
}

func NewVboxPortType(url string, tls bool, auth *BasicAuth) *VboxPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &VboxPortType{
		client: client,
	}
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetResultCode(request *IVirtualBoxErrorInfogetResultCode) (*IVirtualBoxErrorInfogetResultCodeResponse, error) {
	response := new(IVirtualBoxErrorInfogetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetResultDetail(request *IVirtualBoxErrorInfogetResultDetail) (*IVirtualBoxErrorInfogetResultDetailResponse, error) {
	response := new(IVirtualBoxErrorInfogetResultDetailResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetInterfaceID(request *IVirtualBoxErrorInfogetInterfaceID) (*IVirtualBoxErrorInfogetInterfaceIDResponse, error) {
	response := new(IVirtualBoxErrorInfogetInterfaceIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetComponent(request *IVirtualBoxErrorInfogetComponent) (*IVirtualBoxErrorInfogetComponentResponse, error) {
	response := new(IVirtualBoxErrorInfogetComponentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetText(request *IVirtualBoxErrorInfogetText) (*IVirtualBoxErrorInfogetTextResponse, error) {
	response := new(IVirtualBoxErrorInfogetTextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxErrorInfogetNext(request *IVirtualBoxErrorInfogetNext) (*IVirtualBoxErrorInfogetNextResponse, error) {
	response := new(IVirtualBoxErrorInfogetNextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNetworkName(request *INATNetworkgetNetworkName) (*INATNetworkgetNetworkNameResponse, error) {
	response := new(INATNetworkgetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNetworkName(request *INATNetworksetNetworkName) (*INATNetworksetNetworkNameResponse, error) {
	response := new(INATNetworksetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetEnabled(request *INATNetworkgetEnabled) (*INATNetworkgetEnabledResponse, error) {
	response := new(INATNetworkgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetEnabled(request *INATNetworksetEnabled) (*INATNetworksetEnabledResponse, error) {
	response := new(INATNetworksetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNetwork(request *INATNetworkgetNetwork) (*INATNetworkgetNetworkResponse, error) {
	response := new(INATNetworkgetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNetwork(request *INATNetworksetNetwork) (*INATNetworksetNetworkResponse, error) {
	response := new(INATNetworksetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetGateway(request *INATNetworkgetGateway) (*INATNetworkgetGatewayResponse, error) {
	response := new(INATNetworkgetGatewayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetIPv6Enabled(request *INATNetworkgetIPv6Enabled) (*INATNetworkgetIPv6EnabledResponse, error) {
	response := new(INATNetworkgetIPv6EnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetIPv6Enabled(request *INATNetworksetIPv6Enabled) (*INATNetworksetIPv6EnabledResponse, error) {
	response := new(INATNetworksetIPv6EnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetIPv6Prefix(request *INATNetworkgetIPv6Prefix) (*INATNetworkgetIPv6PrefixResponse, error) {
	response := new(INATNetworkgetIPv6PrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetIPv6Prefix(request *INATNetworksetIPv6Prefix) (*INATNetworksetIPv6PrefixResponse, error) {
	response := new(INATNetworksetIPv6PrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworkgetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworkgetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworkgetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworksetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworksetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworksetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetNeedDhcpServer(request *INATNetworkgetNeedDhcpServer) (*INATNetworkgetNeedDhcpServerResponse, error) {
	response := new(INATNetworkgetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetNeedDhcpServer(request *INATNetworksetNeedDhcpServer) (*INATNetworksetNeedDhcpServerResponse, error) {
	response := new(INATNetworksetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetEventSource(request *INATNetworkgetEventSource) (*INATNetworkgetEventSourceResponse, error) {
	response := new(INATNetworkgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetPortForwardRules4(request *INATNetworkgetPortForwardRules4) (*INATNetworkgetPortForwardRules4Response, error) {
	response := new(INATNetworkgetPortForwardRules4Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetLocalMappings(request *INATNetworkgetLocalMappings) (*INATNetworkgetLocalMappingsResponse, error) {
	response := new(INATNetworkgetLocalMappingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetLoopbackIp6(request *INATNetworkgetLoopbackIp6) (*INATNetworkgetLoopbackIp6Response, error) {
	response := new(INATNetworkgetLoopbackIp6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworksetLoopbackIp6(request *INATNetworksetLoopbackIp6) (*INATNetworksetLoopbackIp6Response, error) {
	response := new(INATNetworksetLoopbackIp6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkgetPortForwardRules6(request *INATNetworkgetPortForwardRules6) (*INATNetworkgetPortForwardRules6Response, error) {
	response := new(INATNetworkgetPortForwardRules6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkaddLocalMapping(request *INATNetworkaddLocalMapping) (*INATNetworkaddLocalMappingResponse, error) {
	response := new(INATNetworkaddLocalMappingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkaddPortForwardRule(request *INATNetworkaddPortForwardRule) (*INATNetworkaddPortForwardRuleResponse, error) {
	response := new(INATNetworkaddPortForwardRuleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkremovePortForwardRule(request *INATNetworkremovePortForwardRule) (*INATNetworkremovePortForwardRuleResponse, error) {
	response := new(INATNetworkremovePortForwardRuleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkstart(request *INATNetworkstart) (*INATNetworkstartResponse, error) {
	response := new(INATNetworkstartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkstop(request *INATNetworkstop) (*INATNetworkstopResponse, error) {
	response := new(INATNetworkstopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetEventSource(request *IDHCPServergetEventSource) (*IDHCPServergetEventSourceResponse, error) {
	response := new(IDHCPServergetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetEnabled(request *IDHCPServergetEnabled) (*IDHCPServergetEnabledResponse, error) {
	response := new(IDHCPServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetEnabled(request *IDHCPServersetEnabled) (*IDHCPServersetEnabledResponse, error) {
	response := new(IDHCPServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetIPAddress(request *IDHCPServergetIPAddress) (*IDHCPServergetIPAddressResponse, error) {
	response := new(IDHCPServergetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkMask(request *IDHCPServergetNetworkMask) (*IDHCPServergetNetworkMaskResponse, error) {
	response := new(IDHCPServergetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetNetworkName(request *IDHCPServergetNetworkName) (*IDHCPServergetNetworkNameResponse, error) {
	response := new(IDHCPServergetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetLowerIP(request *IDHCPServergetLowerIP) (*IDHCPServergetLowerIPResponse, error) {
	response := new(IDHCPServergetLowerIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetUpperIP(request *IDHCPServergetUpperIP) (*IDHCPServergetUpperIPResponse, error) {
	response := new(IDHCPServergetUpperIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetGlobalOptions(request *IDHCPServergetGlobalOptions) (*IDHCPServergetGlobalOptionsResponse, error) {
	response := new(IDHCPServergetGlobalOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetVmConfigs(request *IDHCPServergetVmConfigs) (*IDHCPServergetVmConfigsResponse, error) {
	response := new(IDHCPServergetVmConfigsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServeraddGlobalOption(request *IDHCPServeraddGlobalOption) (*IDHCPServeraddGlobalOptionResponse, error) {
	response := new(IDHCPServeraddGlobalOptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServeraddVmSlotOption(request *IDHCPServeraddVmSlotOption) (*IDHCPServeraddVmSlotOptionResponse, error) {
	response := new(IDHCPServeraddVmSlotOptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerremoveVmSlotOptions(request *IDHCPServerremoveVmSlotOptions) (*IDHCPServerremoveVmSlotOptionsResponse, error) {
	response := new(IDHCPServerremoveVmSlotOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetVmSlotOptions(request *IDHCPServergetVmSlotOptions) (*IDHCPServergetVmSlotOptionsResponse, error) {
	response := new(IDHCPServergetVmSlotOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServergetMacOptions(request *IDHCPServergetMacOptions) (*IDHCPServergetMacOptionsResponse, error) {
	response := new(IDHCPServergetMacOptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServersetConfiguration(request *IDHCPServersetConfiguration) (*IDHCPServersetConfigurationResponse, error) {
	response := new(IDHCPServersetConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstart(request *IDHCPServerstart) (*IDHCPServerstartResponse, error) {
	response := new(IDHCPServerstartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDHCPServerstop(request *IDHCPServerstop) (*IDHCPServerstopResponse, error) {
	response := new(IDHCPServerstopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetVersion(request *IVirtualBoxgetVersion) (*IVirtualBoxgetVersionResponse, error) {
	response := new(IVirtualBoxgetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetVersionNormalized(request *IVirtualBoxgetVersionNormalized) (*IVirtualBoxgetVersionNormalizedResponse, error) {
	response := new(IVirtualBoxgetVersionNormalizedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetRevision(request *IVirtualBoxgetRevision) (*IVirtualBoxgetRevisionResponse, error) {
	response := new(IVirtualBoxgetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPackageType(request *IVirtualBoxgetPackageType) (*IVirtualBoxgetPackageTypeResponse, error) {
	response := new(IVirtualBoxgetPackageTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetAPIVersion(request *IVirtualBoxgetAPIVersion) (*IVirtualBoxgetAPIVersionResponse, error) {
	response := new(IVirtualBoxgetAPIVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHomeFolder(request *IVirtualBoxgetHomeFolder) (*IVirtualBoxgetHomeFolderResponse, error) {
	response := new(IVirtualBoxgetHomeFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSettingsFilePath(request *IVirtualBoxgetSettingsFilePath) (*IVirtualBoxgetSettingsFilePathResponse, error) {
	response := new(IVirtualBoxgetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHost(request *IVirtualBoxgetHost) (*IVirtualBoxgetHostResponse, error) {
	response := new(IVirtualBoxgetHostResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSystemProperties(request *IVirtualBoxgetSystemProperties) (*IVirtualBoxgetSystemPropertiesResponse, error) {
	response := new(IVirtualBoxgetSystemPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachines(request *IVirtualBoxgetMachines) (*IVirtualBoxgetMachinesResponse, error) {
	response := new(IVirtualBoxgetMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachineGroups(request *IVirtualBoxgetMachineGroups) (*IVirtualBoxgetMachineGroupsResponse, error) {
	response := new(IVirtualBoxgetMachineGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetHardDisks(request *IVirtualBoxgetHardDisks) (*IVirtualBoxgetHardDisksResponse, error) {
	response := new(IVirtualBoxgetHardDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDVDImages(request *IVirtualBoxgetDVDImages) (*IVirtualBoxgetDVDImagesResponse, error) {
	response := new(IVirtualBoxgetDVDImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetFloppyImages(request *IVirtualBoxgetFloppyImages) (*IVirtualBoxgetFloppyImagesResponse, error) {
	response := new(IVirtualBoxgetFloppyImagesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetProgressOperations(request *IVirtualBoxgetProgressOperations) (*IVirtualBoxgetProgressOperationsResponse, error) {
	response := new(IVirtualBoxgetProgressOperationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSTypes(request *IVirtualBoxgetGuestOSTypes) (*IVirtualBoxgetGuestOSTypesResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetSharedFolders(request *IVirtualBoxgetSharedFolders) (*IVirtualBoxgetSharedFoldersResponse, error) {
	response := new(IVirtualBoxgetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetPerformanceCollector(request *IVirtualBoxgetPerformanceCollector) (*IVirtualBoxgetPerformanceCollectorResponse, error) {
	response := new(IVirtualBoxgetPerformanceCollectorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetDHCPServers(request *IVirtualBoxgetDHCPServers) (*IVirtualBoxgetDHCPServersResponse, error) {
	response := new(IVirtualBoxgetDHCPServersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetNATNetworks(request *IVirtualBoxgetNATNetworks) (*IVirtualBoxgetNATNetworksResponse, error) {
	response := new(IVirtualBoxgetNATNetworksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetEventSource(request *IVirtualBoxgetEventSource) (*IVirtualBoxgetEventSourceResponse, error) {
	response := new(IVirtualBoxgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetInternalNetworks(request *IVirtualBoxgetInternalNetworks) (*IVirtualBoxgetInternalNetworksResponse, error) {
	response := new(IVirtualBoxgetInternalNetworksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGenericNetworkDrivers(request *IVirtualBoxgetGenericNetworkDrivers) (*IVirtualBoxgetGenericNetworkDriversResponse, error) {
	response := new(IVirtualBoxgetGenericNetworkDriversResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcomposeMachineFilename(request *IVirtualBoxcomposeMachineFilename) (*IVirtualBoxcomposeMachineFilenameResponse, error) {
	response := new(IVirtualBoxcomposeMachineFilenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateMachine(request *IVirtualBoxcreateMachine) (*IVirtualBoxcreateMachineResponse, error) {
	response := new(IVirtualBoxcreateMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenMachine(request *IVirtualBoxopenMachine) (*IVirtualBoxopenMachineResponse, error) {
	response := new(IVirtualBoxopenMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxregisterMachine(request *IVirtualBoxregisterMachine) (*IVirtualBoxregisterMachineResponse, error) {
	response := new(IVirtualBoxregisterMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindMachine(request *IVirtualBoxfindMachine) (*IVirtualBoxfindMachineResponse, error) {
	response := new(IVirtualBoxfindMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachinesByGroups(request *IVirtualBoxgetMachinesByGroups) (*IVirtualBoxgetMachinesByGroupsResponse, error) {
	response := new(IVirtualBoxgetMachinesByGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetMachineStates(request *IVirtualBoxgetMachineStates) (*IVirtualBoxgetMachineStatesResponse, error) {
	response := new(IVirtualBoxgetMachineStatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateAppliance(request *IVirtualBoxcreateAppliance) (*IVirtualBoxcreateApplianceResponse, error) {
	response := new(IVirtualBoxcreateApplianceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateHardDisk(request *IVirtualBoxcreateHardDisk) (*IVirtualBoxcreateHardDiskResponse, error) {
	response := new(IVirtualBoxcreateHardDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxopenMedium(request *IVirtualBoxopenMedium) (*IVirtualBoxopenMediumResponse, error) {
	response := new(IVirtualBoxopenMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetGuestOSType(request *IVirtualBoxgetGuestOSType) (*IVirtualBoxgetGuestOSTypeResponse, error) {
	response := new(IVirtualBoxgetGuestOSTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateSharedFolder(request *IVirtualBoxcreateSharedFolder) (*IVirtualBoxcreateSharedFolderResponse, error) {
	response := new(IVirtualBoxcreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveSharedFolder(request *IVirtualBoxremoveSharedFolder) (*IVirtualBoxremoveSharedFolderResponse, error) {
	response := new(IVirtualBoxremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraDataKeys(request *IVirtualBoxgetExtraDataKeys) (*IVirtualBoxgetExtraDataKeysResponse, error) {
	response := new(IVirtualBoxgetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxgetExtraData(request *IVirtualBoxgetExtraData) (*IVirtualBoxgetExtraDataResponse, error) {
	response := new(IVirtualBoxgetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxsetExtraData(request *IVirtualBoxsetExtraData) (*IVirtualBoxsetExtraDataResponse, error) {
	response := new(IVirtualBoxsetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxsetSettingsSecret(request *IVirtualBoxsetSettingsSecret) (*IVirtualBoxsetSettingsSecretResponse, error) {
	response := new(IVirtualBoxsetSettingsSecretResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateDHCPServer(request *IVirtualBoxcreateDHCPServer) (*IVirtualBoxcreateDHCPServerResponse, error) {
	response := new(IVirtualBoxcreateDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindDHCPServerByNetworkName(request *IVirtualBoxfindDHCPServerByNetworkName) (*IVirtualBoxfindDHCPServerByNetworkNameResponse, error) {
	response := new(IVirtualBoxfindDHCPServerByNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveDHCPServer(request *IVirtualBoxremoveDHCPServer) (*IVirtualBoxremoveDHCPServerResponse, error) {
	response := new(IVirtualBoxremoveDHCPServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcreateNATNetwork(request *IVirtualBoxcreateNATNetwork) (*IVirtualBoxcreateNATNetworkResponse, error) {
	response := new(IVirtualBoxcreateNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxfindNATNetworkByName(request *IVirtualBoxfindNATNetworkByName) (*IVirtualBoxfindNATNetworkByNameResponse, error) {
	response := new(IVirtualBoxfindNATNetworkByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxremoveNATNetwork(request *IVirtualBoxremoveNATNetwork) (*IVirtualBoxremoveNATNetworkResponse, error) {
	response := new(IVirtualBoxremoveNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualBoxcheckFirmwarePresent(request *IVirtualBoxcheckFirmwarePresent) (*IVirtualBoxcheckFirmwarePresentResponse, error) {
	response := new(IVirtualBoxcheckFirmwarePresentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetPath(request *IVFSExplorergetPath) (*IVFSExplorergetPathResponse, error) {
	response := new(IVFSExplorergetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorergetType(request *IVFSExplorergetType) (*IVFSExplorergetTypeResponse, error) {
	response := new(IVFSExplorergetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerupdate(request *IVFSExplorerupdate) (*IVFSExplorerupdateResponse, error) {
	response := new(IVFSExplorerupdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercd(request *IVFSExplorercd) (*IVFSExplorercdResponse, error) {
	response := new(IVFSExplorercdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorercdUp(request *IVFSExplorercdUp) (*IVFSExplorercdUpResponse, error) {
	response := new(IVFSExplorercdUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerentryList(request *IVFSExplorerentryList) (*IVFSExplorerentryListResponse, error) {
	response := new(IVFSExplorerentryListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerexists(request *IVFSExplorerexists) (*IVFSExplorerexistsResponse, error) {
	response := new(IVFSExplorerexistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVFSExplorerremove(request *IVFSExplorerremove) (*IVFSExplorerremoveResponse, error) {
	response := new(IVFSExplorerremoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetPath(request *IAppliancegetPath) (*IAppliancegetPathResponse, error) {
	response := new(IAppliancegetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetDisks(request *IAppliancegetDisks) (*IAppliancegetDisksResponse, error) {
	response := new(IAppliancegetDisksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetVirtualSystemDescriptions(request *IAppliancegetVirtualSystemDescriptions) (*IAppliancegetVirtualSystemDescriptionsResponse, error) {
	response := new(IAppliancegetVirtualSystemDescriptionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetMachines(request *IAppliancegetMachines) (*IAppliancegetMachinesResponse, error) {
	response := new(IAppliancegetMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceread(request *IApplianceread) (*IAppliancereadResponse, error) {
	response := new(IAppliancereadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceinterpret(request *IApplianceinterpret) (*IApplianceinterpretResponse, error) {
	response := new(IApplianceinterpretResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IApplianceimportMachines(request *IApplianceimportMachines) (*IApplianceimportMachinesResponse, error) {
	response := new(IApplianceimportMachinesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancecreateVFSExplorer(request *IAppliancecreateVFSExplorer) (*IAppliancecreateVFSExplorerResponse, error) {
	response := new(IAppliancecreateVFSExplorerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancewrite(request *IAppliancewrite) (*IAppliancewriteResponse, error) {
	response := new(IAppliancewriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAppliancegetWarnings(request *IAppliancegetWarnings) (*IAppliancegetWarningsResponse, error) {
	response := new(IAppliancegetWarningsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetCount(request *IVirtualSystemDescriptiongetCount) (*IVirtualSystemDescriptiongetCountResponse, error) {
	response := new(IVirtualSystemDescriptiongetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescription(request *IVirtualSystemDescriptiongetDescription) (*IVirtualSystemDescriptiongetDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetDescriptionByType(request *IVirtualSystemDescriptiongetDescriptionByType) (*IVirtualSystemDescriptiongetDescriptionByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetDescriptionByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptiongetValuesByType(request *IVirtualSystemDescriptiongetValuesByType) (*IVirtualSystemDescriptiongetValuesByTypeResponse, error) {
	response := new(IVirtualSystemDescriptiongetValuesByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionsetFinalValues(request *IVirtualSystemDescriptionsetFinalValues) (*IVirtualSystemDescriptionsetFinalValuesResponse, error) {
	response := new(IVirtualSystemDescriptionsetFinalValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVirtualSystemDescriptionaddDescription(request *IVirtualSystemDescriptionaddDescription) (*IVirtualSystemDescriptionaddDescriptionResponse, error) {
	response := new(IVirtualSystemDescriptionaddDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeIn(request *IBIOSSettingsgetLogoFadeIn) (*IBIOSSettingsgetLogoFadeInResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeIn(request *IBIOSSettingssetLogoFadeIn) (*IBIOSSettingssetLogoFadeInResponse, error) {
	response := new(IBIOSSettingssetLogoFadeInResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoFadeOut(request *IBIOSSettingsgetLogoFadeOut) (*IBIOSSettingsgetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingsgetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoFadeOut(request *IBIOSSettingssetLogoFadeOut) (*IBIOSSettingssetLogoFadeOutResponse, error) {
	response := new(IBIOSSettingssetLogoFadeOutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoDisplayTime(request *IBIOSSettingsgetLogoDisplayTime) (*IBIOSSettingsgetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingsgetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoDisplayTime(request *IBIOSSettingssetLogoDisplayTime) (*IBIOSSettingssetLogoDisplayTimeResponse, error) {
	response := new(IBIOSSettingssetLogoDisplayTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetLogoImagePath(request *IBIOSSettingsgetLogoImagePath) (*IBIOSSettingsgetLogoImagePathResponse, error) {
	response := new(IBIOSSettingsgetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetLogoImagePath(request *IBIOSSettingssetLogoImagePath) (*IBIOSSettingssetLogoImagePathResponse, error) {
	response := new(IBIOSSettingssetLogoImagePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetBootMenuMode(request *IBIOSSettingsgetBootMenuMode) (*IBIOSSettingsgetBootMenuModeResponse, error) {
	response := new(IBIOSSettingsgetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetBootMenuMode(request *IBIOSSettingssetBootMenuMode) (*IBIOSSettingssetBootMenuModeResponse, error) {
	response := new(IBIOSSettingssetBootMenuModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetACPIEnabled(request *IBIOSSettingsgetACPIEnabled) (*IBIOSSettingsgetACPIEnabledResponse, error) {
	response := new(IBIOSSettingsgetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetACPIEnabled(request *IBIOSSettingssetACPIEnabled) (*IBIOSSettingssetACPIEnabledResponse, error) {
	response := new(IBIOSSettingssetACPIEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetIOAPICEnabled(request *IBIOSSettingsgetIOAPICEnabled) (*IBIOSSettingsgetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingsgetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetIOAPICEnabled(request *IBIOSSettingssetIOAPICEnabled) (*IBIOSSettingssetIOAPICEnabledResponse, error) {
	response := new(IBIOSSettingssetIOAPICEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetTimeOffset(request *IBIOSSettingsgetTimeOffset) (*IBIOSSettingsgetTimeOffsetResponse, error) {
	response := new(IBIOSSettingsgetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetTimeOffset(request *IBIOSSettingssetTimeOffset) (*IBIOSSettingssetTimeOffsetResponse, error) {
	response := new(IBIOSSettingssetTimeOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetPXEDebugEnabled(request *IBIOSSettingsgetPXEDebugEnabled) (*IBIOSSettingsgetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingsgetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingssetPXEDebugEnabled(request *IBIOSSettingssetPXEDebugEnabled) (*IBIOSSettingssetPXEDebugEnabledResponse, error) {
	response := new(IBIOSSettingssetPXEDebugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBIOSSettingsgetNonVolatileStorageFile(request *IBIOSSettingsgetNonVolatileStorageFile) (*IBIOSSettingsgetNonVolatileStorageFileResponse, error) {
	response := new(IBIOSSettingsgetNonVolatileStorageFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetBus(request *IPCIAddressgetBus) (*IPCIAddressgetBusResponse, error) {
	response := new(IPCIAddressgetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetBus(request *IPCIAddresssetBus) (*IPCIAddresssetBusResponse, error) {
	response := new(IPCIAddresssetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetDevice(request *IPCIAddressgetDevice) (*IPCIAddressgetDeviceResponse, error) {
	response := new(IPCIAddressgetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetDevice(request *IPCIAddresssetDevice) (*IPCIAddresssetDeviceResponse, error) {
	response := new(IPCIAddresssetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressgetDevFunction(request *IPCIAddressgetDevFunction) (*IPCIAddressgetDevFunctionResponse, error) {
	response := new(IPCIAddressgetDevFunctionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddresssetDevFunction(request *IPCIAddresssetDevFunction) (*IPCIAddresssetDevFunctionResponse, error) {
	response := new(IPCIAddresssetDevFunctionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressasLong(request *IPCIAddressasLong) (*IPCIAddressasLongResponse, error) {
	response := new(IPCIAddressasLongResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPCIAddressfromLong(request *IPCIAddressfromLong) (*IPCIAddressfromLongResponse, error) {
	response := new(IPCIAddressfromLongResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParent(request *IMachinegetParent) (*IMachinegetParentResponse, error) {
	response := new(IMachinegetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIcon(request *IMachinegetIcon) (*IMachinegetIconResponse, error) {
	response := new(IMachinegetIconResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIcon(request *IMachinesetIcon) (*IMachinesetIconResponse, error) {
	response := new(IMachinesetIconResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessible(request *IMachinegetAccessible) (*IMachinegetAccessibleResponse, error) {
	response := new(IMachinegetAccessibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccessError(request *IMachinegetAccessError) (*IMachinegetAccessErrorResponse, error) {
	response := new(IMachinegetAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetName(request *IMachinegetName) (*IMachinegetNameResponse, error) {
	response := new(IMachinegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetName(request *IMachinesetName) (*IMachinesetNameResponse, error) {
	response := new(IMachinesetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDescription(request *IMachinegetDescription) (*IMachinegetDescriptionResponse, error) {
	response := new(IMachinegetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDescription(request *IMachinesetDescription) (*IMachinesetDescriptionResponse, error) {
	response := new(IMachinesetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetId(request *IMachinegetId) (*IMachinegetIdResponse, error) {
	response := new(IMachinegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGroups(request *IMachinegetGroups) (*IMachinegetGroupsResponse, error) {
	response := new(IMachinegetGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGroups(request *IMachinesetGroups) (*IMachinesetGroupsResponse, error) {
	response := new(IMachinesetGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetOSTypeId(request *IMachinegetOSTypeId) (*IMachinegetOSTypeIdResponse, error) {
	response := new(IMachinegetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetOSTypeId(request *IMachinesetOSTypeId) (*IMachinesetOSTypeIdResponse, error) {
	response := new(IMachinesetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareVersion(request *IMachinegetHardwareVersion) (*IMachinegetHardwareVersionResponse, error) {
	response := new(IMachinegetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareVersion(request *IMachinesetHardwareVersion) (*IMachinesetHardwareVersionResponse, error) {
	response := new(IMachinesetHardwareVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHardwareUUID(request *IMachinegetHardwareUUID) (*IMachinegetHardwareUUIDResponse, error) {
	response := new(IMachinegetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHardwareUUID(request *IMachinesetHardwareUUID) (*IMachinesetHardwareUUIDResponse, error) {
	response := new(IMachinesetHardwareUUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUCount(request *IMachinegetCPUCount) (*IMachinegetCPUCountResponse, error) {
	response := new(IMachinegetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUCount(request *IMachinesetCPUCount) (*IMachinesetCPUCountResponse, error) {
	response := new(IMachinesetCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUHotPlugEnabled(request *IMachinegetCPUHotPlugEnabled) (*IMachinegetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinegetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUHotPlugEnabled(request *IMachinesetCPUHotPlugEnabled) (*IMachinesetCPUHotPlugEnabledResponse, error) {
	response := new(IMachinesetCPUHotPlugEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUExecutionCap(request *IMachinegetCPUExecutionCap) (*IMachinegetCPUExecutionCapResponse, error) {
	response := new(IMachinegetCPUExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUExecutionCap(request *IMachinesetCPUExecutionCap) (*IMachinesetCPUExecutionCapResponse, error) {
	response := new(IMachinesetCPUExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemorySize(request *IMachinegetMemorySize) (*IMachinegetMemorySizeResponse, error) {
	response := new(IMachinegetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemorySize(request *IMachinesetMemorySize) (*IMachinesetMemorySizeResponse, error) {
	response := new(IMachinesetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMemoryBalloonSize(request *IMachinegetMemoryBalloonSize) (*IMachinegetMemoryBalloonSizeResponse, error) {
	response := new(IMachinegetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMemoryBalloonSize(request *IMachinesetMemoryBalloonSize) (*IMachinesetMemoryBalloonSizeResponse, error) {
	response := new(IMachinesetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPageFusionEnabled(request *IMachinegetPageFusionEnabled) (*IMachinegetPageFusionEnabledResponse, error) {
	response := new(IMachinegetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPageFusionEnabled(request *IMachinesetPageFusionEnabled) (*IMachinesetPageFusionEnabledResponse, error) {
	response := new(IMachinesetPageFusionEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGraphicsControllerType(request *IMachinegetGraphicsControllerType) (*IMachinegetGraphicsControllerTypeResponse, error) {
	response := new(IMachinegetGraphicsControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGraphicsControllerType(request *IMachinesetGraphicsControllerType) (*IMachinesetGraphicsControllerTypeResponse, error) {
	response := new(IMachinesetGraphicsControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRAMSize(request *IMachinegetVRAMSize) (*IMachinegetVRAMSizeResponse, error) {
	response := new(IMachinegetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVRAMSize(request *IMachinesetVRAMSize) (*IMachinesetVRAMSizeResponse, error) {
	response := new(IMachinesetVRAMSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate3DEnabled(request *IMachinegetAccelerate3DEnabled) (*IMachinegetAccelerate3DEnabledResponse, error) {
	response := new(IMachinegetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate3DEnabled(request *IMachinesetAccelerate3DEnabled) (*IMachinesetAccelerate3DEnabledResponse, error) {
	response := new(IMachinesetAccelerate3DEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAccelerate2DVideoEnabled(request *IMachinegetAccelerate2DVideoEnabled) (*IMachinegetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinegetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAccelerate2DVideoEnabled(request *IMachinesetAccelerate2DVideoEnabled) (*IMachinesetAccelerate2DVideoEnabledResponse, error) {
	response := new(IMachinesetAccelerate2DVideoEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMonitorCount(request *IMachinegetMonitorCount) (*IMachinegetMonitorCountResponse, error) {
	response := new(IMachinegetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetMonitorCount(request *IMachinesetMonitorCount) (*IMachinesetMonitorCountResponse, error) {
	response := new(IMachinesetMonitorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureEnabled(request *IMachinegetVideoCaptureEnabled) (*IMachinegetVideoCaptureEnabledResponse, error) {
	response := new(IMachinegetVideoCaptureEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureEnabled(request *IMachinesetVideoCaptureEnabled) (*IMachinesetVideoCaptureEnabledResponse, error) {
	response := new(IMachinesetVideoCaptureEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureScreens(request *IMachinegetVideoCaptureScreens) (*IMachinegetVideoCaptureScreensResponse, error) {
	response := new(IMachinegetVideoCaptureScreensResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureScreens(request *IMachinesetVideoCaptureScreens) (*IMachinesetVideoCaptureScreensResponse, error) {
	response := new(IMachinesetVideoCaptureScreensResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureFile(request *IMachinegetVideoCaptureFile) (*IMachinegetVideoCaptureFileResponse, error) {
	response := new(IMachinegetVideoCaptureFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureFile(request *IMachinesetVideoCaptureFile) (*IMachinesetVideoCaptureFileResponse, error) {
	response := new(IMachinesetVideoCaptureFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureWidth(request *IMachinegetVideoCaptureWidth) (*IMachinegetVideoCaptureWidthResponse, error) {
	response := new(IMachinegetVideoCaptureWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureWidth(request *IMachinesetVideoCaptureWidth) (*IMachinesetVideoCaptureWidthResponse, error) {
	response := new(IMachinesetVideoCaptureWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureHeight(request *IMachinegetVideoCaptureHeight) (*IMachinegetVideoCaptureHeightResponse, error) {
	response := new(IMachinegetVideoCaptureHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureHeight(request *IMachinesetVideoCaptureHeight) (*IMachinesetVideoCaptureHeightResponse, error) {
	response := new(IMachinesetVideoCaptureHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureRate(request *IMachinegetVideoCaptureRate) (*IMachinegetVideoCaptureRateResponse, error) {
	response := new(IMachinegetVideoCaptureRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureRate(request *IMachinesetVideoCaptureRate) (*IMachinesetVideoCaptureRateResponse, error) {
	response := new(IMachinesetVideoCaptureRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVideoCaptureFPS(request *IMachinegetVideoCaptureFPS) (*IMachinegetVideoCaptureFPSResponse, error) {
	response := new(IMachinegetVideoCaptureFPSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetVideoCaptureFPS(request *IMachinesetVideoCaptureFPS) (*IMachinesetVideoCaptureFPSResponse, error) {
	response := new(IMachinesetVideoCaptureFPSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBIOSSettings(request *IMachinegetBIOSSettings) (*IMachinegetBIOSSettingsResponse, error) {
	response := new(IMachinegetBIOSSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFirmwareType(request *IMachinegetFirmwareType) (*IMachinegetFirmwareTypeResponse, error) {
	response := new(IMachinegetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFirmwareType(request *IMachinesetFirmwareType) (*IMachinesetFirmwareTypeResponse, error) {
	response := new(IMachinesetFirmwareTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPointingHIDType(request *IMachinegetPointingHIDType) (*IMachinegetPointingHIDTypeResponse, error) {
	response := new(IMachinegetPointingHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetPointingHIDType(request *IMachinesetPointingHIDType) (*IMachinesetPointingHIDTypeResponse, error) {
	response := new(IMachinesetPointingHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetKeyboardHIDType(request *IMachinegetKeyboardHIDType) (*IMachinegetKeyboardHIDTypeResponse, error) {
	response := new(IMachinegetKeyboardHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetKeyboardHIDType(request *IMachinesetKeyboardHIDType) (*IMachinesetKeyboardHIDTypeResponse, error) {
	response := new(IMachinesetKeyboardHIDTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHPETEnabled(request *IMachinegetHPETEnabled) (*IMachinegetHPETEnabledResponse, error) {
	response := new(IMachinegetHPETEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHPETEnabled(request *IMachinesetHPETEnabled) (*IMachinesetHPETEnabledResponse, error) {
	response := new(IMachinesetHPETEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetChipsetType(request *IMachinegetChipsetType) (*IMachinegetChipsetTypeResponse, error) {
	response := new(IMachinegetChipsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetChipsetType(request *IMachinesetChipsetType) (*IMachinesetChipsetTypeResponse, error) {
	response := new(IMachinesetChipsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotFolder(request *IMachinegetSnapshotFolder) (*IMachinegetSnapshotFolderResponse, error) {
	response := new(IMachinegetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetSnapshotFolder(request *IMachinesetSnapshotFolder) (*IMachinesetSnapshotFolderResponse, error) {
	response := new(IMachinesetSnapshotFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetVRDEServer(request *IMachinegetVRDEServer) (*IMachinegetVRDEServerResponse, error) {
	response := new(IMachinegetVRDEServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetEmulatedUSBCardReaderEnabled(request *IMachinegetEmulatedUSBCardReaderEnabled) (*IMachinegetEmulatedUSBCardReaderEnabledResponse, error) {
	response := new(IMachinegetEmulatedUSBCardReaderEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetEmulatedUSBCardReaderEnabled(request *IMachinesetEmulatedUSBCardReaderEnabled) (*IMachinesetEmulatedUSBCardReaderEnabledResponse, error) {
	response := new(IMachinesetEmulatedUSBCardReaderEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachments(request *IMachinegetMediumAttachments) (*IMachinegetMediumAttachmentsResponse, error) {
	response := new(IMachinegetMediumAttachmentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllers(request *IMachinegetUSBControllers) (*IMachinegetUSBControllersResponse, error) {
	response := new(IMachinegetUSBControllersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBDeviceFilters(request *IMachinegetUSBDeviceFilters) (*IMachinegetUSBDeviceFiltersResponse, error) {
	response := new(IMachinegetUSBDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAudioAdapter(request *IMachinegetAudioAdapter) (*IMachinegetAudioAdapterResponse, error) {
	response := new(IMachinegetAudioAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllers(request *IMachinegetStorageControllers) (*IMachinegetStorageControllersResponse, error) {
	response := new(IMachinegetStorageControllersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsFilePath(request *IMachinegetSettingsFilePath) (*IMachinegetSettingsFilePathResponse, error) {
	response := new(IMachinegetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSettingsModified(request *IMachinegetSettingsModified) (*IMachinegetSettingsModifiedResponse, error) {
	response := new(IMachinegetSettingsModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionState(request *IMachinegetSessionState) (*IMachinegetSessionStateResponse, error) {
	response := new(IMachinegetSessionStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionType(request *IMachinegetSessionType) (*IMachinegetSessionTypeResponse, error) {
	response := new(IMachinegetSessionTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSessionPID(request *IMachinegetSessionPID) (*IMachinegetSessionPIDResponse, error) {
	response := new(IMachinegetSessionPIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetState(request *IMachinegetState) (*IMachinegetStateResponse, error) {
	response := new(IMachinegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLastStateChange(request *IMachinegetLastStateChange) (*IMachinegetLastStateChangeResponse, error) {
	response := new(IMachinegetLastStateChangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStateFilePath(request *IMachinegetStateFilePath) (*IMachinegetStateFilePathResponse, error) {
	response := new(IMachinegetStateFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetLogFolder(request *IMachinegetLogFolder) (*IMachinegetLogFolderResponse, error) {
	response := new(IMachinegetLogFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentSnapshot(request *IMachinegetCurrentSnapshot) (*IMachinegetCurrentSnapshotResponse, error) {
	response := new(IMachinegetCurrentSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSnapshotCount(request *IMachinegetSnapshotCount) (*IMachinegetSnapshotCountResponse, error) {
	response := new(IMachinegetSnapshotCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCurrentStateModified(request *IMachinegetCurrentStateModified) (*IMachinegetCurrentStateModifiedResponse, error) {
	response := new(IMachinegetCurrentStateModifiedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSharedFolders(request *IMachinegetSharedFolders) (*IMachinegetSharedFoldersResponse, error) {
	response := new(IMachinegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetClipboardMode(request *IMachinegetClipboardMode) (*IMachinegetClipboardModeResponse, error) {
	response := new(IMachinegetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetClipboardMode(request *IMachinesetClipboardMode) (*IMachinesetClipboardModeResponse, error) {
	response := new(IMachinesetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDragAndDropMode(request *IMachinegetDragAndDropMode) (*IMachinegetDragAndDropModeResponse, error) {
	response := new(IMachinegetDragAndDropModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDragAndDropMode(request *IMachinesetDragAndDropMode) (*IMachinesetDragAndDropModeResponse, error) {
	response := new(IMachinesetDragAndDropModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyNotificationPatterns(request *IMachinegetGuestPropertyNotificationPatterns) (*IMachinegetGuestPropertyNotificationPatternsResponse, error) {
	response := new(IMachinegetGuestPropertyNotificationPatternsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestPropertyNotificationPatterns(request *IMachinesetGuestPropertyNotificationPatterns) (*IMachinesetGuestPropertyNotificationPatternsResponse, error) {
	response := new(IMachinesetGuestPropertyNotificationPatternsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterEnabled(request *IMachinegetTeleporterEnabled) (*IMachinegetTeleporterEnabledResponse, error) {
	response := new(IMachinegetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterEnabled(request *IMachinesetTeleporterEnabled) (*IMachinesetTeleporterEnabledResponse, error) {
	response := new(IMachinesetTeleporterEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPort(request *IMachinegetTeleporterPort) (*IMachinegetTeleporterPortResponse, error) {
	response := new(IMachinegetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPort(request *IMachinesetTeleporterPort) (*IMachinesetTeleporterPortResponse, error) {
	response := new(IMachinesetTeleporterPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterAddress(request *IMachinegetTeleporterAddress) (*IMachinegetTeleporterAddressResponse, error) {
	response := new(IMachinegetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterAddress(request *IMachinesetTeleporterAddress) (*IMachinesetTeleporterAddressResponse, error) {
	response := new(IMachinesetTeleporterAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTeleporterPassword(request *IMachinegetTeleporterPassword) (*IMachinegetTeleporterPasswordResponse, error) {
	response := new(IMachinegetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTeleporterPassword(request *IMachinesetTeleporterPassword) (*IMachinesetTeleporterPasswordResponse, error) {
	response := new(IMachinesetTeleporterPasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceState(request *IMachinegetFaultToleranceState) (*IMachinegetFaultToleranceStateResponse, error) {
	response := new(IMachinegetFaultToleranceStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceState(request *IMachinesetFaultToleranceState) (*IMachinesetFaultToleranceStateResponse, error) {
	response := new(IMachinesetFaultToleranceStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultTolerancePort(request *IMachinegetFaultTolerancePort) (*IMachinegetFaultTolerancePortResponse, error) {
	response := new(IMachinegetFaultTolerancePortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultTolerancePort(request *IMachinesetFaultTolerancePort) (*IMachinesetFaultTolerancePortResponse, error) {
	response := new(IMachinesetFaultTolerancePortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceAddress(request *IMachinegetFaultToleranceAddress) (*IMachinegetFaultToleranceAddressResponse, error) {
	response := new(IMachinegetFaultToleranceAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceAddress(request *IMachinesetFaultToleranceAddress) (*IMachinesetFaultToleranceAddressResponse, error) {
	response := new(IMachinesetFaultToleranceAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultTolerancePassword(request *IMachinegetFaultTolerancePassword) (*IMachinegetFaultTolerancePasswordResponse, error) {
	response := new(IMachinegetFaultTolerancePasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultTolerancePassword(request *IMachinesetFaultTolerancePassword) (*IMachinesetFaultTolerancePasswordResponse, error) {
	response := new(IMachinesetFaultTolerancePasswordResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetFaultToleranceSyncInterval(request *IMachinegetFaultToleranceSyncInterval) (*IMachinegetFaultToleranceSyncIntervalResponse, error) {
	response := new(IMachinegetFaultToleranceSyncIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetFaultToleranceSyncInterval(request *IMachinesetFaultToleranceSyncInterval) (*IMachinesetFaultToleranceSyncIntervalResponse, error) {
	response := new(IMachinesetFaultToleranceSyncIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetRTCUseUTC(request *IMachinegetRTCUseUTC) (*IMachinegetRTCUseUTCResponse, error) {
	response := new(IMachinegetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetRTCUseUTC(request *IMachinesetRTCUseUTC) (*IMachinesetRTCUseUTCResponse, error) {
	response := new(IMachinesetRTCUseUTCResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIOCacheEnabled(request *IMachinegetIOCacheEnabled) (*IMachinegetIOCacheEnabledResponse, error) {
	response := new(IMachinegetIOCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIOCacheEnabled(request *IMachinesetIOCacheEnabled) (*IMachinesetIOCacheEnabledResponse, error) {
	response := new(IMachinesetIOCacheEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetIOCacheSize(request *IMachinegetIOCacheSize) (*IMachinegetIOCacheSizeResponse, error) {
	response := new(IMachinegetIOCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetIOCacheSize(request *IMachinesetIOCacheSize) (*IMachinesetIOCacheSizeResponse, error) {
	response := new(IMachinesetIOCacheSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetPCIDeviceAssignments(request *IMachinegetPCIDeviceAssignments) (*IMachinegetPCIDeviceAssignmentsResponse, error) {
	response := new(IMachinegetPCIDeviceAssignmentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBandwidthControl(request *IMachinegetBandwidthControl) (*IMachinegetBandwidthControlResponse, error) {
	response := new(IMachinegetBandwidthControlResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTracingEnabled(request *IMachinegetTracingEnabled) (*IMachinegetTracingEnabledResponse, error) {
	response := new(IMachinegetTracingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTracingEnabled(request *IMachinesetTracingEnabled) (*IMachinesetTracingEnabledResponse, error) {
	response := new(IMachinesetTracingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetTracingConfig(request *IMachinegetTracingConfig) (*IMachinegetTracingConfigResponse, error) {
	response := new(IMachinegetTracingConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetTracingConfig(request *IMachinesetTracingConfig) (*IMachinesetTracingConfigResponse, error) {
	response := new(IMachinesetTracingConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAllowTracingToAccessVM(request *IMachinegetAllowTracingToAccessVM) (*IMachinegetAllowTracingToAccessVMResponse, error) {
	response := new(IMachinegetAllowTracingToAccessVMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAllowTracingToAccessVM(request *IMachinesetAllowTracingToAccessVM) (*IMachinesetAllowTracingToAccessVMResponse, error) {
	response := new(IMachinesetAllowTracingToAccessVMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostartEnabled(request *IMachinegetAutostartEnabled) (*IMachinegetAutostartEnabledResponse, error) {
	response := new(IMachinegetAutostartEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostartEnabled(request *IMachinesetAutostartEnabled) (*IMachinesetAutostartEnabledResponse, error) {
	response := new(IMachinesetAutostartEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostartDelay(request *IMachinegetAutostartDelay) (*IMachinegetAutostartDelayResponse, error) {
	response := new(IMachinegetAutostartDelayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostartDelay(request *IMachinesetAutostartDelay) (*IMachinesetAutostartDelayResponse, error) {
	response := new(IMachinesetAutostartDelayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetAutostopType(request *IMachinegetAutostopType) (*IMachinegetAutostopTypeResponse, error) {
	response := new(IMachinegetAutostopTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutostopType(request *IMachinesetAutostopType) (*IMachinesetAutostopTypeResponse, error) {
	response := new(IMachinesetAutostopTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetDefaultFrontend(request *IMachinegetDefaultFrontend) (*IMachinegetDefaultFrontendResponse, error) {
	response := new(IMachinegetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetDefaultFrontend(request *IMachinesetDefaultFrontend) (*IMachinesetDefaultFrontendResponse, error) {
	response := new(IMachinesetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBProxyAvailable(request *IMachinegetUSBProxyAvailable) (*IMachinegetUSBProxyAvailableResponse, error) {
	response := new(IMachinegetUSBProxyAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinelockMachine(request *IMachinelockMachine) (*IMachinelockMachineResponse, error) {
	response := new(IMachinelockMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinelaunchVMProcess(request *IMachinelaunchVMProcess) (*IMachinelaunchVMProcessResponse, error) {
	response := new(IMachinelaunchVMProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetBootOrder(request *IMachinesetBootOrder) (*IMachinesetBootOrderResponse, error) {
	response := new(IMachinesetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetBootOrder(request *IMachinegetBootOrder) (*IMachinegetBootOrderResponse, error) {
	response := new(IMachinegetBootOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachDevice(request *IMachineattachDevice) (*IMachineattachDeviceResponse, error) {
	response := new(IMachineattachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachDeviceWithoutMedium(request *IMachineattachDeviceWithoutMedium) (*IMachineattachDeviceWithoutMediumResponse, error) {
	response := new(IMachineattachDeviceWithoutMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedetachDevice(request *IMachinedetachDevice) (*IMachinedetachDeviceResponse, error) {
	response := new(IMachinedetachDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinepassthroughDevice(request *IMachinepassthroughDevice) (*IMachinepassthroughDeviceResponse, error) {
	response := new(IMachinepassthroughDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinetemporaryEjectDevice(request *IMachinetemporaryEjectDevice) (*IMachinetemporaryEjectDeviceResponse, error) {
	response := new(IMachinetemporaryEjectDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinenonRotationalDevice(request *IMachinenonRotationalDevice) (*IMachinenonRotationalDeviceResponse, error) {
	response := new(IMachinenonRotationalDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetAutoDiscardForDevice(request *IMachinesetAutoDiscardForDevice) (*IMachinesetAutoDiscardForDeviceResponse, error) {
	response := new(IMachinesetAutoDiscardForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHotPluggableForDevice(request *IMachinesetHotPluggableForDevice) (*IMachinesetHotPluggableForDeviceResponse, error) {
	response := new(IMachinesetHotPluggableForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetBandwidthGroupForDevice(request *IMachinesetBandwidthGroupForDevice) (*IMachinesetBandwidthGroupForDeviceResponse, error) {
	response := new(IMachinesetBandwidthGroupForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetNoBandwidthGroupForDevice(request *IMachinesetNoBandwidthGroupForDevice) (*IMachinesetNoBandwidthGroupForDeviceResponse, error) {
	response := new(IMachinesetNoBandwidthGroupForDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineunmountMedium(request *IMachineunmountMedium) (*IMachineunmountMediumResponse, error) {
	response := new(IMachineunmountMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinemountMedium(request *IMachinemountMedium) (*IMachinemountMediumResponse, error) {
	response := new(IMachinemountMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMedium(request *IMachinegetMedium) (*IMachinegetMediumResponse, error) {
	response := new(IMachinegetMediumResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

type Returnval *IMediumAttachment

func (service *VboxPortType) IMachinegetMediumAttachmentsOfController(request *IMachinegetMediumAttachmentsOfController) (*IMachinegetMediumAttachmentsOfControllerResponse, error) {
	response := new(IMachinegetMediumAttachmentsOfControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetMediumAttachment(request *IMachinegetMediumAttachment) (*IMachinegetMediumAttachmentResponse, error) {
	response := new(IMachinegetMediumAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineattachHostPCIDevice(request *IMachineattachHostPCIDevice) (*IMachineattachHostPCIDeviceResponse, error) {
	response := new(IMachineattachHostPCIDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedetachHostPCIDevice(request *IMachinedetachHostPCIDevice) (*IMachinedetachHostPCIDeviceResponse, error) {
	response := new(IMachinedetachHostPCIDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetNetworkAdapter(request *IMachinegetNetworkAdapter) (*IMachinegetNetworkAdapterResponse, error) {
	response := new(IMachinegetNetworkAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineaddStorageController(request *IMachineaddStorageController) (*IMachineaddStorageControllerResponse, error) {
	response := new(IMachineaddStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByName(request *IMachinegetStorageControllerByName) (*IMachinegetStorageControllerByNameResponse, error) {
	response := new(IMachinegetStorageControllerByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetStorageControllerByInstance(request *IMachinegetStorageControllerByInstance) (*IMachinegetStorageControllerByInstanceResponse, error) {
	response := new(IMachinegetStorageControllerByInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveStorageController(request *IMachineremoveStorageController) (*IMachineremoveStorageControllerResponse, error) {
	response := new(IMachineremoveStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetStorageControllerBootable(request *IMachinesetStorageControllerBootable) (*IMachinesetStorageControllerBootableResponse, error) {
	response := new(IMachinesetStorageControllerBootableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineaddUSBController(request *IMachineaddUSBController) (*IMachineaddUSBControllerResponse, error) {
	response := new(IMachineaddUSBControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveUSBController(request *IMachineremoveUSBController) (*IMachineremoveUSBControllerResponse, error) {
	response := new(IMachineremoveUSBControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllerByName(request *IMachinegetUSBControllerByName) (*IMachinegetUSBControllerByNameResponse, error) {
	response := new(IMachinegetUSBControllerByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetUSBControllerCountByType(request *IMachinegetUSBControllerCountByType) (*IMachinegetUSBControllerCountByTypeResponse, error) {
	response := new(IMachinegetUSBControllerCountByTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetSerialPort(request *IMachinegetSerialPort) (*IMachinegetSerialPortResponse, error) {
	response := new(IMachinegetSerialPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetParallelPort(request *IMachinegetParallelPort) (*IMachinegetParallelPortResponse, error) {
	response := new(IMachinegetParallelPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraDataKeys(request *IMachinegetExtraDataKeys) (*IMachinegetExtraDataKeysResponse, error) {
	response := new(IMachinegetExtraDataKeysResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetExtraData(request *IMachinegetExtraData) (*IMachinegetExtraDataResponse, error) {
	response := new(IMachinegetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetExtraData(request *IMachinesetExtraData) (*IMachinesetExtraDataResponse, error) {
	response := new(IMachinesetExtraDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUProperty(request *IMachinegetCPUProperty) (*IMachinegetCPUPropertyResponse, error) {
	response := new(IMachinegetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUProperty(request *IMachinesetCPUProperty) (*IMachinesetCPUPropertyResponse, error) {
	response := new(IMachinesetCPUPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUIDLeaf(request *IMachinegetCPUIDLeaf) (*IMachinegetCPUIDLeafResponse, error) {
	response := new(IMachinegetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetCPUIDLeaf(request *IMachinesetCPUIDLeaf) (*IMachinesetCPUIDLeafResponse, error) {
	response := new(IMachinesetCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveCPUIDLeaf(request *IMachineremoveCPUIDLeaf) (*IMachineremoveCPUIDLeafResponse, error) {
	response := new(IMachineremoveCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveAllCPUIDLeaves(request *IMachineremoveAllCPUIDLeaves) (*IMachineremoveAllCPUIDLeavesResponse, error) {
	response := new(IMachineremoveAllCPUIDLeavesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetHWVirtExProperty(request *IMachinegetHWVirtExProperty) (*IMachinegetHWVirtExPropertyResponse, error) {
	response := new(IMachinegetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetHWVirtExProperty(request *IMachinesetHWVirtExProperty) (*IMachinesetHWVirtExPropertyResponse, error) {
	response := new(IMachinesetHWVirtExPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetSettingsFilePath(request *IMachinesetSettingsFilePath) (*IMachinesetSettingsFilePathResponse, error) {
	response := new(IMachinesetSettingsFilePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesaveSettings(request *IMachinesaveSettings) (*IMachinesaveSettingsResponse, error) {
	response := new(IMachinesaveSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinediscardSettings(request *IMachinediscardSettings) (*IMachinediscardSettingsResponse, error) {
	response := new(IMachinediscardSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineunregister(request *IMachineunregister) (*IMachineunregisterResponse, error) {
	response := new(IMachineunregisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteConfig(request *IMachinedeleteConfig) (*IMachinedeleteConfigResponse, error) {
	response := new(IMachinedeleteConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineexportTo(request *IMachineexportTo) (*IMachineexportToResponse, error) {
	response := new(IMachineexportToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinefindSnapshot(request *IMachinefindSnapshot) (*IMachinefindSnapshotResponse, error) {
	response := new(IMachinefindSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecreateSharedFolder(request *IMachinecreateSharedFolder) (*IMachinecreateSharedFolderResponse, error) {
	response := new(IMachinecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineremoveSharedFolder(request *IMachineremoveSharedFolder) (*IMachineremoveSharedFolderResponse, error) {
	response := new(IMachineremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecanShowConsoleWindow(request *IMachinecanShowConsoleWindow) (*IMachinecanShowConsoleWindowResponse, error) {
	response := new(IMachinecanShowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineshowConsoleWindow(request *IMachineshowConsoleWindow) (*IMachineshowConsoleWindowResponse, error) {
	response := new(IMachineshowConsoleWindowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestProperty(request *IMachinegetGuestProperty) (*IMachinegetGuestPropertyResponse, error) {
	response := new(IMachinegetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyValue(request *IMachinegetGuestPropertyValue) (*IMachinegetGuestPropertyValueResponse, error) {
	response := new(IMachinegetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetGuestPropertyTimestamp(request *IMachinegetGuestPropertyTimestamp) (*IMachinegetGuestPropertyTimestampResponse, error) {
	response := new(IMachinegetGuestPropertyTimestampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestProperty(request *IMachinesetGuestProperty) (*IMachinesetGuestPropertyResponse, error) {
	response := new(IMachinesetGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinesetGuestPropertyValue(request *IMachinesetGuestPropertyValue) (*IMachinesetGuestPropertyValueResponse, error) {
	response := new(IMachinesetGuestPropertyValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinedeleteGuestProperty(request *IMachinedeleteGuestProperty) (*IMachinedeleteGuestPropertyResponse, error) {
	response := new(IMachinedeleteGuestPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineenumerateGuestProperties(request *IMachineenumerateGuestProperties) (*IMachineenumerateGuestPropertiesResponse, error) {
	response := new(IMachineenumerateGuestPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedGuestScreenInfo(request *IMachinequerySavedGuestScreenInfo) (*IMachinequerySavedGuestScreenInfoResponse, error) {
	response := new(IMachinequerySavedGuestScreenInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedThumbnailSize(request *IMachinequerySavedThumbnailSize) (*IMachinequerySavedThumbnailSizeResponse, error) {
	response := new(IMachinequerySavedThumbnailSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedThumbnailToArray(request *IMachinereadSavedThumbnailToArray) (*IMachinereadSavedThumbnailToArrayResponse, error) {
	response := new(IMachinereadSavedThumbnailToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedThumbnailPNGToArray(request *IMachinereadSavedThumbnailPNGToArray) (*IMachinereadSavedThumbnailPNGToArrayResponse, error) {
	response := new(IMachinereadSavedThumbnailPNGToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequerySavedScreenshotPNGSize(request *IMachinequerySavedScreenshotPNGSize) (*IMachinequerySavedScreenshotPNGSizeResponse, error) {
	response := new(IMachinequerySavedScreenshotPNGSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadSavedScreenshotPNGToArray(request *IMachinereadSavedScreenshotPNGToArray) (*IMachinereadSavedScreenshotPNGToArrayResponse, error) {
	response := new(IMachinereadSavedScreenshotPNGToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotPlugCPU(request *IMachinehotPlugCPU) (*IMachinehotPlugCPUResponse, error) {
	response := new(IMachinehotPlugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinehotUnplugCPU(request *IMachinehotUnplugCPU) (*IMachinehotUnplugCPUResponse, error) {
	response := new(IMachinehotUnplugCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinegetCPUStatus(request *IMachinegetCPUStatus) (*IMachinegetCPUStatusResponse, error) {
	response := new(IMachinegetCPUStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinequeryLogFilename(request *IMachinequeryLogFilename) (*IMachinequeryLogFilenameResponse, error) {
	response := new(IMachinequeryLogFilenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinereadLog(request *IMachinereadLog) (*IMachinereadLogResponse, error) {
	response := new(IMachinereadLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachinecloneTo(request *IMachinecloneTo) (*IMachinecloneToResponse, error) {
	response := new(IMachinecloneToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBgetWebcams(request *IEmulatedUSBgetWebcams) (*IEmulatedUSBgetWebcamsResponse, error) {
	response := new(IEmulatedUSBgetWebcamsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBwebcamAttach(request *IEmulatedUSBwebcamAttach) (*IEmulatedUSBwebcamAttachResponse, error) {
	response := new(IEmulatedUSBwebcamAttachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEmulatedUSBwebcamDetach(request *IEmulatedUSBwebcamDetach) (*IEmulatedUSBwebcamDetachResponse, error) {
	response := new(IEmulatedUSBwebcamDetachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMachine(request *IConsolegetMachine) (*IConsolegetMachineResponse, error) {
	response := new(IConsolegetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetState(request *IConsolegetState) (*IConsolegetStateResponse, error) {
	response := new(IConsolegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuest(request *IConsolegetGuest) (*IConsolegetGuestResponse, error) {
	response := new(IConsolegetGuestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetKeyboard(request *IConsolegetKeyboard) (*IConsolegetKeyboardResponse, error) {
	response := new(IConsolegetKeyboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetMouse(request *IConsolegetMouse) (*IConsolegetMouseResponse, error) {
	response := new(IConsolegetMouseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDisplay(request *IConsolegetDisplay) (*IConsolegetDisplayResponse, error) {
	response := new(IConsolegetDisplayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDebugger(request *IConsolegetDebugger) (*IConsolegetDebuggerResponse, error) {
	response := new(IConsolegetDebuggerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetUSBDevices(request *IConsolegetUSBDevices) (*IConsolegetUSBDevicesResponse, error) {
	response := new(IConsolegetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetRemoteUSBDevices(request *IConsolegetRemoteUSBDevices) (*IConsolegetRemoteUSBDevicesResponse, error) {
	response := new(IConsolegetRemoteUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetSharedFolders(request *IConsolegetSharedFolders) (*IConsolegetSharedFoldersResponse, error) {
	response := new(IConsolegetSharedFoldersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetVRDEServerInfo(request *IConsolegetVRDEServerInfo) (*IConsolegetVRDEServerInfoResponse, error) {
	response := new(IConsolegetVRDEServerInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetEventSource(request *IConsolegetEventSource) (*IConsolegetEventSourceResponse, error) {
	response := new(IConsolegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetAttachedPCIDevices(request *IConsolegetAttachedPCIDevices) (*IConsolegetAttachedPCIDevicesResponse, error) {
	response := new(IConsolegetAttachedPCIDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetUseHostClipboard(request *IConsolegetUseHostClipboard) (*IConsolegetUseHostClipboardResponse, error) {
	response := new(IConsolegetUseHostClipboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesetUseHostClipboard(request *IConsolesetUseHostClipboard) (*IConsolesetUseHostClipboardResponse, error) {
	response := new(IConsolesetUseHostClipboardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetEmulatedUSB(request *IConsolegetEmulatedUSB) (*IConsolegetEmulatedUSBResponse, error) {
	response := new(IConsolegetEmulatedUSBResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUp(request *IConsolepowerUp) (*IConsolepowerUpResponse, error) {
	response := new(IConsolepowerUpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerUpPaused(request *IConsolepowerUpPaused) (*IConsolepowerUpPausedResponse, error) {
	response := new(IConsolepowerUpPausedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerDown(request *IConsolepowerDown) (*IConsolepowerDownResponse, error) {
	response := new(IConsolepowerDownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolereset(request *IConsolereset) (*IConsoleresetResponse, error) {
	response := new(IConsoleresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepause(request *IConsolepause) (*IConsolepauseResponse, error) {
	response := new(IConsolepauseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleresume(request *IConsoleresume) (*IConsoleresumeResponse, error) {
	response := new(IConsoleresumeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolepowerButton(request *IConsolepowerButton) (*IConsolepowerButtonResponse, error) {
	response := new(IConsolepowerButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesleepButton(request *IConsolesleepButton) (*IConsolesleepButtonResponse, error) {
	response := new(IConsolesleepButtonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetPowerButtonHandled(request *IConsolegetPowerButtonHandled) (*IConsolegetPowerButtonHandledResponse, error) {
	response := new(IConsolegetPowerButtonHandledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetGuestEnteredACPIMode(request *IConsolegetGuestEnteredACPIMode) (*IConsolegetGuestEnteredACPIModeResponse, error) {
	response := new(IConsolegetGuestEnteredACPIModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolesaveState(request *IConsolesaveState) (*IConsolesaveStateResponse, error) {
	response := new(IConsolesaveStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleadoptSavedState(request *IConsoleadoptSavedState) (*IConsoleadoptSavedStateResponse, error) {
	response := new(IConsoleadoptSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolediscardSavedState(request *IConsolediscardSavedState) (*IConsolediscardSavedStateResponse, error) {
	response := new(IConsolediscardSavedStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolegetDeviceActivity(request *IConsolegetDeviceActivity) (*IConsolegetDeviceActivityResponse, error) {
	response := new(IConsolegetDeviceActivityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleattachUSBDevice(request *IConsoleattachUSBDevice) (*IConsoleattachUSBDeviceResponse, error) {
	response := new(IConsoleattachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledetachUSBDevice(request *IConsoledetachUSBDevice) (*IConsoledetachUSBDeviceResponse, error) {
	response := new(IConsoledetachUSBDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceByAddress(request *IConsolefindUSBDeviceByAddress) (*IConsolefindUSBDeviceByAddressResponse, error) {
	response := new(IConsolefindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolefindUSBDeviceById(request *IConsolefindUSBDeviceById) (*IConsolefindUSBDeviceByIdResponse, error) {
	response := new(IConsolefindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolecreateSharedFolder(request *IConsolecreateSharedFolder) (*IConsolecreateSharedFolderResponse, error) {
	response := new(IConsolecreateSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleremoveSharedFolder(request *IConsoleremoveSharedFolder) (*IConsoleremoveSharedFolderResponse, error) {
	response := new(IConsoleremoveSharedFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoletakeSnapshot(request *IConsoletakeSnapshot) (*IConsoletakeSnapshotResponse, error) {
	response := new(IConsoletakeSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledeleteSnapshot(request *IConsoledeleteSnapshot) (*IConsoledeleteSnapshotResponse, error) {
	response := new(IConsoledeleteSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledeleteSnapshotAndAllChildren(request *IConsoledeleteSnapshotAndAllChildren) (*IConsoledeleteSnapshotAndAllChildrenResponse, error) {
	response := new(IConsoledeleteSnapshotAndAllChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoledeleteSnapshotRange(request *IConsoledeleteSnapshotRange) (*IConsoledeleteSnapshotRangeResponse, error) {
	response := new(IConsoledeleteSnapshotRangeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsolerestoreSnapshot(request *IConsolerestoreSnapshot) (*IConsolerestoreSnapshotResponse, error) {
	response := new(IConsolerestoreSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IConsoleteleport(request *IConsoleteleport) (*IConsoleteleportResponse, error) {
	response := new(IConsoleteleportResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetName(request *IHostNetworkInterfacegetName) (*IHostNetworkInterfacegetNameResponse, error) {
	response := new(IHostNetworkInterfacegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetShortName(request *IHostNetworkInterfacegetShortName) (*IHostNetworkInterfacegetShortNameResponse, error) {
	response := new(IHostNetworkInterfacegetShortNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetId(request *IHostNetworkInterfacegetId) (*IHostNetworkInterfacegetIdResponse, error) {
	response := new(IHostNetworkInterfacegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkName(request *IHostNetworkInterfacegetNetworkName) (*IHostNetworkInterfacegetNetworkNameResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetDHCPEnabled(request *IHostNetworkInterfacegetDHCPEnabled) (*IHostNetworkInterfacegetDHCPEnabledResponse, error) {
	response := new(IHostNetworkInterfacegetDHCPEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPAddress(request *IHostNetworkInterfacegetIPAddress) (*IHostNetworkInterfacegetIPAddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetNetworkMask(request *IHostNetworkInterfacegetNetworkMask) (*IHostNetworkInterfacegetNetworkMaskResponse, error) {
	response := new(IHostNetworkInterfacegetNetworkMaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Supported(request *IHostNetworkInterfacegetIPV6Supported) (*IHostNetworkInterfacegetIPV6SupportedResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6SupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6Address(request *IHostNetworkInterfacegetIPV6Address) (*IHostNetworkInterfacegetIPV6AddressResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6AddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength(request *IHostNetworkInterfacegetIPV6NetworkMaskPrefixLength) (*IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse, error) {
	response := new(IHostNetworkInterfacegetIPV6NetworkMaskPrefixLengthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetHardwareAddress(request *IHostNetworkInterfacegetHardwareAddress) (*IHostNetworkInterfacegetHardwareAddressResponse, error) {
	response := new(IHostNetworkInterfacegetHardwareAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetMediumType(request *IHostNetworkInterfacegetMediumType) (*IHostNetworkInterfacegetMediumTypeResponse, error) {
	response := new(IHostNetworkInterfacegetMediumTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetStatus(request *IHostNetworkInterfacegetStatus) (*IHostNetworkInterfacegetStatusResponse, error) {
	response := new(IHostNetworkInterfacegetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfacegetInterfaceType(request *IHostNetworkInterfacegetInterfaceType) (*IHostNetworkInterfacegetInterfaceTypeResponse, error) {
	response := new(IHostNetworkInterfacegetInterfaceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIPConfig(request *IHostNetworkInterfaceenableStaticIPConfig) (*IHostNetworkInterfaceenableStaticIPConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableStaticIPConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableStaticIPConfigV6(request *IHostNetworkInterfaceenableStaticIPConfigV6) (*IHostNetworkInterfaceenableStaticIPConfigV6Response, error) {
	response := new(IHostNetworkInterfaceenableStaticIPConfigV6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceenableDynamicIPConfig(request *IHostNetworkInterfaceenableDynamicIPConfig) (*IHostNetworkInterfaceenableDynamicIPConfigResponse, error) {
	response := new(IHostNetworkInterfaceenableDynamicIPConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostNetworkInterfaceDHCPRediscover(request *IHostNetworkInterfaceDHCPRediscover) (*IHostNetworkInterfaceDHCPRediscoverResponse, error) {
	response := new(IHostNetworkInterfaceDHCPRediscoverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetName(request *IHostVideoInputDevicegetName) (*IHostVideoInputDevicegetNameResponse, error) {
	response := new(IHostVideoInputDevicegetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetPath(request *IHostVideoInputDevicegetPath) (*IHostVideoInputDevicegetPathResponse, error) {
	response := new(IHostVideoInputDevicegetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostVideoInputDevicegetAlias(request *IHostVideoInputDevicegetAlias) (*IHostVideoInputDevicegetAliasResponse, error) {
	response := new(IHostVideoInputDevicegetAliasResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetDVDDrives(request *IHostgetDVDDrives) (*IHostgetDVDDrivesResponse, error) {
	response := new(IHostgetDVDDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetFloppyDrives(request *IHostgetFloppyDrives) (*IHostgetFloppyDrivesResponse, error) {
	response := new(IHostgetFloppyDrivesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDevices(request *IHostgetUSBDevices) (*IHostgetUSBDevicesResponse, error) {
	response := new(IHostgetUSBDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUSBDeviceFilters(request *IHostgetUSBDeviceFilters) (*IHostgetUSBDeviceFiltersResponse, error) {
	response := new(IHostgetUSBDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetNetworkInterfaces(request *IHostgetNetworkInterfaces) (*IHostgetNetworkInterfacesResponse, error) {
	response := new(IHostgetNetworkInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetNameServers(request *IHostgetNameServers) (*IHostgetNameServersResponse, error) {
	response := new(IHostgetNameServersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetDomainName(request *IHostgetDomainName) (*IHostgetDomainNameResponse, error) {
	response := new(IHostgetDomainNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetSearchStrings(request *IHostgetSearchStrings) (*IHostgetSearchStringsResponse, error) {
	response := new(IHostgetSearchStringsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCount(request *IHostgetProcessorCount) (*IHostgetProcessorCountResponse, error) {
	response := new(IHostgetProcessorCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorOnlineCount(request *IHostgetProcessorOnlineCount) (*IHostgetProcessorOnlineCountResponse, error) {
	response := new(IHostgetProcessorOnlineCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCoreCount(request *IHostgetProcessorCoreCount) (*IHostgetProcessorCoreCountResponse, error) {
	response := new(IHostgetProcessorCoreCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorOnlineCoreCount(request *IHostgetProcessorOnlineCoreCount) (*IHostgetProcessorOnlineCoreCountResponse, error) {
	response := new(IHostgetProcessorOnlineCoreCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemorySize(request *IHostgetMemorySize) (*IHostgetMemorySizeResponse, error) {
	response := new(IHostgetMemorySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetMemoryAvailable(request *IHostgetMemoryAvailable) (*IHostgetMemoryAvailableResponse, error) {
	response := new(IHostgetMemoryAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOperatingSystem(request *IHostgetOperatingSystem) (*IHostgetOperatingSystemResponse, error) {
	response := new(IHostgetOperatingSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetOSVersion(request *IHostgetOSVersion) (*IHostgetOSVersionResponse, error) {
	response := new(IHostgetOSVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetUTCTime(request *IHostgetUTCTime) (*IHostgetUTCTimeResponse, error) {
	response := new(IHostgetUTCTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetAcceleration3DAvailable(request *IHostgetAcceleration3DAvailable) (*IHostgetAcceleration3DAvailableResponse, error) {
	response := new(IHostgetAcceleration3DAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetVideoInputDevices(request *IHostgetVideoInputDevices) (*IHostgetVideoInputDevicesResponse, error) {
	response := new(IHostgetVideoInputDevicesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorSpeed(request *IHostgetProcessorSpeed) (*IHostgetProcessorSpeedResponse, error) {
	response := new(IHostgetProcessorSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorFeature(request *IHostgetProcessorFeature) (*IHostgetProcessorFeatureResponse, error) {
	response := new(IHostgetProcessorFeatureResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorDescription(request *IHostgetProcessorDescription) (*IHostgetProcessorDescriptionResponse, error) {
	response := new(IHostgetProcessorDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgetProcessorCPUIDLeaf(request *IHostgetProcessorCPUIDLeaf) (*IHostgetProcessorCPUIDLeafResponse, error) {
	response := new(IHostgetProcessorCPUIDLeafResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateHostOnlyNetworkInterface(request *IHostcreateHostOnlyNetworkInterface) (*IHostcreateHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostcreateHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveHostOnlyNetworkInterface(request *IHostremoveHostOnlyNetworkInterface) (*IHostremoveHostOnlyNetworkInterfaceResponse, error) {
	response := new(IHostremoveHostOnlyNetworkInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostcreateUSBDeviceFilter(request *IHostcreateUSBDeviceFilter) (*IHostcreateUSBDeviceFilterResponse, error) {
	response := new(IHostcreateUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostinsertUSBDeviceFilter(request *IHostinsertUSBDeviceFilter) (*IHostinsertUSBDeviceFilterResponse, error) {
	response := new(IHostinsertUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostremoveUSBDeviceFilter(request *IHostremoveUSBDeviceFilter) (*IHostremoveUSBDeviceFilterResponse, error) {
	response := new(IHostremoveUSBDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostDVDDrive(request *IHostfindHostDVDDrive) (*IHostfindHostDVDDriveResponse, error) {
	response := new(IHostfindHostDVDDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostFloppyDrive(request *IHostfindHostFloppyDrive) (*IHostfindHostFloppyDriveResponse, error) {
	response := new(IHostfindHostFloppyDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceByName(request *IHostfindHostNetworkInterfaceByName) (*IHostfindHostNetworkInterfaceByNameResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfaceById(request *IHostfindHostNetworkInterfaceById) (*IHostfindHostNetworkInterfaceByIdResponse, error) {
	response := new(IHostfindHostNetworkInterfaceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindHostNetworkInterfacesOfType(request *IHostfindHostNetworkInterfacesOfType) (*IHostfindHostNetworkInterfacesOfTypeResponse, error) {
	response := new(IHostfindHostNetworkInterfacesOfTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceById(request *IHostfindUSBDeviceById) (*IHostfindUSBDeviceByIdResponse, error) {
	response := new(IHostfindUSBDeviceByIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostfindUSBDeviceByAddress(request *IHostfindUSBDeviceByAddress) (*IHostfindUSBDeviceByAddressResponse, error) {
	response := new(IHostfindUSBDeviceByAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostgenerateMACAddress(request *IHostgenerateMACAddress) (*IHostgenerateMACAddressResponse, error) {
	response := new(IHostgenerateMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestRAM(request *ISystemPropertiesgetMinGuestRAM) (*ISystemPropertiesgetMinGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestRAM(request *ISystemPropertiesgetMaxGuestRAM) (*ISystemPropertiesgetMaxGuestRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestVRAM(request *ISystemPropertiesgetMinGuestVRAM) (*ISystemPropertiesgetMinGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMinGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestVRAM(request *ISystemPropertiesgetMaxGuestVRAM) (*ISystemPropertiesgetMaxGuestVRAMResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinGuestCPUCount(request *ISystemPropertiesgetMinGuestCPUCount) (*ISystemPropertiesgetMinGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMinGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestCPUCount(request *ISystemPropertiesgetMaxGuestCPUCount) (*ISystemPropertiesgetMaxGuestCPUCountResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestCPUCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxGuestMonitors(request *ISystemPropertiesgetMaxGuestMonitors) (*ISystemPropertiesgetMaxGuestMonitorsResponse, error) {
	response := new(ISystemPropertiesgetMaxGuestMonitorsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetInfoVDSize(request *ISystemPropertiesgetInfoVDSize) (*ISystemPropertiesgetInfoVDSizeResponse, error) {
	response := new(ISystemPropertiesgetInfoVDSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetSerialPortCount(request *ISystemPropertiesgetSerialPortCount) (*ISystemPropertiesgetSerialPortCountResponse, error) {
	response := new(ISystemPropertiesgetSerialPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetParallelPortCount(request *ISystemPropertiesgetParallelPortCount) (*ISystemPropertiesgetParallelPortCountResponse, error) {
	response := new(ISystemPropertiesgetParallelPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxBootPosition(request *ISystemPropertiesgetMaxBootPosition) (*ISystemPropertiesgetMaxBootPositionResponse, error) {
	response := new(ISystemPropertiesgetMaxBootPositionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetExclusiveHwVirt(request *ISystemPropertiesgetExclusiveHwVirt) (*ISystemPropertiesgetExclusiveHwVirtResponse, error) {
	response := new(ISystemPropertiesgetExclusiveHwVirtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetExclusiveHwVirt(request *ISystemPropertiessetExclusiveHwVirt) (*ISystemPropertiessetExclusiveHwVirtResponse, error) {
	response := new(ISystemPropertiessetExclusiveHwVirtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultMachineFolder(request *ISystemPropertiesgetDefaultMachineFolder) (*ISystemPropertiesgetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiesgetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultMachineFolder(request *ISystemPropertiessetDefaultMachineFolder) (*ISystemPropertiessetDefaultMachineFolderResponse, error) {
	response := new(ISystemPropertiessetDefaultMachineFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetLoggingLevel(request *ISystemPropertiesgetLoggingLevel) (*ISystemPropertiesgetLoggingLevelResponse, error) {
	response := new(ISystemPropertiesgetLoggingLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetLoggingLevel(request *ISystemPropertiessetLoggingLevel) (*ISystemPropertiessetLoggingLevelResponse, error) {
	response := new(ISystemPropertiessetLoggingLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMediumFormats(request *ISystemPropertiesgetMediumFormats) (*ISystemPropertiesgetMediumFormatsResponse, error) {
	response := new(ISystemPropertiesgetMediumFormatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultHardDiskFormat(request *ISystemPropertiesgetDefaultHardDiskFormat) (*ISystemPropertiesgetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiesgetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultHardDiskFormat(request *ISystemPropertiessetDefaultHardDiskFormat) (*ISystemPropertiessetDefaultHardDiskFormatResponse, error) {
	response := new(ISystemPropertiessetDefaultHardDiskFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceWarning(request *ISystemPropertiesgetFreeDiskSpaceWarning) (*ISystemPropertiesgetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceWarning(request *ISystemPropertiessetFreeDiskSpaceWarning) (*ISystemPropertiessetFreeDiskSpaceWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentWarning(request *ISystemPropertiesgetFreeDiskSpacePercentWarning) (*ISystemPropertiesgetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentWarning(request *ISystemPropertiessetFreeDiskSpacePercentWarning) (*ISystemPropertiessetFreeDiskSpacePercentWarningResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentWarningResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpaceError(request *ISystemPropertiesgetFreeDiskSpaceError) (*ISystemPropertiesgetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpaceError(request *ISystemPropertiessetFreeDiskSpaceError) (*ISystemPropertiessetFreeDiskSpaceErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpaceErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetFreeDiskSpacePercentError(request *ISystemPropertiesgetFreeDiskSpacePercentError) (*ISystemPropertiesgetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiesgetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetFreeDiskSpacePercentError(request *ISystemPropertiessetFreeDiskSpacePercentError) (*ISystemPropertiessetFreeDiskSpacePercentErrorResponse, error) {
	response := new(ISystemPropertiessetFreeDiskSpacePercentErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetVRDEAuthLibrary(request *ISystemPropertiesgetVRDEAuthLibrary) (*ISystemPropertiesgetVRDEAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetVRDEAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetVRDEAuthLibrary(request *ISystemPropertiessetVRDEAuthLibrary) (*ISystemPropertiessetVRDEAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetVRDEAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetWebServiceAuthLibrary(request *ISystemPropertiesgetWebServiceAuthLibrary) (*ISystemPropertiesgetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiesgetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetWebServiceAuthLibrary(request *ISystemPropertiessetWebServiceAuthLibrary) (*ISystemPropertiessetWebServiceAuthLibraryResponse, error) {
	response := new(ISystemPropertiessetWebServiceAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultVRDEExtPack(request *ISystemPropertiesgetDefaultVRDEExtPack) (*ISystemPropertiesgetDefaultVRDEExtPackResponse, error) {
	response := new(ISystemPropertiesgetDefaultVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultVRDEExtPack(request *ISystemPropertiessetDefaultVRDEExtPack) (*ISystemPropertiessetDefaultVRDEExtPackResponse, error) {
	response := new(ISystemPropertiessetDefaultVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetLogHistoryCount(request *ISystemPropertiesgetLogHistoryCount) (*ISystemPropertiesgetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiesgetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetLogHistoryCount(request *ISystemPropertiessetLogHistoryCount) (*ISystemPropertiessetLogHistoryCountResponse, error) {
	response := new(ISystemPropertiessetLogHistoryCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultAudioDriver(request *ISystemPropertiesgetDefaultAudioDriver) (*ISystemPropertiesgetDefaultAudioDriverResponse, error) {
	response := new(ISystemPropertiesgetDefaultAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetAutostartDatabasePath(request *ISystemPropertiesgetAutostartDatabasePath) (*ISystemPropertiesgetAutostartDatabasePathResponse, error) {
	response := new(ISystemPropertiesgetAutostartDatabasePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetAutostartDatabasePath(request *ISystemPropertiessetAutostartDatabasePath) (*ISystemPropertiessetAutostartDatabasePathResponse, error) {
	response := new(ISystemPropertiessetAutostartDatabasePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultAdditionsISO(request *ISystemPropertiesgetDefaultAdditionsISO) (*ISystemPropertiesgetDefaultAdditionsISOResponse, error) {
	response := new(ISystemPropertiesgetDefaultAdditionsISOResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultAdditionsISO(request *ISystemPropertiessetDefaultAdditionsISO) (*ISystemPropertiessetDefaultAdditionsISOResponse, error) {
	response := new(ISystemPropertiessetDefaultAdditionsISOResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultFrontend(request *ISystemPropertiesgetDefaultFrontend) (*ISystemPropertiesgetDefaultFrontendResponse, error) {
	response := new(ISystemPropertiesgetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiessetDefaultFrontend(request *ISystemPropertiessetDefaultFrontend) (*ISystemPropertiessetDefaultFrontendResponse, error) {
	response := new(ISystemPropertiessetDefaultFrontendResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxNetworkAdapters(request *ISystemPropertiesgetMaxNetworkAdapters) (*ISystemPropertiesgetMaxNetworkAdaptersResponse, error) {
	response := new(ISystemPropertiesgetMaxNetworkAdaptersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxNetworkAdaptersOfType(request *ISystemPropertiesgetMaxNetworkAdaptersOfType) (*ISystemPropertiesgetMaxNetworkAdaptersOfTypeResponse, error) {
	response := new(ISystemPropertiesgetMaxNetworkAdaptersOfTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxDevicesPerPortForStorageBus(request *ISystemPropertiesgetMaxDevicesPerPortForStorageBus) (*ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxDevicesPerPortForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMinPortCountForStorageBus(request *ISystemPropertiesgetMinPortCountForStorageBus) (*ISystemPropertiesgetMinPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMinPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxPortCountForStorageBus(request *ISystemPropertiesgetMaxPortCountForStorageBus) (*ISystemPropertiesgetMaxPortCountForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxPortCountForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxInstancesOfStorageBus(request *ISystemPropertiesgetMaxInstancesOfStorageBus) (*ISystemPropertiesgetMaxInstancesOfStorageBusResponse, error) {
	response := new(ISystemPropertiesgetMaxInstancesOfStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDeviceTypesForStorageBus(request *ISystemPropertiesgetDeviceTypesForStorageBus) (*ISystemPropertiesgetDeviceTypesForStorageBusResponse, error) {
	response := new(ISystemPropertiesgetDeviceTypesForStorageBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetDefaultIoCacheSettingForStorageController(request *ISystemPropertiesgetDefaultIoCacheSettingForStorageController) (*ISystemPropertiesgetDefaultIoCacheSettingForStorageControllerResponse, error) {
	response := new(ISystemPropertiesgetDefaultIoCacheSettingForStorageControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISystemPropertiesgetMaxInstancesOfUSBControllerType(request *ISystemPropertiesgetMaxInstancesOfUSBControllerType) (*ISystemPropertiesgetMaxInstancesOfUSBControllerTypeResponse, error) {
	response := new(ISystemPropertiesgetMaxInstancesOfUSBControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetUser(request *IGuestSessiongetUser) (*IGuestSessiongetUserResponse, error) {
	response := new(IGuestSessiongetUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetDomain(request *IGuestSessiongetDomain) (*IGuestSessiongetDomainResponse, error) {
	response := new(IGuestSessiongetDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetName(request *IGuestSessiongetName) (*IGuestSessiongetNameResponse, error) {
	response := new(IGuestSessiongetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetId(request *IGuestSessiongetId) (*IGuestSessiongetIdResponse, error) {
	response := new(IGuestSessiongetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetTimeout(request *IGuestSessiongetTimeout) (*IGuestSessiongetTimeoutResponse, error) {
	response := new(IGuestSessiongetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsetTimeout(request *IGuestSessionsetTimeout) (*IGuestSessionsetTimeoutResponse, error) {
	response := new(IGuestSessionsetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetProtocolVersion(request *IGuestSessiongetProtocolVersion) (*IGuestSessiongetProtocolVersionResponse, error) {
	response := new(IGuestSessiongetProtocolVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetStatus(request *IGuestSessiongetStatus) (*IGuestSessiongetStatusResponse, error) {
	response := new(IGuestSessiongetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetEnvironment(request *IGuestSessiongetEnvironment) (*IGuestSessiongetEnvironmentResponse, error) {
	response := new(IGuestSessiongetEnvironmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsetEnvironment(request *IGuestSessionsetEnvironment) (*IGuestSessionsetEnvironmentResponse, error) {
	response := new(IGuestSessionsetEnvironmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetProcesses(request *IGuestSessiongetProcesses) (*IGuestSessiongetProcessesResponse, error) {
	response := new(IGuestSessiongetProcessesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetDirectories(request *IGuestSessiongetDirectories) (*IGuestSessiongetDirectoriesResponse, error) {
	response := new(IGuestSessiongetDirectoriesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetFiles(request *IGuestSessiongetFiles) (*IGuestSessiongetFilesResponse, error) {
	response := new(IGuestSessiongetFilesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiongetEventSource(request *IGuestSessiongetEventSource) (*IGuestSessiongetEventSourceResponse, error) {
	response := new(IGuestSessiongetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionclose(request *IGuestSessionclose) (*IGuestSessioncloseResponse, error) {
	response := new(IGuestSessioncloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessioncopyFrom(request *IGuestSessioncopyFrom) (*IGuestSessioncopyFromResponse, error) {
	response := new(IGuestSessioncopyFromResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessioncopyTo(request *IGuestSessioncopyTo) (*IGuestSessioncopyToResponse, error) {
	response := new(IGuestSessioncopyToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCreate(request *IGuestSessiondirectoryCreate) (*IGuestSessiondirectoryCreateResponse, error) {
	response := new(IGuestSessiondirectoryCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryCreateTemp(request *IGuestSessiondirectoryCreateTemp) (*IGuestSessiondirectoryCreateTempResponse, error) {
	response := new(IGuestSessiondirectoryCreateTempResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryExists(request *IGuestSessiondirectoryExists) (*IGuestSessiondirectoryExistsResponse, error) {
	response := new(IGuestSessiondirectoryExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryOpen(request *IGuestSessiondirectoryOpen) (*IGuestSessiondirectoryOpenResponse, error) {
	response := new(IGuestSessiondirectoryOpenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryQueryInfo(request *IGuestSessiondirectoryQueryInfo) (*IGuestSessiondirectoryQueryInfoResponse, error) {
	response := new(IGuestSessiondirectoryQueryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryRemove(request *IGuestSessiondirectoryRemove) (*IGuestSessiondirectoryRemoveResponse, error) {
	response := new(IGuestSessiondirectoryRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryRemoveRecursive(request *IGuestSessiondirectoryRemoveRecursive) (*IGuestSessiondirectoryRemoveRecursiveResponse, error) {
	response := new(IGuestSessiondirectoryRemoveRecursiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectoryRename(request *IGuestSessiondirectoryRename) (*IGuestSessiondirectoryRenameResponse, error) {
	response := new(IGuestSessiondirectoryRenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessiondirectorySetACL(request *IGuestSessiondirectorySetACL) (*IGuestSessiondirectorySetACLResponse, error) {
	response := new(IGuestSessiondirectorySetACLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentClear(request *IGuestSessionenvironmentClear) (*IGuestSessionenvironmentClearResponse, error) {
	response := new(IGuestSessionenvironmentClearResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentGet(request *IGuestSessionenvironmentGet) (*IGuestSessionenvironmentGetResponse, error) {
	response := new(IGuestSessionenvironmentGetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentSet(request *IGuestSessionenvironmentSet) (*IGuestSessionenvironmentSetResponse, error) {
	response := new(IGuestSessionenvironmentSetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionenvironmentUnset(request *IGuestSessionenvironmentUnset) (*IGuestSessionenvironmentUnsetResponse, error) {
	response := new(IGuestSessionenvironmentUnsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileCreateTemp(request *IGuestSessionfileCreateTemp) (*IGuestSessionfileCreateTempResponse, error) {
	response := new(IGuestSessionfileCreateTempResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileExists(request *IGuestSessionfileExists) (*IGuestSessionfileExistsResponse, error) {
	response := new(IGuestSessionfileExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileRemove(request *IGuestSessionfileRemove) (*IGuestSessionfileRemoveResponse, error) {
	response := new(IGuestSessionfileRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileOpen(request *IGuestSessionfileOpen) (*IGuestSessionfileOpenResponse, error) {
	response := new(IGuestSessionfileOpenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileOpenEx(request *IGuestSessionfileOpenEx) (*IGuestSessionfileOpenExResponse, error) {
	response := new(IGuestSessionfileOpenExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileQueryInfo(request *IGuestSessionfileQueryInfo) (*IGuestSessionfileQueryInfoResponse, error) {
	response := new(IGuestSessionfileQueryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileQuerySize(request *IGuestSessionfileQuerySize) (*IGuestSessionfileQuerySizeResponse, error) {
	response := new(IGuestSessionfileQuerySizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileRename(request *IGuestSessionfileRename) (*IGuestSessionfileRenameResponse, error) {
	response := new(IGuestSessionfileRenameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionfileSetACL(request *IGuestSessionfileSetACL) (*IGuestSessionfileSetACLResponse, error) {
	response := new(IGuestSessionfileSetACLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessCreate(request *IGuestSessionprocessCreate) (*IGuestSessionprocessCreateResponse, error) {
	response := new(IGuestSessionprocessCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessCreateEx(request *IGuestSessionprocessCreateEx) (*IGuestSessionprocessCreateExResponse, error) {
	response := new(IGuestSessionprocessCreateExResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionprocessGet(request *IGuestSessionprocessGet) (*IGuestSessionprocessGetResponse, error) {
	response := new(IGuestSessionprocessGetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkCreate(request *IGuestSessionsymlinkCreate) (*IGuestSessionsymlinkCreateResponse, error) {
	response := new(IGuestSessionsymlinkCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkExists(request *IGuestSessionsymlinkExists) (*IGuestSessionsymlinkExistsResponse, error) {
	response := new(IGuestSessionsymlinkExistsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkRead(request *IGuestSessionsymlinkRead) (*IGuestSessionsymlinkReadResponse, error) {
	response := new(IGuestSessionsymlinkReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkRemoveDirectory(request *IGuestSessionsymlinkRemoveDirectory) (*IGuestSessionsymlinkRemoveDirectoryResponse, error) {
	response := new(IGuestSessionsymlinkRemoveDirectoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionsymlinkRemoveFile(request *IGuestSessionsymlinkRemoveFile) (*IGuestSessionsymlinkRemoveFileResponse, error) {
	response := new(IGuestSessionsymlinkRemoveFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionwaitFor(request *IGuestSessionwaitFor) (*IGuestSessionwaitForResponse, error) {
	response := new(IGuestSessionwaitForResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionwaitForArray(request *IGuestSessionwaitForArray) (*IGuestSessionwaitForArrayResponse, error) {
	response := new(IGuestSessionwaitForArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetArguments(request *IProcessgetArguments) (*IProcessgetArgumentsResponse, error) {
	response := new(IProcessgetArgumentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetEnvironment(request *IProcessgetEnvironment) (*IProcessgetEnvironmentResponse, error) {
	response := new(IProcessgetEnvironmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetEventSource(request *IProcessgetEventSource) (*IProcessgetEventSourceResponse, error) {
	response := new(IProcessgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetExecutablePath(request *IProcessgetExecutablePath) (*IProcessgetExecutablePathResponse, error) {
	response := new(IProcessgetExecutablePathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetExitCode(request *IProcessgetExitCode) (*IProcessgetExitCodeResponse, error) {
	response := new(IProcessgetExitCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetName(request *IProcessgetName) (*IProcessgetNameResponse, error) {
	response := new(IProcessgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetPID(request *IProcessgetPID) (*IProcessgetPIDResponse, error) {
	response := new(IProcessgetPIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessgetStatus(request *IProcessgetStatus) (*IProcessgetStatusResponse, error) {
	response := new(IProcessgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswaitFor(request *IProcesswaitFor) (*IProcesswaitForResponse, error) {
	response := new(IProcesswaitForResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswaitForArray(request *IProcesswaitForArray) (*IProcesswaitForArrayResponse, error) {
	response := new(IProcesswaitForArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessread(request *IProcessread) (*IProcessreadResponse, error) {
	response := new(IProcessreadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswrite(request *IProcesswrite) (*IProcesswriteResponse, error) {
	response := new(IProcesswriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcesswriteArray(request *IProcesswriteArray) (*IProcesswriteArrayResponse, error) {
	response := new(IProcesswriteArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProcessterminate(request *IProcessterminate) (*IProcessterminateResponse, error) {
	response := new(IProcessterminateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectorygetDirectoryName(request *IDirectorygetDirectoryName) (*IDirectorygetDirectoryNameResponse, error) {
	response := new(IDirectorygetDirectoryNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectorygetFilter(request *IDirectorygetFilter) (*IDirectorygetFilterResponse, error) {
	response := new(IDirectorygetFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectoryclose(request *IDirectoryclose) (*IDirectorycloseResponse, error) {
	response := new(IDirectorycloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDirectoryread(request *IDirectoryread) (*IDirectoryreadResponse, error) {
	response := new(IDirectoryreadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetCreationMode(request *IFilegetCreationMode) (*IFilegetCreationModeResponse, error) {
	response := new(IFilegetCreationModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetDisposition(request *IFilegetDisposition) (*IFilegetDispositionResponse, error) {
	response := new(IFilegetDispositionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetEventSource(request *IFilegetEventSource) (*IFilegetEventSourceResponse, error) {
	response := new(IFilegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetFileName(request *IFilegetFileName) (*IFilegetFileNameResponse, error) {
	response := new(IFilegetFileNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetId(request *IFilegetId) (*IFilegetIdResponse, error) {
	response := new(IFilegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetInitialSize(request *IFilegetInitialSize) (*IFilegetInitialSizeResponse, error) {
	response := new(IFilegetInitialSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetOpenMode(request *IFilegetOpenMode) (*IFilegetOpenModeResponse, error) {
	response := new(IFilegetOpenModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetOffset(request *IFilegetOffset) (*IFilegetOffsetResponse, error) {
	response := new(IFilegetOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilegetStatus(request *IFilegetStatus) (*IFilegetStatusResponse, error) {
	response := new(IFilegetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileclose(request *IFileclose) (*IFilecloseResponse, error) {
	response := new(IFilecloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilequeryInfo(request *IFilequeryInfo) (*IFilequeryInfoResponse, error) {
	response := new(IFilequeryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileread(request *IFileread) (*IFilereadResponse, error) {
	response := new(IFilereadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilereadAt(request *IFilereadAt) (*IFilereadAtResponse, error) {
	response := new(IFilereadAtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFileseek(request *IFileseek) (*IFileseekResponse, error) {
	response := new(IFileseekResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilesetACL(request *IFilesetACL) (*IFilesetACLResponse, error) {
	response := new(IFilesetACLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilewrite(request *IFilewrite) (*IFilewriteResponse, error) {
	response := new(IFilewriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFilewriteAt(request *IFilewriteAt) (*IFilewriteAtResponse, error) {
	response := new(IFilewriteAtResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetAccessTime(request *IFsObjInfogetAccessTime) (*IFsObjInfogetAccessTimeResponse, error) {
	response := new(IFsObjInfogetAccessTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetAllocatedSize(request *IFsObjInfogetAllocatedSize) (*IFsObjInfogetAllocatedSizeResponse, error) {
	response := new(IFsObjInfogetAllocatedSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetBirthTime(request *IFsObjInfogetBirthTime) (*IFsObjInfogetBirthTimeResponse, error) {
	response := new(IFsObjInfogetBirthTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetChangeTime(request *IFsObjInfogetChangeTime) (*IFsObjInfogetChangeTimeResponse, error) {
	response := new(IFsObjInfogetChangeTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetDeviceNumber(request *IFsObjInfogetDeviceNumber) (*IFsObjInfogetDeviceNumberResponse, error) {
	response := new(IFsObjInfogetDeviceNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetFileAttributes(request *IFsObjInfogetFileAttributes) (*IFsObjInfogetFileAttributesResponse, error) {
	response := new(IFsObjInfogetFileAttributesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGenerationId(request *IFsObjInfogetGenerationId) (*IFsObjInfogetGenerationIdResponse, error) {
	response := new(IFsObjInfogetGenerationIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGID(request *IFsObjInfogetGID) (*IFsObjInfogetGIDResponse, error) {
	response := new(IFsObjInfogetGIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetGroupName(request *IFsObjInfogetGroupName) (*IFsObjInfogetGroupNameResponse, error) {
	response := new(IFsObjInfogetGroupNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetHardLinks(request *IFsObjInfogetHardLinks) (*IFsObjInfogetHardLinksResponse, error) {
	response := new(IFsObjInfogetHardLinksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetModificationTime(request *IFsObjInfogetModificationTime) (*IFsObjInfogetModificationTimeResponse, error) {
	response := new(IFsObjInfogetModificationTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetName(request *IFsObjInfogetName) (*IFsObjInfogetNameResponse, error) {
	response := new(IFsObjInfogetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetNodeId(request *IFsObjInfogetNodeId) (*IFsObjInfogetNodeIdResponse, error) {
	response := new(IFsObjInfogetNodeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetNodeIdDevice(request *IFsObjInfogetNodeIdDevice) (*IFsObjInfogetNodeIdDeviceResponse, error) {
	response := new(IFsObjInfogetNodeIdDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetObjectSize(request *IFsObjInfogetObjectSize) (*IFsObjInfogetObjectSizeResponse, error) {
	response := new(IFsObjInfogetObjectSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetType(request *IFsObjInfogetType) (*IFsObjInfogetTypeResponse, error) {
	response := new(IFsObjInfogetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUID(request *IFsObjInfogetUID) (*IFsObjInfogetUIDResponse, error) {
	response := new(IFsObjInfogetUIDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUserFlags(request *IFsObjInfogetUserFlags) (*IFsObjInfogetUserFlagsResponse, error) {
	response := new(IFsObjInfogetUserFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFsObjInfogetUserName(request *IFsObjInfogetUserName) (*IFsObjInfogetUserNameResponse, error) {
	response := new(IFsObjInfogetUserNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetOSTypeId(request *IGuestgetOSTypeId) (*IGuestgetOSTypeIdResponse, error) {
	response := new(IGuestgetOSTypeIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsRunLevel(request *IGuestgetAdditionsRunLevel) (*IGuestgetAdditionsRunLevelResponse, error) {
	response := new(IGuestgetAdditionsRunLevelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsVersion(request *IGuestgetAdditionsVersion) (*IGuestgetAdditionsVersionResponse, error) {
	response := new(IGuestgetAdditionsVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsRevision(request *IGuestgetAdditionsRevision) (*IGuestgetAdditionsRevisionResponse, error) {
	response := new(IGuestgetAdditionsRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetEventSource(request *IGuestgetEventSource) (*IGuestgetEventSourceResponse, error) {
	response := new(IGuestgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetFacilities(request *IGuestgetFacilities) (*IGuestgetFacilitiesResponse, error) {
	response := new(IGuestgetFacilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetSessions(request *IGuestgetSessions) (*IGuestgetSessionsResponse, error) {
	response := new(IGuestgetSessionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetMemoryBalloonSize(request *IGuestgetMemoryBalloonSize) (*IGuestgetMemoryBalloonSizeResponse, error) {
	response := new(IGuestgetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetMemoryBalloonSize(request *IGuestsetMemoryBalloonSize) (*IGuestsetMemoryBalloonSizeResponse, error) {
	response := new(IGuestsetMemoryBalloonSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetStatisticsUpdateInterval(request *IGuestgetStatisticsUpdateInterval) (*IGuestgetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestgetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetStatisticsUpdateInterval(request *IGuestsetStatisticsUpdateInterval) (*IGuestsetStatisticsUpdateIntervalResponse, error) {
	response := new(IGuestsetStatisticsUpdateIntervalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestinternalGetStatistics(request *IGuestinternalGetStatistics) (*IGuestinternalGetStatisticsResponse, error) {
	response := new(IGuestinternalGetStatisticsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetFacilityStatus(request *IGuestgetFacilityStatus) (*IGuestgetFacilityStatusResponse, error) {
	response := new(IGuestgetFacilityStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestgetAdditionsStatus(request *IGuestgetAdditionsStatus) (*IGuestgetAdditionsStatusResponse, error) {
	response := new(IGuestgetAdditionsStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestsetCredentials(request *IGuestsetCredentials) (*IGuestsetCredentialsResponse, error) {
	response := new(IGuestsetCredentialsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragHGEnter(request *IGuestdragHGEnter) (*IGuestdragHGEnterResponse, error) {
	response := new(IGuestdragHGEnterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragHGMove(request *IGuestdragHGMove) (*IGuestdragHGMoveResponse, error) {
	response := new(IGuestdragHGMoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragHGLeave(request *IGuestdragHGLeave) (*IGuestdragHGLeaveResponse, error) {
	response := new(IGuestdragHGLeaveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragHGDrop(request *IGuestdragHGDrop) (*IGuestdragHGDropResponse, error) {
	response := new(IGuestdragHGDropResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragHGPutData(request *IGuestdragHGPutData) (*IGuestdragHGPutDataResponse, error) {
	response := new(IGuestdragHGPutDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragGHPending(request *IGuestdragGHPending) (*IGuestdragGHPendingResponse, error) {
	response := new(IGuestdragGHPendingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragGHDropped(request *IGuestdragGHDropped) (*IGuestdragGHDroppedResponse, error) {
	response := new(IGuestdragGHDroppedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestdragGHGetData(request *IGuestdragGHGetData) (*IGuestdragGHGetDataResponse, error) {
	response := new(IGuestdragGHGetDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestcreateSession(request *IGuestcreateSession) (*IGuestcreateSessionResponse, error) {
	response := new(IGuestcreateSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestfindSession(request *IGuestfindSession) (*IGuestfindSessionResponse, error) {
	response := new(IGuestfindSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestupdateGuestAdditions(request *IGuestupdateGuestAdditions) (*IGuestupdateGuestAdditionsResponse, error) {
	response := new(IGuestupdateGuestAdditionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetId(request *IProgressgetId) (*IProgressgetIdResponse, error) {
	response := new(IProgressgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetDescription(request *IProgressgetDescription) (*IProgressgetDescriptionResponse, error) {
	response := new(IProgressgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetInitiator(request *IProgressgetInitiator) (*IProgressgetInitiatorResponse, error) {
	response := new(IProgressgetInitiatorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCancelable(request *IProgressgetCancelable) (*IProgressgetCancelableResponse, error) {
	response := new(IProgressgetCancelableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetPercent(request *IProgressgetPercent) (*IProgressgetPercentResponse, error) {
	response := new(IProgressgetPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeRemaining(request *IProgressgetTimeRemaining) (*IProgressgetTimeRemainingResponse, error) {
	response := new(IProgressgetTimeRemainingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCompleted(request *IProgressgetCompleted) (*IProgressgetCompletedResponse, error) {
	response := new(IProgressgetCompletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetCanceled(request *IProgressgetCanceled) (*IProgressgetCanceledResponse, error) {
	response := new(IProgressgetCanceledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetResultCode(request *IProgressgetResultCode) (*IProgressgetResultCodeResponse, error) {
	response := new(IProgressgetResultCodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetErrorInfo(request *IProgressgetErrorInfo) (*IProgressgetErrorInfoResponse, error) {
	response := new(IProgressgetErrorInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationCount(request *IProgressgetOperationCount) (*IProgressgetOperationCountResponse, error) {
	response := new(IProgressgetOperationCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperation(request *IProgressgetOperation) (*IProgressgetOperationResponse, error) {
	response := new(IProgressgetOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationDescription(request *IProgressgetOperationDescription) (*IProgressgetOperationDescriptionResponse, error) {
	response := new(IProgressgetOperationDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationPercent(request *IProgressgetOperationPercent) (*IProgressgetOperationPercentResponse, error) {
	response := new(IProgressgetOperationPercentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetOperationWeight(request *IProgressgetOperationWeight) (*IProgressgetOperationWeightResponse, error) {
	response := new(IProgressgetOperationWeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgressgetTimeout(request *IProgressgetTimeout) (*IProgressgetTimeoutResponse, error) {
	response := new(IProgressgetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetTimeout(request *IProgresssetTimeout) (*IProgresssetTimeoutResponse, error) {
	response := new(IProgresssetTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetCurrentOperationProgress(request *IProgresssetCurrentOperationProgress) (*IProgresssetCurrentOperationProgressResponse, error) {
	response := new(IProgresssetCurrentOperationProgressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresssetNextOperation(request *IProgresssetNextOperation) (*IProgresssetNextOperationResponse, error) {
	response := new(IProgresssetNextOperationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForCompletion(request *IProgresswaitForCompletion) (*IProgresswaitForCompletionResponse, error) {
	response := new(IProgresswaitForCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForOperationCompletion(request *IProgresswaitForOperationCompletion) (*IProgresswaitForOperationCompletionResponse, error) {
	response := new(IProgresswaitForOperationCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresswaitForAsyncProgressCompletion(request *IProgresswaitForAsyncProgressCompletion) (*IProgresswaitForAsyncProgressCompletionResponse, error) {
	response := new(IProgresswaitForAsyncProgressCompletionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IProgresscancel(request *IProgresscancel) (*IProgresscancelResponse, error) {
	response := new(IProgresscancelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetId(request *ISnapshotgetId) (*ISnapshotgetIdResponse, error) {
	response := new(ISnapshotgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetName(request *ISnapshotgetName) (*ISnapshotgetNameResponse, error) {
	response := new(ISnapshotgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetName(request *ISnapshotsetName) (*ISnapshotsetNameResponse, error) {
	response := new(ISnapshotsetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetDescription(request *ISnapshotgetDescription) (*ISnapshotgetDescriptionResponse, error) {
	response := new(ISnapshotgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotsetDescription(request *ISnapshotsetDescription) (*ISnapshotsetDescriptionResponse, error) {
	response := new(ISnapshotsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetTimeStamp(request *ISnapshotgetTimeStamp) (*ISnapshotgetTimeStampResponse, error) {
	response := new(ISnapshotgetTimeStampResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetOnline(request *ISnapshotgetOnline) (*ISnapshotgetOnlineResponse, error) {
	response := new(ISnapshotgetOnlineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetMachine(request *ISnapshotgetMachine) (*ISnapshotgetMachineResponse, error) {
	response := new(ISnapshotgetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetParent(request *ISnapshotgetParent) (*ISnapshotgetParentResponse, error) {
	response := new(ISnapshotgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetChildren(request *ISnapshotgetChildren) (*ISnapshotgetChildrenResponse, error) {
	response := new(ISnapshotgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotgetChildrenCount(request *ISnapshotgetChildrenCount) (*ISnapshotgetChildrenCountResponse, error) {
	response := new(ISnapshotgetChildrenCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetId(request *IMediumgetId) (*IMediumgetIdResponse, error) {
	response := new(IMediumgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDescription(request *IMediumgetDescription) (*IMediumgetDescriptionResponse, error) {
	response := new(IMediumgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetDescription(request *IMediumsetDescription) (*IMediumsetDescriptionResponse, error) {
	response := new(IMediumsetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetState(request *IMediumgetState) (*IMediumgetStateResponse, error) {
	response := new(IMediumgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetVariant(request *IMediumgetVariant) (*IMediumgetVariantResponse, error) {
	response := new(IMediumgetVariantResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLocation(request *IMediumgetLocation) (*IMediumgetLocationResponse, error) {
	response := new(IMediumgetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetName(request *IMediumgetName) (*IMediumgetNameResponse, error) {
	response := new(IMediumgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetDeviceType(request *IMediumgetDeviceType) (*IMediumgetDeviceTypeResponse, error) {
	response := new(IMediumgetDeviceTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetHostDrive(request *IMediumgetHostDrive) (*IMediumgetHostDriveResponse, error) {
	response := new(IMediumgetHostDriveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSize(request *IMediumgetSize) (*IMediumgetSizeResponse, error) {
	response := new(IMediumgetSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetFormat(request *IMediumgetFormat) (*IMediumgetFormatResponse, error) {
	response := new(IMediumgetFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMediumFormat(request *IMediumgetMediumFormat) (*IMediumgetMediumFormatResponse, error) {
	response := new(IMediumgetMediumFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetType(request *IMediumgetType) (*IMediumgetTypeResponse, error) {
	response := new(IMediumgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetType(request *IMediumsetType) (*IMediumsetTypeResponse, error) {
	response := new(IMediumsetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetAllowedTypes(request *IMediumgetAllowedTypes) (*IMediumgetAllowedTypesResponse, error) {
	response := new(IMediumgetAllowedTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetParent(request *IMediumgetParent) (*IMediumgetParentResponse, error) {
	response := new(IMediumgetParentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetChildren(request *IMediumgetChildren) (*IMediumgetChildrenResponse, error) {
	response := new(IMediumgetChildrenResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetBase(request *IMediumgetBase) (*IMediumgetBaseResponse, error) {
	response := new(IMediumgetBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetReadOnly(request *IMediumgetReadOnly) (*IMediumgetReadOnlyResponse, error) {
	response := new(IMediumgetReadOnlyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLogicalSize(request *IMediumgetLogicalSize) (*IMediumgetLogicalSizeResponse, error) {
	response := new(IMediumgetLogicalSizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetAutoReset(request *IMediumgetAutoReset) (*IMediumgetAutoResetResponse, error) {
	response := new(IMediumgetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetAutoReset(request *IMediumsetAutoReset) (*IMediumsetAutoResetResponse, error) {
	response := new(IMediumsetAutoResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetLastAccessError(request *IMediumgetLastAccessError) (*IMediumgetLastAccessErrorResponse, error) {
	response := new(IMediumgetLastAccessErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetMachineIds(request *IMediumgetMachineIds) (*IMediumgetMachineIdsResponse, error) {
	response := new(IMediumgetMachineIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetIds(request *IMediumsetIds) (*IMediumsetIdsResponse, error) {
	response := new(IMediumsetIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumrefreshState(request *IMediumrefreshState) (*IMediumrefreshStateResponse, error) {
	response := new(IMediumrefreshStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetSnapshotIds(request *IMediumgetSnapshotIds) (*IMediumgetSnapshotIdsResponse, error) {
	response := new(IMediumgetSnapshotIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockRead(request *IMediumlockRead) (*IMediumlockReadResponse, error) {
	response := new(IMediumlockReadResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumlockWrite(request *IMediumlockWrite) (*IMediumlockWriteResponse, error) {
	response := new(IMediumlockWriteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumclose(request *IMediumclose) (*IMediumcloseResponse, error) {
	response := new(IMediumcloseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperty(request *IMediumgetProperty) (*IMediumgetPropertyResponse, error) {
	response := new(IMediumgetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperty(request *IMediumsetProperty) (*IMediumsetPropertyResponse, error) {
	response := new(IMediumsetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumgetProperties(request *IMediumgetProperties) (*IMediumgetPropertiesResponse, error) {
	response := new(IMediumgetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetProperties(request *IMediumsetProperties) (*IMediumsetPropertiesResponse, error) {
	response := new(IMediumsetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateBaseStorage(request *IMediumcreateBaseStorage) (*IMediumcreateBaseStorageResponse, error) {
	response := new(IMediumcreateBaseStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumdeleteStorage(request *IMediumdeleteStorage) (*IMediumdeleteStorageResponse, error) {
	response := new(IMediumdeleteStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcreateDiffStorage(request *IMediumcreateDiffStorage) (*IMediumcreateDiffStorageResponse, error) {
	response := new(IMediumcreateDiffStorageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediummergeTo(request *IMediummergeTo) (*IMediummergeToResponse, error) {
	response := new(IMediummergeToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcloneTo(request *IMediumcloneTo) (*IMediumcloneToResponse, error) {
	response := new(IMediumcloneToResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcloneToBase(request *IMediumcloneToBase) (*IMediumcloneToBaseResponse, error) {
	response := new(IMediumcloneToBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumsetLocation(request *IMediumsetLocation) (*IMediumsetLocationResponse, error) {
	response := new(IMediumsetLocationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumcompact(request *IMediumcompact) (*IMediumcompactResponse, error) {
	response := new(IMediumcompactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumresize(request *IMediumresize) (*IMediumresizeResponse, error) {
	response := new(IMediumresizeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumreset(request *IMediumreset) (*IMediumresetResponse, error) {
	response := new(IMediumresetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetId(request *IMediumFormatgetId) (*IMediumFormatgetIdResponse, error) {
	response := new(IMediumFormatgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetName(request *IMediumFormatgetName) (*IMediumFormatgetNameResponse, error) {
	response := new(IMediumFormatgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatgetCapabilities(request *IMediumFormatgetCapabilities) (*IMediumFormatgetCapabilitiesResponse, error) {
	response := new(IMediumFormatgetCapabilitiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatdescribeFileExtensions(request *IMediumFormatdescribeFileExtensions) (*IMediumFormatdescribeFileExtensionsResponse, error) {
	response := new(IMediumFormatdescribeFileExtensionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumFormatdescribeProperties(request *IMediumFormatdescribeProperties) (*IMediumFormatdescribePropertiesResponse, error) {
	response := new(IMediumFormatdescribePropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ITokenabandon(request *ITokenabandon) (*ITokenabandonResponse, error) {
	response := new(ITokenabandonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ITokendummy(request *ITokendummy) (*ITokendummyResponse, error) {
	response := new(ITokendummyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardgetEventSource(request *IKeyboardgetEventSource) (*IKeyboardgetEventSourceResponse, error) {
	response := new(IKeyboardgetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancode(request *IKeyboardputScancode) (*IKeyboardputScancodeResponse, error) {
	response := new(IKeyboardputScancodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputScancodes(request *IKeyboardputScancodes) (*IKeyboardputScancodesResponse, error) {
	response := new(IKeyboardputScancodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardputCAD(request *IKeyboardputCAD) (*IKeyboardputCADResponse, error) {
	response := new(IKeyboardputCADResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetAbsoluteSupported(request *IMousegetAbsoluteSupported) (*IMousegetAbsoluteSupportedResponse, error) {
	response := new(IMousegetAbsoluteSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetRelativeSupported(request *IMousegetRelativeSupported) (*IMousegetRelativeSupportedResponse, error) {
	response := new(IMousegetRelativeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetMultiTouchSupported(request *IMousegetMultiTouchSupported) (*IMousegetMultiTouchSupportedResponse, error) {
	response := new(IMousegetMultiTouchSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetNeedsHostCursor(request *IMousegetNeedsHostCursor) (*IMousegetNeedsHostCursorResponse, error) {
	response := new(IMousegetNeedsHostCursorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousegetEventSource(request *IMousegetEventSource) (*IMousegetEventSourceResponse, error) {
	response := new(IMousegetEventSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEvent(request *IMouseputMouseEvent) (*IMouseputMouseEventResponse, error) {
	response := new(IMouseputMouseEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputMouseEventAbsolute(request *IMouseputMouseEventAbsolute) (*IMouseputMouseEventAbsoluteResponse, error) {
	response := new(IMouseputMouseEventAbsoluteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputEventMultiTouch(request *IMouseputEventMultiTouch) (*IMouseputEventMultiTouchResponse, error) {
	response := new(IMouseputEventMultiTouchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseputEventMultiTouchString(request *IMouseputEventMultiTouchString) (*IMouseputEventMultiTouchStringResponse, error) {
	response := new(IMouseputEventMultiTouchStringResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetWidth(request *IFramebuffergetWidth) (*IFramebuffergetWidthResponse, error) {
	response := new(IFramebuffergetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetHeight(request *IFramebuffergetHeight) (*IFramebuffergetHeightResponse, error) {
	response := new(IFramebuffergetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetBitsPerPixel(request *IFramebuffergetBitsPerPixel) (*IFramebuffergetBitsPerPixelResponse, error) {
	response := new(IFramebuffergetBitsPerPixelResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetBytesPerLine(request *IFramebuffergetBytesPerLine) (*IFramebuffergetBytesPerLineResponse, error) {
	response := new(IFramebuffergetBytesPerLineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetPixelFormat(request *IFramebuffergetPixelFormat) (*IFramebuffergetPixelFormatResponse, error) {
	response := new(IFramebuffergetPixelFormatResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetUsesGuestVRAM(request *IFramebuffergetUsesGuestVRAM) (*IFramebuffergetUsesGuestVRAMResponse, error) {
	response := new(IFramebuffergetUsesGuestVRAMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetHeightReduction(request *IFramebuffergetHeightReduction) (*IFramebuffergetHeightReductionResponse, error) {
	response := new(IFramebuffergetHeightReductionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffergetOverlay(request *IFramebuffergetOverlay) (*IFramebuffergetOverlayResponse, error) {
	response := new(IFramebuffergetOverlayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebuffervideoModeSupported(request *IFramebuffervideoModeSupported) (*IFramebuffervideoModeSupportedResponse, error) {
	response := new(IFramebuffervideoModeSupportedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetX(request *IFramebufferOverlaygetX) (*IFramebufferOverlaygetXResponse, error) {
	response := new(IFramebufferOverlaygetXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetY(request *IFramebufferOverlaygetY) (*IFramebufferOverlaygetYResponse, error) {
	response := new(IFramebufferOverlaygetYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetVisible(request *IFramebufferOverlaygetVisible) (*IFramebufferOverlaygetVisibleResponse, error) {
	response := new(IFramebufferOverlaygetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaysetVisible(request *IFramebufferOverlaysetVisible) (*IFramebufferOverlaysetVisibleResponse, error) {
	response := new(IFramebufferOverlaysetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaygetAlpha(request *IFramebufferOverlaygetAlpha) (*IFramebufferOverlaygetAlphaResponse, error) {
	response := new(IFramebufferOverlaygetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaysetAlpha(request *IFramebufferOverlaysetAlpha) (*IFramebufferOverlaysetAlphaResponse, error) {
	response := new(IFramebufferOverlaysetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IFramebufferOverlaymove(request *IFramebufferOverlaymove) (*IFramebufferOverlaymoveResponse, error) {
	response := new(IFramebufferOverlaymoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaygetScreenResolution(request *IDisplaygetScreenResolution) (*IDisplaygetScreenResolutionResponse, error) {
	response := new(IDisplaygetScreenResolutionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetFramebuffer(request *IDisplaysetFramebuffer) (*IDisplaysetFramebufferResponse, error) {
	response := new(IDisplaysetFramebufferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaygetFramebuffer(request *IDisplaygetFramebuffer) (*IDisplaygetFramebufferResponse, error) {
	response := new(IDisplaygetFramebufferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetVideoModeHint(request *IDisplaysetVideoModeHint) (*IDisplaysetVideoModeHintResponse, error) {
	response := new(IDisplaysetVideoModeHintResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaysetSeamlessMode(request *IDisplaysetSeamlessMode) (*IDisplaysetSeamlessModeResponse, error) {
	response := new(IDisplaysetSeamlessModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaytakeScreenShotToArray(request *IDisplaytakeScreenShotToArray) (*IDisplaytakeScreenShotToArrayResponse, error) {
	response := new(IDisplaytakeScreenShotToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplaytakeScreenShotPNGToArray(request *IDisplaytakeScreenShotPNGToArray) (*IDisplaytakeScreenShotPNGToArrayResponse, error) {
	response := new(IDisplaytakeScreenShotPNGToArrayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayinvalidateAndUpdate(request *IDisplayinvalidateAndUpdate) (*IDisplayinvalidateAndUpdateResponse, error) {
	response := new(IDisplayinvalidateAndUpdateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayresizeCompleted(request *IDisplayresizeCompleted) (*IDisplayresizeCompletedResponse, error) {
	response := new(IDisplayresizeCompletedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDisplayviewportChanged(request *IDisplayviewportChanged) (*IDisplayviewportChangedResponse, error) {
	response := new(IDisplayviewportChangedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAdapterType(request *INetworkAdaptergetAdapterType) (*INetworkAdaptergetAdapterTypeResponse, error) {
	response := new(INetworkAdaptergetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetAdapterType(request *INetworkAdaptersetAdapterType) (*INetworkAdaptersetAdapterTypeResponse, error) {
	response := new(INetworkAdaptersetAdapterTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetSlot(request *INetworkAdaptergetSlot) (*INetworkAdaptergetSlotResponse, error) {
	response := new(INetworkAdaptergetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetEnabled(request *INetworkAdaptergetEnabled) (*INetworkAdaptergetEnabledResponse, error) {
	response := new(INetworkAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetEnabled(request *INetworkAdaptersetEnabled) (*INetworkAdaptersetEnabledResponse, error) {
	response := new(INetworkAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetMACAddress(request *INetworkAdaptergetMACAddress) (*INetworkAdaptergetMACAddressResponse, error) {
	response := new(INetworkAdaptergetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetMACAddress(request *INetworkAdaptersetMACAddress) (*INetworkAdaptersetMACAddressResponse, error) {
	response := new(INetworkAdaptersetMACAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetAttachmentType(request *INetworkAdaptergetAttachmentType) (*INetworkAdaptergetAttachmentTypeResponse, error) {
	response := new(INetworkAdaptergetAttachmentTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetAttachmentType(request *INetworkAdaptersetAttachmentType) (*INetworkAdaptersetAttachmentTypeResponse, error) {
	response := new(INetworkAdaptersetAttachmentTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBridgedInterface(request *INetworkAdaptergetBridgedInterface) (*INetworkAdaptergetBridgedInterfaceResponse, error) {
	response := new(INetworkAdaptergetBridgedInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBridgedInterface(request *INetworkAdaptersetBridgedInterface) (*INetworkAdaptersetBridgedInterfaceResponse, error) {
	response := new(INetworkAdaptersetBridgedInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetHostOnlyInterface(request *INetworkAdaptergetHostOnlyInterface) (*INetworkAdaptergetHostOnlyInterfaceResponse, error) {
	response := new(INetworkAdaptergetHostOnlyInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetHostOnlyInterface(request *INetworkAdaptersetHostOnlyInterface) (*INetworkAdaptersetHostOnlyInterfaceResponse, error) {
	response := new(INetworkAdaptersetHostOnlyInterfaceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetInternalNetwork(request *INetworkAdaptergetInternalNetwork) (*INetworkAdaptergetInternalNetworkResponse, error) {
	response := new(INetworkAdaptergetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetInternalNetwork(request *INetworkAdaptersetInternalNetwork) (*INetworkAdaptersetInternalNetworkResponse, error) {
	response := new(INetworkAdaptersetInternalNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNATNetwork(request *INetworkAdaptergetNATNetwork) (*INetworkAdaptergetNATNetworkResponse, error) {
	response := new(INetworkAdaptergetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetNATNetwork(request *INetworkAdaptersetNATNetwork) (*INetworkAdaptersetNATNetworkResponse, error) {
	response := new(INetworkAdaptersetNATNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetGenericDriver(request *INetworkAdaptergetGenericDriver) (*INetworkAdaptergetGenericDriverResponse, error) {
	response := new(INetworkAdaptergetGenericDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetGenericDriver(request *INetworkAdaptersetGenericDriver) (*INetworkAdaptersetGenericDriverResponse, error) {
	response := new(INetworkAdaptersetGenericDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetCableConnected(request *INetworkAdaptergetCableConnected) (*INetworkAdaptergetCableConnectedResponse, error) {
	response := new(INetworkAdaptergetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetCableConnected(request *INetworkAdaptersetCableConnected) (*INetworkAdaptersetCableConnectedResponse, error) {
	response := new(INetworkAdaptersetCableConnectedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetLineSpeed(request *INetworkAdaptergetLineSpeed) (*INetworkAdaptergetLineSpeedResponse, error) {
	response := new(INetworkAdaptergetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetLineSpeed(request *INetworkAdaptersetLineSpeed) (*INetworkAdaptersetLineSpeedResponse, error) {
	response := new(INetworkAdaptersetLineSpeedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetPromiscModePolicy(request *INetworkAdaptergetPromiscModePolicy) (*INetworkAdaptergetPromiscModePolicyResponse, error) {
	response := new(INetworkAdaptergetPromiscModePolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetPromiscModePolicy(request *INetworkAdaptersetPromiscModePolicy) (*INetworkAdaptersetPromiscModePolicyResponse, error) {
	response := new(INetworkAdaptersetPromiscModePolicyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceEnabled(request *INetworkAdaptergetTraceEnabled) (*INetworkAdaptergetTraceEnabledResponse, error) {
	response := new(INetworkAdaptergetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceEnabled(request *INetworkAdaptersetTraceEnabled) (*INetworkAdaptersetTraceEnabledResponse, error) {
	response := new(INetworkAdaptersetTraceEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetTraceFile(request *INetworkAdaptergetTraceFile) (*INetworkAdaptergetTraceFileResponse, error) {
	response := new(INetworkAdaptergetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetTraceFile(request *INetworkAdaptersetTraceFile) (*INetworkAdaptersetTraceFileResponse, error) {
	response := new(INetworkAdaptersetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetNATEngine(request *INetworkAdaptergetNATEngine) (*INetworkAdaptergetNATEngineResponse, error) {
	response := new(INetworkAdaptergetNATEngineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBootPriority(request *INetworkAdaptergetBootPriority) (*INetworkAdaptergetBootPriorityResponse, error) {
	response := new(INetworkAdaptergetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBootPriority(request *INetworkAdaptersetBootPriority) (*INetworkAdaptersetBootPriorityResponse, error) {
	response := new(INetworkAdaptersetBootPriorityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetBandwidthGroup(request *INetworkAdaptergetBandwidthGroup) (*INetworkAdaptergetBandwidthGroupResponse, error) {
	response := new(INetworkAdaptergetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetBandwidthGroup(request *INetworkAdaptersetBandwidthGroup) (*INetworkAdaptersetBandwidthGroupResponse, error) {
	response := new(INetworkAdaptersetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetProperty(request *INetworkAdaptergetProperty) (*INetworkAdaptergetPropertyResponse, error) {
	response := new(INetworkAdaptergetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptersetProperty(request *INetworkAdaptersetProperty) (*INetworkAdaptersetPropertyResponse, error) {
	response := new(INetworkAdaptersetPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdaptergetProperties(request *INetworkAdaptergetProperties) (*INetworkAdaptergetPropertiesResponse, error) {
	response := new(INetworkAdaptergetPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetSlot(request *ISerialPortgetSlot) (*ISerialPortgetSlotResponse, error) {
	response := new(ISerialPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetEnabled(request *ISerialPortgetEnabled) (*ISerialPortgetEnabledResponse, error) {
	response := new(ISerialPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetEnabled(request *ISerialPortsetEnabled) (*ISerialPortsetEnabledResponse, error) {
	response := new(ISerialPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIOBase(request *ISerialPortgetIOBase) (*ISerialPortgetIOBaseResponse, error) {
	response := new(ISerialPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIOBase(request *ISerialPortsetIOBase) (*ISerialPortsetIOBaseResponse, error) {
	response := new(ISerialPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetIRQ(request *ISerialPortgetIRQ) (*ISerialPortgetIRQResponse, error) {
	response := new(ISerialPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetIRQ(request *ISerialPortsetIRQ) (*ISerialPortsetIRQResponse, error) {
	response := new(ISerialPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetHostMode(request *ISerialPortgetHostMode) (*ISerialPortgetHostModeResponse, error) {
	response := new(ISerialPortgetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetHostMode(request *ISerialPortsetHostMode) (*ISerialPortsetHostModeResponse, error) {
	response := new(ISerialPortsetHostModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetServer(request *ISerialPortgetServer) (*ISerialPortgetServerResponse, error) {
	response := new(ISerialPortgetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetServer(request *ISerialPortsetServer) (*ISerialPortsetServerResponse, error) {
	response := new(ISerialPortsetServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortgetPath(request *ISerialPortgetPath) (*ISerialPortgetPathResponse, error) {
	response := new(ISerialPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortsetPath(request *ISerialPortsetPath) (*ISerialPortsetPathResponse, error) {
	response := new(ISerialPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetSlot(request *IParallelPortgetSlot) (*IParallelPortgetSlotResponse, error) {
	response := new(IParallelPortgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetEnabled(request *IParallelPortgetEnabled) (*IParallelPortgetEnabledResponse, error) {
	response := new(IParallelPortgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetEnabled(request *IParallelPortsetEnabled) (*IParallelPortsetEnabledResponse, error) {
	response := new(IParallelPortsetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIOBase(request *IParallelPortgetIOBase) (*IParallelPortgetIOBaseResponse, error) {
	response := new(IParallelPortgetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIOBase(request *IParallelPortsetIOBase) (*IParallelPortsetIOBaseResponse, error) {
	response := new(IParallelPortsetIOBaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetIRQ(request *IParallelPortgetIRQ) (*IParallelPortgetIRQResponse, error) {
	response := new(IParallelPortgetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetIRQ(request *IParallelPortsetIRQ) (*IParallelPortsetIRQResponse, error) {
	response := new(IParallelPortsetIRQResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortgetPath(request *IParallelPortgetPath) (*IParallelPortgetPathResponse, error) {
	response := new(IParallelPortgetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortsetPath(request *IParallelPortsetPath) (*IParallelPortsetPathResponse, error) {
	response := new(IParallelPortsetPathResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetSingleStep(request *IMachineDebuggergetSingleStep) (*IMachineDebuggergetSingleStepResponse, error) {
	response := new(IMachineDebuggergetSingleStepResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetSingleStep(request *IMachineDebuggersetSingleStep) (*IMachineDebuggersetSingleStepResponse, error) {
	response := new(IMachineDebuggersetSingleStepResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRecompileUser(request *IMachineDebuggergetRecompileUser) (*IMachineDebuggergetRecompileUserResponse, error) {
	response := new(IMachineDebuggergetRecompileUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRecompileUser(request *IMachineDebuggersetRecompileUser) (*IMachineDebuggersetRecompileUserResponse, error) {
	response := new(IMachineDebuggersetRecompileUserResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRecompileSupervisor(request *IMachineDebuggergetRecompileSupervisor) (*IMachineDebuggergetRecompileSupervisorResponse, error) {
	response := new(IMachineDebuggergetRecompileSupervisorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRecompileSupervisor(request *IMachineDebuggersetRecompileSupervisor) (*IMachineDebuggersetRecompileSupervisorResponse, error) {
	response := new(IMachineDebuggersetRecompileSupervisorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetExecuteAllInIEM(request *IMachineDebuggergetExecuteAllInIEM) (*IMachineDebuggergetExecuteAllInIEMResponse, error) {
	response := new(IMachineDebuggergetExecuteAllInIEMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetExecuteAllInIEM(request *IMachineDebuggersetExecuteAllInIEM) (*IMachineDebuggersetExecuteAllInIEMResponse, error) {
	response := new(IMachineDebuggersetExecuteAllInIEMResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetPATMEnabled(request *IMachineDebuggergetPATMEnabled) (*IMachineDebuggergetPATMEnabledResponse, error) {
	response := new(IMachineDebuggergetPATMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetPATMEnabled(request *IMachineDebuggersetPATMEnabled) (*IMachineDebuggersetPATMEnabledResponse, error) {
	response := new(IMachineDebuggersetPATMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetCSAMEnabled(request *IMachineDebuggergetCSAMEnabled) (*IMachineDebuggergetCSAMEnabledResponse, error) {
	response := new(IMachineDebuggergetCSAMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetCSAMEnabled(request *IMachineDebuggersetCSAMEnabled) (*IMachineDebuggersetCSAMEnabledResponse, error) {
	response := new(IMachineDebuggersetCSAMEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogEnabled(request *IMachineDebuggergetLogEnabled) (*IMachineDebuggergetLogEnabledResponse, error) {
	response := new(IMachineDebuggergetLogEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetLogEnabled(request *IMachineDebuggersetLogEnabled) (*IMachineDebuggersetLogEnabledResponse, error) {
	response := new(IMachineDebuggersetLogEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgFlags(request *IMachineDebuggergetLogDbgFlags) (*IMachineDebuggergetLogDbgFlagsResponse, error) {
	response := new(IMachineDebuggergetLogDbgFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgGroups(request *IMachineDebuggergetLogDbgGroups) (*IMachineDebuggergetLogDbgGroupsResponse, error) {
	response := new(IMachineDebuggergetLogDbgGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogDbgDestinations(request *IMachineDebuggergetLogDbgDestinations) (*IMachineDebuggergetLogDbgDestinationsResponse, error) {
	response := new(IMachineDebuggergetLogDbgDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelFlags(request *IMachineDebuggergetLogRelFlags) (*IMachineDebuggergetLogRelFlagsResponse, error) {
	response := new(IMachineDebuggergetLogRelFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelGroups(request *IMachineDebuggergetLogRelGroups) (*IMachineDebuggergetLogRelGroupsResponse, error) {
	response := new(IMachineDebuggergetLogRelGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetLogRelDestinations(request *IMachineDebuggergetLogRelDestinations) (*IMachineDebuggergetLogRelDestinationsResponse, error) {
	response := new(IMachineDebuggergetLogRelDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExEnabled(request *IMachineDebuggergetHWVirtExEnabled) (*IMachineDebuggergetHWVirtExEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExNestedPagingEnabled(request *IMachineDebuggergetHWVirtExNestedPagingEnabled) (*IMachineDebuggergetHWVirtExNestedPagingEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExNestedPagingEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExVPIDEnabled(request *IMachineDebuggergetHWVirtExVPIDEnabled) (*IMachineDebuggergetHWVirtExVPIDEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExVPIDEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetHWVirtExUXEnabled(request *IMachineDebuggergetHWVirtExUXEnabled) (*IMachineDebuggergetHWVirtExUXEnabledResponse, error) {
	response := new(IMachineDebuggergetHWVirtExUXEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetOSName(request *IMachineDebuggergetOSName) (*IMachineDebuggergetOSNameResponse, error) {
	response := new(IMachineDebuggergetOSNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetOSVersion(request *IMachineDebuggergetOSVersion) (*IMachineDebuggergetOSVersionResponse, error) {
	response := new(IMachineDebuggergetOSVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetPAEEnabled(request *IMachineDebuggergetPAEEnabled) (*IMachineDebuggergetPAEEnabledResponse, error) {
	response := new(IMachineDebuggergetPAEEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetVirtualTimeRate(request *IMachineDebuggergetVirtualTimeRate) (*IMachineDebuggergetVirtualTimeRateResponse, error) {
	response := new(IMachineDebuggergetVirtualTimeRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetVirtualTimeRate(request *IMachineDebuggersetVirtualTimeRate) (*IMachineDebuggersetVirtualTimeRateResponse, error) {
	response := new(IMachineDebuggersetVirtualTimeRateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpGuestCore(request *IMachineDebuggerdumpGuestCore) (*IMachineDebuggerdumpGuestCoreResponse, error) {
	response := new(IMachineDebuggerdumpGuestCoreResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpHostProcessCore(request *IMachineDebuggerdumpHostProcessCore) (*IMachineDebuggerdumpHostProcessCoreResponse, error) {
	response := new(IMachineDebuggerdumpHostProcessCoreResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerinfo(request *IMachineDebuggerinfo) (*IMachineDebuggerinfoResponse, error) {
	response := new(IMachineDebuggerinfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerinjectNMI(request *IMachineDebuggerinjectNMI) (*IMachineDebuggerinjectNMIResponse, error) {
	response := new(IMachineDebuggerinjectNMIResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogGroups(request *IMachineDebuggermodifyLogGroups) (*IMachineDebuggermodifyLogGroupsResponse, error) {
	response := new(IMachineDebuggermodifyLogGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogFlags(request *IMachineDebuggermodifyLogFlags) (*IMachineDebuggermodifyLogFlagsResponse, error) {
	response := new(IMachineDebuggermodifyLogFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggermodifyLogDestinations(request *IMachineDebuggermodifyLogDestinations) (*IMachineDebuggermodifyLogDestinationsResponse, error) {
	response := new(IMachineDebuggermodifyLogDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerreadPhysicalMemory(request *IMachineDebuggerreadPhysicalMemory) (*IMachineDebuggerreadPhysicalMemoryResponse, error) {
	response := new(IMachineDebuggerreadPhysicalMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerwritePhysicalMemory(request *IMachineDebuggerwritePhysicalMemory) (*IMachineDebuggerwritePhysicalMemoryResponse, error) {
	response := new(IMachineDebuggerwritePhysicalMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerreadVirtualMemory(request *IMachineDebuggerreadVirtualMemory) (*IMachineDebuggerreadVirtualMemoryResponse, error) {
	response := new(IMachineDebuggerreadVirtualMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerwriteVirtualMemory(request *IMachineDebuggerwriteVirtualMemory) (*IMachineDebuggerwriteVirtualMemoryResponse, error) {
	response := new(IMachineDebuggerwriteVirtualMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdetectOS(request *IMachineDebuggerdetectOS) (*IMachineDebuggerdetectOSResponse, error) {
	response := new(IMachineDebuggerdetectOSResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRegister(request *IMachineDebuggergetRegister) (*IMachineDebuggergetRegisterResponse, error) {
	response := new(IMachineDebuggergetRegisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetRegisters(request *IMachineDebuggergetRegisters) (*IMachineDebuggergetRegistersResponse, error) {
	response := new(IMachineDebuggergetRegistersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRegister(request *IMachineDebuggersetRegister) (*IMachineDebuggersetRegisterResponse, error) {
	response := new(IMachineDebuggersetRegisterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggersetRegisters(request *IMachineDebuggersetRegisters) (*IMachineDebuggersetRegistersResponse, error) {
	response := new(IMachineDebuggersetRegistersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpGuestStack(request *IMachineDebuggerdumpGuestStack) (*IMachineDebuggerdumpGuestStackResponse, error) {
	response := new(IMachineDebuggerdumpGuestStackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerresetStats(request *IMachineDebuggerresetStats) (*IMachineDebuggerresetStatsResponse, error) {
	response := new(IMachineDebuggerresetStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggerdumpStats(request *IMachineDebuggerdumpStats) (*IMachineDebuggerdumpStatsResponse, error) {
	response := new(IMachineDebuggerdumpStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDebuggergetStats(request *IMachineDebuggergetStats) (*IMachineDebuggergetStatsResponse, error) {
	response := new(IMachineDebuggergetStatsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersgetDeviceFilters(request *IUSBDeviceFiltersgetDeviceFilters) (*IUSBDeviceFiltersgetDeviceFiltersResponse, error) {
	response := new(IUSBDeviceFiltersgetDeviceFiltersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFilterscreateDeviceFilter(request *IUSBDeviceFilterscreateDeviceFilter) (*IUSBDeviceFilterscreateDeviceFilterResponse, error) {
	response := new(IUSBDeviceFilterscreateDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersinsertDeviceFilter(request *IUSBDeviceFiltersinsertDeviceFilter) (*IUSBDeviceFiltersinsertDeviceFilterResponse, error) {
	response := new(IUSBDeviceFiltersinsertDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersremoveDeviceFilter(request *IUSBDeviceFiltersremoveDeviceFilter) (*IUSBDeviceFiltersremoveDeviceFilterResponse, error) {
	response := new(IUSBDeviceFiltersremoveDeviceFilterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetName(request *IUSBControllergetName) (*IUSBControllergetNameResponse, error) {
	response := new(IUSBControllergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetType(request *IUSBControllergetType) (*IUSBControllergetTypeResponse, error) {
	response := new(IUSBControllergetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBControllergetUSBStandard(request *IUSBControllergetUSBStandard) (*IUSBControllergetUSBStandardResponse, error) {
	response := new(IUSBControllergetUSBStandardResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetId(request *IUSBDevicegetId) (*IUSBDevicegetIdResponse, error) {
	response := new(IUSBDevicegetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVendorId(request *IUSBDevicegetVendorId) (*IUSBDevicegetVendorIdResponse, error) {
	response := new(IUSBDevicegetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProductId(request *IUSBDevicegetProductId) (*IUSBDevicegetProductIdResponse, error) {
	response := new(IUSBDevicegetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRevision(request *IUSBDevicegetRevision) (*IUSBDevicegetRevisionResponse, error) {
	response := new(IUSBDevicegetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetManufacturer(request *IUSBDevicegetManufacturer) (*IUSBDevicegetManufacturerResponse, error) {
	response := new(IUSBDevicegetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetProduct(request *IUSBDevicegetProduct) (*IUSBDevicegetProductResponse, error) {
	response := new(IUSBDevicegetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetSerialNumber(request *IUSBDevicegetSerialNumber) (*IUSBDevicegetSerialNumberResponse, error) {
	response := new(IUSBDevicegetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetAddress(request *IUSBDevicegetAddress) (*IUSBDevicegetAddressResponse, error) {
	response := new(IUSBDevicegetAddressResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPort(request *IUSBDevicegetPort) (*IUSBDevicegetPortResponse, error) {
	response := new(IUSBDevicegetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetVersion(request *IUSBDevicegetVersion) (*IUSBDevicegetVersionResponse, error) {
	response := new(IUSBDevicegetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetPortVersion(request *IUSBDevicegetPortVersion) (*IUSBDevicegetPortVersionResponse, error) {
	response := new(IUSBDevicegetPortVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDevicegetRemote(request *IUSBDevicegetRemote) (*IUSBDevicegetRemoteResponse, error) {
	response := new(IUSBDevicegetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetName(request *IUSBDeviceFiltergetName) (*IUSBDeviceFiltergetNameResponse, error) {
	response := new(IUSBDeviceFiltergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetName(request *IUSBDeviceFiltersetName) (*IUSBDeviceFiltersetNameResponse, error) {
	response := new(IUSBDeviceFiltersetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetActive(request *IUSBDeviceFiltergetActive) (*IUSBDeviceFiltergetActiveResponse, error) {
	response := new(IUSBDeviceFiltergetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetActive(request *IUSBDeviceFiltersetActive) (*IUSBDeviceFiltersetActiveResponse, error) {
	response := new(IUSBDeviceFiltersetActiveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetVendorId(request *IUSBDeviceFiltergetVendorId) (*IUSBDeviceFiltergetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltergetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetVendorId(request *IUSBDeviceFiltersetVendorId) (*IUSBDeviceFiltersetVendorIdResponse, error) {
	response := new(IUSBDeviceFiltersetVendorIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProductId(request *IUSBDeviceFiltergetProductId) (*IUSBDeviceFiltergetProductIdResponse, error) {
	response := new(IUSBDeviceFiltergetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProductId(request *IUSBDeviceFiltersetProductId) (*IUSBDeviceFiltersetProductIdResponse, error) {
	response := new(IUSBDeviceFiltersetProductIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRevision(request *IUSBDeviceFiltergetRevision) (*IUSBDeviceFiltergetRevisionResponse, error) {
	response := new(IUSBDeviceFiltergetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRevision(request *IUSBDeviceFiltersetRevision) (*IUSBDeviceFiltersetRevisionResponse, error) {
	response := new(IUSBDeviceFiltersetRevisionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetManufacturer(request *IUSBDeviceFiltergetManufacturer) (*IUSBDeviceFiltergetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltergetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetManufacturer(request *IUSBDeviceFiltersetManufacturer) (*IUSBDeviceFiltersetManufacturerResponse, error) {
	response := new(IUSBDeviceFiltersetManufacturerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetProduct(request *IUSBDeviceFiltergetProduct) (*IUSBDeviceFiltergetProductResponse, error) {
	response := new(IUSBDeviceFiltergetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetProduct(request *IUSBDeviceFiltersetProduct) (*IUSBDeviceFiltersetProductResponse, error) {
	response := new(IUSBDeviceFiltersetProductResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetSerialNumber(request *IUSBDeviceFiltergetSerialNumber) (*IUSBDeviceFiltergetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltergetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetSerialNumber(request *IUSBDeviceFiltersetSerialNumber) (*IUSBDeviceFiltersetSerialNumberResponse, error) {
	response := new(IUSBDeviceFiltersetSerialNumberResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetPort(request *IUSBDeviceFiltergetPort) (*IUSBDeviceFiltergetPortResponse, error) {
	response := new(IUSBDeviceFiltergetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetPort(request *IUSBDeviceFiltersetPort) (*IUSBDeviceFiltersetPortResponse, error) {
	response := new(IUSBDeviceFiltersetPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetRemote(request *IUSBDeviceFiltergetRemote) (*IUSBDeviceFiltergetRemoteResponse, error) {
	response := new(IUSBDeviceFiltergetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetRemote(request *IUSBDeviceFiltersetRemote) (*IUSBDeviceFiltersetRemoteResponse, error) {
	response := new(IUSBDeviceFiltersetRemoteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltergetMaskedInterfaces(request *IUSBDeviceFiltergetMaskedInterfaces) (*IUSBDeviceFiltergetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltergetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceFiltersetMaskedInterfaces(request *IUSBDeviceFiltersetMaskedInterfaces) (*IUSBDeviceFiltersetMaskedInterfacesResponse, error) {
	response := new(IUSBDeviceFiltersetMaskedInterfacesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDevicegetState(request *IHostUSBDevicegetState) (*IHostUSBDevicegetStateResponse, error) {
	response := new(IHostUSBDevicegetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltergetAction(request *IHostUSBDeviceFiltergetAction) (*IHostUSBDeviceFiltergetActionResponse, error) {
	response := new(IHostUSBDeviceFiltergetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostUSBDeviceFiltersetAction(request *IHostUSBDeviceFiltersetAction) (*IHostUSBDeviceFiltersetActionResponse, error) {
	response := new(IHostUSBDeviceFiltersetActionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetEnabled(request *IAudioAdaptergetEnabled) (*IAudioAdaptergetEnabledResponse, error) {
	response := new(IAudioAdaptergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetEnabled(request *IAudioAdaptersetEnabled) (*IAudioAdaptersetEnabledResponse, error) {
	response := new(IAudioAdaptersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioController(request *IAudioAdaptergetAudioController) (*IAudioAdaptergetAudioControllerResponse, error) {
	response := new(IAudioAdaptergetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioController(request *IAudioAdaptersetAudioController) (*IAudioAdaptersetAudioControllerResponse, error) {
	response := new(IAudioAdaptersetAudioControllerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptergetAudioDriver(request *IAudioAdaptergetAudioDriver) (*IAudioAdaptergetAudioDriverResponse, error) {
	response := new(IAudioAdaptergetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IAudioAdaptersetAudioDriver(request *IAudioAdaptersetAudioDriver) (*IAudioAdaptersetAudioDriverResponse, error) {
	response := new(IAudioAdaptersetAudioDriverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetEnabled(request *IVRDEServergetEnabled) (*IVRDEServergetEnabledResponse, error) {
	response := new(IVRDEServergetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetEnabled(request *IVRDEServersetEnabled) (*IVRDEServersetEnabledResponse, error) {
	response := new(IVRDEServersetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthType(request *IVRDEServergetAuthType) (*IVRDEServergetAuthTypeResponse, error) {
	response := new(IVRDEServergetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthType(request *IVRDEServersetAuthType) (*IVRDEServersetAuthTypeResponse, error) {
	response := new(IVRDEServersetAuthTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthTimeout(request *IVRDEServergetAuthTimeout) (*IVRDEServergetAuthTimeoutResponse, error) {
	response := new(IVRDEServergetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthTimeout(request *IVRDEServersetAuthTimeout) (*IVRDEServersetAuthTimeoutResponse, error) {
	response := new(IVRDEServersetAuthTimeoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAllowMultiConnection(request *IVRDEServergetAllowMultiConnection) (*IVRDEServergetAllowMultiConnectionResponse, error) {
	response := new(IVRDEServergetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAllowMultiConnection(request *IVRDEServersetAllowMultiConnection) (*IVRDEServersetAllowMultiConnectionResponse, error) {
	response := new(IVRDEServersetAllowMultiConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetReuseSingleConnection(request *IVRDEServergetReuseSingleConnection) (*IVRDEServergetReuseSingleConnectionResponse, error) {
	response := new(IVRDEServergetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetReuseSingleConnection(request *IVRDEServersetReuseSingleConnection) (*IVRDEServersetReuseSingleConnectionResponse, error) {
	response := new(IVRDEServersetReuseSingleConnectionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEExtPack(request *IVRDEServergetVRDEExtPack) (*IVRDEServergetVRDEExtPackResponse, error) {
	response := new(IVRDEServergetVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetVRDEExtPack(request *IVRDEServersetVRDEExtPack) (*IVRDEServersetVRDEExtPackResponse, error) {
	response := new(IVRDEServersetVRDEExtPackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetAuthLibrary(request *IVRDEServergetAuthLibrary) (*IVRDEServergetAuthLibraryResponse, error) {
	response := new(IVRDEServergetAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetAuthLibrary(request *IVRDEServersetAuthLibrary) (*IVRDEServersetAuthLibraryResponse, error) {
	response := new(IVRDEServersetAuthLibraryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEProperties(request *IVRDEServergetVRDEProperties) (*IVRDEServergetVRDEPropertiesResponse, error) {
	response := new(IVRDEServergetVRDEPropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServersetVRDEProperty(request *IVRDEServersetVRDEProperty) (*IVRDEServersetVRDEPropertyResponse, error) {
	response := new(IVRDEServersetVRDEPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVRDEServergetVRDEProperty(request *IVRDEServergetVRDEProperty) (*IVRDEServergetVRDEPropertyResponse, error) {
	response := new(IVRDEServergetVRDEPropertyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetState(request *ISessiongetState) (*ISessiongetStateResponse, error) {
	response := new(ISessiongetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetType(request *ISessiongetType) (*ISessiongetTypeResponse, error) {
	response := new(ISessiongetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetMachine(request *ISessiongetMachine) (*ISessiongetMachineResponse, error) {
	response := new(ISessiongetMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessiongetConsole(request *ISessiongetConsole) (*ISessiongetConsoleResponse, error) {
	response := new(ISessiongetConsoleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionunlockMachine(request *ISessionunlockMachine) (*ISessionunlockMachineResponse, error) {
	response := new(ISessionunlockMachineResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetName(request *IStorageControllergetName) (*IStorageControllergetNameResponse, error) {
	response := new(IStorageControllergetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxDevicesPerPortCount(request *IStorageControllergetMaxDevicesPerPortCount) (*IStorageControllergetMaxDevicesPerPortCountResponse, error) {
	response := new(IStorageControllergetMaxDevicesPerPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMinPortCount(request *IStorageControllergetMinPortCount) (*IStorageControllergetMinPortCountResponse, error) {
	response := new(IStorageControllergetMinPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetMaxPortCount(request *IStorageControllergetMaxPortCount) (*IStorageControllergetMaxPortCountResponse, error) {
	response := new(IStorageControllergetMaxPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetInstance(request *IStorageControllergetInstance) (*IStorageControllergetInstanceResponse, error) {
	response := new(IStorageControllergetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetInstance(request *IStorageControllersetInstance) (*IStorageControllersetInstanceResponse, error) {
	response := new(IStorageControllersetInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetPortCount(request *IStorageControllergetPortCount) (*IStorageControllergetPortCountResponse, error) {
	response := new(IStorageControllergetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetPortCount(request *IStorageControllersetPortCount) (*IStorageControllersetPortCountResponse, error) {
	response := new(IStorageControllersetPortCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetBus(request *IStorageControllergetBus) (*IStorageControllergetBusResponse, error) {
	response := new(IStorageControllergetBusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetControllerType(request *IStorageControllergetControllerType) (*IStorageControllergetControllerTypeResponse, error) {
	response := new(IStorageControllergetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetControllerType(request *IStorageControllersetControllerType) (*IStorageControllersetControllerTypeResponse, error) {
	response := new(IStorageControllersetControllerTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetUseHostIOCache(request *IStorageControllergetUseHostIOCache) (*IStorageControllergetUseHostIOCacheResponse, error) {
	response := new(IStorageControllergetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllersetUseHostIOCache(request *IStorageControllersetUseHostIOCache) (*IStorageControllersetUseHostIOCacheResponse, error) {
	response := new(IStorageControllersetUseHostIOCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageControllergetBootable(request *IStorageControllergetBootable) (*IStorageControllergetBootableResponse, error) {
	response := new(IStorageControllergetBootableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefgetInterfaceName(request *IManagedObjectRefgetInterfaceName) (*IManagedObjectRefgetInterfaceNameResponse, error) {
	response := new(IManagedObjectRefgetInterfaceNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IManagedObjectRefrelease(request *IManagedObjectRefrelease) (*IManagedObjectRefreleaseResponse, error) {
	response := new(IManagedObjectRefreleaseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogon(request *IWebsessionManagerlogon) (*IWebsessionManagerlogonResponse, error) {
	response := new(IWebsessionManagerlogonResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagergetSessionObject(request *IWebsessionManagergetSessionObject) (*IWebsessionManagergetSessionObjectResponse, error) {
	response := new(IWebsessionManagergetSessionObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IWebsessionManagerlogoff(request *IWebsessionManagerlogoff) (*IWebsessionManagerlogoffResponse, error) {
	response := new(IWebsessionManagerlogoffResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMetricName(request *IPerformanceMetricgetMetricName) (*IPerformanceMetricgetMetricNameResponse, error) {
	response := new(IPerformanceMetricgetMetricNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetObject(request *IPerformanceMetricgetObject) (*IPerformanceMetricgetObjectResponse, error) {
	response := new(IPerformanceMetricgetObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetDescription(request *IPerformanceMetricgetDescription) (*IPerformanceMetricgetDescriptionResponse, error) {
	response := new(IPerformanceMetricgetDescriptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetPeriod(request *IPerformanceMetricgetPeriod) (*IPerformanceMetricgetPeriodResponse, error) {
	response := new(IPerformanceMetricgetPeriodResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetCount(request *IPerformanceMetricgetCount) (*IPerformanceMetricgetCountResponse, error) {
	response := new(IPerformanceMetricgetCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetUnit(request *IPerformanceMetricgetUnit) (*IPerformanceMetricgetUnitResponse, error) {
	response := new(IPerformanceMetricgetUnitResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMinimumValue(request *IPerformanceMetricgetMinimumValue) (*IPerformanceMetricgetMinimumValueResponse, error) {
	response := new(IPerformanceMetricgetMinimumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceMetricgetMaximumValue(request *IPerformanceMetricgetMaximumValue) (*IPerformanceMetricgetMaximumValueResponse, error) {
	response := new(IPerformanceMetricgetMaximumValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetricNames(request *IPerformanceCollectorgetMetricNames) (*IPerformanceCollectorgetMetricNamesResponse, error) {
	response := new(IPerformanceCollectorgetMetricNamesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorgetMetrics(request *IPerformanceCollectorgetMetrics) (*IPerformanceCollectorgetMetricsResponse, error) {
	response := new(IPerformanceCollectorgetMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorsetupMetrics(request *IPerformanceCollectorsetupMetrics) (*IPerformanceCollectorsetupMetricsResponse, error) {
	response := new(IPerformanceCollectorsetupMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorenableMetrics(request *IPerformanceCollectorenableMetrics) (*IPerformanceCollectorenableMetricsResponse, error) {
	response := new(IPerformanceCollectorenableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectordisableMetrics(request *IPerformanceCollectordisableMetrics) (*IPerformanceCollectordisableMetricsResponse, error) {
	response := new(IPerformanceCollectordisableMetricsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IPerformanceCollectorqueryMetricsData(request *IPerformanceCollectorqueryMetricsData) (*IPerformanceCollectorqueryMetricsDataResponse, error) {
	response := new(IPerformanceCollectorqueryMetricsDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetwork(request *INATEnginegetNetwork) (*INATEnginegetNetworkResponse, error) {
	response := new(INATEnginegetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetwork(request *INATEnginesetNetwork) (*INATEnginesetNetworkResponse, error) {
	response := new(INATEnginesetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetHostIP(request *INATEnginegetHostIP) (*INATEnginegetHostIPResponse, error) {
	response := new(INATEnginegetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetHostIP(request *INATEnginesetHostIP) (*INATEnginesetHostIPResponse, error) {
	response := new(INATEnginesetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPPrefix(request *INATEnginegetTFTPPrefix) (*INATEnginegetTFTPPrefixResponse, error) {
	response := new(INATEnginegetTFTPPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPPrefix(request *INATEnginesetTFTPPrefix) (*INATEnginesetTFTPPrefixResponse, error) {
	response := new(INATEnginesetTFTPPrefixResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPBootFile(request *INATEnginegetTFTPBootFile) (*INATEnginegetTFTPBootFileResponse, error) {
	response := new(INATEnginegetTFTPBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPBootFile(request *INATEnginesetTFTPBootFile) (*INATEnginesetTFTPBootFileResponse, error) {
	response := new(INATEnginesetTFTPBootFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetTFTPNextServer(request *INATEnginegetTFTPNextServer) (*INATEnginegetTFTPNextServerResponse, error) {
	response := new(INATEnginegetTFTPNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetTFTPNextServer(request *INATEnginesetTFTPNextServer) (*INATEnginesetTFTPNextServerResponse, error) {
	response := new(INATEnginesetTFTPNextServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetAliasMode(request *INATEnginegetAliasMode) (*INATEnginegetAliasModeResponse, error) {
	response := new(INATEnginegetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetAliasMode(request *INATEnginesetAliasMode) (*INATEnginesetAliasModeResponse, error) {
	response := new(INATEnginesetAliasModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSPassDomain(request *INATEnginegetDNSPassDomain) (*INATEnginegetDNSPassDomainResponse, error) {
	response := new(INATEnginegetDNSPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSPassDomain(request *INATEnginesetDNSPassDomain) (*INATEnginesetDNSPassDomainResponse, error) {
	response := new(INATEnginesetDNSPassDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSProxy(request *INATEnginegetDNSProxy) (*INATEnginegetDNSProxyResponse, error) {
	response := new(INATEnginegetDNSProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSProxy(request *INATEnginesetDNSProxy) (*INATEnginesetDNSProxyResponse, error) {
	response := new(INATEnginesetDNSProxyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetDNSUseHostResolver(request *INATEnginegetDNSUseHostResolver) (*INATEnginegetDNSUseHostResolverResponse, error) {
	response := new(INATEnginegetDNSUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetDNSUseHostResolver(request *INATEnginesetDNSUseHostResolver) (*INATEnginesetDNSUseHostResolverResponse, error) {
	response := new(INATEnginesetDNSUseHostResolverResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetRedirects(request *INATEnginegetRedirects) (*INATEnginegetRedirectsResponse, error) {
	response := new(INATEnginegetRedirectsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginesetNetworkSettings(request *INATEnginesetNetworkSettings) (*INATEnginesetNetworkSettingsResponse, error) {
	response := new(INATEnginesetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEnginegetNetworkSettings(request *INATEnginegetNetworkSettings) (*INATEnginegetNetworkSettingsResponse, error) {
	response := new(INATEnginegetNetworkSettingsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineaddRedirect(request *INATEngineaddRedirect) (*INATEngineaddRedirectResponse, error) {
	response := new(INATEngineaddRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATEngineremoveRedirect(request *INATEngineremoveRedirect) (*INATEngineremoveRedirectResponse, error) {
	response := new(INATEngineremoveRedirectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetName(request *IBandwidthGroupgetName) (*IBandwidthGroupgetNameResponse, error) {
	response := new(IBandwidthGroupgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetType(request *IBandwidthGroupgetType) (*IBandwidthGroupgetTypeResponse, error) {
	response := new(IBandwidthGroupgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetReference(request *IBandwidthGroupgetReference) (*IBandwidthGroupgetReferenceResponse, error) {
	response := new(IBandwidthGroupgetReferenceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupgetMaxBytesPerSec(request *IBandwidthGroupgetMaxBytesPerSec) (*IBandwidthGroupgetMaxBytesPerSecResponse, error) {
	response := new(IBandwidthGroupgetMaxBytesPerSecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupsetMaxBytesPerSec(request *IBandwidthGroupsetMaxBytesPerSec) (*IBandwidthGroupsetMaxBytesPerSecResponse, error) {
	response := new(IBandwidthGroupsetMaxBytesPerSecResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetNumGroups(request *IBandwidthControlgetNumGroups) (*IBandwidthControlgetNumGroupsResponse, error) {
	response := new(IBandwidthControlgetNumGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlcreateBandwidthGroup(request *IBandwidthControlcreateBandwidthGroup) (*IBandwidthControlcreateBandwidthGroupResponse, error) {
	response := new(IBandwidthControlcreateBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControldeleteBandwidthGroup(request *IBandwidthControldeleteBandwidthGroup) (*IBandwidthControldeleteBandwidthGroupResponse, error) {
	response := new(IBandwidthControldeleteBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetBandwidthGroup(request *IBandwidthControlgetBandwidthGroup) (*IBandwidthControlgetBandwidthGroupResponse, error) {
	response := new(IBandwidthControlgetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthControlgetAllBandwidthGroups(request *IBandwidthControlgetAllBandwidthGroups) (*IBandwidthControlgetAllBandwidthGroupsResponse, error) {
	response := new(IBandwidthControlgetAllBandwidthGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcecreateListener(request *IEventSourcecreateListener) (*IEventSourcecreateListenerResponse, error) {
	response := new(IEventSourcecreateListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcecreateAggregator(request *IEventSourcecreateAggregator) (*IEventSourcecreateAggregatorResponse, error) {
	response := new(IEventSourcecreateAggregatorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceregisterListener(request *IEventSourceregisterListener) (*IEventSourceregisterListenerResponse, error) {
	response := new(IEventSourceregisterListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceunregisterListener(request *IEventSourceunregisterListener) (*IEventSourceunregisterListenerResponse, error) {
	response := new(IEventSourceunregisterListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcefireEvent(request *IEventSourcefireEvent) (*IEventSourcefireEventResponse, error) {
	response := new(IEventSourcefireEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourcegetEvent(request *IEventSourcegetEvent) (*IEventSourcegetEventResponse, error) {
	response := new(IEventSourcegetEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceeventProcessed(request *IEventSourceeventProcessed) (*IEventSourceeventProcessedResponse, error) {
	response := new(IEventSourceeventProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventListenerhandleEvent(request *IEventListenerhandleEvent) (*IEventListenerhandleEventResponse, error) {
	response := new(IEventListenerhandleEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetType(request *IEventgetType) (*IEventgetTypeResponse, error) {
	response := new(IEventgetTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetSource(request *IEventgetSource) (*IEventgetSourceResponse, error) {
	response := new(IEventgetSourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventgetWaitable(request *IEventgetWaitable) (*IEventgetWaitableResponse, error) {
	response := new(IEventgetWaitableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventsetProcessed(request *IEventsetProcessed) (*IEventsetProcessedResponse, error) {
	response := new(IEventsetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventwaitProcessed(request *IEventwaitProcessed) (*IEventwaitProcessedResponse, error) {
	response := new(IEventwaitProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IReusableEventgetGeneration(request *IReusableEventgetGeneration) (*IReusableEventgetGenerationResponse, error) {
	response := new(IReusableEventgetGenerationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IReusableEventreuse(request *IReusableEventreuse) (*IReusableEventreuseResponse, error) {
	response := new(IReusableEventreuseResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineEventgetMachineId(request *IMachineEventgetMachineId) (*IMachineEventgetMachineIdResponse, error) {
	response := new(IMachineEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineStateChangedEventgetState(request *IMachineStateChangedEventgetState) (*IMachineStateChangedEventgetStateResponse, error) {
	response := new(IMachineStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineDataChangedEventgetTemporary(request *IMachineDataChangedEventgetTemporary) (*IMachineDataChangedEventgetTemporaryResponse, error) {
	response := new(IMachineDataChangedEventgetTemporaryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetMediumId(request *IMediumRegisteredEventgetMediumId) (*IMediumRegisteredEventgetMediumIdResponse, error) {
	response := new(IMediumRegisteredEventgetMediumIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetMediumType(request *IMediumRegisteredEventgetMediumType) (*IMediumRegisteredEventgetMediumTypeResponse, error) {
	response := new(IMediumRegisteredEventgetMediumTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumRegisteredEventgetRegistered(request *IMediumRegisteredEventgetRegistered) (*IMediumRegisteredEventgetRegisteredResponse, error) {
	response := new(IMediumRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMachineRegisteredEventgetRegistered(request *IMachineRegisteredEventgetRegistered) (*IMachineRegisteredEventgetRegisteredResponse, error) {
	response := new(IMachineRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISessionStateChangedEventgetState(request *ISessionStateChangedEventgetState) (*ISessionStateChangedEventgetStateResponse, error) {
	response := new(ISessionStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetName(request *IGuestPropertyChangedEventgetName) (*IGuestPropertyChangedEventgetNameResponse, error) {
	response := new(IGuestPropertyChangedEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetValue(request *IGuestPropertyChangedEventgetValue) (*IGuestPropertyChangedEventgetValueResponse, error) {
	response := new(IGuestPropertyChangedEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestPropertyChangedEventgetFlags(request *IGuestPropertyChangedEventgetFlags) (*IGuestPropertyChangedEventgetFlagsResponse, error) {
	response := new(IGuestPropertyChangedEventgetFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISnapshotEventgetSnapshotId(request *ISnapshotEventgetSnapshotId) (*ISnapshotEventgetSnapshotIdResponse, error) {
	response := new(ISnapshotEventgetSnapshotIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetVisible(request *IMousePointerShapeChangedEventgetVisible) (*IMousePointerShapeChangedEventgetVisibleResponse, error) {
	response := new(IMousePointerShapeChangedEventgetVisibleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetAlpha(request *IMousePointerShapeChangedEventgetAlpha) (*IMousePointerShapeChangedEventgetAlphaResponse, error) {
	response := new(IMousePointerShapeChangedEventgetAlphaResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetXhot(request *IMousePointerShapeChangedEventgetXhot) (*IMousePointerShapeChangedEventgetXhotResponse, error) {
	response := new(IMousePointerShapeChangedEventgetXhotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetYhot(request *IMousePointerShapeChangedEventgetYhot) (*IMousePointerShapeChangedEventgetYhotResponse, error) {
	response := new(IMousePointerShapeChangedEventgetYhotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetWidth(request *IMousePointerShapeChangedEventgetWidth) (*IMousePointerShapeChangedEventgetWidthResponse, error) {
	response := new(IMousePointerShapeChangedEventgetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetHeight(request *IMousePointerShapeChangedEventgetHeight) (*IMousePointerShapeChangedEventgetHeightResponse, error) {
	response := new(IMousePointerShapeChangedEventgetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMousePointerShapeChangedEventgetShape(request *IMousePointerShapeChangedEventgetShape) (*IMousePointerShapeChangedEventgetShapeResponse, error) {
	response := new(IMousePointerShapeChangedEventgetShapeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsAbsolute(request *IMouseCapabilityChangedEventgetSupportsAbsolute) (*IMouseCapabilityChangedEventgetSupportsAbsoluteResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsAbsoluteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsRelative(request *IMouseCapabilityChangedEventgetSupportsRelative) (*IMouseCapabilityChangedEventgetSupportsRelativeResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsRelativeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetSupportsMultiTouch(request *IMouseCapabilityChangedEventgetSupportsMultiTouch) (*IMouseCapabilityChangedEventgetSupportsMultiTouchResponse, error) {
	response := new(IMouseCapabilityChangedEventgetSupportsMultiTouchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMouseCapabilityChangedEventgetNeedsHostCursor(request *IMouseCapabilityChangedEventgetNeedsHostCursor) (*IMouseCapabilityChangedEventgetNeedsHostCursorResponse, error) {
	response := new(IMouseCapabilityChangedEventgetNeedsHostCursorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetNumLock(request *IKeyboardLedsChangedEventgetNumLock) (*IKeyboardLedsChangedEventgetNumLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetNumLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetCapsLock(request *IKeyboardLedsChangedEventgetCapsLock) (*IKeyboardLedsChangedEventgetCapsLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetCapsLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IKeyboardLedsChangedEventgetScrollLock(request *IKeyboardLedsChangedEventgetScrollLock) (*IKeyboardLedsChangedEventgetScrollLockResponse, error) {
	response := new(IKeyboardLedsChangedEventgetScrollLockResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStateChangedEventgetState(request *IStateChangedEventgetState) (*IStateChangedEventgetStateResponse, error) {
	response := new(IStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INetworkAdapterChangedEventgetNetworkAdapter(request *INetworkAdapterChangedEventgetNetworkAdapter) (*INetworkAdapterChangedEventgetNetworkAdapterResponse, error) {
	response := new(INetworkAdapterChangedEventgetNetworkAdapterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISerialPortChangedEventgetSerialPort(request *ISerialPortChangedEventgetSerialPort) (*ISerialPortChangedEventgetSerialPortResponse, error) {
	response := new(ISerialPortChangedEventgetSerialPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IParallelPortChangedEventgetParallelPort(request *IParallelPortChangedEventgetParallelPort) (*IParallelPortChangedEventgetParallelPortResponse, error) {
	response := new(IParallelPortChangedEventgetParallelPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IMediumChangedEventgetMediumAttachment(request *IMediumChangedEventgetMediumAttachment) (*IMediumChangedEventgetMediumAttachmentResponse, error) {
	response := new(IMediumChangedEventgetMediumAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IClipboardModeChangedEventgetClipboardMode(request *IClipboardModeChangedEventgetClipboardMode) (*IClipboardModeChangedEventgetClipboardModeResponse, error) {
	response := new(IClipboardModeChangedEventgetClipboardModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IDragAndDropModeChangedEventgetDragAndDropMode(request *IDragAndDropModeChangedEventgetDragAndDropMode) (*IDragAndDropModeChangedEventgetDragAndDropModeResponse, error) {
	response := new(IDragAndDropModeChangedEventgetDragAndDropModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUChangedEventgetCPU(request *ICPUChangedEventgetCPU) (*ICPUChangedEventgetCPUResponse, error) {
	response := new(ICPUChangedEventgetCPUResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUChangedEventgetAdd(request *ICPUChangedEventgetAdd) (*ICPUChangedEventgetAddResponse, error) {
	response := new(ICPUChangedEventgetAddResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ICPUExecutionCapChangedEventgetExecutionCap(request *ICPUExecutionCapChangedEventgetExecutionCap) (*ICPUExecutionCapChangedEventgetExecutionCapResponse, error) {
	response := new(ICPUExecutionCapChangedEventgetExecutionCapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestKeyboardEventgetScancodes(request *IGuestKeyboardEventgetScancodes) (*IGuestKeyboardEventgetScancodesResponse, error) {
	response := new(IGuestKeyboardEventgetScancodesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetMode(request *IGuestMouseEventgetMode) (*IGuestMouseEventgetModeResponse, error) {
	response := new(IGuestMouseEventgetModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetX(request *IGuestMouseEventgetX) (*IGuestMouseEventgetXResponse, error) {
	response := new(IGuestMouseEventgetXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetY(request *IGuestMouseEventgetY) (*IGuestMouseEventgetYResponse, error) {
	response := new(IGuestMouseEventgetYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetZ(request *IGuestMouseEventgetZ) (*IGuestMouseEventgetZResponse, error) {
	response := new(IGuestMouseEventgetZResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetW(request *IGuestMouseEventgetW) (*IGuestMouseEventgetWResponse, error) {
	response := new(IGuestMouseEventgetWResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMouseEventgetButtons(request *IGuestMouseEventgetButtons) (*IGuestMouseEventgetButtonsResponse, error) {
	response := new(IGuestMouseEventgetButtonsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactCount(request *IGuestMultiTouchEventgetContactCount) (*IGuestMultiTouchEventgetContactCountResponse, error) {
	response := new(IGuestMultiTouchEventgetContactCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetXPositions(request *IGuestMultiTouchEventgetXPositions) (*IGuestMultiTouchEventgetXPositionsResponse, error) {
	response := new(IGuestMultiTouchEventgetXPositionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetYPositions(request *IGuestMultiTouchEventgetYPositions) (*IGuestMultiTouchEventgetYPositionsResponse, error) {
	response := new(IGuestMultiTouchEventgetYPositionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactIds(request *IGuestMultiTouchEventgetContactIds) (*IGuestMultiTouchEventgetContactIdsResponse, error) {
	response := new(IGuestMultiTouchEventgetContactIdsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetContactFlags(request *IGuestMultiTouchEventgetContactFlags) (*IGuestMultiTouchEventgetContactFlagsResponse, error) {
	response := new(IGuestMultiTouchEventgetContactFlagsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMultiTouchEventgetScanTime(request *IGuestMultiTouchEventgetScanTime) (*IGuestMultiTouchEventgetScanTimeResponse, error) {
	response := new(IGuestMultiTouchEventgetScanTimeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionEventgetSession(request *IGuestSessionEventgetSession) (*IGuestSessionEventgetSessionResponse, error) {
	response := new(IGuestSessionEventgetSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetId(request *IGuestSessionStateChangedEventgetId) (*IGuestSessionStateChangedEventgetIdResponse, error) {
	response := new(IGuestSessionStateChangedEventgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetStatus(request *IGuestSessionStateChangedEventgetStatus) (*IGuestSessionStateChangedEventgetStatusResponse, error) {
	response := new(IGuestSessionStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionStateChangedEventgetError(request *IGuestSessionStateChangedEventgetError) (*IGuestSessionStateChangedEventgetErrorResponse, error) {
	response := new(IGuestSessionStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestSessionRegisteredEventgetRegistered(request *IGuestSessionRegisteredEventgetRegistered) (*IGuestSessionRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestSessionRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessEventgetProcess(request *IGuestProcessEventgetProcess) (*IGuestProcessEventgetProcessResponse, error) {
	response := new(IGuestProcessEventgetProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessEventgetPid(request *IGuestProcessEventgetPid) (*IGuestProcessEventgetPidResponse, error) {
	response := new(IGuestProcessEventgetPidResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessRegisteredEventgetRegistered(request *IGuestProcessRegisteredEventgetRegistered) (*IGuestProcessRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestProcessRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessStateChangedEventgetStatus(request *IGuestProcessStateChangedEventgetStatus) (*IGuestProcessStateChangedEventgetStatusResponse, error) {
	response := new(IGuestProcessStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessStateChangedEventgetError(request *IGuestProcessStateChangedEventgetError) (*IGuestProcessStateChangedEventgetErrorResponse, error) {
	response := new(IGuestProcessStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessIOEventgetHandle(request *IGuestProcessIOEventgetHandle) (*IGuestProcessIOEventgetHandleResponse, error) {
	response := new(IGuestProcessIOEventgetHandleResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessIOEventgetProcessed(request *IGuestProcessIOEventgetProcessed) (*IGuestProcessIOEventgetProcessedResponse, error) {
	response := new(IGuestProcessIOEventgetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessInputNotifyEventgetStatus(request *IGuestProcessInputNotifyEventgetStatus) (*IGuestProcessInputNotifyEventgetStatusResponse, error) {
	response := new(IGuestProcessInputNotifyEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestProcessOutputEventgetData(request *IGuestProcessOutputEventgetData) (*IGuestProcessOutputEventgetDataResponse, error) {
	response := new(IGuestProcessOutputEventgetDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileEventgetFile(request *IGuestFileEventgetFile) (*IGuestFileEventgetFileResponse, error) {
	response := new(IGuestFileEventgetFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileRegisteredEventgetRegistered(request *IGuestFileRegisteredEventgetRegistered) (*IGuestFileRegisteredEventgetRegisteredResponse, error) {
	response := new(IGuestFileRegisteredEventgetRegisteredResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileStateChangedEventgetStatus(request *IGuestFileStateChangedEventgetStatus) (*IGuestFileStateChangedEventgetStatusResponse, error) {
	response := new(IGuestFileStateChangedEventgetStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileStateChangedEventgetError(request *IGuestFileStateChangedEventgetError) (*IGuestFileStateChangedEventgetErrorResponse, error) {
	response := new(IGuestFileStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileIOEventgetOffset(request *IGuestFileIOEventgetOffset) (*IGuestFileIOEventgetOffsetResponse, error) {
	response := new(IGuestFileIOEventgetOffsetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileIOEventgetProcessed(request *IGuestFileIOEventgetProcessed) (*IGuestFileIOEventgetProcessedResponse, error) {
	response := new(IGuestFileIOEventgetProcessedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestFileReadEventgetData(request *IGuestFileReadEventgetData) (*IGuestFileReadEventgetDataResponse, error) {
	response := new(IGuestFileReadEventgetDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetDevice(request *IUSBDeviceStateChangedEventgetDevice) (*IUSBDeviceStateChangedEventgetDeviceResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetAttached(request *IUSBDeviceStateChangedEventgetAttached) (*IUSBDeviceStateChangedEventgetAttachedResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetAttachedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IUSBDeviceStateChangedEventgetError(request *IUSBDeviceStateChangedEventgetError) (*IUSBDeviceStateChangedEventgetErrorResponse, error) {
	response := new(IUSBDeviceStateChangedEventgetErrorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) ISharedFolderChangedEventgetScope(request *ISharedFolderChangedEventgetScope) (*ISharedFolderChangedEventgetScopeResponse, error) {
	response := new(ISharedFolderChangedEventgetScopeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetFatal(request *IRuntimeErrorEventgetFatal) (*IRuntimeErrorEventgetFatalResponse, error) {
	response := new(IRuntimeErrorEventgetFatalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetId(request *IRuntimeErrorEventgetId) (*IRuntimeErrorEventgetIdResponse, error) {
	response := new(IRuntimeErrorEventgetIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IRuntimeErrorEventgetMessage(request *IRuntimeErrorEventgetMessage) (*IRuntimeErrorEventgetMessageResponse, error) {
	response := new(IRuntimeErrorEventgetMessageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceChangedEventgetListener(request *IEventSourceChangedEventgetListener) (*IEventSourceChangedEventgetListenerResponse, error) {
	response := new(IEventSourceChangedEventgetListenerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IEventSourceChangedEventgetAdd(request *IEventSourceChangedEventgetAdd) (*IEventSourceChangedEventgetAddResponse, error) {
	response := new(IEventSourceChangedEventgetAddResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetMachineId(request *IExtraDataChangedEventgetMachineId) (*IExtraDataChangedEventgetMachineIdResponse, error) {
	response := new(IExtraDataChangedEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetKey(request *IExtraDataChangedEventgetKey) (*IExtraDataChangedEventgetKeyResponse, error) {
	response := new(IExtraDataChangedEventgetKeyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataChangedEventgetValue(request *IExtraDataChangedEventgetValue) (*IExtraDataChangedEventgetValueResponse, error) {
	response := new(IExtraDataChangedEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventaddVeto(request *IVetoEventaddVeto) (*IVetoEventaddVetoResponse, error) {
	response := new(IVetoEventaddVetoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventisVetoed(request *IVetoEventisVetoed) (*IVetoEventisVetoedResponse, error) {
	response := new(IVetoEventisVetoedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVetoEventgetVetos(request *IVetoEventgetVetos) (*IVetoEventgetVetosResponse, error) {
	response := new(IVetoEventgetVetosResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetMachineId(request *IExtraDataCanChangeEventgetMachineId) (*IExtraDataCanChangeEventgetMachineIdResponse, error) {
	response := new(IExtraDataCanChangeEventgetMachineIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetKey(request *IExtraDataCanChangeEventgetKey) (*IExtraDataCanChangeEventgetKeyResponse, error) {
	response := new(IExtraDataCanChangeEventgetKeyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IExtraDataCanChangeEventgetValue(request *IExtraDataCanChangeEventgetValue) (*IExtraDataCanChangeEventgetValueResponse, error) {
	response := new(IExtraDataCanChangeEventgetValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IShowWindowEventgetWinId(request *IShowWindowEventgetWinId) (*IShowWindowEventgetWinIdResponse, error) {
	response := new(IShowWindowEventgetWinIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IShowWindowEventsetWinId(request *IShowWindowEventsetWinId) (*IShowWindowEventsetWinIdResponse, error) {
	response := new(IShowWindowEventsetWinIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetSlot(request *INATRedirectEventgetSlot) (*INATRedirectEventgetSlotResponse, error) {
	response := new(INATRedirectEventgetSlotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetRemove(request *INATRedirectEventgetRemove) (*INATRedirectEventgetRemoveResponse, error) {
	response := new(INATRedirectEventgetRemoveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetName(request *INATRedirectEventgetName) (*INATRedirectEventgetNameResponse, error) {
	response := new(INATRedirectEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetProto(request *INATRedirectEventgetProto) (*INATRedirectEventgetProtoResponse, error) {
	response := new(INATRedirectEventgetProtoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetHostIP(request *INATRedirectEventgetHostIP) (*INATRedirectEventgetHostIPResponse, error) {
	response := new(INATRedirectEventgetHostIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetHostPort(request *INATRedirectEventgetHostPort) (*INATRedirectEventgetHostPortResponse, error) {
	response := new(INATRedirectEventgetHostPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetGuestIP(request *INATRedirectEventgetGuestIP) (*INATRedirectEventgetGuestIPResponse, error) {
	response := new(INATRedirectEventgetGuestIPResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATRedirectEventgetGuestPort(request *INATRedirectEventgetGuestPort) (*INATRedirectEventgetGuestPortResponse, error) {
	response := new(INATRedirectEventgetGuestPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetPlugged(request *IHostPCIDevicePlugEventgetPlugged) (*IHostPCIDevicePlugEventgetPluggedResponse, error) {
	response := new(IHostPCIDevicePlugEventgetPluggedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetSuccess(request *IHostPCIDevicePlugEventgetSuccess) (*IHostPCIDevicePlugEventgetSuccessResponse, error) {
	response := new(IHostPCIDevicePlugEventgetSuccessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetAttachment(request *IHostPCIDevicePlugEventgetAttachment) (*IHostPCIDevicePlugEventgetAttachmentResponse, error) {
	response := new(IHostPCIDevicePlugEventgetAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IHostPCIDevicePlugEventgetMessage(request *IHostPCIDevicePlugEventgetMessage) (*IHostPCIDevicePlugEventgetMessageResponse, error) {
	response := new(IHostPCIDevicePlugEventgetMessageResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IVBoxSVCAvailabilityChangedEventgetAvailable(request *IVBoxSVCAvailabilityChangedEventgetAvailable) (*IVBoxSVCAvailabilityChangedEventgetAvailableResponse, error) {
	response := new(IVBoxSVCAvailabilityChangedEventgetAvailableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IBandwidthGroupChangedEventgetBandwidthGroup(request *IBandwidthGroupChangedEventgetBandwidthGroup) (*IBandwidthGroupChangedEventgetBandwidthGroupResponse, error) {
	response := new(IBandwidthGroupChangedEventgetBandwidthGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetChangeType(request *IGuestMonitorChangedEventgetChangeType) (*IGuestMonitorChangedEventgetChangeTypeResponse, error) {
	response := new(IGuestMonitorChangedEventgetChangeTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetScreenId(request *IGuestMonitorChangedEventgetScreenId) (*IGuestMonitorChangedEventgetScreenIdResponse, error) {
	response := new(IGuestMonitorChangedEventgetScreenIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetOriginX(request *IGuestMonitorChangedEventgetOriginX) (*IGuestMonitorChangedEventgetOriginXResponse, error) {
	response := new(IGuestMonitorChangedEventgetOriginXResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetOriginY(request *IGuestMonitorChangedEventgetOriginY) (*IGuestMonitorChangedEventgetOriginYResponse, error) {
	response := new(IGuestMonitorChangedEventgetOriginYResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetWidth(request *IGuestMonitorChangedEventgetWidth) (*IGuestMonitorChangedEventgetWidthResponse, error) {
	response := new(IGuestMonitorChangedEventgetWidthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestMonitorChangedEventgetHeight(request *IGuestMonitorChangedEventgetHeight) (*IGuestMonitorChangedEventgetHeightResponse, error) {
	response := new(IGuestMonitorChangedEventgetHeightResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetName(request *IGuestUserStateChangedEventgetName) (*IGuestUserStateChangedEventgetNameResponse, error) {
	response := new(IGuestUserStateChangedEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetDomain(request *IGuestUserStateChangedEventgetDomain) (*IGuestUserStateChangedEventgetDomainResponse, error) {
	response := new(IGuestUserStateChangedEventgetDomainResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetState(request *IGuestUserStateChangedEventgetState) (*IGuestUserStateChangedEventgetStateResponse, error) {
	response := new(IGuestUserStateChangedEventgetStateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IGuestUserStateChangedEventgetStateDetails(request *IGuestUserStateChangedEventgetStateDetails) (*IGuestUserStateChangedEventgetStateDetailsResponse, error) {
	response := new(IGuestUserStateChangedEventgetStateDetailsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetStorageDevice(request *IStorageDeviceChangedEventgetStorageDevice) (*IStorageDeviceChangedEventgetStorageDeviceResponse, error) {
	response := new(IStorageDeviceChangedEventgetStorageDeviceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetRemoved(request *IStorageDeviceChangedEventgetRemoved) (*IStorageDeviceChangedEventgetRemovedResponse, error) {
	response := new(IStorageDeviceChangedEventgetRemovedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) IStorageDeviceChangedEventgetSilent(request *IStorageDeviceChangedEventgetSilent) (*IStorageDeviceChangedEventgetSilentResponse, error) {
	response := new(IStorageDeviceChangedEventgetSilentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkChangedEventgetNetworkName(request *INATNetworkChangedEventgetNetworkName) (*INATNetworkChangedEventgetNetworkNameResponse, error) {
	response := new(INATNetworkChangedEventgetNetworkNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkStartStopEventgetStartEvent(request *INATNetworkStartStopEventgetStartEvent) (*INATNetworkStartStopEventgetStartEventResponse, error) {
	response := new(INATNetworkStartStopEventgetStartEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkCreationDeletionEventgetCreationEvent(request *INATNetworkCreationDeletionEventgetCreationEvent) (*INATNetworkCreationDeletionEventgetCreationEventResponse, error) {
	response := new(INATNetworkCreationDeletionEventgetCreationEventResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetEnabled(request *INATNetworkSettingEventgetEnabled) (*INATNetworkSettingEventgetEnabledResponse, error) {
	response := new(INATNetworkSettingEventgetEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetNetwork(request *INATNetworkSettingEventgetNetwork) (*INATNetworkSettingEventgetNetworkResponse, error) {
	response := new(INATNetworkSettingEventgetNetworkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetGateway(request *INATNetworkSettingEventgetGateway) (*INATNetworkSettingEventgetGatewayResponse, error) {
	response := new(INATNetworkSettingEventgetGatewayResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabled(request *INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabled) (*INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabledResponse, error) {
	response := new(INATNetworkSettingEventgetAdvertiseDefaultIPv6RouteEnabledResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkSettingEventgetNeedDhcpServer(request *INATNetworkSettingEventgetNeedDhcpServer) (*INATNetworkSettingEventgetNeedDhcpServerResponse, error) {
	response := new(INATNetworkSettingEventgetNeedDhcpServerResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetCreate(request *INATNetworkPortForwardEventgetCreate) (*INATNetworkPortForwardEventgetCreateResponse, error) {
	response := new(INATNetworkPortForwardEventgetCreateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetIpv6(request *INATNetworkPortForwardEventgetIpv6) (*INATNetworkPortForwardEventgetIpv6Response, error) {
	response := new(INATNetworkPortForwardEventgetIpv6Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetName(request *INATNetworkPortForwardEventgetName) (*INATNetworkPortForwardEventgetNameResponse, error) {
	response := new(INATNetworkPortForwardEventgetNameResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetProto(request *INATNetworkPortForwardEventgetProto) (*INATNetworkPortForwardEventgetProtoResponse, error) {
	response := new(INATNetworkPortForwardEventgetProtoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetHostIp(request *INATNetworkPortForwardEventgetHostIp) (*INATNetworkPortForwardEventgetHostIpResponse, error) {
	response := new(INATNetworkPortForwardEventgetHostIpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetHostPort(request *INATNetworkPortForwardEventgetHostPort) (*INATNetworkPortForwardEventgetHostPortResponse, error) {
	response := new(INATNetworkPortForwardEventgetHostPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetGuestIp(request *INATNetworkPortForwardEventgetGuestIp) (*INATNetworkPortForwardEventgetGuestIpResponse, error) {
	response := new(INATNetworkPortForwardEventgetGuestIpResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - InvalidObjectFault
//   - RuntimeFault

func (service *VboxPortType) INATNetworkPortForwardEventgetGuestPort(request *INATNetworkPortForwardEventgetGuestPort) (*INATNetworkPortForwardEventgetGuestPortResponse, error) {
	response := new(INATNetworkPortForwardEventgetGuestPortResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault `xml:",omitempty"`
	Content interface{}
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	// log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	// log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
