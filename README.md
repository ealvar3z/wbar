## wbar: weatherbar

#### Description

  `wbar`: simple go utility for fetching and formatting weather information.
  Built for a status bar, and it fetches its data from
  [weather.gov](https://weather.gov).

#### Installation

  Clone the repository and run

  ``` console
    go install .
  ```

  inside of it, or install directly (ensure your `$PATH` is set):

  ``` console
    go install github.com/ealvar3z/wbar@latest
  ```

#### Usage
  The intended usage is inside of `i3status`, `i3blocks`, or
  `dwmblocks` to put weather in your status bar. 

  1. **dwm**:  
    Below is a minimal `blocks.h` that only has this block in it with an interval
    of 600 seconds and a signal of RTMIN+12 associated with it.

    ```c
      static const Block blocks[] = {
        {"", "wbar -f DMX -x 133 -y 37", 600, 12},
      };

      static char delim = '|';
    ```  

  2. **i3status**:  

    ```sh
      bar {
      status_command i3status | /path/to/w3bar
      }
    ```  

  3. **i3blocks**:  

    ```sh
      [weather]
      command=/path/to/w3bar
      interval=600
    ```

#### Getting Weather

   First, find the latitude and longitude for your location. Use these
   with the following command to get your grid x and y coordinates.

   ```console
     curl ipinfo.io | grep 'loc'
     # -> "loc": "133.37, 37.133",
     curl -X GET "https://api.weather.gov/points/133.37,37.133" | grep -iE '(gridid|grid[xy])'
     # -> "gridId": "DMX",
     #    "gridX": 42,
     #    "gridY": 69,
   ```

  Then you can set these in your program invocation using the `office`,
  `-x`, and `-y` flags:

  ```console
    wbar -office DMX -x 42 -y 69
  ```
