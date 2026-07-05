@echo off
@REM 添加gendef和dlltool到环境变量
set PATH=E:\go\resources\go-sciter\xp\w64devkit-x86-2.0.0\bin;%PATH%;

@REM 第一步：用 gendef 从 sciter.dll 自动导出所有符号到 sciter.def
gendef dll\windows.xp\x32\sciter.dll

@REM 第二步：用 dlltool 根据 .def 生成静态导入库 libsciter.a
dlltool -d sciter.def -l libsciter.a -k