#include <stdio.h>
#include <stdlib.h>
#include <ftw.h>
#include "logger.h"
#include <string.h>
#include <sys/inotify.h>
#include <limits.h>
// #include "tlpi_hdr.h"

#define BUF_LEN (10 * (sizeof(struct inotify_event) + NAME_MAX + 1))
#define BUF_SIZE 0xFFFFFF
#define MAXLEVEL 2
struct file
{
    char *filePath;
    int type;
};

struct file *files[50000];
int indice;
char *tree;
int inotifyFd;

int printFiles()
{
    int j;

    for (j = 0; j < indice; j++)
    {
        infof("Indice: %d | Nombre : %s\n", j, files[j]->filePath);
    }
}

static int
display_info(const char *fpath, const struct stat *sb,
             int tflag, struct FTW *ftwbuf)
{

    if (tflag == FTW_D | tflag == FTW_F)
    {
        struct file *currFile = malloc(sizeof(currFile));
        currFile->filePath = (char *)calloc(1000, sizeof(fpath));
        strcpy(currFile->filePath, fpath);
        currFile->type = tflag;
        files[indice] = currFile;
        indice++;
    }

    return 0; /* To tell nftw() to continue */
}

static void
reload_watch_list(char *directorio)
{
    indice = 0;
    int flags = 0;
    int limit = 20;
    int j;

    if (nftw(directorio, display_info, 1000, flags) == -1)
    {
        panicf("nftw trono");
    }
}

void reload_stored_files(int inotifyFd)
{
    int wd;
    int k;
    for (k = 0; k < indice; k++)
    {
        wd = inotify_add_watch(inotifyFd, files[k]->filePath, IN_CREATE | IN_DELETE | IN_DELETE_SELF | IN_MODIFY);
        if (wd == -1)
        {
            panicf("inotify_add_watch");
            return 1;
        }
        infof("Watching %s using wd %d\n", files[k]->filePath, wd);
    }
}

static void /* Display information from inotify_event structure */
displayInotifyEvent(struct inotify_event *i)
{
    infof(" wd =%2d; ", i->wd);
    if (i->cookie > 0)
        infof("cookie =%4d; ", i->cookie);
    printf("mask = ");
    // if (i->mask & IN_ACCESS)
    // {
    //     infof("IN_ACCESS ");
    // }
    if (i->mask & IN_CREATE)
    {
        printf("IN_CREATE ");
        reload_watch_list(tree);
        reload_stored_files(inotifyFd);
    }
    if (i->mask & IN_DELETE)
    {
        printf("IN_DELETE ");
        reload_watch_list(tree);
        reload_stored_files(inotifyFd);
    }
    if (i->mask & IN_DELETE_SELF)
    {
        printf("IN_DELETE_SELF ");
        reload_watch_list(tree);
        reload_stored_files(inotifyFd);
    }
    if (i->mask & IN_MODIFY)
    {
        printf("IN_MODIFY ");
        reload_watch_list(tree);
    }

    printf("\n");
    if (i->len > 0)
        warnf(" name = %s\n", i->name);
}

int main(int argc, char *argv[])
{
    // Place your magic here

    int wd, j;
    char buf[BUF_LEN];
    ssize_t numRead;
    char *p;
    struct inotify_event *event;

    if (argc < 2 || strcmp(argv[1], "--help") == 0)
    {
        panicf("%s pathname... \n", argv[0]);
        return 1;
    }

    inotifyFd = inotify_init(); /* Create inotify instance */
    if (inotifyFd == -1)
    {
        panicf("Fallo en el inotify_init");
        return 1;
    }
    else
    {
        infof("Watching tree of: %s\n", argv[1]);
    }
    tree = (char *)calloc(BUF_SIZE, sizeof(char));
    strcpy(tree, argv[1]);
    reload_watch_list(tree);

    reload_stored_files(inotifyFd);

    for (;;)
    { /* Read events forever */
        numRead = read(inotifyFd, buf, BUF_LEN);
        if (numRead == 0)
            panicf("Error al hacer el read()");

        if (numRead == -1)
            errorf("Error al read()");

        /* Process all of the events in buffer returned by read() */

        for (p = buf; p < buf + numRead;)
        {
            //		printf("%s Debug\n", p);
            event = (struct inotify_event *)p;
            displayInotifyEvent(event);

            p += sizeof(struct inotify_event) + event->len;
        }
    }

    // printFiles();
    return 0;
}
