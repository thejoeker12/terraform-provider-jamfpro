---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jamfpro_account Resource - terraform-provider-jamfpro"
subcategory: ""
description: |-
  
---

# jamfpro_account (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_level` (String) The access level of the account. This can be either Full Access, scoped to a jamf pro site with Site Access, or scoped to a jamf pro account group with Group Access
- `enabled` (String) Access status of the account (“enabled” or “disabled”).
- `name` (String) The name of the jamf pro account.

### Optional

- `casper_admin_privileges` (Set of String) Privileges related to Casper Admin.
- `casper_imaging_privileges` (Set of String) Privileges related to Casper Imaging.
- `casper_remote_privileges` (Set of String) Privileges related to Casper Remote.
- `directory_user` (Boolean) Indicates if the user is a directory user.
- `email` (String) The email of the account user.
- `email_address` (String) The email address of the account user.
- `force_password_change` (Boolean) Indicates if the user is forced to change password on next login.
- `full_name` (String) The full name of the account user.
- `groups` (Block Set) A set of group names and IDs associated with the account. (see [below for nested schema](#nestedblock--groups))
- `identity_server` (Block List, Max: 1) LDAP or IdP server associated with the account group. (see [below for nested schema](#nestedblock--identity_server))
- `jss_actions_privileges` (Set of String) Privileges related to JSS Actions.
- `jss_objects_privileges` (Set of String) Privileges related to JSS Objects.
- `jss_settings_privileges` (Set of String) Privileges related to JSS Settings.
- `password` (String, Sensitive) The password for the account.
- `privilege_set` (String) The privilege set assigned to the account.
- `recon_privileges` (Set of String) Privileges related to Recon.
- `site` (Block List, Max: 1) The site information associated with the account group if access_level is set to Site Access. (see [below for nested schema](#nestedblock--site))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The unique identifier of the jamf pro account.

<a id="nestedblock--groups"></a>
### Nested Schema for `groups`

Optional:

- `id` (Number)
- `name` (String)


<a id="nestedblock--identity_server"></a>
### Nested Schema for `identity_server`

Optional:

- `id` (Number) ID is the ID of the LDAP or IdP configuration in Jamf Pro.


<a id="nestedblock--site"></a>
### Nested Schema for `site`

Optional:

- `id` (Number) Jamf Pro Site ID. Value defaults to '0' aka not used.
- `name` (String) Jamf Pro Site Name


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
