---
title: Hello Xamarin.mac
date: 2019-07-05 21:27:59
tags: [Xamarin,Mac,踩坑笔记]
category: .Net
comments: true
---



## 前言



***踩坑笔记是记录遇到的一些问题, 用于个人反思与后续复盘, 并不一定提供解决方案。当然如果你有好的解决思路, 欢迎留言。***



## 背景

最近得到一个新的开发任务：将公司现有的项目移植到 Mac 平台上.

先展示一下项目的结构图

![JmNev8.png](https://s1.ax1x.com/2020/04/18/JmNev8.png)



从上图可以看出, 整个项目思路就是把客户端当作处理web端与硬件设备的数据中转。前端与客户端使用 *SignalR* 进行通讯,  而客户端通过USB与硬件通讯。



## 预期

作者建议的技术方案是：Python 与 .Net Core 混合开发。

为什么会这么思考呢？是基于如下三个点考虑的：

​	1.Mac os 自带Python2.7环境, 利用 Python 绘制 Gui , .Net core 作为它的核心服务。（.Net Core能打包成单文件）减少依赖环境。

​	2.Windows上客户端使用 SignalR 与前端进行通讯, Core是明确支持 SignalR 的, 这样前端无需更改, 缩短开发周期。

​    3.当时 .Net Core 3.0即将发布, 在提前公布的版本特性中, .Net Core 将支持 WPF , 而在 Windows 上客户端就是使用 WPF 开发的。这意味着只要将Windows 上的客户端改写成 .Net Core,  往后的客户端版本迭代, 只需要写一套核心代码就可以支持两个平台。



当然这个方案被否决了,  公司要求使用 Xamarin 。





## 问题一：**Xamarin.Mac 能直接使用SignalR 吗？**



答案是可行的。

在面向搜索编程的基础下,  在Xamarin 社区找到一篇[问答](https://forums.xamarin.com/discussion/35226/signalr-for-xamarin-mac), 从这篇文章可知,有人成功地在Xamarin Mac 项目中使用了 SignalR。

虽然他并未给出具体实现的细节,但是依旧给我极大的鼓舞。

至少目前行进的方向是可行的。



## 问题二：**Xamarin并未涵盖我需要的引用库**　　

依靠着官方文档及旧代码,  我开始了我的移植。

很快地我发现我缺少了一个库 Microsoft.Owin.Host.HttpListener,  在复查了一遍资料以旧代码构建后,  我意识到Windows 上创建本地端口就是依靠的它。

问题是：Xamarin 的 Nuget 并未找到此包。

但是随后，我在查阅Xamarin 文档的时,  发现Xamarin有三个不同的目标框架,  双击解决方案,  打开项目选项, 选择常规可以看到。（如下图） 

![img](https://s1.ax1x.com/2020/04/18/Jexb24.png)

![img](https://s1.ax1x.com/2020/04/18/Jex7PU.png)

　从文档中, 我们可以看出, Xamarin.mac Full 使用的是“桌面”版本的子集, 与.net Framework4.5及更高版本几乎完全相同。

　　选择Xamarin.mac Full 后 重新生成解决方案, 再次打开Nuget包管理, 搜索Dll, 成功找到, 添加。

　　编译通过, 调式模式执行, 打开前端测试页面进行连接测试。

　　顺便说明一下, SignalR通讯可以在NetWork 的 WS 中看到传输内容。在Mac 上使用谷歌浏览器会报异常且无法看到传输协议, 但奇怪的是并不影响正常通讯,  只是无法查看。

<img src="https://s1.ax1x.com/2020/04/18/JexHGF.png"  />



## 问题三：**Xamarin.Mac 兼容问题？**　

在调试完客户端与硬件的交互后, 一个最严重的问题暴露出来：命令执行的时间过长。

在经过反复断点调试, 客户端给硬件发指令, 只有第一次执行的时候是属于正常范围（15-30 S）, 第二次执行则需要三到五分钟不等。

我把注意力集中在与硬件通讯的第三方库上：Hidsharp。查阅相关资料后, 发现这个库是兼容多个平台的。通过阅读源码发现它使用了两套不同的线程调用方式来实现流传输。

我猜测是线程堵塞造成命令执行过慢。

接下来三天我都在调试这个问题, 我查阅了关于线程的资料并尝试更改这部分代码, 但是没有任何改善。我意识到我可能走错了方向,我一心认为是线程的问题, 但是以我目前的水准我没有解决它的能力。

那既然开源库支持Mac平台我相信他们一定做过测试。我重新写了一个客户端, 选择 .Net Framework作为基础框架,使用相同的源码。经过对比测试 .Net Framework 发送指令没有任何问题，无论是响应速度还是命令执行速度……

使用相同的源码, 目标框架不同, 执行的效率相差甚远, 由此推测差异可能出现在目标框架与Hidsharp兼容上。

很遗憾, 这篇博文到这里基本已经翻车了, 我既没能力去更改多线程, 也没能力去更改框架, 甚至找不到继续前进的方向。



> 在Mac 上使用 Sleep(1000)会导致电脑进入睡眠模式。



## 后续

按照最初的想法Python 和 .Net Core 混合开发, 我私下进行了调研, .Net Core 使用Hidsharp 暂时未出现性能上的问题, 而Python 启动 .Net Core则需要获取管理员权限（这一块也是踩了坑的）。

最后用Swift 代替了 Python 的部分, 毕竟Swift 更加适配 Mac os, 获取管理员权限也更加方便。



## 总结

这次踩坑也是暴露了自己自身的一些问题：

1.首先是对多线程的了解太少, 因为不懂, 所以在开发过程中遇到问题容易先入为主, 进而钻进死胡同。

2.其次阅读源码过少, 对原作者的设计不能很好地理解。 



其他地方可能会有人有不同的意见：

1.比如既然使用Swift 为什么不坚持全套写完呢？

这是考虑到后期客户端进行版本迭代的成本（毕竟招一个Swift程序员专门去迭代Mac 客户端也太浪费了）。如果按照作者的预期, 使用.Net Core 只需写一套代码就可以做到多个平台进行迭代, 却别只是不同平台Gui部分的“外壳”。



2.可能会有人觉得这种 “外壳” 加核心程序的方案比较诡异。

其实还是成本问题的考虑, 使用某一种语言进行全书写固然更加健全。但是考虑到公司技术栈和学习成本及开发周期, 算得上是某种妥协。毕竟如果有Linux 移植的需求, 可以直接将核心代码移植过去, 只需编写一个简单的 “外壳”。

>.Net core 3.0 支持WPF, 最开始我还以为能直接在Mac 上运行, 结果当然是失败啦。



3.为什么要使用 Xamarin？

因为公司有 Xamarin 开发App 的经验, 可能觉得使用熟知的技术栈会更加稳妥一点。不过使用Xamarin开发Mac 客户端还是太少了, 相关资料也少, 同事也很难给出有效的建议, 通常为了佐证某一个想法需要写大量的Demo 代码去验证。

> 个人是不会再深入研究Xamarin的