package initializers

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

func ConnectES(config *Config) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username:               config.ESUser,
		Password:               config.ESPassword,
		CertificateFingerprint: config.ESCertFingerPrint,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
