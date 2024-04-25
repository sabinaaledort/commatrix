package commatrix

import "github.com/openshift-kni/commatrix/types"

var generalStaticEntriesWorker = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      22,
		NodeRole:  "worker",
		Service:   "sshd",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9637,
		NodeRole:  "worker",
		Service:   "kube-rbac-proxy-crio",
		Namespace: "openshift-machine-config-operator",
		Pod:       "kube-rbac-proxy-crio",
		Container: "kube-rbac-proxy-crio",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10250,
		NodeRole:  "worker",
		Service:   "kubelet",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9107,
		NodeRole:  "worker",
		Service:   "egressip-node-healthcheck",
		Namespace: "openshift-ovn-kubernetes",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      111,
		NodeRole:  "worker",
		Service:   "rpcbind",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "UDP",
		Port:      111,
		NodeRole:  "worker",
		Service:   "rpcbind",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10256,
		NodeRole:  "worker",
		Service:   "ovnkube",
		Namespace: "openshift-sdn",
		Pod:       "ovnkube",
		Container: "ovnkube-controller",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9001,
		NodeRole:  "worker",
		Service:   "machine-config-daemon",
		Namespace: "openshift-machine-config-operator",
		Pod:       "machine-config-daemon",
		Container: "kube-rbac-proxy",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9537,
		NodeRole:  "worker",
		Service:   "crio-metrics",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	},
}

var generalStaticEntriesMaster = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9637,
		NodeRole:  "master",
		Service:   "kube-rbac-proxy-crio",
		Namespace: "openshift-machine-config-operator",
		Pod:       "kube-rbac-proxy-crio",
		Container: "kube-rbac-proxy-crio",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10256,
		NodeRole:  "master",
		Service:   "openshift-sdn",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9537,
		NodeRole:  "master",
		Service:   "crio-metrics",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10250,
		NodeRole:  "master",
		Service:   "kubelet",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9107,
		NodeRole:  "master",
		Service:   "egressip-node-healthcheck",
		Namespace: "openshift-ovn-kubernetes",
		Pod:       "ovnkube",
		Container: "ovnkube-controller",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      111,
		NodeRole:  "master",
		Service:   "rpcbind",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "UDP",
		Port:      111,
		NodeRole:  "master",
		Service:   "rpcbind",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      22,
		NodeRole:  "master",
		Service:   "sshd",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9192,
		NodeRole:  "master",
		Service:   "machine-approver",
		Namespace: "openshift-cluster-machine-approver",
		Pod:       "machine-approver",
		Container: "kube-rbac-proxy",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9258,
		NodeRole:  "master",
		Service:   "machine-approver",
		Namespace: "openshift-cloud-controller-manager-operator",
		Pod:       "cluster-cloud-controller-manager",
		Container: "cluster-cloud-controller-manager",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9099,
		NodeRole:  "master",
		Service:   "cluster-version-operator",
		Namespace: "openshift-cluster-version",
		Pod:       "cluster-version-operator",
		Container: "cluster-version-operator",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9980,
		NodeRole:  "master",
		Service:   "etcd",
		Namespace: "openshift-etcd",
		Pod:       "etcd",
		Container: "etcd",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9979,
		NodeRole:  "master",
		Service:   "etcd",
		Namespace: "openshift-etcd",
		Pod:       "etcd",
		Container: "etcd-metrics",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9978,
		NodeRole:  "master",
		Service:   "etcd",
		Namespace: "openshift-etcd",
		Pod:       "etcd",
		Container: "etcd-metrics",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10357,
		NodeRole:  "master",
		Service:   "openshift-kube-apiserver-healthz",
		Namespace: "openshift-kube-apiserver",
		Pod:       "kube-apiserver",
		Container: "kube-apiserver-check-endpoints",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      17697,
		NodeRole:  "master",
		Service:   "openshift-kube-apiserver-healthz",
		Namespace: "openshift-kube-apiserver",
		Pod:       "kube-apiserver",
		Container: "kube-apiserver-check-endpoints",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      2380,
		NodeRole:  "master",
		Service:   "healthz",
		Namespace: "openshift-etcd",
		Pod:       "etcd",
		Container: "etcd",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      2379,
		NodeRole:  "master",
		Service:   "etcd",
		Namespace: "openshift-etcd",
		Pod:       "etcd",
		Container: "etcdctl",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      6080,
		NodeRole:  "master",
		Service:   "",
		Namespace: "openshift-kube-apiserver-readyz",
		Pod:       "kube-apiserver",
		Container: "kube-apiserver-insecure-readyz",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      22624,
		NodeRole:  "master",
		Service:   "machine-config-server",
		Namespace: "openshift-machine-config-operator",
		Pod:       "machine-config-server",
		Container: "machine-config-server",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      22623,
		NodeRole:  "master",
		Service:   "machine-config-server",
		Namespace: "openshift-machine-config-operator",
		Pod:       "machine-config-server",
		Container: "machine-config-server",
		Optional:  false,
	},
}

