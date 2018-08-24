#!/bin/bash

openssl genrsa -out tk.key 2048
openssl req -new -x509 -key tk.key -out tk.pem -days 3650
openssl genrsa -out tk.rsa 1024
openssl rsa -in tk.rsa -pubout > tk.rsa.pub


