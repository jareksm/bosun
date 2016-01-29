package main

import (
	"fmt"
	"reflect"

	"bosun.org/cmd/scollector/snmpDev"
	"bosun.org/snmp"
	"github.com/davecgh/go-spew/spew"
)

type HW struct {
	entPhysicalSerialNum string
}

type OIDs struct {
	HW
	entPhysicalDescr      string `oid:"entPhysicalDescr"`
	entPhysicalVendorType string
	entPhysicalClass      int
}

func main() {

	st := reflect.TypeOf(snmpDev.PhysHdw{})
	sv := reflect.ValueOf(snmpDev.PhysHdw{})
	//st := reflect.TypeOf(snmpDev.LldpPort{})
	//sv := reflect.ValueOf(snmpDev.LldpPort{})
	n := st.NumField()
	oid := ""
	for i := 0; i < n; i++ {
		field := st.Field(i)
		kind := sv.Field(i).Kind()
		switch kind {
		case reflect.Map, reflect.Slice:
			continue
		case reflect.Struct:
			fmt.Println("Field is a struct!")
			spew.Dump(sv.Field(i).Kind())
			continue
		}
		oid = field.Tag.Get("oid")
		if oid == "" {
			oid = field.Name
			continue
		}
		fmt.Println(oid)
		v, err := snmp.Walk("todclsp02b", "public", oid)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("spewing V!")
		spew.Dump(v)
		for v.Next() {
			var a []byte
			x, _ := v.Scan(&a)
			if x != nil && a != nil && len(a) > 0 {
				fmt.Println("spewing X!")
				spew.Dump(x)
				fmt.Println("spewing A!")
				spew.Dump(a)
			}
		}
	}
}
