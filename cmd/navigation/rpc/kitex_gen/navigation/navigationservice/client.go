// Code generated by Kitex v0.4.4. DO NOT EDIT.

package navigationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	navigation "github.com/quarkcloudio/quark-micro/cmd/navigation/rpc/kitex_gen/navigation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetNavigationList(ctx context.Context, req *navigation.NavigationListRequest, callOptions ...callopt.Option) (r *navigation.NavigationListResponse, err error)
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
	return &kNavigationServiceClient{
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

type kNavigationServiceClient struct {
	*kClient
}

func (p *kNavigationServiceClient) GetNavigationList(ctx context.Context, req *navigation.NavigationListRequest, callOptions ...callopt.Option) (r *navigation.NavigationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNavigationList(ctx, req)
}
