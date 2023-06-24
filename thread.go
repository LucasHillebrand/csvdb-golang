package main

import (
	filesys "csvDB/FileSys"
	"fmt"
	"net"
	"time"
)

type threads struct {
	OpenConns []*net.Conn
}

func (itSelf *threads) add(new *net.Conn) {
	itSelf.OpenConns = growConnArr(itSelf.OpenConns, new)
}

func (itSelf *threads) isOpen() bool {
	out := false
	if len(itSelf.OpenConns) > 0 {
		out = true
	}
	return out
}

func (itSelf *threads) get() (out *net.Conn, succes bool) {
	if itSelf.isOpen() {
		succes = true
		out = itSelf.OpenConns[0]
		itSelf.OpenConns = removeConnArr(itSelf.OpenConns, 0)
	} else {
		succes = false
	}
	return
}

func thread(all *threads, upd *Update, master *Master, kA *KeepAlive, id int) {
	for !upd.stop {
		for !all.isOpen() {
			time.Sleep(time.Millisecond * 200)
		}
		pconn, suc := all.get()
		if suc {
			conn := *pconn
			fmt.Printf("%s connected to thread %d \n", conn.RemoteAddr(), id)
			lenbuf := make([]byte, 8)
			conn.Read(lenbuf)
			size := ByteToUint(lenbuf)
			buf := make([]byte, size)
			n, _ := conn.Read(buf)

			if n > 0 {
				kA.Add(1)
				str := ""
				for i := 0; i < n; i++ {
					str += string(buf[i])
				}
				SInp := filesys.Split(str, ";")
				args := make([]string, len(SInp)-1)
				for i := 1; i < len(SInp); i++ {
					args[i-1] = SInp[i]
				}

				out := Runtime(SInp[0], args, master, upd)
				outp := toBytes(out[1] + ";" + out[0])

				conn.Write(UintToBytes(uint64(len(outp))))
				conn.Write(outp)
				kA.Done()
			}
			conn.Close()
		}
	}
}
