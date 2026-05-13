// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeAccessListIPv6IcmpNamed(t *testing.T) {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "name", "TFACC_IPV6_ICMP_NAMED"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.sequence", "10"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.remark", "ipv6 icmp named proof"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.sequence", "20"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.ace_rule_action", "permit"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.ace_rule_protocol", "icmp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.source_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.destination_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.icmp_named_msg_type", "packet-too-big"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.sequence", "30"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.ace_rule_action", "permit"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.ace_rule_protocol", "tcp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.source_prefix", "2001:db8:64::/64"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_port_equal", "22"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAccessListIPv6IcmpNamedConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccIosxeAccessListIPv6IcmpNamedConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_access_list_ipv6.named_icmp",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "TFACC_IPV6_ICMP_NAMED",
				ImportStateVerifyIgnore: []string{
					"entries.0.source_any",
					"entries.0.destination_any",
					"entries.0.fragments",
					"entries.0.log",
					"entries.0.log_input",
					"entries.1.fragments",
					"entries.1.log",
					"entries.1.log_input",
					"entries.2.source_any",
					"entries.2.fragments",
					"entries.2.log",
					"entries.2.log_input",
				},
				Check: resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeAccessListIPv6IcmpNamedConfig() string {
	return `resource "iosxe_access_list_ipv6" "named_icmp" {
  name = "TFACC_IPV6_ICMP_NAMED"
  entries = [
    {
      sequence = 10
      remark = "ipv6 icmp named proof"
    },
    {
      sequence = 20
      ace_rule_action = "permit"
      ace_rule_protocol = "icmp"
      source_any = true
      destination_any = true
      icmp_named_msg_type = "packet-too-big"
    },
    {
      sequence = 30
      ace_rule_action = "permit"
      ace_rule_protocol = "tcp"
      source_prefix = "2001:db8:64::/64"
      destination_any = true
      destination_port_equal = "22"
    },
  ]
}
`
}
