package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

const (
	// PostgreSQLPipelinewiseID defines Pipelinewise PostgreSQL Target ID
	PostgreSQLPipelinewiseID PipelinewiseID = "postgresql"

	// RedshiftPipelinewiseID defines redshift target ID
	RedshiftPipelinewiseID PipelinewiseID = "redshift"

	// SnowflakePipelinewiseID defines snowflake target ID
	SnowflakePipelinewiseID PipelinewiseID = "snowflake"

	// S3CSVPipelinewiseID defines snowflake target ID
	S3CSVPipelinewiseID PipelinewiseID = "s3_csv"

	// PostgreSQLTargetType defines PostgreSQL Pipelinewise Target type
	PostgreSQLTargetType PipelinewiseType = "target-postgres"

	// RedshiftTargetType defines Redshift Pipelinewise Target type
	RedshiftTargetType PipelinewiseType = "target-redshift"

	// SnowflakeTargetType defines Snowflake Pipelinewise Target type
	SnowflakeTargetType PipelinewiseType = "target-snowflake"

	// S3CAVTargetType defines Snowflake Pipelinewise Target type
	S3CAVTargetType PipelinewiseType = "target-s3-csv"
)

// PostgreSQLTargetSpec defines PostgreSQL Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/postgres.html)
type PostgreSQLTargetSpec struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	DatabaseName string `yaml:"dbname" json:"dbname"`
}

// RedshiftTargetSpec defines Redshift Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/redshift.html)
type RedshiftTargetSpec struct {
	Host                   string `yaml:"host" json:"host"`
	Port                   int    `yaml:"port" json:"port"`
	User                   string `yaml:"user" json:"user"`
	Password               string `yaml:"password" json:"password"`
	DatabaseName           string `yaml:"dbname" json:"dbname"`
	AWSProfile             string `yaml:"aws_profile,omitempty" json:"aws_profile,omitempty"`
	AWSAccessKeyID         string `yaml:"aws_access_key_id,omitempty" json:"aws_access_key_id,omitempty"`
	AWSAccessSecretKey     string `yaml:"aws_secret_access_key,omitempty" json:"aws_secret_access_key,omitempty"`
	AWSSessionToken        string `yaml:"aws_session_token,omitempty" json:"aws_session_token,omitempty"`
	AWSRedshiftCopyRoleARN string `yaml:"aws_redshift_copy_role_arn,omitempty" json:"aws_redshift_copy_role_arn,omitempty"`
	S3Bucket               string `yaml:"s3_bucket" json:"s3_bucket"`
	S3KeyPrefix            string `yaml:"s3_key_prefix,omitempty" json:"s3_key_prefix,omitempty"`
	S3ACL                  string `yaml:"s3_acl,omitempty" json:"s3_acl,omitempty"`
	CopyOptions            string `yaml:"copy_options" json:"copy_options"`
}

