
This is a small Go program to test the garbage collector functionality of a Go compiler.

It will allocate a lot of []byte (random sizes), fill them with random data, but only keep a few of them referenced.

In the end, it will output how much memory has been allocated and how big the referenced memory was in average.

The shell scripts can be used to monitor the real memory usage of the program.

Then, the output of the program can be compared to the output of the scripts, to see if (and how well) the garbage collector did its job.

