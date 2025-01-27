// Copyright 2023 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

import _ "embed"

// inputFileTemplate is the template for the input file.
//
//go:embed 00-apply.yaml.tmpl
var inputFileTemplate string

// assertFileTemplate is the template for the assert file.
//
//go:embed 00-assert.yaml.tmpl
var assertFileTemplate string

// deleteFileTemplate is the template for the delete file.
//
//go:embed 01-delete.yaml.tmpl
var deleteFileTemplate string

// assertDeletedFileTemplate is the template for delete assert file.
//
//go:embed 01-assert.yaml.tmpl
var assertDeletedFileTemplate string
