// Code generated by Kitex v0.6.1. DO NOT EDIT.

package cloudfileservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Add(ctx context.Context, req *rpc.AddRequest, callOptions ...callopt.Option) (err error)
	CreateDirectory(ctx context.Context, req *rpc.CreateDirectoryRequest, callOptions ...callopt.Option) (err error)
	Remove(ctx context.Context, req *rpc.RemoveRequest, callOptions ...callopt.Option) (err error)
	RemoveDirectory(ctx context.Context, id string, callOptions ...callopt.Option) (err error)
	Query(ctx context.Context, req *rpc.QueryRequest, callOptions ...callopt.Option) (r *rpc.QueryResponse, err error)
	Update(ctx context.Context, req *rpc.UpdateRequest, callOptions ...callopt.Option) (err error)
	Rename(ctx context.Context, id string, newName_ string, callOptions ...callopt.Option) (err error)
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
	return &kCloudFileServiceClient{
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

type kCloudFileServiceClient struct {
	*kClient
}

func (p *kCloudFileServiceClient) Add(ctx context.Context, req *rpc.AddRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Add(ctx, req)
}

func (p *kCloudFileServiceClient) CreateDirectory(ctx context.Context, req *rpc.CreateDirectoryRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateDirectory(ctx, req)
}

func (p *kCloudFileServiceClient) Remove(ctx context.Context, req *rpc.RemoveRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Remove(ctx, req)
}

func (p *kCloudFileServiceClient) RemoveDirectory(ctx context.Context, id string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveDirectory(ctx, id)
}

func (p *kCloudFileServiceClient) Query(ctx context.Context, req *rpc.QueryRequest, callOptions ...callopt.Option) (r *rpc.QueryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Query(ctx, req)
}

func (p *kCloudFileServiceClient) Update(ctx context.Context, req *rpc.UpdateRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Update(ctx, req)
}

func (p *kCloudFileServiceClient) Rename(ctx context.Context, id string, newName_ string, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Rename(ctx, id, newName_)
}
