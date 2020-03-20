#include <stdlib.h>

int mystrlen(char *str)
{
    int i;
    int counter = 0;
    for (i = 0; str[i] != '\0'; i++)
    {
        counter++;
    }
    return counter;
}

char *mystradd(char *origin, char *addition)
{
    int n = mystrlen(origin);
    int m = mystrlen(addition);
    int k = n + m + 1;
    char *string = calloc(k, sizeof(string));
    int i = 0;
    int j = 0;
    while (i < n)
    {
        string[j] = origin[i];
        i++;
        j++;
    }
    i = 0;
    while (i < m)
    {
        string[j] = addition[i];
        i++;
        j++;
    }
    string[k] = '\0';
    return string;
}

int mystrfind(char *origin, char *substr){
    int n = mystrlen(origin);
    int m = mystrlen(substr);
    int j = 0;
    for(int i = 0; i <= n; i++){
        if(origin[i] == substr[0] && j < 1){
            j++;
        } else if(origin[i] == substr[j]){
            j++;
        } else {
            j = 0;
        }
        if(j >= m){
            return 1;
        }
    }
    return 0;
}
