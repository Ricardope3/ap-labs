#include <stdio.h>
#include <syslog.h>
#include <stdarg.h>
#include "logger.h"
#include <string.h>
#define DEFAULT_LOG 0
#define SYS_LOG 1

#define RESET 0
#define BRIGHT 1
#define DIM 2
#define UNDERLINE 3
#define BLINK 4
#define REVERSE 7
#define HIDDEN 8

#define BLACK 0
#define RED 1
#define GREEN 2
#define YELLOW 3
#define BLUE 4
#define MAGENTA 5
#define CYAN 6
#define WHITE 7

int logFlag = DEFAULT_LOG;

int initLogger(char *logType, ...)

{
    if (strcmp(logType, "syslog") == 0)
    {
        logFlag = SYS_LOG;
        openlog("Advanced Logger -", LOG_CONS, LOG_USER);
        printf("Initializing Logger on: %s\n", logType);
    }
    else
    {
        logType = "stdout";
        printf("Initializing Logger on: %s\n", logType);
    }
    return 0;
}

int infof(const char *format, ...)
{
    va_list list;
    va_start(list, format);
    if (logFlag == DEFAULT_LOG)
    {

        printf("\x1b[44m INFO \x1b[0m     ");
        vprintf(format, list);
    }
    else
    {
        syslog(LOG_INFO, format, list);
    }
    return 0;
}

int warnf(const char *format, ...)
{

    va_list list;
    va_start(list, format);
    if (logFlag == DEFAULT_LOG)
    {

        printf("\x1b[30;43m WARNING \x1b[0m  ");
        vprintf(format, list);
    }
    else
    {
        syslog(LOG_WARNING, format, list);
    }

    return 0;
}
int errorf(const char *format, ...)
{

    va_list list;
    va_start(list, format);
    if (logFlag == DEFAULT_LOG)
    {
        printf("\x1b[45m ERROR \x1b[0m    ");
        vprintf(format, list);
    }
    else
    {
        syslog(LOG_ERR, format, list);
    }

    return 0;
}
int panicf(const char *format, ...)
{

    va_list list;
    va_start(list, format);
    if (logFlag == DEFAULT_LOG)
    {
        printf("\x1b[41m PANIC \x1b[0m    ");
        vprintf(format, list);
    }
    else
    {
        syslog(LOG_EMERG, format, list);
    }
    return 0;
}

void openLog()
{
    openlog("Advanced Logger -", LOG_CONS, LOG_USER);
}