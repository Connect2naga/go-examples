// Package k8s_client contains ...
package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

/*
Author : Nagarjuna S
Date : 2/7/22 11:35 AM
Project : go-examples
File : k8s_events_client.go
*/

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	isContinue:="eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6Mjk5MTksInN0YXJ0IjoiZGVmYXVsdC9pcC0xMC0wLTEyOC01Ni51cy13ZXN0LTEuY29tcHV0ZS5pbnRlcm5hbC4xNmQxNmNhZmNjODVhNWQ0XHUwMDAwIn0"
	for {
		pods, err := clientset.EventsV1().Events("external-dns").List(context.TODO(), metav1.ListOptions{
			Limit:                10,
			Continue:isContinue,
		})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d events in the cluster, continue is : %s\n", len(pods.Items), pods.Continue)
		fmt.Printf("Items : %+v", pods.Items)
		isContinue = pods.GetContinue()
		time.Sleep(10 * time.Second)
	}
}
