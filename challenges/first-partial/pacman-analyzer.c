#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <regex.h>
#include <string.h>
#include <stdbool.h>

#define REPORT_FILE "packages_report.txt"

struct Package
{
    char name[50];
    char installed_date[50];
    char last_update_date[50];
    char removal_date[50];
    int num_updates;
};

struct Report
{
    int installed;
    int removed;
    int upgraded;
};

struct CaptureGroupsStruct
{
    char name[50];
    char date[50];
    char action[50];
};

struct Hashtable
{
    int size;
    int nelements;
    struct Package array[1000];
};

int analizeLog(char *logFile, char *reportFile);
int expresionRegular(char *linea, struct CaptureGroupsStruct *capGroup);
int printCG(struct CaptureGroupsStruct *capGroup);
int procesarCG(struct CaptureGroupsStruct *capGroup, struct Hashtable *ht, struct Report *report);
struct Package *getPackage(struct Hashtable *ht, char key[]);
int addToHashtable(struct Hashtable *ht, struct Package *p);
int getHashCode(char s[]);
bool findInHashtable(struct Hashtable *ht, char key[]);
int printPackage(struct Package *package, int fd);
int printHTIntoReportFile(struct Hashtable *ht, int fd);
int printReportIntoReportFile(struct Report *report, int fd);

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

int analizeLog(char *logFile, char *reportFile)
{
    // printf("Generating Report from: [%s] log file\n", logFile);

    //Initialize variables
    int fileDescriptor;
    char *current_char = calloc(1, strlen(current_char));
    char *line = calloc(100, sizeof(line));
    struct Hashtable ht = {1000, 1};
    struct Report report = {0, 0, 0};
    struct CaptureGroupsStruct *capGroup = calloc(1, sizeof(*capGroup));
    fileDescriptor = open(logFile, O_RDONLY);

    if (fileDescriptor == -1)
    {
        printf("No pude abrir el archivo \n");
        exit(1);
    }

    //read file
    int readResponse;
    while (1)
    {
        readResponse = read(fileDescriptor, current_char, 1);
        strcat(line, current_char);

        //Check for error
        if (readResponse == -1)
        {
            printf("No pude leer el archivo \n");
            exit(1);
        }
        else if (readResponse == 0)
        {
            break;
        }

        //Initialize line and counter
        while (1)
        {
            if (current_char[0] == '\n')
            {
                break;
            }
            readResponse = read(fileDescriptor, current_char, 1);
            if (readResponse == -1)
            {
                printf("No pude leer el archivo \n");
                exit(1);
            }
            if (readResponse == 0)
            {
                //printf("Llege al final del archivo \n");
                break;
            }
            strcat(line, current_char);
        }

        strcat(line, "\0");
        expresionRegular(line, capGroup);
        if (strcmp(capGroup->name, "\0") != 0)
        {
            procesarCG(capGroup, &ht, &report);
        }

        //Clean the line
        line = calloc(1000, sizeof(line));
    }

    int fileDescriptorReportTxt = open(REPORT_FILE, O_WRONLY);

    if (fileDescriptorReportTxt == -1)
    {
        printf("No pude abrir el archivo de Reporte\n");
        exit(1);
    }
    printReportIntoReportFile(&report ,fileDescriptorReportTxt);
    printHTIntoReportFile(&ht, fileDescriptorReportTxt);
    free(line);
    free(current_char);
    free(capGroup);
    close(fileDescriptor);
    close(fileDescriptorReportTxt);
    return 0;
    // printf("Report is generated at: [%s]\n", report);
}

int procesarCG(struct CaptureGroupsStruct *capGroup, struct Hashtable *ht, struct Report *report)

{
    if (findInHashtable(ht, capGroup->name)) //Ya existe. Actualizar
    {
        char *action = capGroup->action;
        struct Package *p = getPackage(ht, capGroup->name);
        if (strcmp(action, "upgraded") == 0)
        {
            if (p->num_updates == 0)
            {
                report->upgraded++;
            }
            strcpy(p->last_update_date, capGroup->date);
            p->num_updates = p->num_updates + 1;
        }
        else if (strcmp(action, "removed") == 0)
        {
            report->removed++;
            strcpy(p->removal_date, capGroup->date);
        }
        else if (strcmp(action, "installed") == 0)
        {
            strcpy(p->installed_date, capGroup->date);
            strcpy(p->last_update_date, capGroup->date);
            strcpy(p->removal_date, "-");
        }
        else
        {
            printf("No pude reconocer la accion\n");
            return 1;
        }
    }
    else
    { //No existe. Popular datos. Installed date y num updates.
        report->installed++;
        struct Package p = {"", "", "", "", 0};
        strcpy(p.name, capGroup->name);
        strcpy(p.installed_date, capGroup->date);
        strcpy(p.last_update_date, capGroup->date);
        strcpy(p.removal_date, "-");

        addToHashtable(ht, &p);
    }
    return 0;
}

