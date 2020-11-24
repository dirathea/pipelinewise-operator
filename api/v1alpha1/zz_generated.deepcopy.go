// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLTapConnectionSpec) DeepCopyInto(out *MySQLTapConnectionSpec) {
	*out = *in
	if in.SessionSQLs != nil {
		in, out := &in.SessionSQLs, &out.SessionSQLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLTapConnectionSpec.
func (in *MySQLTapConnectionSpec) DeepCopy() *MySQLTapConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLTapConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLTapSpec) DeepCopyInto(out *MySQLTapSpec) {
	*out = *in
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]TapSchemaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Connection.DeepCopyInto(&out.Connection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLTapSpec.
func (in *MySQLTapSpec) DeepCopy() *MySQLTapSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLTapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OracleTapSpec) DeepCopyInto(out *OracleTapSpec) {
	*out = *in
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]TapSchemaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Connection = in.Connection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OracleTapSpec.
func (in *OracleTapSpec) DeepCopy() *OracleTapSpec {
	if in == nil {
		return nil
	}
	out := new(OracleTapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OracleTargetConnectionSpec) DeepCopyInto(out *OracleTargetConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OracleTargetConnectionSpec.
func (in *OracleTargetConnectionSpec) DeepCopy() *OracleTargetConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(OracleTargetConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelinewiseJob) DeepCopyInto(out *PipelinewiseJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelinewiseJob.
func (in *PipelinewiseJob) DeepCopy() *PipelinewiseJob {
	if in == nil {
		return nil
	}
	out := new(PipelinewiseJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelinewiseJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelinewiseJobList) DeepCopyInto(out *PipelinewiseJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelinewiseJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelinewiseJobList.
func (in *PipelinewiseJobList) DeepCopy() *PipelinewiseJobList {
	if in == nil {
		return nil
	}
	out := new(PipelinewiseJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelinewiseJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelinewiseJobSpec) DeepCopyInto(out *PipelinewiseJobSpec) {
	*out = *in
	in.Tap.DeepCopyInto(&out.Tap)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelinewiseJobSpec.
func (in *PipelinewiseJobSpec) DeepCopy() *PipelinewiseJobSpec {
	if in == nil {
		return nil
	}
	out := new(PipelinewiseJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelinewiseJobStatus) DeepCopyInto(out *PipelinewiseJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelinewiseJobStatus.
func (in *PipelinewiseJobStatus) DeepCopy() *PipelinewiseJobStatus {
	if in == nil {
		return nil
	}
	out := new(PipelinewiseJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLTapSpec) DeepCopyInto(out *PostgreSQLTapSpec) {
	*out = *in
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = make([]TapSchemaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Connection = in.Connection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLTapSpec.
func (in *PostgreSQLTapSpec) DeepCopy() *PostgreSQLTapSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLTapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLTargetConnectionSpec) DeepCopyInto(out *PostgreSQLTargetConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLTargetConnectionSpec.
func (in *PostgreSQLTargetConnectionSpec) DeepCopy() *PostgreSQLTargetConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLTargetConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLTargetSpec) DeepCopyInto(out *PostgreSQLTargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLTargetSpec.
func (in *PostgreSQLTargetSpec) DeepCopy() *PostgreSQLTargetSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedshiftTargetSpec) DeepCopyInto(out *RedshiftTargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedshiftTargetSpec.
func (in *RedshiftTargetSpec) DeepCopy() *RedshiftTargetSpec {
	if in == nil {
		return nil
	}
	out := new(RedshiftTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CSVTargetSpec) DeepCopyInto(out *S3CSVTargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CSVTargetSpec.
func (in *S3CSVTargetSpec) DeepCopy() *S3CSVTargetSpec {
	if in == nil {
		return nil
	}
	out := new(S3CSVTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnowflakeTargetSpec) DeepCopyInto(out *SnowflakeTargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnowflakeTargetSpec.
func (in *SnowflakeTargetSpec) DeepCopy() *SnowflakeTargetSpec {
	if in == nil {
		return nil
	}
	out := new(SnowflakeTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TapSchemaSpec) DeepCopyInto(out *TapSchemaSpec) {
	*out = *in
	if in.Tables != nil {
		in, out := &in.Tables, &out.Tables
		*out = make([]TapTableSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TapSchemaSpec.
func (in *TapSchemaSpec) DeepCopy() *TapSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(TapSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TapSpec) DeepCopyInto(out *TapSpec) {
	*out = *in
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		*out = new(MySQLTapSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PostgreSQL != nil {
		in, out := &in.PostgreSQL, &out.PostgreSQL
		*out = new(PostgreSQLTapSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Oracle != nil {
		in, out := &in.Oracle, &out.Oracle
		*out = new(OracleTapSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TapSpec.
func (in *TapSpec) DeepCopy() *TapSpec {
	if in == nil {
		return nil
	}
	out := new(TapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TapTableSpec) DeepCopyInto(out *TapTableSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TapTableSpec.
func (in *TapTableSpec) DeepCopy() *TapTableSpec {
	if in == nil {
		return nil
	}
	out := new(TapTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
	if in.Redshift != nil {
		in, out := &in.Redshift, &out.Redshift
		*out = new(RedshiftTargetSpec)
		**out = **in
	}
	if in.PostgreSQL != nil {
		in, out := &in.PostgreSQL, &out.PostgreSQL
		*out = new(PostgreSQLTargetSpec)
		**out = **in
	}
	if in.Snowflake != nil {
		in, out := &in.Snowflake, &out.Snowflake
		*out = new(SnowflakeTargetSpec)
		**out = **in
	}
	if in.S3CSV != nil {
		in, out := &in.S3CSV, &out.S3CSV
		*out = new(S3CSVTargetSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}
