// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bcs-data-manager/bcs-data-manager.proto

package datamanager

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DataManager service

func NewDataManagerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "DataManager.GetAllProjectList",
			Path:    []string{"/datamanager/v1/projects"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetProjectInfo",
			Path:    []string{"/datamanager/v1/project"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetAllClusterList",
			Path:    []string{"/datamanager/v1/clusters"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetClusterListByProject",
			Path:    []string{"/datamanager/v1/projects/clusters"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetClusterInfo",
			Path:    []string{"/datamanager/v1/clusters/{clusterID}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetNamespaceInfoList",
			Path:    []string{"/datamanager/v1/clusters/{clusterID}/namespaces"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetNamespaceInfo",
			Path:    []string{"/datamanager/v1/clusters/{clusterID}/namespaces/{namespace}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetWorkloadInfoList",
			Path:    []string{"/datamanager/v1/workloads"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetWorkloadInfo",
			Path:    []string{"/datamanager/v1/workload"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetPodAutoscalerList",
			Path:    []string{"/datamanager/v1/podautoscalers"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetPodAutoscaler",
			Path:    []string{"/datamanager/v1/podautoscaler"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetPowerTrading",
			Path:    []string{"/datamanager/v2/powertrading"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetCloudNativeWorkloadList",
			Path:    []string{"/datamanager/v2/cloudnativeworkload"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetUserOperationDataList",
			Path:    []string{"/datamanager/v2/useroperationdata"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetWorkloadRequestResult",
			Path:    []string{"/datamanager/v1/request"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "DataManager.GetWorkloadOriginRequestResult",
			Path:    []string{"/datamanager/v1/workload/request"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for DataManager service

type DataManagerService interface {
	GetAllProjectList(ctx context.Context, in *GetAllProjectListRequest, opts ...client.CallOption) (*GetAllProjectListResponse, error)
	GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, opts ...client.CallOption) (*GetProjectInfoResponse, error)
	GetAllClusterList(ctx context.Context, in *GetClusterListRequest, opts ...client.CallOption) (*GetClusterListResponse, error)
	GetClusterListByProject(ctx context.Context, in *GetClusterListRequest, opts ...client.CallOption) (*GetClusterListResponse, error)
	GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, opts ...client.CallOption) (*GetClusterInfoResponse, error)
	GetNamespaceInfoList(ctx context.Context, in *GetNamespaceInfoListRequest, opts ...client.CallOption) (*GetNamespaceInfoListResponse, error)
	GetNamespaceInfo(ctx context.Context, in *GetNamespaceInfoRequest, opts ...client.CallOption) (*GetNamespaceInfoResponse, error)
	GetWorkloadInfoList(ctx context.Context, in *GetWorkloadInfoListRequest, opts ...client.CallOption) (*GetWorkloadInfoListResponse, error)
	GetWorkloadInfo(ctx context.Context, in *GetWorkloadInfoRequest, opts ...client.CallOption) (*GetWorkloadInfoResponse, error)
	GetPodAutoscalerList(ctx context.Context, in *GetPodAutoscalerListRequest, opts ...client.CallOption) (*GetPodAutoscalerListResponse, error)
	GetPodAutoscaler(ctx context.Context, in *GetPodAutoscalerRequest, opts ...client.CallOption) (*GetPodAutoscalerResponse, error)
	GetPowerTrading(ctx context.Context, in *GetPowerTradingDataRequest, opts ...client.CallOption) (*GetPowerTradingDataResponse, error)
	GetCloudNativeWorkloadList(ctx context.Context, in *GetCloudNativeWorkloadListRequest, opts ...client.CallOption) (*GetCloudNativeWorkloadListResponse, error)
	GetUserOperationDataList(ctx context.Context, in *GetUserOperationDataListRequest, opts ...client.CallOption) (*GetUserOperationDataListResponse, error)
	GetWorkloadRequestResult(ctx context.Context, in *GetWorkloadRequestRecommendResultReq, opts ...client.CallOption) (*GetWorkloadRequestRecommendResultRsp, error)
	GetWorkloadOriginRequestResult(ctx context.Context, in *GetWorkloadOriginRequestResultReq, opts ...client.CallOption) (*GetWorkloadOriginRequestResultRsp, error)
}

type dataManagerService struct {
	c    client.Client
	name string
}

func NewDataManagerService(name string, c client.Client) DataManagerService {
	return &dataManagerService{
		c:    c,
		name: name,
	}
}

func (c *dataManagerService) GetAllProjectList(ctx context.Context, in *GetAllProjectListRequest, opts ...client.CallOption) (*GetAllProjectListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetAllProjectList", in)
	out := new(GetAllProjectListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, opts ...client.CallOption) (*GetProjectInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetProjectInfo", in)
	out := new(GetProjectInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetAllClusterList(ctx context.Context, in *GetClusterListRequest, opts ...client.CallOption) (*GetClusterListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetAllClusterList", in)
	out := new(GetClusterListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetClusterListByProject(ctx context.Context, in *GetClusterListRequest, opts ...client.CallOption) (*GetClusterListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetClusterListByProject", in)
	out := new(GetClusterListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, opts ...client.CallOption) (*GetClusterInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetClusterInfo", in)
	out := new(GetClusterInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetNamespaceInfoList(ctx context.Context, in *GetNamespaceInfoListRequest, opts ...client.CallOption) (*GetNamespaceInfoListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetNamespaceInfoList", in)
	out := new(GetNamespaceInfoListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetNamespaceInfo(ctx context.Context, in *GetNamespaceInfoRequest, opts ...client.CallOption) (*GetNamespaceInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetNamespaceInfo", in)
	out := new(GetNamespaceInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetWorkloadInfoList(ctx context.Context, in *GetWorkloadInfoListRequest, opts ...client.CallOption) (*GetWorkloadInfoListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetWorkloadInfoList", in)
	out := new(GetWorkloadInfoListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetWorkloadInfo(ctx context.Context, in *GetWorkloadInfoRequest, opts ...client.CallOption) (*GetWorkloadInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetWorkloadInfo", in)
	out := new(GetWorkloadInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetPodAutoscalerList(ctx context.Context, in *GetPodAutoscalerListRequest, opts ...client.CallOption) (*GetPodAutoscalerListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetPodAutoscalerList", in)
	out := new(GetPodAutoscalerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetPodAutoscaler(ctx context.Context, in *GetPodAutoscalerRequest, opts ...client.CallOption) (*GetPodAutoscalerResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetPodAutoscaler", in)
	out := new(GetPodAutoscalerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetPowerTrading(ctx context.Context, in *GetPowerTradingDataRequest, opts ...client.CallOption) (*GetPowerTradingDataResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetPowerTrading", in)
	out := new(GetPowerTradingDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetCloudNativeWorkloadList(ctx context.Context, in *GetCloudNativeWorkloadListRequest, opts ...client.CallOption) (*GetCloudNativeWorkloadListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetCloudNativeWorkloadList", in)
	out := new(GetCloudNativeWorkloadListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetUserOperationDataList(ctx context.Context, in *GetUserOperationDataListRequest, opts ...client.CallOption) (*GetUserOperationDataListResponse, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetUserOperationDataList", in)
	out := new(GetUserOperationDataListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetWorkloadRequestResult(ctx context.Context, in *GetWorkloadRequestRecommendResultReq, opts ...client.CallOption) (*GetWorkloadRequestRecommendResultRsp, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetWorkloadRequestResult", in)
	out := new(GetWorkloadRequestRecommendResultRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataManagerService) GetWorkloadOriginRequestResult(ctx context.Context, in *GetWorkloadOriginRequestResultReq, opts ...client.CallOption) (*GetWorkloadOriginRequestResultRsp, error) {
	req := c.c.NewRequest(c.name, "DataManager.GetWorkloadOriginRequestResult", in)
	out := new(GetWorkloadOriginRequestResultRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataManager service

type DataManagerHandler interface {
	GetAllProjectList(context.Context, *GetAllProjectListRequest, *GetAllProjectListResponse) error
	GetProjectInfo(context.Context, *GetProjectInfoRequest, *GetProjectInfoResponse) error
	GetAllClusterList(context.Context, *GetClusterListRequest, *GetClusterListResponse) error
	GetClusterListByProject(context.Context, *GetClusterListRequest, *GetClusterListResponse) error
	GetClusterInfo(context.Context, *GetClusterInfoRequest, *GetClusterInfoResponse) error
	GetNamespaceInfoList(context.Context, *GetNamespaceInfoListRequest, *GetNamespaceInfoListResponse) error
	GetNamespaceInfo(context.Context, *GetNamespaceInfoRequest, *GetNamespaceInfoResponse) error
	GetWorkloadInfoList(context.Context, *GetWorkloadInfoListRequest, *GetWorkloadInfoListResponse) error
	GetWorkloadInfo(context.Context, *GetWorkloadInfoRequest, *GetWorkloadInfoResponse) error
	GetPodAutoscalerList(context.Context, *GetPodAutoscalerListRequest, *GetPodAutoscalerListResponse) error
	GetPodAutoscaler(context.Context, *GetPodAutoscalerRequest, *GetPodAutoscalerResponse) error
	GetPowerTrading(context.Context, *GetPowerTradingDataRequest, *GetPowerTradingDataResponse) error
	GetCloudNativeWorkloadList(context.Context, *GetCloudNativeWorkloadListRequest, *GetCloudNativeWorkloadListResponse) error
	GetUserOperationDataList(context.Context, *GetUserOperationDataListRequest, *GetUserOperationDataListResponse) error
	GetWorkloadRequestResult(context.Context, *GetWorkloadRequestRecommendResultReq, *GetWorkloadRequestRecommendResultRsp) error
	GetWorkloadOriginRequestResult(context.Context, *GetWorkloadOriginRequestResultReq, *GetWorkloadOriginRequestResultRsp) error
}

func RegisterDataManagerHandler(s server.Server, hdlr DataManagerHandler, opts ...server.HandlerOption) error {
	type dataManager interface {
		GetAllProjectList(ctx context.Context, in *GetAllProjectListRequest, out *GetAllProjectListResponse) error
		GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, out *GetProjectInfoResponse) error
		GetAllClusterList(ctx context.Context, in *GetClusterListRequest, out *GetClusterListResponse) error
		GetClusterListByProject(ctx context.Context, in *GetClusterListRequest, out *GetClusterListResponse) error
		GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, out *GetClusterInfoResponse) error
		GetNamespaceInfoList(ctx context.Context, in *GetNamespaceInfoListRequest, out *GetNamespaceInfoListResponse) error
		GetNamespaceInfo(ctx context.Context, in *GetNamespaceInfoRequest, out *GetNamespaceInfoResponse) error
		GetWorkloadInfoList(ctx context.Context, in *GetWorkloadInfoListRequest, out *GetWorkloadInfoListResponse) error
		GetWorkloadInfo(ctx context.Context, in *GetWorkloadInfoRequest, out *GetWorkloadInfoResponse) error
		GetPodAutoscalerList(ctx context.Context, in *GetPodAutoscalerListRequest, out *GetPodAutoscalerListResponse) error
		GetPodAutoscaler(ctx context.Context, in *GetPodAutoscalerRequest, out *GetPodAutoscalerResponse) error
		GetPowerTrading(ctx context.Context, in *GetPowerTradingDataRequest, out *GetPowerTradingDataResponse) error
		GetCloudNativeWorkloadList(ctx context.Context, in *GetCloudNativeWorkloadListRequest, out *GetCloudNativeWorkloadListResponse) error
		GetUserOperationDataList(ctx context.Context, in *GetUserOperationDataListRequest, out *GetUserOperationDataListResponse) error
		GetWorkloadRequestResult(ctx context.Context, in *GetWorkloadRequestRecommendResultReq, out *GetWorkloadRequestRecommendResultRsp) error
		GetWorkloadOriginRequestResult(ctx context.Context, in *GetWorkloadOriginRequestResultReq, out *GetWorkloadOriginRequestResultRsp) error
	}
	type DataManager struct {
		dataManager
	}
	h := &dataManagerHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetAllProjectList",
		Path:    []string{"/datamanager/v1/projects"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetProjectInfo",
		Path:    []string{"/datamanager/v1/project"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetAllClusterList",
		Path:    []string{"/datamanager/v1/clusters"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetClusterListByProject",
		Path:    []string{"/datamanager/v1/projects/clusters"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetClusterInfo",
		Path:    []string{"/datamanager/v1/clusters/{clusterID}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetNamespaceInfoList",
		Path:    []string{"/datamanager/v1/clusters/{clusterID}/namespaces"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetNamespaceInfo",
		Path:    []string{"/datamanager/v1/clusters/{clusterID}/namespaces/{namespace}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetWorkloadInfoList",
		Path:    []string{"/datamanager/v1/workloads"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetWorkloadInfo",
		Path:    []string{"/datamanager/v1/workload"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetPodAutoscalerList",
		Path:    []string{"/datamanager/v1/podautoscalers"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetPodAutoscaler",
		Path:    []string{"/datamanager/v1/podautoscaler"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetPowerTrading",
		Path:    []string{"/datamanager/v2/powertrading"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetCloudNativeWorkloadList",
		Path:    []string{"/datamanager/v2/cloudnativeworkload"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetUserOperationDataList",
		Path:    []string{"/datamanager/v2/useroperationdata"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetWorkloadRequestResult",
		Path:    []string{"/datamanager/v1/request"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "DataManager.GetWorkloadOriginRequestResult",
		Path:    []string{"/datamanager/v1/workload/request"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&DataManager{h}, opts...))
}

type dataManagerHandler struct {
	DataManagerHandler
}

func (h *dataManagerHandler) GetAllProjectList(ctx context.Context, in *GetAllProjectListRequest, out *GetAllProjectListResponse) error {
	return h.DataManagerHandler.GetAllProjectList(ctx, in, out)
}

func (h *dataManagerHandler) GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, out *GetProjectInfoResponse) error {
	return h.DataManagerHandler.GetProjectInfo(ctx, in, out)
}

func (h *dataManagerHandler) GetAllClusterList(ctx context.Context, in *GetClusterListRequest, out *GetClusterListResponse) error {
	return h.DataManagerHandler.GetAllClusterList(ctx, in, out)
}

func (h *dataManagerHandler) GetClusterListByProject(ctx context.Context, in *GetClusterListRequest, out *GetClusterListResponse) error {
	return h.DataManagerHandler.GetClusterListByProject(ctx, in, out)
}

func (h *dataManagerHandler) GetClusterInfo(ctx context.Context, in *GetClusterInfoRequest, out *GetClusterInfoResponse) error {
	return h.DataManagerHandler.GetClusterInfo(ctx, in, out)
}

func (h *dataManagerHandler) GetNamespaceInfoList(ctx context.Context, in *GetNamespaceInfoListRequest, out *GetNamespaceInfoListResponse) error {
	return h.DataManagerHandler.GetNamespaceInfoList(ctx, in, out)
}

func (h *dataManagerHandler) GetNamespaceInfo(ctx context.Context, in *GetNamespaceInfoRequest, out *GetNamespaceInfoResponse) error {
	return h.DataManagerHandler.GetNamespaceInfo(ctx, in, out)
}

func (h *dataManagerHandler) GetWorkloadInfoList(ctx context.Context, in *GetWorkloadInfoListRequest, out *GetWorkloadInfoListResponse) error {
	return h.DataManagerHandler.GetWorkloadInfoList(ctx, in, out)
}

func (h *dataManagerHandler) GetWorkloadInfo(ctx context.Context, in *GetWorkloadInfoRequest, out *GetWorkloadInfoResponse) error {
	return h.DataManagerHandler.GetWorkloadInfo(ctx, in, out)
}

func (h *dataManagerHandler) GetPodAutoscalerList(ctx context.Context, in *GetPodAutoscalerListRequest, out *GetPodAutoscalerListResponse) error {
	return h.DataManagerHandler.GetPodAutoscalerList(ctx, in, out)
}

func (h *dataManagerHandler) GetPodAutoscaler(ctx context.Context, in *GetPodAutoscalerRequest, out *GetPodAutoscalerResponse) error {
	return h.DataManagerHandler.GetPodAutoscaler(ctx, in, out)
}

func (h *dataManagerHandler) GetPowerTrading(ctx context.Context, in *GetPowerTradingDataRequest, out *GetPowerTradingDataResponse) error {
	return h.DataManagerHandler.GetPowerTrading(ctx, in, out)
}

func (h *dataManagerHandler) GetCloudNativeWorkloadList(ctx context.Context, in *GetCloudNativeWorkloadListRequest, out *GetCloudNativeWorkloadListResponse) error {
	return h.DataManagerHandler.GetCloudNativeWorkloadList(ctx, in, out)
}

func (h *dataManagerHandler) GetUserOperationDataList(ctx context.Context, in *GetUserOperationDataListRequest, out *GetUserOperationDataListResponse) error {
	return h.DataManagerHandler.GetUserOperationDataList(ctx, in, out)
}

func (h *dataManagerHandler) GetWorkloadRequestResult(ctx context.Context, in *GetWorkloadRequestRecommendResultReq, out *GetWorkloadRequestRecommendResultRsp) error {
	return h.DataManagerHandler.GetWorkloadRequestResult(ctx, in, out)
}

func (h *dataManagerHandler) GetWorkloadOriginRequestResult(ctx context.Context, in *GetWorkloadOriginRequestResultReq, out *GetWorkloadOriginRequestResultRsp) error {
	return h.DataManagerHandler.GetWorkloadOriginRequestResult(ctx, in, out)
}
