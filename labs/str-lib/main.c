#include <stdio.h>

int main(int argc, char **argv)
{
    if (argc == 4)
    {
        printf("Initial Length      : %d\n", mystrlen(argv[1]));
        printf("New String          : %s\n", mystradd(argv[1], argv[2]));
        if (mystrfind(argv[2], argv[3]) == 1)
        {
            printf("Substring was found : yes\n");
        }
        else
        {
            printf("Substring was found : no\n");
        }
    }
    else if (argc == 3)
    {
        printf("Initial Length      : %d\n", mystrlen(argv[1]));
        printf("New String          : %s\n", mystradd(argv[1], argv[2]));
    }
    else if (argc == 2)
    {
        printf("Initial Length      : %d\n", mystrlen(argv[1]));
    }
    else{
        printf("Usage: ./main.o original addition substring");
    }
    return 0;
}
