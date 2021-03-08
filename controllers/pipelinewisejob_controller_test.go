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

package controllers

import (
	"context"
	"fmt"
	"time"

	batchv1alpha1 "github.com/dirathea/pipelinewise-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	kbatchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("PipelinewiseJob Controller", func() {
	type TestCase struct {
		JobName              string
		Tap                  batchv1alpha1.TapSpec
		Target               batchv1alpha1.TargetSpec
		ConfigName           string
		ConfigAssertionValue string
	}
	const (
		jobNamespace = "default"
		timeout      = time.Second * 10
		duration     = time.Second * 10
		interval     = time.Millisecond * 250
		cron         = "0 0 * * *"
	)
	var (
		defaultTapSpec = batchv1alpha1.TapSpec{
			MySQL: &batchv1alpha1.MySQLTapSpec{
				Schemas: []batchv1alpha1.TapSchemaSpec{
					{
						Source: "default-source",
						Target: "default-target",
						Tables: []batchv1alpha1.TapTableSpec{
							{
								TableName: "default-table",
							},
						},
					},
				},
				Connection: batchv1alpha1.MySQLTapConnectionSpec{
					Host:   "default-host",
					DBName: "default-db-name",
				},
			},
		}
		defaultTargetSpec = batchv1alpha1.TargetSpec{
			PostgreSQL: &batchv1alpha1.PostgreSQLTargetSpec{
				Host:     "target-postgresql",
				Port:     5432,
				User:     "ordinary-user",
				Password: "ordinary-password",
			},
		}
	)
	Context("When creating PipelinewiseJob", func() {
		DescribeTable("Should create Kubernetes Resources",
			func(tc TestCase) {
				By("Submitting CRD")
				ctx := context.Background()
				pwJob := &batchv1alpha1.PipelinewiseJob{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "batch.pipelinewise.v1alpha1",
						Kind:       "PipelinewiseJob",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      tc.JobName,
						Namespace: jobNamespace,
					},
					Spec: batchv1alpha1.PipelinewiseJobSpec{
						Schedule: cron,
						Tap:      tc.Tap,
						Target:   tc.Target,
					},
				}
				Expect(k8sClient.Create(ctx, pwJob)).Should(Succeed())

				pwJobLookupKey := types.NamespacedName{Name: tc.JobName, Namespace: jobNamespace}
				createdPwJob := &batchv1alpha1.PipelinewiseJob{}

				// We'll need to retry getting this newly created CronJob, given that creation may not immediately happen.
				Eventually(func() bool {
					err := k8sClient.Get(ctx, pwJobLookupKey, createdPwJob)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(createdPwJob).ShouldNot(BeNil())

				By("Creating Pipelinewise configuration as ConfigMap")
				pwConfigLookupKey := types.NamespacedName{Name: tc.ConfigName, Namespace: jobNamespace}
				createdConfigMap := &corev1.ConfigMap{}

				Eventually(func() bool {
					err := k8sClient.Get(ctx, pwConfigLookupKey, createdConfigMap)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
				Expect(createdConfigMap.Data).Should(ContainElements(ContainSubstring(tc.ConfigAssertionValue)))

				By("Creating Cronjob")
				cronJobName := fmt.Sprintf("pw-job-%v", tc.JobName)
				pwCronJobLookupKey := types.NamespacedName{Name: cronJobName, Namespace: jobNamespace}
				createdCronJob := &kbatchv1beta1.CronJob{}

				Eventually(func() bool {
					err := k8sClient.Get(ctx, pwCronJobLookupKey, createdCronJob)
					if err != nil {
						return false
					}
					return true
				}, timeout, interval).Should(BeTrue())
			},
			Entry("Tap Mysql", TestCase{
				JobName: "tap-mysql",
				Tap: batchv1alpha1.TapSpec{
					MySQL: &batchv1alpha1.MySQLTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "Source",
								Target: "Target",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "test",
									},
								},
							},
						},
						Connection: batchv1alpha1.MySQLTapConnectionSpec{
							Host:     "tap-mysql",
							Port:     3306,
							User:     "tap-user",
							Password: "tap-password",
							DBName:   "test",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-mysql",
				ConfigAssertionValue: "tap-user",
			}),
			Entry("Tap PostgreSQL", TestCase{
				JobName: "tap-pgsql",
				Tap: batchv1alpha1.TapSpec{
					PostgreSQL: &batchv1alpha1.PostgreSQLTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "pg-source",
								Target: "pg-target",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "pg-source-table",
									},
								},
							},
						},
						Connection: batchv1alpha1.PostgreSQLTapConnectionSpec{
							Host:     "pg-source-host",
							Port:     5432,
							User:     "pg-source-user",
							Password: "pg-source-password",
							DBName:   "pg-source-name",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-pgsql",
				ConfigAssertionValue: "pg-source-table",
			}),
			Entry("Tap Oracle", TestCase{
				JobName: "tap-oracle",
				Tap: batchv1alpha1.TapSpec{
					Oracle: &batchv1alpha1.OracleTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "oracle-source",
								Target: "oracle-target",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "oracle-source-table",
									},
								},
							},
						},
						Connection: batchv1alpha1.OracleTapConnectionSpec{
							SID:      "oracle-source-sid",
							Host:     "oracle-source-host",
							Port:     12345,
							User:     "oracle-source-user",
							Password: "oracle-source-password",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-oracle",
				ConfigAssertionValue: "oracle-source-sid",
			}),
			Entry("Tap Kafka", TestCase{
				JobName: "tap-kafka",
				Tap: batchv1alpha1.TapSpec{
					Kafka: &batchv1alpha1.KafkaTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-kafka",
								Target: "target-kafka",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "source-table-kafka",
									},
								},
							},
						},
						Connection: batchv1alpha1.KafkaTapConnectionSpec{
							GroupID:          "group-kafka",
							BootstrapServers: "bootstrap-server",
							Topic:            "source-topic",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-kafka",
				ConfigAssertionValue: "kafka-source-topic",
			}),
			Entry("Tap S3CSV", TestCase{
				JobName: "tap-s3csv",
				Tap: batchv1alpha1.TapSpec{
					S3CSV: &batchv1alpha1.S3CSVTapSpec{
						Schemas: []batchv1alpha1.S3CSVTapSchemaSpec{
							{
								Source: "source-csv",
								Target: "target-csv",
								Tables: []batchv1alpha1.S3CSVTapTableSpec{
									{
										TableName: "source-table-csv",
										Mapping: batchv1alpha1.S3CSVTableMappingSpec{
											SearchPattern: "",
											Delimiter:     ",",
										},
									},
								},
							},
						},
						Connection: batchv1alpha1.S3CSVTapConnectionSpec{
							Bucket:    "source-bucket",
							StartDate: "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-s3csv",
				ConfigAssertionValue: "s3-csv-source-bucket",
			}),
			Entry("Tap Snowflake", TestCase{
				JobName: "tap-snowflake",
				Tap: batchv1alpha1.TapSpec{
					Snowflake: &batchv1alpha1.SnowflakeTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-snowflake",
								Target: "target-snowflake",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "source-table-snowflake",
									},
								},
							},
						},
						Connection: batchv1alpha1.SnowflakeTapConnectionSpec{
							Account:   "source-snowflake-account",
							DBName:    "source-snowflake-db",
							User:      "source-snowflake-user",
							Password:  "source-snowflake-password",
							Warehouse: "source-snowflake-warehouse",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-snowflake",
				ConfigAssertionValue: "snowflake-source-snowflake-db",
			}),
			Entry("Tap MongoDB", TestCase{
				JobName: "tap-mongodb",
				Tap: batchv1alpha1.TapSpec{
					MongoDB: &batchv1alpha1.MongoDBTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-mongodb",
								Target: "target-mongodb",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "source-table-mongodb",
									},
								},
							},
						},
						Connection: batchv1alpha1.MongoDBTapConnectionSpec{
							Host:         "source-host",
							Port:         27017,
							User:         "source-user-mongodb",
							Password:     "source-password-mongodb",
							AuthDatabase: "source-auth-db",
							DBName:       "source-db",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-mongodb",
				ConfigAssertionValue: "mongodb-source-db",
			}),
			Entry("Tap Salesforce", TestCase{
				JobName: "tap-salesforce",
				Tap: batchv1alpha1.TapSpec{
					Salesforce: &batchv1alpha1.SalesforceTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-salesforce",
								Target: "target-salesforce",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "source-table-salesforce",
									},
								},
							},
						},
						Connection: batchv1alpha1.SalesforceTapConnectionSpec{
							ClientID:     "sf-client-id",
							ClientSecret: "sf-client-secret",
							RefreshToken: "sf-refresh-token",
							StartDate:    "2021-01-01",
							APIType:      "BULK",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-salesforce",
				ConfigAssertionValue: "salesforce-sf-client-id",
			}),
			Entry("Tap Zendesk", TestCase{
				JobName: "tap-zendesk",
				Tap: batchv1alpha1.TapSpec{
					Zendesk: &batchv1alpha1.ZendeskTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-zendesk",
								Target: "target-zendesk",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "table-zendesk",
									},
								},
							},
						},
						Connection: batchv1alpha1.ZendeskTapConnectionSpec{
							AccessToken: "zendesk-access-token",
							Subdomain:   "zendesk-subdomain",
							StartDate:   "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-zendesk",
				ConfigAssertionValue: "zendesk-subdomain",
			}),
			Entry("Tap Jira", TestCase{
				JobName: "tap-jira",
				Tap: batchv1alpha1.TapSpec{
					Jira: &batchv1alpha1.JiraTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-jira",
								Target: "target-jira",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "table-jira",
									},
								},
							},
						},
						Connection: batchv1alpha1.JiraTapConnectionSpec{
							BaseURL:  "https://awesome-orgs.atlassian.com",
							Username: "awesome-jira-username",
							Password: "awesome-jira-password",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-jira",
				ConfigAssertionValue: "awesome-orgs.atlassian.com",
			}),
			Entry("Tap Zuora", TestCase{
				JobName: "tap-zuora",
				Tap: batchv1alpha1.TapSpec{
					Zuora: &batchv1alpha1.ZuoraTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-zuora",
								Target: "target-zuora",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "table-zuora",
									},
								},
							},
						},
						Connection: batchv1alpha1.ZuoraTapConnectionSpec{
							Username:  "zuora-user",
							Password:  "zuora-pass",
							PartnerID: "zuora-partner-id",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-zuora",
				ConfigAssertionValue: "zuora-zuora-partner-id",
			}),
			Entry("Tap Google Analytics", TestCase{
				JobName: "tap-google-analytics",
				Tap: batchv1alpha1.TapSpec{
					GoogleAnalytics: &batchv1alpha1.GoogleAnalyticsTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-google-analytics",
								Target: "target-google-analytics",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "table-google-analytics",
									},
								},
							},
						},
						Connection: batchv1alpha1.GoogleAnalyticsTapConnectionSpec{
							ViewID:          "home-dashboard",
							KeyFileLocation: "some-path-to-json",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-google-analytics",
				ConfigAssertionValue: "google-analytics-home-dashboard",
			}),
			Entry("Tap Github", TestCase{
				JobName: "tap-github",
				Tap: batchv1alpha1.TapSpec{
					Github: &batchv1alpha1.GithubTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-github",
								Target: "target-github",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "commits",
									},
									{
										TableName: "pull_requests",
									},
								},
							},
						},
						Connection: batchv1alpha1.GithubTapConnectionSpec{
							AccessToken: "awesome-token",
							Repository:  "https://github.com/dirathea/pipelinewise-operator",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-github",
				ConfigAssertionValue: "https://github.com/dirathea/pipelinewise-operator",
			}),
			Entry("Tap Shopify", TestCase{
				JobName: "tap-shopify",
				Tap: batchv1alpha1.TapSpec{
					Shopify: &batchv1alpha1.ShopifyTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-shopify",
								Target: "target-shopify",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "orders",
									},
								},
							},
						},
						Connection: batchv1alpha1.ShopifyTapConnectionSpec{
							Shop:      "awesome-shop",
							APIKey:    "awesome-api-key",
							StartDate: "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-shopify",
				ConfigAssertionValue: "shopify-awesome-shop",
			}),
			Entry("Tap Slack", TestCase{
				JobName: "tap-slack",
				Tap: batchv1alpha1.TapSpec{
					Slack: &batchv1alpha1.SlackTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-slack",
								Target: "target-slack",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "users",
									},
								},
							},
						},
						Connection: batchv1alpha1.SlackTapConnectionSpec{
							Token:     "awesome-slack-token",
							StartDate: "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-slack",
				ConfigAssertionValue: "awesome-slack-token",
			}),
			Entry("Tap Mixpanel", TestCase{
				JobName: "tap-mixpanel",
				Tap: batchv1alpha1.TapSpec{
					Mixpanel: &batchv1alpha1.MixpanelTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-mixpanel",
								Target: "target-mixpanel",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "funnels",
									},
								},
							},
						},
						Connection: batchv1alpha1.MixpanelTapConnectionSpec{
							APISecret: "awesome-api-secret",
							StartDate: "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-mixpanel",
				ConfigAssertionValue: "awesome-api-secret",
			}),
			Entry("Tap Twilio", TestCase{
				JobName: "tap-twilio",
				Tap: batchv1alpha1.TapSpec{
					Twilio: &batchv1alpha1.TwilioTapSpec{
						Schemas: []batchv1alpha1.TapSchemaSpec{
							{
								Source: "source-twilio",
								Target: "target-twilio",
								Tables: []batchv1alpha1.TapTableSpec{
									{
										TableName: "workspaces",
									},
									{
										TableName: "activities",
									},
								},
							},
						},
						Connection: batchv1alpha1.TwilioTapConnectionSpec{
							AccountSID: "twilio-awesome-sid",
							AuthToken:  "twilio-tokens",
							StartDate:  "2021-01-01",
						},
					},
				},
				Target:               defaultTargetSpec,
				ConfigName:           "pw-config-tap-twilio",
				ConfigAssertionValue: "twilio-twilio-awesome-sid",
			}),
			Entry("Target Postgres", TestCase{
				JobName: "target-postgres",
				Tap:     defaultTapSpec,
				Target: batchv1alpha1.TargetSpec{
					PostgreSQL: &batchv1alpha1.PostgreSQLTargetSpec{
						Host:         "postgres-host",
						Port:         5432,
						User:         "postgres-user",
						Password:     "postgres-pass",
						DatabaseName: "postgres-db",
					},
				},
				ConfigName:           "pw-config-target-postgres",
				ConfigAssertionValue: "postgres-postgres-db",
			}),
			Entry("Target Redshift", TestCase{
				JobName: "target-redshift",
				Tap:     defaultTapSpec,
				Target: batchv1alpha1.TargetSpec{
					Redshift: &batchv1alpha1.RedshiftTargetSpec{
						Host:         "redshift-host",
						Port:         5432,
						User:         "redshift-user",
						Password:     "redshift-password",
						DatabaseName: "redshift-db-name",
					},
				},
				ConfigName:           "pw-config-target-redshift",
				ConfigAssertionValue: "redshift-redshift-db-name",
			}),
			Entry("Target Snowflake", TestCase{
				JobName: "target-snowflake",
				Tap:     defaultTapSpec,
				Target: batchv1alpha1.TargetSpec{
					Snowflake: &batchv1alpha1.SnowflakeTargetSpec{
						Account:      "target-snowflake-account",
						DatabaseName: "target-snowflake-db-name",
						User:         "target-snowflake-user",
						Password:     "target-snowflake-password",
						Warehouse:    "target-snowflake-warehouse",
					},
				},
				ConfigName:           "pw-config-target-snowflake",
				ConfigAssertionValue: "snowflake-target-snowflake-db-name",
			}),
			Entry("Target S3CSV", TestCase{
				JobName: "target-s3csv",
				Tap:     defaultTapSpec,
				Target: batchv1alpha1.TargetSpec{
					S3CSV: &batchv1alpha1.S3CSVTargetSpec{
						S3Bucket: "bucket-name",
					},
				},
				ConfigName:           "pw-config-target-s3csv",
				ConfigAssertionValue: "s3-csv-bucket-name",
			}),
		)
	})
})
