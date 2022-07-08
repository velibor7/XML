package startup

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	cfg "github.com/velibor7/XML/api_gateway/startup/config"
	"github.com/velibor7/XML/common/loggers"
	authenticationGw "github.com/velibor7/XML/common/proto/authentication_service"
	commentGw "github.com/velibor7/XML/common/proto/comment_service"
	connectionGw "github.com/velibor7/XML/common/proto/connection_service"
	jobGw "github.com/velibor7/XML/common/proto/job_service"
	postGw "github.com/velibor7/XML/common/proto/post_service"
	profileGw "github.com/velibor7/XML/common/proto/profile_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = loggers.NewGatewayLogger()

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authenticationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	err := authenticationGw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authenticationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	connectionEndPoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	err = connectionGw.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionEndPoint, opts)
	if err != nil {
		panic(err)
	}

	profileEndpoint := fmt.Sprintf("%s:%s", server.config.ProfileHost, server.config.ProfilePort)
	err = profileGw.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, profileEndpoint, opts)
	if err != nil {
		panic(err)
	}

	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if err != nil {
		panic(err)
	}

	jobEndpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	err = jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEndpoint, opts)
	if err != nil {
		panic(err)
	}

	commentEndpoint := fmt.Sprintf("%s:%s", server.config.CommentHost, server.config.CommentPort)
	err = commentGw.RegisterCommentServiceHandlerFromEndpoint(context.TODO(), server.mux, commentEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server Server) Start() {
	// cors := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"http://localhost:8080/**", "http://localhost:3000/**", "http://localhost:3000"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-Origin", ""}),
	// 	handlers.AllowCredentials(),
	// )
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(server.mux)))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		log.WithFields(logrus.Fields{
			"method":     r.Method,
			"url":        r.URL.String(),
			"origin":     r.Header.Get("Origin"),
			"user-agent": r.Header.Get("User-Agent"),
		}).Info("CORS filter")

		if r.Header.Get("Origin") != "" {
			h.Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		//h.Set("Access-Control-Allow-Origin", "https://localhost:7777")

		if r.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
					http.MethodHead,
					http.MethodOptions,
					http.MethodTrace,
				}, ", ",
			))

			h.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
					"Authorization",
					"Location",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}
