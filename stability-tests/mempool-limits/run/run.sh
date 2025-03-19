#!/bin/bash

APPDIR=/tmp/danad-temp
DANAD_RPC_PORT=29587

rm -rf "${APPDIR}"

danad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${DANAD_RPC_PORT}" --profile=6061 &
DANAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${DANAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $DANAD_PID

wait $DANAD_PID
DANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Danad exit code: $DANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DANAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
