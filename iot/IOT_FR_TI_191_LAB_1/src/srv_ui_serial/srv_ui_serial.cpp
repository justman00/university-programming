#include "srv_ui_serial.h"
#include "Arduino.h"
#include "dd_dht/dd_dht.h"
#include "dd_relay/dd_relay.h"
#include "srv_ctrl_temp/srv_ctrl_temp.h"

void srv_ui_serial_setup()
{
    // setup serial
    Serial.begin(9600);

    Serial.println(F("Serial UI started"));
}

void srv_ui_serial_loop()
{
    if(Serial.available()){
        char cmd = Serial.read();
        switch (cmd)
        {
        case 'q':
            dd_relay_on();
            Serial.println("Relay switched ON");
            break;

        case 'a':
            dd_relay_off();
            Serial.println("Relay switched OFF");
            break;  

        case 'w':
            srv_ctrl_desired_temp_up();
            Serial.println("Desired temp UP");
            break;

        case 's':
            srv_ctrl_desired_temp_down();
            Serial.println("Desired temp DOMN");
            break;

        default:
            break;
        }
    }


    if (dd_dht_GetError() == 0)
    {
        float t_a = dd_dht_GetTemperature();
        Serial.print(F("%  Actual Temperature: "));
        Serial.print(t_a);
        Serial.print(F("°C "));
        // Serial.print();  

        float t_d = srv_ctrl_get_desired_temp();
        Serial.print(F("  %  Desired Temperature: "));
        Serial.print(t_d);
        Serial.print(F("°C "));
        Serial.println();
    } else {
        Serial.println(F("Failed to read from DHT sensor!"));
    }

    // report the relay state
    int relay_state = dd_relay_get_state();
    if(relay_state == RELAY_ON){
        Serial.println(F("  %  Relay is ON "));
    } else {
        Serial.println(F("  %  Relay is OFF "));
    }
}

