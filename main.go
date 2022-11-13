package main

import (
	"os"

	"github.com/keyvchan/NetAssist/client"
	"github.com/keyvchan/NetAssist/pkg/flags"
	"github.com/keyvchan/NetAssist/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "NetAssit",
		Short: "NetAssit is a network debugging and testing tool",
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start a server",
		Run: func(_ *cobra.Command, _ []string) {
			// pass the config to it
			flags.Config.Type = "server"

			// concat the protocol and host
			server.Serve()

		},
	}

	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Start a client",
		Run: func(_ *cobra.Command, _ []string) {
			// set type
			flags.Config.Type = "client"
			client.Req()
		},
	}
)

func execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	log.Info().Msg("initConfig")

}

func init() {

	// setup log
	cobra.OnInitialize(initConfig)

	serverCmd.PersistentFlags().StringVarP(&flags.Config.Protocol, "protocol", "p", "tcp", "protocol")
	serverCmd.PersistentFlags().IntVarP(&flags.Config.Port, "port", "P", 8080, "port")
	serverCmd.PersistentFlags().StringVarP(&flags.Config.Host, "host", "H", "127.0.0.1", "host")
	// protocol port host type is all required
	serverCmd.MarkFlagsRequiredTogether("protocol")
	serverCmd.MarkFlagsRequiredTogether("host", "port")
	serverCmd.PersistentFlags().BoolVarP(&flags.Config.Binary, "binary", "b", false, "binary")

	rootCmd.AddCommand(serverCmd)

	clientCmd.PersistentFlags().StringVarP(&flags.Config.Protocol, "protocol", "p", "tcp", "protocol")
	clientCmd.PersistentFlags().IntVarP(&flags.Config.Port, "port", "P", 8080, "port")
	clientCmd.PersistentFlags().StringVarP(&flags.Config.Host, "host", "H", "127.0.0.1", "host")
	// protocol port host type is all required
	clientCmd.MarkFlagsRequiredTogether("protocol")
	clientCmd.MarkFlagsRequiredTogether("port", "host")
	clientCmd.PersistentFlags().BoolVarP(&flags.Config.Binary, "binary", "b", false, "binary")
	rootCmd.AddCommand(clientCmd)

}

func main() {
	// setup log
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// set command line parsing
	rootCmd.Execute()

}
