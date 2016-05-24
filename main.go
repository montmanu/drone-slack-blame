package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "slack-blame"
	app.Usage = "slack-blame plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "slack access token",
			EnvVar: "PLUGIN_TOKEN,SLACK_TOKEN",
		},
		cli.StringFlag{
			Name:   "channel",
			Usage:  "slack channel",
			EnvVar: "PLUGIN_CHANNEL",
		},
		cli.StringFlag{
			Name:   "success.username",
			Usage:  "username for successful builds",
			Value:  "drone",
			EnvVar: "PLUGIN_SUCCESS_USERNAME",
		},
		cli.StringFlag{
			Name:   "success.icon",
			Usage:  "icon for successful builds",
			Value:  ":drone:",
			EnvVar: "PLUGIN_SUCCESS_ICON",
		},
		cli.StringFlag{
			Name:   "success.template",
			Usage:  "template for successful builds",
			EnvVar: "PLUGIN_SUCCESS_TEMPLATE",
		},
		cli.StringSliceFlag{
			Name:   "success.image_attachments",
			Usage:  "image attachments for successful builds",
			EnvVar: "PLUGIN_SUCCESS_IMAGE_ATTACHMENTS",
		},
		cli.StringFlag{
			Name:   "failure.username",
			Usage:  "username for failed builds",
			Value:  "drone",
			EnvVar: "PLUGIN_FAILURE_USERNAME",
		},
		cli.StringFlag{
			Name:   "failure.icon",
			Usage:  "icon for failed builds",
			Value:  ":drone:",
			EnvVar: "PLUGIN_FAILURE_ICON",
		},
		cli.StringFlag{
			Name:   "failure.template",
			Usage:  "template for failed builds",
			EnvVar: "PLUGIN_FAILURE_TEMPLATE",
		},
		cli.StringSliceFlag{
			Name:   "failure.image_attachments",
			Usage:  "image attachments for failed builds",
			EnvVar: "PLUGIN_FAILURE_IMAGE_ATTACHMENTS",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Number: c.Int("build.number"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Commit: c.String("commit.sha"),
			Branch: c.String("commit.branch"),
			Author: c.String("commit.author"),
			Link:   c.String("build.link"),
		},
		Config: Config{
			Token:   c.String("token"),
			Channel:   c.String("channel"),
			Success: MessageOptions{
				Username:         c.String("success.username"),
				Icon:             c.String("success.icon"),
				Template:         c.String("success.template"),
				ImageAttachments: c.StringSlice("success.image_attachments"),
			},
			Failure: MessageOptions{
				Username:         c.String("failure.username"),
				Icon:             c.String("failure.icon"),
				Template:         c.String("failure.template"),
				ImageAttachments: c.StringSlice("failure.image_attachments"),
			},
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}
