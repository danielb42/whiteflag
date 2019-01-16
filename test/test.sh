#!/usr/bin/env bash

function check {
  if [[ "$?" == "0" ]]; then
    echo OK: $1
    echo
  else
    echo FAIL: $1
    exit 1
  fi
}

function checkFail {
  [[ "$?" == "0" ]] && false || true
  check "$@"
}

./test.bin 
check "no flags"

./test.bin -s
check "short bool set"

./test.bin --short
check "long bool set"

./test.bin --long 42
check "long int = 42"

./test.bin --notint foobar
checkFail "foobar is not an int"

./test.bin --long foobar
check "long string = foobar"

echo "============="
echo "OK: all tests"
exit 0
