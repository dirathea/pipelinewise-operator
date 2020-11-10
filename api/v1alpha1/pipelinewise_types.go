package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// PipelinewiseType defines configuration type. Could be a `tap` or a `target` with defined application type
type PipelinewiseType string

type PipelinewiseID string

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseType = "mysql"
	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseType = "tap-mysql"
	// RedshiftPipelinewiseID defines redshift target ID
	RedshiftPipelinewiseID PipelinewiseID = "redshift"
	// RedshiftTargetType defines Pipelinewise Target type
	RedshiftTargetType PipelinewiseType = "target-redshift"
)

// PipelinewiseGenericConfiguration defines generic pipelinewise configuration
type PipelinewiseGenericConfiguration struct {
	ID   string           `json:"id"`
	Name string           `json:"name"`
	Type PipelinewiseType `json:"type"`
}

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pipelinewiseJob *PipelinewiseJob) (string, error) {
	// Find tap
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		// MySQL Tap found
		type MySQLTapConfiguration struct {
			PipelinewiseGenericConfiguration
			*MySQLTapSpec
			Target string `json:"target"`
		}
		mysqlConfiguration := MySQLTapConfiguration{
			MySQLTapSpec: pipelinewiseJob.Spec.Tap.MySQL,
			PipelinewiseGenericConfiguration: PipelinewiseGenericConfiguration{
				ID:   GetTapID(pipelinewiseJob),
				Name: "Tap Mysql to Target with suffix",
				Type: MySQLTapType,
			},
			Target: GetTargetID(pipelinewiseJob),
		}
		configurationBytes, err := yaml.Marshal(mysqlConfiguration)
		if err != nil {
			return "", err
		}
		return string(configurationBytes), err
	}
	return "", fmt.Errorf("No Valid Tap configured")
}

// ConstructTargetConfiguration parse and return a target yaml configuration string
func ConstructTargetConfiguration(pipelinewiseJob *PipelinewiseJob) (string, error) {
	// Find target
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		// Redshift Target found
		type RedshiftTargetConfiguration struct {
			PipelinewiseGenericConfiguration
			*RedshiftTargetSpec
		}
		redshiftConfiguration := RedshiftTargetConfiguration{
			RedshiftTargetSpec: pipelinewiseJob.Spec.Target.Redshift,
			PipelinewiseGenericConfiguration: PipelinewiseGenericConfiguration{
				ID:   GetTargetID(pipelinewiseJob),
				Name: "Redshift Target with suffix",
				Type: MySQLTapType,
			},
		}
		configurationBytes, err := yaml.Marshal(redshiftConfiguration)
		if err != nil {
			return "", err
		}
		return string(configurationBytes), err
	}
	return "", fmt.Errorf("No Valid Tap configured")
}

// GetTargetID calculate pipelinewise target id
func GetTargetID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Target.Redshift != nil {
		return string(RedshiftPipelinewiseID)
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
