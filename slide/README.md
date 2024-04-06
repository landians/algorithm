### 滑动窗口

滑动窗口，简单来讲就是维护一个窗口，然后根据需要进行
滑动来调整窗口的大小，一般用于解决数组类的问题. 
时间复杂度为 O(N). 它的算法的大致逻辑如下：
```
var arr []int

// 初始窗口大小为 0
lo, hi := 0, 0

for hi < len(arr) {
    // 增大窗口
    // 往窗口中增加元素
    window.add(s[hi])
    hi++
    
    for window needs shinl {
        // 缩小窗口
        // 往窗口中删除元素
        window.remove(s[lo])
        lo++
    }
}
```

滑动窗口算法框架
```
void slidingWindow(string s, string t) {
    unordered_map<char, int> need, window;
    for (char c : t) need[c]++;

    int left = 0, right = 0;
    int valid = 0; 
    while (right < s.size()) {
        // c 是将移入窗口的字符
        char c = s[right];
        // 右移窗口
        right++;
        // 进行窗口内数据的一系列更新
        ...

        /*** debug 输出的位置 ***/
        printf("window: [%d, %d)\n", left, right);
        /********************/

        // 判断左侧窗口是否要收缩
        while (window needs shrink) {
            // d 是将移出窗口的字符
            char d = s[left];
            // 左移窗口
            left++;
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```


