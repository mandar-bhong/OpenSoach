@echo off
SETLOCAL ENABLEDELAYEDEXPANSION
echo Starting sql script executing

::-------------Set sql parameter------------------
SET Server=localhost
SET User=root
SET Password=welcome
SET Database=spl_hkt_node_0001
::----------------End ----------------------------

"C:\Program Files\MySQL\MySQL Server 5.7\bin\mysql.exe" -h%Server% -u%User% -p%Password% %Database% < %~dp0\testdata\testdata_1.0.0.sql


SET /P uname=Please enter any key to exit:


