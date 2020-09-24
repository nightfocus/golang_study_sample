// main project main.go
package main

import (
	//"apiclient"
	"container/list"
	"ctoola"
	"encoding/binary"
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"sync"
	"tasks"
	"time"

	//"myGin"

	//	"time"

	//"io"
	//"net/http"
	//"os"
	"strings"
	// "time"
	"os"
	"os/signal"
	"syscall"
	// "github.com/gin-gonic/gin"
)

func modifyList(lst *list.List) {
	lst.PushBack(int(220))
}

type Datam struct {
	id int
	b  *byte
	s  string
}
type Datam2 struct {
	id int
	bb []byte
	s  string
}

type NullStatic struct {
}

func (ns *NullStatic) sFunc() int {
	return 2020
}

type User struct {
	id   int
	name string
}

// 泛型的类型检查
// 判断传入的interface{} 是不是一个*User 指针类型
// 如果要判断是不是User 值类型，直接改为  anything.(User) 即可
func checkInterface(anything interface{}) bool {
	if _, ok := anything.(*User); ok {
		fmt.Println("Yes, it's a User type.")
		return true
	} else {
		fmt.Println("No. it's not a User type.")
		return false
	}
}

type Ier interface {
	GetName() string
}

type Humer struct {
}

func (h *Humer) GetName() string {
	return "IamHumer"
}

type I interface {
	Name()
}
type S struct {
}

func (*S) Name() {
}

const (
	Mon string = "Monday"
	Tue string = "Tuesday"
)

func loopRun() {
	for {
		tid := 1 // ctoola.GetThreadID() // linux
		fmt.Println("this tid: ", tid)
		// time.Sleep(1 * time.Second)
	}
}

// encoding/binary 用法
func bn() {
	var a int32
	p := &a
	b := [10]int64{1}
	s := "adsa"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  // 4
	fmt.Println(binary.Size(p))  // 4
	fmt.Println(binary.Size(b))  // 80
	fmt.Println(binary.Size(s))  // -1
	fmt.Println(binary.Size(bs)) // 10

	// 将数字 0x12345678，以小字节序存入到 []byte中，然后输出验证
	bs1 := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs1, 0x12345678)
	for _, x := range bs1 {
		fmt.Printf("%02X", x)
	}
	fmt.Printf("\n")

}

//*
type Test1 struct {
	Name string
	Pwg  *sync.WaitGroup
}

// 模拟一个执行时长是随机的任务
func (t *Test1) Exec() error {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println("Here is:", t.Name)
	ss := r.Intn(100) + 1
	time.Sleep(time.Duration(ss) * time.Millisecond)
	fmt.Printf("After sleep. Here is:%s, delay %d ms.\n", t.Name, ss)

	t.Pwg.Done()
	return nil
}

// 创建10万个任务，每个任务间隔1毫秒（不间隔也可以，但在循环投递的几秒内，CPU占满）
// 投递给GJobQueue去异步执行
func testTasks() {
	total := 99999
	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		wg.Add(1)
		s := fmt.Sprintf("Is%d", i)

		t := Test1{Name: s, Pwg: &wg}
		work := &t
		// fmt.Println("deliver ", t.Name)
		tasks.GJobQueue <- work
		time.Sleep(1 * time.Millisecond)
		// fmt.Println("deliver is finished. ", t.Name)
	}

	fmt.Println("deliver is completed, total: ", total)
	// 等待所有任务执行完成
	// 如果不等待，那就不调用 Wait()
	wg.Wait()
	fmt.Println("all task is completed, total: ", total)
}

// */

