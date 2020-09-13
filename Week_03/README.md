# 学习笔记

## 递归模版

### golang

```golang
func recursion(level, n int, params ...interface{}) {
    // recursion terminator
    if level > n {
        // process result
        processResult()

        return 
    }

    // process current logic
    process(level, params)

    // drill down
    recursion(level+1, n, params)

    // reverse current status if needed

    return
}
```

### 思维要点

- 不要人肉进行递归
- 找到最近最简单方法，将其拆分成可重复解决的问题
- 数学归纳法思想

## 分治、回溯

### 分治

- problem->subproblems
- split->merge

```golang
func divideConquer(problem interface{}, params ...interface{}){
    // recursion terminator
    if problem == nil {
        processResult()

        return 
    }

    // prepare data
    data := prepareData(problem)
    subproblems := splitProblem(problem, data)

    // conquer subproblems
    subproblem1 := divideConquer(subproblems[0], params)
    subproblem2 := divideConquer(subproblems[1], params)
    ...

    // process and generate final reslut
    result := processResult(subproblem1, subproblem2, ...)

    return
}
```

### 回溯

- 采用试错的思想(尝试分步去解决一个问题)
- 分步解决问题过程中
    - 找到一个可能的答案
    - 找不到答案，取消上一步或几步，通过其他可能分步解答尝试找到答案

## 总结

对递归，分治，回溯的概念还需要加强理解，特别是分治与回溯。