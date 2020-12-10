# Kubernetes Pipelinewise Operator

A kubernetes operator to run Pipelinewise

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

To run pipelinewise job create crd for pipelinewisejob. It is recommended to encrypt your sensitive values like passwords, tokens, and so on using [pipelinewise encrypt_string](https://transferwise.github.io/pipelinewise/user_guide/encrypting_passwords.html).

```yaml
apiVersion: batch.pipelinewise/v1alpha1
kind: PipelinewiseJob
metadata:
  name: pipelinewisejob-sample-mysql-to-postgres
spec:
  # Schedule defines the cron expression
  schedule: "*/2 * * * *"
  # Secret defines the master token secret to be use for decrypting pipelinewise string. You need to create kubernete secret first.
  secret:
    name: pw-master-token
    key: pipelinewise-master-password
  # tap defines the source of the data. Since pipelinewise using singer.io, this fields describe the singer tap
  tap:
    mysql:
      db_conn:
        host: mariadb
        port: 3306
        user: root
        password: |
          !vault |
            $ANSIBLE_VAULT;1.1;AES256
            63326538333564633932343139396262646161636331643930393634393564656663303234623630
            6464613364306637393230343266653232353431653035300a396236346536653366336633323961
            66303766363534666665343065376631316437636164643639316132326538353264623733616233
            3732613037363039370a383838383438653366656462623230633530393331326333373937313566
            3735
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
  # target defines the persistent target. Full example is on https://transferwise.github.io/pipelinewise/connectors/targets.html
  target:
    postgresql:
      host: postgresql
      port: 5432
      user: application-target
      password: |
        !vault |
          $ANSIBLE_VAULT;1.1;AES256
          34363466633039386362636561323561366434373065633639653638346438383566383035643966
          6134643930623133616235363339326131346439363431640a636363623737386661326435383435
          39643833326532393335363831663666333036326333363633393166613737333333663239363034
          3665666363306632320a373233393564343737373430623233643737373732633938343263306166
          35623234363937396536626635343931636263346364366461646164653630336163
      dbname: destination
```

## Roadmap

The following table are list of supported Pipelinewise taps and targets

| Type      | Name       | Status |
|-----------|------------|-------|
| Tap       | **[Postgres](https://github.com/transferwise/pipelinewise-tap-postgres)** | ✔ |
| Tap       | **[MySQL](https://github.com/transferwise/pipelinewise-tap-mysql)** | ✔ |
| Tap       | **[Kafka](https://github.com/transferwise/pipelinewise-tap-kafka)** | ✔ |
| Tap       | **[S3 CSV](https://github.com/transferwise/pipelinewise-tap-s3-csv)** | ✔ |
| Tap       | **[Zendesk](https://github.com/singer-io/tap-zendesk)** | ❌ |
| Tap       | **[Snowflake](https://github.com/transferwise/pipelinewise-tap-snowflake)** | ❌ |
| Tap       | **[Salesforce](https://github.com/singer-io/tap-salesforce)** | ❌ |
| Tap       | **[Jira](https://github.com/singer-io/tap-jira)** | ❌ |
| Tap       | **[MongoDB](https://github.com/transferwise/pipelinewise-tap-mongodb)** | ❌ |
| Tap       | **[AdWords](https://github.com/singer-io/tap-adwords)** | ❌ |
| Tap       | **[Google Analytics](https://github.com/transferwise/pipelinewise-tap-google-analytics)** | ❌ |
| Tap       | **[Oracle](https://github.com/transferwise/pipelinewise-tap-oracle)** | ✔ |
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