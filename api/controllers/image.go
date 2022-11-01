package controllers

import (

	"net/http"
	"github.com/gorilla/mux"
	"fmt"
    "log"
	"github.com/bndr/gojenkins"
	"context"
	"encoding/json"
	"io/ioutil"
	"github.com/OpenNebula/one/src/oca/go/src/goca"
	"time"
    "github.com/OpenNebula/one/src/oca/go/src/goca/schemas/shared"

    "github.com/OpenNebula/one/src/oca/go/src/goca/schemas/vm"

    "github.com/OpenNebula/one/src/oca/go/src/goca/schemas/vm/keys"
)

func opinit() *goca.Controller {
		 con := map[string]string{
			"user":     "oneadmin",
			"password": "Azerty123",
			"endpoint": "http://10.4.0.127:2633/RPC2",
		}
		client := goca.NewDefaultClient(
			goca.NewConfig(con["user"], con["password"], con["endpoint"]),
		)
		controller := goca.NewController(client)
		return controller
}


func JenkinsBuild(w http.ResponseWriter, r *http.Request){

	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://10.4.0.127:8080/", "admin", "azerty")
	_, err := jenkins.Init(ctx)
	body, err := ioutil.ReadAll(r.Body)

    params := make(map[string]string)

    newerr := json.Unmarshal(body, &params)

	if err != nil && newerr != nil{
	panic("Something Went Wrong")
	}
	job, err := jenkins.GetJob(ctx, "terraform")
	queueId, err := job.InvokeSimple(ctx, params)

	if err != nil && newerr != nil{
		panic("Something Went Wrong")
	}

	build, err := jenkins.GetBuildFromQueueID(ctx, queueId)
	if err != nil && newerr != nil{
		panic("Something Went Wrong")
	}

	for build.IsRunning(ctx) {

        time.Sleep(5000 * time.Millisecond)

        build.Poll(ctx)

    }

    var result string = build.GetResult()
	w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(result)
}

func CreateVM(w http.ResponseWriter, r *http.Request) {
	
	controller := opinit()
    tpl := vm.NewTemplate()
    tpl.Add(keys.Name, "testVM")
    tpl.CPU(1).Memory(64).VCPU(2)
    disk := tpl.AddDisk()
    disk.Add(shared.ImageID, "0")
    disk.Add(shared.DevPrefix, "vd")
    nic := tpl.AddNIC()
    nic.Add(shared.NetworkID, "0")
    nic.Add(shared.Model, "virtio")
    vmID, err := controller.VMs().Create(tpl.String(), false)
    if err != nil {
        log.Fatal(err)
    }
    vmCtrl := controller.VM(vmID)
    vm, err := vmCtrl.Info(false)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", vm)
    // Poweroff the VM
    // err = vmCtrl.Poweroff()
    // if err != nil {
    //     log.Fatal(err)
    // }
}


func GetInfos(w http.ResponseWriter, r *http.Request) {
	controller := opinit()
	vms, err := controller.VMs().Info()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(vms.VMs); i++ {
		
		vm, err := controller.VM(vms.VMs[i].ID).Info(false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", vm)
	}
}

func GetVM(w http.ResponseWriter, r *http.Request) () {
	controller := opinit()
	fmt.Println(r)

	vars := mux.Vars(r)
	key := vars["Name"]
	id, err := controller.VMs().ByName(key)
	if err != nil {
		log.Fatal(err)
	}
	vm, err := controller.VM(id).Info(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", vm)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vm)
}