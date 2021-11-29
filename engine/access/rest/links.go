package rest

import (
	"fmt"
	"github.com/onflow/flow-go/engine/access/rest/generated"
	"github.com/onflow/flow-go/model/flow"
	"regexp"
)

var pathRegex = regexp.MustCompile(`{[a-z]*}`)

func accountLink(address flow.Address) *generated.Links {
	link := getLink(getAccountRoute, address.String())
	return linkResponse(link)
}

func linkResponse(self string) *generated.Links {
	return &generated.Links{Self: self}
}

func getLink(name string, value string) string {
	routes := routeDefinitions()
	var route string
	for _, r := range routes {
		if r.name == name {
			route = r.pattern
		}
	}

	if route == "" {
		return ""
	}

	const version = "v1"
	return fmt.Sprintf(
		"/%s%s",
		version,
		pathRegex.ReplaceAllString(route, value),
	)
}
