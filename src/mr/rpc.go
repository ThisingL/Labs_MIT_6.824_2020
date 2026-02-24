package mr

//
// RPC 定义。
//
// 记得所有名称首字母大写。
//

import (
	"os"
	"strconv"
)

//
// 示例：展示如何声明 RPC 的参数和返回值。
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

type Args struct {
	X int
}

type Reply struct {
	Y int
}

// 在此添加你的 RPC 定义。

// 在 /var/tmp 中为协调器生成一个相对唯一的 UNIX 域套接字名称。
// 不能使用当前目录，因为 Athena AFS 不支持 UNIX 域套接字。
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
