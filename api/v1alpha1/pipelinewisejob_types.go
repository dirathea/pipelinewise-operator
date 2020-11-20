/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLTapTableSpec defines MySQL Tap Table configuration
type MySQLTapTableSpec struct {
	TableName         string `yaml:"table_name" json:"table_name"`
	ReplicationMethod string `yaml:"replication_method" json:"replication_method"`
	ReplicationKey    string `yaml:"replication_key,omitempty" json:"replication_key,omitempty"`
}

// MySQLTapSchemaSpec defines MySQL Tap schema configuration
type MySQLTapSchemaSpec struct {
	Source string              `yaml:"source_schema" json:"source_schema"`
	Target string              `yaml:"target_schema" json:"target_schema"`
	Tables []MySQLTapTableSpec `yaml:"tables" json:"tables"`
}

// MySQLTapConnectionSpec defines MySQL Tap connection configuration
type MySQLTapConnectionSpec struct {
	Host            string   `yaml:"host" json:"host"`
	Port            int      `yaml:"port" json:"port"`
	User            string   `yaml:"user" json:"user"`
	Password        string   `yaml:"password" json:"password"`
	DatabaseName    string   `yaml:"dbname" json:"dbname"`
	FilterDatabases string   `yaml:"filter_dbs,omitempty" json:"filter_dbs,omitempty"`
	ExportBatchRows int      `yaml:"export_batch_rows,omitempty" json:"export_batch_rows,omitempty"`
	SessionSQLs     []string `yaml:"session_sqls,omitempty" json:"session_sqls,omitempty"`
}

// MySQLTapSpec defines Tap configuration for MySQL. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/mysql.html)
type MySQLTapSpec struct {
	Schemas          []MySQLTapSchemaSpec   `yaml:"schemas" json:"schemas"`
	Connection       MySQLTapConnectionSpec `yaml:"db_conn" json:"db_conn"`
	Owner            string                 `yaml:"owner" json:"owner"`
	BatchSizeRows    int                    `yaml:"batch_size_rows" json:"batch_size_rows"`
	StreamBufferSize int                    `yaml:"stream_buffer_size" json:"stream_buffer_size"`
}

// TapSpec defines Tap configuration
type TapSpec struct {
	MySQL *MySQLTapSpec `json:"mysql,omitempty"`
}

// TargetSpec defines Target configuration
type TargetSpec struct {
	Redshift   *RedshiftTargetSpec   `json:"redshift,omitempty"`
	PostgreSQL *PostgreSQLTargetSpec `json:"postgresql,omitempty"`
	Snowflake  *SnowflakeTargetSpec  `json:"snowflake,omitempty"`
	S3CSV      *S3CSVTargetSpec      `json:"s3_csv,omitempty"`
}

// PipelinewiseJobSpec defines the desired state of PipelinewiseJob
type PipelinewiseJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Tap    TapSpec    `json:"tap"`
	Target TargetSpec `json:"target"`
}

// PipelinewiseJobStatus defines the observed state of PipelinewiseJob
type PipelinewiseJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PipelinewiseJob is the Schema for the pipelinewisejobs API
type PipelinewiseJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelinewiseJobSpec   `json:"spec,omitempty"`
	Status PipelinewiseJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelinewiseJobList contains a list of PipelinewiseJob
type PipelinewiseJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PipelinewiseJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PipelinewiseJob{}, &PipelinewiseJobList{})
}
