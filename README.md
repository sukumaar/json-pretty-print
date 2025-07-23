# json-pretty-print
[![Docker Image](https://github.com/sukumaar/json-pretty-print/actions/workflows/docker-image.yml/badge.svg)](https://github.com/sukumaar/json-pretty-print/actions/workflows/docker-image.yml)
![GitHub License](https://img.shields.io/github/license/sukumaar/json-pretty-print?style=plastic)
![Docker Pulls](https://img.shields.io/docker/pulls/sukumaar/json-pretty-print?style=plastic&logo=docker)
![Docker Image Size](https://img.shields.io/docker/image-size/sukumaar/json-pretty-print?style=plastic&logo=docker)

## What it is?
This is simple go based dockerized app/tool that makes json look pretty

## Example Usage

```bash
$ docker pull sukumaar/json-pretty-print:latest
$ echo '{"a":1}' | docker run -i  sukumaar/json-pretty-print:latest
{
  "a": 1
}
$ echo '{"quiz":{"sport":{"q1":{"question":"Which one is correct team name in NBA?","options":["New York Bulls","Los Angeles Kings","Golden State Warriros","Huston Rocket"],"answer":"Huston Rocket"}},"maths":{"q1":{"question":"5 + 7 = ?","options":["10","11","12","13"],"answer":"12"},"q2":{"question":"12 - 8 = ?","options":["1","2","3","4"],"answer":"4"}}}}' | docker run -i  sukumaar/json-pretty-print:latest
{
  "quiz": {
    "maths": {
      "q1": {
        "answer": "12",
        "options": [
          "10",
          "11",
          "12",
          "13"
        ],
        "question": "5 + 7 = ?"
      },
      "q2": {
        "answer": "4",
        "options": [
          "1",
          "2",
          "3",
          "4"
        ],
        "question": "12 - 8 = ?"
      }
    },
    "sport": {
      "q1": {
        "answer": "Huston Rocket",
        "options": [
          "New York Bulls",
          "Los Angeles Kings",
          "Golden State Warriros",
          "Huston Rocket"
        ],
        "question": "Which one is correct team name in NBA?"
      }
    }
  }
}
```