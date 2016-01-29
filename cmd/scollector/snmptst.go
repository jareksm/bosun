package main

import (
	"fmt"
	"reflect"

	"strconv"

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
	dev := snmpDev.PhysHdw{}
	spew.Dump(dev)
	var a []byte

	st := reflect.TypeOf(snmpDev.GenericDevice{})
	n := st.NumField()
	sv := reflect.ValueOf(snmpDev.GenericDevice{})
	oid := ""
	for i := 0; i < n; i++ {
		field := st.Field(i)
		kind := sv.Field(i).Kind()
		switch kind {
		case reflect.Map, reflect.Slice:
			continue
		case reflect.Struct:
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
			continue
			fmt.Println(err)
		}

		for v.Next() {
			x, err := v.Scan(&a)
			if a == nil || x == nil {
				continue
			}
			id := x.(int)

			if err != nil {
				fmt.Println(err)
				break
			}
			sph := snmpDev.PhysHdw{}
			switch oid {
			case "entPhysicalDescr":
				sph.Desc = string(a)
			case "entPhysicalVendorType":
				sph.Vendor = string(a)
			case "entPhysicalClass":
				num, _ := strconv.Atoi(string(a))
				sph.Class = snmpDev.PhysClass(num)
			}
			dev.Hardware[id] = sph

		}
		spew.Dump(dev)
	}
}
