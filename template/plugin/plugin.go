// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/secret"
)

// New returns a new secret plugin.
func New(param1, param2 string) secret.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) Find(ctx context.Context, req *secret.Request) (*drone.Secret, error) {
	// TODO this should be modified or removed. This
	// demonstrates how we can limit access to secrets
	// based on repository meta data. For example,
	// only expose secrets to specific organizations
	// or accounts.
	if req.Repo.Namespace != "octocat" {
		return nil, nil
	}

	// TODO this should be modified or removed. This
	// demonstrates how we could return secrets to
	// pipeline based on name and path.
	if req.Path == "docker" {
		switch req.Name {
		case "username":
			return nil, &drone.Secret{
				Name:        "username",
				Data:        "octocat",
				PullRequest: false, // never expose to pulls requests
			}, nil
		case "password":
			return nil, &drone.Secret{
				Name:        "password",
				Data:        "correct-horse-battery-staple",
				PullRequest: false, // never expose to pulls requests
			}, nil
		}
	}

	return nil, nil
}
