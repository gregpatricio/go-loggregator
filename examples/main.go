package main

import (
	"log"
	"os"

	loggregator "code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/v2"
)

func main() {
	metronCfg := v2.MetronConfig{
		UseV2API:      true,
		APIPort:       3458,
		CACertPath:    os.Getenv("CA_CERT_PATH"),
		CertPath:      os.Getenv("CERT_PATH"),
		KeyPath:       os.Getenv("KEY_PATH"),
		JobDeployment: "example-deployment",
		JobName:       "example-job",
		JobIndex:      "example-index",
		JobIP:         "0.0.0.0",
		JobOrigin:     "example-deployment",
	}

	client, err := loggregator.NewClient(metronCfg)

	if err != nil {
		log.Fatal("Could not create client", err)
	}

	err = client.SendMetric("some-metric-name", 1234)
	if err != nil {
		log.Fatal("Unable to send metric", err)
	}
}
