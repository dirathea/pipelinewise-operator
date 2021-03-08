package v1alpha1

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

// PipelinewiseTapID defines tap id
type PipelinewiseTapID string

// PipelinewiseTapType defines tap type
type PipelinewiseTapType string

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseTapID = "mysql"
	// PostgreSQLTapID defines Pipelinewise Postgres Tap ID
	PostgreSQLTapID PipelinewiseTapID = "postgres"
	// OracleTapID defines Pipelinewise Oracle Tap ID
	OracleTapID PipelinewiseTapID = "oracle"
	// KafkaTapID defines Pipelinewise Kafka Tap ID
	KafkaTapID PipelinewiseTapID = "kafka"
	// S3CSVTapID defines Pipelinewise S3 CSV Tap ID
	S3CSVTapID PipelinewiseTapID = "s3-csv"
	// SnowflakeTapID defines Pipelinewise Snowflake Tap ID
	SnowflakeTapID PipelinewiseTapID = "snowflake"
	// MongoDBTapID defines Pipelinewise MongoDB Tap ID
	MongoDBTapID PipelinewiseTapID = "mongodb"
	// SalesforceTapID defines Pipelinewise Salesforce Tap ID
	SalesforceTapID PipelinewiseTapID = "salesforce"
	// ZendeskTapID defines Pipelinewise Zendesk Tap ID
	ZendeskTapID PipelinewiseTapID = "zendesk"
	// JiraTapID defines Pipelinewise Jira Tap ID
	JiraTapID PipelinewiseTapID = "jira"
	// ZuoraTapID defines Pipelinewise Zuora Tap ID
	ZuoraTapID PipelinewiseTapID = "zuora"
	// GoogleAnalyticsTapID defines Pipelinewise Google Analytics Tap ID
	GoogleAnalyticsTapID PipelinewiseTapID = "google-analytics"
	// GithubTapID defines Pipelinewise Github Tap ID
	GithubTapID PipelinewiseTapID = "github"
	// ShopifyTapID defines Pipelinewise Shopify Tap ID
	ShopifyTapID PipelinewiseTapID = "shopify"
	// SlackTapID defines Pipelinewise Slack Tap ID
	SlackTapID PipelinewiseTapID = "slack"
	// MixpanelTapID defines Pipelinewise Mixpanel Tap ID
	MixpanelTapID PipelinewiseTapID = "mixpanel"
	// TwilioTapID defines Pipelinewise Twilio Tap ID
	TwilioTapID PipelinewiseTapID = "twilio"

	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseTapType = "tap-mysql"
	// PostgreSQLTapType defines Pipelinewise Postgres Tap type
	PostgreSQLTapType PipelinewiseTapType = "tap-postgres"
	// OracleTapType defines Pipelinewise Oracle Tap type
	OracleTapType PipelinewiseTapType = "tap-oracle"
	// KafkaTapType defines Pipelinewise Kafka Tap type
	KafkaTapType PipelinewiseTapType = "tap-kafka"
	// S3CSVTapType defines Pipelinewise S3 CSV Tap type
	S3CSVTapType PipelinewiseTapType = "tap-s3-csv"
	// SnowflakeTapType defines Pipelinewise Snowflake Tap type
	SnowflakeTapType PipelinewiseTapType = "tap-snowflake"
	// MongoDBTapType defines Pipelinewise MongoDB Tap type
	MongoDBTapType PipelinewiseTapType = "tap-mongodb"
	// SalesforceTapType defines Pipelinewise Salesforce Tap type
	SalesforceTapType PipelinewiseTapType = "tap-salesforce"
	// ZendeskTapType defines Pipelinewise Zendesk Tap type
	ZendeskTapType PipelinewiseTapType = "tap-zendesk"
	// JiraTapType defines Pipelinewise Jira Tap type
	JiraTapType PipelinewiseTapType = "tap-jira"
	// ZuoraTapType defines Pipelinewise Zuora Tap type
	ZuoraTapType PipelinewiseTapType = "tap-zuora"
	// GoogleAnalyticsTapType defines Pipelinewise Google Analytics Tap type
	GoogleAnalyticsTapType PipelinewiseTapType = "tap-google-analytics"
	// GithubTapType defines Pipelinewise Github Tap type
	GithubTapType PipelinewiseTapType = "tap-github"
	// ShopifyTapType defines Pipelinewise Shopify Tap type
	ShopifyTapType PipelinewiseTapType = "tap-shopify"
	// SlackTapType defines Pipelinewise Slack Tap type
	SlackTapType PipelinewiseTapType = "tap-slack"
	// MixpanelTapType defines Pipelinewise Mixpanel Tap type
	MixpanelTapType PipelinewiseTapType = "tap-mixpanel"
	// TwiliolTapType defines Pipelinewise Twilio Tap type
	TwiliolTapType PipelinewiseTapType = "tap-twilio"
)

