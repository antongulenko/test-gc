
mem=$1
if [ -z "$mem" ]; then
	mem=500
fi

# Will execute the test program and report on the maximum memory usage.
/usr/bin/time -f 'Max resident set size: %M' `pwd`/test -mem $mem -rand -randsize

