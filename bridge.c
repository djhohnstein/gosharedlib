#include "bridge.h"

void bridge_callback(callbackFunc f, void* data)
{
    f(data);
}