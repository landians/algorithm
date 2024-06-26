### 回溯算法

解决一个回溯问题，本质上就是解决一个决策树的遍历过程. 
只需要思考三个问题:
1. 路径，即已经做出的选择
2. 选择列表，即当前可以做的选择
3. 结束条件，即到达决策树底层后，无法再做选择的条件

回溯算法的框架
```
var result []int

func backtrack(路径, 选择列表) {
    if 满足结束条件:
       result = append(result, 路径)
       return
    
    for 选择 range 选择列表 {
        做选择
        backtrack(路径, 选择列表)
        撤销选择
    }  
}
```
