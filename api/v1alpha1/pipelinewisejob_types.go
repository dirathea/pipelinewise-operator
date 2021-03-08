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

// PipelinewiseType defines configuration type. Could be a `tap` or a `target` with defined application type
type PipelinewiseType string

// PipelinewiseID defines configuration ID
type PipelinewiseID string

// TapSpec defines Tap configuration
type TapSpec struct {
	MySQL           *MySQLTapSpec           `json:"mysql,omitempty"`
	PostgreSQL      *PostgreSQLTapSpec      `json:"postgres,omitempty"`
	Oracle          *OracleTapSpec          `json:"oracle,omitempty"`
	Kafka           *KafkaTapSpec           `json:"kafka,omitempty"`
	S3CSV           *S3CSVTapSpec           `json:"s3_csv,omitempty"`
	Snowflake       *SnowflakeTapSpec       `json:"snowflake,omitempty"`
	MongoDB         *MongoDBTapSpec         `json:"mongodb,omitempty"`
	Salesforce      *SalesforceTapSpec      `json:"salesforce,omitempty"`
	Zendesk         *ZendeskTapSpec         `json:"zendesk,omitempty"`
	Jira            *JiraTapSpec            `json:"jira,omitempty"`
	Zuora           *ZuoraTapSpec           `json:"zuora,omitempty"`
	GoogleAnalytics *GoogleAnalyticsTapSpec `json:"google_analytics,omitempty"`
	Github          *GithubTapSpec          `json:"github,omitempty"`
	Shopify         *ShopifyTapSpec         `json:"shopify,omitempty"`
	Slack           *SlackTapSpec           `json:"slack,omitempty"`
	Mixpanel        *MixpanelTapSpec        `json:"mixpanel,omitempty"`
	Twilio          *TwilioTapSpec          `json:"twilio,omitempty"`
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

	// Image override executor image. If not supplied it will be calculated based on tap and target id
	Image *string `json:"image,omitempty"`

	// Schedule defines cron expression of the job
	Schedule string `json:"schedule"`

	// Suspend flags the job to suspend subsequent executions
	Suspend *bool `json:"suspend,omitempty"`

	// SuccessfulJobsHistoryLimit define how many successful finished job to retain
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// FailedJobsHistoryLimit define how many failed finished job to retain
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`

	// All Pipelinewise job spec. Specify your simplified tap and target configuration
	Tap    TapSpec    `json:"tap"`
	Target TargetSpec `json:"target"`

	// Secret defines if the configuration uses [encrypted string](https://transferwise.github.io/pipelinewise/user_guide/encrypting_passwords.html)
	Secret *SecretSpec `json:"secret,omitempty"`
}

// SecretSpec defines secret specification for loading master password for [encrypted string](https://transferwise.github.io/pipelinewise/user_guide/encrypting_passwords.html)
type SecretSpec struct {
	Name string `json:"name"`
	Key  string `json:"key"`
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
