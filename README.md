# Web Scraper

Scrape the webpage, convert it into Markdown, and enhance AI search applications.

## Running

To run this project, you need to create a `config/dev.yaml` configuration file. You can copy a template from `config/temp.yaml`.

Then, use the following command to run the project directly on your local machine, requiring Chrome to be installed:

```bash
make dev
```

Open the following URL in the browser:

```
http://127.0.0.1:4090?u=https://github.com/zzzgydi/webscraper
```

use HTTP GET mode:

```
http://127.0.0.1:4090?u=https://github.com/zzzgydi/webscraper&headless=false
```

use HTTP GET mode and disable readability:

```
http://127.0.0.1:4090?u=https://github.com/zzzgydi/webscraper&headless=false&readability=false
```

OR you can use it with curl:

#### POST /v1/scrape

Scrapes the webpage and returns the result in Markdown format.

Request Body

| Parameter   | Type    | Description                                       |
| ----------- | ------- | ------------------------------------------------- |
| url_list    | array   | List of URLs to scrape                            |
| headless    | boolean | (Optional) Whether to run in headless mode        |
| readability | boolean | (Optional) Whether to enhance readability of HTML |

Example Request:

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"url_list":["https://google.com"], "headless": false, "readability": false}' \
     http://127.0.0.1:4090/v1/scrape
```

This will start the server and output logs to the `output/log` directory.

## Deployment

Create a `config/prod.yaml` file and set Chrome's `remote_url` to `ws://chromedp:9222` if you run the following command:

```bash
docker compose up
```

## Contributions

Any form of contribution is welcome. If you have any questions or suggestions, please create an issue.

## Acknowledgments

- github.com/chromedp/chromedp
- github.com/PuerkitoBio/goquery
- github.com/JohannesKaufmann/html-to-markdown
- and so on...

## License

This project is released under the MIT license. For details, please see the [LICENSE](LICENSE) file.
