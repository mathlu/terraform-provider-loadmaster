---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "loadmaster_real_server Resource - terraform-provider-loadmaster"
subcategory: ""
description: |-
  Provides a resource for Real Servers. Use this to create and manage Real Servers assigned to Virtual Services in the KEMP LoadMaster.
---

# loadmaster_real_server (Resource)

Provides a resource for Real Servers. Use this to create and manage Real Servers assigned to Virtual Services in the KEMP LoadMaster.

## Example Usage

```terraform
resource "loadmaster_virtual_service" "foo" {
  address  = "192.168.1.10"
  protocol = "tcp"
  port     = "8080"
  nickname = "bar"
  type     = "gen"
}


resource "loadmaster_real_server" "baz" {
  virtual_service_id = loadmaster_virtual_serice.foo.id
  address            = "192.168.1.11"
  port               = "8080"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) The Real Server IP address.
- `port` (Number) The port on the Real Server to be used.
- `virtual_service_id` (Number) Virtual Service index number.

### Read-Only

- `id` (String) The ID of this resource.