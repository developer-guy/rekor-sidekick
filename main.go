package main

import (
	"github.com/nsmith5/rekor-sidekick/cli"

	// Loading output drivers
	_ "github.com/nsmith5/rekor-sidekick/outputs"
	_ "github.com/nsmith5/rekor-sidekick/outputs/cloudevents"
	_ "github.com/nsmith5/rekor-sidekick/outputs/pagerduty"
	_ "github.com/nsmith5/rekor-sidekick/outputs/stdout"
)

func main() {
	cmd := cli.New()
	cmd.Execute()
}
