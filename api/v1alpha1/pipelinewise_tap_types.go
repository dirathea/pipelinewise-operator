package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseType = "mysql"
	// PostgreSQLTapID defines Pipelinewise Postgres Tap ID
	PostgreSQLTapID PipelinewiseType = "postgres"
	// OracleTapID defines Pipelinewise Oracle Tap ID
	OracleTapID PipelinewiseType = "oracle"
	// KafkaTapID defines Pipelinewise Kafka Tap ID
	KafkaTapID PipelinewiseType = "kafka"

	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseType = "tap-mysql"
	// PostgreSQLTapType defines Pipelinewise Postgres Tap type
	PostgreSQLTapType PipelinewiseType = "tap-postgres"
	// OracleTapType defines Pipelinewise Oracle Tap type
	OracleTapType PipelinewiseType = "tap-oracle"
	// KafkaTapType defines Pipelinewise Kafka Tap type
	KafkaTapType PipelinewiseType = "tap-kafka"
)

// GenericTapSpec defines generic Pipelinewise Tap configuration
// +kubebuilder:object:generate=false
type GenericTapSpec struct {
	ID    string           `yaml:"id"`
	Name  string           `yaml:"name"`
	Type  PipelinewiseType `yaml:"type"`
	Owner string           `yaml:"owner,omitempty" json:"owner,omitempty"`
	// DefaultTargetSchema only applies on Kafka
	DefaultTargetSchema string      `yaml:"default_target_schema,omitempty" json:"default_target_schema,omitempty"`
	DatabaseConnection  interface{} `yaml:"db_conn"`
	Target              string      `yaml:"target"`
	Schemas             interface{} `yaml:"schemas"`
}

// TapTableSpec defines MySQL Tap Table configuration
type TapTableSpec struct {
	TableName         string `yaml:"table_name" json:"table_name"`
	ReplicationMethod string `yaml:"replication_method" json:"replication_method"`
	ReplicationKey    string `yaml:"replication_key,omitempty" json:"replication_key,omitempty"`
}

// TapSchemaSpec defines MySQL Tap schema configuration
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

