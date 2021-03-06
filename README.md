# Learn Go

1. ทำความรู้จักภาษา Go ผ่าน **A Tour of Go** [ภาษาไทย](https://go-tour-th.appspot.com/list) หรือ [English](https://tour.golang.org/list)
2. จากนั้นลองเขียน Go ผ่านบทเรียน [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/)

![](https://raw.githubusercontent.com/junlapong/learn-go-with-tests/master/red-green-blue-gophers-smaller.png)

## Check List

- [x] Install Go
- [x] Hello, World
- [x] Integers
- [x] Iteration
- [x] Arrays and slices
- [ ] Structs, methods & interfaces
- [ ] Pointers & errors
- [ ] Maps
- [ ] Dependency Injection
- [ ] Mocking
- [ ] Concurrency
- [ ] Select
- [ ] Reflection
- [ ] Sync
- [ ] Context

## Notes

- [Go cheatsheet](https://devhints.io/go)

### somkiat
- http://www.somkiat.cc/learn-golang-01
- http://www.somkiat.cc/learn-golang-02
- http://www.somkiat.cc/learn-golang-03
- http://www.somkiat.cc/learn-golang-04
- http://www.somkiat.cc/learn-golang-05
- http://www.somkiat.cc/learn-golang-06
- [สรุปเรื่องของ Principle 3 ข้อของภาษา Go](http://www.somkiat.cc/go-principles/)
- [ว่าด้วยเรื่องการจัดการ Error ถ้ามันเยอะก็ลดสิ](http://www.somkiat.cc/go-error-handling/)
- [* สรุปเรื่องของการพัฒนา Testable application ด้วยภาษา Go](http://www.somkiat.cc/testable-app-with-golang/)
- more... [go](http://www.somkiat.cc/tag/go/) | [golang](http://www.somkiat.cc/tag/golang/)


### IPorsut
 - [Golang array, slice](https://iporsut.blogspot.com/2014/09/golang-array-slice.html)
 - [Golang string, bytes, runes](https://iporsut.blogspot.com/2014/09/golang-string-bytes-runes.html)
 - [more...](https://iporsut.blogspot.com/search/label/Golang)


### superhardcode
- [ภาษา Golang มีอะไรน่าสนใจ](https://superhardcode.wordpress.com/2015/07/14/%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2-golang-%E0%B8%A1%E0%B8%B5%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3%E0%B8%99%E0%B9%88%E0%B8%B2%E0%B8%AA%E0%B8%99%E0%B9%83%E0%B8%88/)
- [มาเล่น ภาษา Go กันเถอะ!](https://superhardcode.wordpress.com/2015/07/14/%E0%B8%A1%E0%B8%B2%E0%B9%80%E0%B8%A5%E0%B9%88%E0%B8%99-%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2-go-%E0%B8%81%E0%B8%B1%E0%B8%99%E0%B9%80%E0%B8%96%E0%B8%AD%E0%B8%B0/)
- [มาดู Function ใน GO กัน มีไรดี](https://superhardcode.wordpress.com/2015/07/17/%E0%B8%A1%E0%B8%B2%E0%B8%94%E0%B8%B9-function-%E0%B9%83%E0%B8%99-go-%E0%B8%81%E0%B8%B1%E0%B8%99-%E0%B8%A1%E0%B8%B5%E0%B9%84%E0%B8%A3%E0%B8%94%E0%B8%B5/)
- [Data ใน Go มีอะไรบ้างนะ](https://superhardcode.wordpress.com/2015/07/20/data-%E0%B9%83%E0%B8%99-go-%E0%B8%A1%E0%B8%B5%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%87%E0%B8%99%E0%B8%B0/)
- [ว่าด้วยเรื่องของ Struct ใน GO](https://superhardcode.wordpress.com/2015/07/28/%E0%B8%AA%E0%B9%88%E0%B8%AD%E0%B8%87-struct-%E0%B9%83%E0%B8%99-go-%E0%B8%81%E0%B8%B1%E0%B8%99/)


### others
- [สรุปความรู้จากการฟัง Live : Golang Workshop [Part I]](https://medium.com/thipwriteblog/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B8%84%E0%B8%A7%E0%B8%B2%E0%B8%A1%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B8%88%E0%B8%B2%E0%B8%81%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%9F%E0%B8%B1%E0%B8%87-live-golang-workshop-part-i-f00504d4366c)


## Tools

- [Gin - Live reload utility for Go web servers](https://github.com/codegangsta/gin)
```
go get github.com/codegangsta/gin
gin run main.go
```
- [Ginkgo is a BDD-style Go testing framework ](https://onsi.github.io/ginkgo/)
- [GoConvey, a yummy testing tool for gophers](http://goconvey.co/) | [Introduction to GoConvey](https://youtu.be/wlUKRxWEELU) on YouTube
```
go get github.com/smartystreets/goconvey
$GOPATH/bin/goconvey
```

## Test Coverage

```
go test -cover -coverprofile=cover.out
go tool cover -html=cover.out -o coverage.html
open coverage.html
```