// TapInfo basic Tap information
// +kubebuilder:object:generate=false
type TapInfo interface {
	ConnectorID() string
	ID() PipelinewiseTapID
	Type() PipelinewiseTapType
	GetSchemas() interface{}
	GetConnection() interface{}
}

// GenericTapSpec defines generic Pipelinewise Tap configuration
// +kubebuilder:object:generate=false
type GenericTapSpec struct {
	ID                  PipelinewiseTapID    `yaml:"id"`
	Name                string               `yaml:"name"`
	Type                PipelinewiseTapType  `yaml:"type"`
	Owner               string               `yaml:"owner,omitempty" json:"owner,omitempty"`
	DefaultTargetSchema string               `yaml:"default_target_schema,omitempty" json:"default_target_schema,omitempty"`
	DatabaseConnection  interface{}          `yaml:"db_conn"`
	Target              PipelinewiseTargetID `yaml:"target"`
	Schemas             interface{}          `yaml:"schemas"`
}

// TapTableSpec defines Generic Tap Table configuration
type TapTableSpec struct {
	TableName         string `yaml:"table_name" json:"table_name"`
	ReplicationMethod string `yaml:"replication_method" json:"replication_method"`
	ReplicationKey    string `yaml:"replication_key,omitempty" json:"replication_key,omitempty"`
}

// TapSchemaSpec defines Generic Tap schema configuration
type TapSchemaSpec struct {
	Source string         `yaml:"source_schema" json:"source_schema"`
	Target string         `yaml:"target_schema" json:"target_schema"`
	Tables []TapTableSpec `yaml:"tables" json:"tables"`
}

// MySQLTapConnectionSpec defines MySQL Tap connection configuration
type MySQLTapConnectionSpec struct {
	Host            string   `yaml:"host" json:"host"`
	Port            int      `yaml:"port" json:"port"`
	User            string   `yaml:"user" json:"user"`
	Password        string   `yaml:"password" json:"password"`
	DBName          string   `yaml:"dbname" json:"dbname"`
	FilterDatabases string   `yaml:"filter_dbs,omitempty" json:"filter_dbs,omitempty"`
	ExportBatchRows int      `yaml:"export_batch_rows,omitempty" json:"export_batch_rows,omitempty"`
	SessionSQLs     []string `yaml:"session_sqls,omitempty" json:"session_sqls,omitempty"`
}

