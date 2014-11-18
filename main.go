package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"

	"github.com/codegangsta/cli"
	"github.com/cpuguy83/dockerclient"
)

func main() {
	app := cli.NewApp()
	app.Name = "dockerdash"
	app.Usage = "A simple executive dashbaord for Docker"
	app.Version = "0.1.0"
	app.Author = "Charlie Lewis"
	app.Email = "defermat@gmail.com"
	certPath := os.Getenv("DOCKER_CERT_PATH")
	if certPath == "" {
		certPath = filepath.Join(os.Getenv("HOME"), ".docker")
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host, H",
			Value:  "/var/run/docker.sock",
			Usage:  "Location of the Docker socket",
			EnvVar: "DOCKER_HOST",
		},
		cli.BoolFlag{
			Name:   "tls",
			Usage:  "Enable TLS",
			EnvVar: "DOCKER_TLS",
		},
		cli.StringFlag{
			Name:   "tlsverify",
			Usage:  "Enable TLS Server Verification",
			EnvVar: "DOCKER_TLS_VERIFY",
		},
		cli.StringFlag{
			Name:  "tlscacert",
			Value: filepath.Join(certPath, "ca.pem"),
			Usage: "Location of tls ca cert",
		},
		cli.StringFlag{
			Name:  "tlscert",
			Value: filepath.Join(certPath, "cert.pem"),
			Usage: "Location of tls cert",
		},
		cli.StringFlag{
			Name:  "tlskey",
			Value: filepath.Join(certPath, "key.pem"),
			Usage: "Location of tls key",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "List all containers",
			Action: containerList,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "quiet, q",
					Usage: "Display only IDs",
				},
			},
		},
		{
			Name:   "count",
			Usage:  "Get count of containers",
			Action: containerCount,
		},
		{
			Name:   "info",
			Usage:  "Get host info",
			Action: info,
		},
		{
			Name:   "inspect",
			Usage:  "Get details of container",
			Action: containerInspect,
		},
	}

	app.Run(os.Args)
}

func getDockerClient(ctx *cli.Context) docker.Docker {
	docker, err := docker.NewClient(ctx.GlobalString("host"))
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = true
	if ctx.GlobalBool("tls") || ctx.GlobalString("tlsverify") != "" {
		if ctx.GlobalString("tlsverify") != "" {
			certPool := x509.NewCertPool()
			file, err := ioutil.ReadFile(ctx.GlobalString("tlscacert"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			certPool.AppendCertsFromPEM(file)
			tlsConfig.RootCAs = certPool
			tlsConfig.InsecureSkipVerify = false
		}

		_, errCert := os.Stat(ctx.GlobalString("tlscert"))
		_, errKey := os.Stat(ctx.GlobalString("tlskey"))
		if errCert == nil || errKey == nil {
			cert, err := tls.LoadX509KeyPair(ctx.GlobalString("tlscert"), ctx.GlobalString("tlskey"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't load X509 key pair: %s. Key encrpyted?\n", err)
				os.Exit(1)
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
		tlsConfig.MinVersion = tls.VersionTLS10
		docker.SetTlsConfig(&tlsConfig)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return docker
}
