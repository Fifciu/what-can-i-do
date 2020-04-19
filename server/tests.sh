#!/bin/bash
cp .env controllers/.env;
cp .env middlewares/.env;
cp .env models/.env;
cp .env utils/.env;
go test ./...;
rm -f controllers/.env;
rm -f middlewares/.env;
rm -f models/.env;
rm -f utils/.env;