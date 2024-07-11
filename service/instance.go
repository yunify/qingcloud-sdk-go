// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package service

import (
	"fmt"
	"time"

	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/request/data"
	"github.com/yunify/qingcloud-sdk-go/request/errors"
)

var _ fmt.State
var _ time.Time

type InstanceService struct {
	Config     *config.Config
	Properties *InstanceServiceProperties
}

type InstanceServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Instance(zone string) (*InstanceService, error) {
	properties := &InstanceServiceProperties{
		Zone: &zone,
	}

	return &InstanceService{Config: s.Config, Properties: properties}, nil
}

// Documentation URL: https://docs.qingcloud.com/api/instance/cease_instances.html
func (s *InstanceService) CeaseInstances(i *CeaseInstancesInput) (*CeaseInstancesOutput, error) {
	if i == nil {
		i = &CeaseInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CeaseInstances",
		RequestMethod: "GET",
	}

	x := &CeaseInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type CeaseInstancesInput struct {
	Instances []*string `json:"instances" name:"instances" location:"params"` // Required
}

func (v *CeaseInstancesInput) Validate() error {

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "CeaseInstancesInput",
		}
	}

	return nil
}

type CeaseInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/describe_instance_types.html
func (s *InstanceService) DescribeInstanceTypes(i *DescribeInstanceTypesInput) (*DescribeInstanceTypesOutput, error) {
	if i == nil {
		i = &DescribeInstanceTypesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeInstanceTypes",
		RequestMethod: "GET",
	}

	x := &DescribeInstanceTypesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeInstanceTypesInput struct {
	Baremetal *int `json:"baremetal" name:"baremetal" location:"params"`
	// 指定查询的云服务器类型
	InstanceTypes []*string `json:"instance_types" name:"instance_types" location:"params"`
}

func (v *DescribeInstanceTypesInput) Validate() error {

	return nil
}

type DescribeInstanceTypesOutput struct {
	Message         *string         `json:"message" name:"message" location:"elements"`
	Action          *string         `json:"action" name:"action" location:"elements"`
	InstanceTypeSet []*InstanceType `json:"instance_type_set" name:"instance_type_set" location:"elements"`
	RetCode         *int            `json:"ret_code" name:"ret_code" location:"elements"`
	TotalCount      *int            `json:"total_count" name:"total_count" location:"elements"`
}

type InstanceType struct {
	VCPUsCurrent     int    `json:"vcpus_current"`
	Status           string `json:"status"`
	Description      string `json:"description"`
	InstanceTypeName string `json:"instance_type_name"`
	InstanceTypeID   string `json:"instance_type_id"`
	ZoneID           string `json:"zone_id"`
	MemoryCurrent    int    `json:"memory_current"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/describe_instances.html
func (s *InstanceService) DescribeInstances(i *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	if i == nil {
		i = &DescribeInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeInstances",
		RequestMethod: "GET",
	}

	x := &DescribeInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeInstancesInput struct {
	ImageID []*string `json:"image_id" name:"image_id" location:"params"`
	// InstanceClass's available values: 0, 1
	InstanceClass *int      `json:"instance_class" name:"instance_class" location:"params"`
	InstanceType  []*string `json:"instance_type" name:"instance_type" location:"params"`
	Instances     []*string `json:"instances" name:"instances" location:"params"`
	IsClusterNode *int      `json:"is_cluster_node" name:"is_cluster_node" default:"0" location:"params"`
	Limit         *int      `json:"limit" name:"limit" default:"20" location:"params"`
	Offset        *int      `json:"offset" name:"offset" default:"0" location:"params"`
	Owner         *string   `json:"owner" name:"owner" location:"params"`
	ProjectID     *string   `json:"project_id" name:"project_id" location:"params"`
	SearchWord    *string   `json:"search_word" name:"search_word" location:"params"`
	Status        []*string `json:"status" name:"status" location:"params"`
	Tags          []*string `json:"tags" name:"tags" location:"params"`
	// Verbose's available values: 0, 1
	Verbose *int `json:"verbose" name:"verbose" location:"params"`

	// alarm status of the instance
	AlarmStatus *string `json:"alarm_status" name:"alarm_status" location:"params"`
	// filter by border
	Border    *string `json:"border" name:"border" location:"params"`
	ConsoleID *string `json:"console_id" name:"console_id" location:"params"`
	// pitrix or self.
	Controller           *string `json:"controller" name:"controller" location:"params"`
	CreateTime           *string `json:"create_time" name:"create_time" location:"params"`
	DedicatedHostGroupID *string `json:"dedicated_host_group_id" name:"dedicated_host_group_id" location:"params"`
	DedicatedHostID      *string `json:"dedicated_host_id" name:"dedicated_host_id" location:"params"`
	// directory_id
	DirectoryID *string `json:"directory_id" name:"directory_id" location:"params"`
	// for filter reserved resource
	ExcludeReserved *string `json:"exclude_reserved" name:"exclude_reserved" location:"params"`
	ExcludeUser     *string `json:"exclude_user" name:"exclude_user" location:"params"`
	// for filter directory resources
	ExcludedDir *string `json:"excluded_dir" name:"excluded_dir" location:"params"`
	// query instances by excluded_plg_id
	ExcludedPlgID *string `json:"excluded_plg_id" name:"excluded_plg_id" location:"params"`
	// filter by fence
	Fence *string `json:"fence" name:"fence" location:"params"`
	// The graphics protocol
	GraphicsProtocol *string `json:"graphics_protocol" name:"graphics_protocol" location:"params"`
	// the host machine of instances.
	HostMachine *string `json:"host_machine" name:"host_machine" location:"params"`
	// filter by hypervisor.
	Hypervisor *string `json:"hypervisor" name:"hypervisor" location:"params"`
	// filter by instance group
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	// the name of the instance. Support partial match.
	InstanceName *string `json:"instance_name" name:"instance_name" location:"params"`
	IsElastic    *string `json:"is_elastic" name:"is_elastic" location:"params"`
	// memory_current
	MemoryCurrent *string `json:"memory_current" name:"memory_current" location:"params"`
	// output mount detail(1
	MountDetail *string `json:"mount_detail" name:"mount_detail" location:"params"`
	// query instances by mounted image
	MountImageID *string `json:"mount_image_id" name:"mount_image_id" location:"params"`
	// 0: include resources whose transition_status is not empty
	NotTransition *string `json:"not_transition" name:"not_transition" default:"0" location:"params"`
	// os_disk_size
	OSDiskSize *string `json:"os_disk_size" name:"os_disk_size" location:"params"`
	// the place group of instances
	PlaceGroupID *string `json:"place_group_id" name:"place_group_id" location:"params"`
	// Support windows
	Platform *string `json:"platform" name:"platform" location:"params"`
	// filter by replica policy.
	Repl *string `json:"repl" name:"repl" location:"params"`
	// for reverse sorting. 1: reverse
	Reverse *string `json:"reverse" name:"reverse" location:"params"`
	// if role==partner then it is for partner to get the instances of her/his Apps
	Role       *string `json:"role" name:"role" location:"params"`
	RootUserID *string `json:"root_user_id" name:"root_user_id" location:"params"`
	// sort key
	SortKey *string `json:"sort_key" name:"sort_key" location:"params"`
	// sriov_nic_type
	SriovNICType *string `json:"sriov_nic_type" name:"sriov_nic_type" location:"params"`
	// query instances support cdrom(1
	SupportCdrom *string `json:"support_cdrom" name:"support_cdrom" location:"params"`
	// filter by tags
	TransitionStatus *string `json:"transition_status" name:"transition_status" location:"params"`
	// vcpus_current
	VCPUsCurrent *string `json:"vcpus_current" name:"vcpus_current" location:"params"`
	VdcNodeID    *string `json:"vdc_node_id" name:"vdc_node_id" location:"params"`
	// filter by vxnet.
	VxNet *string `json:"vxnet" name:"vxnet" location:"params"`
	// filter by vxnet type.
	VxNetType       *string `json:"vxnet_type" name:"vxnet_type" location:"params"`
	WithoutContract *string `json:"without_contract" name:"without_contract" location:"params"`
	// zone id
	Zone *string `json:"zone" name:"zone" location:"params"`
}

func (v *DescribeInstancesInput) Validate() error {

	if v.InstanceClass != nil {
		instanceClassValidValues := []string{"0", "1"}
		instanceClassParameterValue := fmt.Sprint(*v.InstanceClass)

		instanceClassIsValid := false
		for _, value := range instanceClassValidValues {
			if value == instanceClassParameterValue {
				instanceClassIsValid = true
			}
		}

		if !instanceClassIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "InstanceClass",
				ParameterValue: instanceClassParameterValue,
				AllowedValues:  instanceClassValidValues,
			}
		}
	}

	if v.Verbose != nil {
		verboseValidValues := []string{"0", "1"}
		verboseParameterValue := fmt.Sprint(*v.Verbose)

		verboseIsValid := false
		for _, value := range verboseValidValues {
			if value == verboseParameterValue {
				verboseIsValid = true
			}
		}

		if !verboseIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "Verbose",
				ParameterValue: verboseParameterValue,
				AllowedValues:  verboseValidValues,
			}
		}
	}

	return nil
}

type DescribeInstancesOutput struct {
	Message     *string     `json:"message" name:"message" location:"elements"`
	Action      *string     `json:"action" name:"action" location:"elements"`
	InstanceSet []*Instance `json:"instance_set" name:"instance_set" location:"elements"`
	RetCode     *int        `json:"ret_code" name:"ret_code" location:"elements"`
	TotalCount  *int        `json:"total_count" name:"total_count" location:"elements"`
}

type Instance struct {
	VCPUsCurrent int      `json:"vcpus_current"`
	InstanceID   string   `json:"instance_id"`
	VolumeIDs    []string `json:"volume_ids"`
	Vxnets       []struct {
		VxnetName string `json:"vxnet_name"`
		VxnetType int    `json:"vxnet_type"`
		VxnetID   string `json:"vxnet_id"`
		NICID     string `json:"nic_id"`
		PrivateIP string `json:"private_ip"`
	} `json:"vxnets"`
	EIP struct {
		EIPID     string `json:"eip_id"`
		EIPAddr   string `json:"eip_addr"`
		Bandwidth string `json:"bandwidth"`
	} `json:"eip"`
	MemoryCurrent    int       `json:"memory_current"`
	SubCode          int       `json:"sub_code"`
	TransitionStatus string    `json:"transition_status"`
	InstanceName     string    `json:"instance_name"`
	InstanceType     string    `json:"instance_type"`
	CreateTime       time.Time `json:"create_time"`
	Status           string    `json:"status"`
	Description      string    `json:"description"`
	SecurityGroup    struct {
		IsDefault       int    `json:"is_default"`
		SecurityGroupID string `json:"security_group_id"`
	} `json:"security_group"`
	StatusTime time.Time `json:"status_time"`
	Image      struct {
		ProcessorType string `json:"processor_type"`
		Platform      string `json:"platform"`
		ImageSize     int    `json:"image_size"`
		ImageName     string `json:"image_name"`
		ImageID       string `json:"image_id"`
		OSFamily      string `json:"os_family"`
		Provider      string `json:"provider"`
	} `json:"image"`
	KeypairIDs []string `json:"keypair_ids"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/modify_instance_attributes.html
func (s *InstanceService) ModifyInstanceAttributes(i *ModifyInstanceAttributesInput) (*ModifyInstanceAttributesOutput, error) {
	if i == nil {
		i = &ModifyInstanceAttributesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ModifyInstanceAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyInstanceAttributesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type ModifyInstanceAttributesInput struct {
	Description  *string `json:"description" name:"description" location:"params"`
	Instance     *string `json:"instance" name:"instance" location:"params"` // Required
	InstanceName *string `json:"instance_name" name:"instance_name" location:"params"`
	NICMqueue    *string `json:"nic_mqueue" name:"nic_mqueue" location:"params"`
}

func (v *ModifyInstanceAttributesInput) Validate() error {

	if v.Instance == nil {
		return errors.ParameterRequiredError{
			ParameterName: "Instance",
			ParentName:    "ModifyInstanceAttributesInput",
		}
	}

	return nil
}

type ModifyInstanceAttributesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/reset_instances.html
func (s *InstanceService) ResetInstances(i *ResetInstancesInput) (*ResetInstancesOutput, error) {
	if i == nil {
		i = &ResetInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ResetInstances",
		RequestMethod: "GET",
	}

	x := &ResetInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type ResetInstancesInput struct {
	Instances    []*string `json:"instances" name:"instances" location:"params"` // Required
	LoginKeyPair *string   `json:"login_keypair" name:"login_keypair" location:"params"`
	// LoginMode's available values: keypair, passwd
	LoginMode   *string `json:"login_mode" name:"login_mode" location:"params"` // Required
	LoginPasswd *string `json:"login_passwd" name:"login_passwd" location:"params"`
	// NeedNewSID's available values: 0, 1
	NeedNewSID *int `json:"need_newsid" name:"need_newsid" default:"0" location:"params"`

	// force.
	Force *string `json:"force" name:"force" default:"0" location:"params"`
	// image_id.
	ImageID *string `json:"image_id" name:"image_id" location:"params"`
	// os_disk_size.
	OSDiskSize   *string `json:"os_disk_size" name:"os_disk_size" location:"params"`
	Type         *string `json:"type" name:"type" location:"params"`
	UserData     *string `json:"user_data" name:"user_data" location:"params"`
	Value        *string `json:"value" name:"value" location:"params"`
	VmDefinition *string `json:"vm_definition" name:"vm_definition" location:"params"`
	Zone         *string `json:"zone" name:"zone" location:"params"`
}

func (v *ResetInstancesInput) Validate() error {

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "ResetInstancesInput",
		}
	}

	if v.LoginMode == nil {
		return errors.ParameterRequiredError{
			ParameterName: "LoginMode",
			ParentName:    "ResetInstancesInput",
		}
	}

	if v.LoginMode != nil {
		loginModeValidValues := []string{"keypair", "passwd"}
		loginModeParameterValue := fmt.Sprint(*v.LoginMode)

		loginModeIsValid := false
		for _, value := range loginModeValidValues {
			if value == loginModeParameterValue {
				loginModeIsValid = true
			}
		}

		if !loginModeIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "LoginMode",
				ParameterValue: loginModeParameterValue,
				AllowedValues:  loginModeValidValues,
			}
		}
	}

	if v.NeedNewSID != nil {
		needNewSIDValidValues := []string{"0", "1"}
		needNewSIDParameterValue := fmt.Sprint(*v.NeedNewSID)

		needNewSIDIsValid := false
		for _, value := range needNewSIDValidValues {
			if value == needNewSIDParameterValue {
				needNewSIDIsValid = true
			}
		}

		if !needNewSIDIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "NeedNewSID",
				ParameterValue: needNewSIDParameterValue,
				AllowedValues:  needNewSIDValidValues,
			}
		}
	}

	return nil
}

type ResetInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/resize_instances.html
func (s *InstanceService) ResizeInstances(i *ResizeInstancesInput) (*ResizeInstancesOutput, error) {
	if i == nil {
		i = &ResizeInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ResizeInstances",
		RequestMethod: "GET",
	}

	x := &ResizeInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type ResizeInstancesInput struct {

	// CPU's available values: 1, 2, 4, 8, 16
	CPU          *int      `json:"cpu" name:"cpu" location:"params"`
	CPUModel     *string   `json:"cpu_model" name:"cpu_model" location:"params"`
	Gpu          *int      `json:"gpu" name:"gpu" location:"params"`
	InstanceType *string   `json:"instance_type" name:"instance_type" location:"params"`
	Instances    []*string `json:"instances" name:"instances" location:"params"` // Required
	// Memory's available values: 1024, 2048, 4096, 6144, 8192, 12288, 16384, 24576, 32768
	Memory     *int `json:"memory" name:"memory" location:"params"`
	OSDiskSize *int `json:"os_disk_size" name:"os_disk_size" location:"params"`

	// the boot device
	BootDev *string `json:"boot_dev" name:"boot_dev" location:"params"`
	// the CPU topology
	CPUTopology   *string `json:"cpu_topology" name:"cpu_topology" location:"params"`
	MemoryCurrent *string `json:"memory_current" name:"memory_current" location:"params"`
	MemoryMax     *string `json:"memory_max" name:"memory_max" location:"params"`
	VCPUsCurrent  *string `json:"vcpus_current" name:"vcpus_current" location:"params"`
	VCPUsMax      *string `json:"vcpus_max" name:"vcpus_max" location:"params"`
	VmDefinition  *string `json:"vm_definition" name:"vm_definition" location:"params"`
	Zone          *string `json:"zone" name:"zone" location:"params"`
}

func (v *ResizeInstancesInput) Validate() error {

	if v.CPU != nil {
		cpuValidValues := []string{"1", "2", "4", "8", "16"}
		cpuParameterValue := fmt.Sprint(*v.CPU)

		cpuIsValid := false
		for _, value := range cpuValidValues {
			if value == cpuParameterValue {
				cpuIsValid = true
			}
		}

		if !cpuIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "CPU",
				ParameterValue: cpuParameterValue,
				AllowedValues:  cpuValidValues,
			}
		}
	}

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "ResizeInstancesInput",
		}
	}

	if v.Memory != nil {
		memoryValidValues := []string{"1024", "2048", "4096", "6144", "8192", "12288", "16384", "24576", "32768"}
		memoryParameterValue := fmt.Sprint(*v.Memory)

		memoryIsValid := false
		for _, value := range memoryValidValues {
			if value == memoryParameterValue {
				memoryIsValid = true
			}
		}

		if !memoryIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "Memory",
				ParameterValue: memoryParameterValue,
				AllowedValues:  memoryValidValues,
			}
		}
	}

	return nil
}

type ResizeInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/restart_instances.html
func (s *InstanceService) RestartInstances(i *RestartInstancesInput) (*RestartInstancesOutput, error) {
	if i == nil {
		i = &RestartInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "RestartInstances",
		RequestMethod: "GET",
	}

	x := &RestartInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type RestartInstancesInput struct {
	Instances []*string `json:"instances" name:"instances" location:"params"` // Required
}

func (v *RestartInstancesInput) Validate() error {

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "RestartInstancesInput",
		}
	}

	return nil
}

type RestartInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/run_instances.html
func (s *InstanceService) RunInstances(i *RunInstancesInput) (*RunInstancesOutput, error) {
	if i == nil {
		i = &RunInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "RunInstances",
		RequestMethod: "GET",
	}

	x := &RunInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type RunInstancesInput struct {
	BillingID *string `json:"billing_id" name:"billing_id" location:"params"`
	Count     *int    `json:"count" name:"count" default:"1" location:"params"`
	// CPU's available values: 1, 2, 4, 8, 16
	CPU *int `json:"cpu" name:"cpu" default:"1" location:"params"`
	// CPUMax's available values: 1, 2, 4, 8, 16
	CPUMax *int `json:"cpu_max" name:"cpu_max" location:"params"`
	// CPUModel's available values: Westmere, SandyBridge, IvyBridge, Haswell, Broadwell
	CPUModel *string `json:"cpu_model" name:"cpu_model" default:"Westmere" location:"params"`
	Gpu      *int    `json:"gpu" name:"gpu" default:"0" location:"params"`
	Hostname *string `json:"hostname" name:"hostname" location:"params"`
	ImageID  *string `json:"image_id" name:"image_id" location:"params"` // Required
	// InstanceClass's available values: 0, 1, 2, 3, 4, 5, 6, 100, 101, 200, 201, 300, 301
	InstanceClass *int    `json:"instance_class" name:"instance_class" location:"params"`
	InstanceName  *string `json:"instance_name" name:"instance_name" location:"params"`
	InstanceType  *string `json:"instance_type" name:"instance_type" location:"params"`
	LoginKeyPair  *string `json:"login_keypair" name:"login_keypair" location:"params"`
	// LoginMode's available values: keypair, passwd
	LoginMode   *string `json:"login_mode" name:"login_mode" location:"params"` // Required
	LoginPasswd *string `json:"login_passwd" name:"login_passwd" location:"params"`
	// MemMax's available values: 1024, 2048, 4096, 6144, 8192, 12288, 16384, 24576, 32768
	MemMax *int `json:"mem_max" name:"mem_max" location:"params"`
	// Memory's available values: 1024, 2048, 4096, 6144, 8192, 12288, 16384, 24576, 32768
	Memory *int `json:"memory" name:"memory" default:"1024" location:"params"`
	// NeedNewSID's available values: 0, 1
	NeedNewSID *int `json:"need_newsid" name:"need_newsid" default:"0" location:"params"`
	// NeedUserdata's available values: 0, 1
	NeedUserdata  *int    `json:"need_userdata" name:"need_userdata" default:"0" location:"params"`
	OSDiskSize    *int    `json:"os_disk_size" name:"os_disk_size" location:"params"`
	SecurityGroup *string `json:"security_group" name:"security_group" location:"params"`
	UIType        *string `json:"ui_type" name:"ui_type" location:"params"`
	UserdataFile  *string `json:"userdata_file" name:"userdata_file" default:"/etc/rc.local" location:"params"`
	UserdataPath  *string `json:"userdata_path" name:"userdata_path" default:"/etc/qingcloud/userdata" location:"params"`
	// UserdataType's available values: plain, exec, tar
	UserdataType     *string   `json:"userdata_type" name:"userdata_type" location:"params"`
	UserdataValue    *string   `json:"userdata_value" name:"userdata_value" location:"params"`
	Volumes          []*string `json:"volumes" name:"volumes" location:"params"`
	VxNets           []*string `json:"vxnets" name:"vxnets" location:"params"`
	OsDiskEncryption *int      `json:"os_disk_encryption" name:"os_disk_encryption" location:"params"`
	NicMqueue        *int      `json:"nic_mqueue" name:"nic_mqueue" location:"params"`
	Platform         *string   `json:"platform" name:"platform" location:"params"`
	FResetpwd        *int      `json:"f_resetpwd" name:"f_resetpwd" location:"params"`
	ProcessorType    *string   `json:"processor_type" name:"processor_type" location:"params"`
	DefaultUser      *string   `json:"default_user" name:"default_user" location:"params"`
	DefaultPasswd    *string   `json:"default_passwd" name:"default_passwd" location:"params"`
	Hypervisor       *string   `json:"hypervisor" name:"hypervisor" location:"params"`
	GpuClass         *string   `json:"gpu_class" name:"gpu_class" location:"params"`
	PlaceGroupID     *string   `json:"place_group_id" name:"place_group_id" location:"params"`

	AutoRenew            *string `json:"auto_renew" name:"auto_renew" location:"params"`
	AutoVolumes          *string `json:"auto_volumes" name:"auto_volumes" location:"params"`
	Backups              *string `json:"backups" name:"backups" location:"params"`
	ChargeMode           *string `json:"charge_mode" name:"charge_mode" location:"params"`
	CipherAlg            *string `json:"cipher_alg" name:"cipher_alg" location:"params"`
	ContractDescription  *string `json:"contract_description" name:"contract_description" location:"params"`
	ContractEntries      *string `json:"contract_entries" name:"contract_entries" location:"params"`
	ContractID           *string `json:"contract_id" name:"contract_id" location:"params"`
	CouponID             *string `json:"coupon_id" name:"coupon_id" location:"params"`
	CPUTopology          *int    `json:"cpu_topology" name:"cpu_topology" location:"params"`
	DedicatedHostGroupID *string `json:"dedicated_host_group_id" name:"dedicated_host_group_id" location:"params"`
	DedicatedHostID      *string `json:"dedicated_host_id" name:"dedicated_host_id" location:"params"`
	DirectoryID          *string `json:"directory_id" name:"directory_id" location:"params"`
	// dry_run
	DryRun *string `json:"dry_run" name:"dry_run" default:"0" location:"params"`
	// eip_bandwidth.
	EIPBandwidth *string `json:"eip_bandwidth" name:"eip_bandwidth" location:"params"`
	// eip_bandwidth.
	EIPBillingMode *string `json:"eip_billing_mode" name:"eip_billing_mode" location:"params"`
	// eip_group.
	EIPGroup *string `json:"eip_group" name:"eip_group" location:"params"`
	// eips.
	EIPIDs *string `json:"eip_ids" name:"eip_ids" location:"params"`
	// eip billing
	EIPIgnoreContract *string `json:"eip_ignore_contract" name:"eip_ignore_contract" location:"params"`
	EIPVirgin         *string `json:"eip_virgin" name:"eip_virgin" location:"params"`
	Entries           *string `json:"entries" name:"entries" location:"params"`
	ExpirationTime    *string `json:"expiration_time" name:"expiration_time" location:"params"`
	// type of the physical GPU where the vGPU will be allocated
	GpuType            *string `json:"gpu_type" name:"gpu_type" location:"params"`
	InResourceGroupIDs *string `json:"in_resource_group_ids" name:"in_resource_group_ids" location:"params"`
	// instance_ext_type
	InstanceExtType *string `json:"instance_ext_type" name:"instance_ext_type" location:"params"`
	// instance_ext_type
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	// login_keypair
	LoginKeyPairList *string `json:"login_keypair_list" name:"login_keypair_list" location:"params"`
	MemoryCurrent    *string `json:"memory_current" name:"memory_current" location:"params"`
	MemoryMax        *string `json:"memory_max" name:"memory_max" location:"params"`
	Months           *string `json:"months" name:"months" location:"params"`
	NextChargeMode   *string `json:"next_charge_mode" name:"next_charge_mode" location:"params"`
	ProjectID        *string `json:"project_id" name:"project_id" location:"params"`
	PromotionID      *string `json:"promotion_id" name:"promotion_id" location:"params"`
	ReservedContract *string `json:"reserved_contract" name:"reserved_contract" location:"params"`
	// Whether to stop on error
	StopOnError *string `json:"stop_on_error" name:"stop_on_error" location:"params"`
	Tags        *string `json:"tags" name:"tags" location:"params"`
	// the user who will own this instance
	TargetUser            *string `json:"target_user" name:"target_user" location:"params"`
	Type                  *string `json:"type" name:"type" location:"params"`
	UserData              *string `json:"user_data" name:"user_data" location:"params"`
	Value                 *string `json:"value" name:"value" location:"params"`
	VCPUsCurrent          *string `json:"vcpus_current" name:"vcpus_current" location:"params"`
	VCPUsMax              *string `json:"vcpus_max" name:"vcpus_max" location:"params"`
	VdcNodeID             *string `json:"vdc_node_id" name:"vdc_node_id" location:"params"`
	VmDefinition          *string `json:"vm_definition" name:"vm_definition" location:"params"`
	VolumeContractEntries *string `json:"volume_contract_entries" name:"volume_contract_entries" location:"params"`
	// volume_encryption.
	VolumeEncryption *string `json:"volume_encryption" name:"volume_encryption" location:"params"`
	// volume_filesystem_type.
	VolumeFilesystemType *string `json:"volume_filesystem_type" name:"volume_filesystem_type" location:"params"`
	// volume_mount_point.
	VolumeMountPoint *string `json:"volume_mount_point" name:"volume_mount_point" location:"params"`
	VolumeRepl       *string `json:"volume_repl" name:"volume_repl" location:"params"`
	// volume_size.
	VolumeSize *string `json:"volume_size" name:"volume_size" location:"params"`
	// volume_type.
	VolumeType *string `json:"volume_type" name:"volume_type" location:"params"`
	// zone id to run instance to
	Zone *string `json:"zone" name:"zone" location:"params"`
}

func (v *RunInstancesInput) Validate() error {

	if v.CPU != nil {
		cpuValidValues := []string{"1", "2", "4", "8", "16"}
		cpuParameterValue := fmt.Sprint(*v.CPU)

		cpuIsValid := false
		for _, value := range cpuValidValues {
			if value == cpuParameterValue {
				cpuIsValid = true
			}
		}

		if !cpuIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "CPU",
				ParameterValue: cpuParameterValue,
				AllowedValues:  cpuValidValues,
			}
		}
	}

	if v.CPUMax != nil {
		cpuMaxValidValues := []string{"1", "2", "4", "8", "16"}
		cpuMaxParameterValue := fmt.Sprint(*v.CPUMax)

		cpuMaxIsValid := false
		for _, value := range cpuMaxValidValues {
			if value == cpuMaxParameterValue {
				cpuMaxIsValid = true
			}
		}

		if !cpuMaxIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "CPUMax",
				ParameterValue: cpuMaxParameterValue,
				AllowedValues:  cpuMaxValidValues,
			}
		}
	}

	if v.CPUModel != nil {
		cpuModelValidValues := []string{"Westmere", "SandyBridge", "IvyBridge", "Haswell", "Broadwell", "EPYC",
			"Skylake", "CascadeLake", "IceLake", "SapphireRapids", "Haswell-noTSX",
			"EPYC-Rome"}
		cpuModelParameterValue := fmt.Sprint(*v.CPUModel)

		cpuModelIsValid := false
		for _, value := range cpuModelValidValues {
			if value == cpuModelParameterValue {
				cpuModelIsValid = true
			}
		}

		if !cpuModelIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "CPUModel",
				ParameterValue: cpuModelParameterValue,
				AllowedValues:  cpuModelValidValues,
			}
		}
	}

	if v.ImageID == nil {
		return errors.ParameterRequiredError{
			ParameterName: "ImageID",
			ParentName:    "RunInstancesInput",
		}
	}

	if v.InstanceClass != nil {
		instanceClassValidValues := []string{"0", "1", "2", "3", "4", "5", "6", "100", "101", "200", "201", "300", "301"}
		instanceClassParameterValue := fmt.Sprint(*v.InstanceClass)

		instanceClassIsValid := false
		for _, value := range instanceClassValidValues {
			if value == instanceClassParameterValue {
				instanceClassIsValid = true
			}
		}

		if !instanceClassIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "InstanceClass",
				ParameterValue: instanceClassParameterValue,
				AllowedValues:  instanceClassValidValues,
			}
		}
	}

	if v.LoginMode == nil {
		return errors.ParameterRequiredError{
			ParameterName: "LoginMode",
			ParentName:    "RunInstancesInput",
		}
	}

	if v.LoginMode != nil {
		loginModeValidValues := []string{"keypair", "passwd"}
		loginModeParameterValue := fmt.Sprint(*v.LoginMode)

		loginModeIsValid := false
		for _, value := range loginModeValidValues {
			if value == loginModeParameterValue {
				loginModeIsValid = true
			}
		}

		if !loginModeIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "LoginMode",
				ParameterValue: loginModeParameterValue,
				AllowedValues:  loginModeValidValues,
			}
		}
	}

	if v.MemMax != nil {
		memMaxValidValues := []string{"1024", "2048", "4096", "6144", "8192", "12288", "16384", "24576", "32768"}
		memMaxParameterValue := fmt.Sprint(*v.MemMax)

		memMaxIsValid := false
		for _, value := range memMaxValidValues {
			if value == memMaxParameterValue {
				memMaxIsValid = true
			}
		}

		if !memMaxIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "MemMax",
				ParameterValue: memMaxParameterValue,
				AllowedValues:  memMaxValidValues,
			}
		}
	}

	if v.Memory != nil {
		memoryValidValues := []string{"1024", "2048", "4096", "6144", "8192", "12288", "16384", "24576", "32768"}
		memoryParameterValue := fmt.Sprint(*v.Memory)

		memoryIsValid := false
		for _, value := range memoryValidValues {
			if value == memoryParameterValue {
				memoryIsValid = true
			}
		}

		if !memoryIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "Memory",
				ParameterValue: memoryParameterValue,
				AllowedValues:  memoryValidValues,
			}
		}
	}

	if v.NeedNewSID != nil {
		needNewSIDValidValues := []string{"0", "1"}
		needNewSIDParameterValue := fmt.Sprint(*v.NeedNewSID)

		needNewSIDIsValid := false
		for _, value := range needNewSIDValidValues {
			if value == needNewSIDParameterValue {
				needNewSIDIsValid = true
			}
		}

		if !needNewSIDIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "NeedNewSID",
				ParameterValue: needNewSIDParameterValue,
				AllowedValues:  needNewSIDValidValues,
			}
		}
	}

	if v.NeedUserdata != nil {
		needUserdataValidValues := []string{"0", "1"}
		needUserdataParameterValue := fmt.Sprint(*v.NeedUserdata)

		needUserdataIsValid := false
		for _, value := range needUserdataValidValues {
			if value == needUserdataParameterValue {
				needUserdataIsValid = true
			}
		}

		if !needUserdataIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "NeedUserdata",
				ParameterValue: needUserdataParameterValue,
				AllowedValues:  needUserdataValidValues,
			}
		}
	}

	if v.UserdataType != nil {
		userdataTypeValidValues := []string{"plain", "exec", "tar"}
		userdataTypeParameterValue := fmt.Sprint(*v.UserdataType)

		userdataTypeIsValid := false
		for _, value := range userdataTypeValidValues {
			if value == userdataTypeParameterValue {
				userdataTypeIsValid = true
			}
		}

		if !userdataTypeIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "UserdataType",
				ParameterValue: userdataTypeParameterValue,
				AllowedValues:  userdataTypeValidValues,
			}
		}
	}

	return nil
}

