/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetPortForwardingRuleListResponse struct {

	// 포트포워딩설정번호
PortForwardingConfigurationNo *string `json:"portForwardingConfigurationNo,omitempty"`

	// 포트포워딩공인IP
PortForwardingPublicIp *string `json:"portForwardingPublicIp,omitempty"`

	// ZONE객체
Zone *Zone `json:"zone,omitempty"`

	// 인터넷라인구분
InternetLineType *CommonCode `json:"internetLineType,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

PortForwardingRuleList []*PortForwardingRule `json:"portForwardingRuleList,omitempty"`
}
