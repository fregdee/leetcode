# 13. Roman to Integer

https://leetcode.com/problems/roman-to-integer/

# Idea
問題は、事前に与えられた roman number と integer の対応表が
```
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```
となっており、与えられた文字からその integer としての数値を求める。
ただ、一つ条件があって、基本的に文字列は↑の表で大きいものから並んでいるが、そうなっていない場合に次の文字から値を引く必要がある。

## 問題のアプローチ方法
1. 文字列を左から1つずつ見ていき、次の文字の値のほうが大きい場合には引く
2. 引くパターンの文字列も一つの文字と見なす ( `CM` )
    ```
    Symbol       Value
    I             1
    IV            4
    V             5
    IX            9
    X             10
    XL            40
    L             50
    XC            90
    C             100
    CD            400
    D             500
    CM            900
    M             1000
    ```

# Complexity Analysis
- Time complexity
- Space complexity
