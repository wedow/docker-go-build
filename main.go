package main

// this file is meant to simulate building our real app code which
// requires the same deps as found in ./dummy/main.go

// without the prebuild step using that file, this file can take
// several minutes to compile when building a docker image. with
// the prebuild step, this file only takes a few seconds.

import (
	_ "github.com/hashicorp/terraform/helper/schema"
	_ "github.com/hashicorp/terraform/plugin"
	_ "github.com/hashicorp/terraform/terraform"
)

func main() {

}
