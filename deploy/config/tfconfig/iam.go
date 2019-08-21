// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tfconfig

// ProjectIAMMember represents a Terraform project IAM member.
type ProjectIAMMember struct {
	Role   string `json:"role"`
	Member string `json:"member"`

	// The following fields should not be set by users.

	// ForEach is used to let a single iam member expand to reference multiple iam members
	// through the use of terraform's for_each iterator.
	ForEach map[string]*ProjectIAMMember `json:"for_each,omitempty"`
	Project string                       `json:"project,omitempty"`
}

// Init initializes the resource.
func (m *ProjectIAMMember) Init(string) error {
	return nil
}

// ID returns the resource unique identifier.
func (m *ProjectIAMMember) ID() string {
	return m.Project
}

// ResourceType returns the resource terraform provider type.
func (m *ProjectIAMMember) ResourceType() string {
	return "google_project_iam_member"
}
