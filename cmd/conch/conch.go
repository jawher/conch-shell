// Copyright 2017 Joyent, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/briandowns/spinner"
	"github.com/joyent/conch-shell/pkg/commands"
	"github.com/joyent/conch-shell/pkg/config"
	"github.com/joyent/conch-shell/pkg/util"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/jawher/mow.cli.v1"
	"os"
	"strconv"
	"time"
)

// These variables are provided by the build environment
var (
	Version   string
	BuildTime string
	GitRev    string
)

func main() {
	util.UserAgent = fmt.Sprintf("conch shell v%s-%s", Version, GitRev)
	app := cli.App("conch", "Command line interface for Conch")
	app.Version("version", Version)

	app.Command(
		"version",
		"Get more detailed version info than --version",
		func(cmd *cli.Cmd) {
			buildTime := BuildTime
			t, err := strconv.ParseInt(BuildTime, 10, 64)
			if err == nil {
				buildTime = util.TimeStr(time.Unix(t, 0))
			}

			cmd.Action = func() {
				fmt.Printf(
					"Conch Shell v%s\n"+
						"  Git Revision: %s\n"+
						"  Build Time: %s\n",
					Version,
					GitRev,
					buildTime,
				)
			}
		},
	)

	var (
		useJSON    = app.BoolOpt("json j", false, "Output JSON")
		configFile = app.StringOpt("config c", "~/.conch.json", "Path to config file")
		pretty     = app.BoolOpt("pretty", false, "Pretty CLI output, including spinners")
		noVersion  = app.BoolOpt("no-version-check", false, "Skip Github version check")
	)

	app.Before = func() {
		if *useJSON {
			util.JSON = true
		} else {
			util.JSON = false
		}

		if !*noVersion {
			gh, err := util.LatestGithubRelease("joyent", "conch-shell")
			if err != nil {
				util.Bail(err)
			}
			sem := semver.MustParse(Version)
			if gh.SemVer.GT(sem) {
				fmt.Printf(
					"** A new release is available! You have v%s and %s is available.\n",
					Version,
					gh.TagName,
				)
				fmt.Printf("** Download the new release at %s\n\n", gh.URL)
			}
		}

		util.Pretty = *pretty
		if *pretty {
			util.Spin = spinner.New(spinner.CharSets[10], 100*time.Millisecond)
			util.Spin.FinalMSG = "Complete.\n"
		}

		expandedPath, err := homedir.Expand(*configFile)
		if err != nil {
			util.Bail(err)
		}

		cfg, err := config.NewFromJSONFile(expandedPath)
		if err != nil {
			cfg.Path = expandedPath
		}
		util.Config = cfg

		for _, prof := range cfg.Profiles {
			if prof.Active {
				util.ActiveProfile = prof
				break
			}
		}
	}

	commands.Init(app)

	_ = app.Run(os.Args)
}
