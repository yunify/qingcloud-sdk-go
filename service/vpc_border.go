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
)

var _ fmt.State
var _ time.Time

type VpcBorderService struct {
	Config     *config.Config
	Properties *VpcBorderServiceProperties
}

type VpcBorderServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) VpcBorder(zone string) (*VpcBorderService, error) {
	properties := &VpcBorderServiceProperties{
		Zone: &zone,
	}

	return &VpcBorderService{Config: s.Config, Properties: properties}, nil
}

// AddBorderStatics: AddBorderStatics

func (s *VpcBorderService) AddBorderStatics(i *AddBorderStaticsInput) (*AddBorderStaticsOutput, error) {
	if i == nil {
		i = &AddBorderStaticsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "AddBorderStatics",
		RequestMethod: "GET",
	}

	x := &AddBorderStaticsOutput{}
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

type AddBorderStaticsInput struct {

	// the ID of intranet router whose statics you want to add.
	Border *string `json:"border" name:"border" location:"params"`
	// a json string of rules list. e.g.[{static_type:0
	Statics []*string `json:"statics" name:"statics" location:"params"`
}

func (v *AddBorderStaticsInput) Validate() error {

	return nil
}

type AddBorderStaticsOutput struct {
	Message *string `json:"message" name:"message"`
	Action  *string `json:"action" name:"action" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// AssociateBorder: AssociateBorder

func (s *VpcBorderService) AssociateBorder(i *AssociateBorderInput) (*AssociateBorderOutput, error) {
	if i == nil {
		i = &AssociateBorderInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "AssociateBorder",
		RequestMethod: "GET",
	}

	x := &AssociateBorderOutput{}
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

type AssociateBorderInput struct {

	// the intranet router you want to associate.
	Border *string `json:"border" name:"border" location:"params"`
	// the id of the vpc router.
	Router *string `json:"router" name:"router" location:"params"`
}

func (v *AssociateBorderInput) Validate() error {

	return nil
}

type AssociateBorderOutput struct {
	Message  *string `json:"message" name:"message"`
	Action   *string `json:"action" name:"action" location:"elements"`
	RetCode  *int    `json:"ret_code" name:"ret_code" location:"elements"`
	RouterID *string `json:"router_id" name:"router_id" location:"elements"`
	BorderID *string `json:"border_id" name:"border_id" location:"elements"`
}

// ConfigBorder: ConfigBorder

func (s *VpcBorderService) ConfigBorder(i *ConfigBorderInput) (*ConfigBorderOutput, error) {
	if i == nil {
		i = &ConfigBorderInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ConfigBorder",
		RequestMethod: "GET",
	}

	x := &ConfigBorderOutput{}
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

type ConfigBorderInput struct {

	// the ID of border you want to configure.
	Border *string `json:"border" name:"border" location:"params"`
	// configuration data in json format
	Data *string `json:"data" name:"data" location:"params"`
	// operation such as CreateSubIf
	Operation *string `json:"operation" name:"operation" location:"params"`
}

func (v *ConfigBorderInput) Validate() error {

	return nil
}

type ConfigBorderOutput struct {
	Message  *string `json:"message" name:"message"`
	Action   *string `json:"action" name:"action" location:"elements"`
	RetCode  *int    `json:"ret_code" name:"ret_code" location:"elements"`
	BorderID *string `json:"border_id" name:"border_id" location:"elements"`
}

// CreateVpcBorders: CreateVpcBorders

func (s *VpcBorderService) CreateVpcBorders(i *CreateVpcBordersInput) (*CreateVpcBordersOutput, error) {
	if i == nil {
		i = &CreateVpcBordersInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "CreateVpcBorders",
		RequestMethod: "GET",
	}

	x := &CreateVpcBordersOutput{}
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

type CreateVpcBordersInput struct {
	BorderName  *string `json:"border_name" name:"border_name" location:"params"`
	BorderType  *string `json:"border_type" name:"border_type" location:"params"`
	Description *string `json:"description" name:"description" location:"params"`
}

func (v *CreateVpcBordersInput) Validate() error {

	return nil
}

type CreateVpcBordersOutput struct {
	Message    *string   `json:"message" name:"message"`
	Action     *string   `json:"action" name:"action" location:"elements"`
	RetCode    *int      `json:"ret_code" name:"ret_code" location:"elements"`
	VpcBorders []*string `json:"vpc_borders" name:"vpc_borders" location:"elements"`
}

// DeleteBorderStatics: DeleteBorderStatics

func (s *VpcBorderService) DeleteBorderStatics(i *DeleteBorderStaticsInput) (*DeleteBorderStaticsOutput, error) {
	if i == nil {
		i = &DeleteBorderStaticsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DeleteBorderStatics",
		RequestMethod: "GET",
	}

	x := &DeleteBorderStaticsOutput{}
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

type DeleteBorderStaticsInput struct {

	// the comma separated IDs of border statics you want to delete.
	BorderStatics []*string `json:"border_statics" name:"border_statics" location:"params"`
}

func (v *DeleteBorderStaticsInput) Validate() error {

	return nil
}

type DeleteBorderStaticsOutput struct {
	Message *string `json:"message" name:"message"`
	Action  *string `json:"action" name:"action" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// DeleteVpcBorders: DeleteVpcBorders

func (s *VpcBorderService) DeleteVpcBorders(i *DeleteVpcBordersInput) (*DeleteVpcBordersOutput, error) {
	if i == nil {
		i = &DeleteVpcBordersInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DeleteVpcBorders",
		RequestMethod: "GET",
	}

	x := &DeleteVpcBordersOutput{}
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

type DeleteVpcBordersInput struct {
	VpcBorders []*string `json:"vpc_borders" name:"vpc_borders" location:"params"`
}

func (v *DeleteVpcBordersInput) Validate() error {

	return nil
}

type DeleteVpcBordersOutput struct {
	Message *string `json:"message" name:"message"`
	Action  *string `json:"action" name:"action" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// DescribeBorderStatics: DescribeBorderStatics

func (s *VpcBorderService) DescribeBorderStatics(i *DescribeBorderStaticsInput) (*DescribeBorderStaticsOutput, error) {
	if i == nil {
		i = &DescribeBorderStaticsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeBorderStatics",
		RequestMethod: "GET",
	}

	x := &DescribeBorderStaticsOutput{}
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

type DescribeBorderStaticsInput struct {

	// filter by border.
	Border *string `json:"border" name:"border" location:"params"`
	// the comma separated IDs of border_statics you want to list.
	BorderStatics []*string `json:"border_statics" name:"border_statics" location:"params"`
	// specify the number of the returning results.
	Limit *int `json:"limit" name:"limit" location:"params"`
	// the starting offset of the returning results.
	Offset *int `json:"offset" name:"offset" location:"params"`
	// filter by owner
	Owner *string `json:"owner" name:"owner" location:"params"`
	// a list of static type. 0: route.
	StaticType []*string `json:"static_type" name:"static_type" location:"params"`
	// the number to specify the verbose level
	Verbose *int `json:"verbose" name:"verbose" location:"params"`
}

func (v *DescribeBorderStaticsInput) Validate() error {

	return nil
}

type DescribeBorderStaticsOutput struct {
	Message *string `json:"message" name:"message"`
	Action  *string `json:"action" name:"action" location:"elements"`
	RetCode *int    `json:"ret_code" name:"ret_code" location:"elements"`
}

// DescribeBorderVxNets: DescribeBorderVxnets

func (s *VpcBorderService) DescribeBorderVxNets(i *DescribeBorderVxNetsInput) (*DescribeBorderVxNetsOutput, error) {
	if i == nil {
		i = &DescribeBorderVxNetsInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeBorderVxnets",
		RequestMethod: "GET",
	}

	x := &DescribeBorderVxNetsOutput{}
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

type DescribeBorderVxNetsInput struct {
	Border *string `json:"border" name:"border" location:"params"`
	// specify the number of the returning results.
	Limit *int `json:"limit" name:"limit" location:"params"`
	// the starting offset of the returning results.
	Offset *int `json:"offset" name:"offset" location:"params"`
	// filter by owner.
	Owner *string `json:"owner" name:"owner" location:"params"`
	VxNet *string `json:"vxnet" name:"vxnet" location:"params"`
}

func (v *DescribeBorderVxNetsInput) Validate() error {

	return nil
}

type DescribeBorderVxNetsOutput struct {
	Message     *string        `json:"message" name:"message"`
	Action      *string        `json:"action" name:"action" location:"elements"`
	RetCode     *int           `json:"ret_code" name:"ret_code" location:"elements"`
	BorderVxnet []*BorderVxnet `json:"border_vxnet_set"  name:"border_vxnet_set" location:"elements"`
}

// DescribeVpcBorders: DescribeVpcBorders

func (s *VpcBorderService) DescribeVpcBorders(i *DescribeVpcBordersInput) (*DescribeVpcBordersOutput, error) {
	if i == nil {
		i = &DescribeVpcBordersInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeVpcBorders",
		RequestMethod: "GET",
	}

	x := &DescribeVpcBordersOutput{}
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

type DescribeVpcBordersInput struct {
	BorderName *string   `json:"border_name" name:"border_name" location:"params"`
	BorderType *string   `json:"border_type" name:"border_type" location:"params"`
	L3vni      *int      `json:"l3vni" name:"l3vni" location:"params"`
	Limit      *int      `json:"limit" name:"limit" location:"params"`
	Offset     *int      `json:"offset" name:"offset" location:"params"`
	Owner      *string   `json:"owner" name:"owner" location:"params"`
	RouterID   *string   `json:"router_id" name:"router_id" location:"params"`
	SearchWord *string   `json:"search_word" name:"search_word" location:"params"`
	Status     []*string `json:"status" name:"status" location:"params"`
	Verbose    *int      `json:"verbose" name:"verbose" location:"params"`
	VpcBorders []*string `json:"vpc_borders" name:"vpc_borders" location:"params"`
}

func (v *DescribeVpcBordersInput) Validate() error {

	return nil
}

type DescribeVpcBordersOutput struct {
	Message      *string      `json:"message" name:"message"`
	Action       *string      `json:"action" name:"action" location:"elements"`
	RetCode      *int         `json:"ret_code" name:"ret_code" location:"elements"`
	TotalCount   *int         `json:"total_count" name:"total_count" location:"elements"`
	VpcBorderSet []*VpcBorder `json:"vpc_border_set" name:"vpc_border_set" location:"elements"`
}

// DissociateBorder: DissociateBorder

func (s *VpcBorderService) DissociateBorder(i *DissociateBorderInput) (*DissociateBorderOutput, error) {
	if i == nil {
		i = &DissociateBorderInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DissociateBorder",
		RequestMethod: "GET",
	}

	x := &DissociateBorderOutput{}
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

type DissociateBorderInput struct {

	// the intranet router you want to dissociate.
	Border *string `json:"border" name:"border" location:"params"`
	// the id of the vpc router.
	Router *string `json:"router" name:"router" location:"params"`
}

func (v *DissociateBorderInput) Validate() error {

	return nil
}

type DissociateBorderOutput struct {
	Message  *string `json:"message" name:"message"`
	Action   *string `json:"action" name:"action" location:"elements"`
	RetCode  *int    `json:"ret_code" name:"ret_code" location:"elements"`
	RouterID *string `json:"router_id" name:"router_id" location:"elements"`
	BorderID *string `json:"border_id" name:"border_id" location:"elements"`
}

// JoinBorder: JoinBorder

func (s *VpcBorderService) JoinBorder(i *JoinBorderInput) (*JoinBorderOutput, error) {
	if i == nil {
		i = &JoinBorderInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "JoinBorder",
		RequestMethod: "GET",
	}

	x := &JoinBorderOutput{}
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

type JoinBorderInput struct {

	// the intranet router you want to join.
	Border *string `json:"border" name:"border" location:"params"`
	// specify the border private ip for each vxnet
	BorderPrivateIPs []*string `json:"border_private_ips" name:"border_private_ips" location:"params"`
	// the ids of the vxnets that will join the intranet router.
	VxNets []*string `json:"vxnets" name:"vxnets" location:"params"`
}

func (v *JoinBorderInput) Validate() error {

	return nil
}

type JoinBorderOutput struct {
	Message  *string `json:"message" name:"message"`
	Action   *string `json:"action" name:"action" location:"elements"`
	RetCode  *int    `json:"ret_code" name:"ret_code" location:"elements"`
	BorderID *string `json:"border_id" name:"border_id" location:"elements"`
	VxNetID  *string `json:"vxnet_id" name:"vxnet_id" location:"elements"`
}

// LeaveBorder: LeaveBorder

func (s *VpcBorderService) LeaveBorder(i *LeaveBorderInput) (*LeaveBorderOutput, error) {
	if i == nil {
		i = &LeaveBorderInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "LeaveBorder",
		RequestMethod: "GET",
	}

	x := &LeaveBorderOutput{}
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

type LeaveBorderInput struct {

	// the intranet router you want to leave.
	Border *string `json:"border" name:"border" location:"params"`
	// force leave
	Force *int `json:"force" name:"force" default:"0" location:"params"`
	// the comm separated IDs of the vxnets that will leave the intranet router.
	VxNets []*string `json:"vxnets" name:"vxnets" location:"params"`
	Zone   *string   `json:"zone" name:"zone" location:"params"`
}

func (v *LeaveBorderInput) Validate() error {

	return nil
}

type LeaveBorderOutput struct {
	Message  *string   `json:"message" name:"message"`
	Action   *string   `json:"action" name:"action" location:"elements"`
	RetCode  *int      `json:"ret_code" name:"ret_code" location:"elements"`
	BorderID *string   `json:"border_id" name:"border_id" location:"elements"`
	Vxnets   []*string `json:"vxnets" name:"vxnets" location:"elements"`
}

// ModifyBorderAttributes: ModifyBorderAttributes

func (s *VpcBorderService) ModifyBorderAttributes(i *ModifyBorderAttributesInput) (*ModifyBorderAttributesOutput, error) {
	if i == nil {
		i = &ModifyBorderAttributesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "ModifyBorderAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyBorderAttributesOutput{}
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

type ModifyBorderAttributesInput struct {

	// the intranet router you want to modify.
	Border *string `json:"border" name:"border" location:"params"`
	// the new border_name.
	BorderName *string `json:"border_name" name:"border_name" location:"params"`
	// the new description.
	Description *string `json:"description" name:"description" location:"params"`
}

func (v *ModifyBorderAttributesInput) Validate() error {

	return nil
}

type ModifyBorderAttributesOutput struct {
	Message     *string `json:"message" name:"message"`
	Action      *string `json:"action" name:"action" location:"elements"`
	RetCode     *int    `json:"ret_code" name:"ret_code" location:"elements"`
	GlobalJobID *string `json:"global_job_id" name:"global_job_id" location:"elements"`
}
