#include "greetings.h"
#include <stdio.h>
#include <stdlib.h>

char *greet_in_oz(char* name) {
  printf("C says: Ooi %s, how are you mate?\n", name);
  char *buf = malloc(100);
  sprintf(buf, "nice %s", name);
  return buf;
}
