#!/bin/bash
set -euo pipefail

echo "🔧 Replacing 'title:' with 'page_title:' in all Markdown files..."
find .work/bpg/proxmox/docs/resources/ -type f -name '*.md' -exec sed -i 's/^title:/page_title:/' {} +

insert_description() {
  local file=$1
  local description=$2

  if grep -q "^description:" "$file"; then
    echo "Skipping $file (already has a description so no needd to insert)"
  else
    echo "Inserting description into $file"
    sed -i "/^subcategory: Virtual Environment/a description: |\
  $description" "$file"
  fi
}

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_certificate.md \
"Manages the custom SSL/TLS certificate for a specific node."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_cluster_firewall.md \
"Manages firewall options on the cluster level."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_cluster_firewall_security_group.md \
"A security group is a collection of rules, defined at cluster level, which can be used in all VMs' rules. For example, you can define a group named webserver with rules to open the http and https ports."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_container.md \
"Manages a container."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_dns.md \
"Manages the DNS configuration for a specific node."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_file.md \
"Use this resource to upload files to a Proxmox VE node. The file can be a backup, an ISO image, a snippet, or a container template depending on the content_type attribute."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_firewall_alias.md \
"Aliases are used to see what devices or group of devices are affected by a rule. We can create aliases to identify an IP address or a network. Aliases can be created on the cluster level, on VM / Container level."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_firewall_ipset.md \
"An IPSet allows us to group multiple IP addresses, IP subnets and aliases. Aliases can be created on the cluster level, on VM / Container level."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_firewall_options.md \
"Manages firewall options on VM / Container level."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_firewall_rules.md \
"A security group is a collection of rules, defined at cluster level, which can be used in all VMs' rules. For example, you can define a group named “webserver” with rules to open the http and https ports. Rules can be created on the cluster level, on VM / Container level."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_group.md \
"Manages a user group."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_hosts.md \
"Manages the host entries on a specific node."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_pool.md \
"Manages a resource pool."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_role.md \
"Manages a role."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_time.md \
"Manages the time for a specific node."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_user.md \
"Manages a user."

insert_description .work/bpg/proxmox/docs/resources/virtual_environment_vm.md \
"Manages a virtual machine."

echo "✅✅✅✅✅ Script completed successfully. Descriptions added only when missing."
