---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kopicloud_user_list Data Source - kopicloud-ad-tf-provider"
subcategory: ""
description: |-
  Read Users from Active Directory
---

# kopicloud_user_list (Data Source)

Read Users from Active Directory



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ou_path` (String)
- `recursive` (Boolean)
- `show_fields` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `result` (List of Object) List of User (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `change_password_next_logon` (Boolean)
- `city` (String)
- `company` (String)
- `country` (String)
- `department` (String)
- `description` (String)
- `display_name` (String)
- `email_address` (String)
- `first_name` (String)
- `home_folder_directory` (String)
- `home_folder_drive` (String)
- `home_folder_path` (String)
- `home_phone` (String)
- `initials` (String)
- `job_title` (String)
- `last_name` (String)
- `manager` (String)
- `middle_name` (String)
- `mobile_phone` (String)
- `office` (String)
- `office_phone` (String)
- `ou_path` (String)
- `postal_code` (String)
- `profile_logon_script` (String)
- `profile_path` (String)
- `rds_allow_logon` (Boolean)
- `rds_connect_drive` (Boolean)
- `rds_home_folder_drive` (String)
- `rds_home_folder_path` (String)
- `rds_profile_path` (String)
- `sam_username` (String)
- `state` (String)
- `street_address` (String)
- `street_po_box` (String)
- `username` (String)


