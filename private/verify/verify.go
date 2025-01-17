// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	internalcompile "github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage/disk"
	"github.com/cerbos/cerbos/internal/storage/hub"
	"github.com/cerbos/cerbos/internal/storage/index"
	"github.com/cerbos/cerbos/internal/verify"
	"github.com/cerbos/cerbos/private/compile"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"go.uber.org/zap"
)

// Files runs tests using the policy files in the given file system.
func Files(ctx context.Context, fsys fs.FS, idx compile.Index) (*policyv1.TestResults, error) {
	if idx == nil {
		var err error
		idx, err = index.Build(ctx, fsys, index.WithBuildFailureLogLevel(zap.DebugLevel))
		if err != nil {
			idxErrs := new(index.BuildError)
			if errors.As(err, &idxErrs) {
				return nil, &compile.Errors{
					Errors: &runtimev1.Errors{
						Kind: &runtimev1.Errors_IndexBuildErrors{IndexBuildErrors: idxErrs.IndexBuildErrors},
					},
				}
			}

			return nil, fmt.Errorf("failed to build index: %w", err)
		}
	}

	store := disk.NewFromIndexWithConf(idx, &disk.Conf{})
	schemaMgr := schema.NewFromConf(ctx, store, schema.NewConf(schema.EnforcementReject))
	compiler := internalcompile.NewManagerFromDefaultConf(ctx, store, schemaMgr)
	eng := engine.NewEphemeral(nil, compiler, schemaMgr)

	results, err := verify.Verify(ctx, fsys, eng, verify.Config{Trace: true})
	if err != nil {
		return nil, fmt.Errorf("failed to run tests: %w", err)
	}

	return results, nil
}

type BundleParams struct {
	BundlePath string
	TestsDir   string
	WorkDir    string
}

// Bundle runs tests using the given policy bundle.
func Bundle(ctx context.Context, params BundleParams) (*policyv1.TestResults, error) {
	bundleSrc, err := hub.NewLocalSource(hub.LocalParams{
		BundlePath: params.BundlePath,
		TempDir:    params.WorkDir,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create local bundle source from %q: %w", params.BundlePath, err)
	}

	schemaMgr := schema.NewFromConf(ctx, bundleSrc, schema.NewConf(schema.EnforcementReject))
	eng := engine.NewEphemeral(nil, bundleSrc, schemaMgr)

	results, err := verify.Verify(ctx, os.DirFS(params.TestsDir), eng, verify.Config{Trace: true})
	if err != nil {
		return nil, fmt.Errorf("failed to run tests: %w", err)
	}

	return results, nil
}

type CheckOptions interface {
	Globals() map[string]any
	NowFunc() func() time.Time
	DefaultPolicyVersion() string
	LenientScopeSearch() bool
}

type Checker interface {
	Check(ctx context.Context, inputs []*enginev1.CheckInput, opts CheckOptions) ([]*enginev1.CheckOutput, error)
}

type Opt func(config *verify.Config)

func WithExcludedResourcePolicyFQNs(fqns ...string) Opt {
	return func(config *verify.Config) {
		if config.ExcludedResourcePolicyFQNs == nil {
			config.ExcludedResourcePolicyFQNs = make(map[string]struct{}, len(fqns))
		}

		for _, fqn := range fqns {
			config.ExcludedResourcePolicyFQNs[fqn] = struct{}{}
		}
	}
}

func WithExcludedPrincipalPolicyFQNs(fqns ...string) Opt {
	return func(config *verify.Config) {
		if config.ExcludedPrincipalPolicyFQNs == nil {
			config.ExcludedPrincipalPolicyFQNs = make(map[string]struct{}, len(fqns))
		}

		for _, fqn := range fqns {
			config.ExcludedPrincipalPolicyFQNs[fqn] = struct{}{}
		}
	}
}

func WithCustomChecker(ctx context.Context, fsys fs.FS, eng Checker, opts ...Opt) (*policyv1.TestResults, error) {
	config := new(verify.Config)
	for _, opt := range opts {
		opt(config)
	}
	results, err := verify.Verify(ctx, fsys, checkFunc(eng.Check), *config)
	if err != nil {
		return nil, fmt.Errorf("failed to run tests: %w", err)
	}

	return results, nil
}

type checkFunc func(context.Context, []*enginev1.CheckInput, CheckOptions) ([]*enginev1.CheckOutput, error)

func (f checkFunc) Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...engine.CheckOpt) ([]*enginev1.CheckOutput, error) {
	return f(ctx, inputs, engine.ApplyCheckOptions(opts...))
}
