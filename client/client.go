package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/service"
)

const (
	INSTANCE_STATUS_PENDING    = "pending"
	INSTANCE_STATUS_RUNNING    = "running"
	INSTANCE_STATUS_STOPPED    = "stopped"
	INSTANCE_STATUS_SUSPENDED  = "suspended"
	INSTANCE_STATUS_TERMINATED = "terminated"
	INSTANCE_STATUS_CEASED     = "ceased"

	defaultOpTimeout = 180*time.Second
)

type QingCloudClient interface {
	RunInstance(arg *service.RunInstancesInput) (*service.Instance, error)
	DescribeInstance(instanceID string) (*service.Instance, error)
	StartInstance(instanceID string) error
	StopInstance(instanceID string, force bool) error
	RestartInstance(instanceID string) error
	TerminateInstance(instanceID string) error
	WaitInstanceStatus(instanceID string, status string) (*service.Instance, error)
}

func NewClient(config *config.Config, zone string) (QingCloudClient, error) {
	qcService, err := service.Init(config)
	if err != nil {
		return nil, err
	}
	instanceService, err := qcService.Instance(zone)
	if err != nil {
		return nil, err
	}
	jobService, err := qcService.Job(zone)
	if err != nil {
		return nil, err
	}

	c := &client{
		InstanceService:      instanceService,
		JobService:           jobService,
		OperationTimeout:            defaultOpTimeout,
		zone:                 zone,
	}
	return c, nil
}

type client struct {
	InstanceService  *service.InstanceService
	JobService       *service.JobService
	OperationTimeout time.Duration
	WaitInterval   time.Duration
	zone             string
}

func (c *client) RunInstance(input *service.RunInstancesInput) (*service.Instance, error) {

	output, err := c.InstanceService.RunInstances(input)
	if err != nil {
		return nil, err
	}
	if len(output.Instances) == 0 {
		return nil, errors.New("Create instance response error.")
	}
	jobID := output.JobID
	jobErr := c.waitJob(*jobID)
	if jobErr != nil {
		return nil, jobErr
	}
	instanceID := *output.Instances[0]
	_, waitErr := c.WaitInstanceStatus(instanceID, INSTANCE_STATUS_RUNNING)
	if waitErr != nil {
		return nil, waitErr
	}
	ins, waitErr := c.waitInstanceNetwork(instanceID)
	if waitErr != nil {
		return nil, waitErr
	}
	return ins, nil
}

func (c *client) DescribeInstance(instanceID string) (*service.Instance, error) {
	input := &service.DescribeInstancesInput{Instances: []*string{&instanceID}}
	output, err := c.InstanceService.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	if len(output.InstanceSet) == 0 {
		return nil, fmt.Errorf("Instance with id [%s] not exist.", instanceID)
	}
	return output.InstanceSet[0], nil
}

func (c *client) StartInstance(instanceID string) error {
	input := &service.StartInstancesInput{Instances: []*string{&instanceID}}
	output, err := c.InstanceService.StartInstances(input)
	if err != nil {
		return err
	}
	jobID := output.JobID
	waitErr := c.waitJob(*jobID)
	if waitErr != nil {
		return waitErr
	}
	_, err = c.WaitInstanceStatus(instanceID, INSTANCE_STATUS_RUNNING)
	return err
}

func (c *client) StopInstance(instanceID string, force bool) error {
	var forceParam int
	if force {
		forceParam = 1
	} else {
		forceParam = 0
	}
	input := &service.StopInstancesInput{Instances: []*string{&instanceID}, Force: &forceParam}
	output, err := c.InstanceService.StopInstances(input)
	if err != nil {
		return err
	}
	jobID := output.JobID
	waitErr := c.waitJob(*jobID)
	if waitErr != nil {
		return waitErr
	}
	_, err = c.WaitInstanceStatus(instanceID, INSTANCE_STATUS_STOPPED)
	return err
}

func (c *client) RestartInstance(instanceID string) error {
	input := &service.RestartInstancesInput{Instances: []*string{&instanceID}}
	output, err := c.InstanceService.RestartInstances(input)
	if err != nil {
		return err
	}
	jobID := output.JobID
	waitErr := c.waitJob(*jobID)
	if waitErr != nil {
		return waitErr
	}
	_, err = c.WaitInstanceStatus(instanceID, INSTANCE_STATUS_RUNNING)
	return err
}

func (c *client) TerminateInstance(instanceID string) error {
	input := &service.TerminateInstancesInput{Instances: []*string{&instanceID}}
	output, err := c.InstanceService.TerminateInstances(input)
	if err != nil {
		return err
	}
	jobID := output.JobID
	waitErr := c.waitJob(*jobID)
	if waitErr != nil {
		return waitErr
	}
	_, err = c.WaitInstanceStatus(instanceID, INSTANCE_STATUS_TERMINATED)
	return err
}

func (c *client) waitJob(jobID string) error {
	return WaitJob(c.JobService, jobID, c.OperationTimeout, c.WaitInterval)
}

func (c *client) WaitInstanceStatus(instanceID string, status string) (*service.Instance, error) {
	return WaitInstanceStatus(c.InstanceService, instanceID, status, c.OperationTimeout, c.WaitInterval)
}

func (c *client) waitInstanceNetwork(instanceID string) (*service.Instance, error) {
	return WaitInstanceNetwork(c.InstanceService, instanceID, c.OperationTimeout, c.WaitInterval)
}
