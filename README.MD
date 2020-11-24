# Kubernetes Pipelinewise Operator

A kubernetes operator to run Pipelinewise natively

## Installation

### Install CRDs

```bash
kustomize build config/crd | kubectl apply -f -
```

### Install Controller

To install the controller to your cluster, simply execute

```bash
kustomize build config/default | kubectl apply -f -
```

## Usage

To run pipelinewise job create crd for pipelinewisejob

```yaml
apiVersion: batch.pipelinewise/v1alpha1
kind: PipelinewiseJob
metadata:
  name: mariadb-postgresql-etl
spec:
  tap:
    mysql:
      db_conn:
        host: mariadb
        port: 3306
        user: root
        password: rKAVX4rQrQ
        dbname: test
      owner: dira.thea@gmail.com
      batch_size_rows: 20000
      stream_buffer_size: 0
      schemas:
        - source_schema: test
          target_schema: etl
          tables:
            - table_name: MOCK_DATA
              replication_method: FULL_TABLE
  target:
    postgresql:
      host: postgresql
      port: 5432
      user: application-target
      password: application-secret-password
      dbname: destination
```

## Roadmap

The following table are list of supported Pipelinewise taps and targets

| Type      | Name       | Status |
|-----------|------------|-------|
| Tap       | **[Postgres](https://github.com/transferwise/pipelinewise-tap-postgres)** | ❌ |
| Tap       | **[MySQL](https://github.com/transferwise/pipelinewise-tap-mysql)** | ✔ |
| Tap       | **[Kafka](https://github.com/transferwise/pipelinewise-tap-kafka)** | ❌ |
| Tap       | **[S3 CSV](https://github.com/transferwise/pipelinewise-tap-s3-csv)** | ❌ |
| Tap       | **[Zendesk](https://github.com/singer-io/tap-zendesk)** | ❌ |
| Tap       | **[Snowflake](https://github.com/transferwise/pipelinewise-tap-snowflake)** | ❌ |
| Tap       | **[Salesforce](https://github.com/singer-io/tap-salesforce)** | ❌ |
| Tap       | **[Jira](https://github.com/singer-io/tap-jira)** | ❌ |
| Tap       | **[MongoDB](https://github.com/transferwise/pipelinewise-tap-mongodb)** | ❌ |
| Tap       | **[AdWords](https://github.com/singer-io/tap-adwords)** | ❌ |
| Tap       | **[Google Analytics](https://github.com/transferwise/pipelinewise-tap-google-analytics)** | ❌ |
| Tap       | **[Oracle](https://github.com/transferwise/pipelinewise-tap-oracle)** | ❌ |
| Tap       | **[Zuora](https://github.com/transferwise/pipelinewise-tap-zuora)** | ❌ |
| Tap       | **[GitHub](https://github.com/singer-io/tap-github)** | ❌ |
| Tap       | **[Shopify](https://github.com/singer-io/tap-shopify)** | ❌ |
| Tap       | **[Slack](https://github.com/transferwise/pipelinewise-tap-slack)** | ❌ |
| Tap       | **[Mixpanel](https://github.com/transferwise/pipelinewise-tap-mixpanel)** | ❌ |
| Target    | **[Postgres](https://github.com/transferwise/pipelinewise-target-postgres)** | ✔ |
| Target    | **[Redshift](https://github.com/transferwise/pipelinewise-target-redshift)** | ✔ |
| Target    | **[Snowflake](https://github.com/transferwise/pipelinewise-target-snowflake)** | ✔ |
| Target    | **[S3 CSV](https://github.com/transferwise/pipelinewise-target-s3-csv)** | ✔ |

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache](LICENSE)