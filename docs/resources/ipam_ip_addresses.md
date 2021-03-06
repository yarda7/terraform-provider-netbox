# netbox\_ipam\_ip\_addresses Resource

Manage an ip address within Netbox.

## Example Usage

```hcl
resource "netbox_ipam_ip_addresses" "ip_test" {
  address = "192.168.56.0/24"
  description = "IP created by terraform"
  status = "active"
  
  tag {
    name = "tag1"
    slug = "tag1"
  }
}
```

## Argument Reference

The following arguments are supported:
* ``address`` - (Required) The IP address (with mask) used for this object.
* ``description`` - (Optional) The description of this object.
* ``dns_name`` - (Optional) The DNS name of this object.
* ``nat_inside_id`` - (Optional) The ID of the NAT inside of this object.
* ``nat_outside_id`` - (Optional) The ID of the NAT outside of this object.
* ``object_id`` - (Optional) The ID of the object where this resource is attached to.
* ``object_type`` - (Optional) The object type among virtualization.vminterface
or dcim.interface (virtualization.vminterface by default)
* ``primary_ip4`` - (Optional) Set this resource as primary IPv4 (false by default)
* ``role`` - (Optional) The role among loopback, secondary, anycast, vip, vrrp, hsrp, glbp, carp of this object.
* ``status`` - (Optional) The status among container, active, reserved, deprecated (active by default).
* ``tenant_id`` - (Optional) ID of the tenant where this object is attached.
* ``vrf_id`` - (Optional) The ID of the vrf attached to this object.

The ``tag`` block supports:
* ``name`` - (Required) Name of the existing tag to associate with this resource.
* ``slug`` - (Required) Slug of the existing tag to associate with this resource.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:
* ``id`` - The id (ref in Netbox) of this object.
