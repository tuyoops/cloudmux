// Copyright 2023 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/volcengine"
	"yunion.io/x/pkg/util/shellutils"
)

func init() {
	type NetworkInterfaceListOptions struct {
		InstanceId string `help:"Id or instance"`
		PageSize   int
		PageNumber int
	}
	shellutils.R(&NetworkInterfaceListOptions{}, "network-interface-list", "List networkinterfaces", func(cli *volcengine.SRegion, args *NetworkInterfaceListOptions) error {
		interfaces, _, err := cli.GetNetworkInterfaces(args.InstanceId, args.PageNumber, args.PageSize)
		if err != nil {
			return err
		}
		printList(interfaces, 0, 0, 0, nil)
		return nil
	})
}
