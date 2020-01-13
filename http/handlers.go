// Copyright 2020 SpotHero
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spothero/tools/http/ready"
	"github.com/spothero/tools/log"
	"go.uber.org/zap"
)

// healthHandler is a simple HTTP handler that returns 200 OK
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "OK")
}

func readyHandler(indicators map[string]ready.Indicator) http.HandlerFunc {
	monitor := ready.NewMonitor(indicators)
	return func(w http.ResponseWriter, r *http.Request) {
		report := monitor.ReadyCheck()
		if err := json.NewEncoder(w).Encode(report); err != nil {
			log.Get(context.Background()).Error("Error encoding json response", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		if report.Overall != ready.Ready {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