// PostgreSQLTargetConnectionSpec defines Postgres tap connection configuration
type PostgreSQLTargetConnectionSpec struct {
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
	Schemas          []TapSchemaSpec                `yaml:"schemas" json:"schemas"`
	Connection       PostgreSQLTargetConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows    int                            `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize int                            `yaml:"stream_buffer_size" json:"stream_buffer_size"`
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

// KafkaTapPrimaryKey defines Kafka tap connection primary key
type KafkaTapPrimaryKey struct {
	TransferID string `yaml:"transfer_id" json:"transfer_id"`
}

// KafkaTapConnectionSpec defines Kafka tap connection configuration
type KafkaTapConnectionSpec struct {
	GroupID                 string             `yaml:"group_id" json:"group_id"`
	BootstrapServers        string             `yaml:"bootstrap_servers" json:"bootstrap_servers"`
	Topic                   string             `yaml:"topic" json:"topic"`
	PrimaryKeys             KafkaTapPrimaryKey `yaml:"primary_keys,omitempty" json:"primary_keys,omitempty"`
	MaxRuntimeMs            int                `yaml:"max_runtime_ms,omitempty" json:"max_runtime_ms,omitempty"`
	ConsumerTimeoutMs       int                `yaml:"consumer_timeout_ms,omitempty" json:"consumer_timeout_ms,omitempty"`
	SessionTimeoutMs        int                `yaml:"session_timeout_ms,omitempty" json:"session_timeout_ms,omitempty"`
	HeartbeatIntervalMs     int                `yaml:"heartbeat_interval_ms,omitempty" json:"heartbeat_interval_ms,omitempty"`
	MaxPollIntervalMs       int                `yaml:"max_poll_interval_ms,omitempty" json:"max_poll_interval_ms,omitempty"`
	MaxPollRecords          int                `yaml:"max_poll_records,omitempty" json:"max_poll_records,omitempty"`
	CommitIntervalMs        int                `yaml:"commit_interval_ms,omitempty" json:"commit_interval_ms,omitempty"`
	LocalStoreDir           string             `yaml:"local_store_dir,omitempty" json:"local_store_dir,omitempty"`
	LocalStoreBatchSizeRows int                `yaml:"local_store_batch_size_rows,omitempty" json:"local_store_batch_size_rows,omitempty"`
}

// KafkaTapSpec defines Tap configuration for Kafka. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/kafka.html)
type KafkaTapSpec struct {
	Schemas             []TapSchemaSpec        `yaml:"schemas" json:"schemas"`
	Connection          KafkaTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	BatchSizeRows       int                    `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize    int                    `yaml:"stream_buffer_size" json:"stream_buffer_size"`
	DefaultTargetSchema string                 `yaml:"default_target_schema" json:"default_target_schema"`
}

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	// Find tap
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		return constructMysqlTap(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Tap.PostgreSQL != nil {
		return constructPostgreSQLTap(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Tap.Oracle != nil {
		return constructOracleTap(pipelinewiseJob)
	}
	if pipelinewiseJob.Spec.Tap.Kafka != nil {
		return constructKafkaTap(pipelinewiseJob)
	}
	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

func constructMysqlTap(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	tapConfiguration := GenericTapSpec{
		DatabaseConnection: pipelinewiseJob.Spec.Tap.MySQL.Connection,
		ID:                 GetTapID(pipelinewiseJob),
		Name:               string(MySQLTapType),
		Type:               MySQLTapType,
		Target:             GetTargetID(pipelinewiseJob),
		Schemas:            pipelinewiseJob.Spec.Tap.MySQL.Schemas,
	}
	return yaml.Marshal(tapConfiguration)
}

func constructPostgreSQLTap(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	tapConfiguration := GenericTapSpec{
		DatabaseConnection: pipelinewiseJob.Spec.Tap.PostgreSQL.Connection,
		ID:                 GetTapID(pipelinewiseJob),
		Name:               string(PostgreSQLTapType),
		Type:               PostgreSQLTapType,
		Target:             GetTargetID(pipelinewiseJob),
		Schemas:            pipelinewiseJob.Spec.Tap.PostgreSQL.Schemas,
	}
	return yaml.Marshal(tapConfiguration)
}

func constructOracleTap(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	tapConfiguration := GenericTapSpec{
		DatabaseConnection: pipelinewiseJob.Spec.Tap.Oracle.Connection,
		ID:                 GetTapID(pipelinewiseJob),
		Name:               string(OracleTapID),
		Type:               OracleTapType,
		Target:             GetTargetID(pipelinewiseJob),
		Schemas:            pipelinewiseJob.Spec.Tap.Oracle.Schemas,
	}
	return yaml.Marshal(tapConfiguration)
}

func constructKafkaTap(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	tapConfiguration := GenericTapSpec{
		DatabaseConnection:  pipelinewiseJob.Spec.Tap.Kafka.Connection,
		ID:                  GetTapID(pipelinewiseJob),
		Name:                string(KafkaTapID),
		Type:                KafkaTapType,
		Target:              GetTargetID(pipelinewiseJob),
		DefaultTargetSchema: pipelinewiseJob.Spec.Tap.Kafka.DefaultTargetSchema,
		Schemas:             pipelinewiseJob.Spec.Tap.Kafka.Schemas,
	}
	return yaml.Marshal(tapConfiguration)
}

// GetTapID calculate pipelinewise tap id
func GetTapID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		return fmt.Sprintf("%v-%v", MySQLTapID, pipelinewiseJob.Spec.Tap.MySQL.Connection.DBName)
	}
	if pipelinewiseJob.Spec.Tap.PostgreSQL != nil {
		return fmt.Sprintf("%v-%v", PostgreSQLTapID, pipelinewiseJob.Spec.Tap.PostgreSQL.Connection.DBName)
	}
	if pipelinewiseJob.Spec.Tap.Oracle != nil {
		return fmt.Sprintf("%v-%v", OracleTapID, pipelinewiseJob.Spec.Tap.Oracle.Connection.SID)
	}
	if pipelinewiseJob.Spec.Tap.Kafka != nil {
		return fmt.Sprintf("%v-%v", KafkaTapID, pipelinewiseJob.Spec.Tap.Kafka.Connection.Topic)
	}
	return ""
}
