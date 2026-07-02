## 项目说明
- 本项目,是我clone github上的go-sciter的master版本,原项目网址是:
https://github.com/sciter-sdk/go-sciter
- 最后一个可以兼容XP的commit是:https://github.com/sciter-sdk/go-sciter/commit/a04e052a28133d8a79c82b53fc861d1e473c0499,在这个之后的commit(99cd4de65a26163ff93872ef7bba888b479081dc),都不兼容XP了。
- 本项目是go的cgo的sciter的绑定,因为sciter的头文件中有inline等c++的,如 sciter-x-api.h中,存在inclie,而cgo只能编译C,所以go-sciter的作者,写了sciter-x-api.c文件,是为了cgo能够正常编译.
- 原项目已经很久没有更新了,现在的sciter官方已经早就开始兼容XP了.
现在sciter的官方版本是:6.0.4.5,仓库是:
https://gitlab.com/sciter-engine/sciter-js-sdk
- 本项目的include6目录,就是sciter官方版本的include目录.
- 本项目,是基于go-sciter的最新master版本修改后的.
- 我已经试过,使用官方的6.0.4.5版本的的scapp.exe,搭配官方的sciter.dll,在XP上运行,没有问题.
- 因为go-sciter长时间不更新,本项目目前在xp上运行有bug
- 本项目使用的环境是:go 1.10.8 i386(兼容XP最后版本),gcc:w64devkit-x86-2.0.0(兼容XP)
- 当前,本项目在windows10上编译,运行,会成功显示窗口,但有个小问题,关闭程序后,黑色窗口不会自动关闭.但在XP上运行,连窗口都不显示,直接报错.

## 目标
- 我想把本项目改为兼容XP的.
- 编译本项目,请执行"E:\go\resources\go-sciter\xp\gopath\src\go-sciter-xp-test\一键重新build.bat",并且,如果编译成功,会在 E:\go\resources\go-sciter\xp\gopath\src\go-sciter-xp-test\bin 目录下生成go-sciter-xp-test.exe文件.才算编译成功,至于是否在XP上成功运行,需要我手动拷贝到XP电脑上运行.

