#!/bin/bash
rm -rf /tmp/danad-temp

danad --devnet --appdir=/tmp/danad-temp --profile=6061 &
DANAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $DANAD_PID

wait $DANAD_PID
DANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Danad exit code: $DANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DANAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
