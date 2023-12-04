# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "IONOS 1&1"
  description = "The 1&1 Packer Plugin is able to create virtual machines for IONOS cloud."
  identifier = "packer/hashicorp/oneandone"
  flags = ["archived"]
  component {
    type = "builder"
    name = "1&1"
    slug = "oneandone"
  }
}
