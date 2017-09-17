#!/bin/sh

make build
docker build -t raven-api:latest .
