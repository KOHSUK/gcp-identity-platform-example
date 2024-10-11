#!/bin/bash

# Download and run the ORY installation script
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -b . ory

sudo mv ./ory /usr/local/bin/

# Ensure the /usr/local/bin directory is in the PATH
export PATH=$PATH:/usr/local/bin
