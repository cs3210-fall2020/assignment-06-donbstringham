// Package pkg provides commands for the application.
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

func TestGetFilesSuccess(t *testing.T) {
	//arrange
	curDir := "./"
	expected := []string{"canary_test.go", "dir.go", "dir_test.go"}
	//act
	actual, err := GetFiles(curDir)
	//assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, expected, actual, "should be identical")
}

func TestGetFilesFailure(t *testing.T) {
	//arrange
	curDir := "./foobar"
	expected := []string{"dir.go", "dir_test.go"}
	//act
	actual, err := GetFiles(curDir)
	//assert
	assert.NotNil(t, err, "expecting an error")
	assert.NotEqual(t, expected, actual, "should not be identical")
}

func TestFindFilesSuccessFullName(t *testing.T) {
	//arrange
	curDir := "./"
	expected := []string{"canary_test.go"}
	seed := "canary_test.go"
	//act
	actual, err := FindFiles(curDir, seed)
	//assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, expected, actual, "should be identical")
}

func TestFindFilesSuccessPartialName(t *testing.T) {
	//arrange
	curDir := "./"
	expected := []string{"canary_test.go"}
	seed := "can"
	//act
	actual, err := FindFiles(curDir, seed)
	//assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, expected, actual, "should be identical")
}

func TestFindFilesSuccessWildcardName(t *testing.T) {
	//arrange
	curDir := "./"
	expected := []string{"canary_test.go"}
	seed := "c"
	//act
	actual, err := FindFiles(curDir, seed)
	//assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, expected, actual, "should be identical")
}

func TestFindFilesSuccessWildcardExt(t *testing.T) {
	//arrange
	curDir := "./"
	expected := []string{"canary_test.go", "dir.go", "dir_test.go"}
	seed := "go"
	//act
	actual, err := FindFiles(curDir, seed)
	//assert
	assert.Nil(t, err, "not expecting an error")
	assert.Equal(t, expected, actual, "should be identical")
}
