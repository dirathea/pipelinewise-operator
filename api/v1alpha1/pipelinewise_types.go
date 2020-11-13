package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// PipelinewiseType defines configuration type. Could be a `tap` or a `target` with defined application type
type PipelinewiseType string

type PipelinewiseID string

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseType = "mysql"
	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseType = "tap-mysql"
	// PostgreSQLPipelinewiseID defines Pipelinewise PostgreSQL Target ID
	PostgreSQLPipelinewiseID PipelinewiseID = "postgresql"
	// PostgreSQLTargetType defines Pipelinewise PostgreSQL Target type
	PostgreSQLTargetType PipelinewiseType = "target-postgresql"
	// RedshiftPipelinewiseID defines redshift target ID
	RedshiftPipelinewiseID PipelinewiseID = "redshift"
	// RedshiftTargetType defines Pipelinewise Target type
	RedshiftTargetType PipelinewiseType = "target-redshift"
)

// PipelinewiseGenericConfiguration defines generic pipelinewise configuration
type PipelinewiseGenericConfiguration struct {
	ID   string           `json:"id"`
	Name string           `json:"name"`
	Type PipelinewiseType `json:"type"`
}

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pipelinewiseJob *PipelinewiseJob) (string, error) {
	// Find tap
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		// MySQL Tap found
		type MySQLTapConfiguration struct {
			ID                 string           `json:"id"`
			Name               string           `json:"name"`
			Type               PipelinewiseType `json:"type"`
			DatabaseConnection *MySQLTapSpec    `json:"db_conn"`
			Target             string           `json:"target"`
		}
		mysqlConfiguration := MySQLTapConfiguration{
			DatabaseConnection: pipelinewiseJob.Spec.Tap.MySQL,
			ID:                 GetTapID(pipelinewiseJob),
			Name:               "Tap Mysql to Target with suffix",
			Type:               MySQLTapType,
			Target:             GetTargetID(pipelinewiseJob),
		}
		configurationBytes, err := yaml.Marshal(mysqlConfiguration)
		if err != nil {
			return "", err
		}
		return string(configurationBytes), err
	}
	return "", fmt.Errorf("No Valid Tap configured")
}

// ConstructTargetConfiguration parse and return a target yaml configuration string
func ConstructTargetConfiguration(pipelinewiseJob *PipelinewiseJob) (string, error) {
	// Find target
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return constructRedshiftTarget(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Target.PostgreSQL != nil {
		return constructPostgreSQLTarget(pipelinewiseJob)
	}
	return "", fmt.Errorf("No Valid Tap configured")
}

func constructPostgreSQLTarget(pipelinewiseJob *PipelinewiseJob) (string, error) {
	type PostgreSQLTargetConfiguration struct {
		ID                 string                `json:"id"`
		Name               string                `json:"name"`
		Type               PipelinewiseType      `json:"type"`
		DatabaseConnection *PostgreSQLTargetSpec `json:"db_conn"`
	}
	postgreSQLConfiguration := PostgreSQLTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.PostgreSQL,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "PostgreSQL Target with suffix",
		Type:               PostgreSQLTargetType,
	}
	configurationBytes, err := yaml.Marshal(postgreSQLConfiguration)
	if err != nil {
		return "", err
	}
	return string(configurationBytes), err
}

func constructRedshiftTarget(pipelinewiseJob *PipelinewiseJob) (string, error) {
	type RedshiftTargetConfiguration struct {
		ID                 string              `json:"id"`
		Name               string              `json:"name"`
		Type               PipelinewiseType    `json:"type"`
		DatabaseConnection *RedshiftTargetSpec `json:"db_conn"`
	}
	redshiftConfiguration := RedshiftTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.Redshift,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "Redshift Target with suffix",
		Type:               MySQLTapType,
	}
	configurationBytes, err := yaml.Marshal(redshiftConfiguration)
	if err != nil {
		return "", err
	}
	return string(configurationBytes), err
}

// GetTargetID calculate pipelinewise target id
func GetTargetID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return string(RedshiftPipelinewiseID)
	}
	if pipelinewiseJob.Spec.Target.PostgreSQL != nil {
		return string(PostgreSQLPipelinewiseID)
	}
	return ""
}

// GetTapID calculate pipelinewise tap id
func GetTapID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		return fmt.Sprintf("%v-%v", MySQLTapID, pipelinewiseJob.Spec.Tap.MySQL.Connection.DatabaseName)
	}
	return ""
}
