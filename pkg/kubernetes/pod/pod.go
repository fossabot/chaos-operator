package pod

import (
	corev1 "k8s.io/api/core/v1"
)

// Pod holds the api's pod objects
type Pod struct {
	object *corev1.Pod
}

// PodList holds the list of API pod instances
type PodList struct {
	items []*Pod
}
