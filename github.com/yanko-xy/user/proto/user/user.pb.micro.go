// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	// 对外提供添加服务
	AddUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*Response, error)
	FindUserByID(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfo, error)
	FindAllUser(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllUser, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AddUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.AddUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *UserId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) FindUserByID(ctx context.Context, in *UserId, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "User.FindUserByID", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) FindAllUser(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllUser, error) {
	req := c.c.NewRequest(c.name, "User.FindAllUser", in)
	out := new(AllUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 对外提供添加服务
	AddUser(context.Context, *UserInfo, *Response) error
	DeleteUser(context.Context, *UserId, *Response) error
	UpdateUser(context.Context, *UserInfo, *Response) error
	FindUserByID(context.Context, *UserId, *UserInfo) error
	FindAllUser(context.Context, *FindAll, *AllUser) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		AddUser(ctx context.Context, in *UserInfo, out *Response) error
		DeleteUser(ctx context.Context, in *UserId, out *Response) error
		UpdateUser(ctx context.Context, in *UserInfo, out *Response) error
		FindUserByID(ctx context.Context, in *UserId, out *UserInfo) error
		FindAllUser(ctx context.Context, in *FindAll, out *AllUser) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) AddUser(ctx context.Context, in *UserInfo, out *Response) error {
	return h.UserHandler.AddUser(ctx, in, out)
}

func (h *userHandler) DeleteUser(ctx context.Context, in *UserId, out *Response) error {
	return h.UserHandler.DeleteUser(ctx, in, out)
}

func (h *userHandler) UpdateUser(ctx context.Context, in *UserInfo, out *Response) error {
	return h.UserHandler.UpdateUser(ctx, in, out)
}

func (h *userHandler) FindUserByID(ctx context.Context, in *UserId, out *UserInfo) error {
	return h.UserHandler.FindUserByID(ctx, in, out)
}

func (h *userHandler) FindAllUser(ctx context.Context, in *FindAll, out *AllUser) error {
	return h.UserHandler.FindAllUser(ctx, in, out)
}
