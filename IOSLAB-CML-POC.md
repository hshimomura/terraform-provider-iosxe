# ioslab CML Proof Notes

This branch was validated against the shared ioslab CML `cml-cr1` proof router at `172.24.0.34` using NETCONF and a locally built provider binary.

## Scope Tested

- ZBFW: IPv4 ACL, IPv6 ACL, `class-map type inspect`, `policy-map type inspect`, `zone security`, `zone-pair security`, and `zone_member_security` on `GigabitEthernet1`.
- QoS: IPv4 ACL, QoS class-map, policy-map with `priority level 1`, `fair-queue`, `set dscp default`, and `service_policy_output` on `GigabitEthernet1`.
- CoPP: IPv4 ACL, IPv6 ACL, class-maps, policy-map policing with both `exceed-action drop` and `exceed-action transmit`, and `control_plane_service_policy_input` via `iosxe_system`.

## Provider Changes Proven

- Added `iosxe_access_list_ipv6` so IPv6 ACLs can be managed directly instead of relying on pre-existing CLI/YAML-owned objects.
- Added `police_cir_exceed_transmit` so CoPP `class-default` can round-trip without using the unrelated target-bitrate police action.

## Verification Performed

- Built the provider locally with Go 1.25.8.
- Ran Terraform `init`, `plan`, `apply`, then a second `plan -detailed-exitcode` with the local provider override.
- The second plan returned no changes.
- Fetched running-config over NETCONF and confirmed the temporary `TF_CML_*` ZBFW/QoS/CoPP objects were present with expected IOS XE CLI.
- Ran Terraform destroy for all temporary `TF_CML_*` objects.
- Restored iosxecore YAML baseline for ZBFW, QoS, QoS interface attachment, ZBFW interface zones, and CoPP using the existing selective-replace tooling.
- Confirmed no `TF_CML_*` objects remained on `cml-cr1`.
- Ran `go test ./...` successfully.

The shared CML lab was returned to the iosxecore baseline after the proof.
