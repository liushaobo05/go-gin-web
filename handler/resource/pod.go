package resource

import (
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"go-gin-web/pkg/common"
	"go-gin-web/pkg/errMsg"

	"github.com/gin-gonic/gin"
)

func GetPodCount(c *gin.Context) {
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

	pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d nodes in the cluster\n", len(pods.Items))

	// 请求返回
	resData = gin.H{
		"podCount": len(pods.Items),
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}

func ListPods(c *gin.Context) {
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

	pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pod in the cluster\n", len(pods.Items))

	// 请求返回
	resData = gin.H{
		"pods": pods.Items,
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}
