package cmd

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serveCmd = &cobra.Command{
		Use:           "serve",
		Short:         "Serve a generated result site on this machine",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := http.FileServer(http.Dir(viper.GetString("output")))
			http.Handle("/", fs)

			// TODO: make this configurable
			port := 8000

			log.WithField("port", port).Info("Serving results on localhost")
			return http.ListenAndServe(fmt.Sprintf(":%d", port), logRequest(http.DefaultServeMux))
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