type RunInstancesOutput struct {
	Message   *string   `json:"message" name:"message" location:"elements"`
	Action    *string   `json:"action" name:"action" location:"elements"`
	Instances []*string `json:"instances" name:"instances" location:"elements"`
	JobID     *string   `json:"job_id" name:"job_id" location:"elements"`
	RetCode   *int      `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/start_instances.html
func (s *InstanceService) StartInstances(i *StartInstancesInput) (*StartInstancesOutput, error) {
	if i == nil {
		i = &StartInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "StartInstances",
		RequestMethod: "GET",
	}

	x := &StartInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type StartInstancesInput struct {
	Instances []*string `json:"instances" name:"instances" location:"params"` // Required
	Volumes   *string   `json:"volumes" name:"volumes" location:"params"`
}

func (v *StartInstancesInput) Validate() error {

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "StartInstancesInput",
		}
	}

	return nil
}

type StartInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/stop_instances.html
func (s *InstanceService) StopInstances(i *StopInstancesInput) (*StopInstancesOutput, error) {
	if i == nil {
		i = &StopInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "StopInstances",
		RequestMethod: "GET",
	}

	x := &StopInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type StopInstancesInput struct {

	// Force's available values: 0, 1
	Force     *int      `json:"force" name:"force" default:"0" location:"params"`
	Instances []*string `json:"instances" name:"instances" location:"params"` // Required
}

func (v *StopInstancesInput) Validate() error {

	if v.Force != nil {
		forceValidValues := []string{"0", "1"}
		forceParameterValue := fmt.Sprint(*v.Force)

		forceIsValid := false
		for _, value := range forceValidValues {
			if value == forceParameterValue {
				forceIsValid = true
			}
		}

		if !forceIsValid {
			return errors.ParameterValueNotAllowedError{
				ParameterName:  "Force",
				ParameterValue: forceParameterValue,
				AllowedValues:  forceValidValues,
			}
		}
	}

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "StopInstancesInput",
		}
	}

	return nil
}

type StopInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/terminate_instances.html
func (s *InstanceService) TerminateInstances(i *TerminateInstancesInput) (*TerminateInstancesOutput, error) {
	if i == nil {
		i = &TerminateInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "TerminateInstances",
		RequestMethod: "GET",
	}

	x := &TerminateInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type TerminateInstancesInput struct {
	Instances []*string `json:"instances" name:"instances" location:"params"` // Required
}

func (v *TerminateInstancesInput) Validate() error {

	if len(v.Instances) == 0 {
		return errors.ParameterRequiredError{
			ParameterName: "Instances",
			ParentName:    "TerminateInstancesInput",
		}
	}

	return nil
}

type TerminateInstancesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// Documentation URL: https://docs.qingcloud.com/api/instance/clone_instances.html
func (s *InstanceService) CloneInstances(i *CloneInstancesInput) (*CloneInstancesOutput, error) {
	if i == nil {
		i = &CloneInstancesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CloneInstances",
		RequestMethod: "GET",
	}

	x := &CloneInstancesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type CloneInstancesInput struct {
	Instances *string `json:"instances" name:"instances" location:"params"`
	VxNets    *string `json:"vxnets" name:"vxnets" location:"params"`
}

func (v *CloneInstancesInput) Validate() error {

	return nil
}

type CloneInstancesOutput struct {
	Message      *string                    `json:"message" name:"message" location:"elements"`
	Action       *string                    `json:"action" name:"action" location:"elements"`
	JobID        *string                    `json:"job_id" name:"job_id" location:"elements"`
	RetCode      *int                       `json:"ret_code" name:"ret_code" location:"elements"`
	InstancesSet map[string]InstanceDetails `json:"instance_set"  name:"instance_set" location:"elements"`
	Instances    []*string                  `json:"instances" name:"instances" location:"elements"`
}

type InstanceDetails struct {
	InstanceMap map[string]string `json:"instance_map"`
	VolumesMap  map[string]string `json:"volumes_map"`
}

// CreateBrokers: CreateBrokers

func (s *InstanceService) CreateBrokers(i *CreateBrokersInput) (*CreateBrokersOutput, error) {
	if i == nil {
		i = &CreateBrokersInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CreateBrokers",
		RequestMethod: "GET",
	}

	x := &CreateBrokersOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type CreateBrokersInput struct {
	Instances *string `json:"instances" name:"instances" location:"params"`
}

func (v *CreateBrokersInput) Validate() error {

	return nil
}

type CreateBrokersOutput struct {
	Message *string  `json:"message" name:"message" location:"elements"`
	Action  *string  `json:"action" name:"action" location:"elements"`
	JobID   *string  `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int     `json:"ret_code" name:"ret_code" location:"elements"`
	Brokers []Broker `json:"brokers"`
}
type Broker struct {
	InstanceID string `json:"instance_id"`
	BrokerPort int    `json:"broker_port"`
	BrokerHost string `json:"broker_host"`
}

