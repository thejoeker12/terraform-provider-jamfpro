---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "jamfpro_mobile_device_configuration_profile Resource - terraform-provider-jamfpro"
subcategory: ""
description: |-
  
---

# jamfpro_mobile_device_configuration_profile (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the mobile device configuration profile.

### Optional

- `category` (Block List, Max: 1) The jamf pro category information for the mobile device configuration profile. (see [below for nested schema](#nestedblock--category))
- `deployment_method` (String) The deployment method for the mobile device configuration profile, can be either 'Install Automatically' or 'Make Available in Self Service'.
- `description` (String) The description of the mobile device configuration profile.
- `level` (String) The level at which the mobile device configuration profile is applied, can be either 'Device Level' or 'User Level'.
- `payloads` (String) The iOS / iPadOS / tvOS configuration profile payload. Can be a file path to a .mobileconfig or a string with an embedded mobileconfig plist.
- `redeploy_days_before_cert_expires` (Number) The number of days before certificate expiration when the profile should be redeployed.
- `redeploy_on_update` (String) Defines the redeployment behaviour when a mobile device config profile update occurs.This is always 'Newly Assigned' on new profile objects, but may be set 'All' on profile update requests and in TF state
- `scope` (Block List, Max: 1) The scope in which the mobile device configuration profile is applied. (see [below for nested schema](#nestedblock--scope))
- `site` (Block List, Max: 1) The site information associated with the mobile device configuration profile. (see [below for nested schema](#nestedblock--site))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The unique identifier for the mobile device configuration profile.
- `uuid` (String) The universally unique identifier for the profile.

<a id="nestedblock--category"></a>
### Nested Schema for `category`

Optional:

- `id` (Number) The unique identifier for the Jamf Pro category.


<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `all_jss_users` (Boolean) If true, the profile is applied to all JSS users.
- `all_mobile_devices` (Boolean) If true, the profile is applied to all mobile devices.
- `building_ids` (List of Number) A list of building IDs associated with the profile.
- `department_ids` (List of Number) A list of department IDs associated with the profile.
- `exclusions` (Block List, Max: 1) Exclusions for the profile. (see [below for nested schema](#nestedblock--scope--exclusions))
- `jss_user_group_ids` (List of Number) A list of JSS user group IDs associated with the profile.
- `jss_user_ids` (List of Number) A list of JSS user IDs associated with the profile.
- `limitations` (Block List, Max: 1) Limitations for the profile. (see [below for nested schema](#nestedblock--scope--limitations))
- `mobile_device_group_ids` (List of Number) A list of mobile device group IDs associated with the profile.
- `mobile_device_ids` (List of Number) A list of mobile device IDs associated with the profile.

<a id="nestedblock--scope--exclusions"></a>
### Nested Schema for `scope.exclusions`

Optional:

- `building_ids` (List of Number) A list of building IDs for exclusions.
- `department_ids` (List of Number) A list of department IDs for exclusions.
- `ibeacon_ids` (List of Number) A list of iBeacon IDs for exclusions.
- `jss_user_group_ids` (List of Number) A list of JSS user group IDs for exclusions.
- `jss_user_ids` (List of Number) A list of user names for exclusions.
- `mobile_device_group_ids` (List of Number) A list of mobile device group IDs for exclusions.
- `mobile_device_ids` (List of Number) A list of mobile device IDs for exclusions.
- `network_segment_ids` (List of Number) A list of network segment IDs for exclusions.


<a id="nestedblock--scope--limitations"></a>
### Nested Schema for `scope.limitations`

Optional:

- `directory_service_or_local_usernames` (List of String) A list of directory service / local usernames for scoping limitations.
- `ibeacon_ids` (List of Number) A list of iBeacon IDs for limitations.
- `network_segment_ids` (List of Number) A list of network segment IDs for limitations.
- `user_group_ids` (List of Number) A list of user group IDs for limitations.



<a id="nestedblock--site"></a>
### Nested Schema for `site`

Optional:

- `id` (Number)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
