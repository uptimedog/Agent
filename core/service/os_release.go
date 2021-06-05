// Copyright 2020 Uptimedog. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"os"

	envparse "github.com/hashicorp/go-envparse"
)

const (
	// etcOSRelease path
	etcOSRelease = "/etc/os-release"
	// usrLibOSRelease path
	usrLibOSRelease = "/usr/lib/os-release"
)

// OsRelease types
type OsRelease struct {
	Name            string
	ID              string
	IDLike          string
	PrettyName      string
	Variant         string
	VariantID       string
	Version         string
	VersionID       string
	VersionCodename string
	BuildID         string
	ImageID         string
	ImageVersion    string
}

// GetOsRelease gets OsRelease instance
func GetOsRelease(paths []string) (*OsRelease, error) {

	if len(paths) == 0 {
		paths = []string{etcOSRelease, usrLibOSRelease}
	}

	release := &OsRelease{}
	err := release.parse(paths)

	return release, err
}

// parse parses the OS release file
func (o *OsRelease) parse(paths []string) error {

	for _, path := range paths {

		if o.Name != "" {
			return nil
		}

		releaseFile, err := os.Open(path)

		if err != nil {
			return err
		}

		defer releaseFile.Close()

		_, err = releaseFile.Stat()

		if err != nil {
			return err
		}

		env, err := envparse.Parse(releaseFile)

		if err != nil {
			return err
		}

		o.Name = env["NAME"]
		o.ID = env["ID"]
		o.IDLike = env["ID_LIKE"]
		o.PrettyName = env["PRETTY_NAME"]
		o.Variant = env["VARIANT"]
		o.VariantID = env["VARIANT_ID"]
		o.Version = env["VERSION"]
		o.VersionID = env["VERSION_ID"]
		o.VersionCodename = env["VERSION_CODENAME"]
		o.BuildID = env["BUILD_ID"]
		o.ImageID = env["IMAGE_ID"]
		o.ImageVersion = env["IMAGE_VERSION"]
	}

	if o.Name == "" {
		return errors.New("could parse os release files")
	}

	return nil
}
