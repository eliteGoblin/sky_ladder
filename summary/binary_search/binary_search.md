

总结: 

*  自己consistent: end 标记不存在的位置: len(nums)
*  防止溢出写成: `mid = start + (end - start) / 2`: mid落在哪个index, 其实不重要。每次只要reduce就好
*  精确匹配简单: 加上: `if nums[mid] == target`
*  规范变量名: start, end

略去明显的精确匹配, 无论是

*  lower_bound: 第一个 >= val
*  upper_bound: 第一个 > val.

抽象为给定predicate: 将升序的数列分为两部分: 

false false false true 或者
true true true false

寻找突变的变量: 第一个true, 或者最后一个false.

可以转换为用end标定result: end为不存在的位置，存储符合条件的, "最右侧" 选项. 

最后return end (或者 return end + 1)?

用right适合找 first xx: right标定的: 

左边不满足predicate, 但是右边一直满足, 一直用right记录. 最后得到第一个满足的位置(在left发生越界之后).

如果是 左边一直满足，右边一直不满足: 反转逻辑: 让左边不满足，右边满足, 求出end. 最后根据end的相对位置求出result:

e.g 基于lower bound, 条件: last one < target

true true true false

完全翻转逻辑判断: 转换为first >= target: false false false true, 求出right; right - 1就是result.

总之自己的模板就是全部转为右边

## lower bound

```py
# first >= target
# < < < = = > > > > 
# 用right标定: right >= 
def lower_bound(nums: int, target: int) -> int:
    start, end = 0, len(nums)
    while start < end:
        mid = (start + end) // 2
        if nums[mid] >= target:
            end = mid
        else:
            start = mid + 1
    return end
```

## upper bound 
那么我们就要考虑使用类似于二分查找法来缩短时间，由于只是需要找到任意一个峰值，那么我们在确定二分查找折半后中间那个元素后，和紧跟的那个元素比较下大小，如果大于，则说明峰值在前面，如果小于则在后面。这样就
        start = mid + 1
return end
```


TODO: 
二分查找: 学习计划
https://leetcode-cn.com/leetbook/read/binary-search/xeog5j/