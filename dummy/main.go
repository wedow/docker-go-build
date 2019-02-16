package main

// this file is used purely for compiling project dependencies in
// the absence of the project source. this allows us to cache those
// deps in a separate docker layer and avoids having to do a full
// recompile when building new images after every source change.

// add here any deps with large build times
import (
	_ "github.com/hashicorp/terraform/helper/schema"
	_ "github.com/hashicorp/terraform/plugin"
	_ "github.com/hashicorp/terraform/terraform"
)

func main() {

}
