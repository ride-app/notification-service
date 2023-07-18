package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ride-app/notification-service/api/gen/ride/notification/v1alpha1/notificationv1alpha1connect"
	"github.com/ride-app/notification-service/config"
	"github.com/ride-app/notification-service/di"
	"github.com/ride-app/notification-service/interceptors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	log "github.com/sirupsen/logrus"
)

func main() {
	service, err := di.InitializeService()

	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	log.Info("Service Initialized")

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authInterceptor, err := interceptors.NewAuthInterceptor(ctx)

	if err != nil {
		log.Fatalf("Failed to initialize auth interceptor: %v", err)
	}

	connectInterceptors := connect.WithInterceptors(authInterceptor)

	path, handler := notificationv1alpha1connect.NewNotificationServiceHandler(service, connectInterceptors)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Env.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))

}

func init() {
	log.SetReportCaller(true)
	// if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableLevelTruncation: true,
	// 	PadLevelText:           true,
	// 	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
	// 		dir, err := os.Getwd()
	// 		if err != nil {
	// 			dir = ""
	// 		} else {
	// 			dir = dir + "/"
	// 		}

	// 		filename := strings.Replace(f.File, dir, "", -1)

	// 		return fmt.Sprintf("(%s)", path.Base(f.Function)), fmt.Sprintf(" %s:%d", filename, f.Line)

	// 	},
	// })
	// }

	log.SetFormatter(&log.JSONFormatter{})

	err := cleanenv.ReadEnv(&config.Env)

	if err != nil {
		log.Warnf("Could not load config: %v", err)
	}
}
