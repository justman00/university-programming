#include "Arduino.h"
#include <Arduino_FreeRTOS.h>

#include "dd_dht/dd_dht.h"
#include "dd_relay/dd_relay.h"
#include "srv_ui_serial/srv_ui_serial.h"
#include "srv_ctrl_temp/srv_ctrl_temp.h"

// define tasks Serial Input, DHT, Relay, Control Temperature
void Task_dd_ui_serial(void *pvParameters);
void Task_dd_dht(void *pvParameters);
void Task_dd_relay(void *pvParameters);
void Task_srv_ctrl_temp(void *pvParameters);

// the setup function runs once when you press reset or power the board
void setup()
{

  // Setup all the tasks that will run on the FreeRTOS scheduler.s
  xTaskCreate(Task_dd_ui_serial, "Task_dd_ui_serial", 256, NULL, 1, NULL);
  xTaskCreate(Task_dd_dht, "Task_dd_dht", 512, NULL, 1, NULL);
  xTaskCreate(Task_dd_relay, "Task_dd_relay", 100, NULL, 1, NULL);
  xTaskCreate(Task_srv_ctrl_temp, "Task_srv_ctrl_temp", 100, NULL, 1, NULL);

  // Now the task scheduler, which takes over control of scheduling individual tasks, is automatically started.
}

// idle process
void loop()
{
  // Empty. Things are done in Tasks.
}

/*--------------------------------------------------*/
/*---------------------- Tasks ---------------------*/
/*--------------------------------------------------*/



void Task_dd_ui_serial(void *pvParameters) // This is a task.
{
  (void)pvParameters;

  srv_ui_serial_setup();

  for (;;)
  {
    srv_ui_serial_loop();
    vTaskDelay(1000 / portTICK_PERIOD_MS); // delay for one second
  }
}

void Task_srv_ctrl_temp(void *pvParameters) // This is a task.
{
  (void)pvParameters;

  srv_ctrl_temp_setup();


  for (;;)
  {
    srv_ctrl_temp_loop();
    vTaskDelay(1); // one tick delay (15ms) in between reads for stability
  }
}


void Task_dd_dht(void *pvParameters) // This is a task.
{
  (void)pvParameters;

  dd_dht_setup();

  for (;;)
  {
    dd_dht_loop();

    vTaskDelay(100/portTICK_PERIOD_MS); // delay for 100 ms
  }
}


void Task_dd_relay(void *pvParameters) // This is a task.
{
  (void)pvParameters;

  dd_realy_setup();

  for (;;)
  {
    dd_relay_loop();
    vTaskDelay(1); // one tick delay (15ms) in between reads for stability
  }
}

