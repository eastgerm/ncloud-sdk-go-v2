/*
 * vmssql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmssql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmssql

type GetCloudMssqlBackupDetailListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

	// CloudMssql백업상세리스트
CloudMssqlBackupDetailList *CloudMssqlBackupDetailList `json:"cloudMssqlBackupDetailList,omitempty"`
}
