# Gverify

一款用于Golang的结构体校验工具。它不通过标签来指定校验规则，而是由代码指定。使用标签的校验工具需要修改结构体，在一些代码生成器生成的结构体上使用时，可能导致代码覆盖或冲突问题。gverify可实现零代码入侵，支持自定义错误码和错误消息，不使用反射，处理速度理论上比使用标签的校验工具快。

[![Go Reference](https://pkg.go.dev/badge/github.com/jishaocong0910/gverify.svg)](https://pkg.go.dev/github.com/jishaocong0910/gverify)
[![Go Report Card](https://goreportcard.com/badge/github.com/jishaocong0910/gverify)](https://goreportcard.com/report/github.com/jishaocong0910/gverify)
![coverage](https://raw.githubusercontent.com/jishaocong0910/gverify/badges/.badges/main/coverage.svg)

# 安装

```shell
go get github.com/jishaocong0910/gverify
```

# 用法与例子

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

func (b *Book) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &b.Title, "title").NotBlank().Max(10)
    vfy.String(vc, &b.Isbn, "isbn").NotBlank().Regex(regexp.MustCompile(`^[0-9]{13}$`))
    vfy.String(vc, b.Description, "description").Required()
    vfy.Int(vc, &b.Stock, "stock").Range(0, 100)
    vfy.Float64(vc, &b.Price, "price").Gt(0)
    vfy.String(vc, b.Language, "language").Enum([]string{"zh-cn", "en-US", "ja-JP"})
    vfy.Struct(vc, b.Author, "author").Required().Dive()
    vfy.Slice(vc, b.Categories, "categories").NotEmpty().Dive(func(e *Category) {
        vfy.Struct(vc, e, "").Dive()
    })
}

type Author struct {
    Name    string
    Resumes *string
}

func (a Author) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &a.Name, "name").NotBlank()
    vfy.String(vc, &a.Name, "resumes").NotBlank()
}

type Category struct {
    Id   string
    Sort int
}

func (c Category) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &c.Id, "id").NotBlank()
    vfy.Int(vc, &c.Sort, "sort").Max(999)
}

func main() {
    b := Book{Author: &Author{}, Categories: []*Category{{Sort: 127}, {Id: "c1", Sort: 1000}}}
    code, _, msgs := vfy.Check(context.Background(), &b, vfy.All())
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // 0 title must not be blank
    // 1 isbn must not be blank
    // 2 isbn's format is illegal
    // 3 description is required
    // 4 price must be greater than 0
    // 5 language must be zh-cn, en-US or ja-JP
    // 6 author.name must not be blank
    // 7 author.resumes must not be blank
    // 8 categories[0].id must not be blank
    // 9 categories[1].sort must not be greater than 999
}
```

# 校验入口

***校验函数***：`vfy.Check`

<table>
    <th>入参</th>
    <th>描述</th>
    <tr>
        <td width=170px>ctx context.Context</td>
        <td>可为nil</td>
    </tr>
    <tr>
        <td>s vfy.Verifiable</td>
        <td>待校验的结构体</td>
    </tr>
    <tr>
        <td>opts ...vfy.Option</td>
        <td>选项</td>
    </tr>
    
</table>

<table>
    <th>返回</th>
    <th>描述</th>
    <tr>
        <td width=160px>code string</td>
        <td>错误码，"SUCCESS"为校验成功，"FAIL"为校验失败，可使用常量<i>vfy.SUCCESS</i>和<i>vfy.FAIL</i>进行比较。</td>
    </tr>
    <tr>
        <td>msg string</td>
        <td>首个错误消息</td>
    </tr>
    <tr>
        <td>msgs []string</td>
        <td>所有错误消息</td>
    </tr>
</table>

| 选项      | 描述                              |
|---------|---------------------------------|
| vfy.All | 校验所有校验规则，不使用此选项时，首个校验错误出现既停止校验。 |

# Checklist方法

结构体通过实现*vfy.Verifiable*接口的*Checklist*方法，在其中编写字段校验规则，如果接收者为指针，**无需担心它为nil**，*校验函数*检测到传入的值为nil会创建零值进行校验。

# 字段校验

校验字段时，根据字段类型，调用对应的***字段类型函数***，传入<i>Checklist</i>方法的<i>vc</i>参数、字段值的指针和字段名称，再链式调用***校验方法***，可指定一些选项。

*e.g.*

```go
func (b *Book) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &b.Title, "title", vfy.Omittable()).NotBlank().Max(10, vfy.Code("TITLE_TOO_LONG"))
}
```

| 字段类型函数      | 对应类型    |
|-------------|---------|
| vfy.Bool    | bool    |
| vfy.Byte    | byte    |
| vfy.Int     | int     |
| vfy.Int8    | int8    |
| vfy.Int16   | int16   |
| vfy.Int32   | int32   |
| vfy.Int64   | int64   |
| vfy.Uint    | uint    |
| vfy.Uint8   | uint8   |
| vfy.Uint16  | uint16  |
| vfy.Uint32  | uint32  |
| vfy.Uint64  | uint64  |
| vfy.Float32 | float32 |
| vfy.Float64 | float64 |
| vfy.String  | string  |
| vfy.Slice   | slice   |
| vfy.Map     | map     |
| vfy.Struct  | struct  |
| vfy.Any     | any     |

<table>
    <tr>
        <th>校验方法</th>
        <th>描述</th>
    </tr>
    <tr>
        <td>Required</td>
        <td>必填，即不能为nil。</td>
    </tr>
    <tr>
        <td>NotBlank</td>
        <td>string不能为空白。nil值处理：视为空字符串。</td>
    </tr>
    <tr>
        <td>Regex</td>
        <td>string必须匹配正则表达式。nil值处理：视为空字符串。</td>
    </tr>
    <tr>
        <td>NotEmpty</td>
        <td>slice和map长度必须大于0。nil值处理：长度视为0。</td>
    </tr>
    <tr>
        <td>Length</td>
        <td>string字符长度、slice和map长度的固定值。nil值处理：视为0长度。</td>
    </tr>
    <tr>
        <td>Min</td>
        <td>数值、string字符长度、slice和map长度的最小值。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Max</td>
        <td>数值、string字符长度、slice和map长度的最大值。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Range</td>
        <td>数值、string字符长度、slice和map长度的范围值，包含边界。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Gt</td>
        <td>数值、string字符长度、slice和map长度，必须大于指定值。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Lt</td>
        <td>数值、string字符长度、slice和map长度，必须小于指定值。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Within</td>
        <td>数值、string字符长度、slice和map长度的范围值，不包含边界。nil值处理：数字视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Enum</td>
        <td>枚举值</td>
    </tr>
    <tr>
        <td>Custom</td>
        <td>自定义校验</td>
    </tr>
    <tr>
        <td>Dive</td>
        <td>下沉校验struct字段、slice元素、map的key和value。</td>
    </tr>
</table>

| 字段类型函数选项      | 描述       |
|---------------|----------|
| vfy.Omittable | 值为nil时忽略 |

| 校验方法选项   | 描述      |
|----------|---------|
| vfy.Code | 自定义错误码  |
| vfy.Msg  | 自定义错误消息 |

# 自定义错误码&错误消息

<i>vfy.Code</i>只在不使用<i>vfy.All()</i>选项的情况下有效。

<i>vfy.Msg</i>选项须传入一个函数，通过函数参数<i> f </i>的<i>Msg</i>方法设置错误消息，同时参数<i> f </i>还提供了一些动态参数用于构建错误消息。

*e.g.*

```go
func (b *Book) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &b.Title, "title").Max(10, vfy.Msg(func(f *vfy.FieldInfo) {
        f.Msg("%s is too long", f.FieldName(), f.Confine(0))
    }))
}
```

<table>
    <th width="130px">动态参数方法</th>
    <th>描述</th>
    <tr>
        <td>FieldName</td>
        <td>返回具有路径的字段名称。例如：<i>title</i>、<i>author.name</i>、<i>categoryId[2].sort</i>。</td>
    </tr>
    <tr>
        <td>Confine</td>
        <td>返回<i>校验方法</i>的指定索引的限制值的字符串形式。例如：对于<i>Max(10)</i>，<i>vc.Confine(0)</i>返回<i>10</i>；对于<i>Range(5, 15)</i>，<i>vc.Confine(0)</i>返回<i>5</i>，<i>vc.Confine(1)</i>返回<i>15</i>。</td>
    </tr>
    <tr>
        <td>Confines</td>
        <td>返回<i>校验项方法</i>的所有限制值的字符串形式，用","拼接，若数量超过两个，则最后一个用"or"拼接。例如：对于<i>Enum("zh-cn", "en-US")</i>，返回"zh-cn, en-US"；对于<i>Enum("zh-cn", "en-US", "ja-JP")</i>，返回"zh-cn, en-US or ja-JP"。</td>
    </tr>
</table>

# 元素的字段名称

在slice和map的<i>Dive</i>方法内校验元素值时，会预设字段名称并且无法更改，slice元素为<i><字段名称>[<索引>]</i>，对于map，key为<i><字段名称>$key</i>，value为<i><字段名称>$value</i>。

*e.g.*

```go
package main

