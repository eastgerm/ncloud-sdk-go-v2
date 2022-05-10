/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetElasticsearchMonitoringData struct {

	// 조회 시작 시간(Millisecond)
	TimeStart int32 `json:"timeStart"`

	// 조회 종료 시간(Millisecond)
	TimeEnd int32 `json:"timeEnd"`

	// 조회하려는 metric 명
	Metric string `json:"metric"`

	// 조회할 서버의 Compute Instance No
	ComputeInstanceNo string `json:"computeInstanceNo"`

	// 집계 주기(default: Min1)
	Interval string `json:"interval,omitempty"`
}
