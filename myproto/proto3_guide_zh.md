# Proto3 语言指南

本文介绍如何在你的项目中使用 Protocol Buffers 语言的 proto3 版本。

本指南说明如何使用 protocol buffer 语言来组织你的 protocol buffer 数据，包括 .proto 文件语法，以及如何从 .proto 文件生成数据访问类。本文涵盖的是 protocol buffers 语言的 proto3 版本。

有关 editions 语法的信息，请参见 Protobuf Editions 语言指南。

有关 proto2 语法的信息，请参见 Proto2 语言指南。

这是一份参考指南。如果你想看一个循序渐进的示例，其中会用到本文描述的许多特性，请查看你所选语言的教程。

## 定义消息类型

先来看一个非常简单的例子。假设你想定义一种搜索请求消息格式，其中每个搜索请求都包含查询字符串、你感兴趣的结果页码，以及每页结果数量。下面就是定义该消息类型所使用的 .proto 文件。

```protobuf
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 results_per_page = 3;
}
```

文件的第一行说明你使用的是 protobuf 语言规范的 proto3 版本。

- edition（或者在 proto2/proto3 中使用的 syntax）必须是文件中第一条非空且非注释的语句。
- 如果没有指定 edition 或 syntax，protocol buffer 编译器会默认你使用的是 proto2。

SearchRequest 消息定义指定了 3 个字段（名称/值对），分别对应你希望在这种消息中包含的 3 类数据。每个字段都有名称和类型。

### 指定字段类型

在前面的示例中，所有字段都是标量类型：两个整数（page_number 和 results_per_page）以及一个字符串（query）。你也可以为字段指定枚举类型，以及其他消息类型这类复合类型。

### 分配字段编号

你必须为消息定义中的每个字段指定一个介于 1 和 536,870,911 之间的编号，同时满足以下限制：

- 该编号在该消息的所有字段中必须唯一。
- 19,000 到 19,999 之间的字段编号保留给 Protocol Buffers 实现使用。如果你在消息中使用这些保留编号，protocol buffer 编译器会报错。
- 你不能使用之前已保留的字段编号，也不能使用已经分配给扩展的字段编号。

一旦某个消息类型已经投入使用，这个编号就不能再更改，因为它用于在线上格式中标识字段。“修改”字段编号，等价于删除原字段，再创建一个类型相同但编号不同的新字段。如何正确执行此类操作，请参见后文的“删除字段”。

字段编号绝对不应重复使用。不要把已经放入 reserved 列表的字段编号再取出来，用于新的字段定义。请参见“重复使用字段编号的后果”。

最常设置的字段，建议使用 1 到 15 之间的字段编号。较小的字段编号在线上格式中占用的空间更小。例如，1 到 15 范围内的字段编号编码时只占 1 个字节；16 到 2047 范围内的字段编号编码时占 2 个字节。有关更多细节，请参见 Protocol Buffer Encoding。

#### 重复使用字段编号的后果

重复使用字段编号会使线上格式消息的解码产生歧义。

protobuf 的线上格式非常精简，无法检测一个字段是按某种定义编码、却按另一种定义解码的。

一个字段如果按某种定义进行编码，再按另一种定义解码，可能导致：

- 开发者花费大量时间调试
- 解析或合并错误（这还是最好情况）
- PII/SPII 泄露
- 数据损坏

字段编号被重复使用的常见原因包括：

- 给字段重新编号（有时只是为了让字段编号看起来更美观）。重新编号本质上等于删除并重新添加所有涉及到的字段，因此会造成不兼容的线上格式变更。
- 删除某个字段后，没有把它的编号保留下来，导致未来可能被重复使用。

字段编号之所以限制为 29 位而不是 32 位，是因为其中有 3 位要用于指定字段的线上格式。详情请参见 Encoding 主题。

### 指定字段基数

消息字段可以是以下几种之一：

- 单数字段：

在 proto3 中，单数字段有两种形式：

  - optional：（推荐）optional 字段只有两种可能状态：

    - 字段已设置，且包含一个显式设置或从线上数据中解析得到的值。它会被序列化到线上格式中。
    - 字段未设置，此时会返回默认值，并且不会被序列化到线上格式中。

    你可以检查该值是否曾被显式设置。

    为了最大程度兼容 protobuf editions 和 proto2，建议优先使用 optional，而不是隐式字段。

  - implicit：（不推荐）隐式字段没有显式的基数标签，其行为如下：

    - 如果该字段是消息类型，它的行为与 optional 字段完全相同。
    - 如果该字段不是消息类型，它有两种状态：

      - 字段被设置为非默认值（非零值），并且该值是显式设置或从线上解析得到的。它会被序列化到线上格式中。
      - 字段被设置为默认值（零值）。它不会被序列化到线上格式中。事实上，你无法判断这个默认值究竟是显式设置的、从线上解析得到的，还是根本没有提供。关于这一点的更多信息，请参见 Field Presence。

- repeated：这种字段类型在格式良好的消息中可以重复出现零次或多次。repeated 值的顺序会被保留。
- map：这是一种成对的键值字段类型。更多内容请参见“Maps”。

#### repeated 字段默认使用 packed 编码

在 proto3 中，标量数值类型的 repeated 字段默认采用 packed 编码。

有关 packed 编码的更多信息，请参见 Protocol Buffer Encoding。

#### 消息类型字段始终具有字段存在性

在 proto3 中，消息类型字段本身就已经具备字段存在性。因此，为该字段添加 optional 修饰符，并不会改变它的字段存在性行为。

下面代码中的 Message2 和 Message3，在所有语言中都会生成相同的代码，而且它们在二进制、JSON 和 TextFormat 中的表示也没有区别：

```protobuf
syntax="proto3";

package foo.bar;

message Message1 {}

message Message2 {
  Message1 foo = 1;
}

message Message3 {
  optional Message1 bar = 1;
}
```

#### 格式良好的消息

当“well-formed（格式良好）”一词用于 protobuf 消息时，它指的是被序列化或反序列化的字节。protoc 解析器会校验给定的 proto 定义文件是否可被正确解析。

单数字段在一段线上格式字节中可以出现多次。解析器会接受这样的输入，但在生成的绑定中，只有该字段最后一次出现的值是可访问的。关于这一点的更多信息，请参见 Last One Wins。

### 添加更多消息类型

一个 .proto 文件中可以定义多个消息类型。如果你定义的是多个彼此相关的消息，这会很有用。比如，如果你想定义与 SearchResponse 消息类型对应的回复消息格式，就可以把它添加到同一个 .proto 文件中：

```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 results_per_page = 3;
}

message SearchResponse {
 ...
}
```

把多个消息放在一起可能导致膨胀。虽然多个消息类型（例如 message、enum 和 service）都可以定义在同一个 .proto 文件中，但如果一个文件里定义了大量依赖各异的消息，也会造成依赖膨胀。因此，建议每个 .proto 文件中包含尽可能少的消息类型。

### 添加注释

要在 .proto 文件中添加注释：

- 建议在 .proto 代码元素前一行使用 C/C++/Java 风格的行尾注释 //。
- 也接受 C 风格的行内/多行注释 /* ... */。
- 使用多行注释时，建议每一行都带有前导的 *。

