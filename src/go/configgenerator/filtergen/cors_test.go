// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filtergen_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/esp-v2/src/go/configgenerator/filtergen"
	"github.com/GoogleCloudPlatform/esp-v2/src/go/options"
)

func TestNewCORSFilterGensFromOPConfig_GenConfig(t *testing.T) {
	testdata := []SuccessOPTestCase{
		{
			Desc: "Generate with CORS is set to basic mode",
			OptsIn: options.ConfigGeneratorOptions{
				CorsPreset: "basic",
			},
			WantFilterConfigs: []string{
				`
{
   "name":"envoy.filters.http.cors",
   "typedConfig":{
      "@type":"type.googleapis.com/envoy.extensions.filters.http.cors.v3.Cors"
   }
}
`,
			},
		},
		{
			Desc: "Generate with CORS is set to regex mode",
			OptsIn: options.ConfigGeneratorOptions{
				CorsPreset: "cors_with_regex",
			},
			WantFilterConfigs: []string{
				`
{
   "name":"envoy.filters.http.cors",
   "typedConfig":{
      "@type":"type.googleapis.com/envoy.extensions.filters.http.cors.v3.Cors"
   }
}
`,
			},
		},
		{
			Desc: "No-op when CORS preset is not enabled",
			OptsIn: options.ConfigGeneratorOptions{
				CorsPreset: "",
			},
			WantFilterConfigs: nil,
		},
	}

	for _, tc := range testdata {
		tc.RunTest(t, filtergen.NewCORSFilterGensFromOPConfig)
	}
}
