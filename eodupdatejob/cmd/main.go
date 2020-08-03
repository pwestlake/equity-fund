package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pwestlake/aemo/userservice/pkg/config"
	"github.com/pwestlake/equity-fund/eodupdatejob/pkg/jobs"
)

func main() {
	configPtr := flag.String("server", "http://localhost:8888", "the url of the cloud config server")
	profilePtr := flag.String("profile", "dev", "the configuration profile")
	labelPtr := flag.String("label", "development", "the configuration label")
	typePtr := flag.String("type", "latest", "the type of operation latest|backfill")
	symbolPtr := flag.String("symbol", "", "the symbol to backfill")
	flag.Parse()

	configURL := fmt.Sprintf("%s/equity-fund-eodupdatejob", *configPtr)
	config.NewConfig(&config.Params{Server: configURL, Profile: *profilePtr, Label: *labelPtr})

	backfillJob := jobs.InitializeBackFillJob()
	switch *typePtr {
	case "backfill":
		backfillJob.Run(*symbolPtr)
	case "latest":
		backfillJob.UpdateWithLatest()
	default:
		log.Print("Nothing to do")
	}
}
