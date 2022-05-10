/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

type GetClusterRequest struct {

	// 조회할 Cluster의 이름
	InputText string `json:"inputText,omitempty"`

	// 해당 VPC를 사용하고 있는 Cluster 조회
	VpcName string `json:"vpcName,omitempty"`

	// 페이지 번호
	PageNo int32 `json:"pageNo,omitempty"`

	// 페이지 사이즈
	PageSize int32 `json:"pageSize,omitempty"`
}
