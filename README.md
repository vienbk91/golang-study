# Golang Study

## Lession01: Install Go


## Lession02: Hello word by Go

```
$ go version
go version go1.16.4 darwin/amd64
```

#### Check GOROOT 
```
$ go env GOROOT
/usr/local/go
```

#### Create $GOPATH
```
$ cd develop/study/02_golang
~/develop/study/02_golang                                              22:33:13
$ pwd
/Users/chuvien/develop/study/02_golang
$ export GOPATH=/Users/chuvien/develop/study/02_golang
$ go env GOPATH
/Users/chuvien/develop/study/02_golang
```

#### Create helloword folder
```
/golang-study
    /src
        /basic
            helloword.go
```

#### Create helloword content
```
package main

import "fmt"

func main() {
	fmt.Println("Helloword")
}
```

#### Run hellowrld
```
$ go run $GOPATH/golang-study/src/basic/helloword.go
Helloword
```

#### Build hellowrld
```
$ cd $GOPATH/golang-study/src/basic
$ go build helloword.go
$ ls -al
total 3976
drwxr-xr-x  5 chuvien  staff      160 May 17 01:17 .
drwxr-xr-x  3 chuvien  staff       96 May 17 01:06 ..
-rwxr-xr-x  1 chuvien  staff  2023840 May 17 01:17 helloword
-rw-r--r--  1 chuvien  staff      869 May 16 22:51 helloword.go
```

#### Run after build 
```
$ go ./helloword
Helloword
```
