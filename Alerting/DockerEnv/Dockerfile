FROM python:3

COPY app/ /app/

WORKDIR /app/

RUN pip install -r requirements.txt

CMD ["/app/alerting.py"]