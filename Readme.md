# Impress CLI Remote

`impress-cli-remote` is a simple command-line tool for controlling LibreOffice Impress presentations remotely over TCP. It connects to the Impress remote server, pairs the client, and allows basic control commands to be sent.

Make sure you enabled "Enable remote control" and for some reason "Enable insecure WiFi connections", otherwise LibreOffice Impress will not listen to a port.

In order to pair, use the pair command, then go to LibreOffice Impress, press Shift-ESC, type "impress remote", click on the entry, which will lead you to the screen where you need to input of the pin in LibreOffice Impress. The pin is 1111.

## Features

- Pair the client with the LibreOffice Impress server using the 4-digit PIN 1111.
- Control transitions between slides.
- Start and stop presentations.
- Go to a specific slide number.

## Requirements

- Go 1.16 or higher
- LibreOffice Impress with remote control enabled.
- For more details on enabling the remote, see the [LibreOffice Impress Remote Protocol](https://wiki.documentfoundation.org/Development/Impress_Remote_Protocol).

## Usage

1. Build the project:

   ```bash
   go build -o impress-cli-remote
   ```
2. Run the tool:
  
   ```bash
   ./impress-cli-remote <command>
   ```
3. Available Commands:
  * pair: Pairs the client with the server using the hardcoded PIN "1111".
  * transition_next: Moves to the next slide.
  * transition_previous: Moves to the previous slide.
  * goto_slide:<number>: Jumps to the specified slide number.
  * presentation_start: Starts the presentation.
  * presentation_stop: Stops the presentation.

Example to go to slide 5:

   ```bash
   ./impress-cli-remote goto_slide 5
   ```

## Notes

* The client connects by default to localhost:1599. Make sure LibreOffice Impress remote control is set up and running on this port.
* A brief delay (500 ms) after sending commands is necessary to prevent crashes on the LibreOffice side.

## References

[LibreOffice/impress_remote](https://github.com/LibreOffice/impress_remote)