// MySQLTapSpec defines Tap configuration for MySQL. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/mysql.html)
type MySQLTapSpec struct {
	Schemas          []TapSchemaSpec        `yaml:"schemas" json:"schemas"`
	Connection       MySQLTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows    int                    `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize int                    `yaml:"stream_buffer_size" json:"stream_buffer_size"`
}

// ConnectorID return MySQL connector ID
func (ts *MySQLTapSpec) ConnectorID() string {
	return string(MySQLTapID)
}

// ID return MySQL Tap ID
func (ts *MySQLTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", MySQLTapID, ts.Connection.DBName))
}

// Type implement TapInfo interface to return Pipelinewise Type
func (ts *MySQLTapSpec) Type() PipelinewiseTapType {
	return MySQLTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *MySQLTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *MySQLTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// PostgreSQLTapConnectionSpec defines Postgres tap connection configuration
type PostgreSQLTapConnectionSpec struct {
	Host                    string `yaml:"host" json:"host"`
	Port                    int    `yaml:"port" json:"port"`
	User                    string `yaml:"user" json:"user"`
	Password                string `yaml:"password" json:"password"`
	DBName                  string `yaml:"dbname" json:"dbname"`
	FilterSchemas           string `yaml:"filter_schemas,omitempty" json:"filter_schemas,omitempty"`
	MaxRunSeconds           int    `yaml:"max_run_seconds,omitempty" json:"max_run_seconds,omitempty"`
	LogicalPollTotalSeconds int    `yaml:"logical_poll_total_seconds,omitempty" json:"logical_poll_total_seconds,omitempty"`
	BreakAtEndLSN           bool   `yaml:"break_at_end_lsn,omitempty" json:"break_at_end_lsn,omitempty"`
	SSL                     bool   `yaml:"ssl,omitempty" json:"ssl,omitempty"`
}

// PostgreSQLTapSpec defines Tap configuration for PostgreSQL. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/postgres.html)
type PostgreSQLTapSpec struct {
	Schemas          []TapSchemaSpec             `yaml:"schemas" json:"schemas"`
	Connection       PostgreSQLTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows    int                         `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize int                         `yaml:"stream_buffer_size" json:"stream_buffer_size"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *PostgreSQLTapSpec) ConnectorID() string {
	return string(PostgreSQLTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *PostgreSQLTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", PostgreSQLTapID, ts.Connection.DBName))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *PostgreSQLTapSpec) Type() PipelinewiseTapType {
	return PostgreSQLTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *PostgreSQLTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *PostgreSQLTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// OracleTapConnectionSpec defines Oracle tap connection configuration
type OracleTapConnectionSpec struct {
	SID           string `yaml:"sid" json:"sid"`
	Host          string `yaml:"host" json:"host"`
	Port          int    `yaml:"port" json:"port"`
	User          string `yaml:"user" json:"user"`
	Password      string `yaml:"password" json:"password"`
	FilterSchemas string `yaml:"filter_schemas,omitempty" json:"filter_schemas,omitempty"`
}

// OracleTapSpec defines Tap configuration for Oracle. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/oracle.html)
type OracleTapSpec struct {
	Schemas          []TapSchemaSpec         `yaml:"schemas" json:"schemas"`
	Connection       OracleTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows    int                     `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize int                     `yaml:"stream_buffer_size" json:"stream_buffer_size"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *OracleTapSpec) ConnectorID() string {
	return string(OracleTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *OracleTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", OracleTapID, ts.Connection.SID))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *OracleTapSpec) Type() PipelinewiseTapType {
	return OracleTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *OracleTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *OracleTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// KafkaTapPrimaryKey defines Kafka tap connection primary key
type KafkaTapPrimaryKey struct {
	TransferID string `yaml:"transfer_id" json:"transfer_id"`
}

// KafkaTapConnectionSpec defines Kafka tap connection configuration
type KafkaTapConnectionSpec struct {
	GroupID                 string              `yaml:"group_id" json:"group_id"`
	BootstrapServers        string              `yaml:"bootstrap_servers" json:"bootstrap_servers"`
	Topic                   string              `yaml:"topic" json:"topic"`
	PrimaryKeys             *KafkaTapPrimaryKey `yaml:"primary_keys,omitempty" json:"primary_keys,omitempty"`
	MaxRuntimeMs            *int                `yaml:"max_runtime_ms,omitempty" json:"max_runtime_ms,omitempty"`
	ConsumerTimeoutMs       *int                `yaml:"consumer_timeout_ms,omitempty" json:"consumer_timeout_ms,omitempty"`
	SessionTimeoutMs        *int                `yaml:"session_timeout_ms,omitempty" json:"session_timeout_ms,omitempty"`
	HeartbeatIntervalMs     *int                `yaml:"heartbeat_interval_ms,omitempty" json:"heartbeat_interval_ms,omitempty"`
	MaxPollIntervalMs       *int                `yaml:"max_poll_interval_ms,omitempty" json:"max_poll_interval_ms,omitempty"`
	MaxPollRecords          *int                `yaml:"max_poll_records,omitempty" json:"max_poll_records,omitempty"`
	CommitIntervalMs        *int                `yaml:"commit_interval_ms,omitempty" json:"commit_interval_ms,omitempty"`
	LocalStoreDir           string              `yaml:"local_store_dir,omitempty" json:"local_store_dir,omitempty"`
	LocalStoreBatchSizeRows *int                `yaml:"local_store_batch_size_rows,omitempty" json:"local_store_batch_size_rows,omitempty"`
}

// KafkaTapSpec defines Tap configuration for Kafka. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/kafka.html)
type KafkaTapSpec struct {
	Schemas    []TapSchemaSpec        `yaml:"schemas" json:"schemas"`
	Connection KafkaTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *KafkaTapSpec) ConnectorID() string {
	return string(KafkaTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *KafkaTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", KafkaTapID, ts.Connection.Topic))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *KafkaTapSpec) Type() PipelinewiseTapType {
	return KafkaTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *KafkaTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *KafkaTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// S3CSVTableMappingSpec defines S3 CSV Table Mapping
type S3CSVTableMappingSpec struct {
	SearchPattern string   `yaml:"search_pattern" json:"search_pattern"`
	SearchPrefix  string   `yaml:"search_prefix,omitempty" json:"search_prefix,omitempty"`
	KeyProperties []string `yaml:"key_properties,omitempty" json:"key_properties,omitempty"`
	Delimiter     string   `yaml:"delimiter" json:"delimiter"`
}

// S3CSVTapTableSpec defines S3 CSV Tap Table configuration
type S3CSVTapTableSpec struct {
	TableName string                `yaml:"table_name" json:"table_name"`
	Mapping   S3CSVTableMappingSpec `yaml:"s3_csv_mapping" json:"s3_csv_mapping"`
}

// S3CSVTapSchemaSpec defines S3 CSV Tap schema configuration
type S3CSVTapSchemaSpec struct {
	Source string              `yaml:"source_schema" json:"source_schema"`
	Target string              `yaml:"target_schema" json:"target_schema"`
	Tables []S3CSVTapTableSpec `yaml:"tables" json:"tables"`
}

// S3CSVTapConnectionSpec defines S3 CSV Tap connection specification
type S3CSVTapConnectionSpec struct {
	AWSProfile         string `yaml:"aws_profile,omitempty" json:"aws_profile,omitempty"`
	AWSAccessKeyID     string `yaml:"aws_access_key_id,omitempty" json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey string `yaml:"aws_secret_access_key,omitempty" json:"aws_secret_access_key,omitempty"`
	AWSSessionToken    string `yaml:"aws_session_token,omitempty" json:"aws_session_token,omitempty"`
	AWSEndpointURI     string `yaml:"aws_endpoint_uri,omitempty" json:"aws_endpoint_uri,omitempty"`
	Bucket             string `yaml:"bucket" json:"bucket"`
	StartDate          string `yaml:"start_date" json:"start_date"`
}

