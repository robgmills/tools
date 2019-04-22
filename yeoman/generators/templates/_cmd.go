package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/spothero/tools/service"
	"github.com/spothero/<%=appName%>/pkg/<%=appName%>"
)

// These variables should be set during build with the Go link tool
// e.x.: when running go build, provide -ldflags="-X main.version=1.0.0"
var version = "not-set"
var gitSHA = "not-set"

// This is the main entrypoint of the program. Here we create our root command and then execute it.
func main() {
	serverCmd := service.HTTPConfig{
		Name:             "<%=appName%>",
		Version:          version,
		GitSHA:           gitSHA,
		RegisterHandlers: <%=appName%>.RegisterHandlers,
	}
	if err := serverCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
