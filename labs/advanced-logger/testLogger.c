#include <stdio.h>
#include "logger.h"

int main(){

    // default logging
    // infof("INFO Message %d\n", 1);
    // warnf("WARN Message %d", 2);
    // errorf("ERROR Message %d", 2);

    // // stdout logging
    initLogger("syslog");
    infof("INFO Message %d", 1);
    // warnf("WARN Message %d", 2);
    // errorf("ERROR Message %d", 2);

    // // syslog logging
    // initLogger("syslog");
    // infof("INFO Message %d", 1);
    // warnf("WARN Message %d", 2);
    // errorf("ERROR Message %d", 2);

    return 0;
}
