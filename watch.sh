
# Will monitor the memory usage of running executables named 'test'
watch -n1 ps -C test -o cmd,args,%cpu,%mem,size

