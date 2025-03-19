#!/bin/bash
rm -rf /tmp/danad-temp

danad --devnet --appdir=/tmp/danad-temp --profile=6061 --loglevel=debug &
DANAD_PID=$!
DANAD_KILLED=0
function killDanadIfNotKilled() {
    if [ $DANAD_KILLED -eq 0 ]; then
      kill $DANAD_PID
    fi
}
trap "killDanadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DANAD_PID

wait $DANAD_PID
DANAD_KILLED=1
DANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Danad exit code: $DANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DANAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
