#! /bin/bash

pack build \
	--builder gcr.io/buildpacks/builder:v1 \
	--env GOOGLE_FUNCTION_SIGNATURE_TYPE=http \
	--env GOOGLE_FUNCTION_TARGET=$1 \
	${1,,}-function