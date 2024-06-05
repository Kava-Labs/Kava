package main

import (
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/cmd/kava/cmd"

	_ "net/http/pprof"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	go func() {
		if err := http.ListenAndServe("0.0.0.0:6061", nil); err != nil {
			panic(err)
		}
	}()

	if err := svrcmd.Execute(rootCmd, cmd.EnvPrefix, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
