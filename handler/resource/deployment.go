package resource

import (
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"go-gin-web/pkg/common"
	"go-gin-web/pkg/errMsg"

	"github.com/gin-gonic/gin"
)

func ListDeployments(c *gin.Context) {
	var (
		// resObj  siginRes
		resData gin.H
		res     = common.Res{C: c}
	)

	var kubeconfig string = "/Users/liushaobo/.kube/config"

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 请求返回
	resData = gin.H{
		"namespaces": ns.Items,
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}
