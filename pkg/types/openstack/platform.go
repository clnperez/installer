package openstack

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// Region specifies the OpenStack region where the cluster will be created.
	// Deprecated: this value is not used by the installer.
	// +optional
	DeprecatedRegion string `json:"region,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on OpenStack for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// Cloud is the name of OpenStack cloud to use from clouds.yaml.
	Cloud string `json:"cloud"`

	// ExternalNetwork is name of the external network in your OpenStack cluster.
	// +optional
	ExternalNetwork string `json:"externalNetwork,omitempty"`

	// DeprecatedFlavorName is the name of the flavor to use for instances in this cluster.
	// Deprecated: use FlavorName in DefaultMachinePlatform to define default flavor.
	// +optional
	DeprecatedFlavorName string `json:"computeFlavor,omitempty"`

	// LbFloatingIP is the IP address of an available floating IP in your OpenStack cluster
	// to associate with the OpenShift load balancer.
	// Deprecated: this value has been renamed to apiFloatingIP.
	// +optional
	DeprecatedLbFloatingIP string `json:"lbFloatingIP,omitempty"`

	// APIFloatingIP is the IP address of an available floating IP in your OpenStack cluster
	// to associate with the OpenShift API load balancer.
	// +optional
	APIFloatingIP string `json:"apiFloatingIP,omitempty"`

	// IngressFloatingIP is the ID of an available floating IP in your OpenStack cluster
	// that will be associated with the OpenShift ingress port
	// +optional
	IngressFloatingIP string `json:"ingressFloatingIP,omitempty"`

	// ExternalDNS holds the IP addresses of dns servers that will
	// be added to the dns resolution of all instances in the cluster.
	// +optional
	ExternalDNS []string `json:"externalDNS"`

	// TrunkSupport holds a `0` or `1` value that indicates whether or not to use trunk ports
	// in your OpenShift cluster.
	// Deprecated: this value is set by the installer automatically.
	// +optional
	DeprecatedTrunkSupport string `json:"trunkSupport,omitempty"`

	// OctaviaSupport holds a `0` or `1` value that indicates whether your OpenStack
	// cluster supports Octavia Loadbalancing.
	// Deprecated: this value is set by the installer automatically.
	// +optional
	DeprecatedOctaviaSupport string `json:"octaviaSupport,omitempty"`

	// ClusterOSImage is either a URL with `http(s)` or `file` scheme to override
	// the default OS image for cluster nodes, or an existing Glance image name.
	// +optional
	ClusterOSImage string `json:"clusterOSImage,omitempty"`

	// ClusterOSImageProperties is a list of properties to be added to the metadata of the uploaded Glance ClusterOSImage.
	// Default: the default is to not set any properties.
	// +optional
	ClusterOSImageProperties map[string]string `json:"clusterOSImageProperties,omitempty"`

	// APIVIP is the static IP on the nodes subnet that the api port for openshift will be assigned
	// Default: will be set to the 5 on the first entry in the machineNetwork CIDR
	// +optional
	APIVIP string `json:"apiVIP,omitempty"`

	// IngressVIP is the static IP on the nodes subnet that the apps port for openshift will be assigned
	// Default: will be set to the 7 on the first entry in the machineNewtwork CIDR
	// +optional
	IngressVIP string `json:"ingressVIP,omitempty"`

	// MachinesSubnet is the UUIDv4 of an openstack subnet. This subnet will be used by all nodes created by the installer.
	// By setting this, the installer will no longer create a network and subnet.
	// The subnet and network specified in MachinesSubnet will not be deleted or modified by the installer.
	// +optional
	MachinesSubnet string `json:"machinesSubnet,omitempty"`

	// DisableSecurityGroups will allow clusters to be created without the use of OpenStack Security Groups. Some Openstack
	// configurations do not support security groups. This is generally not recommended where security groups are supported.
	DisableSecurityGroups bool `json:"disableSecurityGroups,omitempty"`
}
