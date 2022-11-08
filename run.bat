SET BINARY_NAME=imdb-searcher
SET PROGRAM_TYPE=cmd
SET OUTPUT_PATH=bin\%BINARY_NAME%

set "PROJECT_DIR=%CD%"

@echo off 
echo %PROJECT_DIR%

go build -o %OUTPUT_PATH%.exe %PROGRAM_TYPE%\%BINARY_NAME%\main.go

start %OUTPUT_PATH%.exe

@echo on
