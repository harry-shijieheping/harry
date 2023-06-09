# harry
# 1.1 起源与发展

Go 语言起源 2007 年，并于 2009 年正式对外发布。它从 2009 年 9 月 21 日开始作为谷歌公司 20% 兼职项目，即相关员工利用 20% 的空余时间来参与 Go 语言的研发工作。该项目的三位领导者均是著名的 IT 工程师：Robert Griesemer，参与开发 Java HotSpot 虚拟机；Rob Pike，Go 语言项目总负责人，贝尔实验室 Unix 团队成员，参与的项目包括 Plan 9，Inferno 操作系统和 Limbo 编程语言；Ken Thompson，贝尔实验室 Unix 团队成员，C 语言、Unix 和 Plan 9 的创始人之一，与 Rob Pike 共同开发了 UTF-8 字符集规范。自 2008 年 1 月起，Ken Thompson 就开始研发一款以 C 语言为目标结果的编译器来拓展 Go 语言的设计思想。

**这是一个由计算机领域 “发明之父” 所组成的黄金团队，他们对系统编程语言，操作系统和并行都有着非常深刻的见解**

![](images/1.1.designers_of_Go.jpg?raw=true)

<center>图 1.1 Go 语言设计者：Griesemer、Thompson 和 Pike</center>

在 2008 年年中，Go 语言的设计工作接近尾声，一些员工开始以全职工作状态投入到这个项目的编译器和运行实现上。Ian Lance Taylor 也加入到了开发团队中，并于 2008 年 5 月创建了一个 gcc 前端。

Russ Cox 加入开发团队后着手语言和类库方面的开发，也就是 Go 语言的标准包。在 2009 年 10 月 30 日，Rob Pike 以 Google Techtalk 的形式第一次向人们宣告了 Go 语言的存在。

直到 2009 年 11 月 10 日，开发团队将 Go 语言项目以 BSD-style 授权（完全开源）正式公布了 Linux 和 Mac OS X 平台上的版本。Hector Chu 于同年 11 月 22 日公布了 Windows 版本。

