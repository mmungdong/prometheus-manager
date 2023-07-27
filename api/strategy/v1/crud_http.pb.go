// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.4
// source: strategy/v1/crud.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCrudCreateCrud = "/api.strategy.v1.Crud/CreateCrud"
const OperationCrudDeleteCrud = "/api.strategy.v1.Crud/DeleteCrud"
const OperationCrudGetCrud = "/api.strategy.v1.Crud/GetCrud"
const OperationCrudUpdateCrud = "/api.strategy.v1.Crud/UpdateCrud"

type CrudHTTPServer interface {
	CreateCrud(context.Context, *CreateCrudRequest) (*CreateCrudReply, error)
	DeleteCrud(context.Context, *DeleteCrudRequest) (*DeleteCrudReply, error)
	GetCrud(context.Context, *GetCrudRequest) (*GetCrudReply, error)
	UpdateCrud(context.Context, *UpdateCrudRequest) (*UpdateCrudReply, error)
}

func RegisterCrudHTTPServer(s *http.Server, srv CrudHTTPServer) {
	r := s.Route("/")
	r.POST("/op/v1/add", _Crud_CreateCrud0_HTTP_Handler(srv))
	r.PUT("/op/v1/edit/{id}", _Crud_UpdateCrud0_HTTP_Handler(srv))
	r.DELETE("/op/v1/delete/{id}", _Crud_DeleteCrud0_HTTP_Handler(srv))
	r.GET("/op/v1/get/{id}", _Crud_GetCrud0_HTTP_Handler(srv))
}

func _Crud_CreateCrud0_HTTP_Handler(srv CrudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCrudRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCrudCreateCrud)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCrud(ctx, req.(*CreateCrudRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCrudReply)
		return ctx.Result(200, reply)
	}
}

func _Crud_UpdateCrud0_HTTP_Handler(srv CrudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCrudRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCrudUpdateCrud)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCrud(ctx, req.(*UpdateCrudRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCrudReply)
		return ctx.Result(200, reply)
	}
}

func _Crud_DeleteCrud0_HTTP_Handler(srv CrudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCrudRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCrudDeleteCrud)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCrud(ctx, req.(*DeleteCrudRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCrudReply)
		return ctx.Result(200, reply)
	}
}

func _Crud_GetCrud0_HTTP_Handler(srv CrudHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCrudRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCrudGetCrud)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCrud(ctx, req.(*GetCrudRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCrudReply)
		return ctx.Result(200, reply)
	}
}

type CrudHTTPClient interface {
	CreateCrud(ctx context.Context, req *CreateCrudRequest, opts ...http.CallOption) (rsp *CreateCrudReply, err error)
	DeleteCrud(ctx context.Context, req *DeleteCrudRequest, opts ...http.CallOption) (rsp *DeleteCrudReply, err error)
	GetCrud(ctx context.Context, req *GetCrudRequest, opts ...http.CallOption) (rsp *GetCrudReply, err error)
	UpdateCrud(ctx context.Context, req *UpdateCrudRequest, opts ...http.CallOption) (rsp *UpdateCrudReply, err error)
}

type CrudHTTPClientImpl struct {
	cc *http.Client
}

func NewCrudHTTPClient(client *http.Client) CrudHTTPClient {
	return &CrudHTTPClientImpl{client}
}

func (c *CrudHTTPClientImpl) CreateCrud(ctx context.Context, in *CreateCrudRequest, opts ...http.CallOption) (*CreateCrudReply, error) {
	var out CreateCrudReply
	pattern := "/op/v1/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCrudCreateCrud))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CrudHTTPClientImpl) DeleteCrud(ctx context.Context, in *DeleteCrudRequest, opts ...http.CallOption) (*DeleteCrudReply, error) {
	var out DeleteCrudReply
	pattern := "/op/v1/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCrudDeleteCrud))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CrudHTTPClientImpl) GetCrud(ctx context.Context, in *GetCrudRequest, opts ...http.CallOption) (*GetCrudReply, error) {
	var out GetCrudReply
	pattern := "/op/v1/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCrudGetCrud))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CrudHTTPClientImpl) UpdateCrud(ctx context.Context, in *UpdateCrudRequest, opts ...http.CallOption) (*UpdateCrudReply, error) {
	var out UpdateCrudReply
	pattern := "/op/v1/edit/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCrudUpdateCrud))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
