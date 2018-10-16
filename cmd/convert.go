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
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	//"k8s.io/api/policy/v1beta1"

	//"github.com/nuance-mobility/zookeeper-operator/pkg/generated/clientset/versioned/scheme"

	//"github.com/wirelessregistry/signal-graph/backend/pkg/dep/sources/https---github.com-kubernetes-api/apps/v1beta1"

	//_ "k8s.io/client-go/pkg/api/install"
	//_ "k8s.io/client-go/pkg/apis/extensions/install"

	//_ "k8s.io/client-go/pkg/api/install"
	"k8s.io/api/apps/v1beta1"
	_ "k8s.io/client-go/kubernetes"
	api "k8s.io/client-go/kubernetes/scheme"

	//_ "k8s.io/client-go/pkg/apis/extensions/install"

	//"github.com/ghodss/yaml"

	"github.com/spf13/cobra"
	//"github.com/wirelessregistry/signal-graph/backend/pkg/dep/sources/https---github.com-kubernetes-apimachinery/pkg/util/yaml"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a kubernetes yaml file to a golang struct.",
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

		decode := api.Codecs.UniversalDeserializer().Decode
		obj, groupVersionKind, err := decode([]byte(fileContent), nil, nil)
		fmt.Println(groupVersionKind.Kind)

		if err != nil {
			log.Fatal(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
		}
		if obj == nil {
			log.Fatal(fmt.Sprintf("Runtime object is empty"))
		}
		//obj2, ok := obj.(*v1beta1.Deployment)
		//obj3 := *obj2
		objTypeString := reflect.TypeOf(obj).String()
		fmt.Println(objTypeString)
		switch objTypeString {
		case "*v1.Pod":
			// o is a pod
		case "*v1alpha1.Role":
			// o is the actual role Object with all fields etc
		case "*v1alpha1.RoleBinding":
		case "*v1alpha1.ClusterRole":
		case "*v1alpha1.ClusterRoleBinding":
		case "*v1.ServiceAccount":
			break
		case "*v1beta1.Deployment":
			fmt.Printf("DEPLOYMENT DETECTED")
			obj2, ok := obj.(*v1beta1.Deployment)
			if ok == false {
				fmt.Println("Type assertion is incorrect")
				break
			}
			fmt.Printf("%+v\n", obj2)
		default:
			//fmt.Printf("%+v\n", obj)
			//fmt.Printf("default")
			//fmt.Println(reflect.TypeOf(o))
			//o is unknown for us
		}
	},
}

func init() {
	var File string
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	convertCmd.Flags().StringVarP(&File, "file", "f", "", "File location of Kubernetes yaml file (required)")
	convertCmd.MarkPersistentFlagRequired("file")
}
