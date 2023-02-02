package main

import (
	"bytes"
	"path/filepath"
	"text/template"

	log "github.com/sirupsen/logrus"
	core_v1 "k8s.io/api/core/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a Pod object to pull out relevant data
	pod := obj.(*core_v1.Pod)
	log.Infof("    ResourceVersion: %s", pod.ObjectMeta.ResourceVersion)
	log.Infof("    NodeName: %s", pod.Spec.NodeName)
	log.Infof("    Phase: %s", pod.Status.Phase)

	isPodReady := false

	for _, condition := range pod.Status.Conditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			isPodReady = true
			break
		}
	}
	if isPodReady {
		for _, volume := range pod.Spec.Volumes {
			if volume.Secret != nil {
				for _, volumeMount := range pod.Spec.Containers[0].VolumeMounts {
					if volumeMount.Name == volume.Name {
						// log.Infof("Label: %v", pod.ObjectMeta.Labels["pod-template-hash"])
						// 	labels := pod.ObjectMeta.Labels
						// 	podTemplateHash := labels["pod-template-hash"]

						// 	keyValLabels := ""
						// 	for key, val := range labels {
						// 		if key == "pod-template-hash" {
						// 			continue
						// 		}
						// 		keyValLabels += "      " + key + ": " + val + "\n"
						// 	}
						// 	dataPolicy := map[string]string{
						// 		"PolicyName":      fmt.Sprintf("pod-%s-disable-secret-access", podTemplateHash),
						// 		"KeyValLabel":     keyValLabels,
						// 		"SecretMountPath": fmt.Sprintf("%s/", volumeMount.MountPath),
						// 	}
						// 	outPolicy, _ := generateKarmorPolicy("./example/template/k8s-secret-karmor.yaml", dataPolicy)
						// 	fileNameOut := "./generated/" + dataPolicy["PolicyName"] + ".yaml"
						// 	ioutil.WriteFile(fileNameOut, outPolicy, 0644)

						// 	time.Sleep(time.Minute) // replace with probiness check
						// 	cmd := exec.Command("kubectl", "apply", "-f", fileNameOut)
						// 	out, err := cmd.Output()

						// 	log.Infof("Info: %v", string(out))
						// 	if err != nil {
						// 		log.Errorf("Error: %v", err.Error())
						// 		return
						// 	}
					}
				}
			}
		}

		log.Infof("pod.Spec.Volumes.Secret: %v\n", pod.Spec.Volumes[0].Secret.SecretName)
		log.Infof("pod.Spec.Containers: %v\n", pod.Spec.Containers[0].VolumeMounts)
	}
}

func generateKarmorPolicy(filePath string, availableData map[string]string) ([]byte, error) {
	tmpl, err := template.New(filepath.Base(filePath)).ParseFiles(filePath)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, availableData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