```protobuf
/**
 * SearchRequest 表示一个搜索查询，并包含分页选项，
 * 用于指示响应中应包含哪些结果。
 */
message SearchRequest {
  string query = 1;

  // 我们想要哪一页？
  int32 page_number = 2;

  // 每页返回多少条结果。
  int32 results_per_page = 3;
}
```

### 删除字段

如果删除字段的方式不正确，可能会引发严重问题。

当某个字段不再需要，并且客户端代码中对它的所有引用都已删除后，你可以从消息中删除该字段定义。但是，你必须保留这个已删除字段的字段编号。如果不保留，就有可能在将来被其他开发者重新使用。

你还应该保留该字段名，以便消息的 JSON 和 TextFormat 编码仍然可以继续被解析。

#### 保留字段编号

如果你更新某个消息类型时彻底删除了一个字段，或者只是把它注释掉，那么未来的开发者可能会在更新该类型时重复使用这个字段编号。这会造成严重问题，具体请参见“重复使用字段编号的后果”。为了避免这种情况，请把已删除字段的编号加入 reserved 列表。

如果将来有开发者尝试使用这些保留字段编号，protoc 编译器会生成错误信息。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
}
```

保留字段编号区间是包含端点的，也就是说，9 to 11 与 9、10、11 完全等价。

#### 保留字段名称

以后再次使用旧字段名通常是安全的，但在使用 TextProto 或 JSON 编码时，字段名会被序列化，因此会带来风险。为了避免这种风险，你可以把已删除字段的名称加入 reserved 列表。

保留名称只会影响 protoc 编译器的行为，而不会影响运行时行为，但有一个例外：TextProto 实现可能会在解析时直接丢弃使用保留名称的未知字段，而不会像处理其他未知字段那样抛出错误（目前只有 C++ 和 Go 实现会这样做）。运行时的 JSON 解析不受保留名称影响。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

注意：同一个 reserved 语句中不能混用字段名和字段编号。

### 从 .proto 中会生成什么？

当你对某个 .proto 运行 protocol buffer 编译器时，编译器会为你选择的语言生成对应代码，以便处理文件中描述的消息类型，包括获取和设置字段值、将消息序列化到输出流，以及从输入流解析消息。

- 对于 C++，编译器会根据每个 .proto 生成一个 .h 和 .cc 文件，其中为文件中描述的每个消息类型生成一个类。
- 对于 Java，编译器会生成一个 .java 文件，其中为每个消息类型生成一个类，同时还会生成一个专用的 Builder 类，用于创建消息类实例。
- 对于 Kotlin，除了 Java 生成代码外，编译器还会为每个消息类型生成一个 .kt 文件，提供改进后的 Kotlin API，其中包括简化消息实例创建的 DSL、可空字段访问器以及 copy 函数。
- Python 略有不同。Python 编译器会生成一个模块，其中包含你在 .proto 中定义的每个消息类型的静态描述符，然后通过元类在运行时创建所需的 Python 数据访问类。
- 对于 Go，编译器会生成一个 .pb.go 文件，其中为文件中的每个消息类型生成一个类型。
- 对于 Ruby，编译器会生成一个 .rb 文件，其中包含一个 Ruby 模块，模块中定义了你的消息类型。
- 对于 Objective-C，编译器会根据每个 .proto 生成一个 pbobjc.h 和一个 pbobjc.m 文件，其中为文件中描述的每个消息类型生成一个类。
- 对于 C#，编译器会根据每个 .proto 生成一个 .cs 文件，其中为文件中描述的每个消息类型生成一个类。
- 对于 PHP，编译器会为每个消息类型生成一个 .php 消息文件，并为每个被编译的 .proto 文件生成一个 .php 元数据文件。该元数据文件用于将合法消息类型加载到描述符池中。
- 对于 Dart，编译器会生成一个 .pb.dart 文件，其中为文件中的每个消息类型生成一个类。

你可以通过所选语言的教程进一步了解各语言 API 的用法。若需要更详细的 API 细节，请参见相应的 API 参考文档。

## 标量值类型

标量消息字段可以具有以下类型。下表列出了 .proto 文件中指定的类型以及相关说明：

| Proto 类型 | 说明 |
| --- | --- |
| double | 使用 IEEE 754 双精度格式。 |
| float | 使用 IEEE 754 单精度格式。 |
| int32 | 使用变长编码。对负数编码效率不高，如果字段可能出现负值，请改用 sint32。 |
| int64 | 使用变长编码。对负数编码效率不高，如果字段可能出现负值，请改用 sint64。 |
| uint32 | 使用变长编码。 |
| uint64 | 使用变长编码。 |
| sint32 | 使用变长编码。有符号整型值。相比普通 int32，它能更高效地编码负数。 |
| sint64 | 使用变长编码。有符号整型值。相比普通 int64，它能更高效地编码负数。 |
| fixed32 | 固定 4 字节。如果数值经常大于 2^28，则比 uint32 更高效。 |
| fixed64 | 固定 8 字节。如果数值经常大于 2^56，则比 uint64 更高效。 |
| sfixed32 | 固定 4 字节。 |
| sfixed64 | 固定 8 字节。 |
| bool | 布尔值。 |
| string | 字符串必须始终包含 UTF-8 编码或 7 位 ASCII 文本，长度不能超过 2^32。 |
| bytes | 可包含任意字节序列，长度不能超过 2^32。 |

下表展示了这些 Proto 类型在自动生成类中对应的语言类型：

| Proto 类型 | C++ 类型 | Java/Kotlin 类型[1] | Python 类型[3] | Go 类型 | Ruby 类型 | C# 类型 | PHP 类型 | Dart 类型 | Rust 类型 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| double | double | double | float | float64 | Float | double | float | double | f64 |
| float | float | float | float | float32 | Float | float | float | double | f32 |
| int32 | int32_t | int | int | int32 | Fixnum 或 Bignum（按需） | int | integer | int | i32 |
| int64 | int64_t | long | int/long[4] | int64 | Bignum | long | integer/string[6] | Int64 | i64 |
| uint32 | uint32_t | int[2] | int/long[4] | uint32 | Fixnum 或 Bignum（按需） | uint | integer | int | u32 |
| uint64 | uint64_t | long[2] | int/long[4] | uint64 | Bignum | ulong | integer/string[6] | Int64 | u64 |
| sint32 | int32_t | int | int | int32 | Fixnum 或 Bignum（按需） | int | integer | int | i32 |
| sint64 | int64_t | long | int/long[4] | int64 | Bignum | long | integer/string[6] | Int64 | i64 |
| fixed32 | uint32_t | int[2] | int/long[4] | uint32 | Fixnum 或 Bignum（按需） | uint | integer | int | u32 |
| fixed64 | uint64_t | long[2] | int/long[4] | uint64 | Bignum | ulong | integer/string[6] | Int64 | u64 |
| sfixed32 | int32_t | int | int | int32 | Fixnum 或 Bignum（按需） | int | integer | int | i32 |
| sfixed64 | int64_t | long | int/long[4] | int64 | Bignum | long | integer/string[6] | Int64 | i64 |
| bool | bool | boolean | bool | bool | TrueClass/FalseClass | bool | boolean | bool | bool |
| string | std::string | String | str/unicode[5] | string | String（UTF-8） | string | string | String | ProtoString |
| bytes | std::string | ByteString | str（Python 2）/bytes（Python 3） | []byte | String（ASCII-8BIT） | ByteString | string | List | ProtoBytes |

[1] Kotlin 对无符号类型也沿用 Java 中对应的类型，以确保在 Java/Kotlin 混合代码库中的兼容性。

[2] 在 Java 中，无符号 32 位和 64 位整数通过其有符号对应类型来表示，最高位只是存放在符号位中。

[3] 无论哪种情况，给字段赋值时都会进行类型检查，以确保值是合法的。

[4] 64 位整数或无符号 32 位整数在解码后总是表示为 long；但在设置字段时，如果传入的是 int，也可以是 int。无论哪种情况，设置的值都必须能放进对应表示类型的范围内。参见 [2]。

[5] Python 字符串在解码时表示为 unicode；如果给定的是 ASCII 字符串，也可以是 str（这一点未来可能会改变）。

[6] 在 64 位机器上使用 integer，在 32 位机器上使用 string。

有关这些类型在消息序列化时如何编码的更多信息，请参见 Protocol Buffer Encoding。

## 默认字段值

当消息被解析时，如果编码后的消息字节中不包含某个字段，那么访问解析后对象中的该字段时，会得到该字段的默认值。默认值随类型而不同：

- 对于字符串，默认值为空字符串。
- 对于 bytes，默认值为空字节串。
- 对于 bool，默认值为 false。
- 对于数值类型，默认值为 0。
- 对于消息字段，该字段处于未设置状态。其具体值依赖于语言实现。详情请参见生成代码指南。
- 对于枚举，默认值是枚举中定义的第一个值，而且该值必须为 0。参见“枚举默认值”。

repeated 字段的默认值为空（通常是对应语言中的空列表）。

map 字段的默认值也为空（通常是对应语言中的空映射）。

需要注意的是，对于隐式存在性的标量字段，一旦消息被解析，就无法判断该字段究竟是被显式设置为了默认值，还是压根没有设置。例如，布尔字段是被设为了 false，还是根本没提供，这是无法区分的。因此在定义消息类型时必须考虑这一点。比如，如果你不希望某个行为在默认情况下发生，就不要设计一个“当布尔值为 false 时触发该行为”的字段，因为默认值本身就是 false。另外还要注意，如果标量消息字段被设置为其默认值，该值不会被序列化到线上格式中。如果 float 或 double 被设置为 +0，它不会被序列化；但 -0 被认为是不同的值，因此会被序列化。

关于默认值在各语言生成代码中的更多细节，请参见对应语言的生成代码指南。

## 枚举

在定义消息类型时，你可能希望某个字段只能取预定义值列表中的一个。例如，假设你想给每个 SearchRequest 增加一个 corpus 字段，它的取值可以是 UNIVERSAL、WEB、IMAGES、LOCAL、NEWS、PRODUCTS 或 VIDEO。要做到这一点，只需在消息定义中加入一个 enum，并为每个可能的值定义一个常量。

下面的示例中，我们增加了一个名为 Corpus 的 enum，并定义了一个类型为 Corpus 的字段：

```protobuf
enum Corpus {
  CORPUS_UNSPECIFIED = 0;
  CORPUS_UNIVERSAL = 1;
  CORPUS_WEB = 2;
  CORPUS_IMAGES = 3;
  CORPUS_LOCAL = 4;
  CORPUS_NEWS = 5;
  CORPUS_PRODUCTS = 6;
  CORPUS_VIDEO = 7;
}

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 results_per_page = 3;
  Corpus corpus = 4;
}
```

### 为枚举值添加前缀

给枚举值添加前缀时，去掉前缀后的剩余名称仍应是合法且符合风格约定的枚举名称。例如，应避免下面这种写法：

```protobuf
enum DeviceTier {
  DEVICE_TIER_UNKNOWN = 0;
  DEVICE_TIER_1 = 1;
  DEVICE_TIER_2 = 2;
}
```

更好的做法是使用 DEVICE_TIER_TIER1 这样的名称，其中 DEVICE_TIER_ 被看作枚举值的作用域前缀，而不是单个枚举值名称的一部分。某些 Protobuf 实现会在安全的情况下自动剥离与所属枚举名称相匹配的前缀，但在这个例子中无法这么做，因为单独的 1 不是合法的枚举值名称。

未来某个 Edition 计划增加对作用域枚举的支持，这样就不再需要为每个枚举值手工加前缀，也就可以更简洁地写成 TIER1 = 1。

### 枚举默认值

SearchRequest.corpus 字段的默认值是 CORPUS_UNSPECIFIED，因为它是该枚举中定义的第一个值。

在 proto3 中，枚举定义中的第一个值必须是 0，并且名称应为 ENUM_TYPE_NAME_UNSPECIFIED 或 ENUM_TYPE_NAME_UNKNOWN。原因如下：

- 必须有一个值为 0 的枚举项，这样我们才能把 0 作为数值默认值。
- 这个零值必须是第一个元素，以兼容 proto2 的语义。在 proto2 中，除非显式指定其他值，否则第一个枚举值就是默认值。

同时也建议这个第一个默认值不要承载任何语义，只表示“该值未指定”。

### 枚举值别名

你可以通过为不同的枚举常量赋予相同数值来定义别名。要这样做，需要把 allow_alias 选项设置为 true。否则，当发现别名时，protocol buffer 编译器会生成警告信息。虽然所有别名值都可以合法序列化，但反序列化时只会使用第一个值。

```protobuf
enum EnumAllowingAlias {
  option allow_alias = true;
  EAA_UNSPECIFIED = 0;
  EAA_STARTED = 1;
  EAA_RUNNING = 1;
  EAA_FINISHED = 2;
}

