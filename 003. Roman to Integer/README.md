___
## 003. Roman to Integer 
#### Easy, **`Go`**

___
Roman numerals are represented by seven different symbols: , , , , , and .`I` `V` `X` `L` `C` `D` `M`

**Symbol = Value**
* `I`   =   1
* `V`   =   5
* `X`   =   10
* `L`   =   50
* `C`   =   100
* `D`   =   500
* `M`   =   1000

For example,  is written as  in Roman numeral, just two ones added together. is written as , which is simply . The number is written as , which is .
`2`->`II`
`12`->`XII`=`X + II`
`27`->`XXVII`=`XX + V + II`

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not . Instead, the number four is
written as . Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as . 
There are six instances where subtraction is used:`IIII` `IV` `IX`

* `I` can be placed before (5) and (10) to make 4 and 9. `V` `X`
* `X` can be placed before (50) and (100) to make 40 and 90. `L` `C`
* `C` can be placed before (500) and (1000) to make 400 and 900. `D` `M`
Given a roman numeral, convert it to an integer.

 
___
**Example 1:**

* **Input:** `s` = "III"
* **Output:** 3
* **Explanation:** `III` = 3.

**Example 2:**

* **Input:** `s` = "LVIII"
* **Output:** 58
* **Explanation:** `L` = 50, `V` = 5, `III` = 3.

**Example 3:**

* **Input:** `s` = "MCMXCIV"
* **Output:** 1994
* **Explanation:** `M` = 1000, `CM` = 900, `XC` = 90 and `IV` = 4.
 

**Constraints:**

* 1 <= s.length <= 15
* s contains only the characters .('I', 'V', 'X', 'L', 'C', 'D', 'M')
* It is guaranteed that is a valid roman numeral in the range .s[1, 3999]
---
#### Go:
```Go
func romanToInt(s string) int {
    //...
}
```
---
