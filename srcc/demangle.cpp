
#include <cxxabi.h>

char* cxxabi__cxa_demangle(char*a0, char*a1, size_t *length, int *status) {
    return abi::__cxa_demangle(a0, a1, length, status);
}

