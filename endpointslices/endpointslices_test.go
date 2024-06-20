package endpointslices

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TestCase struct {
	desc      string
	podName   string
	nodeName  string
	ownerRefs []metav1.OwnerReference
	expected  string
}

func TestGetPodName(t *testing.T) {
	tests := []TestCase{
		{
			desc:     "with-no-owner-reference",
			nodeName: "worker-node",
			podName:  "kube-rbac-proxy-worker-node",
			expected: "kube-rbac-proxy",
		},
		{
			desc:     "with-owner-reference-kind-node",
			nodeName: "worker-node",
			podName:  "kube-rbac-proxy-worker-node",
			ownerRefs: []metav1.OwnerReference{
				{
					Kind: "Node",
				},
			},
			expected: "kube-rbac-proxy",
		},
		{
			desc: "with-owner-reference-kind-ReplicaSet",
			ownerRefs: []metav1.OwnerReference{
				{
					Kind: "ReplicaSet",
					Name: "kube-rbac-proxy-7b7df454c7",
				},
			},
			expected: "kube-rbac-proxy",
		},
		{
			desc: "with-owner-reference-kind-DaemonSet",
			ownerRefs: []metav1.OwnerReference{
				{
					Kind: "DaemonSet",
					Name: "kube-rbac-proxy",
				},
			},
			expected: "kube-rbac-proxy",
		},
		{
			desc: "with-owner-reference-kind-StatefulSet",
			ownerRefs: []metav1.OwnerReference{
				{
					Kind: "StatefulSet",
					Name: "kube-rbac-proxy",
				},
			},
			expected: "kube-rbac-proxy",
		},
	}
	for _, test := range tests {
		p := defineTestPod(&test)
		res, err := getPodName(p)
		if err != nil {
			t.Fatalf("test %s failed. got error: %s", test.desc, err)
		}
		if res != test.expected {
			t.Fatalf("test %s failed. expected %v got %v", test.desc, test.expected, res)
		}
	}
}

func defineTestPod(t *TestCase) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: t.podName, OwnerReferences: t.ownerRefs},
		Spec:       corev1.PodSpec{NodeName: t.nodeName},
	}
}