enum EnumNotAllowingAlias {
  ENAA_UNSPECIFIED = 0;
  ENAA_STARTED = 1;
  // ENAA_RUNNING = 1;  // 取消这一行注释会导致出现警告信息。
  ENAA_FINISHED = 2;
}
```

枚举常量必须位于 32 位整数范围之内。由于 enum 在线上使用 varint 编码，因此负值编码效率较低，不推荐使用。你既可以像前面的例子那样在消息定义内部定义 enum，也可以在外部定义。定义在外部的 enum 可以在同一 .proto 文件中的任意消息定义里复用。你还可以把某个消息中声明的 enum 类型，作为另一个消息字段的类型，语法为 _MessageType_._EnumType_。

当你对使用了 enum 的 .proto 运行 protocol buffer 编译器时，Java、Kotlin 和 C++ 生成的代码中会有对应的 enum；而 Python 则会生成一个特殊的 EnumDescriptor 类，在运行时生成的类中据此创建一组带整数值的符号常量。

#### 重要说明

生成的代码可能会受到各语言对枚举成员数量限制的影响（有些语言只支持几千个枚举成员）。请检查你计划使用的语言所对应的限制。

反序列化过程中，未识别的枚举值会保留在消息中，但它在反序列化后的表现形式依赖于具体语言。在支持开放枚举类型、允许取值超出已定义符号范围的语言中，例如 C++ 和 Go，未知的枚举值会直接以其底层整数表示存储。在使用封闭枚举类型的语言中，例如 Java，枚举中会有一个特殊分支来表示未识别值，并通过专用访问器访问其底层整数值。无论哪种情况，如果消息再次被序列化，这个未识别值仍会随消息一起被序列化。

#### 重要说明

如果你想了解“枚举应该如何工作”与“当前各语言中实际如何工作”之间的差异，请参见 Enum Behavior。

如果你想进一步了解如何在应用程序中使用消息 enum，请参见你所选语言的生成代码指南。

### 保留值

如果你更新某个 enum 类型时彻底移除了一个枚举项，或者只是把它注释掉，那么未来的用户可能会在更新该类型时复用这个数值。这会在他们之后加载同一个 .proto 的旧实例时引发严重问题，包括数据损坏、隐私漏洞等。为了避免这种情况，可以把已删除项的数值（以及名称，因为名称在 JSON 序列化时也可能引发问题）声明为 reserved。如果将来有人尝试使用这些标识符，protocol buffer 编译器会报错。你还可以使用 max 关键字，把保留数值范围一直写到可能的最大值。

```protobuf
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

