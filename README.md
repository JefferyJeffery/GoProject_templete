# Go Project Template

## 說明

一般Go項目在GOPATH下，會有如下三個目錄
~~~
|--bin
|--pkg
|--src
~~~

## Setting

### mkdir
~~~
mkdir Go_Project
cd Go_Project
~~~

### 將目前位置設定為專案位置

~~~
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
~~~

### Install govendor to manage vendor

~~~
go get -u github.com/kardianos/govendor

> [~/Documents/WorkSpace/MyStudy/Go/test_Govender]
> $ ll
> total 0
> drwxr-xr-x  5 mac  staff  170 Jul 18 19:01 ./
> drwxr-xr-x  6 mac  staff  204 Jul 18 12:42 ../
> drwxr-xr-x  3 mac  staff  102 Jul 18 19:01 bin/
> drwxr-xr-x  3 mac  staff  102 Jul 18 19:01 pkg/
> drwxr-xr-x  3 mac  staff  102 Jul 18 19:01 src/
~~~

### 接下來，增加一個包 config , 和一個main程序。
~~~
cd src
mkdir config
touch config/config.go

mkdir backend 
touch backend/main.go
~~~

~~~
$ cat config/config.go
package config

import (
	"fmt"
)

func Run() {
	fmt.Println(
		"i'm config",
	)
}
~~~

~~~
$ cat backend/main.go
package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println(
		"i,m backend/main",
	)

	config.Run()
}
~~~

### 目錄結構如下：

~~~
test_Govender
|-- bin
|-- pkg
`-- src
    |-- config
    |   `-- config.go
    `-- backend
        `-- main.go
~~~

### 編譯並安裝指定的代碼包及它們的依賴包 
go install backend


#### 目錄結構如下：

~~~
test_Govender
|-- bin
    `-- backend
|-- pkg
`-- src
    |-- config
    |   `-- config.go
    `-- backend
        `-- main.go
~~~

可以直接執行

~~~
$ backend
i,m backend/main
i'm config
~~~

## 參考
http://blog.studygolang.com/2012/12/go%E9%A1%B9%E7%9B%AE%E7%9A%84%E7%9B%AE%E5%BD%95%E7%BB%93%E6%9E%84/

https://github.com/kilfu0701/go-vendor-example

http://www.evanlin.com/about-go-vendoring/

http://wiki.jikexueyuan.com/project/go-command-tutorial/0.1.html
