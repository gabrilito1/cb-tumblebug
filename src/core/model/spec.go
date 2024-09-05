/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package model is to handle object of CB-Tumblebug
package model

import (

	//"github.com/cloud-barista/cb-tumblebug/src/core/mci"

	_ "github.com/go-sql-driver/mysql"
)

// SpiderSpecInfo is a struct to create JSON body of 'Get spec request'
type SpiderSpecInfo struct {
	// https://github.com/cloud-barista/cb-spider/blob/master/cloud-control-manager/cloud-driver/interfaces/resources/VMSpecHandler.go

	Region string
	Name   string
	VCpu   SpiderVCpuInfo
	Mem    string
	Gpu    []SpiderGpuInfo

	KeyValueList []KeyValue
}

// SpiderVCpuInfo is a struct to handle vCPU Info from CB-Spider.
type SpiderVCpuInfo struct {
	Count string
	Clock string // GHz
}

// SpiderGpuInfo is a struct to handle GPU Info from CB-Spider.
type SpiderGpuInfo struct {
	Count string
	Mfr   string
	Model string
	Mem   string
}

// TbSpecReq is a struct to handle 'Register spec' request toward CB-Tumblebug.
type TbSpecReq struct { // Tumblebug
	Name           string `json:"name" validate:"required"`
	ConnectionName string `json:"connectionName" validate:"required"`
	CspResourceId  string `json:"cspResourceId" validate:"required"`
	Description    string `json:"description"`
}

// TbSpecInfo is a struct that represents TB spec object.
type TbSpecInfo struct { // Tumblebug
	// Id is unique identifier for the object
	Id string `json:"id" example:"aws-ap-southeast-1"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`
	// CspResourceHandlingName is identifier to handle CSP resource
	CspResourceHandlingName string `json:"cspResourceHandlingName,omitempty" example:"we12fawefadf1221edcf"`
	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`

	// Name is human-readable string to represent the object
	Name           string `json:"name" example:"aws-ap-southeast-1"`
	Namespace      string `json:"namespace,omitempty" example:"default"`
	ConnectionName string `json:"connectionName,omitempty"`
	ProviderName   string `json:"providerName,omitempty"`
	RegionName     string `json:"regionName,omitempty"`
	// vm|k8s|kubernetes|container, etc.
	InfraType             string   `json:"infraType,omitempty"`
	OsType                string   `json:"osType,omitempty"`
	VCPU                  uint16   `json:"vCPU,omitempty"`
	MemoryGiB             float32  `json:"memoryGiB,omitempty"`
	StorageGiB            uint32   `json:"storageGiB,omitempty"`
	MaxTotalStorageTiB    uint16   `json:"maxTotalStorageTiB,omitempty"`
	NetBwGbps             uint16   `json:"netBwGbps,omitempty"`
	AcceleratorModel      string   `json:"acceleratorModel,omitempty"`
	AcceleratorCount      uint8    `json:"acceleratorCount,omitempty"`
	AcceleratorMemoryGB   float32  `json:"acceleratorMemoryGB,omitempty"`
	AcceleratorType       string   `json:"acceleratorType,omitempty"`
	CostPerHour           float32  `json:"costPerHour,omitempty"`
	Description           string   `json:"description,omitempty"`
	OrderInFilteredResult uint16   `json:"orderInFilteredResult,omitempty"`
	EvaluationStatus      string   `json:"evaluationStatus,omitempty"`
	EvaluationScore01     float32  `json:"evaluationScore01"`
	EvaluationScore02     float32  `json:"evaluationScore02"`
	EvaluationScore03     float32  `json:"evaluationScore03"`
	EvaluationScore04     float32  `json:"evaluationScore04"`
	EvaluationScore05     float32  `json:"evaluationScore05"`
	EvaluationScore06     float32  `json:"evaluationScore06"`
	EvaluationScore07     float32  `json:"evaluationScore07"`
	EvaluationScore08     float32  `json:"evaluationScore08"`
	EvaluationScore09     float32  `json:"evaluationScore09"`
	EvaluationScore10     float32  `json:"evaluationScore10"`
	RootDiskType          string   `json:"rootDiskType"`
	RootDiskSize          string   `json:"rootDiskSize"`
	AssociatedObjectList  []string `json:"associatedObjectList,omitempty"`
	IsAutoGenerated       bool     `json:"isAutoGenerated,omitempty"`

	// SystemLabel is for describing the Resource in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel,omitempty" example:"Managed by CB-Tumblebug" default:""`
}

// FilterSpecsByRangeRequest is for 'FilterSpecsByRange'
type FilterSpecsByRangeRequest struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	ConnectionName      string `json:"connectionName"`
	ProviderName        string `json:"providerName"`
	RegionName          string `json:"regionName"`
	CspResourceId       string `json:"cspResourceId"`
	InfraType           string `json:"infraType"`
	OsType              string `json:"osType"`
	VCPU                Range  `json:"vCPU"`
	MemoryGiB           Range  `json:"memoryGiB"`
	StorageGiB          Range  `json:"storageGiB"`
	MaxTotalStorageTiB  Range  `json:"maxTotalStorageTiB"`
	NetBwGbps           Range  `json:"netBwGbps"`
	AcceleratorModel    string `json:"acceleratorModel"`
	AcceleratorCount    Range  `json:"acceleratorCount"`
	AcceleratorMemoryGB Range  `json:"acceleratorMemoryGB"`
	AcceleratorType     string `json:"acceleratorType"`
	CostPerHour         Range  `json:"costPerHour"`
	Description         string `json:"description"`
	EvaluationStatus    string `json:"evaluationStatus"`
	EvaluationScore01   Range  `json:"evaluationScore01"`
	EvaluationScore02   Range  `json:"evaluationScore02"`
	EvaluationScore03   Range  `json:"evaluationScore03"`
	EvaluationScore04   Range  `json:"evaluationScore04"`
	EvaluationScore05   Range  `json:"evaluationScore05"`
	EvaluationScore06   Range  `json:"evaluationScore06"`
	EvaluationScore07   Range  `json:"evaluationScore07"`
	EvaluationScore08   Range  `json:"evaluationScore08"`
	EvaluationScore09   Range  `json:"evaluationScore09"`
	EvaluationScore10   Range  `json:"evaluationScore10"`
}

// SpiderSpecList is a struct to handle spec list from the CB-Spider's REST API response
type SpiderSpecList struct {
	Vmspec []SpiderSpecInfo `json:"vmspec"`
}

// Range struct is for 'FilterSpecsByRange'
type Range struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}
