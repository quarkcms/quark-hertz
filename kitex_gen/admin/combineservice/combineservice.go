// Code generated by Kitex v0.4.4. DO NOT EDIT.

package combineservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	admin "github.com/quarkcms/quark-hertz/kitex_gen/admin"
)

type CombineService interface {
	admin.Dashboard
	admin.ResourceIndex
}

func serviceInfo() *kitex.ServiceInfo {
	return combineServiceServiceInfo
}

var combineServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CombineService"
	handlerType := (*CombineService)(nil)
	methods := map[string]kitex.MethodInfo{
		"dashboardHandle":     kitex.NewMethodInfo(dashboardHandleHandler, newDashboardDashboardHandleArgs, newDashboardDashboardHandleResult, false),
		"resourceIndexhandle": kitex.NewMethodInfo(resourceIndexhandleHandler, newResourceIndexResourceIndexhandleArgs, newResourceIndexResourceIndexhandleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "admin",
	}
	extra["combine_service"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func dashboardHandleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*admin.DashboardDashboardHandleArgs)
	realResult := result.(*admin.DashboardDashboardHandleResult)
	success, err := handler.(admin.Dashboard).DashboardHandle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDashboardDashboardHandleArgs() interface{} {
	return admin.NewDashboardDashboardHandleArgs()
}

func newDashboardDashboardHandleResult() interface{} {
	return admin.NewDashboardDashboardHandleResult()
}

func resourceIndexhandleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*admin.ResourceIndexResourceIndexhandleArgs)
	realResult := result.(*admin.ResourceIndexResourceIndexhandleResult)
	success, err := handler.(admin.ResourceIndex).ResourceIndexhandle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newResourceIndexResourceIndexhandleArgs() interface{} {
	return admin.NewResourceIndexResourceIndexhandleArgs()
}

func newResourceIndexResourceIndexhandleResult() interface{} {
	return admin.NewResourceIndexResourceIndexhandleResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DashboardHandle(ctx context.Context, req *admin.DashboardRequest) (r *admin.DashboardResponse, err error) {
	var _args admin.DashboardDashboardHandleArgs
	_args.Req = req
	var _result admin.DashboardDashboardHandleResult
	if err = p.c.Call(ctx, "dashboardHandle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ResourceIndexhandle(ctx context.Context, req *admin.ResourceIndexRequest) (r *admin.ResourceIndexResponse, err error) {
	var _args admin.ResourceIndexResourceIndexhandleArgs
	_args.Req = req
	var _result admin.ResourceIndexResourceIndexhandleResult
	if err = p.c.Call(ctx, "resourceIndexhandle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
