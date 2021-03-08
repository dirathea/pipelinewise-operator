package v1alpha1

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

// PipelinewiseTargetID defines pipelinewise target id
type PipelinewiseTargetID string

// PipelinewiseTargetType defines pipelinewise target type
type PipelinewiseTargetType string

// TargetInfo basic Target information
// +kubebuilder:object:generate=false
type TargetInfo interface {
	ConnectorID() string
	ID() PipelinewiseTargetID
	Type() PipelinewiseTargetType
	GetConnection() interface{}
}

const (
	// PostgreSQLTargetID defines Pipelinewise PostgreSQL Target ID
	PostgreSQLTargetID PipelinewiseTargetID = "postgres"

	// RedshiftTargetID defines redshift target ID
	RedshiftTargetID PipelinewiseTargetID = "redshift"

	// SnowflakeTargetID defines snowflake target ID
	SnowflakeTargetID PipelinewiseTargetID = "snowflake"

	// S3CSVTargetID defines snowflake target ID
	S3CSVTargetID PipelinewiseTargetID = "s3-csv"

	// PostgreSQLTargetType defines PostgreSQL Pipelinewise Target type
	PostgreSQLTargetType PipelinewiseTargetType = "target-postgres"

	// RedshiftTargetType defines Redshift Pipelinewise Target type
	RedshiftTargetType PipelinewiseTargetType = "target-redshift"

	// SnowflakeTargetType defines Snowflake Pipelinewise Target type
	SnowflakeTargetType PipelinewiseTargetType = "target-snowflake"

	// S3CSVTargetType defines Snowflake Pipelinewise Target type
	S3CSVTargetType PipelinewiseTargetType = "target-s3-csv"
)

// GenericTargetSpec defines generic Pipelinewise Target configuration
// +kubebuilder:object:generate=false
type GenericTargetSpec struct {
	ID                 PipelinewiseTargetID   `yaml:"id"`
	Name               string                 `yaml:"name"`
	Type               PipelinewiseTargetType `yaml:"type"`
	DatabaseConnection interface{}            `yaml:"db_conn"`
}

// PostgreSQLTargetSpec defines PostgreSQL Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/postgres.html)
type PostgreSQLTargetSpec struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	DatabaseName string `yaml:"dbname" json:"dbname"`
}

// ConnectorID implements TargetInfo interface to return connection id
func (ts *PostgreSQLTargetSpec) ConnectorID() string {
	return string(PostgreSQLTargetID)
}

// ID implements TargetInfo interface to return target id
func (ts *PostgreSQLTargetSpec) ID() PipelinewiseTargetID {
	return PipelinewiseTargetID(fmt.Sprintf("%v-%v", PostgreSQLTargetID, ts.DatabaseName))
}

// Type implements TargetInfo interface to return target type
func (ts *PostgreSQLTargetSpec) Type() PipelinewiseTargetType {
	return PostgreSQLTargetType
}

// GetConnection implements TargetInfo interface to return connection info
func (ts *PostgreSQLTargetSpec) GetConnection() interface{} {
	return ts
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

// ConnectorID implements TargetInfo interface to return connection id
func (ts *RedshiftTargetSpec) ConnectorID() string {
	return string(RedshiftTargetID)
}

// ID implements TargetInfo interface to return target id
func (ts *RedshiftTargetSpec) ID() PipelinewiseTargetID {
	return PipelinewiseTargetID(fmt.Sprintf("%v-%v", RedshiftTargetID, ts.DatabaseName))
}

// Type implements TargetInfo interface to return target type
func (ts *RedshiftTargetSpec) Type() PipelinewiseTargetType {
	return RedshiftTargetType
}

// GetConnection implements TargetInfo interface to return connection info
func (ts *RedshiftTargetSpec) GetConnection() interface{} {
	return ts
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

// ConnectorID implements TargetInfo interface to return connection id
func (ts *SnowflakeTargetSpec) ConnectorID() string {
	return string(SnowflakeTargetID)
}

// ID implements TargetInfo interface to return target id
func (ts *SnowflakeTargetSpec) ID() PipelinewiseTargetID {
	return PipelinewiseTargetID(fmt.Sprintf("%v-%v", SnowflakeTargetID, ts.DatabaseName))
}

// Type implements TargetInfo interface to return target type
func (ts *SnowflakeTargetSpec) Type() PipelinewiseTargetType {
	return SnowflakeTargetType
}

// GetConnection implements TargetInfo interface to return connection info
func (ts *SnowflakeTargetSpec) GetConnection() interface{} {
	return ts
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
	QuoteChar          string `yaml:"quotechar,omitempty" json:"quotechar,omitempty"`
	EncryptionType     string `yaml:"encryption_type,omitempty" json:"encryption_type,omitempty"`
	EncryptionKey      string `yaml:"encryption_key,omitempty" json:"encryption_key,omitempty"`
}

// ConnectorID implements TargetInfo interface to return connection id
func (ts *S3CSVTargetSpec) ConnectorID() string {
	return string(S3CSVTargetID)
}

// ID implements TargetInfo interface to return target id
func (ts *S3CSVTargetSpec) ID() PipelinewiseTargetID {
	return PipelinewiseTargetID(fmt.Sprintf("%v-%v", S3CSVTargetID, ts.S3Bucket))
}

// Type implements TargetInfo interface to return target type
func (ts *S3CSVTargetSpec) Type() PipelinewiseTargetType {
	return S3CSVTargetType
}

// GetConnection implements TargetInfo interface to return connection info
func (ts *S3CSVTargetSpec) GetConnection() interface{} {
	return ts
}

func getTargetInfo(pwJob *PipelinewiseJob) TargetInfo {
	pwVal := reflect.ValueOf(pwJob.Spec.Target)
	for fieldNth := 0; fieldNth < pwVal.NumField(); fieldNth++ {
		field := pwVal.Field(fieldNth)
		if !field.IsNil() {
			return field.Interface().(TargetInfo)
		}
	}
	return nil
}

// GetTargetConnectorID defines pipelinewise target connector id
func GetTargetConnectorID(pwJob *PipelinewiseJob) string {
	targetInfo := getTargetInfo(pwJob)
	if targetInfo != nil {
		return targetInfo.ConnectorID()
	}
	return ""
}

// GetTargetID calculate pipelinewise target id
func GetTargetID(pipelinewiseJob *PipelinewiseJob) PipelinewiseTargetID {
	targetInfo := getTargetInfo(pipelinewiseJob)
	if targetInfo != nil {
		return targetInfo.ID()
	}
	return ""
}

// ConstructTargetConfiguration parse and return a target yaml configuration string
func ConstructTargetConfiguration(pwJob *PipelinewiseJob) ([]byte, error) {
	targetInfo := getTargetInfo(pwJob)

	if targetInfo != nil {
		return constructTarget(targetInfo.ID(), targetInfo.Type(), targetInfo.GetConnection())
	}

	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

func constructTarget(pwID PipelinewiseTargetID, pwType PipelinewiseTargetType, dbConn interface{}) ([]byte, error) {
	targetConfiguration := GenericTargetSpec{
		DatabaseConnection: dbConn,
		ID:                 pwID,
		Name:               string(pwID),
		Type:               pwType,
	}
	return yaml.Marshal(targetConfiguration)
}
