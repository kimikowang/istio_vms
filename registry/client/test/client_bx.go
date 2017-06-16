package main

import(
//	"github.com/amalgam8/amalgam8/registry/client/test/platform/vms"
	"github.com/amalgam8/amalgam8/registry/client"
	"github.com/Sirupsen/logrus"
	"github.com/amalgam8/amalgam8/pkg/api"
	"fmt"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	config := client.Config {
		URL: "http://a8registry-server.mybluemix.net:80",
	}

	c, err := client.New(config)

	if err != nil {
		logrus.WithError(err).Error("Failure creating config")
	}

//	controller := vms.NewController(c, nil)

//	if res, err := c.GetServiceObject("details"); err == nil {
//		for i, item := range res {
//			fmt.Println(i, ": ", item.ServiceName)
//			fmt.Println(item.Address)
//			fmt.Println(item.ExternalName)
//			for _, p := range item.Ports {
//				fmt.Println("port name: ", p.Name)
//				fmt.Println("port num: ", p.Port)
//			}
//		}
//	} else {
//		fmt.Println("Result is nil with error: ", err)
//	}

	if res, err := c.ListServices(); err == nil {
		fmt.Println(res)
	}
//	if res, err := c.ListServiceInstances("productpage"); err == nil {
//		fmt.Println(res)
//		for i, ins := range res {
//			fmt.Println(i)
///			fmt.Println(ins.Service.ServiceName)
//			fmt.Println(ins.Service.Address)
//			fmt.Println(ins.Endpoint.ServicePort.Port)
//			fmt.Println(ins.Endpoint.ServicePort.Protocol)
//		}
//	}
}

func PrintInstance(ins *api.ServiceInstance) {
	logrus.WithFields(logrus.Fields{
		"service": ins.Service.ServiceName,
		"address": ins.Service.Address,
		"service_port_num": ins.Endpoint.ServicePort.Port,
		"service_port_proto": ins.Endpoint.ServicePort.Protocol,
	}).Info("")
}
