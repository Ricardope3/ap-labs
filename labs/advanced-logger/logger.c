#include <stdio.h>
#include "logger.h"

int initLogger(char *logType) {
    printf("Initializing Logger on: %s\n", logType);
    return 0;
}

int infof(const char *format)
{
    printf("\x1b[44m INFO \x1b[0m");
    printf("\x1b[34m %s \x1b[0m  \n", format);
    return 1;
}

int warnf(const char *format)
{
    printf("\x1b[30;43m WARNING \x1b[0m");
    printf("\x1b[33m %s \x1b[0m  \n", format);
    return 0;
}
int errorf(const char *format)
{
    printf("\x1b[45m ERROR \x1b[0m");
    printf("\x1b[35m %s \x1b[0m  \n", format);
    return 0;
}
int panicf(const char *format)
{
    printf("\x1b[41m PANIC \x1b[0m");
    printf("\x1b[31m %s \x1b[0m  \n", format);
    return 0;
}