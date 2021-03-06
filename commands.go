package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageList(ctx *cli.Context) {
	fmt.Println("todo")
	// TODO
}

func imageInspect(ctx *cli.Context) {
	fmt.Println("todo")
	// TODO
}

func containerCount(ctx *cli.Context) {
	docker := getDockerClient(ctx)
	containers, err := docker.FetchAllContainers(true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(len(containers))
}

func info(ctx *cli.Context) {
	docker := getDockerClient(ctx)
	info, err := docker.Info()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(info)
}

func containerList(ctx *cli.Context) {
	docker := getDockerClient(ctx)
	containers, err := docker.FetchAllContainers(true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, c := range containers {
		fmt.Println(c.Id)
		fmt.Println(c.Name)
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
