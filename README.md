# Gverify

一款用于Golang的结构体校验工具。它通过手动编排校验过程进行校验，而非使用标签。使用标签的验证工具，无法用在经常被代码生成器覆盖的结构体，例如grpc、gorm生成的代码。gverify无需修改结构体，并且简单易用，处理速度理论上比使用标签的校验工具更快。

[![Go Reference](https://pkg.go.dev/badge/github.com/jishaocong0910/gverify.svg)](https://pkg.go.dev/github.com/jishaocong0910/gverify)
[![Go Report Card](https://goreportcard.com/badge/github.com/jishaocong0910/gverify)](https://goreportcard.com/report/github.com/jishaocong0910/gverify)
![coverage](https://raw.githubusercontent.com/jishaocong0910/gverify/badges/.badges/main/coverage.svg)

# 安装

```shell
go get github.com/jishaocong0910/gverify
```

# 用法与例子

*代码示例*

```go
package main

import (
    "context"
    "fmt"
    "regexp"

    vfy "github.com/jishaocong0910/gverify"
)

type Book struct {
    Title       string
    Isbn        string
    Description *string
    Stock       int
    Price       float64
    Language    *string
    Author      *Author
    Categories  []*Category
}

func (b *Book) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &b.Title, "title").
        NotBlank().Msg("%s must not be blank", ctx.FieldName()).
        Max(10).Msg("%s's length exceed %s", ctx.FieldName(), ctx.Confine(0))

    vfy.String(ctx, &b.Isbn, "isbn").
        NotBlank().DefaultMsg().
        Regex(regexp.MustCompile(`^[0-9]{13}$`)).DefaultMsg()

    vfy.String(ctx, b.Description, "description").
        NotNil().DefaultMsg()

    vfy.Int(ctx, &b.Stock, "stock").
        Range(0, 100).Msg("%s must be %s to %s", ctx.FieldName(), ctx.Confine(0), ctx.Confine(1))

    vfy.Float64(ctx, &b.Price, "price").
        Gt(0).DefaultMsg()

    vfy.String(ctx, b.Language, "language").
        Options([]string{"zh-cn", "en-US", "ja-JP"}).Msg("%s must be %s", ctx.FieldName(), ctx.Confines())

    vfy.Struct(ctx, b.Author, "author").
        NotNil().DefaultMsg().
        Dive()

    vfy.Slices(ctx, b.Categories, "categories").
        NotEmpty().Msg("%s must not be empty", ctx.FieldName()).
            Dive(func(e *Category) {
                vfy.Struct(ctx, e, "").
                    NotNil().Msg("%s must not be nil", ctx.FieldName()).
                    Dive()
            })
}

type Author struct {
    Name    string
    Resumes *string
}

func (a Author) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &a.Name, "name").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    vfy.String(ctx, &a.Name, "resumes").NotBlank_(true).Msg("%s must not be blank", ctx.FieldName())
}

type Category struct {
    Id   string
    Sort int
}

func (c Category) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &c.Id, "id").NotBlank_(true).Msg("%s must not be blank", ctx.FieldName())
    vfy.Int(ctx, &c.Sort, "sort").Max(999).Msg("%s exceeded limit %s", ctx.FieldName(), ctx.Confines())
}

func main() {
    b := Book{Author: &Author{}, Categories: []*Category{{Sort: 127}, {Id: "c1", Sort: 1000}}}
    code, _, msgs := vfy.Check_(context.Background(), &b, true)
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output：
    // 0 title must not be blank
    // 1 isbn must not be blank
    // 2 isbn's format is illegal
    // 3 description must not be nil
    // 4 price must be greater than 0
    // 5 language must be zh-cn, en-US or ja-JP
    // 6 author.name must not be blank
    // 7 author.resumes must not be blank
    // 8 categories[0].id must not be blank
    // 9 categories[1].sort exceeded limit 999
}
```

# 验证函数

| 函数         | 说明                                        | 参数                                                                                                                      | 返回                                                                                                                           |
|------------|-------------------------------------------|-------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|
| vfy.Check_ | 检查结构体，可指定检查至首个错误的字段，或所有字段。                | `ctx`：`context.Context`，可为nil。<br/><br/>`v`：待验证的结构体，须实现`vfy.Verifiable`接口。<br/><br/>`all`：false则检查至首个错误字段 ，true则检查所有字段。 | `code`：错误码，字符串`SUCCESS`表示验证成功，`ERROR`表示有验证错误，可使用常量`vfy.SUCCESS`和`vfy.ERROR`进行比较。<br/><br/>`first`：首个错误消息。<br/>`msgs`：所有错误消息。 |
| vfy.Check  | 检查至首个错误的字段，相当于`vfy.Check_(ctx, v, false)` | 比`vfy.Check_`少了`all`参数                                                                                                  | 与`vfy.Check_`函数相同。                                                                                                           |

# 编写结构体校验过程

`vfy.Verifiable`接口的`Checklist`方法，用于编写每个字段的校验过程。方法的接收者为指针时，**无需担心它为nil**
，*验证函数*会判断传入的值是否为nil，若为nil则会创建零值进行验证。

# 字段验证函数

字段校验时，须根据自身类型，调用对应的*字段验证函数*，再链式调用*校验方法*。所有*字段验证函数*的第一个参数都传入`Checklist`方法的`ctx`参数。

| 字段验证函数      | 对应类型                   |
|-------------|------------------------|
| vfy.Bool    | bool                   |
| vfy.Byte    | byte                   |
| vfy.Int     | int                    |
| vfy.Int8    | int8                   |
| vfy.Int16   | int16                  |
| vfy.Int32   | int32                  |
| vfy.Int64   | int64                  |
| vfy.Float32 | float32                |
| vfy.Float64 | float64                |
| vfy.Uint    | uint                   |
| vfy.Uint8   | uint8                  |
| vfy.Uint16  | uint16                 |
| vfy.Uint32  | uint32                 |
| vfy.Uint64  | uint64                 |
| vfy.String  | string                 |
| vfy.Slices  | 切片                     |
| vfy.Map     | map                    |
| vfy.Struct  | 结构体（需实现vfy.Verifiable） |
| vfy.Any     | any                    |

# 校验方法

有些*校验方法*之间的名称，仅区别于是否具有下划线，例如`NotBlank`和`NotBlank_`
。带下划线的方法多一个参数用于指定值为nil时是否跳过，不带下划线的方法默认不跳过，例如`NotBlank()`相当于`NotBlank_(false)`
。下面列举所有*校验方法*，为节省篇幅不列举带下划线的方法。

<table>
    <tr>
        <th>校验方法</th><th>类型</th><th>说明</th>
    </tr>
    <tr>
        <td>NotNil</td>
        <td>所有</td>
        <td>不能为nil</td>
    </tr>
    <tr>
        <td>NotBlank</td>
        <td rowspan="2">string</td>
        <td>字符串不能为空白</td>
    </tr>
    <tr>
        <td>Regex</td>
        <td>字符串格式必须符合正则表达式</td>
    </tr>
    <tr>
        <td>NotEmpty</td>
        <td>切片 map</td>
        <td>元素数量必须大于0</td>
    </tr>
    <tr>
        <td>Length</td>
        <td>string 切片 map</td>
        <td>指定长度（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Min</td>
        <td rowspan="6">byte int int8 int16 int32 int64 float32 float64 uint uint8 uint16 uint32 uint64 string 切片 map</td>
        <td>必须大于等于指定值（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Max</td><td>必须小于等于指定值（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Range</td><td>必须在指定范围，包含边界（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Gt</td><td>必须大于指定值（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Lt</td><td>必须小于指定值（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Within</td><td>必须在指定范围内，即不包含边界（string类型校验UTF8字符长度，切片和map类型校验元素长度）</td>
    </tr>
    <tr>
        <td>Options</td>
        <td >byte int int8 int16 int32 int64 float32 float64 uint uint8 uint16 uint32 uint64 string</td>
        <td>必须在指定选项中选择</td>
    </tr>
    <tr>
        <td>Custom</td>
        <td>bool byte int int8 int16 int32 int64 float32 float64 uint uint8 uint16 uint32 uint64 string any</td>
        <td>自定义校验</td>
    </tr>
    <tr>
        <td>Dive</td>
        <td>struct 切片 map</td>
        <td>校验struct字段、切片元素、map元素的key和value</td>
    </tr>
</table>

# 错误消息&错误码

## 自定义消息

*校验方法*调用后可链式调用`Msg`方法来指定错误消息。错误消息使用Golang自带的格式化字符串，`Checklist`方法的`ctx`
参数提供了一些方法，可用于填充错误消息。

| ctx的方法    | 说明                                                                                                                                                                   |
|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FieldName | 返回具有路径的字段名称。例如：`title`、`author.name`、`categoryId[2].sort`。                                                                                                           |
| Confine   | 返回*校验方法*的指定索引的限制值的字符串形式。例如：对于`Max(10)`，`ctx.Confine(0)`返回`10`；对于`Range(5, 15)`，`ctx.Confine(0)`返回`5`，`ctx.Confine(1)`返回`15`。                                         |
| Confines  | 返回*校验方法*的所有限制值的字符串形式，用`,`拼接，若数量超过两个，则最后一个用`or`拼接。例如：对于`Options("zh-cn", "en-US")`，返回`zh-cn, en-US`；对于`Options("zh-cn", "en-US", "ja-JP")`，返回`zh-cn, en-US or ja-JP`。 |
| Index     | 返回切片元素索引，必须在切片的`Dive`*校验方法*中使用才有效，否则返回-1。                                                                                                                            |

## 元素字段名称

在切片和map的`Dive`*校验方法*内校验元素值时，会预设*元素字段名称*，此时若*字段验证函数*的`fieldName`参数为空字符串时则自动使用。对于切片设置为`[<索引>]`；对于map，key设置为`$key`，value的设置为`$value`。

*代码示例*

```go
package main

import (
    "context"
    "fmt"
    "strconv"

    vfy "github.com/jishaocong0910/gverify"
)

type Demo struct {
    MySlice  []string
    MyMap    map[string]int
    MySlice2 []string
}

func (d Demo) Checklist(ctx *vfy.Context) {
    vfy.Slices(ctx, d.MySlice, "mySlice").Dive(func(e string) {
        // fieldName为空字符串时默认为"[<索引>]"
        vfy.String(ctx, &e, "").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    })
    vfy.Slices(ctx, d.MySlice2, "mySlice2").Dive(func(e string) {
        // fieldName不为空字符串，则使用该值
        vfy.String(ctx, &e, "#"+strconv.Itoa(ctx.Index()+1)).NotBlank().Msg("%s must not be blank", ctx.FieldName())
    })
    vfy.Map(ctx, d.MyMap, "myMap").Dive(func(k string) {
        // fieldName为空字符串时默认为"$key"
        vfy.String(ctx, &k, "").NotBlank().Msg("%s must not be blank", ctx.FieldName()).
            Gt(2).Msg("%s's length must greater than %s", ctx.FieldName(), ctx.Confine(0))
    }, func(v int) {
        // fieldName不为空字符串，则使用该值
        vfy.Int(ctx, &v, "@value").Within(0, 100).Msg("%s must be > %s and < %s", ctx.FieldName(), ctx.Confine(0), ctx.Confine(1))
    })
}

func main() {
    d := &Demo{
        MySlice:  []string{"", "", ""},
        MySlice2: []string{"", "", ""},
        MyMap:    map[string]int{"": 0, "a": 100},
    }
    code, _, msgs := vfy.Check_(context.Background(), d, true)
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 mySlice[0] must not be blank
    // 1 mySlice[1] must not be blank
    // 2 mySlice[2] must not be blank
    // 3 mySlice2#1 must not be blank
    // 4 mySlice2#2 must not be blank
    // 5 mySlice2#3 must not be blank
    // 6 myMap$key must not be blank
    // 7 myMap$key's length must greater than 2
    // 8 myMap@value must be > 0 and < 100
    // 9 myMap$key's length must greater than 2
    // 10 myMap@value must be > 0 and < 100
}

```

## 默认消息

有些*校验方法*调用后可链式调用`DefaultMsg`方法，使用默认的错误消息。可自定义默认消息，设置方式为`vfy.SetDefaultMsg().<字段验证函数名>().<校验方法名>(<默认消息处理函数>）`

*代码示例*

```go
package main

import (
    "fmt"
    "regexp"

    vfy "github.com/jishaocong0910/gverify"
)

type Demo struct {
    Name  string
    Phone string
    Age   int
}

func init() {
    // 自定义默认消息
    vfy.SetDefaultMsg().String().NotBlank(func(ctx *vfy.Context) string {
        return fmt.Sprintf(`%s is blank`, ctx.FieldName())
    }).Gt(func(ctx *vfy.Context) string {
        return fmt.Sprintf(`%s is to short`, ctx.FieldName())
    })
}

func (d Demo) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &d.Name, "name").NotBlank().DefaultMsg()
    vfy.String(ctx, &d.Phone, "phone").Gt(10).DefaultMsg()
    vfy.Int(ctx, &d.Age, "age").Gt(0).DefaultMsg()
}

func main() {
    d := Demo{}
    code, _, msgs := vfy.Check_(nil, d, true)
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 name is blank
    // 1 phone is too short
    // 2 age must be greater than 0
}
```

## 错误码

有字段验证错误时，*验证函数*返回的错误码默认为字符串`ERROR`。只验证至首个错误字段时，即调用`vfy.Check`或`vfy.Check_(ctx, v, false)`，可通过链式调用`Msg_`和`DefaultMsg_`方法可指定错误码，它们与不带下划线的函数区别是，多了第一个参数用于指定错误码。

*代码示例*

```go
package main

import (
    "fmt"
    "regexp"

    vfy "github.com/jishaocong0910/gverify"
)

type Demo struct {
    Name  string
    Email string
    Phone string
}

func init() {
    vfy.SetDefaultMsg().String().NotBlank(func(ctx *vfy.Context) string {
        return fmt.Sprintf(`%s must not be blank`, ctx.FieldName())
    })
}

func (d Demo) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &d.Name, "name").NotBlank().DefaultMsg_("NAME_ERROR")
    vfy.String(ctx, &d.Email, "email").Regex(regexp.MustCompile(`\w+@\w+.\w+`)).Msg_("EMAIL_ERROR", `%s's format is illegal`, ctx.FieldName())
    vfy.String(ctx, &d.Phone, "phone").Gt(10).Msg_("PHONE_ERROR", "%s's length must be greater than %s", ctx.FieldName(), ctx.Confine(0))
}

func main() {
    d := Demo{}
    code, msg := vfy.Check(nil, d)
    fmt.Println("code:", code, ", msg:", msg)

    d2 := Demo{Name: "demo"}
    code2, msg2, _ := vfy.Check_(nil, d2, false)
    fmt.Println("code:", code2, ", msg:", msg2)

    d3 := Demo{Name: "demo", Email: "demo@demo.com"}
    code3, msg3, _ := vfy.Check_(nil, d3, true) // 验证所有字段时，自定义错误码无效。
    fmt.Println("code:", code3, ", msg:", msg3)
    // Output:
    // code: NAME_ERROR , msg: name must not be blank
    // code: EMAIL_ERROR , msg: email's format is illegal
    // code: ERROR , msg: phone's length must be greater than 10
}
```

# 代码风格

gverify的特色之一，是可避免结构体文件被代码生成器覆盖，导致校验规则丢失，其原理是通过在其他文件实现`vfy.Verifiable`
接口来避免，以下介绍各种实现接口的代码风格。

## 在结构体所在文件

对于非代码生成器生成的结构体，可直接在结构体所在文件实现`vfy.Verifiable`接口，上文中的代码示例都是这种风格。

## 在同目录其他文件

为避免代码生成器覆盖结构体所在的文件，可以在同包下另一个文件实现`vfy.Verifiable`接口。

*目录结构*

```
.
├─ gen_struct.go
└─ gen_struct_checklist.go
```

*gen_struct.go*

```go
package demo

type GenStruct struct {
    Id   string
    Name string
}
```

*gen_struct_checklist.go*

```go
package demo

import vfy "github.com/jishaocong0910/gverify"

func (g GenStruct) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &g.Id, "id").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    vfy.String(ctx, &g.Name, "name").NotBlank().Msg("%s must not be blank", ctx.FieldName())
}
```

## 在其他目录的文件

如果代码生成器会覆盖整个目录，或希望目录只存放生成的代码，则可以在其他包实现`vfy.Verifiable`接口。
由于Golang语法禁止在其他包增加结构体方法，因此需增加一个内嵌原结构体的新的结构体，通过新的结构体来验证。

*目录结构*

```
.
├─ model
│  └─ gen_struct.go
└─ check
   └─ gen_struct_checklist.go
```

*gen_struct.go*

```go
package model

type GenStruct struct {
    Id   string
    Name string
}
```

*gen_struct_checklist.go*

```go
package check

import (
    "demo/model"

    vfy "github.com/jishaocong0910/gverify"
)

type GenStruct struct {
    model.GenStruct
}

func (g GenStruct) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &g.Id, "id").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    vfy.String(ctx, &g.Name, "name").NotBlank().Msg("%s must not be blank", ctx.FieldName())
}
```

*代码示例*

```go
package main

import (
    "context"
    "fmt"

    "demo/check"
    "demo/model"

    vfy "github.com/jishaocong0910/gverify"
)

func main() {
    g := model.GenStruct{}
    code, _, msgs := vfy.Check_(context.Background(), &check.GenStruct{g}, true)
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 id must not be blank
    // 1 name must not be blank
}
```