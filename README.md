# prometheus-ecs-sd

Prometheus service discovery for Amazon ECS using `file_sd_config` format.

## Usage

```
./prometheus-ecs-sd -path /path/to/file_sd_configs/ecs_sd.yml
```

## Default labels

`cluster`, `service`, `task`, and `instance` will set as below:

```
- targets: [ecs.external.link.example.com:port]
  labels:
    cluster: ecs-example-cluster
    service: ecs-example-service
    task: xxxxxxxx-0000-0000-xxxxxxxxxxxxxxxxx
    instance: ecs-example-instance-i-xxxxxxxxxxxxxxxxx
```

## Custom metrics path

If container has `ECS_SD_LABEL_METRICS_PATH` environment variable like `/custom/metrics/path`, prometheus-ecs-sd will generate config as below:

```
- targets: [ecs.external.link.example.com:port]
  labels:
    __metrics_path__: ${ECS_SD_METRICS_PATH}
```