import (
    "context"
    "fmt"

    vfy "github.com/jishaocong0910/gverify"
)

type Demo struct {
    MySlice []int
    MyMap   map[string]int
}

func (d Demo) Checklist(vc *vfy.VContext) {
    vfy.Slice(vc, d.MySlice, "mySlice").Dive(func(t int) {
        // fieldName参数无效
        vfy.Int(vc, &t, "elem").Gt(5)
    })
    vfy.Map(vc, d.MyMap, "myMap").Dive(func(k string) {
        // fieldName参数无效
        vfy.String(vc, &k, "mapKey").NotBlank()
    }, func(v int) {
        // fieldName参数无效
        vfy.Int(vc, &v, "mapValue").Gt(5)
    })
}

func main() {
    d := &Demo{
        MySlice: []int{4},
        MyMap:   map[string]int{"": 4},
    }
    code, _, msgs := vfy.Check(context.Background(), d, vfy.All())
    if code != vfy.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 mySlice[0] must be greater than 5
    // 1 myMap$key must not be blank
    // 2 myMap$value must be greater than 5
}
```

# 零代码入侵

零代码入侵是gverify的特色，可在不修改结构文件的情况下实现校验，原理是<i>Checklist</i>方法可以与结构体分开不同文件，有以下三种场景，一些场景可实现零代码入侵。

* 在结构体所在文件

对于非代码生成器生成的结构体，可直接在结构体所在文件实现<i>vfy.Verifiable</i>接口，上文中的代码示例都是这种风格。

* 在同目录其他文件

为了避免代码生成器覆盖文件，可以在同包下另一个文件实现<i>vfy.Verifiable</i>接口。

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

func (g GenStruct) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &g.Id, "id").NotBlank()
    vfy.String(vc, &g.Name, "name").NotBlank()
}
```

* 在其他目录的文件

也可以在其他包实现`vfy.Verifiable`接口。由于Golang语法禁止在其他包增加结构体方法，因此需增加一个内嵌原结构体的新的结构体，通过新的结构体来校验。

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

func (g GenStruct) Checklist(vc *vfy.VContext) {
    vfy.String(vc, &g.Id, "id").NotBlank()
    vfy.String(vc, &g.Name, "name").NotBlank()
}
```

*校验示例*

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