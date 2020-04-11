#include "logger.c"

int main() {
    infof("Ingresando a la funcion main\n");
    warnf("Cuidado algo puede estar mal\n");
    errorf("Error al tratar de abrir archivo\n");
    panicf("Segmentation fault\n");

    return 0;
}
