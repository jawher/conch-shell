// Copyright 2017 Joyent, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"encoding/json"
)

type ConchConfig struct {
	Api     string
	User    string
	Session string
}

func (c *ConchConfig) Serialize() (s string, err error) {

	j, err := json.Marshal(c)

	if err != nil {
		return "", err
	}

	return string(j), nil
}
