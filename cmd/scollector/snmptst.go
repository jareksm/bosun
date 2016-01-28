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
	entPhysicalDescr      string `oid:"entPhysicalDescr" snmp:"octal"`
	entPhysicalVendorType string
	entPhysicalClass      int
}

func main() {
	dev := snmpDev.GenericDevice{Hardware: make(map[int]snmpDev.PhysHdw, 100)}
	spew.Dump(dev)
	var a []byte

	st := reflect.TypeOf(OIDs{})
	n := st.NumField()
	sv := reflect.ValueOf(OIDs{})
	oid := ""
	for i := 0; i < n; i++ {
		field := st.Field(i)
		switch sv.Field(i).Kind() {
		case reflect.Struct:
			continue
		}
		oid = field.Tag.Get("oid")
		if oid == "" {
			oid = field.Name
		}
		fmt.Println(oid)
		v, err := snmp.Walk("todclsp02b", "public", oid)
		fmt.Println(err)

		for v.Next() {
			x, err := v.Scan(&a)
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
