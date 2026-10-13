@echo off
title gitgwtuh - Terminal Git Explorer
cls
if not exist "gitgwtuh.exe" (
    echo [INFO] Mengompilasi binary gitgwtuh...
    go build -ldflags="-s -w" -o gitgwtuh.exe ./cmd/gitgwtuh
)
gitgwtuh.exe
