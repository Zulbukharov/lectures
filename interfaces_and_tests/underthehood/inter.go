package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func drainAndClose(r io.ReadCloser) {
	if r == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, r)
	_ = r.Close()
}

func getCurrentData(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		drainAndClose(res.Body)
		return nil, fmt.Errorf("unexpected response status code: %d", res.StatusCode)
	}
	return res.Body, nil
}

func ex00() {
	getCurrentData("hello")
}

type Mather interface {
	Add(a, b int32) int32
	Sub(a, b int32) int32
}

type Adder struct {
	id int32
}

func (adder Adder) Add(a, b int32) int32 {
	return a + b
}

func (adder Adder) Sub(a, b int32) int32 {
	return a - b
}

// tab holds the address of an itab object, which embeds the
// datastructures that describe both the type of the interface as well as the type of the data it points to.
// data is a raw (i.e. unsafe) pointer to the value held by the interface.
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter uintptr // interfacetype, указатель на тип интерфейса
	_type uintptr // _type, указатель на тип
	hash  uint32
	_     [4]byte
	fun   [1]uintptr // указатель на первый метод реализации
}

// type interfacetype struct {
// 	typ     _type // _type
// 	pkgpath name
// 	mhdr    []imethod
// }

// type _type struct {
// 	size       uintptr
// 	ptrdata    uintptr // size of memory prefix holding all pointers
// 	hash       uint32
// 	tflag      uintptr
// 	align      uint8
// 	fieldalign uint8
// 	kind       uint8
// 	// alg        *typeAlg
// 	// gcdata stores the GC type data for the garbage collector.
// 	// If the KindGCProg bit is set in kind, gcdata is a GC program.
// 	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
// 	gcdata    *byte
// 	str       nameOff
// 	ptrToThis typeOff
// }

func ex01() {
	m := Mather(Adder{id: 6754})

	iface := (*iface)(unsafe.Pointer(&m))
	fmt.Printf("iface.tab.hash = %#x\n", iface.tab.hash) // 0x615f3d8a
	fmt.Printf("iface.data = %v\n", iface.data)
}

// type Stringer interface {
//     String() string
// }

// func ToString(any interface{}) string {
//     if v, ok := any.(Stringer); ok {
//         return v.String()
//     }
//     switch v := any.(type) {
//     case int:
//         return strconv.Itoa(v)
//     case float:
//         return strconv.Ftoa(v, 'g', -1)
//     }
//     return "???"
// }

type Bucks uint32

func (b Bucks) String() string {
	return fmt.Sprintf("%d bucks", b)
}

func ex02() {
	fmt.Println(Bucks(300))
}

func main() {
	ex00()
	ex01()
	ex02()
}