注意：同一个 reserved 语句中不能混用字段名和数值。

## 使用其他消息类型

你可以把其他消息类型用作字段类型。例如，假设你希望每个 SearchResponse 消息中都包含 Result 消息，那么就可以在同一个 .proto 中定义一个 Result 消息类型，然后在 SearchResponse 中声明一个类型为 Result 的字段：

```protobuf
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

### 导入定义

在前面的例子中，Result 消息类型与 SearchResponse 定义在同一个文件中。如果你想用作字段类型的消息已经定义在另一个 .proto 文件里，该怎么办？

你可以通过 import 其他 .proto 文件来使用其中的定义。要导入另一个 .proto 中的定义，只需要在文件顶部添加 import 语句：

```protobuf
import "myproject/other_protos.proto";
```

protobuf 编译器会在一组由 -I 或 --proto_path 标志指定的目录中查找被导入文件。import 语句中给出的路径会相对于这些目录进行解析。关于如何使用编译器的更多信息，请参见“生成代码”。

例如，假设有如下目录结构：

```text
my_project/
├── protos/
│   ├── main.proto
│   └── common/
│       └── timestamp.proto
```

如果要在 main.proto 中使用 timestamp.proto 的定义，应当在 my_project 目录下运行编译器，并设置 --proto_path=protos。这样一来，main.proto 中的 import 语句应写为：

```protobuf
// 位于 my_project/protos/main.proto
import "common/timestamp.proto";
```

一般来说，你应该把 --proto_path 设置为包含 protos 的最高层目录。很多时候这就是项目根目录，而在这个例子里则是单独的 /protos 目录。

默认情况下，你只能使用直接 import 的 .proto 文件中的定义。不过，有时候你可能需要把某个 .proto 文件移动到新位置。此时，与其直接移动文件并在一次改动中更新所有调用方，不如在旧位置保留一个占位用的 .proto 文件，并用 import public 将所有导入转发到新位置。

注意：Java 中的 public import 功能，在迁移整个 .proto 文件或使用 java_multiple_files = true 时效果最好。在这些情况下，生成名称可以保持稳定，从而避免你更新代码中的引用。如果只迁移 .proto 文件中的一部分内容，且未使用 java_multiple_files = true，虽然技术上也能工作，但通常仍需要同时更新大量引用，因此对迁移帮助可能并不明显。该功能在 Kotlin、TypeScript、JavaScript 和 GCL 中不可用。

带有 import public 的依赖，可以被任何导入该 proto 的代码以传递方式依赖。例如：

```protobuf
// new.proto
// 所有定义都迁移到这里

// old.proto
// 这是所有客户端正在导入的 proto。
import public "new.proto";
import "other.proto";

// client.proto
import "old.proto";
// 你可以使用 old.proto 和 new.proto 中的定义，但不能使用 other.proto 中的定义
```

### 使用 proto2 消息类型

你可以导入 proto2 的消息类型，并在 proto3 消息中使用它们，反之亦然。但是，proto2 的 enum 不能直接在 proto3 语法中使用（如果某个导入的 proto2 消息内部使用了这些 enum，则没问题）。

## 嵌套类型

你可以在消息类型内部定义并使用其他消息类型，例如下面这个例子中，Result 消息定义在 SearchResponse 消息内部：

```protobuf
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

如果你希望在父消息类型外部复用这个消息类型，可以用 _Parent_._Type_ 的形式引用它：

