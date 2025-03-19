#!/bin/bash
rm -rf /tmp/danad-temp

NUM_CLIENTS=128
danad --devnet --appdir=/tmp/danad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
DANAD_PID=$!
DANAD_KILLED=0
function killDanadIfNotKilled() {
  if [ $DANAD_KILLED -eq 0 ]; then
    kill $DANAD_PID
  fi
}
trap "killDanadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $DANAD_PID

wait $DANAD_PID
DANAD_EXIT_CODE=$?
DANAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Danad exit code: $DANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DANAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
