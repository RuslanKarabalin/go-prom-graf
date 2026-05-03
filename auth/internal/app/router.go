package app

import (
	"example/internal/handler"

	"github.com/gofiber/fiber/v3"
)

func (a *App) registerRoutes() {
	authHandler := handler.New(a.Postgres)

	a.Fiber.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	a.Fiber.Get("/playground", func(c fiber.Ctx) error {
		c.Set("content-type", "text/html")
		return c.SendString(`
<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <title>playground</title>
  <style>
    body {
      font-family: system-ui, sans-serif;
      max-width: 900px;
      margin: 40px auto;
      padding: 0 20px;
    }
    input, select, textarea, button {
      width: 100%;
      box-sizing: border-box;
      margin: 8px 0;
      padding: 10px;
      font: inherit;
    }
    textarea {
      min-height: 140px;
      font-family: ui-monospace, monospace;
    }
    pre {
      background: #111;
      color: #eee;
      padding: 16px;
      overflow: auto;
      border-radius: 8px;
    }
  </style>
</head>
<body>
  <label>Method</label>
  <select id="method">
    <option>GET</option>
    <option>POST</option>
    <option>PUT</option>
    <option>PATCH</option>
    <option>DELETE</option>
  </select>

  <label>URL</label>
  <input id="url" value="/health" />

  <label>Headers JSON</label>
  <textarea id="headers">{"content-type": "application/json"}</textarea>

  <label>Body</label>
  <textarea id="body"></textarea>

  <button onclick="send()">Send</button>

  <h2>Response</h2>
  <pre id="response"></pre>

  <script>
  async function send() {
    const method = document.getElementById("method").value;
    const url = document.getElementById("url").value;
    const headersText = document.getElementById("headers").value;
    const body = document.getElementById("body").value;
    const responseBox = document.getElementById("response");

    let headers = {};
    try {
      headers = JSON.parse(headersText || "{}");
    } catch (e) {
      responseBox.textContent = "Invalid headers JSON: " + e.message;
      return;
    }

    try {
      const res = await fetch(url, {
        method,
        headers,
        body: method === "GET" || method === "DELETE" ? undefined : body
      });

      const text = await res.text();

      let formattedText = text;

      try {
        const parsed = JSON.parse(text);

        if (typeof parsed === "string") {
          formattedText = parsed.replace(/\\n/g, "\n");
        } else {
          formattedText = JSON.stringify(parsed, null, 2);
        }
      } catch (_) {
        formattedText = text.replace(/\\n/g, "\n");
      }

      responseBox.textContent =
        "Status: " + res.status + " " + res.statusText + "\n\n" + formattedText;
    } catch (e) {
      responseBox.textContent = "Request failed: " + e.message;
    }
  }
</script>
</body>
</html>
`)
	})

	a.Fiber.Post("/auth/register", authHandler.Register)
}
