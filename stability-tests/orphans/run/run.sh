#!/bin/bash
rm -rf /tmp/danad-temp

danad --simnet --appdir=/tmp/danad-temp --profile=6061 &
DANAD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $DANAD_PID

wait $DANAD_PID
DANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Danad exit code: $DANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $DANAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
