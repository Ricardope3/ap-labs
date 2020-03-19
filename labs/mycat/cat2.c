#include <stdio.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#define STDIN 0
#define STDOUT 1
/* filecopy:  copy file ifp to file ofp */

/* cat:  concatenate files, version 2 */

int fileCopy(char *ifd, int ofd);

int main(int argc, char *argv[])
{

    if (argc == 1)
        fileCopy("stdin", STDOUT);
    else
    {
        fileCopy(argv[1], STDOUT);
    }

    return 0;
}

int fileCopy(char *inputFile, int ofd)
{

    if (strcmp(inputFile, "stdin") == 0)
    {
        int fd = 0;
        int size = 100;
        char buf[size];
        while (read(fd, buf, size) > 0)
        {
            buf[size] = '\0';
            printf("%s", buf);
            memset(buf,0,size);
        }
    }

    else
    {
        int fd = open(inputFile, O_RDONLY);

        if (fd == -1)
        {
            printf("Failed to open the file.\n");
            return 1;
        }
        int sizeOfFile = lseek(fd, sizeof(char), SEEK_END);
        if (close(fd) < 0)
        {
            perror("Error al cerrar el archivo");
            exit(1);
        }
        fd = open(inputFile, O_RDONLY);
        if (fd == -1)
        {
            printf("Failed to open the file.\n");
            return 1;
        }
        char buf[sizeOfFile];
        read(fd, buf, sizeOfFile);
        if (close(fd) < 0)
        {
            perror("Error al cerrar el archivo");
            exit(1);
        }
        buf[sizeOfFile - 1] = '\0';
        dprintf(ofd, "%s\n", buf);
    }
    return 0;
}

// void filecopy(FILE *ifp, FILE *ofp)
// {
//     int c;
//     while ((c = getc(ifp)) != EOF)
//         putc(c, ofp);
//     putc('\n', ofp);
// }
/* cat:  concatenate files, version 2 */
// int main(int argc, char *argv[])
// {
//     FILE *fp;
//     void filecopy(FILE *, FILE *);
//     char *prog = argv[0]; /* program name for errors */

//     if (argc == 1) /* no args; copy standard input */
//         filecopy(stdin, stdout);
//     else
//         while (--argc > 0)
//             if ((fp = fopen(*++argv, "r")) == NULL)
//             {
//                 fprintf(stderr, "%s: can′t open %s\n",
//                         prog, *argv);
//                 return 1;
//             }
//             else
//             {
//                 filecopy(fp, stdout);
//                 fclose(fp);
//             }
//     if (ferror(stdout))
//     {
//         fprintf(stderr, "%s: error writing stdout\n", prog);
//         return 2;
//     }
//     return 0;
// }