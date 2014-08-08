package main

/*
#include <stdio.h>
#include <stdlib.h>

char* GoVersion();

void print_version()
{
	char *ver = GoVersion();
	printf("Go Version is %s\n",  ver);
	free(ver);
}
*/
import "C"

func main() {
	C.print_version()
}
