#!/bin/bash

openssl genrsa -out flui.key 2048
openssl req -new -x509 -key flui.key -out flui.pem -days 3650
