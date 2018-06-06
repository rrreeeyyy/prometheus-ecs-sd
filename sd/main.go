package sd

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

const (
	ecsClusterLabel        = model.MetaLabelPrefix + "ecs_cluster"
	ecsServiceLabel        = model.MetaLabelPrefix + "ecs_service"
	ecsTaskLabel           = model.MetaLabelPrefix + "ecs_task"
	ecsTaskDefinitionLabel = model.MetaLabelPrefix + "ecs_task_definition"
	ecsInstanceLabel       = model.MetaLabelPrefix + "ecs_instance"
)

type SDConfig struct {
	RefreshInterval int
	OnlyECSEnable   bool
}

type discovery struct {
	ecs             *ecs.ECS
	ec2             *ec2.EC2
	refreshInterval int
}

func newDiscovery(conf sdConfig) (*discovery, error) {
	ecssd := &discovery{}
}

func (d *discovery) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	for c := time.Tick(time.Duration(d.refreshInterval) * time.Second); ; {
		select {
		case <-c:
			continue
		case <-ctx.Done():
			return
		}
	}
}
