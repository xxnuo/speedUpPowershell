# Powershell 启动加速器

## 介绍

在 Windows 10 及以上系统最适合的 Shell 只有 Powershell，

但是 Powershell 的初次启动速度令人难以忍受的慢。

所以写了这个简单的程序。

## 使用方法

从`release`下载并将`speedUpPowershell.exe`放入`启动`目录。

`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`

即可，每次开机后程序会自动启动，

程序会在后台启动一个 Powershell 并打印一个字符串后退出，

达到加速 Powershell 启动的目的。

## 注意

杀毒软件可能会误报或者拦截，因为程序会在后台启动一个 Powershell

自行添加白名单即可。