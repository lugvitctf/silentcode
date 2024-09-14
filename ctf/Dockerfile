FROM python:3.11-alpine
RUN apk update && apk add --no-cache curl bash
COPY question /root/question
COPY gen_flag /root/gen_flag
RUN chmod +x /root/gen_flag
# RUN /bin/bash /root/gen_flag 'c0d{somethingthatyoudontknow}'
RUN pip install flask
EXPOSE 8080
WORKDIR /root/question
CMD ["python3", "-m", "http.server", "8080"]