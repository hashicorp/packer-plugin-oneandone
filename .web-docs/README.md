The IONOS 1&1 Packer Plugin is able to create virtual machines for IONOS cloud.

### Installation
To install this plugin add this code into your Packer configuration and run [packer init](/packer/docs/commands/init)

```hcl
packer {
  required_plugins {
    oneandone = {
      version = "~> 1"
      source = "github.com/hashicorp/oneandone"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
packer plugins install github.com/hashicorp/oneandone
```

### Components

#### Builders
- [1&1 Builder](/packer/integrations/hashicorp/oneandone/latest/components/builder/oneandone)

