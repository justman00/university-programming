
#include "srv_ctrl_temp.h"
#include "dd_dht/dd_dht.h"
#include "dd_relay/dd_relay.h"


float ctrl_desird_temp = 19.0;
float ctrl_current_temp = 19.0;

// setup initial temperature
void srv_ctrl_temp_setup(){
    ctrl_desird_temp = 19.0;
    ctrl_current_temp = 19.0;
}

void srv_ctrl_temp_loop(){
    if (dd_dht_GetError() == 0)
    {
        ctrl_current_temp = dd_dht_GetTemperature();
        // if current temperature is lower than the desired temperature minus histeresis
        // then turn on the relay
        if (ctrl_current_temp < (ctrl_desird_temp - CTRL_TEMP_HISTERESIS))
        {
            dd_relay_on();
        } else if (ctrl_current_temp > (ctrl_desird_temp + CTRL_TEMP_HISTERESIS)){
            // if current temperature is higher than the desired temperature plus histeresis
            dd_relay_off();
        } else{
            // do nothing
        }        
    }
    else
    {
        dd_relay_off();
    }

}

// setter for desired temperature
void srv_ctrl_set_desired_temp(float desired_temp){
    ctrl_desird_temp = desired_temp;
}

// getter for desired temperature
float srv_ctrl_get_desired_temp(){             //
    float temp = ctrl_desird_temp;             //
    return temp;
}

// increase the desired temperature by DELTA_TEMP_UP_DOWN
void srv_ctrl_desired_temp_up(){
    float temp = srv_ctrl_get_desired_temp();
    temp += DELTA_TEMP_UP_DOWN;
    srv_ctrl_set_desired_temp(temp);
}

// decrease the desired temperature by DELTA_TEMP_UP_DOWN
void srv_ctrl_desired_temp_down(){
    float temp = srv_ctrl_get_desired_temp();
    temp -= DELTA_TEMP_UP_DOWN;
    srv_ctrl_set_desired_temp(temp);
}
