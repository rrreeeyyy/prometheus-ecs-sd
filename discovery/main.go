package discovery

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
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
	Region    string
	AccessKey string
	SecretKey string
	RoleARN   string
	Profile   string

	RefreshInterval model.Duration
	OnlyECSEnable   bool
}

type Discovery struct {
	aws      *aws.Config
	interval time.Duration
	profile  string
	roleARN  string
	logger   log.Logger
}

func NewDiscovery(conf SDConfig, logger log.Logger) (*Discovery, error) {
	if conf.Region == "" {
		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}
		metadata := ec2metadata.New(sess)
		region, err := metadata.Region()
		if err != nil {
			return nil, fmt.Errorf("ECS SD Configuration requires a region")
		}
		conf.Region = region
	}

	creds := credentials.NewStaticCredentials(conf.AccessKey, conf.SecretKey, "")
	if conf.AccessKey == "" && conf.SecretKey == "" {
		creds = nil
	}

	return &Discovery{
		aws: &aws.Config{
			Region: &config.Regions,
			Credentials: creds,
		}
		profile: conf.Profile,
		roleARN: conf.RoleARN,
		interval: time.Duration(conf.RefreshInterval),
		port: conf.Port,
	}
}

func (d *discovery) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	ticker := time.NewTicker(d.interval)
	defer ticker.Stop()

	for c := time.Tick(time.Duration(d.interval) * time.Second); ; {
		select {
		case <-c:
			continue
		case <-ctx.Done():
			return
		}
	}
}
