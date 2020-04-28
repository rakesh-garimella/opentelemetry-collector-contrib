// Copyright 2020 OpenTelemetry Authors
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

package regexp

// RegexpConfig represents the options for a NewRegexpFilterSet.
type Config struct {
	// CacheEnabled determines whether match results are LRU cached to make subsequent matches faster.
	// Cache size is unlimited unless CacheMaxNumEntries is also specified.
	CacheEnabled bool `mapstructure:"cacheenabled"`
	// CacheMaxNumEntries is the max number of entries of the LRU cache that stores match results.
	// CacheMaxNumEntries is ignored if CacheEnabled is false.
	CacheMaxNumEntries int `mapstructure:"cachemaxnumentries"`
}

// CreateFilterSet creates a regexp FilterSet from yaml config.
func CreateRegexpFilterSet(filters []string, cfg *Config) (*FilterSet, error) {
	opts := []Option{}
	if cfg != nil && cfg.CacheEnabled {
		opts = append(opts, WithCache(cfg.CacheMaxNumEntries))
	}

	return NewRegexpFilterSet(filters, opts...)
}
