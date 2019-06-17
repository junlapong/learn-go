
# Notes

## Array

Array is a compound object with fix length of object of the same type.

```go
package main

import "fmt"

func main() {
    //This defines a variable that holds 4 integers
    var a [4]int = [4]int{1,2,3,3}
    //go array index value can be mutated unlike rust
    a[2] = 10
    b := [5]int{1,2,3,4,5}
    fmt.Println(a)
    fmt.Println(b)
    //If you try to assign array of length 5 to a, it errors
}
```

## Slice

A slice is a data type that is backed by an array but with mutable size. It consists of a type and size basically pointing to a specific part of array, length of the array that the segment has and maximum capability of the underlying array starting from the the slice head.

```go
package main

import "fmt"

func main() {
    arr := [4]int{1,2,3,4}
    s := arr[1:] //The defaults to arr[1:4] with last one excluded
    //You can also make slice using `make`
    m := make([]int,1,2) //The creates a slice with length 1 but capacity 2.
    fmt.Println(m)
    //Ok, lets just copy one slice to another. This only copies one element because of m length
    copy(m, arr[1:])
    fmt.Println(m)
    fmt.Println(s)
}
```

You can define a slice literal or go the route of using make function specifying the length as well as capacity as we have seen above.

### Null Slice

When you define a slice without setting a value, what you have is a null slice. There is no underlying array that is backing the slice. One thing that stands out with null slice is they can act like zero length slice which means you can treat them like a slice that you defined with zero length.

```go
package main

import "fmt"

func main() {
   //This slice defined below will be treated like an empty slice of type int
   var s []int
   fmt.Println(s)
   //Lets add some int to this null slice; Null because it does not have underlying array
   s = append(s,1,2,3,4)
   fmt.Println(s)
}
```

```
[]
[1 2 3 4]
```
You have to be careful with slice as they can hold large objects in memory hostage even if they refer to only a small part of the data.

### Slice Operations

```go
a := [5]int{1, 2, 3, 4, 5}
b := a[0:3]

fmt.Println(a)
fmt.Println(b)
// b จะกลายเป็นข้อมูลแบบ slice เพราะเกิดจากการ slice array a
// จากตำแหน่งที่ 0 จนถึงตำแหน่งก่อนหน้า 3 ดังนั้นข้อมูลของ b คือ [1,2,3]

fmt.Println(a[1:5])
fmt.Println(a[2:4])
```

```
[1 2 3 4 5]
[1 2 3]
[2 3 4 5]
[3 4]
```

```go
var a []int
a = append(a, 1)
fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))

a = []int{1, 2, 3, 4}
fmt.Printf("a: %v\n", a)

// การ copy slice โดยไม่เปลี่ยนค่าของ slice เดิม
c := make([]int, len(a))
copy(c, a)

b := a
// การกำหนดค่าตัวแปร slice ให้กับตัวแปรอื่น จะเป็นการสร้าง slice ใหม่
// แต่ข้อมูล array ภายในของทั้ง 2 slice จะมีค่าเป็นตัวเดียวกัน

b[0] = 5
// เมื่อเปลี่ยนค่า b[0] ขอมูลใน a จะกลายเป็น [5,2,3,4] ไปด้วย

fmt.Printf("a: %v\n", a)
fmt.Printf("b: %v\n", b)
fmt.Printf("c: %v\n", c)
```

```
a: [1], len: 1, cap: 1
a: [1 2 3 4]
a: [5 2 3 4]
b: [5 2 3 4]
c: [1 2 3 4]
```