```protobuf
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

消息可以任意深度嵌套。在下面的例子中，两个名为 Inner 的嵌套类型彼此完全独立，因为它们定义在不同的消息内部：

```protobuf
message Outer {       // 第 0 层
  message MiddleAA {  // 第 1 层
    message Inner {   // 第 2 层
      int64 ival = 1;
      bool  booly = 2;
    }
  }
  message MiddleBB {  // 第 1 层
    message Inner {   // 第 2 层
      int32 ival = 1;
      bool  booly = 2;
    }
  }
}
```

## 更新消息类型

如果某个现有消息类型已经不能满足你的需求了，例如你想为消息格式新增一个字段，但又希望继续使用旧格式生成的代码，不用担心。只要你使用的是二进制线上格式，更新消息类型而不破坏已有代码是很容易做到的。

### 注意

如果你使用 ProtoJSON 或 proto text format 来存储 protocol buffer 消息，那么你在 proto 定义中可以进行的变更会有所不同。ProtoJSON 线上格式中的安全变更有专门说明。

请参考 Proto 最佳实践，并遵守以下规则：

### 二进制线上不安全的变更

线上不安全的变更，是指当你用新 schema 的解析器去解析由旧 schema 序列化的数据时，或者反过来时，会直接出错的 schema 变更。只有在你明确知道所有数据的序列化器和反序列化器都已经切换到新 schema 时，才应该进行这种变更。

- 修改任何现有字段的字段编号都不安全。
  - 修改字段编号等价于删除该字段并添加一个同类型的新字段。如果你确实想重新编号，请参见删除字段的说明。
- 将字段移动进一个现有的 oneof 中是不安全的。

### 二进制线上安全的变更

线上安全的变更，是指按照这种方式演进 schema 时，不会带来数据丢失风险，也不会引入新的解析失败。

请注意，任何线上安全的变更，在某种语言的应用代码层面依然可能是破坏性变更。例如，向已有枚举新增一个值，会让所有对该枚举执行穷举式 switch 的代码出现编译错误。因此，Google 在公共消息上有时会避免进行这类变更；相关 AIP 文档中对哪些变更在这种场景下是安全的给出了指导。

- 添加新字段是安全的。
  - 如果你添加了新字段，使用旧消息格式的代码所序列化出来的任何消息，仍然可以被新生成的代码解析。你应当留意这些字段的默认值，以便新代码能够正确处理由旧代码生成的消息。同样地，由新代码创建的消息也可以被旧代码解析：旧二进制在解析时会直接忽略新字段。详情请参见“未知字段”。
- 删除字段是安全的。
  - 被删除字段对应的字段编号绝不能在更新后的消息类型中再次使用。你可以考虑改为给该字段重命名，例如加上 OBSOLETE_ 前缀，或者直接把该字段编号设为 reserved，以避免未来 .proto 的使用者意外重用这个编号。
- 向枚举添加额外的值是安全的。
- 将单个显式存在性字段或扩展，变成一个新的 oneof 的成员，是安全的。
- 将只包含一个字段的 oneof 改为显式存在性字段，是安全的。
- 将某个字段改为具有相同编号和相同类型的扩展，是安全的。

### 二进制线上兼容的变更（有条件安全）

与“线上安全”不同，“线上兼容”表示同一份数据在变更前后都能被解析，但在这种变更形态下，解析过程可能会有损。例如，把 int32 改成 int64 是兼容变更；但是如果写入了一个大于 INT32_MAX 的值，那么仍按 int32 读取它的客户端会丢弃高位比特。

只有在你能非常小心地控制系统中的发布过程时，才应该进行兼容性变更。比如，你可以先把 int32 改成 int64，但在新 schema 部署到所有端点之前，继续只写入合法的 int32 值；等全量发布完成后，再开始写入更大的值。

如果你的 schema 会在组织外部发布，那么通常不应进行线上兼容变更，因为你无法控制新 schema 的部署进度，也就无法知道何时才能安全地使用更宽的取值范围。

- int32、uint32、int64、uint64 和 bool 彼此兼容。
  - 如果从线上解析出的数值无法放进目标类型中，那么结果就和在 C++ 中把这个数字强制转换成该类型时一样。例如，一个 64 位数如果被按 int32 读取，就会被截断到 32 位。
- sint32 和 sint64 彼此兼容，但它们与其他整数类型不兼容。
  - 如果写入的值在 INT_MIN 到 INT_MAX 之间（包含端点），那么无论按哪种类型解析，得到的值都相同。如果一个超出该范围的 sint64 值被按 sint32 解析，varint 会先被截断到 32 位，然后再执行 zigzag 解码，这将导致观察到的值发生变化。
- 只要字节内容是合法 UTF-8，string 与 bytes 就是兼容的。
- 如果 bytes 中包含某个消息实例的编码结果，那么嵌入消息与 bytes 也是兼容的。
- fixed32 与 sfixed32 兼容，fixed64 与 sfixed64 兼容。
- 对于 string、bytes 和消息字段，singular 与 repeated 兼容。
  - 如果把 repeated 字段的序列化数据输入给一个期待该字段为 singular 的客户端，那么如果该字段是基本类型，客户端会取最后一个输入值；如果该字段是消息类型，客户端会将所有输入元素进行合并。注意，这对数值类型通常并不安全，包括 bool 和 enum。数值类型的 repeated 字段默认采用 packed 格式序列化，而当解析器期待的是 singular 字段时，它无法正确解析这种格式。
- enum 与 int32、uint32、int64 和 uint64 兼容。
  - 请注意，消息反序列化后，客户端代码对它们的处理可能不同。例如，未识别的 proto3 枚举值会被保留在消息中，但具体如何表示依赖于语言实现。
- 在 map<K, V> 与对应的 repeated 消息字段之间切换，是二进制兼容的（下面的 Maps 部分会介绍其消息布局及其他限制）。
  - 不过，这种变更是否安全仍取决于应用本身：在反序列化后再次序列化时，使用 repeated 字段定义的客户端会产生语义等价的结果；但使用 map 字段定义的客户端可能会重排条目顺序，并丢弃重复键对应的条目。

## 未知字段

未知字段是格式良好的 protocol buffer 序列化数据，它表示解析器无法识别的字段。例如，当旧版本二进制解析由新版本二进制发送、且包含新字段的数据时，这些新字段在旧版本中就会变成未知字段。

Proto3 消息会保留未知字段，并在解析和序列化输出时一并包含它们，这一点与 proto2 的行为一致。

### 保留未知字段

某些操作会导致未知字段丢失。例如，执行以下任一操作时，未知字段都会丢失：

- 把 proto 序列化成 JSON。
- 遍历消息中的所有字段，并逐字段填充到一个新消息中。

为了避免丢失未知字段，请遵循以下做法：

- 使用二进制格式；在数据交换中尽量避免使用文本格式。
- 使用面向消息的 API（例如 CopyFrom() 和 MergeFrom()）复制数据，而不要按字段逐个拷贝。

TextFormat 是一个稍有特殊的情况。序列化为 TextFormat 时，未知字段会以字段编号的形式输出；但如果再把这种 TextFormat 数据解析回二进制 proto，只要其中存在按字段编号表示的条目，解析就会失败。

## Any

Any 消息类型允许你把消息作为嵌入类型使用，即使你没有该消息的 .proto 定义。Any 中包含一个任意消息的序列化 bytes，以及一个 URL，该 URL 既作为该消息类型的全局唯一标识符，也用于解析该类型。要使用 Any 类型，你需要导入 google/protobuf/any.proto。

```protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

给定某个消息类型时，默认的类型 URL 形式为 type.googleapis.com/_packagename_._messagename_。

不同语言实现通常都会在运行时库中提供类型安全的辅助方法，用于打包和解包 Any 值。例如，在 Java 中，Any 类型会提供专门的 pack() 和 unpack() 访问器；在 C++ 中则有 PackFrom() 和 UnpackTo() 方法：

```cpp
// 将任意消息类型存入 Any。
NetworkErrorDetails details = ...;
ErrorStatus status;
status.add_details()->PackFrom(details);

// 从 Any 中读取任意消息。
ErrorStatus status = ...;
for (const google::protobuf::Any& detail : status.details()) {
  if (detail.Is<NetworkErrorDetails>()) {
    NetworkErrorDetails network_error;
    detail.UnpackTo(&network_error);
    ... processing network_error ...
  }
}
```

## Oneof

如果一个消息中有很多单数字段，并且任意时刻最多只会设置其中一个字段，那么你可以使用 oneof 特性来强制这一约束，并节省内存。

oneof 字段类似 optional 字段，但 oneof 中的所有字段共享同一块内存，并且同一时刻最多只能设置一个字段。设置 oneof 中的任意一个成员时，其他成员都会被自动清除。你可以根据所选语言，使用特殊的 case() 或 WhichOneof() 方法检查 oneof 中当前设置的是哪个值（如果有的话）。

请注意，如果同时设置了多个值，那么最终生效的是根据 proto 中顺序判定的最后一个值，它会覆盖此前所有值。

oneof 中各字段的字段编号，必须在其外层消息中保持唯一。

### 使用 Oneof

要在 .proto 中定义 oneof，可以使用 oneof 关键字，后面跟上 oneof 的名称，这里是 test_oneof：

