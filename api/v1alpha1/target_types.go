package v1alpha1

// PostgreSQLTargetSpec defines PostgreSQL Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/postgres.html)
type PostgreSQLTargetSpec struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"dbname"`
}

// RedshiftTargetSpec defines Redshift Target configuration. [Read more](https://transferwise.github.io/pipelinewise/connectors/targets/redshift.html)
type RedshiftTargetSpec struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	User                   string `json:"user"`
	Password               string `json:"password"`
	DatabaseName           string `json:"dbname"`
	AWSProfile             string `json:"aws_profile,omitempty"`
	AWSAccessKeyID         string `json:"aws_access_key_id,omitempty"`
	AWSAccessSecretKey     string `json:"aws_secret_access_key,omitempty"`
	AWSSessionToken        string `json:"aws_session_token,omitempty"`
	AWSRedshiftCopyRoleARN string `json:"aws_redshift_copy_role_arn,omitempty"`
	S3Bucket               string `json:"s3_bucket"`
	S3KeyPrefix            string `json:"s3_key_prefix,omitempty"`
	S3ACL                  string `json:"s3_acl,omitempty"`
	CopyOptions            string `json:"copy_options"`
}
