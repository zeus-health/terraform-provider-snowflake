// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type ViewModel struct {
	AggregationPolicy   tfconfig.Variable `json:"aggregation_policy,omitempty"`
	ChangeTracking      tfconfig.Variable `json:"change_tracking,omitempty"`
	Columns             tfconfig.Variable `json:"columns,omitempty"`
	Comment             tfconfig.Variable `json:"comment,omitempty"`
	CopyGrants          tfconfig.Variable `json:"copy_grants,omitempty"`
	DataMetricFunctions tfconfig.Variable `json:"data_metric_functions,omitempty"`
	DataMetricSchedule  tfconfig.Variable `json:"data_metric_schedule,omitempty"`
	Database            tfconfig.Variable `json:"database,omitempty"`
	IsRecursive         tfconfig.Variable `json:"is_recursive,omitempty"`
	IsSecure            tfconfig.Variable `json:"is_secure,omitempty"`
	IsTemporary         tfconfig.Variable `json:"is_temporary,omitempty"`
	Name                tfconfig.Variable `json:"name,omitempty"`
	OrReplace           tfconfig.Variable `json:"or_replace,omitempty"`
	RowAccessPolicy     tfconfig.Variable `json:"row_access_policy,omitempty"`
	Schema              tfconfig.Variable `json:"schema,omitempty"`
	Statement           tfconfig.Variable `json:"statement,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func View(
	resourceName string,
	database string, name string, schema string, statement string,
) *ViewModel {
	v := &ViewModel{ResourceModelMeta: config.Meta(resourceName, resources.View)}
	v.WithDatabase(database)
	v.WithName(name)
	v.WithSchema(schema)
	v.WithStatement(statement)
	return v
}

func ViewWithDefaultMeta(
	database string, name string, schema string, statement string,
) *ViewModel {
	v := &ViewModel{ResourceModelMeta: config.DefaultMeta(resources.View)}
	v.WithDatabase(database)
	v.WithName(name)
	v.WithSchema(schema)
	v.WithStatement(statement)
	return v
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

// aggregation_policy attribute type is not yet supported, so WithAggregationPolicy can't be generated

func (v *ViewModel) WithChangeTracking(changeTracking string) *ViewModel {
	v.ChangeTracking = tfconfig.StringVariable(changeTracking)
	return v
}

// columns attribute type is not yet supported, so WithColumns can't be generated

func (v *ViewModel) WithComment(comment string) *ViewModel {
	v.Comment = tfconfig.StringVariable(comment)
	return v
}

func (v *ViewModel) WithCopyGrants(copyGrants bool) *ViewModel {
	v.CopyGrants = tfconfig.BoolVariable(copyGrants)
	return v
}

// data_metric_functions attribute type is not yet supported, so WithDataMetricFunctions can't be generated

// data_metric_schedule attribute type is not yet supported, so WithDataMetricSchedule can't be generated

func (v *ViewModel) WithDatabase(database string) *ViewModel {
	v.Database = tfconfig.StringVariable(database)
	return v
}

func (v *ViewModel) WithIsRecursive(isRecursive string) *ViewModel {
	v.IsRecursive = tfconfig.StringVariable(isRecursive)
	return v
}

func (v *ViewModel) WithIsSecure(isSecure string) *ViewModel {
	v.IsSecure = tfconfig.StringVariable(isSecure)
	return v
}

func (v *ViewModel) WithIsTemporary(isTemporary string) *ViewModel {
	v.IsTemporary = tfconfig.StringVariable(isTemporary)
	return v
}

func (v *ViewModel) WithName(name string) *ViewModel {
	v.Name = tfconfig.StringVariable(name)
	return v
}

func (v *ViewModel) WithOrReplace(orReplace bool) *ViewModel {
	v.OrReplace = tfconfig.BoolVariable(orReplace)
	return v
}

// row_access_policy attribute type is not yet supported, so WithRowAccessPolicy can't be generated

func (v *ViewModel) WithSchema(schema string) *ViewModel {
	v.Schema = tfconfig.StringVariable(schema)
	return v
}

func (v *ViewModel) WithStatement(statement string) *ViewModel {
	v.Statement = tfconfig.StringVariable(statement)
	return v
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (v *ViewModel) WithAggregationPolicyValue(value tfconfig.Variable) *ViewModel {
	v.AggregationPolicy = value
	return v
}

func (v *ViewModel) WithChangeTrackingValue(value tfconfig.Variable) *ViewModel {
	v.ChangeTracking = value
	return v
}

func (v *ViewModel) WithColumnsValue(value tfconfig.Variable) *ViewModel {
	v.Columns = value
	return v
}

func (v *ViewModel) WithCommentValue(value tfconfig.Variable) *ViewModel {
	v.Comment = value
	return v
}

func (v *ViewModel) WithCopyGrantsValue(value tfconfig.Variable) *ViewModel {
	v.CopyGrants = value
	return v
}

func (v *ViewModel) WithDataMetricFunctionsValue(value tfconfig.Variable) *ViewModel {
	v.DataMetricFunctions = value
	return v
}

func (v *ViewModel) WithDataMetricScheduleValue(value tfconfig.Variable) *ViewModel {
	v.DataMetricSchedule = value
	return v
}

func (v *ViewModel) WithDatabaseValue(value tfconfig.Variable) *ViewModel {
	v.Database = value
	return v
}

func (v *ViewModel) WithIsRecursiveValue(value tfconfig.Variable) *ViewModel {
	v.IsRecursive = value
	return v
}

func (v *ViewModel) WithIsSecureValue(value tfconfig.Variable) *ViewModel {
	v.IsSecure = value
	return v
}

func (v *ViewModel) WithIsTemporaryValue(value tfconfig.Variable) *ViewModel {
	v.IsTemporary = value
	return v
}

func (v *ViewModel) WithNameValue(value tfconfig.Variable) *ViewModel {
	v.Name = value
	return v
}

func (v *ViewModel) WithOrReplaceValue(value tfconfig.Variable) *ViewModel {
	v.OrReplace = value
	return v
}

func (v *ViewModel) WithRowAccessPolicyValue(value tfconfig.Variable) *ViewModel {
	v.RowAccessPolicy = value
	return v
}

func (v *ViewModel) WithSchemaValue(value tfconfig.Variable) *ViewModel {
	v.Schema = value
	return v
}

func (v *ViewModel) WithStatementValue(value tfconfig.Variable) *ViewModel {
	v.Statement = value
	return v
}