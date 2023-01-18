package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubernetesClient() kubernetes.Interface {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("[Error] Create the kubeconfig: %v", err)
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("[Error] Create k8s client: %v", err)
	}

	log.Println("Successfully created k8s client")
	return k8sClient
}

func main() {
	client := getKubernetesClient()

	fmt.Print(client.AppsV1().RESTClient().APIVersion())
}
