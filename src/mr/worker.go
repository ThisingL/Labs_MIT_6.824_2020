package mr

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/rpc"
)

// Map 函数返回一个 KeyValue 切片。
type KeyValue struct {
	Key   string
	Value string
}

// 使用 ihash(key) % NReduce 为 Map 发出的每个 KeyValue 选择 reduce 任务编号。
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go 会调用此函数。
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// 在这里实现你的 worker。

	// 取消注释以向协调器发送示例 RPC。
	// CallExample()

}

// 向主节点请求一个任务
func Call_ask_for_task() {

}

// 向主节点报告完成任务
func Call_response_ok() {

}

// 示例函数，展示如何向协调器发起 RPC 调用。
//
// RPC 的参数和回复类型定义在 rpc.go 中。
func CallExample() {

	// 声明一个参数结构体。
	args := ExampleArgs{}

	// 填入参数。
	args.X = 99

	// 声明一个回复结构体。
	reply := ExampleReply{}

	// 发送 RPC 请求，等待回复。
	// "Coordinator.Example" 告诉接收服务器我们想调用 Coordinator 结构体的 Example() 方法。
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y 应该是 100。
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

// 向协调器发送 RPC 请求，等待响应。
// 通常返回 true。
// 出错时返回 false。
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
