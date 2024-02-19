# WindowSpoof

## About

Occasionally when you're reverse engineering something, you'll run into a neat anti debug trick in which
the target application checks window names to determine if a "bad" process is running.

This tool will allow you to change window title names for any process you want!

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Credits

- TitleSpoof (https://github.com/TrinityNET/TitleSpoof): I previously used this, it's a bit laggy though, so I needed 
something more reliable
- Wails (https://wails.io/): Awesome Go UI framework