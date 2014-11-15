package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func containerList(ctx *cli.Context) {
	docker := getDockerClient(ctx)
	containers, err := docker.FetchAllContainers(true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, c := range containers {
		fmt.Println(c.Id)
	}
	// TODO
}

func containerInspect(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Malformed argument. Please supply 1 and only 1 argument")
		os.Exit(1)
	}

	docker := getDockerClient(ctx)
	container, err := docker.FetchContainer(ctx.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(container)
	// TODO
}