func main() {
	// 设置可使用的最大CPU核数，设为1，那么这个程序的CPU占用最多为100%
	// 在Linux实测，设置这个值和程序启动后的工作线程数无关
	runtime.GOMAXPROCS(4) // 设定 GMP 概念里 P 的数量
	// 设置允许的最大工作线程数。在4核CPU的Linux上，这个值不能低于5
	// 返回原来的值
	oldv := debug.SetMaxThreads(16) // 设定 GMP 概念里 M 的数量

	fmt.Println("------------------------ Begin... maxthreads:", oldv)
	//bn()

	// 一个调用face++的最简单示例
	// apiclient.PostMegFacepp()

	// 用来结合GOMAXPROCS(), SetMaxThreads() 测试协程和线程的关系.
	// go loopRun()

	/* 测试限制并发协程数
	ngl := NewGoLimit(10) // 表示最多允许10协程并发
	for ic := 0; ic < 100; ic++ {
		ngl.Add() // 登记一个协程，对应的用 ngl.Done 释放一个协程
		go echoClientLimited("101.200.188.59:20206", 10, ngl)
	}
	*/

	// 测试tasks包里的并发处理系统
	// go testTasks()

	// 阻塞9999秒后继续.
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func(wg1 *sync.WaitGroup) {
		// time.Sleep(9999 * time.Second)
		wg1.Done()
	}(&wg1)
	wg1.Wait()

	/*
		ps := new(string)  // ps是一个指向string类型的指针
		pps := ps
		*ps = "abcd"
		*pps = "xyzxyz"
		fmt.Println(*ps) // 会输出xyzxyz
	*/

	// testPanic()

	value := S{}
	value.Name() //可以调用
	var point = &value
	point.Name() //可以调用

	h := Humer{}
	var ph = &h
	fmt.Println(ph.GetName())
	/*
		for _, v := range h {
			fmt.Println(v.GetName())
		}
	*/

	// 测试最简单的context
	TestContext()

	/*
		for {
			go echoClient("39.107.105.159:9999", 9999)
			break
		}
	*/

	//echoServer()
	//time.Sleep(24 * time.Hour)
	//return

	// mge := myGin.NewMyGin()
	/*
		mge.Ge.GET("/hello", func(context *gin.Context) {
			// log.Println(">>>> hello gin start <<<<")
			context.JSON(200, gin.H{
				"code":    200,
				"success": true,
			})
		})
	*/
	// mge.Ge.GET("/hello", myGin.Handle(myGin.HandleHello))

	// go testPanic2()
	// mge.Run() // default :8080

	five := (int32)(50)
	fmt.Printf("%d\n", SquareFunc(&five)) // 传指针类型到函数

	cstr := ctoola.Toola("aabb")
	fmt.Printf("ctoola.Toola() return : %s\n", cstr)

	// 仿class的调用.
	pmi2 := new(ctoola.MyInfo)
	mi8 := ctoola.MyInfo{}
	pmi2.SetAge(40)
	fmt.Printf("pmi2 age : %d\n", pmi2.GetAge())
	mi8.SetAge(80)
	fmt.Printf("mi8 age : %d\n", mi8.GetAge())

	//* 多态性示例
	name := "Lee"
	pa := A{}
	pa.Hello(name) //hello Lee, i am a
	pa.Hello2("Yu")

	var ih IHello = &pa
	ih.Hello2("Coo")

	pb := B{&pa}
	pb.Hello(name) //hello Lee, i am a
	//*/
	fmt.Println("1=======================================")

	// 只要传入的类型，实现了 MyCallbacker 的所有接口，就可以传递.
	testCallback(&Hello{})

	// 链表写
	lst := list.New()
	lst.PushBack(int(3))
	lst.PushBack(float64(4.3))
	modifyList(lst)
	fmt.Printf("len lst:%d\n", lst.Len())
	fmt.Println("2=======================================")

	// 像操作C一样直接偏移内存操作和void*类似的转换
	unsafePointOper()
	fmt.Println("3=======================================")
	voidPointOper()
	fmt.Println("4=======================================")

	/* 循环标签
	   L1:
	   	for x := 0; x < 3; x++ {
	   		fmt.Printf("x:%d\n", x)
	   	L2:
	   		for y := 0; y < 5; y++ {
	   			if y > 2 {
	   				continue L2
	   			}

	   			if x > 1 {
	   				break L1
	   			}

	   			fmt.Println(x, ":", y, " ")
	   		}
	   		println()
	   	}
	//*/

	dm := Datam{}
	fmt.Println(dm) // 都会是初始值
	dm.id = 5
	onebyte := (byte)('c')
	dm.b = &onebyte
	dm.s = Mon
	fmt.Printf("Datam:%v, %c\n", dm, *dm.b)

	dm2 := Datam2{}
	bb := make([]byte, 2)
	dm2.bb = bb     // 两个bb会共用同一块内存
	bb[0] = 'a'     // a ascii is 97
	bb[1] = 'B'     // B ascii is 66
	dm2.bb[0] = 'A' // 会覆盖前面的赋值
	fmt.Printf("Datams2:%v, %s\n", dm2, dm2.bb)

	fmt.Println("5=======================================")
	var pf Platform = Dos
	fmt.Println("this is ", pf, pf.ShowText())
	pfLst := pf.List()
	for _, v := range pfLst {
		fmt.Println("foreach Platform:", v.Key, v.Val)
	}

	// 仿 静态类
	ns := NullStatic{}
	fmt.Println("ns.sFunc():", ns.sFunc())
	fmt.Println("6=======================================")

	chanSample()
	fmt.Println("8=======================================")

	// 通过interface来隐藏一个实现的内部细节
	mw := NewWidget()
	mwid := mw.GetId()
	fmt.Println("mwid:", mwid)
	fmt.Println("9=======================================")

	// 一个简单的http访问请求
	/*
		resp, err := http.Get("https://www.google.com")
		if err != nil {
			return
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	//*/

	// 这个类似将对象1传入给一个interface，然后在interface调用对象1的方法
	// 但这个方法在interface内是没有定义的.
	thisisaa := AA{}
	Callit(thisisaa)

	fmt.Println("\n\n\n11=======================================")
	chanSyncSample()

	fmt.Println("12=======================================")
	// 类型安全检查
	u := User{}
	// checkInterface(quit)
	checkInterface(&u)
	checkInterface(u)

	fmt.Println("13=======================================")

	fmt.Println(strings.TrimLeft("hello Tom", "hl")) // output: ello Tome
	words := "mongodb://oof"
	prefix := "mondb://"
	fmt.Println(strings.TrimLeft(words, prefix)) // output: godb://oof
	words = "mongodb://oooff"
	prefix = "mongodb://"
	fmt.Println(strings.TrimLeft(words, prefix)) // output: ff

	fmt.Println("100=======================================")

	// 创建一个net server.
	// netserver() 会永久阻塞
	go netserver()

	// 定时执行，ezTimer()会永久阻塞
	go ezTimer()

	fmt.Println("------------------------ End ...")

	// 等待信号来退出
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 0)
	// 将对应的信号通知sigs
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println("Recv signal: ", sig)
		done <- struct{}{}
	}()
	fmt.Println("I am waiting Signal")
	<-done
	fmt.Println("Exiting")
}
