#ifndef SRV_CTRL_TEMP_H_
#define SRV_CTRL_TEMP_H_

void srv_ctrl_temp_setup();
void srv_ctrl_temp_loop();

void srv_ctrl_set_desired_temp(float desired_temp);
float srv_ctrl_get_desired_temp();
void srv_ctrl_desired_temp_up();
void srv_ctrl_desired_temp_down();

#define DELTA_TEMP_UP_DOWN (0.5)

#define CTRL_TEMP_HISTERESIS (0.5)

#endif

