// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type ViewShowOutputAssert struct {
	*assert.ResourceAssert
}

func ViewShowOutput(t *testing.T, name string) *ViewShowOutputAssert {
	t.Helper()

	v := ViewShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	v.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &v
}

func ImportedViewShowOutput(t *testing.T, id string) *ViewShowOutputAssert {
	t.Helper()

	v := ViewShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	v.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &v
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (v *ViewShowOutputAssert) HasCreatedOn(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected))
	return v
}

func (v *ViewShowOutputAssert) HasName(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return v
}

func (v *ViewShowOutputAssert) HasKind(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("kind", expected))
	return v
}

func (v *ViewShowOutputAssert) HasReserved(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("reserved", expected))
	return v
}

func (v *ViewShowOutputAssert) HasDatabaseName(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return v
}

func (v *ViewShowOutputAssert) HasSchemaName(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return v
}

func (v *ViewShowOutputAssert) HasOwner(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return v
}

func (v *ViewShowOutputAssert) HasComment(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return v
}

func (v *ViewShowOutputAssert) HasText(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("text", expected))
	return v
}

func (v *ViewShowOutputAssert) HasIsSecure(expected bool) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_secure", expected))
	return v
}

func (v *ViewShowOutputAssert) HasIsMaterialized(expected bool) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_materialized", expected))
	return v
}

func (v *ViewShowOutputAssert) HasOwnerRoleType(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return v
}

func (v *ViewShowOutputAssert) HasChangeTracking(expected string) *ViewShowOutputAssert {
	v.AddAssertion(assert.ResourceShowOutputValueSet("change_tracking", expected))
	return v
}