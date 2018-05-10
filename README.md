# prometheus-ecs-sd

Prometheus service discovery for Amazon ECS using `file_sd_config` format.

## Install

TODO

## Usage

```
./prometheus-ecs-sd -path /path/to/file_sd_configs/ecs_sd.yml -cluster example-ecs-cluster
```

## Service discovery

If there is an `ECS_SD_ENABLE=1` environment variable in one of the containers in the task definition, it will be discovered as a target.
And if the `ECS_SD_CONTAINER_PORT` environment variable is in one of the containers in the task definition, it discovers container that has the specified port as Container Port.

For example, if service has task definitions like bellow:

```
$ aws ecs describe-task-definition --task-definition example-ecs-task-definition:1
{
    "taskDefinition": {
        "taskDefinitionArn": "arn:aws:ecs:ap-northeast-1:xxxxxxxxxxxx:task-definition/example-ecs-task-definition:1",
        "containerDefinitions": [
            {
                "name": "front",
                "image": "nginx:latest",
                "cpu": 64,
                "memory": 32,
                "links": [
                    "app:backend"
                ],
                "portMappings": [
                    {
                        "containerPort": 80,
                        "hostPort": 0,
                        "protocol": "tcp"
                    }
                ],
                :
            }
            {
                "name": "app",
                "image": "rrreeeyyy/example-docker-app:latest",
                "cpu": 32,
                "memory": 64,
                "links": [],
                "portMappings": [],
                "essential": true,
                "environment": [
                    {
                        "name": "ECS_SD_ENABLE",
                        "value": "1"
                    },
                    {
                        "name": "ECS_SD_CONTAINER_PORT",
                        "value": "80"
                    },
                ],
                :
            },
        ],
        :
```

And if tasks created from this task definition is running on the container instance with IP address `192.0.2.1` and binded `32768`, the following yaml is output.

```
- targets: [192.0.2.1:32768]
  labels:
    cluster: ecs-example-cluster
    service: ecs-example-service
    task: xxxxxxxx-0000-0000-xxxxxxxxxxxxxxxxx
    task_definitison: example-ecs-task-definition:1
    instance: ecs-example-instance-i-xxxxxxxxxxxxxxxxx
```

## Custom metrics path

If container has `ECS_SD_LABEL_METRICS_PATH` environment variable like `/custom/metrics/path`, prometheus-ecs-sd will generate config as below:

```
- targets: [ecs.external.link.example.com:port]
  labels:
    __metrics_path__: ${ECS_SD_METRICS_PATH}
```

## Author

Ryota Yoshikawa <yoshikawa@rrreeeyyy.com>