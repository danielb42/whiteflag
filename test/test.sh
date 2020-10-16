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

./test.bin --testdefaults | grep -q "defint = 123"
check "default int"

./test.bin --testdefaults | grep -q "defstr = foobar"
check "default str"

./test.bin --testdefaultalreadyset1 | grep -q "default value for --defint already set"
check "default value already set"

./test.bin --testdefaultalreadyset2 | grep -q "default value for --defint already set"
check "default value already set (across types)"

./test.bin --testrequired | grep -q "integer flag --required missing or no integer value given"
check "missing required param"

./test.bin -a | grep -q "bool = true"
check "unaliased short bool set"

./test.bin -b | grep -q "bool = true"
check "aliased short bool set"

./test.bin --testboolfalse | grep -q "bool = false"
check "absent bool param = false"

./test.bin --bool notbool | grep -q "followed by a non-bool value"
check "bool in non-bool context recognized"

./test.bin --int 42 | grep -q "integer = 42"
check "integer = 42"

./test.bin --string foobar | grep -q "string = foobar"
check "string = foobar"

./test.bin --bool | grep -q "bool = true"
check "long flag correctly resolved (1)"

./test.bin --cflag 42 | grep -q "cflag = 42"
check "long flag correctly resolved (2)"

./test.bin --notint foobar | grep -q "integer flag --notint missing or no integer value given"
check "type mismatch recognized (1)"

./test.bin --notstring | grep -q "string flag --notstring missing or no string value given"
check "type mismatch recognized (2)"

./test.bin --help | egrep "Usage|bool|Another|emptydesc" | wc -l | grep -q 6
check "usage text looks complete"

./test.bin -a -b -a | grep -q "\-a specified more than once"
check "duplicate short flags recognized"

./test.bin --foobar -a --foobar | grep -q "\--foobar specified more than once"
check "duplicate long flags recognized"

./test.bin --testredefineh | grep -q "cannot re-define"
check "re-defining -h prevented"

./test.bin --testredefinehelp | grep -q "cannot re-define"
check "re-defining -help prevented"

./test.bin --testshorttoolong | grep -q "must not be longer than 1 char"
check "short name too long"

./test.bin --testlongtooshort | grep -q "must be longer than 1 char"
check "long name too short"

./test.bin --testlongalreadyaliased | grep -q "already aliased to another short flag"
check "long flag already aliased"

./test.bin --testshortalreadyaliased | grep -q "already has an associated long flag"
check "short flag already aliased"

./test.bin --testemptyaliasfields1 | grep -q "short flag aliased to --foobar must not be empty"
check "tried to alias empty short flag"

./test.bin --testemptyaliasfields2 | grep -q "long flag aliased to -f must not be empty"
check "tried to alias empty long flag"

./test.bin --testemptyaliasfields3 | grep -q "long flag aliased to -f must not be empty"
check "tried to alias empty (space) long flag"

./test.bin --testemptyaliasfields4 | grep -q "long flag aliased to -f must not be empty"
check "tried to alias empty (spaces) long flag"

echo "============="
echo "OK: all tests"
exit 0
