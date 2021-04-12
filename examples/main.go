package main

import (
	"fmt"
	"log"

	"github.com/bjornm82/go-swagger-schema-registry/client"
	"github.com/bjornm82/go-swagger-schema-registry/client/operations"
	"github.com/bjornm82/go-swagger-schema-registry/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client struct {
	*client.ConfluentSchemaRegistry
}

var cl = Client{}

func init() {
	host := "localhost"
	port := 8081
	usessl := false

	var scheme = "http"
	if usessl {
		scheme = "https"
	}

	transport := httptransport.New(fmt.Sprintf("%s:%d", host, port), "", []string{scheme})
	// Needs to be set
	transport.Consumers["application/vnd.schemaregistry.v1+json"] = runtime.JSONConsumer()
	transport.Producers["application/vnd.schemaregistry.v1+json"] = runtime.JSONProducer()

	// Create client
	cl.ConfluentSchemaRegistry = client.New(transport, strfmt.Default)
}

func main() {

	b, err := cl.setTopLevelConfig("BACKWARD")
	if err != nil {
		log.Fatalln(err)
	}
	if string(b) != `{"compatibility":"BACKWARD"}` {
		log.Fatalln("should print BACKWARD")
	}

	b, err = cl.getTopLevelConfig()
	if err != nil {
		log.Fatalln(err)
	}
	if string(b) != `{"compatibilityLevel":"BACKWARD"}` {
		log.Fatalln("should print BACKWARD")
	}

	b, err = cl.setTopLevelConfig("FULL_TRANSITIVE")
	if err != nil {
		log.Fatalln(err)
	}
	if string(b) != `{"compatibility":"FULL_TRANSITIVE"}` {
		log.Fatalln("should print FULL_TRANSITIVE")
	}

	b, err = cl.getTopLevelConfig()
	if err != nil {
		log.Fatalln(err)
	}
	if string(b) != `{"compatibilityLevel":"FULL_TRANSITIVE"}` {
		log.Fatalln("should print FULL_TRANSITIVE")
	}
}

// GetTopLevelConfig will return the bytestring of current set config
func (c *Client) getTopLevelConfig() ([]byte, error) {
	getTopLevelConfig := operations.NewGetTopLevelConfigParams()
	v, err := cl.Operations.GetTopLevelConfig(getTopLevelConfig)
	if err != nil {
		return nil, err
	}
	return v.Payload.MarshalBinary()
}

// SetTopLevelConfig will set a new toplevel config
func (c *Client) setTopLevelConfig(level string) ([]byte, error) {
	mod := models.ConfigUpdateRequest{}
	params := operations.NewUpdateTopLevelConfigParams()
	mod.Compatibility = level
	params.SetBody(&mod)

	confOK, err := cl.Operations.UpdateTopLevelConfig(params)
	if err != nil {
		return nil, err
	}

	return confOK.Payload.MarshalBinary()
}
