package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	//	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"

	"k8s.io/client-go/rest"
)

var (

	//Set during build
	version string
	commit  string
	date    string

	versionFlag = flag.Bool("version", false, "Print the version, git-commit hash and build date and exit")

	original_ingress_name      = flag.String("original_ingress_name", "test", "Define the original ingress")
	new_ingress_class_name     = flag.String("new_ingress_classname", "nginx-plus", "Define the new ingress class")
	ingress_namespace_name     = flag.String("namespace_name", "default", "Define the namespace")
	new_ingress_namespace_name = flag.String("new_namespace_name", "default", "Define the namespace")
	new_ingress_name           = flag.String("new_ingress_name", "test-migration-edtion", "Define the new ingress name")
	//enable_megertable          = flag.Bool("enable_mgertable", false, "Define whether enable convert to Megertable Ingress Resource")
	//enable_crd                 = flag.Bool("enable_crd", false, "Define whether convert all ingress to Customer Resource Definition")
)

func main() {
	flag.Parse()

	versionInfo := fmt.Sprintf("Version=%v GitCommit=%v Date=%v Arch=%v/%v", version, commit, date, runtime.GOOS, runtime.GOARCH)
	if *versionFlag {
		fmt.Println(versionInfo)
		os.Exit(0)
	}

	//config, kubeClient := createConfigAndKubeClient()

	// var kubeconfig *string
	//ctx := context.Background()
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	klog.Fatal(err)
	// 	return
	// }
	//klog.Fatal("config = ", config)
	//create clientset

	var config *rest.Config
	ctx := context.Background()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return
	}

	//read config json into keyinfos map
	f, err := ioutil.ReadFile("./cfg/ingresskey.json")
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}

	keyinfos := make(map[string]string)

	errr := json.Unmarshal([]byte(f), &keyinfos)
	if errr != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	//ingress informer
	informerFactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	ingressInformer := informerFactory.Extensions().V1beta1().Ingresses()
	informer := ingressInformer.Informer()
	ingressLister := ingressInformer.Lister()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	stopper := make(chan struct{})
	defer close(stopper)
	// start Informer List & Watch
	informerFactory.Start(stopper)
	// wait till Informer cache sync
	informerFactory.WaitForCacheSync(stopper)
	ingressesInfom, err := ingressLister.Ingresses(*new_ingress_namespace_name).List(labels.Everything())
	// edit the Ingresses list
	for index, ingress := range ingressesInfom {
		fmt.Printf("%d -> ===== processing name %s ===== \n", index, ingress.Name)
		if ingress.Name == *original_ingress_name {
			newingress := ingress
			newingress.Name = *new_ingress_name
			newingress.Namespace = *new_ingress_namespace_name
			newingress.ResourceVersion = ""

			err_delete := clientset.ExtensionsV1beta1().Ingresses(*new_ingress_namespace_name).Delete(ctx, newingress.Name, metav1.DeleteOptions{})
			if err_delete == nil {
				fmt.Printf("Name %s already exist, delete\n", newingress.Name)
			}

			//replacing value
			fmt.Println("replacing value:")
			for key, value := range ingress.Annotations {
				fmt.Printf("replace value %s \n", key)
				switch key {
				case "nginx.ingress.kubernetes.io/upstream-hash-by":
					newingress.Annotations[key] = "hash " + value + " consistent"
				}
			}

			//replacing key
			fmt.Println("replacing key:")
			for key, value := range ingress.Annotations {
				replace_key, ok := keyinfos[key]
				if ok {
					fmt.Printf("replace key %s to %s\n", key, replace_key)
					delete(newingress.Annotations, key)
					newingress.Annotations[replace_key] = value
				}
			}

			fmt.Println("replacing ingressClassName")
			newingress.Spec.IngressClassName = new_ingress_class_name

			_, err_create := clientset.ExtensionsV1beta1().Ingresses("default").Create(ctx, newingress, metav1.CreateOptions{})
			if err_create != nil {
				fmt.Printf("Encounter error in creating newingress %s\n", err_create)
			}
		}
	}
	<-stopper
}

func onAdd(obj interface{}) {
	ingress := obj.(*v1beta1.Ingress)
	klog.Infoln("add a Ingress: ", ingress.Name)
}

func onUpdate(old, new interface{}) {
	oldIngress := old.(*v1beta1.Ingress)
	newIngress := new.(*v1beta1.Ingress)
	klog.Infoln("update deploy: ", oldIngress.Spec.IngressClassName, newIngress.Spec.IngressClassName)
}

func onDelete(obj interface{}) {
	ingress := obj.(*v1beta1.Ingress)
	klog.Infoln("delete a Ingress: ", ingress.Name)
}
