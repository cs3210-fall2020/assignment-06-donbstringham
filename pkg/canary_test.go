// Package pkg provides support for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanary(t *testing.T) {
	assert.True(t, true, "canary test passing")
}
