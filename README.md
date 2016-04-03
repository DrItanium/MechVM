This is a port of the MechVM 2011-04-04 source code from C++ to golang.

The code was originally written by Björn Ganster but does not build correctly
on modern linux distros. The original code attempted to take wide characters
into account and thus had a lot of NIH in its design. Attempting to port that
ball of code is an exercise for a later date. Instead I want to port the code
to golang
