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

package config

import (
	"fmt"

	"github.com/GoogleCloudPlatform/healthcare/deploy/config/tfconfig"
)

func (p *Project) initTerraform() error {
	for _, r := range p.TerraformResources() {
		if err := r.Init(p.ID); err != nil {
			return err
		}
	}
	return nil
}

// TerraformResources gets all terraform data resources in this project.
func (p *Project) TerraformResources() []tfconfig.Resource {
	var rs []tfconfig.Resource
	for _, r := range p.StorageBuckets {
		rs = append(rs, r)
	}
	for _, r := range p.BigqueryDatasets {
		rs = append(rs, r)
	}

	if len(p.IAMMembers) > 0 {
		forEach := make(map[string]*tfconfig.ProjectIAMMember)
		for _, m := range p.IAMMembers {
			key := fmt.Sprintf("%s %s", m.Role, m.Member)
			forEach[key] = m
		}

		rs = append(rs, &tfconfig.ProjectIAMMember{
			ForEach: forEach,
			Project: p.ID,
			Role:    "${each.value.role}",
			Member:  "${each.value.member}",
		})
	}
	return rs
}
