#!/bin/bash

# Copyright The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

echo "Running e2e-tests.sh"

unset GOFLAGS
tmp="$(mktemp -d)"

# Default branch for both CCAPIO and this repo should be `main`
CCAPIO_BASE_REF=$PULL_BASE_REF

echo "cloning github.com/openshift/cluster-capi-operator at branch '$CCAPIO_BASE_REF'"
git clone --single-branch --branch="$CCAPIO_BASE_REF" --depth=1 "https://github.com/openshift/cluster-capi-operator.git" "$tmp"

echo "running cluster-capi-operator's: make e2e"
exec make -C "$tmp" e2e
