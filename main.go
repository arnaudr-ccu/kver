// Copyright 2019 Collabora Ltd.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Find a release-like string in a file, using a regexp.
// Returns the first match.
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

// Find a static string in a file.
// Returns true if found.
func find_static_string(path string, str string) (bool, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	return bytes.Contains(buf, []byte(str)), nil
}

// List the files /boot/vmlinuz-*.
// Returns the files found.
func list_vmlinuzs() ([]string, error) {
	files, err := ioutil.ReadDir("/boot")
	if err != nil {
		return nil, err
	}

	kernels := []string{}
	for _, f := range files {
		if !f.IsDir() && strings.HasPrefix(f.Name(), "vmlinuz-") {
			kernels = append(kernels, filepath.Join("/boot", f.Name()))
		}
	}

	return kernels, nil
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

	} else if release != "" {
		if ! strings.HasSuffix(release, " ") {
			release += " "
		}

		kernels, err := list_vmlinuzs()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		kernel_found := ""
		for i := len(kernels) - 1; i >= 0; i-- {
			//fmt.Fprintln(os.Stderr, "Scanning", filename, "...")
			found, err := find_static_string(kernels[i], release)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error while scanning:", err)
				continue
			}
			if ! found {
				continue
			}

			kernel_found = kernels[i]
			break
		}

		if kernel_found == "" {
			fmt.Fprintln(os.Stderr, "No kernel found")
			os.Exit(1)
		}

		fmt.Println(kernel_found)

	} else {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

