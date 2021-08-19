package builder

import (
	"context"
	"fmt"

	"github.com/ansemjo/openwrt-imagebuilder/config"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
)

type Config struct {
	Name string `yaml:"name"`
}

func Build(ctx context.Context, c client.Client) (*client.Result, error) {

	cfg, err := GetConfigFile(ctx, c)
	if err != nil {
		return nil, errors.Wrapf(err, "failed reading config")
	}

	conf, err := config.FromBytes(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal config")
	}

	st := ConfigToLLB(*conf)

	def, err := st.Marshal()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal local source: %s", err)
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve dockerfile: %s", err)
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	res.SetRef(ref)
	return res, nil

}

func GetConfigFile(ctx context.Context, c client.Client) ([]byte, error) {
	opts := c.BuildOpts().Opts
	filename := opts["filename"]
	if filename == "" {
		filename = "owrtfile"
	}

	src := llb.Local("dockerfile",
		llb.IncludePatterns([]string{filename}),
		llb.SessionID(c.BuildOpts().SessionID),
		llb.SharedKeyHint("owrtfile"),
		llb.WithCustomName("[owrtbuilder] loading configuration from "+filename),
	)

	def, err := src.Marshal()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}

	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	dtFile, err := ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file")
	}

	return dtFile, nil

}
