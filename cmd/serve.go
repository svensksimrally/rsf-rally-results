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
			serveFolder := viper.GetString("output")
			port := viper.GetUint16("port")

			fs := http.FileServer(http.Dir(serveFolder))
			http.Handle("/", fs)

			log.Infof("Serving results from '%s' on localhost:%d", serveFolder, port)
			return http.ListenAndServe(fmt.Sprintf(":%d", port), logRequest(http.DefaultServeMux))
		},
	}
)

func init() {
	serveCmd.Flags().Uint16("port", 8000, "http server listening port")

	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(serveCmd)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
