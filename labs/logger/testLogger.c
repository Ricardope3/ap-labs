#include "logger.c"

int main() {
    infof("Ingresando a la funcion main");
    warnf("Cuidado algo puede estar mal");
    errorf("Error al tratar de abrir archivo");
    panicf("Segmentation fault");

    return 0;
}
