package num

import (
	"sync"
	"unsafe"
)

// 简单的链表
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 数组
type Array struct {
	Data      string
	NextIndex int64
}

// 切片
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// 可变长数组
type Arrays struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 数组栈，先进后出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 链表栈，先进后出
type LinkStack struct {
	root *LinkNodes // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNodes struct {
	Next  *LinkNodes
	Value string
}

// 数组队列，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 链表队列，先进先出
type LinkQueue struct {
	root *LinkNoded // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNoded struct {
	Next  *LinkNoded
	Value string
}

// 双端列表，双端队列
type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进行并发安全pop操作
}

// 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}
