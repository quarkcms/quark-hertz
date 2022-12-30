// Code generated by Kitex v0.4.4. DO NOT EDIT.

package resourceindex

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	admin "github.com/quarkcms/quark-hertz/cmd/admin/rpc/kitex_gen/admin"
)

func serviceInfo() *kitex.ServiceInfo {
	return resourceIndexServiceInfo
}

var resourceIndexServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ResourceIndex"
	handlerType := (*admin.ResourceIndex)(nil)
	methods := map[string]kitex.MethodInfo{
		"resourceIndexhandle": kitex.NewMethodInfo(resourceIndexhandleHandler, newResourceIndexResourceIndexhandleArgs, newResourceIndexResourceIndexhandleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "admin",
	}
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

func (p *kClient) ResourceIndexhandle(ctx context.Context, req *admin.ResourceIndexRequest) (r *admin.ResourceIndexResponse, err error) {
	var _args admin.ResourceIndexResourceIndexhandleArgs
	_args.Req = req
	var _result admin.ResourceIndexResourceIndexhandleResult
	if err = p.c.Call(ctx, "resourceIndexhandle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}