#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>

int infof(const char *format, ...)
{
    va_list list;
    va_start(list, format);
    printf("\x1b[44m INFO \x1b[0m");
    vprintf(format, list);
    return 0;
}

int warnf(const char *format, ...)
{
    va_list list;
    va_start(list, format);
    printf("\x1b[30;43m WARNING \x1b[0m");
    vprintf(format, list);
    return 0;
}
int errorf(const char *format, ...)
{
    va_list list;
    va_start(list, format);
    printf("\x1b[45m ERROR \x1b[0m");
    vprintf(format, list);
    return 0;
}
int panicf(const char *format, ...)
{
    va_list list;
    va_start(list, format);
    printf("\x1b[41m PANIC \x1b[0m");
    vprintf(format, list);
    abort();
    return 0;
}