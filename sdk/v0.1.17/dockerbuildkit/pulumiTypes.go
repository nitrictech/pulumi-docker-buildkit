// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dockerbuildkit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Docker build argument.
type BuildArg struct {
	// The name of the Docker build argument.
	Name string `pulumi:"name"`
	// The value of the Docker build argument.
	Value string `pulumi:"value"`
}

// BuildArgInput is an input type that accepts BuildArgArgs and BuildArgOutput values.
// You can construct a concrete instance of `BuildArgInput` via:
//
//          BuildArgArgs{...}
type BuildArgInput interface {
	pulumi.Input

	ToBuildArgOutput() BuildArgOutput
	ToBuildArgOutputWithContext(context.Context) BuildArgOutput
}

// Describes a Docker build argument.
type BuildArgArgs struct {
	// The name of the Docker build argument.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the Docker build argument.
	Value pulumi.StringInput `pulumi:"value"`
}

func (BuildArgArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildArg)(nil)).Elem()
}

func (i BuildArgArgs) ToBuildArgOutput() BuildArgOutput {
	return i.ToBuildArgOutputWithContext(context.Background())
}

func (i BuildArgArgs) ToBuildArgOutputWithContext(ctx context.Context) BuildArgOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildArgOutput)
}

// BuildArgArrayInput is an input type that accepts BuildArgArray and BuildArgArrayOutput values.
// You can construct a concrete instance of `BuildArgArrayInput` via:
//
//          BuildArgArray{ BuildArgArgs{...} }
type BuildArgArrayInput interface {
	pulumi.Input

	ToBuildArgArrayOutput() BuildArgArrayOutput
	ToBuildArgArrayOutputWithContext(context.Context) BuildArgArrayOutput
}

type BuildArgArray []BuildArgInput

func (BuildArgArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildArg)(nil)).Elem()
}

func (i BuildArgArray) ToBuildArgArrayOutput() BuildArgArrayOutput {
	return i.ToBuildArgArrayOutputWithContext(context.Background())
}

func (i BuildArgArray) ToBuildArgArrayOutputWithContext(ctx context.Context) BuildArgArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildArgArrayOutput)
}

// Describes a Docker build argument.
type BuildArgOutput struct{ *pulumi.OutputState }

func (BuildArgOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildArg)(nil)).Elem()
}

func (o BuildArgOutput) ToBuildArgOutput() BuildArgOutput {
	return o
}

func (o BuildArgOutput) ToBuildArgOutputWithContext(ctx context.Context) BuildArgOutput {
	return o
}

// The name of the Docker build argument.
func (o BuildArgOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BuildArg) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the Docker build argument.
func (o BuildArgOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v BuildArg) string { return v.Value }).(pulumi.StringOutput)
}

type BuildArgArrayOutput struct{ *pulumi.OutputState }

func (BuildArgArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BuildArg)(nil)).Elem()
}

func (o BuildArgArrayOutput) ToBuildArgArrayOutput() BuildArgArrayOutput {
	return o
}

func (o BuildArgArrayOutput) ToBuildArgArrayOutputWithContext(ctx context.Context) BuildArgArrayOutput {
	return o
}

func (o BuildArgArrayOutput) Index(i pulumi.IntInput) BuildArgOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BuildArg {
		return vs[0].([]BuildArg)[vs[1].(int)]
	}).(BuildArgOutput)
}

// Describes a Docker container registry.
type Registry struct {
	// The password to authenticate with.
	Password *string `pulumi:"password"`
	// The URL of the Docker registry server.
	Server string `pulumi:"server"`
	// The username to authenticate with.
	Username *string `pulumi:"username"`
}

// RegistryInput is an input type that accepts RegistryArgs and RegistryOutput values.
// You can construct a concrete instance of `RegistryInput` via:
//
//          RegistryArgs{...}
type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(context.Context) RegistryOutput
}

// Describes a Docker container registry.
type RegistryArgs struct {
	// The password to authenticate with.
	Password pulumi.StringPtrInput `pulumi:"password"`
	// The URL of the Docker registry server.
	Server pulumi.StringInput `pulumi:"server"`
	// The username to authenticate with.
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil)).Elem()
}

func (i RegistryArgs) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i RegistryArgs) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

// Describes a Docker container registry.
type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

// The password to authenticate with.
func (o RegistryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Registry) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// The URL of the Docker registry server.
func (o RegistryOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v Registry) string { return v.Server }).(pulumi.StringOutput)
}

// The username to authenticate with.
func (o RegistryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Registry) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BuildArgInput)(nil)).Elem(), BuildArgArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuildArgArrayInput)(nil)).Elem(), BuildArgArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryInput)(nil)).Elem(), RegistryArgs{})
	pulumi.RegisterOutputType(BuildArgOutput{})
	pulumi.RegisterOutputType(BuildArgArrayOutput{})
	pulumi.RegisterOutputType(RegistryOutput{})
}
