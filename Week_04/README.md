# 学习笔记

## 深度优先搜索 & 广度优先搜索

- 遍历所有结点，按不同方式一般分为
    - 深度优先搜索，Depth First Search
    - 非递归DFS，需要借助Stack和visited结构体(一般为set or map保证访问时为O(1))
    - 广度优先搜索，Board First Search
    - BFS，一般借助队列进行层次遍历

## 贪心算法

- 在每一步中选取当前状态最好或最优的的选择，从而希望导致最终结果全局最优
- 适合场景
    - 问题能够分解为子问题，子问题的最优解，能够递推得到整个问题的最优解
- 与动态规划对比
    - 贪心：当下做局部最优，不能回退
    - 动态规划：根据以前的结果对当前进行选择，即最优判断+能够回退

## 二分查找

- 前提条件
    - 目标函数单调递增或递减
    - 存在上下界
    - 能够通过索引访问
- mid = l + (r - l)/2
    - mid = (l + r)/2可能导致数字越界
- 根据具体场景决定 mid+1 或者 mid-1

## 解题技巧

- 爬楼梯问题，实际为斐布拉切，避免子问题重复计算，推荐从下至上递推求解
- 判断二叉树是否有效，可采用判断当前元素是否在上下界之内
    - 初始上下界可定为 math.MinInt64 和 math.MaxInt64
- 二叉树是最适合采用递归的数据结构体之一
- 单词接龙问题，典型BFS使用场景
    - 使用map存储所有单词，便于后续O(1)访问
    - 使用对列存储当前层访问的单词
    - 遍历当前单词，并用‘a-z’进行替换
    - 判断&处理不同情况
    - return结果
- n皇后问题，递归+回溯+剪枝
- 二分查找典型题目，搜索二维矩阵

# 总结

使用递归(分治、回溯)、动态规划、BFS和DFS，以及二分查找是最重要的解题方法，需要通过刷更多的题来积累经验，后续再总结。
