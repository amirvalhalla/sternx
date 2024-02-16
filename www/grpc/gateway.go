package grpc

//nolint:all
import (
	"fmt"
	"mime"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	sternx "sternx/www/grpc/gen/go"
	_ "sternx/www/grpc/statik"
)

func (s *Server) startGateway(grpcAddr string) error {
	conn, err := grpc.DialContext(
		s.ctx,
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwMux := runtime.NewServeMux()
	err = sternx.RegisterUserHandler(s.ctx, gwMux, conn)
	if err != nil {
		return err
	}

	oa, err := s.getOpenAPIHandler()
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:              s.config.GRPCGatewayListen,
		ReadHeaderTimeout: 3 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/v1") {
				gwMux.ServeHTTP(w, r)

				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	gwServer.Handler = allowCORS(gwServer.Handler)

	listener, err := net.Listen("tcp", s.config.GRPCGatewayListen)
	if err != nil {
		return err
	}

	s.logger.Info("grpc-gateway started listening", "address", listener.Addr().String())

	go func() {
		if err := gwServer.Serve(listener); err != nil {
			s.logger.Error("error on grpc-gateway serve", "error", err)
		}
	}()

	return nil
}

func (s *Server) getOpenAPIHandler() (http.Handler, error) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		return nil, err
	}

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	return http.FileServer(statikFS), nil
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w)

				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}
