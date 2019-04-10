#!/usr/bin/env python
from __future__ import absolute_import
import threading
import logging
import time
import json
import os
import sendgrid
from sendgrid.helpers.mail import *
from string import Template
from kafka import KafkaConsumer, KafkaProducer


class AlertEmail():
  def __init__(self, api_key, from_email, to_email, subject_template, content_template):
    self.api_key          = api_key
    self.from_email       = Email(from_email)
    self.to_email         = Email(to_email)
    self.subject_template = Template(subject_template)
    self.content_template = Template(content_template)
    self.sg = sendgrid.SendGridAPIClient(apikey=api_key)
  def SendMail(self, sub, con):
    mail = Mail(self.from_email,
                self.subject_template.substitute(sub=sub),
                self.to_email,
                Content("text/plain",self.content_template.substitute(con=con)))
    response = self.sg.client.mail.send.post(request_body=mail.get())
    print(response.status_code)
    print(response.body)
    print(response.headers)

class Consumer(threading.Thread):
  daemon = True

  def __init__(self, callback=None):
    self.callback = callback
    super(Consumer, self).__init__()

  def run(self):
    consumer = KafkaConsumer(bootstrap_servers='kf-service:9092',
                 group_id='alerting-mailer',
                 value_deserializer=self.parseJSON)
    consumer.subscribe(['alerts'])
    for message in consumer:
      self.callback(message.value)
  def parseJSON(self, m):
    try:
      return json.loads(m.decode('utf-8'))
    except:
      return json.loads('{"parseError":"true"}')


class Manager():
  def __init__(self):
    self.msg_count = 0
    with open('/etc/sendgrid/sendgrid-key', 'r') as f:
      sendgrid_key=f.read().replace('\n', '')
    print("Sendgrid Key:", sendgrid_key)
    self.alert_mail = AlertEmail(
              sendgrid_key,
              "noreply@varnost.io",
              "tylerhoyt93@gmail.com",
              "Varnost Alert: $sub",
              "Email Body:\n\n$con")

  def start(self):
    Consumer(self.new_alert).start()
    while True:
      print("Current Count:", self.msg_count)
      time.sleep(10)

  def new_alert(self, alert):
    self.msg_count = self.msg_count + 1
    print(alert)
    print(type(alert))
    print("Sending Email")
    try:
      alert_name = alert['name']
    except:
      alert_name = "Unknown"
    self.alert_mail.SendMail(alert_name, json.dumps(alert, indent=2))



if __name__ == "__main__":
  m = Manager()
  m.start()