作为一个开源项目，Go 语言借助开源社区的有生力量达到快速地发展，并吸引更多的开发者来使用并改善它。自该开源项目发布以来，超过 200 名非谷歌员工的贡献者对 Go 语言核心部分提交了超过 1000 个修改建议。在过去的 18 个月里，又有 150 开发者贡献了新的核心代码。这俨然形成了世界上最大的开源团队，并使该项目跻身 [Ohloh](http://www.ohloh.net) 前 2% 的行列。大约在 2011 年 4 月 10 日，谷歌开始抽调员工进入全职开发 Go 语言项目。开源化的语言显然能够让更多的开发者参与其中并加速它的发展速度。Andrew Gerrand 在 2010 年加入到开发团队中成为共同开发者与支持者。

在 Go 语言在 2010 年 1 月 8 日被 [Tiobe](http://www.tiobe.com)（闻名于它的编程语言流行程度排名）宣布为 “2009 年年度语言” 后，引起各界很大的反响。目前 Go 语言在这项排名中的最高记录是在 2017 年 1 月创下的第13名，流行程度 2.325%。

### 时间轴：

- 2007 年 9 月 21 日：雏形设计
- 2009 年 11 月 10日：首次公开发布
- 2010 年 1 月 8 日：当选 2009 年年度语言
- 2010 年 5 月：谷歌投入使用
- 2011 年 5 月 5 日：Google App Engine 支持 Go 语言

从 2010 年 5 月起，谷歌开始将 Go 语言投入到后端基础设施的实际开发中，例如开发用于管理后端复杂环境的项目。有句话叫 “吃你自己的狗食”，这也体现了谷歌确实想要投资这门语言，并认为它是有生产价值的。

Go 语言的官方网站是 [golang.org](http://golang.org)，这个站点采用 Python 作为前端，并且使用 Go 语言自带的工具 godoc 运行在 Google App Engine 上来作为 Web 服务器提供文本内容。在官网的首页有一个功能叫做 Go Playground，是一个 Go 代码的简单编辑器的沙盒，它可以在没有安装 Go 语言的情况下在你的浏览器中编译并运行 Go，它提供了一些示例，其中包括国际惯例 “Hello, World!”。

更多的信息详见 [github.com/golang/go](https://github.com/golang/go)，Go 项目 Bug 追踪和功能预期详见 [github.com/golang/go/issues](https://github.com/golang/go/issues)。

Go 通过以下的 Logo 来展示它的速度，并以囊地鼠 (Gopher) 作为它的吉祥物。

![](images/1.2.Go_logo.jpg?raw=true)

<center>图1.2 Go 语言 Logo</center>

谷歌邮件列表 [golang-nuts](http://groups.google.com/group/golang-nuts/) 非常活跃，每天的讨论和问题解答数以百计。

关于 Go 语言在 Google App Engine 的应用，这里有一个单独的邮件列表 [google-appengine-go](https://groups.google.com/forum/#!forum/google-appengine-go)，不过 2 个邮件列表的讨论内容并不是分得很清楚，都会涉及到相关的话题。[go-lang.cat-v.org/](http://go-lang.cat-v.org/) 是 Go 语言开发社区的资源站，[irc.freenode.net](http://irc.freenode.net) 的 #go-nuts 是官方的 Go IRC 频道。

[@golang](https://twitter.com/golang) 是 Go 语言在 Twitter 的官方帐号，大家一般使用 #golang 作为话题标签。

这里还有一个在 Linked-in 的小组：[www.linkedin.com/groups?gid=2524765&trk=myg_ugrp_ovr](http://www.linkedin.com/groups?gid=2524765&trk=myg_ugrp_ovr)。

Go 编程语言的维基百科：[en.wikipedia.org/wiki/Go_(programming_language)](http://en.wikipedia.org/wiki/Go_\(programming_language\))

Go 语言相关资源的搜索引擎页面：[gowalker.org](https://gowalker.org)

Go 语言还有一个运行在 Google App Engine 上的 [Go Tour](http://tour.golang.org/)，你也可以通过执行命令 `go install go-tour.googlecode.com/hg/gotour` 安装到你的本地机器上。对于中文读者，可以访问该指南的 [中文版本](https://tour.go-zh.org/welcome/1)，或通过命令 `go install https://bitbucket.org/mikespook/go-tour-zh/gotour` 进行安装。

## 链接

- [目录](directory.md)
- 上一部分：[前言](preface.md)
- 下一节: [语言的主要特性与发展的环境和影响因素](01.2.md)
- # 1.2 语言的主要特性与发展的环境和影响因素

## 1.2.1 影响 Go 语言发展的早期编程语言

正如 “21 世纪的 C 语言” 这句话所说，Go 语言并不是凭空而造的，而是和 C++、Java 和 C# 一样属于 C 系。不仅如此，设计者们还汲取了其它编程语言的精粹部分融入到 Go 语言当中。

在声明和包的设计方面，Go 语言受到 Pascal、Modula 和 Oberon 系语言的影响；在并发原理的设计上，Go 语言从同样受到 Tony Hoare 的 CSP（通信序列进程 *Communicating Sequential Processes*）理论影响的 Limbo 和 Newsqueak 的实践中借鉴了一些经验，并使用了和 Erlang 类似的机制。

这是一门完全开源的编程语言，因为它使用 BSD 授权许可，所以任何人都可以进行商业软件的开发而不需要支付任何费用。

尽管为了能够让目前主流的开发者们能够对 Go 语言中的类 C 语言的语法感到非常亲切而易于转型，但是它在极大程度上简化了这些语法，使得它们比 C/C++ 的语法更加简洁和干净。同时，Go 语言也拥有一些动态语言的特性，这使得使用 Python 和 Ruby 的开发者们在使用 Go 语言的时候感觉非常容易上手。

下图展示了一些其它编程语言对 Go 语言的影响：

![](images/1.3.influences_on_go.jpg?raw=true)

图 1.3 其它编程语言对 Go 语言的影响

## 1.2.2 为什么要创造一门编程语言

- C/C++ 的发展速度无法跟上计算机发展的脚步，十多年来也没有出现一门与时代相符的主流系统编程语言，因此人们需要一门新的系统编程语言来弥补这个空缺，尤其是在计算机信息时代。
- 相比计算机性能的提升，软件开发领域不被认为发展得足够快或者比硬件发展得更加成功（有许多项目均以失败告终），同时应用程序的体积始终在不断地扩大，这就迫切地需要一门具备更高层次概念的低级语言来突破现状。
- 在 Go 语言出现之前，开发者们总是面临非常艰难的抉择，究竟是使用执行速度快但是编译速度并不理想的语言（如：C++），还是使用编译速度较快但执行效率不佳的语言（如：.NET、Java），或者说开发难度较低但执行速度一般的动态语言呢？显然，Go 语言在这 3 个条件之间做到了最佳的平衡：快速编译，高效执行，易于开发。

## 1.2.3 Go 语言的发展目标

Go 语言的主要目标是将静态语言的安全性和高效性与动态语言的易开发性进行有机结合，达到完美平衡，从而使编程变得更加有乐趣，而不是在艰难抉择中痛苦前行。

因此，Go 语言是一门类型安全和内存安全的编程语言。虽然 Go 语言中仍有指针的存在，但并不允许进行指针运算。

Go 语言的另一个目标是对于网络通信、并发和并行编程的极佳支持，从而更好地利用大量的分布式和多核的计算机，这一点对于谷歌内部的使用来说就非常重要了。设计者通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信。他们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。

这个特性显然是 Go 语言最强有力的部分，不仅支持了日益重要的多核与多处理器计算机，也弥补了现存编程语言在这方面所存在的不足。

Go 语言中另一个非常重要的特性就是它的构建速度（编译和链接到机器代码的速度），一般情况下构建一个程序的时间只需要数百毫秒到几秒。作为大量使用 C++ 来构建基础设施的谷歌来说，无疑从根本上摆脱了 C++ 在构建速度上非常不理想的噩梦。这不仅极大地提升了开发者的生产力，同时也使得软件开发过程中的代码测试环节更加紧凑，而不必浪费大量的时间在等待程序的构建上。

依赖管理是现今软件开发的一个重要组成部分，但是 C 语言中“头文件”的概念却导致越来越多因为依赖关系而使得构建一个大型的项目需要长达几个小时的时间。人们越来越需要一门具有严格的、简洁的依赖关系分析系统从而能够快速编译的编程语言。这正是 Go 语言采用包模型的根本原因，这个模型通过严格的依赖关系检查机制来加快程序构建的速度，提供了非常好的可量测性。

整个 Go 语言标准库的编译时间一般都在 20 秒以内，其它的常规项目也只需要半秒钟的时间来完成编译工作。这种闪电般的编译速度甚至比编译 C 语言或者 Fortran 更加快，使得编译这一环节不再成为在软件开发中困扰开发人员的问题。在这之前，动态语言将快速编译作为自身的一大亮点，像 C++ 那样的静态语言一般都有非常漫长的编译和链接工作。而同样作为静态语言的 Go 语言，通过自身优良的构建机制，成功地去除了这个弊端，使得程序的构建过程变得微不足道，拥有了像脚本语言和动态语言那样的高效开发的能力。

另外，Go 语言在执行速度方面也可以与 C/C++ 相提并论。

由于内存问题（通常称为内存泄漏）长期以来一直伴随着 C++ 的开发者们，Go 语言的设计者们认为内存管理不应该是开发人员所需要考虑的问题。因此尽管 Go 语言像其它静态语言一样执行本地代码，但它依旧运行在某种意义上的虚拟机，以此来实现高效快速的垃圾回收（使用了一个简单的标记-清除算法）。

尽管垃圾回收并不容易实现，但考虑这将是未来并发应用程序发展的一个重要组成部分，Go 语言的设计者们还是完成了这项艰难的任务。

Go 语言还能够在运行时进行反射相关的操作。

使用 `go install` 能够很轻松地对第三方包进行部署。

此外，Go 语言还支持调用由 C 语言编写的海量库文件（[第 3.9 节](03.9.md)），从而能够将过去开发的软件进行快速迁移。

## 1.2.4 指导设计原则

Go语言通过减少关键字的数量（25 个）来简化编码过程中的混乱和复杂度。干净、整齐和简洁的语法也能够提高程序的编译速度，因为这些关键字在编译过程中少到甚至不需要符号表来协助解析。

这些方面的工作都是为了减少编码的工作量，甚至可以与 Java 的简化程度相比较。

Go 语言有一种极简抽象艺术家的感觉，因为它只提供了一到两种方法来解决某个问题，这使得开发者们的代码都非常容易阅读和理解。众所周知，代码的可读性是软件工程里最重要的一部分（ **译者注：代码是写给人看的，不是写给机器看的** ）。

这些设计理念没有建立其它概念之上，所以并不会因为牵扯到一些概念而将某个概念复杂化，他们之间是相互独立的。

Go 语言有一套完整的编码规范，你可以在 [Go 语言编码规范](http://golang.org/doc/go_spec.html) 页面进行查看。

它不像 Ruby 那样通过实现过程来定义编码规范。作为一门具有明确编码规范的语言，它要求可以采用不同的编译器如 gc 和 gccgo（[第 2.1 节](02.1.md)）进行编译工作，这对语言本身拥有更好的编码规范起到很大帮助。

[LALR](http://en.wikipedia.org/wiki/LALR_parser) 是 Go 语言的语法标准，你也可以在 [`src/cmd/internal/gc/go.y`](https://github.com/golang/go/blob/master/src%2Fcmd%2Finternal%2Fgc%2Fgo.y) 中查看到，这种语法标准在编译时不需要符号表来协助解析。

## 1.2.5 语言的特性

Go 语言从本质上（程序和结构方面）来实现并发编程。

因为 Go 语言没有类和继承的概念，所以它和 Java 或 C++ 看起来并不相同。但是它通过接口 (interface) 的概念来实现多态性。Go 语言有一个清晰易懂的轻量级类型系统，在类型之间也没有层级之说。因此可以说这是一门混合型的语言。

在传统的面向对象语言中，使用面向对象编程技术显得非常臃肿，它们总是通过复杂的模式来构建庞大的类型层级，这违背了编程语言应该提升生产力的宗旨。

函数是 Go 语言中的基本构件，它们的使用方法非常灵活。在[第六章](06.0.md)，我们会看到 Go 语言在函数式编程方面的基本概念。


Go 语言使用静态类型，所以它是类型安全的一门语言，加上通过构建到本地代码，程序的执行速度也非常快。

作为强类型语言，隐式的类型转换是不被允许的，记住一条原则：让所有的东西都是显式的。

Go 语言其实也有一些动态语言的特性（通过关键字 `var`），所以它对那些逃离 Java 和 .Net 世界而使用 Python、Ruby、PHP 和 JavaScript 的开发者们也具有很大的吸引力。

Go 语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发运行 Windows 下运行的应用程序。这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，就连它的源码文件格式都是使用的 UTF-8 编码。Go 语言做到了真正的国际化！

## 1.2.6 语言的用途

Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。

Go 语言一个非常好的目标就是实现所谓的复杂事件处理（[CEP](http://en.wikipedia.org/wiki/Complex_event_processing)），这项技术要求海量并行支持，高度的抽象化和高性能。当我们进入到物联网时代，CEP 必然会成为人们关注的焦点。

但是 Go 语言同时也是一门可以用于实现一般目标的语言，例如对于文本的处理，前端展现，甚至像使用脚本一样使用它。

值得注意的是，因为垃圾回收和自动内存分配的原因，Go 语言不适合用来开发对实时性要求很高的软件。

越来越多的谷歌内部的大型分布式应用程序都开始使用 Go 语言来开发，例如谷歌地球的一部分代码就是由 Go 语言完成的。

如果你想知道一些其它组织使用Go语言开发的实际应用项目，你可以到 [使用 Go 的组织](http://go-lang.cat-v.org/organizations-using-go) 页面进行查看。出于隐私保护的考虑，许多公司的项目都没有展示在这个页面。我们将会在[第 21 章](21.0.md) 讨论到一个使用 Go 语言开发的大型存储区域网络 (SAN) 案例。


在 Chrome 浏览器中内置了一款 Go 语言的编译器用于本地客户端 (NaCl)，这很可能会被用于在 Chrome OS 中执行 Go 语言开发的应用程序。

Go 语言可以在 Intel 或 ARM 处理器上运行，因此它也可以在安卓系统下运行，例如 Nexus 系列的产品。

在 Google App Engine 中使用 Go 语言：2011 年 5 月 5 日，官方发布了用于开发运行在 Google App Engine 上的 Web 应用的 Go SDK，在此之前，开发者们只能选择使用 Python 或者 Java。这主要是 David Symonds 和 Nigel Tao 努力的成果。目前最新的稳定版是基于 Go 1.4 的 SDK 1.9.18，于 2015 年 2 月 18 日发布。当前 Go 语言的稳定版本是 Go 1.4.2。

## 1.2.7 关于特性缺失

许多能够在大多数面向对象语言中使用的特性 Go 语言都没有支持，但其中的一部分可能会在未来被支持。

- 为了简化设计，不支持函数重载和操作符重载
- 为了避免在 C/C++ 开发中的一些 Bug 和混乱，不支持隐式转换
- Go 语言通过另一种途径实现面向对象设计（第 [10](10.0.md)-[11](11.0.md) 章）来放弃类和类型的继承
- 尽管在接口的使用方面（[第 11 章](11.0.md)）可以实现类似变体类型的功能，但本身不支持变体类型
- 不支持动态加载代码
- 不支持动态链接库
- 不支持泛型
- 通过 `recover()` 和 `panic()` 来替代异常机制（第 [13.2](13.2.md)-[13.3](13.3.md) 节）
- 不支持静态变量

关于 Go 语言开发团队对于这些方面的讨论，你可以通过 [Go 常见问题](http://golang.org/doc/go_faq.html) 页面查看。

## 1.2.8 使用 Go 语言编程

如果你有其它语言的编程经历（面向对象编程语言，如：Java、C#、Object-C、Python、Ruby），在你进入到 Go 语言的世界之后，你将会像迷恋你的 X 语言一样无法自拔。Go 语言使用了与其它语言不同的设计模式，所以当你尝试将你的X语言的代码迁移到 Go 语言时，你将会非常失望，所以你需要从头开始，用 Go 的理念来思考。

如果你在至高点使用 Go 的理念来重新审视和分析一个问题，你通常会找到一个适用于 Go 语言的优雅的解决方案。

## 1.2.9 小结

这里列举一些 Go 语言的必杀技：

- 简化问题，易于学习
- 内存管理，简洁语法，易于使用
- 快速编译，高效开发
- 高效执行
- 并发支持，轻松驾驭
- 静态类型
- 标准类库，规范统一
- 易于部署
- 文档全面
- 免费开源

## 链接

- [目录](directory.md)
- 上一节：[起源与发展](01.1.md)
- 下一章：[安装与运行环境](02.1.md)
- # 2.1 平台与架构

Go 语言开发团队开发了适用于以下操作系统的编译器：

- Linux
- FreeBSD
- Mac OS X（也称为 Darwin）

目前有2个版本的编译器：Go 原生编译器 gc 和非原生编译器 gccgo，这两款编译器都是在类 Unix 系统下工作 。其中，gc 版本的编译器已经被移植到 Windows 平台上，并集成在主要发行版中，你也可以通过安装 MinGW 从而在 Windows 平台下使用 gcc 编译器。这两个编译器都是以单通道的形式工作。

你可以获取以下平台上的 Go 1.4 源码和二进制文件：

- Linux 2.6+：amd64、386 和 arm 架构
- Mac OS X (Snow Leopard + Lion) ：amd64 和 386 架构
- Windows 2000+：amd64 和 386 架构

对于非常底层的纯 Go 语言代码或者包而言，在各个操作系统平台上的可移植性是非常强的，只需要将源码拷贝到相应平台上进行编译即可，或者可以使用交叉编译来构建目标平台的应用程序（[第 2.2 节](02.2.md)）。但如果你打算使用 cgo 或者类似文件监控系统的软件，就需要根据实际情况进行相应地修改了。

1. Go 原生编译器 gc：

   主要基于 Ken Thompson 先前在 Plan 9 操作系统上使用的 C 工具链。

   Go 语言的编译器和链接器都是使用 C 语言编写并产生本地代码，Go 不存在自我引导之类的功能。因此如果使用一个有不同指令集的编译器来构建 Go 程序，就需要针对操作系统和处理器架构（32 位操作系统或 64 位操作系统）进行区别对待。（ **译者注：Go从1.5版本开始已经实现自举。Go语言的编译器和链接器都是Go语言编写的**）

   这款编译器使用非分代、无压缩和并行的方式进行编译，它的编译速度要比 gccgo 更快，产生更好的本地代码，但编译后的程序不能够使用 gcc 进行链接。

   编译器目前支持以下基于 Intel 或 AMD 处理器架构的程序构建。

   ![](images/2.1.gc.jpg?raw=true)

   <center>图2.1 gc 编译器支持的处理器架构</center>

   当你第一次看到这套命名系统的时候你会觉得很奇葩，不过这些命名都是来自于 Plan 9 项目。

   	g = 编译器：将源代码编译为项目代码（程序文本）
   	l = 链接器：将项目代码链接到可执行的二进制文件（机器代码）

   （相关的 C 编译器名称为 6c、8c 和 5c，相关的汇编器名称为 6a、8a 和 5a）

   **标记 (Flags)** 是指可以通过命令行设置可选参数来影响编译器或链接器的构建过程或得到一个特殊的目标结果。

   可用的编译器标记如下：

   	flags:
   	-I 针对包的目录搜索
   	-d 打印声明信息
   	-e 不限制错误打印的个数
   	-f 打印栈结构
   	-h 发生错误时进入恐慌（panic）状态
   	-o 指定输出文件名 // 详见第3.4节
   	-S 打印产生的汇编代码
   	-V 打印编译器版本 // 详见第2.3节
   	-u 禁止使用 unsafe 包中的代码
   	-w 打印归类后的语法解析树
   	-x 打印 lex tokens

   从 Go 1.0.3 版本开始，不再使用 `8g`，`8l` 之类的指令进行程序的构建，取而代之的是统一的 `go build` 和 `go install` 等命令，而这些指令会自动调用相关的编译器或链接器。


	如果你想获得更深层次的信息，你可以在目录 [`$GOROOT/src/cmd`](https://github.com/golang/go/tree/master/src/cmd) 下找到编译器和链接器的源代码。Go 语言本身是由 C 语言开发的，而不是 Go 语言（Go 1.5 开始自举）。词法分析程序是 GNU bison，语法分析程序是名为 [`$GOROOT/src/cmd/gc/go.y`](https://github.com/golang/go/blob/master/src%2Fcmd%2Finternal%2Fgc%2Fgo.y) 的 yacc 文件，它会在同一目录输出 `y.tab.{c,h}` 文件。如果你想知道更多有关构建过程的信息，你可以在 [`$GOROOT/src/make.bash`](https://github.com/golang/go/blob/master/src/make.bash) 中找到。
	
	大部分的目录都包含了名为 `doc.go` 的文件，这个文件提供了更多详细的信息。

2. gccgo 编译器：

   一款相对于 gc 而言更加传统的编译器，使用 GCC 作为后端。GCC 是一款非常流行的 GNU 编译器，它能够构建基于众多处理器架构的应用程序。编译速度相对 gc 较慢，但产生的本地代码运行要稍微快一点。它同时也提供一些与 C 语言之间的互操作性。

   从 Go 1 版本开始，gc 和 gccgo 在编译方面都有等价的功能。

3. 文件扩展名与包 (package)：

   Go 语言源文件的扩展名很显然就是 `.go`。

   C 文件使用后缀名 `.c`，汇编文件使用后缀名 `.s`。所有的源代码文件都是通过包 (packages) 来组织。包含可执行代码的包文件在被压缩后使用扩展名 `.a`（AR 文档）。

   Go 语言的标准库（[第 9.1 节](09.1.md)）包文件在被安装后就是使用这种格式的文件。

   **注意** 当你在创建目录时，文件夹名称永远不应该包含空格，而应该使用下划线 "_" 或者其它一般符号代替。

## 链接

- [目录](directory.md)
- 上一章：[语言的主要特性与发展的环境和影响因素](01.2.md)
- 下一节：[Go 环境变量](02.2.md)
- # 2.2 Go 环境变量

Go 开发环境依赖于一些操作系统环境变量，你最好在安装 Go 之前就已经设置好他们。如果你使用的是 Windows 的话，你完全不用进行手动设置，Go 将被默认安装在目录 `c:/go` 下。这里列举几个最为重要的环境变量：

- **$GOROOT** 表示 Go 在你的电脑上的安装位置，它的值一般都是 `$HOME/go`，当然，你也可以安装在别的地方。
- **$GOARCH** 表示目标机器的处理器架构，它的值可以是 386、amd64 或 arm。
- **$GOOS** 表示目标机器的操作系统，它的值可以是 darwin、freebsd、linux 或 windows。
- **$GOBIN** 表示编译器和链接器的安装位置，默认是 `$GOROOT/bin`，如果你使用的是 Go 1.0.3 及以后的版本，一般情况下你可以将它的值设置为空，Go 将会使用前面提到的默认值。

目标机器是指你打算运行你的 Go 应用程序的机器。

Go 编译器支持交叉编译，也就是说你可以在一台机器上构建能够在不同操作系统和处理器架构上运行的应用程序，也就是说编写源代码的机器可以和目标机器有完全不同的特性（操作系统与处理器架构）。

为了区分本地机器和目标机器，你可以使用 `$GOHOSTOS` 和 `$GOHOSTARCH` 设置本地机器的操作系统名称和编译体系结构，这两个变量只有在进行交叉编译的时候才会用到，如果你不进行显示设置，他们的值会和本地机器（`$GOOS` 和 `$GOARCH`）一样。

- **$GOPATH** 默认采用和 `$GOROOT` 一样的值，但从 Go 1.1 版本开始，你必须修改为其它路径。它可以包含多个 Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：`src`、`pkg` 和 `bin`，这三个目录分别用于存放源码文件、包文件和可执行文件。
- **$GOARM** 专门针对基于 arm 架构的处理器，它的值可以是 5 或 6，默认为 6。
- **$GOMAXPROCS** 用于设置应用程序可使用的处理器个数与核数，详见 [第 14.1.3 节](14.1.md)。

在接下来的章节中，我们将会讨论如何在 Linux、Mac OS X 和 Windows 上安装 Go 语言。在 FreeBSD 上的安装和 Linux 非常类似。开发团队正在尝试将 Go 语言移植到其它例如 OpenBSD、DragonFlyBSD、NetBSD、Plan 9、Haiku 和 Solaris 操作系统上，你可以在这个页面找到最近的动态：[Go Porting Efforts](http://go-lang.cat-v.org/os-ports)。

## 链接

- [目录](directory.md)
- 上一节：[平台与架构](02.1.md)
- 下一节：[在 Linux 上安装 Go](02.3.md)
- # 2.3 在 Linux 上安装 Go

如果你能够自己下载并编译 Go 的源代码的话,对你来说是非常有教育意义的，你可以根据这个页面找到安装指南和下载地址：[Download the Go distribution](http://golang.org/doc/install)。

我们接下来也会带你一步步地完成安装过程。

1. 设置 Go 环境变量

   我们在 Linux 系统下一般通过文件 `$HOME/.bashrc` 配置自定义环境变量，根据不同的发行版也可能是文件 `$HOME/.profile`，然后使用 gedit 或 vi 来编辑文件内容。

   	export GOROOT=$HOME/go

   为了确保相关文件在文件系统的任何地方都能被调用，你还需要添加以下内容：

   	export PATH=$PATH:$GOROOT/bin

   在开发 Go 项目时，你还需要一个环境变量来保存你的工作目录。

   	export GOPATH=$HOME/Applications/Go

   `$GOPATH` 可以包含多个工作目录，取决于你的个人情况。如果你设置了多个工作目录，那么当你在之后使用 `go get`（远程包安装命令）时远程包将会被安装在第一个目录下。

   在完成这些设置后，你需要在终端输入指令 `source .bashrc` 以使这些环境变量生效。然后重启终端，输入 `go env` 和 `env` 来检查环境变量是否设置正确。

2. 安装 C 工具

   Go 的工具链是用 C 语言编写的，因此在安装 Go 之前你需要先安装相关的 C 工具。如果你使用的是 Ubuntu 的话，你可以在终端输入以下指令（ **译者注：由于网络环境的特殊性，你可能需要将每个工具分开安装** ）。

   	sudo apt-get install bison ed gawk gcc libc6-dev make

   你可以在其它发行版上使用 RPM 之类的工具。

3. 获取 Go 源代码

   从 [官方页面](https://golang.org/dl/) 或 [国内镜像](http://www.golangtc.com/download) 下载 Go 的源码包到你的计算机上，然后将解压后的目录 `go` 通过命令移动到 `$GOROOT` 所指向的位置。

   	wget https://storage.googleapis.com/golang/go<VERSION>.src.tar.gz
   	tar -zxvf go<VERSION>.src.tar.gz
   	sudo mv go $GOROOT

4. 构建 Go

   在终端使用以下指令来进行编译工作。

   	cd $GOROOT/src
   	./all.bash

   **编译注意事项**

   编译时如果出现如下报错：

   ![](images/2.3.allbasherror.png?raw=true)

   可能是因为 `$GOROOT_BOOTSTRAP` 变量没有设置。这个目录在安装 Go 1.5 版本及之后的版本时需要设置。

   由于在 1.4 版本后，Go 编译器实现了自举，即通过 1.4 版本来编译安装之后版本的编译器。如果不设置该环境变量的话，会产生这样一个错误 `Set $GOROOT_BOOTSTRAP to a working Go tree >= Go 1.4.` 。

   设置 `$GOROOT_BOOTSTRAP` 变量：


       export GOROOT_BOOTSTRAP=$HOME/go1.4

设置完成后，下载 1.4 版本的源码到该目录：

       git clone https://github.com/golang/go.git $HOME/go1.4
       git checkout -b release-branch.go1.4 origin/release-branch.go1.4

进入 1.4 的文件夹后，进行编译：

       cd $HOME/go1.4/src
       ./make.bash

1.4 编译安装好之后，进入 `$GOROOT` 文件夹，真正开始编译安装 Go：

       cd $HOME/go/src
       ./all.bash

在完成编译之后（通常在 1 分钟以内，如果你在 B 型树莓派上编译，一般需要 1 个小时），你会在终端看到如下信息被打印：

![](images/2.3.allbash.png?raw=true)

   <center>图 2.3 完成编译后在终端打印的信息</center>

**注意事项**

在测试 `net/http` 包时有一个测试会尝试连接 `google.com`，你可能会看到如下所示的一个无厘头的错误报告：

   	‘make[2]: Leaving directory `/localusr/go/src/pkg/net’

如果你正在使用一个带有防火墙的机器，我建议你可以在编译过程中暂时关闭防火墙，以避免不必要的错误。

解决这个问题的另一个办法是通过设置环境变量 `$DISABLE_NET_TESTS` 来告诉构建工具忽略 `net/http` 包的相关测试：

   	export DISABLE_NET_TESTS=1

如果你完全不想运行包的测试，你可以直接运行 `./make.bash` 来进行单纯的构建过程。

5. 测试安装

   使用你最喜爱的编辑器来输入以下内容，并保存为文件名 `hello_world1.go`。

   示例 2.1 [hello_world1.go](examples/chapter_2/hello_world1.go)

   ```go
   package main
   
   func main() {
   	println("Hello", "world")
   }
   ```

   切换相关目录到下，然后执行指令 `go run hello_world1.go`，将会打印信息：`Hello, world`。

6. 验证安装版本

   你可以通过在终端输入指令 `go version` 来打印 Go 的版本信息。

   如果你想要通过 Go 代码在运行时检测版本，可以通过以下例子实现。

   示例 2.2 [version.go](examples/chapter_2/version.go)

   ```go
   package main
   
   import (
   	"fmt"
   	"runtime"
   )
   
   func main() {
   	fmt.Printf("%s", runtime.Version())
   }
   ```

   这段代码将会输出 `go1.4.2` 或类似字符串。

7. 更新版本

   你可以在 [发布历史](http://golang.org/doc/devel/release.html) 页面查看到最新的稳定版。

   当前最新的稳定版 Go 1 系列于 2012 年 3 月 28 日发布。

   Go 的源代码有以下三个分支：

   	- Go release：最新稳定版，实际开发最佳选择
   	- Go weekly：包含最近更新的版本，一般每周更新一次
   	- Go tip：永远保持最新的版本，相当于内测版

   当你在使用不同的版本时，注意官方博客发布的信息，因为你所查阅的文档可能和你正在使用的版本不相符。

## 链接

- [目录](directory.md)
- 上一节：[Go 环境变量](02.2.md)
- 下一节：[在 Mac OS X 上安装 Go](02.4.md)
- 