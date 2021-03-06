package main

import (
	"github.com/artlovecode/wordlists.tech/functions/riot-api/pkg/handlers"
	"github.com/artlovecode/wordlists.tech/functions/riot-api/pkg/service"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	RIOT_DATADRAGON_VERSION string `envconfig:"RIOT_DATADRAGON_URL" default:"12.5.1"`
}

func main() {
	var env environment
	envconfig.MustProcess("", &env)

	service := service.New(env.RIOT_DATADRAGON_VERSION)
	champListHandler := handlers.NewChampionListHandler(service)

	lambda.Start(champListHandler)
}
