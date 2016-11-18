/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012-2015
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

/*

Main entry point for the `hekad` daemon. Loads the specified config and calls
`pipeline.Run` to launch the PluginRunners and all additional goroutines.

*/
package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"github.com/bbangert/toml"
)

const (
	VERSION = "0.11.0"
)

type Config struct {
	Address string `toml:"adress"`
	BaseDir string `toml: "base_dir"`
	SuffixFileName string `toml:"suffix_filename"`
	PrefixFileName string `toml:"prefix_filename"`
	SuffixIndexName string `toml:"suffix_index_name"`
	PrefixIndexName string `toml:"prefix_index_name"`
	MaxFileSize string `toml:"max_file_size"`
	FileName string `toml:"file_name"`
	StreamIndex string `toml:"stream_index"`
}

func main() {
	var config Config
	config_file := flag.String("config", filepath.FromSlash("/etc/heka-server.toml"), "Config file for log server")

	if _, err := toml.DecodeFile(*config_file, &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)
}
