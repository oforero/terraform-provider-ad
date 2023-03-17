---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kopicloud_ou Resource - kopicloud-ad-tf-provider"
subcategory: ""
description: |-
  
---

# kopicloud_ou (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String)
- `ou_name` (String)

### Optional

- `ou_path` (String)
- `protected` (Boolean)

### Read-Only

- `id` (String) The ID of this resource.
- `result` (List of Object) Single Element List of OU (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `description` (String)
- `guid` (String)
- `name` (String)
- `path` (String)
- `protected` (Boolean)

