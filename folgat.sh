#!/bin/sh

if [ -f "$HOME/.folgat_chdir" ]; then
	FOLGAT_CHDIR_AS_DOTFILE_SAYS="$(cat "$HOME/.folgat_chdir")"
	cd $HOME && (>&2 echo "Chdired to $HOME because $HOME/.folgat_chdir was present.")
	cd "$FOLGAT_CHDIR_AS_DOTFILE_SAYS" &&
	        (>&2 echo "Chdired to path from $HOME/.folgat_chdir which was $FOLGAT_CHDIR_AS_DOTFILE_SAYS") ||
		(>&2 echo "Failed to chdir to path from $HOME/.folgat_chdir which was $FOLGAT_CHDIR_AS_DOTFILE_SAYS")
fi
TWITFOL_PATH=$(command -v twitfol)
if [ ! -x "$TWITFOL_PATH" ]; then
	TWITFOL_PATH="$GOPATH/bin/twitfol"
	if [ ! -x "$TWITFOL_PATH" ]; then
		exit 1
	fi
fi
set -o noclobber
fndate=$(date +"%FT%H%M%S%z")
FILENAME="twitfol_scrnam-tab-userid_${fndate}.dat"
"$TWITFOL_PATH" > "$FILENAME" &&
	>&2 echo "Finished writing to ${FILENAME}" || { twitfol_failure_error_code=${?}
	>&2 echo "Failure, error code ${?}. Filename ${FILENAME}"; exit $twitfol_failure_error_code;}
