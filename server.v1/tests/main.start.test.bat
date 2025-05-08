@echo off
setlocal enabledelayedexpansion

@REM Переменные
set GREEN= [32m
set RED= [31m
set RESET= [0m

set GO=go
set TEST_FILES=tests\routes\*.go

@REM Начала тестов
echo Running tests...
%GO% test %TEST_FILES% 2 > &1 | findstr "FAIL" > nul
if %errorlevel% equ 0 (
    echo .
    %GO% test %TEST_FILES% | findstr /V "PASS"
    %GO% test %TEST_FILES% | findstr /V "FAIL" | (
        for /F "tokens=*" %%i in ('findstr "."') do (
            echo %RED%%%1%RESET%
        )
    )

    echo %RED%Tests failed%RESET%
    exit /b 1
) else (
    echo .
    %GO% test %TEST_FILES% | findstr /V "FAIL"
    %GO% test %TEST_FILES% | findstr "PASS" | (
        for /F "tokens=*" %%i in ('findstr "."') do (
            echo %GREEN%%%i%RESET%
        )
    )

    echo %GREEN%All tests passed%RESET%
)

echo .
echo Starting server...
%GO% build -o bin/flicksfi main.go
%GO% run main.go

endlocal
pause
