/*
 * Copyright 2018 Sergiy Kovtunenko
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	_ "net/http/pprof"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echopprof "github.com/sevenNt/echo-pprof"

	"github.com/skovtunenko/microservice-training/pkg"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong!")
	})
	echopprof.Wrap(e) // automatically add routers for net/http/pprof e.g. /debug/pprof, /debug/pprof/heap, etc.

	e.GET("/greet", pkg.GreetingHandler)
	e.POST("/api/users", pkg.CreateUserHandler)
	e.GET("/api/users/:id", pkg.GetUserHandler)
	e.PUT("/api/users/:id", pkg.UpdateUserHandler)
	e.DELETE("/api/users/:id", pkg.DeleteUserHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
