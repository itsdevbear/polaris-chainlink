# Chainlink Integration Tests Template

[![Go Version](https://img.shields.io/github/go-mod/go-version/smartcontractkit/integrations-framework)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Use this template directory as a quick start to writing your own Chainlink tests using the [integrations framework](https://smartcontractkit.github.io/integrations-framework/).

## Make Commands

`make go_mod`

`make test_environments`

`make test_contracts`

## Crucial Libraries

The intent is to keep this repo as simple as possible so others can read it with minimal issue and can modify it as they please.

* [zerolog](https://github.com/rs/zerolog): Efficient and pretty logging
* [testify](https://github.com/stretchr/testify): A minimal test assertion library (we're fans of [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) if you feel like getting fancy)
* [helmenv](https://github.com/smartcontractkit/helmenv): Used to deploy and connect to test environments
* [integrations framework](https://smartcontractkit.github.io/integrations-framework/): Used to run tests and interact with the test environments
