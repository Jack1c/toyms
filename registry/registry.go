package registry

import "context"

type Register interface {
	Register(ctx context.Context, instance *ServiceInstance)
	DeRegister(ctx context.Context, instance *ServiceInstance)
}

type Discover interface {
	GetService(ctx context.Context, serverName string) []*ServiceInstance
	Watch(ctx context.Context, serverName string) []*ServiceInstance
}
