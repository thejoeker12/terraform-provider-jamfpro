// computergroup_data_validation.go
package computergroups

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// customDiffComputeGroups is the top-level custom diff function.
func customDiffComputeGroups(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if err := validateComputersNotAllowedWithSmart(ctx, d, meta); err != nil {
		return err
	}

	if err := validateCriteriaFields(ctx, d, meta); err != nil {
		return err
	}

	return nil
}

// validateComputersNotAllowedWithSmart checks that 'computers' is not set when 'is_smart' is true and a site is set
func validateComputersNotAllowedWithSmart(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	isSmart := d.Get("is_smart").(bool)

	if isSmart {
		if computers, ok := d.GetOk("computers"); ok {
			if len(computers.([]interface{})) > 0 {
				return fmt.Errorf("'computers' field not allowed when is_smart is true")
			}
		}
	}

	return nil
}

// validateCriteriaFields validates that 'name', 'and_or', and 'search_type' are set in each 'criteria' if 'criteria' is populated.
func validateCriteriaFields(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	isSmart := d.Get("is_smart").(bool)
	if !isSmart {
		if _, ok := d.GetOk("criteria"); ok {
			return fmt.Errorf("'criteria' not allowed when is_smart is false")
		}
	}

	return nil
}
