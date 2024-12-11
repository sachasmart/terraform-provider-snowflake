// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

// imports edited manually
import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type ProcedureShowOutputAssert struct {
	*assert.ResourceAssert
}

func ProcedureShowOutput(t *testing.T, name string) *ProcedureShowOutputAssert {
	t.Helper()

	p := ProcedureShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	p.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &p
}

func ImportedProcedureShowOutput(t *testing.T, id string) *ProcedureShowOutputAssert {
	t.Helper()

	p := ProcedureShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	p.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &p
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (p *ProcedureShowOutputAssert) HasCreatedOn(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasName(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasSchemaName(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasIsBuiltin(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_builtin", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasIsAggregate(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_aggregate", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasIsAnsi(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_ansi", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasMinNumArguments(expected int) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputIntValueSet("min_num_arguments", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasMaxNumArguments(expected int) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputIntValueSet("max_num_arguments", expected))
	return p
}

// HasArgumentsOld removed manually

func (p *ProcedureShowOutputAssert) HasArgumentsRaw(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("arguments_raw", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasDescription(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("description", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasCatalogName(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("catalog_name", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasIsTableFunction(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_table_function", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasValidForClustering(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("valid_for_clustering", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasIsSecure(expected bool) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputBoolValueSet("is_secure", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasSecrets(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("secrets", expected))
	return p
}

func (p *ProcedureShowOutputAssert) HasExternalAccessIntegrations(expected string) *ProcedureShowOutputAssert {
	p.AddAssertion(assert.ResourceShowOutputValueSet("external_access_integrations", expected))
	return p
}