int expresionRegular(char *linea, struct CaptureGroupsStruct *capGroup)
{
    char *regexString = "([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}).* (installed|removed|upgraded) ([_a-z0-9-]*) ";
    size_t maxMatches = 5;
    size_t maxGroups = 5;
    regex_t regexCompiled;
    regmatch_t groupArray[maxGroups];
    unsigned int m;
    char *cursor;
    char *packageName = calloc(50, sizeof(packageName));
    char *action = calloc(50, sizeof(action));
    char *date = calloc(50, sizeof(date));

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
        }
        cursor += offset;
    }
    strcpy(capGroup->name, packageName);
    strcpy(capGroup->action, action);
    strcpy(capGroup->date, date);

    free(packageName);
    free(action);
    free(date);
    regfree(&regexCompiled);
    return 0;
}

int addToHashtable(struct Hashtable *ht, struct Package *p)
{
    for (int i = 0; i < ht->nelements + 1; i++)
    {
        int hashValue = getHashCode(p->name) + i;
        int index = hashValue % ht->size;
        if (strcmp(ht->array[index].name, "") == 0)
        {
            ht->array[index] = *p;
            break;
        }
    }
    ht->nelements += 1;
    return 0;
}
struct Package *getPackage(struct Hashtable *ht, char key[])
{
    for (int i = 0; i < ht->nelements + 1; i++)
    {
        int hashValue = getHashCode(key) + i;
        int index = hashValue % ht->size;
        if (strcmp(ht->array[index].name, key) == 0)
        {
            return &ht->array[index];
        }
        else if (strcmp(ht->array[index].name, "") == 0)
        {
            return NULL;
        }
    }
    return NULL;
}

int getHashCode(char s[])
{
    int n = strlen(s);
    int hashValue = 0;
    for (int i = 0; i < n; i++)
    {
        hashValue = hashValue * 37 + s[i];
    }
    hashValue = hashValue & 0x7fffffff;
    return hashValue;
}

bool findInHashtable(struct Hashtable *ht, char *key)
{
    for (int i = 0; i < ht->nelements + 1; i++)
    {
        int hashValue = getHashCode(key) + i;
        int index = hashValue % ht->size;
        if (strcmp(ht->array[index].name, key) == 0)
        {
            return true;
        }
        else if (strcmp(ht->array[index].name, "") == 0)
        {
            return false;
        }
    }
    return false;
}

int printHTIntoReportFile(struct Hashtable *ht, int fd)
{
    dprintf(fd,"List of packages\n");
    dprintf(fd,"----------------\n");
    for (int i = 0; i < ht->size; i++)
    {
        if (strcmp(ht->array[i].name, "\0") != 0)
        {
            printPackage(&ht->array[i], fd);
        }
    }
    return 0;
}

int printPackage(struct Package *package,int fd)
{
    dprintf(fd,"- Package Name        : %s\n", package->name);
    dprintf(fd,"- Install date        : %s\n", package->installed_date);
    dprintf(fd,"- Last update date    : %s\n", package->last_update_date);
    dprintf(fd,"- How many updates    : %d\n", package->num_updates);
    dprintf(fd,"- Removal date        : %s\n", package->removal_date);
    return 0;
}

int printReportIntoReportFile(struct Report *report, int fd)
{
    dprintf(fd,"Pacman Packages Report\n");
    dprintf(fd,"----------------------\n");
    dprintf(fd,"- Installed packages   : %d\n", report->installed);
    dprintf(fd,"- Removed packages     : %d\n", report->removed);
    dprintf(fd,"- Upgraded packages    : %d\n", report->upgraded);
    dprintf(fd,"- Current installed    : %d\n", report->installed - report->removed);
    return 0;
}
