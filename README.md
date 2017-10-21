## **GOsUgar**

> GoCn 2017 goHackathon

Golang 语法糖 管理工具

## 愿景

希望作为一个类 babel 至于 js 的工具。 GOsUgar 期望通过标准化，流程化，插件化，一种基于 golang ast + golang generate + golang comment的轻量级无侵入的扩展 Golang 语法**管理工具**。

// 又一个包管理工具 😓



## 什么？ Golang 还有语法糖

是的 你没有看错。虽然这个语法糖基于 comment。

例如:

一种优雅的 error != nil

```
#try err
{
  ......
  return err
}
or
#? err
{
  ......
  return err
}


等效于

if err != err {
  ......
  return err
}
```



一种优雅的装饰器  [Todo]

```
#@ required_login arg0 arg1
func admin(...){
  {detail}
}


等效于
func admin_decorator_0(){
  {detail}
}

func admin(...){
  required(arg0, arg1, admin_decorator_0) 
}
```



## 团队

GoSugar 目前就一个工作中用 Python，个人项目 90% 用 Python ，痴迷效率，懒驱动开发，节约时间玩吃鸡的,迷恋 Python 的高效和简洁。又想要 Golang 的并发友好和性能的弱鸡程序猿— wph95



## 参考和启发

https://github.com/riolet/rix

[todo]