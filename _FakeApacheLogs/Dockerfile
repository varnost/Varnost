FROM python:2.7
# Install Filebeat
COPY filebeat-6.5.0-amd64.deb /filebeat-6.5.0-amd64.deb
RUN dpkg -i /filebeat-6.5.0-amd64.deb
RUN rm -rf /filebeat-6.5.0-amd64.deb

COPY Fake-Apache-Log-Generator-master /usr/share/log-gen
RUN pip install -r /usr/share/log-gen/requirements.txt
RUN mkdir /var/log/apache

COPY filebeat.yml /etc/filebeat/filebeat.yml

COPY entrypoint.sh /bin/entrypoint.sh

ENTRYPOINT ["/bin/entrypoint.sh"]
