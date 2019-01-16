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

./test.bin 
check "no flags"

./test.bin -b | grep -q "bool = true"
check "bool set"

./test.bin --int 42 | grep -q "integer = 42"
check "integer = 42"

./test.bin --string foobar | grep -q "string = foobar"
check "string = foobar"

./test.bin --bool | grep -q "bool = true"
check "long flag correctly resolved"

./test.bin --notint foobar | grep -q "integer flag --notint missing or no integer value given"
check "type mismatch recognized"

./test.bin --help | egrep "Usage|bool|Another" | wc -l | grep -q 3
check "usage text looks complete"

echo "============="
echo "OK: all tests"
exit 0
