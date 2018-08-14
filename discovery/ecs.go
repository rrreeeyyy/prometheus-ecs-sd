package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"

	"github.com/rrreeeyyy/prometheus-ecs-sd/log"
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
			Region:      &conf.Region,
			Credentials: creds,
		},
		profile:  conf.Profile,
		roleARN:  conf.RoleARN,
		interval: time.Duration(conf.RefreshInterval),
	}
}

func (d *Discovery) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
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

func (d *Discovery) refresh() (tg *targetgroup.Group, err error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  *d.aws,
		Profile: d.profile,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create aws session: %s", err)
	}

	if d.roleARN != "" {
		creds := stscreds.NewCredentials(sess, d.roleARN)
	}
	tg = &targetgroup.Group{}

	return tg, nil
}
