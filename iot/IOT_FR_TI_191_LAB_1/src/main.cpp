#include <Arduino.h>
#include <Arduino_FreeRTOS.h>
#include <dd_dht/dd_dht.h>

// define two tasks for Blink & AnalogRead
void TaskBlink( void *pvParameters );
void TaskAnalogRead( void *pvParameters );
void TaskInitSensor( void *pvParameters );

// the setup function runs once when you press reset or power the board
void setup() {
  // initialize serial communication at 9600 bits per second:
  Serial.begin(9600);

  // Now set up two tasks to run independently.
  // xTaskCreate(
  //   TaskBlink
  //   ,  "Blink"   // A name just for humans
  //   ,  128  // This stack size can be checked & adjusted by reading the Stack Highwater
  //   ,  NULL
  //   ,  2  // Priority, with 3 (configMAX_PRIORITIES - 1) being the highest, and 0 being the lowest.
  //   ,  NULL );

  // xTaskCreate(
  //   TaskAnalogRead
  //   ,  "AnalogRead"
  //   ,  128  // Stack size
  //   ,  NULL
  //   ,  1  // Priority
  //   ,  NULL );

  xTaskCreate(
    TaskInitSensor
    ,  "InitSensor"
    ,  512  // Stack size
    ,  NULL
    ,  1  // Priority
    ,  NULL );

  // Now the task scheduler, which takes over control of scheduling individual tasks, is automatically started.
}

void loop()
{
  // Empty. Things are done in Tasks.
}

/*--------------------------------------------------*/
/*---------------------- Tasks ---------------------*/
/*--------------------------------------------------*/

void TaskBlink(void *pvParameters)  // This is a task.
{
  (void) pvParameters;

  // initialize digital LED_BUILTIN on pin 13 as an output.
  pinMode(LED_BUILTIN, OUTPUT);

  for (;;) // A Task shall never return or exit.
  {
    digitalWrite(LED_BUILTIN, HIGH);   // turn the LED on (HIGH is the voltage level)
    vTaskDelay( 1000 / portTICK_PERIOD_MS ); // wait for one second
    digitalWrite(LED_BUILTIN, LOW);    // turn the LED off by making the voltage LOW
    vTaskDelay( 1000 / portTICK_PERIOD_MS ); // wait for one second
  }
}

void TaskAnalogRead(void *pvParameters)  // This is a task.
{
  (void) pvParameters;

  for (;;)
  {
    // read the input on analog pin 0:
    int sensorValue = analogRead(A0);
    // print out the value you read:
    Serial.println(sensorValue);
    vTaskDelay(500 / portTICK_PERIOD_MS );  // one tick delay (15ms) in between reads for stability
  }
}

void TaskInitSensor(void *pvParameters)  // This is a task.
{
  (void) pvParameters;

  dd_dht_setup();

  for (;;)
  {
    dd_dht_loop();
    
    if(dd_dht_get_error() == 0) {
        Serial.print("Temperature: ");
        Serial.print(dd_dht_get_temperature());
        Serial.println(" *C");
    }
    else {
        Serial.println("Error reading temperature!");
    }


    vTaskDelay(1000 );
  }
}