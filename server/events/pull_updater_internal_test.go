// Copyright 2026 The Atlantis Authors
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"testing"

	"github.com/runatlantis/atlantis/server/events/command"
	. "github.com/runatlantis/atlantis/testing"
)

func TestHidePrevCommentsDir(t *testing.T) {
	one := command.Result{ProjectResults: []command.ProjectResult{{RepoRelDir: "envs/a", ProjectName: "a"}}}
	two := command.Result{ProjectResults: []command.ProjectResult{
		{RepoRelDir: "envs/a", ProjectName: "a"},
		{RepoRelDir: "envs/b", ProjectName: "b"},
	}}

	cases := []struct {
		description string
		cmd         PullCommand
		res         command.Result
		exp         string
	}{
		{
			description: "dir flag is used as is",
			cmd:         &CommentCommand{Name: command.Plan, RepoRelDir: "envs/b"},
			res:         one,
			exp:         "envs/b",
		},
		{
			description: "project flag uses the dir of the project that ran",
			cmd:         &CommentCommand{Name: command.Plan, ProjectName: "a"},
			res:         one,
			exp:         "envs/a",
		},
		{
			description: "project flag on a non-pointer command",
			cmd:         CommentCommand{Name: command.Plan, ProjectName: "a"},
			res:         one,
			exp:         "envs/a",
		},
		{
			description: "project flag matching several projects is not limited",
			cmd:         &CommentCommand{Name: command.Plan, ProjectName: ".*"},
			res:         two,
			exp:         "",
		},
		{
			description: "project flag with no project result is not limited",
			cmd:         &CommentCommand{Name: command.Plan, ProjectName: "a"},
			res:         command.Result{},
			exp:         "",
		},
		{
			description: "plan without flags is not limited",
			cmd:         &CommentCommand{Name: command.Plan},
			res:         one,
			exp:         "",
		},
		{
			description: "autoplan is not limited",
			cmd:         AutoplanCommand{},
			res:         one,
			exp:         "",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			Equals(t, c.exp, hidePrevCommentsDir(c.cmd, c.res))
		})
	}
}
