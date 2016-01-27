/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2015 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// TODO - NAME OF THE PACKAGE SHOULD REFLECT COLLECTOR TYPE EG. ceph, haproxy, mysql
// TODO - COLLECTOR PACKAGE SHOULD ONLY CONTAIN METHODS AND STRUCTS DIRECTLY RELATED TO COLLECTOR ITSELF
package pkgname

// TODO - ORDERING OF IMPORTS
// TODO - 1. STANDARD LIBRARY IMPORTS EXAMPLE: "fmt", "os", "encoding/json"
// TODO - 2. 3RD PARTY LIBRARIES EXAMPLE: "github.com/opencontainers/runc/libcontainer/cgroups"
// TODO - 3. FRAMEWORK PACKAGES EXAMPLE: "github.com/intelsdi-x/snap/control/plugin"
// TODO - 4. PLUGIN RELATED PACKAGES
import (
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
)

// TODO - CONST SECTION SHOULD ALWAYS CONTAIN BELOW VALUES
// TODO - 1. PLUGIN NAME
// TODO - 2. PLUGIN VERSION
// TODO - 3. PLUGIN TYPE
// TODO - THESE CONSTANTS DON'T NEED TO BE EXPORTED, SO KEEP THEM LOWERCASED
const (
	name = "YOUR_COLLECTOR_NAME"
	version = 1
	plgtype = plugin.CollectorPluginType
)

// TODO - EVERY EXPORTED CONSTANT, VARIABLE, FUNCTION, STRUCT, METHOD OR INTERFACE SHOULD BE COMMENTED
// ImportantOutsidePackage is very important
var ImportantOutsidePackage = "bar"
// TODO - LIMIT THE NUMBER OF GLOBAL VARIABLES, ESPECIALLY THOSE WITH UPPER CASE
// TODO - IF GLOBAL VARIABLE IS NEEDED FOR SOME REASON, TRY TO KEEP IT PACKAGE LOCAL (LOWERCASED)
var importantForPackage = "foo"

// TODO - THERE'S NO NEED TO EXPORT RECEIVER FOR CollectorPlugin INTERFACE.
// TODO - CollectorPlugin INTERFACE RECEIVER SHOULD HAVE CONSTRUCTOR FUNCTION
// TODO - MOST OF INITIALIZATION FOR RECEIVER SHOULD BE DONE HERE
// TODO - (NOT EVERYTHING IS POSSIBLE HERE, FOR INSTANCE, PLUGIN CONFIG IS ACCESSIBLE AT THIS POINT)
// Commenting exported items is very important
func NewCollector() *collector {
	return &collector{}
}

// TODO - INTERFACE FUNCTIONS SHOULD BE CLEAN AND SIMPLE, SERVING ONLY IT'S DESIRED PURPOSE.
// TODO - AVOID ANY KIND OF INITIALIZATIONS HERE
// TODO - TRY TO GET ADVANTAGE OF ALREADY CHECKED SOLUTIONS
// TODO - USE HANDY FUNCTIONS FROM snap-plugin-utilities LIKE config, ns, pipelines AS NEEDED
// Commenting exported items is very important
func (p *collector) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mts := []plugin.PluginMetricType{}
	// Gather available metrics here
	return mts, nil
}

// TODO - IN CASE COLLECTING IS COSTLY:
// TODO - 1. TRY GO CONCURRENCY PATTERNS IF POSSIBLE
// TODO - 2. DO NOT COLLECT EVERYTHING IF POSSIBLE, COLLECT ONLY REQUIRED METRICS
// Commenting exported items is very important
func (p *collector) CollectMetrics(metricTypes []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	metrics := []plugin.PluginMetricType{}

	for _, metricType := range metricTypes {
		// Collect metrics here
	}

	return metrics, nil
}

// Commenting exported items is very important
func (p *collector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

//TODO - USE Meta FUNCTION TO RETURN COLLECTOR METADATA INSTEAD OF EXPORTING VARIABLES TO MAIN PACKAGE
// Commenting exported items is very important
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		name,
		version,
		plgtype,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
	)
}

// TODO - THERE'S NO NEED TO EXPORT RECEIVER FOR CollectorPlugin INTERFACE.
// TODO - USE INTERFACES FOR EASIER TESTING AND MOCKING
// TODO - USE DEPENDENCY INJECTION WHEN NEEDED TO ACHIEVE TESTABILITY
type collector struct {}


