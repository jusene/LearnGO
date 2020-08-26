package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var KUBEConfig = flag.StringP("config", "c", "./kubeconfig", "kubernetes kubeconfig")

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", *KUBEConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 根据指定config创建新的clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientSet.CoreV1().Pods("arch-pre").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for index, pod := range pods.Items {
		fmt.Println(index, pod.Name)
	}

	deploys, err := clientSet.AppsV1().Deployments("arch-pre").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for index, deploy := range deploys.Items {
		fmt.Println(index, deploy.Name)
	}

	services, err := clientSet.CoreV1().Services("arch-pre").List(metav1.ListOptions{})
	for index, service := range services.Items {
		fmt.Println(index, service.Name)
	}

	clientSet.CoreV1().Services("arch-pre").Create(&v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.ServiceSpec{},
		Status:     v1.ServiceStatus{},
	})
}
