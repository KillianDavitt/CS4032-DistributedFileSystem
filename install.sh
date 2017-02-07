#!/bin/bash
cd lockServer
./certs_gen.sh
cd ../authServer
./certs_gen.sh
cd ../dirServer
./certs_gen.sh
cd ../transactionServer
./certs_gen.sh
