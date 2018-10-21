FROM golang:latest

RUN useradd tsurubee && \
    echo "tsurubee:testpass" | chpasswd && \
    mkdir -p /home/tsurubee/.ssh

ADD misc/testdata/authorized_keys /home/tsurubee/.ssh

RUN chmod 600 /home/tsurubee/.ssh/authorized_keys