var baremetalStaticEntriesWorker = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      53,
		NodeRole:  "worker",
		Service:   "dns-default",
		Namespace: "openshift-dns",
		Pod:       "dnf-default",
		Container: "dns",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "UDP",
		Port:      53,
		NodeRole:  "worker",
		Service:   "dns-default",
		Namespace: "openshift-dns",
		Pod:       "dnf-default",
		Container: "dns",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      18080,
		NodeRole:  "worker",
		Service:   "openshift-kni-infra-coredns",
		Namespace: "openshift-kni-infra",
		Pod:       "coredns",
		Container: "coredns",
		Optional:  false,
	},
}

var baremetalStaticEntriesMaster = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      53,
		NodeRole:  "master",
		Service:   "dns-default",
		Namespace: "openshift-dns",
		Pod:       "dnf-default",
		Container: "dns",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "UDP",
		Port:      53,
		NodeRole:  "master",
		Service:   "dns-default",
		Namespace: "openshift-dns",
		Pod:       "dnf-default",
		Container: "dns",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      5050,
		NodeRole:  "master",
		Service:   "metal3",
		Namespace: "openshift-machine-api",
		Pod:       "ironic-proxy",
		Container: "ironic-proxy",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9444,
		NodeRole:  "master",
		Service:   "openshift-kni-infra-haproxy-haproxy",
		Namespace: "openshift-kni-infra",
		Pod:       "haproxy",
		Container: "haproxy",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9445,
		NodeRole:  "master",
		Service:   "haproxy-openshift-dsn-internal-loadbalancer",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9191,
		NodeRole:  "master",
		Service:   "machine-approver",
		Namespace: "machine-approver",
		Pod:       "machine-approver",
		Container: "machine-approver-controller",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      6385,
		NodeRole:  "master",
		Service:   "",
		Namespace: "openshift-machine-api",
		Pod:       "ironic-proxy",
		Container: "ironic-proxy",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      29445,
		NodeRole:  "master",
		Service:   "haproxy-openshift-dsn",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  true,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      18080,
		NodeRole:  "master",
		Service:   "openshift-kni-infra-coredns",
		Namespace: "openshift-kni-infra",
		Pod:       "coredns",
		Container: "coredns",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      9447,
		NodeRole:  "master",
		Service:   "baremetal-operator-webhook-baremetal provisioning",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	},
}

var awsCloudStaticEntriesWorker = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10304,
		NodeRole:  "worker",
		Service:   "csi-node-driver",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10300,
		NodeRole:  "worker",
		Service:   "csi-livenessprobe",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	},
}

var awsCloudStaticEntriesMaster = []types.ComDetails{
	{
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      8080,
		NodeRole:  "master",
		Service:   "cluster-network",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10260,
		NodeRole:  "master",
		Service:   "aws-cloud-controller",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10258,
		NodeRole:  "master",
		Service:   "aws-cloud-controller",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10304,
		NodeRole:  "master",
		Service:   "csi-node-driver",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "Ingress",
		Protocol:  "TCP",
		Port:      10300,
		NodeRole:  "master",
		Service:   "csi-livenessprobe",
		Namespace: "",
		Pod:       "",
		Container: "",
		Optional:  false,
	},
}

var MNOStaticEntries = []types.ComDetails{
	{
		Direction: "ingress",
		Protocol:  "UDP",
		Port:      6081,
		NodeRole:  "worker",
		Service:   "ovn-kubernetes geneve",
		Namespace: "openshift-ovn-kubernetes",
		Pod:       "",
		Container: "",
		Optional:  false,
	}, {
		Direction: "ingress",
		Protocol:  "UDP",
		Port:      6081,
		NodeRole:  "master",
		Service:   "ovn-kubernetes geneve",
		Namespace: "openshift-ovn-kubernetes",
		Pod:       "",
		Container: "",
		Optional:  false,
	},
}
