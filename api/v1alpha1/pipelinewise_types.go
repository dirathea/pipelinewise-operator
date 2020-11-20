package v1alpha1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// PipelinewiseType defines configuration type. Could be a `tap` or a `target` with defined application type
type PipelinewiseType string

// PipelinewiseID defines configuration ID
type PipelinewiseID string

const (
	// MySQLTapID defines Pipelinewise Mysql Tap ID
	MySQLTapID PipelinewiseType = "mysql"
	// MySQLTapType defines Pipelinewise Mysql Tap type
	MySQLTapType PipelinewiseType = "tap-mysql"
)

// ConstructTapConfiguration parse and return a tap yaml configuration string
func ConstructTapConfiguration(pipelinewiseJob *PipelinewiseJob) ([]byte, error) {
	// Find tap
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		// MySQL Tap found
		type MySQLTapConfiguration struct {
			ID                 string                 `yaml:"id"`
			Name               string                 `yaml:"name"`
			Type               PipelinewiseType       `yaml:"type"`
			DatabaseConnection MySQLTapConnectionSpec `yaml:"db_conn"`
			Target             string                 `yaml:"target"`
			Schemas            []MySQLTapSchemaSpec   `yaml:"schemas"`
		}
		mysqlConfiguration := MySQLTapConfiguration{
			DatabaseConnection: pipelinewiseJob.Spec.Tap.MySQL.Connection,
			ID:                 GetTapID(pipelinewiseJob),
			Name:               "Tap Mysql to Target with suffix",
			Type:               MySQLTapType,
			Target:             GetTargetID(pipelinewiseJob),
			Schemas:            pipelinewiseJob.Spec.Tap.MySQL.Schemas,
		}
		return yaml.Marshal(mysqlConfiguration)
	}
	return []byte{}, fmt.Errorf("No Valid Tap configured")
}

// GetTapID calculate pipelinewise tap id
func GetTapID(pipelinewiseJob *PipelinewiseJob) string {
	if pipelinewiseJob.Spec.Tap.MySQL != nil {
		return fmt.Sprintf("%v-%v", MySQLTapID, pipelinewiseJob.Spec.Tap.MySQL.Connection.DatabaseName)
	}
	return ""
}
