log 系统设计

应该做什么？
    分级打印
    文件分割
    其他的暂时还没有想到

设计的很简单，lib.go 里面定义了 manager 和 printer 两个对象

manager 负责管理 打印协程 的开启和关闭

printer 负责对接输出设备

具体请见 example/example_test.go