// Code generated by protoc-gen-goext. DO NOT EDIT.

package vpc

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetRouteTableRequest) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *ListRouteTablesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListRouteTablesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRouteTablesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRouteTablesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListRouteTablesResponse) SetRouteTables(v []*RouteTable) {
	m.RouteTables = v
}

func (m *ListRouteTablesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateRouteTableRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateRouteTableRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateRouteTableRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateRouteTableRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateRouteTableRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateRouteTableRequest) SetStaticRoutes(v []*StaticRoute) {
	m.StaticRoutes = v
}

func (m *CreateRouteTableMetadata) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *UpdateRouteTableRequest) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *UpdateRouteTableRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRouteTableRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateRouteTableRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateRouteTableRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateRouteTableRequest) SetStaticRoutes(v []*StaticRoute) {
	m.StaticRoutes = v
}

func (m *UpdateRouteTableMetadata) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *DeleteRouteTableRequest) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *DeleteRouteTableMetadata) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *ListRouteTableOperationsRequest) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *ListRouteTableOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRouteTableOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRouteTableOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListRouteTableOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *MoveRouteTableRequest) SetRouteTableId(v string) {
	m.RouteTableId = v
}

func (m *MoveRouteTableRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveRouteTableMetadata) SetRouteTableId(v string) {
	m.RouteTableId = v
}
