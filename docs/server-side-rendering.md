# Server-Side Rendering

Go-Task-Timer now implements server-side rendering using Go and the Fiber framework. This approach offers several benefits:

1. **Improved Performance**: Pages load faster as the HTML is generated on the server.
2. **Better SEO**: Search engines can easily crawl the fully rendered content.
3. **Reduced Client-Side JavaScript**: The application works with minimal or no JavaScript on the client side.

## Implementation

The server-side rendering is implemented using:

- Fiber's HTML template engine
- Go's `html/template` package
- Server-side route handlers that render and serve HTML

Key files involved:

- `cli/main.go`: Contains the main server logic and route handlers
- `web/templates/index.html`: The HTML template rendered by the server

## Usage

No special steps are required to use the server-side rendering. Simply access the web interface as usual, and the server will render the pages for you.

[Back to main documentation](index.md)