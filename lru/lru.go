package main

import (
	"reflect"
)

// todo 并发安全 	atomic.XXX

type node struct {
	key  string
	val  interface{}
	next INode
	prev INode
}

func NewNode(key string, val interface{}) INode {
	return &node{key: key, val: val}
}

func (n *node) Next() INode {
	return n.next
}

func (n *node) SetNext(node INode) {
	n.next = node
}

func (n *node) InsertNext(node INode) {
	node.SetNext(n.next)
	n.next.SetPrev(node)
	n.next = node
	node.SetPrev(n)
}

func (n *node) HasNext() bool {
	return n.next != nil
}

func (n *node) Prev() INode {
	return n.prev
}

func (n *node) HasPrev() bool {
	return n.prev != nil
}

func (n *node) SetPrev(node INode) {
	n.prev = node
}

func (n *node) InsertPrev(node INode) {
	n.prev.SetNext(node)
	node.SetPrev(n.prev)
	node.SetNext(n)
	n.prev = node
}

func (n *node) Key() string {
	return n.key
}

func (n *node) SetKey(s string) {
	n.key = s
}

func (n *node) Value() interface{} {
	return n.val
}

func (n *node) SetValue(i interface{}) {
	n.val = i
}

type cache map[string]INode

type Lru struct {
	start      INode
	end        INode
	maxLen     uint64
	currentLen uint64
	cache      cache
}

func (l *Lru) Len() uint64 {
	return l.currentLen
}

func (l *Lru) MaxLen() uint64 {
	return l.maxLen
}
func (l *Lru) Cache() map[string]INode {
	return l.cache
}

func (l *Lru) Val(key string) (interface{}, bool) {
	v, ok := l.cache[key]
	if !ok {
		return nil, false
	}
	l.setStart(v)

	return v.Value(), true
}


func (l *Lru) setStart(node INode) {

	if reflect.DeepEqual(l.start, node){
		return
	}
	if reflect.DeepEqual(l.end,node){
		l.Del()
		l.Put(node)
	}else {
		node.Prev().SetNext(node.Next())
		node.Next().SetPrev(node.Prev())
		l.put(node)
	}
}




func NewLru(maxLen uint64) ILruCache {
	if maxLen < 1 {
		maxLen = 1
	}
	return &Lru{maxLen: maxLen, cache: cache{}}
}

func (l *Lru) Start() (INode, bool) {
	return l.start, l.start != nil
}

func (l *Lru) End() (INode, bool) {
	return l.end, l.end != nil
}

func (l *Lru) put(n INode) bool{
	l.cache[n.Key()] = n
	if l.start == nil || l.maxLen == 1 {
		n.SetNext(n)
		n.SetPrev(n)
		l.start = n
		l.end = n
		l.currentLen = 1
		return false
	}

	old := l.start

	if l.IsOverflow() {
		// 长度已满
		l.Del()
		l.end.SetNext(nil)
	} else if l.currentLen == 1 {
		old.SetNext(nil)
	}

	old.SetPrev(n)
	n.SetNext(old)
	n.SetPrev(nil)
	l.start = n
	return true
}


func (l *Lru) Put(n INode) {
	if l.put(n){
		l.currentLen += 1
	}
}


func (l *Lru) Del() {
	if l.end != nil {
		delete(l.cache, l.end.Key())
		if l.maxLen == 1 || l.currentLen == 1 {
			delete(l.cache, l.start.Key())
			l.start = nil
			l.end = nil
		} else {

			end := l.end.Prev()
			end.SetNext(nil)
			l.end = end
			if reflect.DeepEqual(l.start, l.end) {
				l.start.SetPrev(l.start)
				l.start.SetNext(l.start)
			}
		}
		l.currentLen -= 1
	}
}

func (l *Lru) IsOverflow() bool {
	return l.currentLen == l.maxLen
}

type INode interface {
	Next() INode
	SetNext(node INode)
	InsertNext(node INode)
	HasNext() bool
	Prev() INode
	HasPrev() bool
	SetPrev(node INode)
	InsertPrev(node INode)
	Key() string
	SetKey(string)
	Value() interface{}
	SetValue(interface{})
}

type ILruCache interface {
	// 获取头部节点
	Start() (INode, bool)
	// 获取尾部节点
	End() (INode, bool)
	// 新增节点
	Put(n INode)
	// 删除尾部节点
	Del()
	// 长度
	Len() uint64
	MaxLen() uint64
	// 是否可以新增
	IsOverflow() bool
	// 查询cache
	Val(key string) (interface{}, bool)
	// 所有缓存数据
	Cache() map[string]INode
}
