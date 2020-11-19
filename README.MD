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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache](LICENSE)