// Copyright 2019 Collabora Ltd.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func find_version_string(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	match := regexp.MustCompile(`\d\d*\.\d\d*\.\d\d*\S*`).Find(buf)
	if match == nil {
		return "", nil
	}

	// cleanup any trailing non-printable characters
	match = regexp.MustCompile(`[[:ascii:]]*`).Find(match)
	return string(match), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "KERNEL_PATH")
		os.Exit(1)
	}

	path := os.Args[1]
	ver, err := find_version_string(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if ver == "" {
		fmt.Fprintln(os.Stderr, "No version found")
		os.Exit(1)
	}

	fmt.Println(ver)
}
