package cmd

import (
	"gitlab.com/wiggins.jonathan/plutus/server"

	"github.com/spf13/cobra"
)

func init() {
	serveCmd.Flags().IntP("port", "p", 8080, "Port on which to serve plutus.")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"server"},
	Short:   "Serve plutus webserver",
	Long:    "Serve plutus API via webserver over HTTP",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// ListenAndServe will deal with port errors
		port, _ := cmd.Flags().GetInt("port")

		s := server.NewServer(
			server.WithDebug(debug),
			server.WithPort(port),
		)

		if err := s.Serve(); err != nil {
			return err
		}

		return nil
	},
}
