#!/bin/bash

# This script is executed after the container is created.
# You can add any setup commands you need here.

if [ ! -f .env ]; then cp .env.example .env; fi
