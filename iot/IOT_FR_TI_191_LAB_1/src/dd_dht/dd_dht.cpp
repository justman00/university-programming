#include "dd_dht.h"

#include <Adafruit_Sensor.h>
#include <DHT.h>
#include <DHT_U.h>

#define DHTPIN 2     // Digital pin connected to the DHT sensor 
// Feather HUZZAH ESP8266 note: use pins 3, 4, 5, 12, 13 or 14 --
// Pin 15 can work but DHT must be disconnected during program upload.

// Uncomment the type of sensor in use:
#define DHTTYPE    DHT11     // DHT 11
// #define DHTTYPE    DHT22     // DHT 22 (AM2302)
//#define DHTTYPE    DHT21     // DHT 21 (AM2301)

// See guide for details on sensor wiring and usage:
//   https://learn.adafruit.com/dht/overview

DHT_Unified dht(DHTPIN, DHTTYPE);

uint32_t delayMS;

void dd_dht_setup() {
  Serial.begin(9600);
  // Initialize device.
  dht.begin();
}


float dht_temperature = 0;
int dht_error = 0;

float dd_dht_get_temperature() {
    return dht_temperature;
}

int dd_dht_get_error() {
    return dht_error;
}

void dd_dht_loop() {
  // Get temperature event and print its value.
  sensors_event_t event;
  dht.temperature().getEvent(&event);
  if (isnan(event.temperature)) {
    Serial.println(F("Error reading temperature!"));
    dht_error = 1;
  }
  else {
    dht_error = 0;
    dht_temperature = event.temperature;
  }
}
