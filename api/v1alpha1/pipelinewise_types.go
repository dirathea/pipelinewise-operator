package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// PipelinewiseType defines configuration type. Could be a `tap` or a `target` with defined application type
type PipelinewiseType string

// PipelinewiseID defines configuration ID
type PipelinewiseID string

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseType = "mysql"
	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseType = "tap-mysql"
	// PostgreSQLPipelinewiseID defines Pipelinewise PostgreSQL Target ID
	PostgreSQLPipelinewiseID PipelinewiseID = "postgresql"
	// PostgreSQLTargetType defines Pipelinewise PostgreSQL Target type
	PostgreSQLTargetType PipelinewiseType = "target-postgres"
	// RedshiftPipelinewiseID defines redshift target ID
	RedshiftPipelinewiseID PipelinewiseID = "redshift"
	// RedshiftTargetType defines Pipelinewise Target type
	RedshiftTargetType PipelinewiseType = "target-redshift"
)

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	// Find tap
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		// MySQL Tap found
		type MySQLTapConfiguration struct {
			ID                 string                 `yaml:"id"`
			Name               string                 `yaml:"name"`
			Type               PipelinewiseType       `yaml:"type"`
			DatabaseConnection MySQLTapConnectionSpec `yaml:"db_conn"`
			Target             string                 `yaml:"target"`
			Schemas            []MySQLTapSchemaSpec   `yaml:"schemas"`
		}
		mysqlConfiguration := MySQLTapConfiguration{
			DatabaseConnection: pipelinewiseJob.Spec.Tap.MySQL.Connection,
			ID:                 GetTapID(pipelinewiseJob),
			Name:               "Tap Mysql to Target with suffix",
			Type:               MySQLTapType,
			Target:             GetTargetID(pipelinewiseJob),
			Schemas:            pipelinewiseJob.Spec.Tap.MySQL.Schemas,
		}
		return yaml.Marshal(mysqlConfiguration)
	}
	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

// ConstructTargetConfiguration parse and return a target yaml configuration string
func ConstructTargetConfiguration(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	// Find target
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return constructRedshiftTarget(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Target.PostgreSQL != nil {
		return constructPostgreSQLTarget(pipelinewiseJob)
	}
	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

func constructPostgreSQLTarget(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	type PostgreSQLTargetConfiguration struct {
		ID                 string                `yaml:"id"`
		Name               string                `yaml:"name"`
		Type               PipelinewiseType      `yaml:"type"`
		DatabaseConnection *PostgreSQLTargetSpec `yaml:"db_conn"`
	}
	postgreSQLConfiguration := PostgreSQLTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.PostgreSQL,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "PostgreSQL Target with suffix",
		Type:               PostgreSQLTargetType,
	}
	return yaml.Marshal(postgreSQLConfiguration)
}

func constructRedshiftTarget(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	type RedshiftTargetConfiguration struct {
		ID                 string              `yaml:"id"`
		Name               string              `yaml:"name"`
		Type               PipelinewiseType    `yaml:"type"`
		DatabaseConnection *RedshiftTargetSpec `yaml:"db_conn"`
	}
	redshiftConfiguration := RedshiftTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.Redshift,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "Redshift Target with suffix",
		Type:               MySQLTapType,
	}
	return yaml.Marshal(redshiftConfiguration)
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
