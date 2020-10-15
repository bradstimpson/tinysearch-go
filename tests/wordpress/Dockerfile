FROM centos:7
LABEL author="kishitat"

#install php,httpd,mysql
RUN yum install -y epel-release && \
    rpm -Uvh http://rpms.famillecollet.com/enterprise/remi-release-7.rpm && \
    yum install -y --enablerepo=remi,remi-php71 php php-devel php-mbstring php-pdo php-gd httpd mariadb mariadb-server php-mysql wget

#create directory
RUN mkdir -p /opt/logs && \
    chmod -R a+rwX /opt/

#modify php config
RUN sed -i -e "s@;date.timezone =@date.timezone = \"Asia/Tokyo\"@" /etc/php.ini

#edit httpd setting
RUN mkdir -p /opt/httpd/logs && \
    sed -i -e "s/Listen 80/Listen 8080/" /etc/httpd/conf/httpd.conf && \
    sed -i -e "s@ErrorLog .*@ErrorLog /opt/logs/httpd_error_log@" /etc/httpd/conf/httpd.conf && \
    sed -i -e "s@    CustomLog .*@    CustomLog "/opt/logs/httpd_access_log" combined@" /etc/httpd/conf/httpd.conf && \
    chmod a+rwX /run/httpd/

#change mariadb setting
RUN /usr/bin/mysql_install_db && \
    sed -i -e "s@log-error=.*@log-error=/opt/logs/mariadb.log@" /etc/my.cnf && \
    chmod -R a+rwX /run/mariadb/ && \
    chmod -R a+rwX /var/lib/mysql

#download wordpress
RUN cd /tmp/ && \
    curl -LO https://wordpress.org/latest.tar.gz && \
    tar -xzvf latest.tar.gz && \
    cp -R /tmp/wordpress/* /var/www/html && \
    cp /var/www/html/wp-config-sample.php /var/www/html/wp-config.php

#edit wordpress config
RUN sed -i -e "s/database_name_here/db_wordpress/" /var/www/html/wp-config.php && \
    sed -i -e "s/username_here/wp_user/" /var/www/html/wp-config.php && \
    sed -i -e "s/password_here/12345/" /var/www/html/wp-config.php && \
    chmod -R a+rwX /var/www && \
    chown -R 1999:1999 /var/www

#ADD startup shells
RUN mkdir /startup/
COPY startup/* /startup/
RUN chmod -R a+rw /startup && \
    chmod a+x /startup/*.sh

USER 1999

EXPOSE 8080
# CMD tail -f /dev/null
ENTRYPOINT /startup/start.sh