/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreateNetworkInterfaceRequest struct {

	// Private Subnet인스턴스번호
PrivateSubnetInstanceNo *string `json:"privateSubnetInstanceNo"`

	// Network Interface이름
NetworkInterfaceName *string `json:"networkInterfaceName"`

	// Network Interface IP
NetworkInterfaceIp *string `json:"networkInterfaceIp"`

	// Network Interface설명
NetworkInterfaceDescription *string `json:"networkInterfaceDescription,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
}
