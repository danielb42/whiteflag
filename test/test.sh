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

./test.bin -b | grep -q "bool set"
check "bool set"

./test.bin --bool | grep -q "bool set"
check "long flag correctly resolved"

./test.bin --int 42 | grep -q "integer = 42"
check "integer = 42"

./test.bin --notint foobar | grep -q "flag --notint missing or no integer value given. did you mean the string flag --notint?"
check "type mismatch recognized and correct type guessed"

./test.bin --string foobar | grep -q "string = foobar"
check "string = foobar"

./test.bin --help | egrep "Usage|bool|Another" | wc -l | grep -q 3
check "usage text looks complete"

echo "============="
echo "OK: all tests"
exit 0
