# gverify

一款用于Golang的结构体校验工具。它通过手动编排校验过程进行校验，而非使用标签。使用标签的验证工具，无法用在经常被代码生成器覆盖的结构体，例如grpc、gorm生成的代码。gverify没有限制，并且简单易用，处理速度理论上比使用标签的校验工具更快。


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
        NotBlank().Msg("%s must not be blank", ctx.FieldName()).
        Regex(regexp.MustCompile(`^[0-9]{13}$`)).Msg("%s's format is illegal", ctx.FieldName())

    vfy.String(ctx, b.Description, "description").
        NotNil().Msg("%s must not be nil", ctx.FieldName())

    vfy.Int(ctx, &b.Stock, "stock").
        Range(0, 100).Msg("%s must between %s and %s", ctx.FieldName(), ctx.Confine(0), ctx.Confine(1))

    vfy.Float64(ctx, &b.Price, "price").
        Gt(0).Msg("%s must be greater than %s", ctx.FieldName(), ctx.Confine(0))

    vfy.String(ctx, b.Language, "language").
        Options([]string{"zh-cn", "en-US", "ja-JP"}).Msg("%s must be %s", ctx.FieldName(), ctx.Confines())

    vfy.Struct(ctx, b.Author, "author").
        NotNil().Msg("%s must not be nil", ctx.FieldName()).
        Dive()

    vfy.Slices(ctx, b.Categories, "categories").
        NotEmpty().Msg("%s must not be empty", ctx.FieldName()).
        Dive(func(t *Category) {
            vfy.Struct(ctx, t, "").
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
    ok, _, msgs := vfy.Check_(&b, true)
    if !ok {
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

# 入口函数

| 函数            | 说明                                      |
|---------------|-----------------------------------------|
| vfy.Check  | 检查至首个错误的字段，相当于`vfy.Check_(?, false)` |
| vfy.Check_ | 参数`all`为false则检查至首个错误字段 ，true则检查所有字段    |

# 编写结构体校验过程

结构体须实现`vfy.Verifiable`接口的`Checklist`方法，在其中编写每个字段的校验过程。方法的接收者为指针时，**无需担心它为nil**，入口函数会判断传入的值是否为nil，若为nil则会创建零值进行验证。

# 类型入口函数

每个字段须根据自身类型，选择对应的*类型入口函数*，创建具有*校验方法*的*校验变量*，再根据需要使用其中的方法。所有*类型入口函数*的第一个参数都传入`Checklist`方法的`ctx`参数。

| 类型入口函数      | 对应类型                      |
|-------------|---------------------------|
| vfy.Bool    | bool                      |
| vfy.Byte    | byte                      |
| vfy.Int     | int                       |
| vfy.Int8    | int8                      |
| vfy.Int16   | int16                     |
| vfy.Int32   | int32                     |
| vfy.Int64   | int64                     |
| vfy.Float32 | float32                   |
| vfy.Float64 | float64                   |
| vfy.Uint    | uint                      |
| vfy.Uint8   | uint8                     |
| vfy.Uint16  | uint16                    |
| vfy.Uint32  | uint32                    |
| vfy.Uint64  | uint64                    |
| vfy.String  | string                    |
| vfy.Slices  | slices                    |
| vfy.Map     | map                       |
| vfy.Struct  | struct（需实现vfy.Verifiable） |
| vfy.Any     | any                       |

# 校验方法

有些*校验方法*之间的名称，仅区别于是否具有下划线，例如`NotBlank`和`NotBlank_`
。带下划线的方法多一个参数用于指定值为nil时是否跳过，不带下划线的方法默认不跳过，例如`NotBlank()`相当于`NotBlank_(false)`
。下面列举不同类型创建的*校验变量*具有的*校验方法*，为节省篇幅不列举带下划线的方法。

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
        <td>slices map</td>
        <td>元素数量必须大于0</td>
    </tr>
    <tr>
        <td>Length</td>
        <td>string slices map</td>
        <td>必须等于指定值。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Min</td>
        <td rowspan="6">byte int int8 int16 int32 int64 float32 float64 uint uint8 uint16 uint32 uint64 string slices map</td>
        <td>必须大于等于指定值。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Max</td><td>必须小于等于指定值。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Range</td><td>必须在指定范围，包含边界。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Gt</td><td>必须大于指定值。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Lt</td><td>必须小于指定值。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
    </tr>
    <tr>
        <td>Within</td><td>必须在指定范围内，即不包含边界。对于string类型则比较utf8字符数，slices和map则比较元素数量</td>
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
        <td>struct slices map</td>
        <td>校验struct字段、slices元素、map元素的key和value</td>
    </tr>
</table>

# 错误消息

## 自定义消息

*校验方法*调用后可链式调用`Msg`方法来指定错误消息。错误消息使用Golang自带的格式化字符串，`Checklist`方法的`ctx`
参数提供了一些方法，可用于填充错误消息。

| ctx的方法    | 说明                                                                                                                                                                                                   |
|-----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FieldName | 返回具有路径的字段名称。例如：`title`、`author.name`、`categoryId[2].sort`。                                                                                                                                           |
| Confine   | 返回*校验方法*的指定索引的限制值的字符串形式。例如：对于`Max(10)`，`c.Confine(0)`返回`10`；对于`Range(5, 15)`，`c.Confine(0)`返回`5`，`c.Confine(1)`返回`15`。                                                                               |
| Confines  | 返回*校验方法*的所有限制值的字符串形式，用`,`拼接，若数量超过两个，则最后一个用`or`拼接。例如：对于`Options("zh-cn", "en-US")`，`ctx.Confines()`返回`zh-cn, en-US`；对于`Options("zh-cn", "en-US", "ja-JP")`，`ctx.Confines()`返回`zh-cn, en-US or ja-JP`。 |

## 元素字段名称

在slices和map的*校验变量*的`Dive`方法内校验元素值时，会预设*元素字段名称*，并且*类型入口函数*的`fieldName`参数无效。对于slices
设置为`<slices字段名称>[<索引>]`；对于map，key设置为`<map字段名称>$key`，value的设置为`<map字段名称>$value`。

*代码示例*
```go
package main

import (
    "fmt"

    vfy "github.com/jishaocong0910/gverify"
)

type Demo struct {
    Slice []string
    Map   map[string]int
}

func (d Demo) Checklist(ctx *vfy.Context) {
    vfy.Slices(ctx, d.Slice, "slice_field_name").Dive(func(t string) {
        // fieldName参数无效
        vfy.String(ctx, &t, "elem").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    })
    vfy.Map(ctx, d.Map, "map_field_name").Dive(func(k string) {
        // fieldName参数无效
        vfy.String(ctx, &k, "key").NotBlank().Msg("%s must not be blank", ctx.FieldName())
    }, func(v int) {
        // fieldName参数无效
        vfy.Int(ctx, &v, "value").Within(0, 100).Msg("%s must be > %s and < %s", ctx.FieldName(), ctx.Confine(0), ctx.Confine(1))
    })
}

func main() {
    d := &Demo{
        Slice: []string{"a", "", ""},
        Map:   map[string]int{"": 32, "a": 0, "b": 101},
    }
    ok, _, msgs := vfy.Check_(d, true)
    if !ok {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 slice_field_name[1] must not be blank
    // 1 slice_field_name[2] must not be blank
    // 2 map_field_name$key must not be blank
    // 3 map_field_name$value must be > 0 and < 100
    // 4 map_field_name$value must be > 0 and < 100
}
```

## 默认消息

有些*校验方法*调用后可链式调用`DefaultMsg`方法，使用默认的错误消息。**每个*校验方法*的默认错误消息必须进行设置**，否则默认消息是空字符串。默认消息的设置方式为`vfy.DefaultMsg().<类型入口函数名>.<校验方法名>(<默认消息处理函数>)）`

*代码示例*
```go
package main

import (
    "fmt"
    "regexp"

    vfy "github.com/jishaocong0910/gverify"
)

func init() {
    vfy.DefaultMsg().String().NotBlank(func(ctx *vfy.Context) string {
        return fmt.Sprintf("%s must not be blank", ctx.FieldName())
    })
}

type Book struct {
    Title string
    Isbn  string
}

func (b Book) Checklist(ctx *vfy.Context) {
    vfy.String(ctx, &b.Title, "title").NotBlank().DefaultMsg()
    vfy.String(ctx, &b.Isbn, "isbn").Regex(regexp.MustCompile(`^[0-9]{13}$`)).DefaultMsg()
}

func main() {
    b := Book{}
    ok, _, msgs := vfy.Check_(b, true)
    if !ok {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 title must not be blank
    // 1
    //
    // 由于没有设置string的Regex校验方法的默认消息，isbn字段的错误消息为空字符串。
}
```
# 代码风格

gverify的特色之一，是可避免结构体代码被生成器覆盖，导致校验规则丢失，它通过在其他文件实现`vfy.Verifiable`接口来避免，以下介绍在各种文件位置下实现的代码风格。

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

如果代码生成器会覆盖整个目录，或者想分开存放生成的和自定义的代码，可以在其他包实现`vfy.Verifiable`接口。 由于Golang语法禁止在其他增加结构体方法，因此需增加一个内嵌原结构体的新的结构体，通过新的结构体来验证。

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
    vfy "github.com/jishaocong0910/gverify"
    "demo/model"
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
    "fmt"

    vfy "github.com/jishaocong0910/gverify"
    "demo/check"
    "demo/model"
)

func main() {
    g := model.GenStruct{}
    ok, _, msgs := vfy.Check_(&check.GenStruct{g}, true)
    if !ok {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 id must not be blank
    // 1 name must not be blank
}
```