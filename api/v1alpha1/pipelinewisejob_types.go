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
	TableName         string `json:"table_name"`
	ReplicationMethod string `json:"replication_method"`
	ReplicationKey    string `json:"replication_key,omitempty"`
}

// MySQLTapSchemaSpec defines MySQL Tap schema configuration
type MySQLTapSchemaSpec struct {
	Source string              `json:"source_schema"`
	Target string              `json:"target_schema"`
	Tables []MySQLTapTableSpec `json:"tables"`
}

// MySQLTapConnectionSpec defines MySQL Tap connection configuration
type MySQLTapConnectionSpec struct {
	Host            string   `json:"host"`
	Port            int      `json:"port"`
	User            string   `json:"user"`
	Password        string   `json:"password"`
	DatabaseName    string   `json:"dbname"`
	FilterDatabases string   `json:"filter_dbs,omitempty"`
	ExportBatchRows int      `json:"export_batch_rows,omitempty"`
	SessionSQLs     []string `json:"session_sqls,omitempty"`
}

// MySQLTapSpec defines Tap configuration for MySQL. [Read more](https://transferwise.github.io/pipelinewise/connectors/taps/mysql.html)
type MySQLTapSpec struct {
	Schemas          []MySQLTapSchemaSpec   `json:"schemas"`
	Connection       MySQLTapConnectionSpec `json:"db_conn"`
	Owner            string                 `json:"owner"`
	BatchSizeRows    int                    `json:"batch_size_rows"`
	StreamBufferSize int                    `json:"stream_buffer_size"`
}

// TapSpec defines Tap configuration
type TapSpec struct {
	MySQL *MySQLTapSpec `json:"mysql,omitempty"`
}

// TargetSpec defines Target configuration
type TargetSpec struct {
	Redshift   *RedshiftTargetSpec   `json:"redshift,omitempty"`
	PostgreSQL *PostgreSQLTargetSpec `json:"postgresql,omitempty"`
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
