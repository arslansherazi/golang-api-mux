#!/bin/bash
git add .
git commit -m "$1"
git push origin "$2"
git push gitlab "$2"
git push bitbucket "$2"
