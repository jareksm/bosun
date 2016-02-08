package nlink

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"net"
	"syscall"
	"testing"
	"unsafe"

	"github.com/hkwi/nlgo"
)

func BenchmarkNetlink(b *testing.B) {
	var hub *nlgo.RtHub
	var err error
	if hub, err = nlgo.NewRtHub(); err != nil {
		//fmt.Println(os.NewSyscallError("netlinkrib", err))
	}
	defer hub.Close()
	req := syscall.NetlinkMessage{
		Header: syscall.NlMsghdr{
			Type:  syscall.RTM_GETLINK,
			Flags: syscall.NLM_F_DUMP,
		},
	}
	(*nlgo.IfInfoMessage)(&req).Set(
		syscall.IfInfomsg{
			Index: int32(1),
		},
		nlgo.AttrSlice{
			nlgo.Attr{
				Header: syscall.NlAttr{
					Type: syscall.IFLA_IFNAME,
				},
				Value: nlgo.NulString("em1"),
			},
		},
	)
	//for {
	if res, err := hub.Sync(req); err != nil {
		//fmt.Println(err)
	} else {

		b.ResetTimer()
		for z := 0; z < b.N; z++ {
		loop:
			for _, r := range res {
				switch r.Header.Type {
				case syscall.RTM_NEWLINK:
					msg := nlgo.IfInfoMessage(r)
					if msg.IfInfo().Index != int32(2) {
						//pass
					}
					attrs, _ := msg.Attrs()
					switch attrs.(type) {
					case nlgo.AttrMap:
						stat := attrs.(nlgo.AttrMap).Get(nlgo.IFLA_STATS64).(nlgo.Binary)
						_ = (*nlgo.RtnlLinkStats64)(unsafe.Pointer(&stat[0]))

						_ = string(attrs.(nlgo.AttrMap).Get(nlgo.IFLA_IFNAME).(nlgo.NulString))
						_ = (net.HardwareAddr)([]byte(attrs.(nlgo.AttrMap).Get(nlgo.IFLA_ADDRESS).(nlgo.Binary)))
						//fmt.Println(i, mac, s)

						linkinfo := attrs.(nlgo.AttrMap).Get(nlgo.IFLA_LINKINFO)
						var linki nlgo.AttrMap
						switch linkinfo.(type) {
						case nlgo.AttrMap:
							linki = linkinfo.(nlgo.AttrMap)
							_ = string(linki.Get(nlgo.IFLA_INFO_KIND).(nlgo.NulString))
						}
					}
				case syscall.NLMSG_DONE:
					break loop
				}
			}
		}
	}
}

var ifstatRE = regexp.MustCompile(`\s+(enp2s0|docker0|eth\d+|em\d+_\d+/\d+|em\d+_\d+|em\d+|` +
        `bond\d+|team\d+|` + `p\d+p\d+_\d+/\d+|p\d+p\d+_\d+|p\d+p\d+):(.*)`)

func readLine(fname string, line func(string) error) error {
        f, err := os.Open(fname)
        if err != nil {
                return err
        }
        defer f.Close()
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
                if err := line(scanner.Text()); err != nil {
                        return err
                }
        }
        return scanner.Err()
}

func BenchmarkProcfs(b *testing.B) {
	for j := 0; j < b.N; j++ {
        readLine("/proc/net/dev", func(s string) error {
                m := ifstatRE.FindStringSubmatch(s)
                if m == nil {
                        return nil
                }
			fmt.Println(s)
                intf := m[1]
				fmt.Println(m[1])
				fmt.Println(m[2])
                stats := strings.Fields(m[2])
                if strings.HasPrefix(intf, "bond") || strings.HasPrefix(intf, "team") {
                }
                // Detect speed of the interface in question
                _ = readLine("/sys/class/net/"+intf+"/speed", func(speed string) error {
                        return nil
                })
				fmt.Println(stats)
                for i, _ := range stats {
                        if i < 4 || (i >= 8 && i < 12) {
                        }
                }
                return nil
        })
	}
}
