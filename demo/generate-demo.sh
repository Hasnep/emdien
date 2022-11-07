#! /usr/bin/env bash

repo_root=$(git rev-parse --show-toplevel) \
    && cd $repo_root \
    && docker build --tag=emdien --file=demo/emdien.dockerfile $repo_root \
    && docker build --tag=vhs-emdien --file=demo/Dockerfile $repo_root \
    && docker run --rm --volume="$repo_root/demo/output:/vhs/output" vhs-emdien emdien.tape
