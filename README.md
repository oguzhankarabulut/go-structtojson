## GO Struct to JSON

I hate creating example JSON object fit to struct when testing something especially working with huge structs. So I developed this tool. It takes struct as a string cli argument then print example JSON object with random values.

PS: It does not work with nested data types.

## Usage ##
Input:
```ssh
$ go-structtojson "type test struct {
        s string
        i int
        i2 int32
        i3 int64
        f float32
        f2 float64
        l []bool
        m map[string]int
}"
```
Output:
```ssh
$ {"f":0.19,"f2":0.99,"i":164,"i2":1104271735,"i3":5854854829731596507,"l":[false,false],"m":{"mXSGYgENYK":96},"s":"hkhRnwAkRh"}
```

## Installation ##

```ssh
$ go install github.com/oguzhankarabulut/go-structtojson@latest
```
Then copy from `$GOPATH/bin/go-structtogo` to `$GOROOT/bin`

You are ready to use!
