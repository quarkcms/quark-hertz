// Code generated by Kitex v0.4.4. DO NOT EDIT.

package postservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	post "github.com/quarkcloudio/quark-micro/cmd/content/rpc/kitex_gen/post"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetPage(ctx context.Context, req *post.PageRequest, callOptions ...callopt.Option) (r *post.PageResponse, err error)
	GetArticleDetail(ctx context.Context, req *post.ArticleDetailRequest, callOptions ...callopt.Option) (r *post.ArticleDetailResponse, err error)
	GetArticleList(ctx context.Context, req *post.ArticleListRequest, callOptions ...callopt.Option) (r *post.ArticleListResponse, err error)
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
	return &kPostServiceClient{
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

type kPostServiceClient struct {
	*kClient
}

func (p *kPostServiceClient) GetPage(ctx context.Context, req *post.PageRequest, callOptions ...callopt.Option) (r *post.PageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPage(ctx, req)
}

func (p *kPostServiceClient) GetArticleDetail(ctx context.Context, req *post.ArticleDetailRequest, callOptions ...callopt.Option) (r *post.ArticleDetailResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetArticleDetail(ctx, req)
}

func (p *kPostServiceClient) GetArticleList(ctx context.Context, req *post.ArticleListRequest, callOptions ...callopt.Option) (r *post.ArticleListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetArticleList(ctx, req)
}
