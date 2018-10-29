package buildserver

import (
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/api/auth"
	"github.com/concourse/concourse/atc/db"
)

type EventHandlerFactory func(lager.Logger, db.Build) http.Handler

type SchedulerFactory interface {
	BuildScheduler(db.Pipeline, string, creds.Variables) scheduler.BuildScheduler
}

type Server struct {
	logger lager.Logger

	externalURL string
	peerURL     string

	engine              engine.Engine
	workerClient        worker.Client
	teamFactory         db.TeamFactory
	buildFactory        db.BuildFactory
	eventHandlerFactory EventHandlerFactory
	drain               <-chan struct{}
	rejector            auth.Rejector

	// Used for the creation of rebuild builds
	schedulerFactory SchedulerFactory
	variablesFactory creds.VariablesFactory
}

func NewServer(
	logger lager.Logger,
	externalURL string,
	peerURL string,
	engine engine.Engine,
	workerClient worker.Client,
	teamFactory db.TeamFactory,
	buildFactory db.BuildFactory,
	eventHandlerFactory EventHandlerFactory,
	drain <-chan struct{},
) *Server {
	return &Server{
		logger: logger,

		externalURL: externalURL,
		peerURL:     peerURL,

		engine:              engine,
		workerClient:        workerClient,
		teamFactory:         teamFactory,
		buildFactory:        buildFactory,
		eventHandlerFactory: eventHandlerFactory,
		drain:               drain,

		rejector: auth.UnauthorizedRejector{},
	}
}
