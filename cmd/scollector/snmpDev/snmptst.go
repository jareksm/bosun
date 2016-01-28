package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime/pprof"

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
	var a interface{}
	f, _ := os.Create("snmptst.cpu")
	defer f.Close()
	fh, _ := os.Create("snmptst.heap")
	defer fh.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

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
			id, err := v.Scan(&a)
			if err != nil {
				fmt.Println(err)
				break
			}
			spew.Dump(id)
			spew.Dump(a)
		}
	}
	pprof.WriteHeapProfile(fh)
}
