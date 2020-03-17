#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <regex.h>
#include <string.h>

#define REPORT_FILE "packages_report.txt"

struct Package
{
    char *name[50];
    char *installed_date[17];
    char *last_update_date[17];
    char *removal_date[17];
    int num_updates;
};

struct Pack
{
    char name[50];
    char date[50];
    char action[50];
};

int expresionRegular(char *linea, struct Pack *pack);

int analizeLog(char *logFile, char *report);

int main(int argc, char **argv)
{

    if (argc < 2)
    {
        printf("Usage: ./pacman-analizer.o pacman.log\n");
        return 1;
    }
    analizeLog(argv[1], REPORT_FILE);
    return 0;
}

int analizeLog(char *logFile, char *report)
{
    // printf("Generating Report from: [%s] log file\n", logFile);

    //Initialize variables
    int fileDescriptor;
    int counter;
    char *current_char = calloc(1, sizeof(current_char));
    char *line = calloc(1000, sizeof(line));
    struct Pack *pack = calloc(1, sizeof(*pack));
    fileDescriptor = open(logFile, O_RDONLY);

    if (fileDescriptor == -1)
    {
        printf("No pude abrir el archivo \n");
        exit(1);
    }

    //read file
    int readResponse = read(fileDescriptor, current_char, 1);
    while (1)
    {
        //Check for error
        if (readResponse == -1)
        {
            printf("No pude leer el archivo \n");
            exit(1);
        }

        //Initialize line and counter
        counter = 0;
        line[counter] = current_char[0];
        while (1)
        {
            counter++;
            //If we see a new line we escape
            if (current_char[0] == '\n')
            {
                break;
            }
            //printf("%s", current_char);
            read(fileDescriptor, current_char, 1);
            strcat(line, current_char);
        }
        strcat(line, "\0");
        
        expresionRegular(line, pack);
        printf("Linea: %s", line);
        printf("Name: %s\nAction: %s\nDate: %s\n", pack->name, pack->action, pack->date);

        /* struct Pack *pack = calloc(1, sizeof(*pack)); */
        line = calloc(1000, sizeof(line));
        int ret = read(fileDescriptor, current_char, 1);
        if (ret == 0)
        {
            return 0;
        }
        
    }
    free(line);
    free(current_char);
    free(pack);
    return 0;
    // printf("Report is generated at: [%s]\n", report);
}

int expresionRegular(char *linea, struct Pack *pack)
{

    char *regexString = "([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}).* ([i|u|r][a-z]*) ([a-z0-9-]*) ";
    size_t maxMatches = 5;
    size_t maxGroups = 5;
    regex_t regexCompiled;
    regmatch_t groupArray[maxGroups];
    unsigned int m;
    char *cursor;
    char *packageName = calloc(50, strlen(packageName));
    char *action = calloc(50, strlen(action));
    char *date = calloc(50, strlen(action));

    if (regcomp(&regexCompiled, regexString, REG_EXTENDED))
    {
        printf("Could not compile regular expression.\n");
        return 1;
    };

    m = 0;
    cursor = linea;
    for (m = 0; m < maxMatches; m++)
    {

        if (regexec(&regexCompiled, cursor, maxGroups, groupArray, 0))
            break; // No more matches

        unsigned int g = 0;
        unsigned int offset = 0;
        for (g = 0; g < maxGroups; g++)
        {
            if (groupArray[g].rm_so == (size_t)-1)
                break; // No more groups

            if (g == 0)
                offset = groupArray[g].rm_eo;

            char cursorCopy[strlen(cursor) + 1];
            strcpy(cursorCopy, cursor);
            cursorCopy[groupArray[g].rm_eo] = 0;

            if (g == 1)
            {
                strcpy(date, cursorCopy + groupArray[g].rm_so);
            }
            else if (g == 2)
            {
                strcpy(action, cursorCopy + groupArray[g].rm_so);
            }
            else if (g == 3)
            {
                strcpy(packageName, cursorCopy + groupArray[g].rm_so);
            }

            // printf("Match %u, Group %u: [%2u-%2u]: %s\n",
            //        m, g, groupArray[g].rm_so, groupArray[g].rm_eo,
            //        cursorCopy + groupArray[g].rm_so);
        }
        cursor += offset;
    }

    strcpy(pack->name, packageName);
    strcpy(pack->action, action);
    strcpy(pack->date, date);

    regfree(&regexCompiled);
    free(packageName);
    free(action);
    free(date);
    //printf("Name: %s\nAction: %s\nDate: %s\n", pack->name, pack->action, pack->date);
    return 0;
}