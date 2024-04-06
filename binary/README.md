### 二分搜索法

二分搜索法的作用一般有以下三种:
1. 准确查找某个值
2. 查找左边界
3. 查找右边界

一般用的最多的还是查找左右边界, 即找最值。

二分搜索法的框架
```
// 函数 f 是关于自变量 x 的单调函数
// 例如: 随着 x 的增大，f(x) 的值越来越小, 就是单调递减
func (x int) int {
    // ...
}

// 主函数，在 f(x) == target 的约束下求 x 的最值
func solution(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
     // 确定左右边界
     lo, hi := ..., ...
     
     // 开始查找
     for lo < hi {
        mid := lo + (hi - lo) / 2
        if f(mid) == target {
            // 问自己：题目是求左边界还是右边界？
            // ...            
        } else if f(mid) < target {
            // 问自己：怎么让 f(x) 大一点？
            // ...
        } else if f(mid) > target {
            // 问自己：怎么让 f(x) 小一点？
            // ...            
        }
     }
     
     return lo
}
```