```protobuf
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

然后把各个 oneof 字段添加到该 oneof 定义中。你可以向 oneof 中添加除 map 字段和 repeated 字段之外的任意类型。如果确实需要把 repeated 字段放进 oneof，可以通过引入一个包含 repeated 字段的消息来实现。

在生成的代码中，oneof 字段和普通字段拥有相同的 getter 与 setter。你还会获得一个专门的方法，用来检查当前 oneof 中设置的是哪个值（如果有的话）。更多内容请参见所选语言对应的 API 参考文档。

### Oneof 特性

- 设置 oneof 中的某个字段时，会自动清除 oneof 中的所有其他成员。因此，如果你连续设置多个 oneof 字段，最终只有最后一次设置的字段仍然有值。

```cpp
SampleMessage message;
message.set_name("name");
CHECK_EQ(message.name(), "name");
// 调用 mutable_sub_message() 会清除 name 字段，
// 并把 sub_message 设置为一个新的 SubMessage 实例，且其字段都未设置。
message.mutable_sub_message();
CHECK(message.name().empty());
```

- 如果解析器在线上数据中遇到了同一个 oneof 的多个成员，那么解析后消息中只会保留最后看到的那个成员。在线上数据解析时，从字节流开头开始，按顺序处理下一个值，并应用以下解析规则：

  - 首先，检查同一个 oneof 中是否已经设置了另一个字段；如果有，就先清除它。
  - 然后，按“该字段不在 oneof 中”时的规则来应用其内容：
    - 基本类型会覆盖已经设置的值。
    - 消息类型会与已设置值进行合并。

- oneof 不能是 repeated。
- 反射 API 对 oneof 字段同样适用。
- 如果你把 oneof 字段设置为默认值（例如将 int32 类型的 oneof 字段设为 0），该字段的“case”仍会被设置，并且该值会被序列化到线上格式中。
- 如果你使用的是 C++，请确保代码不会引发内存崩溃。下面这段示例代码会崩溃，因为调用 set_name() 方法后，sub_message 已经被删除了。

```cpp
SampleMessage message;
SubMessage* sub_message = message.mutable_sub_message();
message.set_name("name");      // 会删除 sub_message
sub_message->set_...            // 这里会崩溃
```

- 同样在 C++ 中，如果你对两个带有 oneof 的消息执行 Swap()，每个消息最终都会带上另一个消息的 oneof case。下面这个例子里，msg1 最终会拥有一个 sub_message，而 msg2 会拥有一个 name。

```cpp
SampleMessage msg1;
msg1.set_name("name");
SampleMessage msg2;
msg2.mutable_sub_message();
msg1.swap(&msg2);
CHECK(msg1.has_sub_message());
CHECK_EQ(msg2.name(), "name");
```

### 向后兼容问题

在添加或删除 oneof 字段时一定要小心。如果检查一个 oneof 的值返回 None 或 NOT_SET，这既可能意味着该 oneof 从未被设置，也可能意味着它被设置成了该 oneof 的另一个版本中的某个字段。由于你无法知道线上数据里的未知字段是否属于该 oneof，因此没有办法区分这两种情况。

#### 标签复用问题

- 把单数字段移入或移出 oneof：消息序列化并再次解析后，你可能会丢失一部分信息（某些字段会被清除）。不过，把单个字段安全地移入一个新的 oneof 是可行的；如果已知多个字段中始终只会设置一个，那么也许也能安全迁移多个字段。详情请参见“更新消息类型”。
- 删除一个 oneof 字段后又重新添加：这可能会导致当前已设置的 oneof 字段在消息序列化并再次解析后被清除。
- 拆分或合并 oneof：这类问题与移动单数字段类似。

## Maps

如果你想在数据定义中创建关联映射，protocol buffers 提供了一种便捷的简写语法：

```protobuf
map<key_type, value_type> map_field = N;
```

其中，key_type 可以是任意整数类型或字符串类型，也就是任意标量类型，但不包括浮点类型和 bytes。注意，enum 和 proto 消息都不能作为 key_type。value_type 则可以是除另一个 map 之外的任意类型。

例如，如果你想创建一个项目映射，使每个 Project 消息都关联到一个字符串键，可以这样定义：

```protobuf
map<string, Project> projects = 3;
```

### Maps 的特性

- map 字段不能是 repeated。
- 线上格式中的顺序以及 map 值的迭代顺序都未定义，因此你不能依赖 map 条目具有某种固定顺序。
- 生成 .proto 的文本格式时，map 会按键排序。数字键按数值顺序排序。
- 从线上格式解析或执行合并时，如果出现重复的 map 键，最后看到的那个键对应的值会生效。若从文本格式解析 map，遇到重复键时可能会解析失败。
- 如果你为某个 map 字段提供了 key，但没有提供 value，那么该字段在序列化时的行为依赖于语言。在 C++、Java、Kotlin 和 Python 中，会序列化该类型的默认值；而在其他语言中则不会序列化任何内容。
- 在与 map foo 相同的作用域中，不能存在名为 FooEntry 的符号，因为 FooEntry 已被 map 的底层实现占用。

目前所有受支持语言都提供了生成后的 map API。更多内容请参见你所选语言的 API 参考文档。

### 向后兼容性

在线上格式中，map 语法等价于下面这种写法，因此即便某些 protocol buffers 实现本身不支持 map，也仍然可以处理你的数据：

```protobuf
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}

