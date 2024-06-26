/* This program is licensed under the MIT license:
 *
 * Copyright 2024 Simon de Vlieger
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/petspalace/anteater"
)

func main() {
	_, err := anteater.NewEnvironmentFromEnv()

	if err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	state, err := anteater.NewState()
	c := make(chan anteater.Item, 16)

	if err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	go anteater.Loop(60*time.Second, state, c)

	for item := range c {
		slog.Info(fmt.Sprintf("main: %v\n", item))
	}
}

// SPDX-License-Identifier: MIT
// vim: ts=4 sw=4 noet