// SnowflakeTargetSpec defines Snowflake Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/snowflake.html)
type SnowflakeTargetSpec struct {
	Account                       string `yaml:"account" json:"account"`
	DatabaseName                  string `yaml:"dbname" json:"dbname"`
	User                          string `yaml:"user" json:"user"`
	Password                      string `yaml:"password" json:"password"`
	Warehouse                     string `yaml:"warehouse" json:"warehouse"`
	AWSProfile                    string `yaml:"aws_profile,omitempty" json:"aws_profile,omitempty"`
	AWSAccessKeyID                string `yaml:"aws_access_key_id,omitempty" json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey            string `yaml:"aws_secret_access_key,omitempty" json:"aws_secret_access_key,omitempty"`
	AWSSessionToken               string `yaml:"aws_session_token,omitempty" json:"aws_session_token,omitempty"`
	AWSEndpointURL                string `yaml:"aws_session_url" json:"aws_session_url"`
	S3Bucket                      string `yaml:"s3_bucket" json:"s3_bucket"`
	S3KeyPrefix                   string `yaml:"s3_key_prefix,omitempty" json:"s3_key_prefix,omitempty"`
	S3ACL                         string `yaml:"s3_acl,omitempty" json:"s3_acl,omitempty"`
	Stage                         string `yaml:"schema" json:"schema"`
	FileFormat                    string `yaml:"file_format" json:"file_format"`
	ClientSideEncryptionMasterKey string `yaml:"client_side_encryption_master_key,omitempty" json:"client_side_encryption_master_key,omitempty"`
}

// S3CSVTargetSpec defines S3 CSV Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/s3_csv.html)
type S3CSVTargetSpec struct {
	AWSProfile         string `yaml:"aws_profile,omitempty" json:"aws_profile,omitempty"`
	AWSAccessKeyID     string `yaml:"aws_access_key_id,omitempty" json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey string `yaml:"aws_secret_access_key,omitempty" json:"aws_secret_access_key,omitempty"`
	AWSSessionToken    string `yaml:"aws_session_token,omitempty" json:"aws_session_token,omitempty"`
	S3Bucket           string `yaml:"s3_bucket" json:"s3_bucket"`
	S3KeyPrefix        string `yaml:"s3_key_prefix,omitempty" json:"s3_key_prefix,omitempty"`
	Delimiter          string `yaml:"delimiter,omitempty" json:"delimiter,omitempty"`
	QuoteChar          string `yaml:"quotechar" json:"quotechar"`
	EncryptionType     string `yaml:"encryption_type,omitempty" json:"encryption_type,omitempty"`
	EncryptionKey      string `yaml:"encryption_key,omitempty" json:"encryption_key,omitempty"`
}

// GetTargetID calculate pipelinewise target id
func GetTargetID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Target.Snowflake != nil {
		return string(SnowflakePipelinewiseID)
	}
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return string(RedshiftPipelinewiseID)
	}
	if pipelinewiseJob.Spec.Target.PostgreSQL != nil {
		return string(PostgreSQLPipelinewiseID)
	}
	if pipelinewiseJob.Spec.Target.S3CSV != nil {
		return string(S3CSVPipelinewiseID)
	}
	return ""
}

// ConstructTargetConfiguration parse and return a target yaml configuration string
func ConstructTargetConfiguration(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	// Find target
	if pipelinewiseJob.Spec.Target.Snowflake != nil {
		return constructSnowflakeTarget(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return constructRedshiftTarget(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Target.PostgreSQL != nil {
		return constructPostgreSQLTarget(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Target.S3CSV != nil {
		return constructS3CSVTarget(pipelinewiseJob)
	}
	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

func constructSnowflakeTarget(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	type SnowflakeTargetConfiguration struct {
		ID                 string               `yaml:"id"`
		Name               string               `yaml:"name"`
		Type               PipelinewiseType     `yaml:"type"`
		DatabaseConnection *SnowflakeTargetSpec `yaml:"db_conn"`
	}
	snowflakeConfiguration := SnowflakeTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.Snowflake,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "Snowflake",
		Type:               SnowflakeTargetType,
	}
	return yaml.Marshal(snowflakeConfiguration)
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
		Name:               "PostgreSQL",
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
		Name:               "Redshift",
		Type:               MySQLTapType,
	}
	return yaml.Marshal(redshiftConfiguration)
}

func constructS3CSVTarget(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	type S3CSVTargetConfiguration struct {
		ID                 string           `yaml:"id"`
		Name               string           `yaml:"name"`
		Type               PipelinewiseType `yaml:"type"`
		DatabaseConnection *S3CSVTargetSpec `yaml:"db_conn"`
	}
	s3csvConfiguration := S3CSVTargetConfiguration{
		DatabaseConnection: pipelinewiseJob.Spec.Target.S3CSV,
		ID:                 GetTargetID(pipelinewiseJob),
		Name:               "S3 CSV",
		Type:               S3CAVTargetType,
	}
	return yaml.Marshal(s3csvConfiguration)
}
