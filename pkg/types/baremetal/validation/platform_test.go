package validation

import (
	"testing"

	"github.com/openshift/installer/pkg/ipnet"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/baremetal"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidatePlatform(t *testing.T) {
	interfaceValidator := func(p *baremetal.Platform, fldPath *field.Path) field.ErrorList {
		errorList := field.ErrorList{}

		if p.ExternalBridge != "br0" {
			errorList = append(errorList, field.Invalid(fldPath.Child("externalBridge"), p.ExternalBridge,
				"invalid external bridge"))
		}

		if p.ProvisioningBridge != "br1" {
			errorList = append(errorList, field.Invalid(fldPath.Child("provisioningBridge"), p.ProvisioningBridge,
				"invalid provisioning bridge"))
		}

		return errorList
	}

	dynamicValidators = append(dynamicValidators, interfaceValidator)
	network := &types.Networking{MachineNetwork: []types.MachineNetworkEntry{{CIDR: *ipnet.MustParseCIDR("192.168.111.0/24")}}}
	cases := []struct {
		name     string
		platform *baremetal.Platform
		network  *types.Networking
		expected string
	}{
		{
			name: "valid",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				ProvisioningNetworkInterface: "ens3",
			},
			network: network,
		},
		{
			name: "valid_ipv6_provisioning",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("fd2e:6f44:5dd8:b856::/64"),
				ClusterProvisioningIP:        "fd2e:6f44:5dd8:b856::3",
				BootstrapProvisioningIP:      "fd2e:6f44:5dd8:b856::2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				ProvisioningNetworkInterface: "ens3",
			},
			network: network,
		},

		{
			name: "valid_with_os_image_overrides",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "http://192.168.111.1/images/qemu.x86_64.qcow2.gz?sha256=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http://192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=340dfa4d92450f2eee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
				ProvisioningNetworkInterface: "ens3",
			},
			network: network,
		},
		{
			name: "valid_provisioningDHCPRange",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ProvisioningDHCPRange:        "172.22.0.10,172.22.0.50",
				ProvisioningNetworkInterface: "ens3",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network: network,
		},
		{
			name: "invalid_provisioningDHCPRange_missing_pair",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ProvisioningDHCPRange:        "172.22.0.10,",
				ProvisioningNetworkInterface: "ens3",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "provisioningDHCPRange: Invalid value: \"172.22.0.10,\": : \"\" is not a valid IP",
		},
		{
			name: "invalid_provisioningDHCPRange_not_a_range",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ProvisioningDHCPRange:        "172.22.0.19",
				ProvisioningNetworkInterface: "ens3",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"172.22.0.19\": provisioning dhcp range should be in format: start_ip,end_ip",
		},
		{
			name: "invalid_provisioningDHCPRange_wrong_CIDR",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningDHCPRange:        "192.168.128.1,172.22.0.100",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.128.1,172.22.0.100\": \"192.168.128.1\" is not in the provisioning network",
		},
		{
			name: "invalid_apivip",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.222.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ProvisioningNetworkInterface: "ens3",
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.222.2\": the virtual IP is expected to be in one of the machine networks",
		},
		{
			name: "invalid_dnsvip",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.222.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ProvisioningNetworkInterface: "ens3",
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.222.3\": the virtual IP is expected to be in one of the machine networks",
		},
		{
			name: "invalid_ingressvip",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.222.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.222.4\": the virtual IP is expected to be in one of the machine networks",
		},
		{
			name: "invalid_hosts",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        nil,
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "bare metal hosts are missing",
		},
		{
			name: "invalid_libvirturi",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "invalid URI \"\"",
		},
		{
			name: "invalid_extbridge",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "noexist",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"noexist\": invalid external bridge",
		},
		{
			name: "invalid_provbridge",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "noexist",
			},
			network:  network,
			expected: "Invalid value: \"noexist\": invalid provisioning bridge",
		},

		{
			name: "invalid_clusterprovip_machineCIDR",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "192.168.111.5",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.111.5\": the IP must not be in one of the machine networks",
		},
		{
			name: "invalid_clusterprovip_wrongCIDR",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "192.168.128.1",
				BootstrapProvisioningIP:      "172.22.0.2",
				ProvisioningDHCPRange:        "172.22.0.10,172.22.0.100",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.128.1\": \"192.168.128.1\" is not in the provisioning network",
		},
		{
			name: "invalid_bootstrapprovip_machineCIDR",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "192.168.111.5",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.111.5\": the IP must not be in one of the machine networks",
		},
		{
			name: "invalid_bootstraposimage",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "192.168.111.1/images/qemu.x86_64.qcow2.gz?sha256=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http://192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=340dfa4d92450f2eee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
			},
			network:  network,
			expected: "the URI provided.*is invalid",
		},
		{
			name: "invalid_clusterosimage",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ClusterProvisioningIP:        "172.22.0.3",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "http://192.168.111.1/images/qemu.x86_64.qcow2.gz?sha256=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http//192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=340dfa4d92450f2eee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
			},
			network:  network,
			expected: "the URI provided.*is invalid",
		},
		{
			name: "invalid_bootstraposimage_checksum",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "http://192.168.111.1/images/qemu.x86_64.qcow2.gz?md5sum=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http://192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=340dfa4d92450f2eee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
			},
			network:  network,
			expected: "the sha256 parameter in the.*URI is missing",
		},
		{
			name: "invalid_clusterosimage_checksum",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "http://192.168.111.1/images/qemu.x86_64.qcow2.gz?sha256=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http://192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=3ee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
			},
			network:  network,
			expected: "the sha256 parameter in the.*URI is invalid",
		},
		{
			name: "invalid_bootstraposimage_uri_scheme",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ProvisioningNetworkInterface: "ens3",
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "172.22.0.2",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
				BootstrapOSImage:             "xttp://192.168.111.1/images/qemu.x86_64.qcow2.gz?sha256=3b5a882c2af3e19d515b961855d144f293cab30190c2bdedd661af31a1fc4e2f",
				ClusterOSImage:               "http://192.168.111.1/images/metal.x86_64.qcow2.gz?sha256=340dfa4d92450f2eee852ed1e2d02e3138cc68d824827ef9cf0a40a7ea2f93da",
			},
			network:  network,
			expected: "the URI provided.*must begin with http/https",
		},
		{
			name: "invalid_bootstrapprovip_wrongCIDR",
			platform: &baremetal.Platform{
				APIVIP:                       "192.168.111.2",
				DNSVIP:                       "192.168.111.3",
				IngressVIP:                   "192.168.111.4",
				Hosts:                        []*baremetal.Host{},
				LibvirtURI:                   "qemu://system",
				ProvisioningNetworkInterface: "ens3",
				ProvisioningNetworkCIDR:      ipnet.MustParseCIDR("172.22.0.0/24"),
				ClusterProvisioningIP:        "172.22.0.3",
				BootstrapProvisioningIP:      "192.168.128.1",
				ProvisioningDHCPRange:        "172.22.0.10,172.22.0.100",
				ExternalBridge:               "br0",
				ProvisioningBridge:           "br1",
			},
			network:  network,
			expected: "Invalid value: \"192.168.128.1\": \"192.168.128.1\" is not in the provisioning network",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePlatform(tc.platform, tc.network, field.NewPath("test-path")).ToAggregate()
			if tc.expected == "" {
				assert.NoError(t, err)
			} else {
				assert.Regexp(t, tc.expected, err)
			}
		})
	}
}
