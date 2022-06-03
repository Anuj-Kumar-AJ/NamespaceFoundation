// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubeobserver

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
)

func WaiterFromPodStatus(ns, name string, ps v1.PodStatus) ops.WaitStatus {
	if ps.Phase == v1.PodPending && len(ps.ContainerStatuses) == 0 {
		return pendingWaitStatus{}
	}

	cw := runtime.ContainerWaitStatus{}
	for _, container := range ps.ContainerStatuses {
		if lbl := containerStateLabel(&ps, container.State); lbl != "" {
			cw.Containers = append(cw.Containers, runtime.ContainerUnitWaitStatus{
				Reference:   kubedef.MakePodRef(ns, name, container.Name),
				Name:        container.Name,
				StatusLabel: lbl,
				Status:      StatusToDiagnostic(container),
			})
		}
	}

	for _, init := range ps.InitContainerStatuses {
		if lbl := containerStateLabel(nil, init.State); lbl != "" {
			cw.Initializers = append(cw.Initializers, runtime.ContainerUnitWaitStatus{
				Reference:   kubedef.MakePodRef(ns, name, init.Name),
				Name:        init.Name,
				StatusLabel: lbl,
				Status:      StatusToDiagnostic(init),
			})
		}
	}

	return cw
}

type pendingWaitStatus struct{}

func (pendingWaitStatus) WaitStatus() string { return "Pending..." }

func containerStateLabel(ps *v1.PodStatus, st v1.ContainerState) string {
	if st.Running != nil {
		label := "Running"
		if ps != nil {
			if !matchPodCondition(*ps, v1.PodReady) {
				label += " (not ready)"
			}
		}
		return label
	}

	if st.Waiting != nil {
		return st.Waiting.Reason
	}

	if st.Terminated != nil {
		if st.Terminated.ExitCode == 0 {
			return ""
		}
		return fmt.Sprintf("Terminated: %s (exit code %d)", st.Terminated.Reason, st.Terminated.ExitCode)
	}

	return "(Unknown)"
}

func podWaitingStatus(ctx context.Context, cli *k8s.Clientset, ns string, replicaset string) ([]ops.WaitStatus, error) {
	// TODO explore how to limit the list here (e.g. through labels or by using a different API)
	pods, err := cli.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var statuses []ops.WaitStatus
	for _, pod := range pods.Items {
		owned := false
		for _, owner := range pod.ObjectMeta.OwnerReferences {
			if owner.Name == replicaset {
				owned = true
			}
		}
		if !owned {
			continue
		}

		statuses = append(statuses, WaiterFromPodStatus(pod.Namespace, pod.Name, pod.Status))
	}

	return statuses, nil
}