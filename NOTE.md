
## Slice

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