repeated MapFieldEntry map_field = N;
```

任何支持 map 的 protocol buffers 实现，都必须既能生成，也能接受可被这种较早定义方式接受的数据。

## 包

你可以在 .proto 文件中添加一个可选的 package 指定符，以避免 protocol 消息类型之间发生命名冲突。

```protobuf
package foo.bar;
message Open { ... }
```

之后，你就可以在定义消息字段时使用 package 指定符：

```protobuf
message Foo {
  ...
  foo.bar.Open open = 1;
  ...
}
```

package 指定符对生成代码的影响取决于你所选的语言：

- 在 C++ 中，生成的类会放在对应的 C++ 命名空间下。例如，Open 会位于 foo::bar 命名空间中。
- 在 Java 和 Kotlin 中，package 会被当作 Java 包名，除非你在 .proto 文件中显式提供了 java_package 选项。
- 在 Python 中，package 指令会被忽略，因为 Python 模块是根据其在文件系统中的位置来组织的。
- 在 Go 中，package 指令会被忽略，生成的 .pb.go 文件将位于与对应 go_proto_library Bazel 规则同名的包中。对于开源项目，你必须提供 go_package 选项，或者设置 Bazel 的 -M 标志。
- 在 Ruby 中，生成的类会被包裹在嵌套的 Ruby 命名空间中，并转换为 Ruby 所要求的大小写风格（首字母大写；如果首字符不是字母，则前置 PB_）。例如，Open 会位于 Foo::Bar 命名空间中。
- 在 PHP 中，package 会先转换为 PascalCase，然后作为命名空间使用，除非你在 .proto 文件中显式提供了 php_namespace 选项。例如，Open 会位于 Foo\Bar 命名空间中。
- 在 C# 中，package 会先转换为 PascalCase，然后作为命名空间使用，除非你在 .proto 文件中显式提供了 csharp_namespace 选项。例如，Open 会位于 Foo.Bar 命名空间中。

请注意，即便 package 指令不会直接影响生成代码，例如在 Python 中，它仍然强烈建议被显式指定。否则，它可能导致描述符中的命名冲突，也会让该 proto 难以被移植到其他语言中使用。

### 包与名称解析

protocol buffer 语言中的类型名解析方式与 C++ 类似：先查找最内层作用域，再查找外一层，以此类推；每一级 package 都被视为其父 package 的“内部”作用域。前导的 .（例如 .foo.bar.Baz）表示应当从最外层作用域开始查找。

protocol buffer 编译器会通过解析导入的 .proto 文件来解析所有类型名。各语言的代码生成器都知道如何在该语言中引用每个类型，即使它们的作用域规则不同。

## 定义服务

如果你想把消息类型用于 RPC（远程过程调用）系统，可以在 .proto 文件中定义 RPC 服务接口，protocol buffer 编译器会为你所选的语言生成服务接口代码和桩代码。例如，如果你想定义一个 RPC 服务方法，它接收 SearchRequest 并返回 SearchResponse，可以在 .proto 文件中这样写：

```protobuf
service SearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}
```

与 protocol buffers 配合使用时，最直接的 RPC 系统是 gRPC。它是 Google 开发的一个语言无关、平台无关的开源 RPC 系统。gRPC 与 protocol buffers 配合得尤其好，并且允许你通过专用的 protocol buffer 编译器插件，直接从 .proto 文件中生成相应的 RPC 代码。

如果你不想使用 gRPC，也可以将 protocol buffers 与你自己的 RPC 实现结合使用。更多内容请参见 Proto2 语言指南。

此外，还有许多正在进行中的第三方项目，致力于为 Protocol Buffers 开发 RPC 实现。相关项目链接列表请参见 third-party add-ons wiki 页面。

## JSON 映射

对于两个都使用 protobuf 的系统之间的通信，标准的 protobuf 二进制线上格式是首选序列化格式。若要与使用 JSON 而非 protobuf 线上格式的系统通信，Protobuf 也支持一种规范化的 JSON 编码。

## 选项

.proto 文件中的单个声明都可以带上一组 option。option 不会改变声明本身的整体语义，但可能会影响该声明在特定上下文中的处理方式。所有可用 option 的完整列表定义在 /google/protobuf/descriptor.proto 中。

有些 option 是文件级的，也就是说它们应该写在顶层作用域，而不是写在 message、enum 或 service 定义内部。有些 option 是消息级的，应写在消息定义内部。还有些 option 是字段级的，应写在字段定义内部。option 也可以写在 enum 类型、enum 值、oneof 字段、service 类型和 service 方法上；不过目前这些位置还没有真正实用的内建 option。

下面列出几个最常用的 option：

- java_package（文件级 option）：用于指定生成的 Java/Kotlin 类所使用的包名。如果 .proto 文件中没有显式给出 java_package，那么默认会使用 proto package（也就是 .proto 文件中通过 package 关键字指定的值）。不过，proto package 往往并不适合作为 Java package，因为 proto package 通常不会使用反向域名开头。如果你不生成 Java 或 Kotlin 代码，此选项不会产生任何效果。

```protobuf
option java_package = "com.example.foo";
```

- java_outer_classname（文件级 option）：用于指定你希望生成的 Java 外层包装类名称（同时也决定文件名）。如果 .proto 文件中没有显式指定 java_outer_classname，那么类名会通过把 .proto 文件名转换成驼峰形式得到（例如 foo_bar.proto 会变成 FooBar.java）。如果 java_multiple_files 选项被禁用，那么该 .proto 文件生成的其他类、枚举等都会作为这个外层包装类的嵌套类或嵌套枚举存在。如果不生成 Java 代码，此选项不会产生任何效果。

```protobuf
option java_outer_classname = "Ponycopter";
```

- java_multiple_files（文件级 option）：如果为 false，则只会为这个 .proto 文件生成一个 .java 文件，而且所有顶层消息、服务和枚举生成的 Java 类等都将嵌套在一个外层类中（参见 java_outer_classname）。如果为 true，则会为顶层消息、服务和枚举分别生成独立的 .java 文件，并且该 .proto 文件生成的外层包装类中不会再包含这些嵌套类或枚举。这是一个布尔选项，默认值为 false。如果不生成 Java 代码，此选项不会产生任何效果。

```protobuf
option java_multiple_files = true;
```

- optimize_for（文件级 option）：可以设置为 SPEED、CODE_SIZE 或 LITE_RUNTIME。它会通过以下方式影响 C++ 和 Java 代码生成器（以及可能的第三方生成器）：

  - SPEED（默认）：protocol buffer 编译器会生成用于序列化、解析以及执行其他常见操作的消息代码。这些代码经过高度优化。
  - CODE_SIZE：protocol buffer 编译器会生成最小化的类，并依赖共享的、基于反射的代码来实现序列化、解析以及各种其他操作。因此生成代码的体积会比 SPEED 模式小得多，但操作速度也更慢。尽管如此，这些类对外提供的公共 API 与 SPEED 模式下完全一致。此模式特别适合包含大量 .proto 文件、但并不要求所有消息类型都极致高性能的应用。
  - LITE_RUNTIME：protocol buffer 编译器会生成只依赖“lite”运行时库的类，也就是依赖 libprotobuf-lite 而不是 libprotobuf。lite 运行时比完整库小得多（大约小一个数量级），但也省略了描述符、反射等某些特性。这对运行在手机等受限平台上的应用尤其有用。编译器仍会像 SPEED 模式一样，为所有方法生成高性能实现。生成的类在各语言中只会实现 MessageLite 接口，也就是完整 Message 接口的一个子集。

```protobuf
option optimize_for = CODE_SIZE;
```

- cc_generic_services、java_generic_services、py_generic_services（文件级 option）：通用服务已经被弃用。它们分别控制 protocol buffer 编译器是否应根据 service 定义，为 C++、Java 和 Python 生成抽象服务代码。出于历史原因，这些选项默认值为 true。不过，自 2.3.0 版本（2010 年 1 月）起，更推荐的做法是让 RPC 实现提供各自的代码生成插件，生成更贴合具体系统的代码，而不是依赖这种“抽象”服务。

```protobuf
// 该文件依赖插件来生成服务代码。
option cc_generic_services = false;
option java_generic_services = false;
option py_generic_services = false;
```

- cc_enable_arenas（文件级 option）：为 C++ 生成代码启用 arena 分配。
- objc_class_prefix（文件级 option）：设置 Objective-C 生成类和枚举统一追加的类名前缀。这个选项没有默认值。你应当使用 3 到 5 个大写字母组成的前缀，这是 Apple 推荐的做法。注意，所有 2 个字母的前缀都被 Apple 保留。
- packed（字段级 option）：对于基本数值类型的 repeated 字段，默认值为 true，这会启用更紧凑的编码方式。如果要使用非 packed 的线上格式，可以将其设为 false。这主要用于兼容 2.3.0 之前版本的解析器（现在很少需要），如下所示：

```protobuf
repeated int32 samples = 4 [packed = false];
```

- deprecated（字段级 option）：如果设为 true，表示该字段已废弃，不应再被新代码使用。在大多数语言中，这个选项本身没有实际运行时效果。在 Java 中，它会变成 @Deprecated 注解。在 C++ 中，clang-tidy 会在使用已废弃字段时发出警告。未来，其他语言的代码生成器也可能在该字段的访问器上生成废弃注解，从而在编译尝试使用该字段的代码时发出警告。如果该字段已经没有任何人在使用，并且你希望阻止新用户继续使用它，可以考虑直接用 reserved 语句替换掉该字段声明。

```protobuf
int32 old_field = 6 [deprecated = true];
```

### 枚举值选项

枚举值支持 option。你可以使用 deprecated 选项来表明某个值不应再被使用。你也可以通过扩展来创建自定义 option。

下面的例子展示了添加这些 option 的语法：

```protobuf
import "google/protobuf/descriptor.proto";

extend google.protobuf.EnumValueOptions {
  optional string string_name = 123456789;
}

enum Data {
  DATA_UNSPECIFIED = 0;
  DATA_SEARCH = 1 [deprecated = true];
  DATA_DISPLAY = 2 [
    (string_name) = "display_value"
  ];
}
```

读取 string_name 选项的 C++ 代码可能大致如下：

```cpp
const absl::string_view foo = proto2::GetEnumDescriptor<Data>()
    ->FindValueByName("DATA_DISPLAY")->options().GetExtension(string_name);
