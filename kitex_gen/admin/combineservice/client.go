// Code generated by Kitex v0.4.4. DO NOT EDIT.

package combineservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	admin "github.com/quarkcms/quark-hertz/kitex_gen/admin"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DashboardHandle(ctx context.Context, req *admin.DashboardRequest, callOptions ...callopt.Option) (r *admin.DashboardResponse, err error)
	ResourceIndexhandle(ctx context.Context, req *admin.ResourceIndexRequest, callOptions ...callopt.Option) (r *admin.ResourceIndexResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCombineServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCombineServiceClient struct {
	*kClient
}

func (p *kCombineServiceClient) DashboardHandle(ctx context.Context, req *admin.DashboardRequest, callOptions ...callopt.Option) (r *admin.DashboardResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DashboardHandle(ctx, req)
}

func (p *kCombineServiceClient) ResourceIndexhandle(ctx context.Context, req *admin.ResourceIndexRequest, callOptions ...callopt.Option) (r *admin.ResourceIndexResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ResourceIndexhandle(ctx, req)
}
