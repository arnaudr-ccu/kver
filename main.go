// Copyright 2019 Collabora Ltd.
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func find_release_string(path string) (string, error) {
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
	var kernel string
	var release string

	flag.StringVar(&kernel, "kernel", "", "print the release for this kernel")
	flag.StringVar(&release, "release", "", "find the kernel that matches this release")
	flag.Parse()

	if kernel != "" {
		ret, err := find_release_string(kernel)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		if ret == "" {
			fmt.Fprintln(os.Stderr, "No release found")
			os.Exit(1)
		}
		fmt.Println(ret)
	} else {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

