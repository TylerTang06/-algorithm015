# 学习笔记

## 哈希表、映射、集合

### Hash Table

- 哈希表(Hash Table)/散列表
- 哈希函数(Hash Function)
- 哈希碰撞常用拉链法解决
- 缓存(LRU Cache)、健值对存储(Redis)等
- Average: Search O(1), Insertion O(1), Delete O(1)
- Worst: Search O(n), Insertion O(n), Delete O(1)

### Map: key-value, key不重合

- new HashMap()/new TreeMap()
- map.set(key, value)
- map.get(key)
- map.has(key)
- map.size()
- map.clear()

### Set: 不重合元素的集合

- new HashSet()/new TreeSet()
- set.add(value)
- set.delete(value)
- set.has(value)

## 树、二叉树、二叉搜索树

### 二叉搜索树

- 左子树上所有结点的值均小于它的根结点的值
- 右子树上所有结点的值均大于它的根结点的值
- 左右子树分别为二叉搜索树
- 中序遍历: 升序遍历
- 查询、插入、删除平均时间复杂度 O(logn)

## 堆、二叉堆、图

### 堆 Heap

- 可以迅速找到一堆数中的最大或者最小值的数据结构
- 大顶堆 与 小顶堆
- 大顶堆: find-max O(1); delete-max O(logn); insert O(logn) or O(1)

### 二叉堆

- 通过完全二叉树实现
- 是一棵完全树
- 树中任意结点的值总是>=其子结点的值
- 数组实现二叉树, 假设第一个结点在索引0位置:
    - 索引为i的左子树索引为 2*i+1
    - 索引为i的右子树索引为 2*i+2
    - 所以为i的结点的父结点索引为 floor((i-1)/2)


### 图

- 临接矩阵
- 临接表

## HashMap

- Golang中可基于Map实现
- 参考链接: https://github.com/emirpasic/gods

### Golang Map实现原理

- 哈希表作为底层数据结构
- 哈希表结点: bucket, 每个bucket含有一对或多对key-value
- 每个bucket可以存储8个key-value
- 哈希冲突: 链地址法
- 负载因子: 负载因子 = 键数量/bucket数量
- 扩容条件：
    - 负载因子 > 6.5时，也即平均每个bucket存储的键值对达到6.5个
    - overflow数量 > 2^15时，也即overflow数量超过32768时
- 扩容方式: 增量扩容、等量扩容

```
// map 数据结构
type hmap struct {
    count int // 当前保存的元素个数
    ...
    B unit
    ...
    buckets unsafe.Pointer // bucket数组指针, 数组的大小为2^B
}

// bucket 数据结构
type bucket struct {
    tophash [8]unit // 存储哈希值的高8位
    data byte[1] // key-value数据: key/key/key/.../value/value/value/...;如此存放是为了节省字节对齐带来的空间浪费
    overflow *bmap // 溢出bucket的位置, 指针指向的是下一个bucket, 据此将所有冲突的键连接起来
}
```

