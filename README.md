# Coop Norge cdk8s library

> very much POC!


## Requirements

* cdk8s 

### Goal

Make a library which can support many usecases

#### Go Library

A library to build default kubernetes templates. A demo consumer can be found at [coopnorge/cdk8s-demo-consumer](https://github.com/coopnorge/cdk8s-demo-consumer)

#### Yaml input

This library also has a binary which can render manifests based on yaml input. Its in the `cmd/cdk8slibrary` folder. Edit the test.yaml and do a `go run .` and it will render the manifest in the `dist` directory.
