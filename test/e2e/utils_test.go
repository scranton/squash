package e2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gokubeutils "github.com/solo-io/go-utils/kubeutils"
	"k8s.io/client-go/kubernetes"
)

func check(err error) {
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
}

// wrapper around check
// suggestion: use check(err) for functions that return multiple values
// suggestion: use must(myFunction()) for functions that return error only
func must(err error) {
	check(err)
}

func MustGetClientset() *kubernetes.Clientset {
	cs := &kubernetes.Clientset{}
	By("should get a kube client")
	restCfg, err := gokubeutils.GetConfig("", "")
	check(err)
	cs, err = kubernetes.NewForConfig(restCfg)
	check(err)
	return cs
}
