package v1alpha1

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
