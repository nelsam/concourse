package v2

import (
	"context"

	"github.com/concourse/concourse/atc"
)

type putRequest struct {
	Source atc.Source `json:"source"`
	Params atc.Params `json:"params,omitempty"`
}

func (r *resource) Put(
	ctx context.Context,
	ioConfig atc.IOConfig,
	source atc.Source,
	params atc.Params,
) ([]atc.SpaceVersion, error) {
	return []atc.SpaceVersion{}, nil
	// resourceDir := atc.ResourcesDir("put")

	// err := RunScript(
	// 	ctx,
	// 	"/opt/resource/out",
	// 	[]string{resourceDir},
	// 	putRequest{
	// 		Params: params,
	// 		Source: source,
	// 	},
	// 	&versionResult,
	// 	ioConfig.Stderr,
	// 	true,
	// 	r.container,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// return res.NewPutVersionedSource(versionResult, r.container, resourceDir), nil
}
