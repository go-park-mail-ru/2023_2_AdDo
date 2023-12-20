package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"google.golang.org/grpc"
	"main/internal/microservices/candidate/proto"
	candidate_service_server "main/internal/microservices/candidate/service/server"
	activity_repository "main/internal/pkg/activity/repository/memcached"
	cluster_repository2 "main/internal/pkg/cluster/cluster_repository/in_memory"

	// microservices_init "main/cmd/microservices"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8088
const PathToClusterTracks = "clustering_data.json"

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	//logger.SetLevel(logrus.ErrorLevel)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting candidate micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	clusterRepo, err := cluster_repository2.NewInMemory(PathToClusterTracks, logger)
	if err != nil {
		logger.Errorln("error loading cluster tracks dump ", err)
	}

	mc := memcache.New("memcached:11211")

	recentActivityRepo := activity_repository.NewMemCached(mc, logger)
	trackRepository := track_repository.NewPostgres(pool, logger)

	candidateManager := candidate_service_server.NewCandidateManager(&recentActivityRepo, trackRepository, clusterRepo, logger)

	server := grpc.NewServer()
	proto.RegisterCandidateServiceServer(server, &candidateManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
