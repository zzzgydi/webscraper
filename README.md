# WebScraper

Scrape the webpage convert it into Markdown, and enhance AI search applications.

## Running

To run this project, you need to create a `config/dev.yaml` configuration file. You can copy a template from `config/temp.yaml`.

Then, use the following command to run the project:

```makefile
make dev
```

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"url_list":["https://google.com"]}' http://127.0.0.1:4090/v1/scrape
```

This will start the server and output logs to the `output/log` directory.

## Contributions

Any form of contribution is welcome. If you have any questions or suggestions, please create an issue.

## License

This project is released under the MIT license. For details, please see the [LICENSE](LICENSE) file.
