# Powershell 启动加速器

## 介绍

在 Windows 10 及以上系统最适合的 Shell 只有 Powershell，

但是 Powershell 的初次启动速度令人难以忍受的慢。

所以写了这个简单的程序。

## 使用方法

从 [release](https://github.com/xxnuo/speedUpPowershell/releases) 下载`speedUpPowershell.exe`并将其放入`启动`目录。

`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`

即可，每次开机后程序会自动启动，

程序会在后台启动一个 Powershell 并打印一个字符串后退出，

达到加速 Powershell 启动的目的。

## 注意

杀毒软件可能会误报或者拦截，因为程序会在后台启动一个 Powershell

自行添加白名单即可。

# English

# Powershell Startup Accelerator

## Introduction

The most suitable shell for Windows 10 and above is only Powershell,

However, the initial startup speed of Powershell is unbearably slow.

So I wrote this simple program.

## How to use

Download `speedUpPowershell.exe` from [release](https://github.com/xxnuo/speedUpPowershell/releases) and place it in the `StartUp` directory.

`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`

That's it, the program will automatically start after each login,

The program will start a Powershell in the background and print a string before exiting,

Achieving the purpose of accelerating the second startup of Powershell.

## Note

Antivirus software may give false positives or intercept because the program will start a Powershell in the background.

Just add it to the whitelist on your own.