```

关于如何把自定义 option 应用于枚举值和字段，请参见“自定义选项”。

### 自定义选项

Protocol Buffers 也允许你定义并使用自己的 option。请注意，这是一个高级特性，大多数人并不需要。如果你确实认为自己需要创建自定义 option，请参见 Proto2 语言指南中的相关说明。还要注意，在 proto3 中，创建自定义 option 要使用扩展，而扩展只被允许用于自定义 option。

### 选项保留策略

option 具有 retention（保留策略）这一概念，用于控制某个 option 是否会保留在生成代码中。默认情况下，option 具有运行时保留策略，这意味着它们会保留在生成代码中，因此可以在运行时通过生成的描述符池看到。不过，你也可以设置 retention = RETENTION_SOURCE，表示该 option（或 option 内部的字段）不应在运行时保留。这种方式称为源码级保留。

选项保留策略是一个高级特性，大多数用户并不需要关心它。但如果你希望使用某些 option，同时又不想为它们在二进制中保留额外代码体积，这个特性会很有用。具有源码级保留策略的 option 仍然对 protoc 和 protoc 插件可见，因此代码生成器仍可利用它们来定制自己的行为。

可以直接在 option 上设置保留策略，例如：

```protobuf
extend google.protobuf.FileOptions {
  optional int32 source_retention_option = 1234
      [retention = RETENTION_SOURCE];
}
```

也可以把它设置在普通字段上，此时只有当该字段出现在某个 option 内部时，这个设置才会生效：

```protobuf
message OptionsMessage {
  int32 source_retention_field = 1 [retention = RETENTION_SOURCE];
}
```

如果你愿意，也可以设置 retention = RETENTION_RUNTIME，不过这不会产生任何效果，因为它本来就是默认行为。当某个消息字段被标记为 RETENTION_SOURCE 时，它的整个内容都会被丢弃；其内部字段无法通过试图设置 RETENTION_RUNTIME 来覆盖这一行为。

#### 注意

截至 Protocol Buffers 22.0，对 option retention 的支持仍在持续推进中，目前只有 C++ 和 Java 已支持。Go 从 1.29.0 开始支持。Python 端的支持已经完成，但尚未进入正式发布版本。

### 选项目标

字段可以带有 targets 选项，用于控制该字段在作为 option 使用时可应用到哪些实体类型上。例如，如果某个字段带有 targets = TARGET_TYPE_MESSAGE，那么该字段就不能在 enum 的自定义 option 中被设置（或者更一般地，不能用于任何非 message 实体）。protoc 会强制执行这一限制，如果违反目标约束就会报错。

乍一看，这个特性似乎有些多余，因为每个自定义 option 本身就是某种特定实体的 options 消息的一个扩展，这本来就已经把它限制在该实体上了。不过，当你有一个可应用于多种实体类型的共享 options 消息，并且希望控制其中各字段在不同实体上的使用范围时，option targets 就很有用了。例如：

```protobuf
message MyOptions {
  string file_only_option = 1 [targets = TARGET_TYPE_FILE];
  int32 message_and_enum_option = 2 [targets = TARGET_TYPE_MESSAGE,
                                     targets = TARGET_TYPE_ENUM];
}

extend google.protobuf.FileOptions {
  optional MyOptions file_options = 50000;
}

extend google.protobuf.MessageOptions {
  optional MyOptions message_options = 50000;
}

extend google.protobuf.EnumOptions {
  optional MyOptions enum_options = 50000;
}

// 正确：该字段允许用于 file options
option (file_options).file_only_option = "abc";

message MyMessage {
  // 正确：该字段允许用于 message 和 enum options
  option (message_options).message_and_enum_option = 42;
}

enum MyEnum {
  MY_ENUM_UNSPECIFIED = 0;
  // 错误：file_only_option 不能设置在 enum 上。
  option (enum_options).file_only_option = "xyz";
}
```

## 生成代码

如果你想生成 Java、Kotlin、Python、C++、Go、Ruby、Objective-C 或 C# 代码，以便处理 .proto 文件中定义的消息类型，就需要对该 .proto 运行 protocol buffer 编译器 protoc。如果你尚未安装编译器，请先下载安装包并按照 README 中的说明进行安装。对于 Go，你还需要为编译器安装一个专门的代码生成插件，相关内容及安装说明可以在 GitHub 上的 golang/protobuf 仓库中找到。

protobuf 编译器的调用形式如下：

```bash
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto
```

IMPORT_PATH 指定了解析 import 指令时查找 .proto 文件的目录。如果省略，则使用当前目录。你可以多次传入 --proto_path 以指定多个导入目录。-I=IMPORT_PATH 是 --proto_path 的简写形式。

注意：相对于各自 proto_path 的文件路径，在同一个二进制中必须全局唯一。例如，如果你有 proto/lib1/data.proto 和 proto/lib2/data.proto，那么不能同时使用 -I=proto/lib1 和 -I=proto/lib2，因为 import "data.proto" 时会无法判断应引用哪一个文件。正确做法是使用 -Iproto/，这样它们的全局名称就分别是 lib1/data.proto 和 lib2/data.proto。

如果你正在发布一个库，并且其他用户可能会直接使用你的消息，那么你应当在它们预期被使用的路径中包含一个唯一的库名，以避免文件名冲突。如果一个项目中包含多个目录，最佳实践通常是把一个最高层目录设置为 -I。

你可以提供一个或多个输出指令：

- --cpp_out：在 DST_DIR 中生成 C++ 代码。详见 C++ generated code reference。
- --java_out：在 DST_DIR 中生成 Java 代码。详见 Java generated code reference。
- --kotlin_out：在 DST_DIR 中生成额外的 Kotlin 代码。详见 Kotlin generated code reference。
- --python_out：在 DST_DIR 中生成 Python 代码。详见 Python generated code reference。
- --go_out：在 DST_DIR 中生成 Go 代码。详见 Go generated code reference。
- --ruby_out：在 DST_DIR 中生成 Ruby 代码。详见 Ruby generated code reference。
- --objc_out：在 DST_DIR 中生成 Objective-C 代码。详见 Objective-C generated code reference。
- --csharp_out：在 DST_DIR 中生成 C# 代码。详见 C# generated code reference。
- --php_out：在 DST_DIR 中生成 PHP 代码。详见 PHP generated code reference。

作为额外便利，如果 DST_DIR 以 .zip 或 .jar 结尾，编译器会把输出写入指定名称的单个 ZIP 格式归档文件中。以 .jar 结尾的输出还会额外包含符合 Java JAR 规范要求的 manifest 文件。请注意，如果该归档文件已经存在，它会被直接覆盖。

你必须提供一个或多个 .proto 文件作为输入。可以一次指定多个 .proto 文件。虽然文件名是相对于当前目录给出的，但每个文件都必须位于某个 IMPORT_PATH 之下，这样编译器才能确定它的规范名称。

## 文件位置

尽量不要把 .proto 文件与其他语言的源文件放在同一个目录中。可以考虑在项目根包下面创建一个 proto 子包，专门存放 .proto 文件。

### 位置应当与语言无关

在使用 Java 代码时，把相关 .proto 文件放在 Java 源文件所在目录中会比较方便。但如果以后有任何非 Java 代码也要使用同样的 proto，这个路径前缀就会显得不合适。因此一般来说，应当把 proto 放在与语言无关的相关目录中，例如 //myteam/mypackage。

这一规则的例外，是你非常明确这些 proto 只会在 Java 环境中使用，例如仅用于测试时。

## 支持的平台

关于以下内容的信息：

- 受支持的操作系统、编译器、构建系统以及 C++ 版本，请参见 Foundational C++ Support Policy。
- 受支持的 PHP 版本，请参见 Supported PHP versions。
