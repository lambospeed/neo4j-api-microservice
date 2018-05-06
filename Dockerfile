FROM python:3.7-rc

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . ./

RUN pip install -r requirements.txt

CMD ["python3.7", "app.py"]
