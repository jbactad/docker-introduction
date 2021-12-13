FROM python:3.11.0a3-alpine3.15

EXPOSE 3000

WORKDIR /app

COPY . .

CMD python -m http.server 3000