// S3CSVTapSpec defines Tap configuration for S3 CSV. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/s3_csv.html)
type S3CSVTapSpec struct {
	Schemas             []S3CSVTapSchemaSpec   `yaml:"schemas" json:"schemas"`
	Connection          S3CSVTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows       *int                   `yaml:"batch_size_rows,omitempty" json:"batch_size_rows,omitempty"`
	StreamBufferSize    *int                   `yaml:"stream_buffer_size,omitempty" json:"stream_buffer_size,omitempty"`
	DefaultTargetSchema string                 `yaml:"default_target_schema,omitempty" json:"default_target_schema,omitempty"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *S3CSVTapSpec) ConnectorID() string {
	return string(S3CSVTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *S3CSVTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", S3CSVTapID, ts.Connection.Bucket))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *S3CSVTapSpec) Type() PipelinewiseTapType {
	return S3CSVTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *S3CSVTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *S3CSVTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// SnowflakeTapConnectionSpec defines Snowflake tap connection
type SnowflakeTapConnectionSpec struct {
	Account   string `yaml:"account" json:"account"`
	DBName    string `yaml:"dbname" json:"dbname"`
	User      string `yaml:"user" json:"user"`
	Password  string `yaml:"password" json:"password"`
	Warehouse string `yaml:"warehouse" json:"warehouse"`
}

// SnowflakeTapSpec defines Tap configuration for Snowflake. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/snowflake.html)
type SnowflakeTapSpec struct {
	Schemas    []TapSchemaSpec            `yaml:"schemas" json:"schemas"`
	Connection SnowflakeTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SnowflakeTapSpec) ConnectorID() string {
	return string(SnowflakeTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SnowflakeTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", SnowflakeTapID, ts.Connection.DBName))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *SnowflakeTapSpec) Type() PipelinewiseTapType {
	return SnowflakeTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *SnowflakeTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *SnowflakeTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// MongoDBTapConnectionSpec defines MongoDB Tap connection
type MongoDBTapConnectionSpec struct {
	Host           string `yaml:"host" json:"host"`
	Port           int    `yaml:"port" json:"port"`
	User           string `yaml:"user" json:"user"`
	Password       string `yaml:"password" json:"password"`
	AuthDatabase   string `yaml:"auth_database" json:"auth_database"`
	DBName         string `yaml:"dbname" json:"dbname"`
	ReplicaSet     string `yaml:"replica_set,omitempty" json:"replica_set,omitempty"`
	WriteBatchRows *int   `yaml:"write_batch_rows,omitempty" json:"write_batch_rows,omitempty"`
}

// MongoDBTapSpec defines Tap configuration for MongoDB. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/mongodb.html)
type MongoDBTapSpec struct {
	Schemas    []TapSchemaSpec          `yaml:"schemas" json:"schemas"`
	Connection MongoDBTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *MongoDBTapSpec) ConnectorID() string {
	return string(MongoDBTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *MongoDBTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", MongoDBTapID, ts.Connection.DBName))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *MongoDBTapSpec) Type() PipelinewiseTapType {
	return MongoDBTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *MongoDBTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *MongoDBTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// SalesforceTapConnectionSpec defines Salesforce Tap connection
type SalesforceTapConnectionSpec struct {
	ClientID     string `yaml:"client_id" json:"client_id"`
	ClientSecret string `yaml:"client_secret" json:"client_secret"`
	RefreshToken string `yaml:"refresh_token" json:"refresh_token"`
	StartDate    string `yaml:"start_date" json:"start_date"`
	APIType      string `yaml:"api_type" json:"api_type"`
}

// SalesforceTapSpec defines Tap configuration for Salesforce. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/salesforce.html)
type SalesforceTapSpec struct {
	Schemas    []TapSchemaSpec             `yaml:"schemas" json:"schemas"`
	Connection SalesforceTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SalesforceTapSpec) ConnectorID() string {
	return string(SalesforceTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SalesforceTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", SalesforceTapID, ts.Connection.ClientID))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *SalesforceTapSpec) Type() PipelinewiseTapType {
	return SalesforceTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *SalesforceTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *SalesforceTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// ZendeskTapConnectionSpec defines Zendesk Tap connection
type ZendeskTapConnectionSpec struct {
	AccessToken string `yaml:"access_token" json:"access_token"`
	Subdomain   string `yaml:"subdomain" json:"subdomain"`
	StartDate   string `yaml:"start_date" json:"start_date"`
	RateLimit   *int   `yaml:"rate_limit,omitempty" json:"rate_limit,omitempty"`
	MaxWorkers  *int   `yaml:"max_workers,omitempty" json:"max_workers,omitempty"`
	BatchSize   *int   `yaml:"batch_size,omitempty" json:"batch_size,omitempty"`
}

// ZendeskTapSpec defines Tap configuration for Zendesk. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/zendesk.html)
type ZendeskTapSpec struct {
	Schemas    []TapSchemaSpec          `yaml:"schemas" json:"schemas"`
	Connection ZendeskTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ZendeskTapSpec) ConnectorID() string {
	return string(ZendeskTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ZendeskTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", ZendeskTapID, ts.Connection.Subdomain))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *ZendeskTapSpec) Type() PipelinewiseTapType {
	return ZendeskTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *ZendeskTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *ZendeskTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// JiraTapConnectionSpec defines Jira Tap connection
type JiraTapConnectionSpec struct {
	BaseURL           string `yaml:"base_url" json:"base_url"`
	Username          string `yaml:"username,omitempty" json:"username,omitempty"`
	Password          string `yaml:"password,omitempty" json:"password,omitempty"`
	OauthClientSecret string `yaml:"oauth_client_secret,omitempty" json:"oauth_client_secret,omitempty"`
	OauthClientID     string `yaml:"oauth_client_id,omitempty" json:"oauth_client_id,omitempty"`
	AccessToken       string `yaml:"access_token,omitempty" json:"access_token,omitempty"`
	CloudID           string `yaml:"cloud_id,omitempty" json:"cloud_id,omitempty"`
	RefreshToken      string `yaml:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	StartData         string `yaml:"start_date" json:"start_date"`
}

