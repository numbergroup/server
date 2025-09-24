package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/numbergroup/errors"
	"github.com/sirupsen/logrus"
)

// ListenWithGracefulShutdown listens on the given address with the given gin router, and will shutdown gracefully when a SIGINT or SIGTERM signal is received.
// Also setups up a health check endpoint on the given path, which will be unhealthy when the server is shutting down.
// This allows for graceful shutdown of the server in a Kubernetes environment, where the health check is used to determine
// if the server is ready to accept traffic.
func ListenWithGracefulShutdown(ctx context.Context, log logrus.Ext1FieldLogger, router *gin.Engine, conf Config) error {
	// Wrap the gin router in http.Server so we can call Shutdown
	hc := NewHealthCheck()
	srv := &http.Server{
		Addr:              conf.Listen,
		Handler:           router.Handler(),
		ReadTimeout:       conf.ReadTimeout,
		ReadHeaderTimeout: conf.ReadTimeout,
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	router.GET(conf.HealthCheckPath, hc.Health)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("failed to listen and serve")
		}
	}()

	<-ctx.Done()

	log.WithField("reason", ctx.Err()).Warn("shutting down server")
	ctx, cancel := context.WithTimeout(ctx, conf.ShutdownTimeout)
	defer cancel()
	// Set the health check to unhealthy, so we can stop accepting new requests
	hc.IsUnhealthy.Store(true)
	if err := srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shutdown server")
	}

	return nil

}
