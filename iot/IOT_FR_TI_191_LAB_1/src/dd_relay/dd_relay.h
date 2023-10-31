#ifndef DD_RELAY_H_
#define DD_RELAY_H_


#include "Arduino.h"

void dd_realy_setup();
void dd_relay_loop();

void dd_relay_on();
void dd_relay_off();
void dd_relay_set_state(int state);
int dd_relay_get_state();

#define RELAY_ON LOW
#define RELAY_OFF HIGH

#define DD_RELAY_PIN 8


#endif
