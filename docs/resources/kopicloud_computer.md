---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kopicloud_computer Resource - kopicloud-ad-tf-provider"
subcategory: ""
description: |-
  
---

# kopicloud_computer (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ad_computer_name` (String)

### Optional

- `description` (String)
- `ou_path` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `result` (List of Object) Single Element List of Computer (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `computer_name` (String)
- `created` (String)
- `description` (String)
- `dns_name` (String)
- `operating_system` (String)
- `path` (String)
- `sid` (String)

