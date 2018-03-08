
package main

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
)


var devices     = make(map[string]*pluginapi.Device)
var deviceFiles = make(map[string]string)
var devs []*pluginapi.Device

var devName = "cambricon"
var devPath = "/dev"

func check(err error) {
        if err != nil {
                log.Panicln("Fatal:", err)
        }
}

func ListDir(dirPth string, suffix string) (files []string, err error) {
        files = make([]string, 0, 10)
        dir, err := ioutil.ReadDir(dirPth)
        if err != nil {
                return nil, err
        }
        PthSep := string(os.PathSeparator)
        for _, fi := range dir {
                if fi.IsDir() { // ignore dirs
                        continue
                }
                if strings.Contains(fi.Name(), suffix) { //match files
                        files = append(files, dirPth+PthSep+fi.Name())
                }
        }

        return files, nil
}

func doesExist(devpath string) bool {

        result := false
        for _, v := range deviceFiles{
                if 0 == strings.Compare(v, devpath){
                        result = true
                        break
                }
        }

        return  result
}


func getDevices() []*pluginapi.Device {

	fmt.Printf("Discover Cambricon Card Resourcesi\n")

        camCards, err := ListDir(devPath, devName)

        if err != nil {
                fmt.Printf("Error while discovering: %v\n", err)
        }

        for _, card := range camCards {
                //fmt.Printf("detect card is: %s \n", card)
                //fmt.Printf("check result is %s\n", doesExist(card))
                if doesExist(card) == false {
                        //fmt.Printf("devicefiles %s", deviceFiles)
                        u1 := uuid.Must(uuid.NewV4())
                        fmt.Printf("Creating UUID for %s : %s\n", card, u1)
                        out := fmt.Sprint(u1)
                        dev := pluginapi.Device{ID: out, Health: pluginapi.Healthy}
			//fmt.Printf("Before Devices: %v \n", devs)
			devs = append(devs, &dev)
			//fmt.Printf("After Devices: %v \n", devs)
                        devices[out] = &dev
                        deviceFiles[out] = card
                        fmt.Printf("devs %s\n", devs)
                }
        }

        fmt.Printf("Devices: %v \n", devs)
        //if len(cam.deviceFiles) > 0{
        //        found = true
        //}


	return devs
}

func deviceExists(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			return true
		}
	}
	return false
}

func watchXIDs(ctx context.Context, devs []*pluginapi.Device, xids chan<- *pluginapi.Device) {

	//for _, d := range devs {
		//if err != nil && strings.HasSuffix(err.Error(), "Not Supported") {
		//	log.Printf("Warning: Cambricon Card with UUID %s is too old to support healtchecking with error: %s. Marking it unhealthy.", d.ID)

		//	xids <- d
		//	continue
		//}

		//if err != nil {
		//	log.Panicln("Fatal:", err)
		//}
	//}

	//for {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	default:
	//	}


	//	//for _, d := range devs {
	//		//if d.ID == *e.UUID {
	//	//		xids <- d
	//		//}
	//	//}
	//}
}
