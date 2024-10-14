// TODO: Change Copyright to your company if open sourcing or remove header
//
// Copyright (c) 2021 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"edgex-app-template/functions"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
)

const (
	serviceKey = "edgex-app-template"
)

// TODO: Define your app's struct
type myApp struct {
	service interfaces.ApplicationService
	lc      logger.LoggingClient
}

func main() {
	// TODO: See https://docs.edgexfoundry.org/latest/microservices/application/ApplicationServices/
	//       for documentation on application services.
	app := myApp{}
	code := app.CreateAndRunAppService(serviceKey, pkg.NewAppService)
	os.Exit(code)
}

// CreateAndRunAppService wraps what would normally be in main() so that it can be unit tested
// TODO: Remove and just use regular main() if unit tests of main logic not needed.
func (app *myApp) CreateAndRunAppService(serviceKey string, newServiceFactory func(string) (interfaces.ApplicationService, bool)) int {
	var ok bool
	app.service, ok = newServiceFactory(serviceKey)
	if !ok {
		return -1
	}

	app.lc = app.service.LoggingClient()

	sample := functions.NewSample()

	// TODO: Replace below functions with built in and/or your custom functions for your use case
	//       or remove if using Pipeline By Topics below.
	//       See https://docs.edgexfoundry.org/latest/microservices/application/BuiltIn/ for list of built-in functions
	err := app.service.SetDefaultFunctionsPipeline(
		sample.LogEventDetails)

	if err != nil {
		app.lc.Errorf("SetFunctionsPipeline returned error: %s", err.Error())
		return -1
	}

	// TODO: Add any custom routes your service may have for its REST API
	if err := app.service.AddCustomRoute("/api/v3/hello", true, app.helloHandler, http.MethodGet); err != nil {
		app.lc.Errorf("AddCustomRoute returned error: %s", err.Error())
		return -1
	}

	if err := app.service.Run(); err != nil {
		app.lc.Errorf("Run returned error: %s", err.Error())
		return -1
	}

	// TODO: Do any required cleanup here, if needed

	return 0
}

func (app *myApp) helloHandler(c echo.Context) error {
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write([]byte("hello"))
	return nil
}
