#!/bin/bash

EXERCISM_DIR="$1"

#
# fetch upstream exercism_completion.bash in case it's missing
#
if [ ! -f "${EXERCISM_DIR}/config/exercism_completion.bash" ]; then
	curl -s http://cli.exercism.io/exercism_completion.bash > "${EXERCISM_DIR}/config/exercism_completion.bash"
fi

#
# enable bash completion
#
source "${EXERCISM_DIR}/config/exercism_completion.bash"

#
# wrap gpg-encrypted configuration decryption at exercism-cli runtime
#
function exercism() {
	#
	# dirty overwrite to not expose the API key publically
	#
	EXERCISM_JSON_CONTENT=`gpg --decrypt --quiet --batch "${EXERCISM_DIR}/config/exercism.json.gpg"`
	/usr/local/bin/exercism --config <(echo ${EXERCISM_JSON_CONTENT}) "$@"
}
