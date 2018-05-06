#!/bin/bash 
echo "Curent Path:" pwd

export builddir=/opt/build/spl

cd $builddir

chmod  +777 $builddir/SPLBuild/spl 


service spl stop

#this should be remove, currently service spl stop not working hence explicity kill
killall -9 spl

tar xvzf SPLBuild.tar.gz


#executing command from shell prompt 
mysql -u root -pwelcome -Bse "drop database spl_master;"
mysql -u root -pwelcome -Bse "create database spl_master;"
#session will be terminiated hence executing all the commands in same session
mysql -u root -pwelcome -Bse "use spl_master;source /opt/build/spl/SPLBuild/spl_master/schema/schema_1.0.0.sql;source /opt/build/spl/SPLBuild/spl_master/data/data_1.0.0.sql;source /opt/build/spl/SPLBuild/spl_master/testdata/testdata_1.0.0.sql;"

service spl start

