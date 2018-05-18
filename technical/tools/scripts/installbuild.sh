#!/bin/bash 
echo "Curent Path:" pwd

export builddir=/opt/build/spl

cd $builddir

chmod  +777 $builddir/Build/spl 


service spl stop
service splserver stop
service hkt stop


systemctl daemon-reload


#this should be remove, currently service spl stop not working hence explicity kill
killall -9 spl
killall -9 splserver
killall -9 hkt

tar xvzf Build.tar.gz


#executing command from shell prompt 
mysql -u root -pwelcome -Bse "drop database spl_master;"
mysql -u root -pwelcome -Bse "create database spl_master;"
#session will be terminiated hence executing all the commands in same session
mysql -u root -pwelcome -Bse "use spl_master;source /opt/build/spl/Build/database/spl_master/schema/schema_1.0.0.sql;source /opt/build/spl/Build/database/spl_master/data/data_1.0.0.sql;"


#executing command from shell prompt 
mysql -u root -pwelcome -Bse "drop database spl_hkt_master;"
mysql -u root -pwelcome -Bse "create database spl_hkt_master;"
#session will be terminiated hence executing all the commands in same session
mysql -u root -pwelcome -Bse "use spl_hkt_master;source /opt/build/spl/Build/database/spl_hkt_master/schema/schema_1.0.0.sql;source /opt/build/spl/Build/database/spl_hkt_master/data/data_1.0.0.sql;"


#executing command from shell prompt 
mysql -u root -pwelcome -Bse "drop database spl_hkt_node_0001;"
mysql -u root -pwelcome -Bse "create database spl_hkt_node_0001;"
#session will be terminiated hence executing all the commands in same session
mysql -u root -pwelcome -Bse "use spl_hkt_node_0001;source /opt/build/spl/Build/database/spl_hkt_node/schema/schema_1.0.0.sql;source /opt/build/spl/Build/database/spl_hkt_node/data/data_1.0.0.sql;"

echo "ExecTestData Value:" $1
if [ "$1" == "true" ]; then
	mysql -u root -pwelcome -Bse "use spl_master;source /opt/build/spl/Build/database/spl_master/testdata/testdata_1.0.0.sql;"
  mysql -u root -pwelcome -Bse "use spl_hkt_master;source /opt/build/spl/Build/database/spl_hkt_master/testdata/testdata_1.0.0.sql;"
  mysql -u root -pwelcome -Bse "use spl_hkt_node_0001;source /opt/build/spl/Build/database/spl_hkt_node/testdata/testdata_1.0.0.sql;"
fi



sleep 5
service spl start
service splserver start

service hkt start

