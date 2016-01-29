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
	var a []byte

	//st := reflect.TypeOf(snmpDev.GenericDevice{})
	//sv := reflect.ValueOf(snmpDev.GenericDevice{})
	st := reflect.TypeOf(snmpDev.LldpPort{})
	sv := reflect.ValueOf(snmpDev.LldpPort{})
	n := st.NumField()
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
		v, err := snmp.Walk("todclsp02b", "public", "lldpLocPortTable")
		if err != nil {
			continue
			fmt.Println(err)
		}

		for v.Next() {
			x, _ := v.Scan(&a)
			id := x.(int)

			spew.Dump(x)
			spew.Dump(id)
			spew.Dump(a)
		}
	}
}
