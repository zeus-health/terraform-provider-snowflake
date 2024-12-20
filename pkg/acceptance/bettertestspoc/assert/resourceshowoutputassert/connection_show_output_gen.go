// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type ConnectionShowOutputAssert struct {
	*assert.ResourceAssert
}

func ConnectionShowOutput(t *testing.T, name string) *ConnectionShowOutputAssert {
	t.Helper()

	c := ConnectionShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	c.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &c
}

func ImportedConnectionShowOutput(t *testing.T, id string) *ConnectionShowOutputAssert {
	t.Helper()

	c := ConnectionShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	c.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &c
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (c *ConnectionShowOutputAssert) HasRegionGroup(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("region_group", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasSnowflakeRegion(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("snowflake_region", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasCreatedOn(expected time.Time) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return c
}

func (c *ConnectionShowOutputAssert) HasAccountName(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("account_name", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasName(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasComment(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasIsPrimary(expected bool) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_primary", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasPrimary(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("primary", expected))
	return c
}

/*
func (c *ConnectionShowOutputAssert) HasFailoverAllowedToAccounts(expected []string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("failover_allowed_to_accounts", expected))
	return c
}
*/

func (c *ConnectionShowOutputAssert) HasConnectionUrl(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("connection_url", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasOrganizationName(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("organization_name", expected))
	return c
}

func (c *ConnectionShowOutputAssert) HasAccountLocator(expected string) *ConnectionShowOutputAssert {
	c.AddAssertion(assert.ResourceShowOutputValueSet("account_locator", expected))
	return c
}
