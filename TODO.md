# 并查集

用于解决集合分组问题

# 动态规划常见类型

- 线性动态规划 dp[i] = f(dp[i-1], dp[i-1], ... dp[0])
- 前缀和 dp[i] = f(sum[i], sum[i-1], ... sum[0])
- 区间动态规划 dp[i][j] = f(dp[i][j-1], dp[i][j-1], dp[i-1][j-1], ...)
- 背包问题 dp[i][v] = f(dp[i-1][v - k*V[i]] + k*W[i])
- 组合计数

## 最长递增序列

输入 nums[i]
输出 dp[i] 以 nums[i] 为结尾的最长递增序列
```if j < i && nums[j] < nums[i] {dp[i] = max(dp[j])+1} else {dp[i] = 1}```

```index = search(dp, nums[i]) if len(dp) == index {dp = append(dp, nums[i])} else {dp[i] = nums[i]}```

## 连续子数组和最大值

输入 nums[i]
输出 max(sum) ```sum += nums[i] if sum <= 0 {sum = 0}```

## 最长公共子序列

输入 nums1[i] nums2[j]
输出 dp[i][j] nums[i] 和 nums[j] 的最长公共子序列
```dp[i][j] = if nums[i] == nums[j] {dp[i-1][j-1]+1} else {max(dp[i-1][j], dp[i][j-1])} ```

## 背包问题

- 0-1背包 ```dp[i][v] = max(dp[i-1][v], dp[i-1][v-V[i]] + W[i])```
- 多重背包/完全背包 ```dp[i][v] = max(dp[i-1][v - k*V[i]] + k*W[i])```
- 多价值背包 ```dp[i][v1][v2] = max(dp[i-1][v1][v2], dp[i-1][v1-V1[i]][v2-V2[i]] + W[i])```

## 零钱问题

可以将问题转化为背包问题

## 将集合划分为两个组，计算两组和的最小差

将问题转化为给每个数字添加 +/- 号，求和最小。
也可以将问题转化为求出和为所有元素和一半的最大背包问题 A = max(sum/2)，S = sum - 2*A

# 拓扑排序

初始状态下，集合 S 装着所有入度为 0 的点，L 是一个空列表。

每次从 S 中取出一个点 u（可以随便取）放入 L, 然后将 u 的所有边 (u, v_1), (u, v_2), (u, v_3) ... 删除。对于边 (u, v)
，若将该边删除后点 v 的入度变为 0，则将 v 放入 S 中。

不断重复以上过程，直到集合 S 为空。检查图中是否存在任何边，如果有，那么这个图一定有环路，否则返回 L，L 中顶点的顺序就是拓扑排序的结果。

# KMP

| index   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
|---------|---|---|---|---|---|---|---|---|---|
| kmp     | 0 | 0 | 0 | 1 | 2 | 3 | 0 | 1 | 2 |
| pattern | a | b | c | a | b | c | c | a | b |

```go
package main

type InfiniteStream interface {
  Next() int
}

// 在无限流中寻找匹配 https://leetcode.cn/problems/find-pattern-in-infinite-stream-ii/description/
func findPattern(stream InfiniteStream, pattern []int) int {
  kmp := make([]int, len(pattern))
  for i := 1; i < len(kmp); i++ {
    if pattern[i] == pattern[kmp[i-1]] {
      kmp[i] = kmp[i] + 1
    }
  }

  for i, s := 0, make([]int, 0, 2*1e5); ; i++ {
    if len(s) < len(pattern) {
      s = append(s, stream.Next())
    }

    for j := 0; j < len(s); j++ {
      if s[j] != pattern[j] {
        s = s[kmp[j]+1:]
        break
      }
    }

    if len(s) == len(pattern) {
      return i - len(pattern) + 1
    }
  }
}

```

# 单调栈

# 单调队列

# 字典树

查找字符串和字符串前缀

# 线段树/前缀和/树状数组

# 双指针

# 快速选择

# 图论

## 最小生成树

- Kruskal 算法
  按照边的权重顺序（从小到大）将边加入生成树中，
  但是若加入该边会与生成树形成环则不加入该边。
  直到树中含有 V-1 条边为止。
  这些边组成的就是该图的最小生成树。