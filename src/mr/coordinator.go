package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// 主节点结构体
type Coordinator struct {
	// 在此处定义你的字段
	m_status []bool // map 任务状态数组
	r_status []bool // reduce 任务状态数组
}

// 在此处编写供 worker 调用的 RPC 处理函数

// 示例 RPC 处理函数
//
// RPC 的参数和返回类型在 rpc.go 中定义
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) Get_unstarted_task(args *Args, reply *Reply) error {

}

// 启动一个线程，监听来自 worker.go 的 RPC 请求
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go 会周期性调用 Done() 来判断
// 整个作业是否已完成
func (c *Coordinator) Done() bool {
	ret := false

	// 在此处添加你的代码

	return ret
}

// 创建一个 Coordinator
// main/mrcoordinator.go 会调用此函数
// nReduce 表示要使用的 reduce 任务数量
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}
	nMap := len(files)
	// 启动 nMap + nReduce 个 worker线程监听器
	for i := 0; i < nMap+nReduce; i++ {
		go c.server()
	}

	c.server()
	return &c
}
