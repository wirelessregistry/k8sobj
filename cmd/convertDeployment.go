// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"k8s.io/api/policy/v1beta1"

	//"github.com/nuance-mobility/zookeeper-operator/pkg/generated/clientset/versioned/scheme"

	//"github.com/wirelessregistry/signal-graph/backend/pkg/dep/sources/https---github.com-kubernetes-api/apps/v1beta1"

	//_ "k8s.io/client-go/pkg/api/install"
	//_ "k8s.io/client-go/pkg/apis/extensions/install"

	//_ "k8s.io/client-go/pkg/api/install"

	"k8s.io/api/apps/v1beta1"
	_ "k8s.io/client-go/kubernetes"

	//_ "k8s.io/client-go/pkg/apis/extensions/install"

	//"github.com/ghodss/yaml"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	//"github.com/wirelessregistry/signal-graph/backend/pkg/dep/sources/https---github.com-kubernetes-apimachinery/pkg/util/yaml"
)

// convertCmd represents the convert command
var convertDeploymentCmd = &cobra.Command{
	Use:   "convertDeployment",
	Short: "Convert a kubernetes Deployment yaml file to a golang struct.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No file name given.")
			return
		}
		fileContent, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Print(err)
		}

		jsonBytes, err := yaml.YAMLToJSON(fileContent)
		if err != nil {
			fmt.Printf("Error converting YAML to JSON")
		}
		// unmarshal the json into the kube struct
		var deployment = v1beta1.Deployment{}
		err = json.Unmarshal(jsonBytes, &deployment)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON")
		}
		fmt.Printf("%+v\n", deployment)
	},
}

func init() {
	var File string
	rootCmd.AddCommand(convertDeploymentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	convertDeploymentCmd.Flags().StringVarP(&File, "file", "f", "", "File location of Kubernetes yaml file (required)")
	convertDeploymentCmd.MarkPersistentFlagRequired("file")
}
