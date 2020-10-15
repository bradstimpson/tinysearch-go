#!/bin/bash


HOME_DIR="/startup/"
INIT="${HOME_DIR}init.fin"

if [ ! -f ${INIT} ] ; then
 echo "initalize mariadb"
 mysql_install_db
 mysqld_safe &
 sleep 5
 mysql -uroot < /startup/initdb.sql
 touch ${INIT}

else
  mysqld_safe &

fi

httpd -DFOREGROUND