#!/bin/bash

# Download and run the ORY installation script
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b . hydra v2.2.0

sudo mv ./hydra /usr/local/bin/

curl -sSf https://atlasgo.sh | sh -y