// DeleteBrokers: DeleteBrokers

func (s *InstanceService) DeleteBrokers(i *DeleteBrokersInput) (*DeleteBrokersOutput, error) {
	if i == nil {
		i = &DeleteBrokersInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DeleteBrokers",
		RequestMethod: "GET",
	}

	x := &DeleteBrokersOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DeleteBrokersInput struct {
	Instances *string `json:"instances" name:"instances" location:"params"`
}

func (v *DeleteBrokersInput) Validate() error {

	return nil
}

type DeleteBrokersOutput struct {
	Message *string  `json:"message" name:"message" location:"elements"`
	Action  *string  `json:"action" name:"action" location:"elements"`
	JobID   *string  `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int     `json:"ret_code" name:"ret_code" location:"elements"`
	Brokers []Broker `json:"brokers"  name:"brokers" location:"elements"`
}

// ApplyInstanceGroup: ApplyInstanceGroup
func (s *InstanceService) ApplyInstanceGroup(i *ApplyInstanceGroupInput) (*ApplyInstanceGroupOutput, error) {
	if i == nil {
		i = &ApplyInstanceGroupInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ApplyInstanceGroup",
		RequestMethod: "GET",
	}

	x := &ApplyInstanceGroupOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type ApplyInstanceGroupInput struct {
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	Zone          *string `json:"zone" name:"zone" location:"params"`
}

func (v *ApplyInstanceGroupInput) Validate() error {

	return nil
}

type ApplyInstanceGroupOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// CreateInstanceGroups: CreateInstanceGroups

func (s *InstanceService) CreateInstanceGroups(i *CreateInstanceGroupsInput) (*CreateInstanceGroupsOutput, error) {
	if i == nil {
		i = &CreateInstanceGroupsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CreateInstanceGroups",
		RequestMethod: "GET",
	}

	x := &CreateInstanceGroupsOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type CreateInstanceGroupsInput struct {

	// The number of instance group to create.
	Count *int `json:"count" name:"count" default:"1" location:"params"`
	// The description of the instance group
	Description *string `json:"description" name:"description" location:"params"`
	// the short name of instance_group you want to create.
	InstanceGroupName *string `json:"instance_group_name" name:"instance_group_name" location:"params"`
	ProjectID         *string `json:"project_id" name:"project_id" location:"params"`
	// The instance group relation. Supported relations are `repel` or `attract`
	Relation *string `json:"relation" name:"relation" location:"params"`
	Zone     *string `json:"zone" name:"zone" location:"params"`
}

func (v *CreateInstanceGroupsInput) Validate() error {

	return nil
}

type CreateInstanceGroupsOutput struct {
	InstanceGroups []string `json:"instance_groups"  name:"instance_groups"  location:"elements"`
	Message        *string  `json:"message" name:"message"`
	Action         *string  `json:"action" name:"action" location:"elements"`
	JobID          *string  `json:"job_id" name:"job_id" location:"elements"`
	RetCode        *int     `json:"ret_code" name:"ret_code" location:"elements"`
}

// DeleteInstanceGroups: DeleteInstanceGroups

func (s *InstanceService) DeleteInstanceGroups(i *DeleteInstanceGroupsInput) (*DeleteInstanceGroupsOutput, error) {
	if i == nil {
		i = &DeleteInstanceGroupsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DeleteInstanceGroups",
		RequestMethod: "GET",
	}

	x := &DeleteInstanceGroupsOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DeleteInstanceGroupsInput struct {

	// the IDs of instance groups you want to delete.
	InstanceGroups []*string `json:"instance_groups" name:"instance_groups" location:"params"`
	Zone           *string   `json:"zone" name:"zone" location:"params"`
}

func (v *DeleteInstanceGroupsInput) Validate() error {

	return nil
}

type DeleteInstanceGroupsOutput struct {
	Message        *string  `json:"message" name:"message" location:"elements"`
	Action         *string  `json:"action" name:"action" location:"elements"`
	JobID          *string  `json:"job_id" name:"job_id" location:"elements"`
	InstanceGroups []string `json:"instance_groups"  name:"instance_groups"  location:"elements"`
	RetCode        *int     `json:"ret_code" name:"ret_code" location:"elements"`
}

// DescribeInstanceGroups: DescribeInstanceGroups

func (s *InstanceService) DescribeInstanceGroups(i *DescribeInstanceGroupsInput) (*DescribeInstanceGroupsOutput, error) {
	if i == nil {
		i = &DescribeInstanceGroupsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeInstanceGroups",
		RequestMethod: "GET",
	}

	x := &DescribeInstanceGroupsOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeInstanceGroupsInput struct {
	ConsoleID *string `json:"console_id" name:"console_id" location:"params"`
	// the name of the instance group. Support partial match.
	InstanceGroupName *string `json:"instance_group_name" name:"instance_group_name" location:"params"`
	// the comma separated IDs of instance_groups you want to list.
	InstanceGroups []*string `json:"instance_groups" name:"instance_groups" location:"params"`
	// specify the number of the returning results.
	Limit *int `json:"limit" name:"limit" default:"20" location:"params"`
	// the starting offset of the returning results.
	Offset *int `json:"offset" name:"offset" location:"params"`
	// the owner id of instance_groups
	Owner []*string `json:"owner" name:"owner" location:"params"`
	// project id
	ProjectID *string `json:"project_id" name:"project_id" location:"params"`
	// filter by instance_group relation
	Relation []*string `json:"relation" name:"relation" location:"params"`
	// for reverse sorting. 1: reverse
	Reverse    *int    `json:"reverse" name:"reverse" location:"params"`
	RootUserID *string `json:"root_user_id" name:"root_user_id" location:"params"`
	SearchWord *string `json:"search_word" name:"search_word" location:"params"`
	// sort key
	SortKey *string `json:"sort_key" name:"sort_key" location:"params"`
	// filter by tags
	Tags []*string `json:"tags" name:"tags" location:"params"`
	// the number to specify the verbose level
	Verbose *int    `json:"verbose" name:"verbose" default:"0" location:"params"`
	Zone    *string `json:"zone" name:"zone" location:"params"`
}

func (v *DescribeInstanceGroupsInput) Validate() error {

	return nil
}

type DescribeInstanceGroupsOutput struct {
	Message        *string          `json:"message" name:"message" location:"elements"`
	Action         *string          `json:"action" name:"action" location:"elements"`
	JobID          *string          `json:"job_id" name:"job_id" location:"elements"`
	InstanceGroups []*InstanceGroup `json:"instance_group_set"  name:"instance_group_set"  location:"elements"`
	RetCode        *int             `json:"ret_code" name:"ret_code" location:"elements"`
}
type InstanceGroup struct {
	InstanceGroupName *string   `json:"instance_group_name"`
	Description       *string   `json:"description"`
	Tags              []*string `json:"tags"`
	Controller        *string   `json:"controller"`
	ConsoleID         *string   `json:"console_id"`
	RootUserID        *string   `json:"root_user_id"`
	CreateTime        *string   `json:"create_time"`
	Relation          *string   `json:"relation"`
	Owner             *string   `json:"owner"`
	InstanceGroupID   *string   `json:"instance_group_id"`
}

// ModifyInstanceGroupAttributes: ModifyInstanceGroupAttributes

func (s *InstanceService) ModifyInstanceGroupAttributes(i *ModifyInstanceGroupAttributesInput) (*ModifyInstanceGroupAttributesOutput, error) {
	if i == nil {
		i = &ModifyInstanceGroupAttributesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ModifyInstanceGroupAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyInstanceGroupAttributesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type ModifyInstanceGroupAttributesInput struct {

	// The detailed description of the resource
	Description   *string `json:"description" name:"description" location:"params"`
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	// specify the new instance_group name.
	InstanceGroupName *string `json:"instance_group_name" name:"instance_group_name" location:"params"`
	// The instance group relation. Supported relations are `repel` or `attract`
	Relation *string `json:"relation" name:"relation" location:"params"`
	Zone     *string `json:"zone" name:"zone" location:"params"`
}

func (v *ModifyInstanceGroupAttributesInput) Validate() error {

	return nil
}

type ModifyInstanceGroupAttributesOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// JoinInstanceGroup: JoinInstanceGroup

func (s *InstanceService) JoinInstanceGroup(i *JoinInstanceGroupInput) (*JoinInstanceGroupOutput, error) {
	if i == nil {
		i = &JoinInstanceGroupInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "JoinInstanceGroup",
		RequestMethod: "GET",
	}

	x := &JoinInstanceGroupOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type JoinInstanceGroupInput struct {

	// the id of instance_group the instances will join.
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	// the IDs of instances that will join a instance_group.
	Instances []*string `json:"instances" name:"instances" location:"params"`
	Zone      *string   `json:"zone" name:"zone" location:"params"`
}

func (v *JoinInstanceGroupInput) Validate() error {

	return nil
}

type JoinInstanceGroupOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// LeaveInstanceGroup: LeaveInstanceGroup

func (s *InstanceService) LeaveInstanceGroup(i *LeaveInstanceGroupInput) (*LeaveInstanceGroupOutput, error) {
	if i == nil {
		i = &LeaveInstanceGroupInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "LeaveInstanceGroup",
		RequestMethod: "GET",
	}

	x := &LeaveInstanceGroupOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type LeaveInstanceGroupInput struct {

	// the id of the instance_group the instances will leave.
	InstanceGroup *string `json:"instance_group" name:"instance_group" location:"params"`
	// the IDs of instances that will leave a instance_group.
	Instances []*string `json:"instances" name:"instances" location:"params"`
	Zone      *string   `json:"zone" name:"zone" location:"params"`
}

func (v *LeaveInstanceGroupInput) Validate() error {

	return nil
}

type LeaveInstanceGroupOutput struct {
	Message *string `json:"message" name:"message" location:"elements"`
	Action  *string `json:"action" name:"action" location:"elements"`
	JobID   *string `json:"job_id" name:"job_id" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}
