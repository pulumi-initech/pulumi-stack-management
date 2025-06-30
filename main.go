package main

import (
	"context"
	"fmt"
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	goGen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
)

func main() {
	provider, err := provider()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	err = provider.Run(context.Background(), "stack-management", "0.0.9")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() (p.Provider, error) {
	return infer.NewProviderBuilder().
		WithNamespace("pulumi-initech").
		WithComponents(
			infer.ComponentF(NewStackManagement),
		).
		WithConfig(infer.Config(Config{})).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"stack-management": "index",
		}).
		WithLanguageMap(map[string]any{
			"go": goGen.GoPackageInfo{
				ImportBasePath: "github.com/pulumi-initech/pulumi-stack-management/sdk/go/stackmanagement",
			},
		}).
		WithPluginDownloadURL("https://github.com/pulumi-initech/pulumi-stack-management/releases/download/v0.0.9/").
		Build()
}

type Config struct {
	Scream *bool `pulumi:"scream,optional"`
}

type StackManagementArgs struct {
}

type StackManagement struct {
	pulumi.ResourceState
	StackManagementArgs
}

func NewStackManagement(ctx *pulumi.Context, name string, args *StackManagementArgs, opts ...pulumi.ResourceOption) (*StackManagement, error) {
	comp := &StackManagement{}
	err := ctx.RegisterComponentResource(p.GetTypeToken(ctx), name, comp, opts...)
	if err != nil {
		return nil, err
	}

	return comp, nil
}
