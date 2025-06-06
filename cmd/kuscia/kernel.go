// Copyright 2024 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// kernel.go defines the kernel-check subcommand.
// This command is used to check kernel parameters to ensure they meet the requirements.
// Use runtime.CurrentKernel().List() to get the current kernel parameter list and output detailed information for each parameter.
package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/secretflow/kuscia/pkg/utils/runtime"
)

func NewKernelCheckCommand(ctx context.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:          "kernel-check",
		Short:        "Check Kernel Params",
		Long:         `Check Kernel Params, make sure satisfy the requirements`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, param := range runtime.CurrentKernel().List() {
				println(fmt.Sprintf("Key=%s, Match(%v), Value(%s), Description(%s)", param.Key, param.IsMatch, param.Value, param.Description))
			}
			return nil
		},
	}

	return cmd
}