// JiraTapSpec defines Tap configuration for Jira. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/jira.html)
type JiraTapSpec struct {
	Schemas    []TapSchemaSpec       `yaml:"schemas" json:"schemas"`
	Connection JiraTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *JiraTapSpec) ConnectorID() string {
	return string(JiraTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *JiraTapSpec) ID() PipelinewiseTapID {
	return JiraTapID
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *JiraTapSpec) Type() PipelinewiseTapType {
	return JiraTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *JiraTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *JiraTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// ZuoraTapConnectionSpec defines Zuora Tap connection
type ZuoraTapConnectionSpec struct {
	Username  string `yaml:"username,omitempty" json:"username,omitempty"`
	Password  string `yaml:"password,omitempty" json:"password,omitempty"`
	PartnerID string `yaml:"partner_id,omitempty" json:"partner_id,omitempty"`
	APIType   string `yaml:"api_type" json:"api_type"`
	Sandbox   bool   `yaml:"sandbox,omitempty" json:"sandbox,omitempty"`
	European  bool   `yaml:"european,omitempty" json:"european,omitempty"`
	StartData string `yaml:"start_date" json:"start_date"`
}

// ZuoraTapSpec defines Tap configuration for Zuora. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/zuora.html)
type ZuoraTapSpec struct {
	Schemas    []TapSchemaSpec        `yaml:"schemas" json:"schemas"`
	Connection ZuoraTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ZuoraTapSpec) ConnectorID() string {
	return string(ZuoraTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ZuoraTapSpec) ID() PipelinewiseTapID {
	if ts.Connection.PartnerID != "" {
		return PipelinewiseTapID(fmt.Sprintf("%s-%s", ZuoraTapID, ts.Connection.PartnerID))
	}

	return ZuoraTapID
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *ZuoraTapSpec) Type() PipelinewiseTapType {
	return ZuoraTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *ZuoraTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *ZuoraTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// GoogleAnalyticsOauthCredentials defines Google Analytics Oauth Credentials
type GoogleAnalyticsOauthCredentials struct {
	ClientID     string `yaml:"client_id" json:"client_id"`
	ClientSecret string `yaml:"client_secret" json:"client_secret"`
	AccessToken  string `yaml:"access_token" json:"access_token"`
	RefreshToken string `yaml:"refresh_token" json:"refresh_token"`
}

// GoogleAnalyticsTapConnectionSpec defines Google Analytics Tap connection
type GoogleAnalyticsTapConnectionSpec struct {
	ViewID           string                           `yaml:"view_id" json:"view_id"`
	OauthCredentials *GoogleAnalyticsOauthCredentials `yaml:"oauth_credentials,omitempty" json:"oauth_credentials,omitempty"`
	KeyFileLocation  string                           `yaml:"key_file_location,omitempty" json:"key_file_location,omitempty"`
	StartDate        string                           `yaml:"start_date" json:"start_date"`
}

// GoogleAnalyticsTapSpec defines Tap configuration for Google Analytics. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/google_analytics.html)
type GoogleAnalyticsTapSpec struct {
	Schemas    []TapSchemaSpec                  `yaml:"schemas" json:"schemas"`
	Connection GoogleAnalyticsTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *GoogleAnalyticsTapSpec) ConnectorID() string {
	return string(GoogleAnalyticsTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *GoogleAnalyticsTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", GoogleAnalyticsTapID, ts.Connection.ViewID))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *GoogleAnalyticsTapSpec) Type() PipelinewiseTapType {
	return GoogleAnalyticsTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *GoogleAnalyticsTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *GoogleAnalyticsTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// GithubTapConnectionSpec defines Github Tap connection
type GithubTapConnectionSpec struct {
	AccessToken string `yaml:"access_token" json:"access_token"`
	Repository  string `yaml:"repository" json:"repository"`
}

// GithubTapSpec defines Tap configuration for Github. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/github.html)
type GithubTapSpec struct {
	Schemas    []TapSchemaSpec         `yaml:"schemas" json:"schemas"`
	Connection GithubTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *GithubTapSpec) ConnectorID() string {
	return string(GithubTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *GithubTapSpec) ID() PipelinewiseTapID {
	return GithubTapID
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *GithubTapSpec) Type() PipelinewiseTapType {
	return GithubTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *GithubTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *GithubTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// ShopifyTapConnectionSpec defines Shopify Tap connection
type ShopifyTapConnectionSpec struct {
	Shop      string `yaml:"shop" json:"shop"`
	APIKey    string `yaml:"api_key" json:"api_key"`
	StartDate string `yaml:"start_date" json:"start_date"`
}

// ShopifyTapSpec defines Tap configuration for Shopify. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/shopify.html)
type ShopifyTapSpec struct {
	Schemas    []TapSchemaSpec          `yaml:"schemas" json:"schemas"`
	Connection ShopifyTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ShopifyTapSpec) ConnectorID() string {
	return string(ShopifyTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *ShopifyTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", ShopifyTapID, ts.Connection.Shop))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *ShopifyTapSpec) Type() PipelinewiseTapType {
	return ShopifyTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *ShopifyTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *ShopifyTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// SlackTapConnectionSpec defines Slack Tap connection
type SlackTapConnectionSpec struct {
	Token              string   `yaml:"token" json:"token"`
	StartDate          string   `yaml:"start_date" json:"start_date"`
	Channels           []string `yaml:"channels,omitempty" json:"channels,omitempty"`
	ExcludeArchived    string   `yaml:"exclude_archived,omitempty" json:"exclude_archived,omitempty"`
	PrivateChannels    string   `yaml:"private_channels,omitempty" json:"private_channels,omitempty"`
	JoinPublicChannels string   `yaml:"join_public_channels,omitempty" json:"join_public_channels,omitempty"`
	DateWindowSize     string   `yaml:"date_window_size,omitempty" json:"date_window_size,omitempty"`
	LookbackWindow     int      `yaml:"lookback_window,omitempty" json:"lookback_window,omitempty"`
}

// SlackTapSpec defines Tap configuration for Slack. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/slack.html)
type SlackTapSpec struct {
	Schemas    []TapSchemaSpec        `yaml:"schemas" json:"schemas"`
	Connection SlackTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SlackTapSpec) ConnectorID() string {
	return string(SlackTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *SlackTapSpec) ID() PipelinewiseTapID {
	return SlackTapID
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *SlackTapSpec) Type() PipelinewiseTapType {
	return SlackTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *SlackTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *SlackTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// MixpanelTapConnectionSpec defines Mixpanel Tap connection
type MixpanelTapConnectionSpec struct {
	APISecret         string   `yaml:"api_secret" json:"api_secret"`
	StartDate         string   `yaml:"start_date" json:"start_date"`
	DateWindowSize    int      `yaml:"date_window_size,omitempty" json:"date_window_size,omitempty"`
	AttributionWindow int      `yaml:"attribution_window,omitempty" json:"attribution_window,omitempty"`
	ProjectTimezone   string   `yaml:"project_timezone,omitempty" json:"project_timezone,omitempty"`
	UserAgent         string   `yaml:"user_agent,omitempty" json:"user_agent,omitempty"`
	DenestProperties  string   `yaml:"denest_properties,omitempty" json:"denest_properties,omitempty"`
	ExportEvents      []string `yaml:"export_events,omitempty" json:"export_events,omitempty"`
}

// MixpanelTapSpec defines Tap configuration for Mixpanel. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/mixpanel.html)
type MixpanelTapSpec struct {
	Schemas    []TapSchemaSpec           `yaml:"schemas" json:"schemas"`
	Connection MixpanelTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *MixpanelTapSpec) ConnectorID() string {
	return string(MixpanelTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *MixpanelTapSpec) ID() PipelinewiseTapID {
	return MixpanelTapID
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *MixpanelTapSpec) Type() PipelinewiseTapType {
	return MixpanelTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *MixpanelTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *MixpanelTapSpec) GetConnection() interface{} {
	return ts.Connection
}

// TwilioTapConnectionSpec defines Twilio Tap connection
type TwilioTapConnectionSpec struct {
	AccountSID string `yaml:"account_sid" json:"account_sid"`
	AuthToken  string `yaml:"auth_token" json:"auth_token"`
	StartDate  string `yaml:"start_date" json:"start_date"`
	UserAgent  string `yaml:"user_agent,omitempty" json:"user_agent,omitempty"`
}

// TwilioTapSpec defines Tap configuration for Twilio. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/twilio.html)
type TwilioTapSpec struct {
	Schemas    []TapSchemaSpec         `yaml:"schemas" json:"schemas"`
	Connection TwilioTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
}

// ConnectorID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *TwilioTapSpec) ConnectorID() string {
	return string(TwilioTapID)
}

// ID implement TapInfo interface to return Pipelinewise Tap ID
func (ts *TwilioTapSpec) ID() PipelinewiseTapID {
	return PipelinewiseTapID(fmt.Sprintf("%v-%v", TwilioTapID, ts.Connection.AccountSID))
}

// Type implement TapInfo interface to return Pipelinewise Tap Type
func (ts *TwilioTapSpec) Type() PipelinewiseTapType {
	return TwiliolTapType
}

// GetSchemas implement TapInfo interface to return schemas object
func (ts *TwilioTapSpec) GetSchemas() interface{} {
	return ts.Schemas
}

// GetConnection implement TapInfo interface to return connection object
func (ts *TwilioTapSpec) GetConnection() interface{} {
	return ts.Connection
}

func getTapInfo(pwJob *PipelinewiseJob) TapInfo {
	pwVal := reflect.ValueOf(pwJob.Spec.Tap)
	for fieldNth := 0; fieldNth < pwVal.NumField(); fieldNth++ {
		field := pwVal.Field(fieldNth)
		if !field.IsNil() {
			return field.Interface().(TapInfo)
		}
	}
	return nil
}

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pwJob *PipelinewiseJob) ([]byte, error) {
	tapInfo := getTapInfo(pwJob)
	targetID := GetTargetID(pwJob)

	if tapInfo != nil {
		return constructTap(tapInfo.ID(), tapInfo.Type(), targetID, tapInfo.GetConnection(), tapInfo.GetSchemas())
	}

	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

func constructTap(tapID PipelinewiseTapID, tapType PipelinewiseTapType, targetID PipelinewiseTargetID, dbConn interface{}, schemas interface{}) ([]byte, error) {
	tapConfiguration := GenericTapSpec{
		DatabaseConnection: dbConn,
		ID:                 tapID,
		Name:               string(tapID),
		Type:               tapType,
		Target:             targetID,
		Schemas:            schemas,
	}
	return yaml.Marshal(tapConfiguration)
}

// GetTapID calculate pipelinewise tap id
func GetTapID(pwJob *PipelinewiseJob) PipelinewiseTapID {
	tapInfo := getTapInfo(pwJob)
	if tapInfo != nil {
		return tapInfo.ID()
	}
	return ""
}

// GetTapConnectorID calculate pipelinewise connector id
func GetTapConnectorID(pwJob *PipelinewiseJob) string {
	tapInfo := getTapInfo(pwJob)
	if tapInfo != nil {
		return tapInfo.ConnectorID()
	}
	return ""
}
