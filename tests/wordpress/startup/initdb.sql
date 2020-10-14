#delete all users
DELETE FROM mysql.user;
DROP DATABASE test;

CREATE DATABASE db_wordpress;
#GRANT ALL PRIVILEGES ON db_wordpress.* to wp_user@'%' identified by '12345';
GRANT ALL PRIVILEGES ON db_wordpress.* to wp_user@'localhost' identified by '12345';