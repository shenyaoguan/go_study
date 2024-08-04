package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3307)/casbin", true) // Your driver and data source.
	if err != nil {
		fmt.Errorf("%s", err)
	}
	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	added, err := e.AddGroupingPolicy("alice", "admin_1")
	fmt.Println(added)
	//filteredPolicy, _ := e.GetFilteredPolicy(0, "alice")
	//fmt.Println(filteredPolicy)
	if err != nil {
		// handle err
		fmt.Errorf("%s", err)
	}

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("permit")
	} else {
		// deny the request, show an error
		fmt.Println("deny")
	}

}
