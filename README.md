<div align=center>
<img src="./docs/golang_1.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.11-blue"/>
<img src="https://img.shields.io/badge/goframe-1.13.1-lightBlue"/>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
</div>
English | [简体中文](./README-zh_CN.md)

# Project documentation

[Online documentation (to be improved)](https://github.com/flipped-aurora/gf-vue-admin)

- Front-end UI framework：[element-ui](https://github.com/ElemeFE/element) 

- Background framework：[GoFrame](https://goframe.org/index)

## 1. basic introduction

### 1.1 Project Introduction

[Online preview (to be improved)](https://github.com/flipped-aurora/gf-vue-admin)

> Gf-vue-admin is a full-stack back-end management system based on vue and GoFrame. It integrates jwt authentication, dynamic routing, dynamic menu, casbin authentication, form generator, code generator and other functions, providing multiple 
>
> This kind of sample files allows you to devote more time to business development.

### 1.2 Contribution guide

Hi! First of all, thank you for using gf-vue-admin.

GoFrame-vue-admin is a set of open source frameworks prepared for the back-end management platform with a separated architecture of front and back ends, aiming to quickly build a back-end management system.

The growth of GoFrame-vue-admin is inseparable from everyone's support. If you are willing to contribute code or provide suggestions for GoFrame-vue-admin, please read the following.

#### 1.2.1 Issue specification

- Issues are only used to submit bugs or features and design-related content, other content may be directly closed。If you have questions while using，Go to Slack or [Gitter](https://gitter.im/ElemeFE/element) Consulting。

- Before submitting an issue, please search whether the relevant content has been submitted。

#### 1.2.2 Pull Request specification

- Please fork a copy to your own project first, do not directly branch under the warehouse。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- <font color=red>确保 PR 是提交到 `develop` 分支，而不是 `master` 分支。</font>

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

### 1.3 Version list

- master: 1.0, function is currently being tested

## 2. Instructions for use

```
-node version> v8.6.0
-golang version >= v1.11
-IDE recommendation: Goland
-After the clone project, after you import the db file into the library you created, it is best to go to Qiniu Cloud to apply for your own space address.
-Replace the Qiniu Cloud public key, private key, warehouse name and default URL address in the project to avoid data confusion in the test file
```

### 2.1 web

```bash
# clone the project
git clone https://github.com/piexlmax/gin-vue-admin.git

# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 2.2 server

```bash
# Use go.mod
# Install go dependencies
go list (go mod tidy)

# Compile
go build
```

### 2.3 swagger automation API documentation

- Reasons to remove swagger
	- Comment redundant code, resulting in bloated code
	- Recommend alternative tool apipost

## 3. Technical selection

- 前端：用基于`vue`的`Element-UI`构建基础页面。
- 后端：用`GoFrame`快速搭建基础restful风格API，`GF(Go Frame)`是一款模块化、高性能、生产级的Go基础开发框架。实现了比较完善的基础设施建设以及开发工具链，提供了常用的基础开发模块，如：缓存、日志、队列、数组、集合、容器、定时器、命令行、内存锁、对象池、配置管理、资源管理、数据校验、数据编码、定时任务、数据库ORM、TCP/UDP组件、进程管理/通信等等。并提供了Web服务开发的系列核心组件，如：Router、Cookie、Session、Middleware、服务注册、模板引擎等等，支持热重启、热更新、域名绑定、TLS/HTTPS、Rewrite等特性。
- 数据库：采用`MySql`(8.0.19)版本，使用`gdb`实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- 配置文件：使用`fsnotify`和`viper`实现`yaml`格式的配置文件。
- 日志：使用`go-logging`实现日志记录。

## 4. Project structure

### 4.1 System architecture diagram

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 Front-end detailed design drawing （provider:<a href="https://github.com/baobeisuper">baobeisuper</a>）

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 目录结构

```
    ├─server  	     （backend）
    └─web            （frontend）
        ├─public        （deploy templates）
        └─src           （source code）
            ├─api       （frontend APIs）
            ├─assets	（static files）
            ├─components（components）
            ├─router	（frontend routers）
            ├─store     （vuex state management）
            ├─style     （common styles）
            ├─utils     （frontend common utilitie）
            └─view      （pages）


```

[See this directory for the backend](./server/document/后端目录结构.md)

## 5. The main function

- authority management：based on [gf-jwt](https://github.com/gogf/gf-jwt) And [casbin](https://github.com/casbin/casbin) Implemented rights management
-  File upload and download：Realize file upload operation based on Qiniu Cloud
- Paging package：The front end uses mixins to encapsulate paging, and the paging method calls mixins
- User Management：System administrators assign user roles and role permissions。
- Role management：Create the main object of permission control, you can assign different api permissions and menu permissions to roles。
- Menu management：Realize user dynamic menu configuration, realize different menus for different roles。
- api management：Different users have different permissions on the api interface that can be called。
-  Rich text editor：MarkDown Editor function embedded。
- Conditional search：Add conditional search example。
- restful Example：You can refer to the sample API in the user management module。 
- Multi-sign-in restrictions：需要在`config.toml`中把`system`中的`UseMultipoint`修改为true
- Split long pass：Provide examples of file segment upload and large file segment upload function
- Form builder：The form builder uses [@form-generator](https://github.com/JakHuang/form-generator)。
- Code generator：Basic background logic and simple curd code generator。 

## 6. Scheduled Tasks

- [ ] Import, export to Excel
- [ ] Echart chart support
- [ ] Workflow, task handover function development
- [ ] Separate front-end usage mode and data simulation
- [ ] User multi-role
- [ ] Self-written api batch import permission assignment list
- [ ] Automatic import of generated files

## 7. knowledge base

## 7.1 Team blog

> https://www.yuque.com/flipped-aurora
>
> There are front-end framework instructional videos inside. 
>
> If you think the project is helpful to you, you can add my personal WeChat: shouzi_1994, welcome your valuable needs。

## 7.2 Teaching video

（1）Golang basic instructional video recording...
> https://space.bilibili.com/322210472/channel/detail?cid=108884

## 8. Contact information

### 8.1 Technology Group

| QQ群 |
|  :---:  |
| <img src="./docs/gf-vue-admin开源项目交流群.jpg" width="180"/> |

### QQ exchange group：1040044540

### WeChat exchange group: add WeChat account SliverHorn, note "join gf-vue-admin exchange group"

### 8.2 Project team members

| Jiang | Yi | Yan | Du | Yin | Song |
|  :---:  |  :---: | :---: | :---:  |  :---: | :---: |
| <img width="150" src="http://qmplusimg.henrongyi.top/qrjjz.png"> | <img width="150" src="http://qmplusimg.henrongyi.top/qryx.png"> | <img width="150" src="http://qmplusimg.henrongyi.top/qryr.png"> | <img width="150" src="http://qmplusimg.henrongyi.top/qrdjl.png"> | <img width="150" src="http://qmplusimg.henrongyi.top/qrygl.png"> | <img width="150" src="http://qmplusimg.henrongyi.top/qrsong.png"> |

| Nick name | Project position | First name |
|  ----  | ----  | ----  |
| [@piexlmax](https://github.com/piexlmax)  | Project sponsor | Jiang |
| [@granty1](https://github.com/granty1)  | Backend developer | Yin |
| [@Ruio9244](https://github.com/Ruio9244)  | Full-stack developer | Yan |
| [@1319612909](https://github.com/1319612909)  | UI developer |  Du |
| [@krank666](https://github.com/krank666)  | Frontend developer | Yin |
| [@chen-chen-up](https://github.com/chen-chen-up)  | Novice developer | Song |
| [@SliverHorn](https://github.com/SliverHorn)  | Community Administrator | Lai |

## 9. Donate

If you find this project useful, you can buy author a glass of juice 🍹 [here](http://doc.henrongyi.top/more/coffee.html)

## 10. Commercial considerations

If you use this project for commercial purposes, please comply with the Apache2.0 agreement and retain the author's technical support statement.