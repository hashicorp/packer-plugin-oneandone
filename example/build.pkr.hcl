packer {
  required_plugins {
    oneandone = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/oneandone"
    }
  }
}

source "oneandone" "example" {
  disk_size     = "50"
  image         = "ubuntu1604-64min"
  snapshot_name = "test5"
}

build {
  sources = ["source.oneandone.example"]
}
