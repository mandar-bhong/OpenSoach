@echo off
setlocal 
  setx GOPATH %~dp0
  
  echo %GOPATH%
  
  timeout 10
  
endlocal