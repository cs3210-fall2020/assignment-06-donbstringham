// Package util provides support for the application.
// Copyright 2020 Don B. Stringham All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.
//
// @author donbstringham <donbstringham@icloud.com>
//
package util

import (
	"io/ioutil"
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// GetFiles returns an array/slice of file names for the passed in directory.
func GetFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filesStr []string
	// Convert os.Fileinfo into strings
	for _, file := range files {
		filesStr = append(filesStr, file.Name())
	}
	// Alphabetically sort the result set
	sort.Strings(filesStr)

	return filesStr, nil
}

// FindFiles retrieves the name of the file if found
func FindFiles(dir string, seed string) ([]string, error) {
	files, err := GetFiles(dir)
	if err != nil {
		return nil, err
	}

	fndFiles := fuzzy.Find(seed, files)

	return fndFiles, nil
}
