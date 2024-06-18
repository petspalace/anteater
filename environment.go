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

package anteater

import (
	"fmt"
	"log"
	"os"
)

type Environment struct {
	MqttHost     string
	MqttPrefix   string
	MqttTopic    string
	MqttTemplate string
}

func NewEnvironmentFromEnv() (*Environment, error) {
	hostFromEnv, hostExists := os.LookupEnv("MQTT_HOST")

	if !hostExists {
		return nil, fmt.Errorf("anteater needs `MQTT_HOST` set in the environment to a value such as `tcp://127.0.0.1:1883`.")
	}

	prefixFromEnv, prefixExists := os.LookupEnv("MQTT_PREFIX")

	if !prefixExists {
		log.Println("`MQTT_PREFIX` undefined using default `home.arpa`-prefix.")
		prefixFromEnv = "/home.arpa"
	} else {
		log.Printf("`MQTT_PREFIX` set to `%s`.\n", prefixFromEnv)
	}

	topicFromEnv, topicExists := os.LookupEnv("MQTT_PREFIX")

	if !topicExists {
		log.Println("`MQTT_TOPIC` undefined using default `anteater`-topic.")
		topicFromEnv = "/anteater"
	} else {
		log.Printf("`MQTT_TOPIC` set to `%s`.\n", topicFromEnv)
	}

	templateFromEnv, templateExists := os.LookupEnv("MQTT_PREFIX")

	if !templateExists {
		log.Println("`MQTT_TEMPLATE` undefined using default `{{.URL}}`-template.")
		templateFromEnv = "{{.URL}}"
	} else {
		log.Printf("`MQTT_TEMPLATE` set to `%s`.\n", templateFromEnv)
	}

	configuration := Environment{
		MqttHost:     hostFromEnv,
		MqttPrefix:   prefixFromEnv,
		MqttTopic:    topicFromEnv,
		MqttTemplate: templateFromEnv,
	}

	log.Printf("%s %s %s %s", hostFromEnv, prefixFromEnv, topicFromEnv, templateFromEnv)

	return &configuration, nil
}

// SPDX-License-Identifier: MIT
// vim: ts=4 sw=4 noet
