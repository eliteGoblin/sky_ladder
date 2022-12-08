
# CSDN

400题列表和250精选

*  [400题分类顺序](https://cspiration.com/leetcodeClassification)

## CLI

### 登录

*  Grab cookie: 
    1.  Chrome make request: `https://leetcode.com/problemset/all/`
    2.  XHR request -> 找到URL: `https://leetcode.com/api/problems/all/`
    3.  Requst Headers: 找到: cookie: xxx, copy value

Note:

*  Cookie似乎包含了过时的session信息: 如果AC了新题，leetcode-cli看不到，仍显示为未AC. 解决办法: 先logout: `leetcode user -L`, 然后重新grab cookie, 登录`leetcode user -c`

Ref: 

[Failed to log in with a leetcode.com account](https://github.com/LeetCode-OpenSource/vscode-leetcode/issues/478)  

### 使用

```s
# e: easy D: not done, L: non lock p: plan category; all/array means follow csdn plan; if empty, means not follow csdn: all problem in leetcode
leetcode ls -q eDL -p array
```

#### 背景
*  leetcode-cli处于`~/local/lib/node_modules`, 被全局安装的node package., leetcode位于: `/home/frank.sun/local/bin/leetcode` executable指向:
```s
#!/usr/bin/env node
require('../lib/cli').run();
```
用的是soft link, 最终指向的是: `git_repo`的版本