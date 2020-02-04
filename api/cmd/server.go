package main

import (
	"fmt"
	ctrl "github.com/panda-next-team/poolrank-basedata/api/internal/app"
	"github.com/panda-next-team/poolrank-basedata/api/internal/pkg"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var config pkg.Config
	app := cli.NewApp()
	app.Name = "poolrank-basedata-api"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2019 poolrank"
	app.Usage = "Poolrank Basedata Grpc server"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Alex Zhou",
			Email: "333266664@qq.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port, p",
			Value:       80,
			Usage:       "server port",
			EnvVar:      "BASEDATA_SERVER_PORT",
			Destination: &config.Port,
		},
		cli.StringFlag{
			Name:        "dbHost, dbh",
			Value:       "127.0.0.1",
			Usage:       "mysql host",
			EnvVar:      "BASEDATA_DB_HOST",
			Destination: &config.Mysql.Host,
		},
		cli.StringFlag{
			Name:        "dbUser, dbu",
			Value:       "root",
			Usage:       "mysql user",
			EnvVar:      "BASEDATA_DB_USER",
			Destination: &config.Mysql.User,
		},
		cli.StringFlag{
			Name:        "dbPassword, dbp",
			Value:       "",
			Usage:       "mysql password",
			EnvVar:      "BASEDATA_DB_PASSWORD",
			Destination: &config.Mysql.Password,
		},
		cli.StringFlag{
			Name:        "dbName, dbn",
			Value:       "pr_basedata",
			Usage:       "mysql database name",
			EnvVar:      "BASEDATA_DB_NAME",
			Destination: &config.Mysql.Database,
		},
		cli.IntFlag{
			Name:        "dbPort",
			Value:       3306,
			Usage:       "mysql db port",
			EnvVar:      "BASEDATA_DB_PORT",
			Destination: &config.Mysql.Port,
		},
		cli.StringFlag{
			Name:        "dbPrefix",
			Value:       "bm_",
			Usage:       "mysql db prefix",
			EnvVar:      "BASEDATA_DB_PREFIX",
			Destination: &config.Mysql.Prefix,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run grpc server",
			Action: func(c *cli.Context) {
				log.Println("poolrank basedata grpc server start, listen: ", config.Port)

				engine, engineErr := pkg.NewMysqlEngine(pkg.DriverName, pkg.GenerateMysqlSource(
					config.Mysql.User, config.Mysql.Password, config.Mysql.Host,
					config.Mysql.Port, config.Mysql.Database), config.Mysql.Prefix)

				if engineErr  != nil {
					log.Fatal(engineErr)
				}

				port := fmt.Sprintf(":%d", config.Port)
				listener, err := net.Listen("tcp", port)
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				server := grpc.NewServer()
				pb.RegisterPOWCoinServiceServer(server, &ctrl.PowCoinService{engine})
				pb.RegisterPoolServiceServer(server, &ctrl.PoolService{engine})
				pb.RegisterAlgorithmServiceServer(server, &ctrl.AlgorithmService{engine})

				if err := server.Serve(listener); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}

			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
