// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/MartinSahlen/go-cloud-fn/app"
	"github.com/MartinSahlen/go-cloud-fn/shim"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	/*
		r := chi.NewRouter()

		// A good base middleware stack
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.StripSlashes)
		// When a client closes their connection midway through a request, the
		// http.CloseNotifier will cancel the request context (ctx).
		r.Use(middleware.CloseNotify)

		// Set a timeout value on the request context (ctx), that will signal
		// through ctx.Done() that the request has timed out and further
		// processing should be stopped.
		r.Use(middleware.Timeout(60 * time.Second))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			shim.DebugRequest(r)

			w.Write([]byte("hi"))
		})
		http.ListenAndServe(":8080", r)
		shim.ServeHTTP(r.ServeHTTP)*/
	service := goa.New("adder")
	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	adderController := NewOperandsController(service)
	app.MountOperandsController(service, adderController)

	shim.ServeHTTP(service.Mux.ServeHTTP)

}
