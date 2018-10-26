package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/concourse/atc"
)

// CheckResponse contains a default space and a list of resource versions. This response is returned from the check of the resource v2 interface. The default space can be empty for resources that do not have a default space (ex. PR resource).
type CheckResponse struct {
	DefaultSpace string `json:"default_space"`
	Versions     []ResourceVersion
}

type ResourceVersion struct {
	Space    string       `json:"space"`
	Version  atc.Version  `json:"version"`
	Metadata atc.Metadata `json:"metadata"`
}

type checkRequest struct {
	Config       map[string]interface{}    `json:"config"`
	From         map[atc.Space]atc.Version `json:"from"`
	ResponsePath string                    `json:"response_path"`
}

func (r *resource) Check(
	ctx context.Context,
	src atc.Source,
	from map[atc.Space]atc.Version,
) error {
	var spaces atc.Spaces

	tmpfile, err := ioutil.TempFile("", "response")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name())

	path := r.info.Artifacts.Check
	input := checkRequest{src, from, tmpfile.Name()}

	request, err := json.Marshal(input)
	if err != nil {
		return err
	}

	stderr := new(bytes.Buffer)

	processIO := garden.ProcessIO{
		Stdin:  bytes.NewBuffer(request),
		Stderr: stderr,
	}

	process, err := r.container.Run(garden.ProcessSpec{
		Path: path,
	}, processIO)
	if err != nil {
		return err
	}

	processExited := make(chan struct{})

	var processStatus int
	var processErr error

	go func() {
		processStatus, processErr = process.Wait()
		close(processExited)
	}()

	select {
	case <-processExited:
		if processErr != nil {
			return processErr
		}

		if processStatus != 0 {
			return ErrResourceScriptFailed{
				Path:       path,
				ExitStatus: processStatus,

				Stderr: stderr.String(),
			}
		}

	case <-ctx.Done():
		r.container.Stop(false)
		<-processExited
		return ctx.Err()
	}

	fileReader, err := os.Open(tmpfile.Name())
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(fileReader)

	if decoder.More() {
		var defaultSpace atc.DefaultSpaceResponse
		err := decoder.Decode(&defaultSpace)
		if err != nil {
			return err
		}

		err = r.resourceConfig.SaveDefaultSpace(defaultSpace.DefaultSpace)
		if err != nil {
			return err
		}
	}

	for decoder.More() {
		var version atc.SpaceVersion
		err := decoder.Decode(&version)
		if err != nil {
			return err
		}

		if len(spaces.AllSpaces) == 0 || spaces.AllSpaces[len(spaces.AllSpaces)-1] != version.Space {
			spaces.AllSpaces = append(spaces.AllSpaces, version.Space)
		}

		err = r.resourceConfig.SaveVersion(version)
		if err != nil {
			return err
		}
	}

	if len(spaces.AllSpaces) != 0 {
		err = r.resourceConfig.SaveSpaces(spaces.AllSpaces)
		if err != nil {
			return err
		}
	}

	return nil
}
