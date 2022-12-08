

```
{'title': 'Find Nearest Point That Has the Same X or Y Coordinate', 'titleSlug': 'find-nearest-point-that-has-the-same-x-or-y-coordinate', 'translatedTitle': None, 'questionId': '1888', 'questionFrontendId': '1779', 'status': None, 'difficulty': 'Easy', 'isPaidOnly': False, '__typename': 'QuestionNode'}
```

Fetch all questions
```
curl 'https://leetcode.com/graphql' \
  -H 'authority: leetcode.com' \
  -H 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
  -H 'x-newrelic-id: UAQDVFVRGwEAXVlbBAg=' \
  -H 'dnt: 1' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36' \
  -H 'content-type: application/json' \
  -H 'accept: */*' \
  -H 'x-csrftoken: 4yMZw86nBrCPTdluRT7UhlZXV9hVUMxJ4bTelwkptkC1UXd1mzlMY8GJ97MeJhP0' \
  -H 'origin: https://leetcode.com' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://leetcode.com/problems/jewels-and-stones/' \
  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6' \
  -H 'cookie: __cfduid=d77f603eea62f80c53cc89d4633edc07e1617862697; _ga=GA1.2.816940264.1617862699; __stripe_mid=a3bfd9ed-1fa9-4f68-91bb-087c09ff7028889b8e; _gid=GA1.2.532937487.1619642789; csrftoken=4yMZw86nBrCPTdluRT7UhlZXV9hVUMxJ4bTelwkptkC1UXd1mzlMY8GJ97MeJhP0; messages="5f19f4eb73d7b76725693b4dfd81c75b4fba0f50$[[\"__json_message\"\0540\05425\054\"You have signed out.\"]\054[\"__json_message\"\0540\05425\054\"Successfully signed in as maorui_way_up.\"]\054[\"__json_message\"\0540\05425\054\"You have signed out.\"]\054[\"__json_message\"\0540\05425\054\"Successfully signed in as maorui_way_up.\"]]"; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMTU0NzY1MiIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNDg2OThiODFjOTU5YmRlMTY3Y2ExN2JhMjM4YmJjZmJkNzg0Njg4MSIsImlkIjoxNTQ3NjUyLCJlbWFpbCI6ImZyYW5rLnN1bi5yZWdpc3RlckBnbWFpbC5jb20iLCJ1c2VybmFtZSI6Im1hb3J1aV93YXlfdXAiLCJ1c2VyX3NsdWciOiJtYW9ydWlfd2F5X3VwIiwiYXZhdGFyIjoiaHR0cHM6Ly93d3cuZ3JhdmF0YXIuY29tL2F2YXRhci83OTg5YmFmMDM4NDM3ZmNmMDEzNDAwMDIzZmVmMjMzOC5wbmc_cz0yMDAiLCJyZWZyZXNoZWRfYXQiOjE2MTk2NDc1MzksImlwIjoiNDkuMjU1LjUuMjQyIiwiaWRlbnRpdHkiOiJiMjE2OTM2Y2ZjODgyOGNhODE5NzdkYzA5M2ZiODU5YyIsInNlc3Npb25faWQiOjg0MTY0NTQsIl9zZXNzaW9uX2V4cGlyeSI6MTIwOTYwMH0.CeyWyZhU_SeyW7iF7R5P7Vjf_skqu57S3xvezb9OPD0; __cf_bm=1671c355481d2852161094e245e218e964559730-1619651921-1800-AQoxDptedpypqHv7mUhZg6ZZCWsDqU86vrddP+PqGDzbCLsk0O/sdU0JCW5j3Z3vkPQhy6tsRZYyVsQ/mhxyNJM=; _gat=1; c_a_u="bWFvcnVpX3dheV91cA==:1lbtTC:8CzsKrzjt27qJkn-UejdDTwW4hA"' \
  --data-raw '{"operationName":"allQuestionsStatusesRaw","variables":{},"query":"query allQuestionsStatusesRaw {\n  allQuestions: allQuestionsRaw {\n    questionId\n    status\n    __typename\n  }\n}\n"}' \
  --compressed
```


Fetch submission detail
```
curl 'https://leetcode.com/graphql' \
  -H 'authority: leetcode.com' \
  -H 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
  -H 'x-newrelic-id: UAQDVFVRGwEAXVlbBAg=' \
  -H 'dnt: 1' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36' \
  -H 'content-type: application/json' \
  -H 'accept: */*' \
  -H 'x-csrftoken: 4yMZw86nBrCPTdluRT7UhlZXV9hVUMxJ4bTelwkptkC1UXd1mzlMY8GJ97MeJhP0' \
  -H 'origin: https://leetcode.com' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://leetcode.com/problems/jewels-and-stones/' \
  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6' \
  -H 'cookie: __cfduid=d77f603eea62f80c53cc89d4633edc07e1617862697; _ga=GA1.2.816940264.1617862699; __stripe_mid=a3bfd9ed-1fa9-4f68-91bb-087c09ff7028889b8e; _gid=GA1.2.532937487.1619642789; csrftoken=4yMZw86nBrCPTdluRT7UhlZXV9hVUMxJ4bTelwkptkC1UXd1mzlMY8GJ97MeJhP0; messages="5f19f4eb73d7b76725693b4dfd81c75b4fba0f50$[[\"__json_message\"\0540\05425\054\"You have signed out.\"]\054[\"__json_message\"\0540\05425\054\"Successfully signed in as maorui_way_up.\"]\054[\"__json_message\"\0540\05425\054\"You have signed out.\"]\054[\"__json_message\"\0540\05425\054\"Successfully signed in as maorui_way_up.\"]]"; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMTU0NzY1MiIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNDg2OThiODFjOTU5YmRlMTY3Y2ExN2JhMjM4YmJjZmJkNzg0Njg4MSIsImlkIjoxNTQ3NjUyLCJlbWFpbCI6ImZyYW5rLnN1bi5yZWdpc3RlckBnbWFpbC5jb20iLCJ1c2VybmFtZSI6Im1hb3J1aV93YXlfdXAiLCJ1c2VyX3NsdWciOiJtYW9ydWlfd2F5X3VwIiwiYXZhdGFyIjoiaHR0cHM6Ly93d3cuZ3JhdmF0YXIuY29tL2F2YXRhci83OTg5YmFmMDM4NDM3ZmNmMDEzNDAwMDIzZmVmMjMzOC5wbmc_cz0yMDAiLCJyZWZyZXNoZWRfYXQiOjE2MTk2NDc1MzksImlwIjoiNDkuMjU1LjUuMjQyIiwiaWRlbnRpdHkiOiJiMjE2OTM2Y2ZjODgyOGNhODE5NzdkYzA5M2ZiODU5YyIsInNlc3Npb25faWQiOjg0MTY0NTQsIl9zZXNzaW9uX2V4cGlyeSI6MTIwOTYwMH0.CeyWyZhU_SeyW7iF7R5P7Vjf_skqu57S3xvezb9OPD0; __cf_bm=1671c355481d2852161094e245e218e964559730-1619651921-1800-AQoxDptedpypqHv7mUhZg6ZZCWsDqU86vrddP+PqGDzbCLsk0O/sdU0JCW5j3Z3vkPQhy6tsRZYyVsQ/mhxyNJM=; _gat=1; c_a_u="bWFvcnVpX3dheV91cA==:1lbtTC:8CzsKrzjt27qJkn-UejdDTwW4hA"' \
  --data-raw $'{"operationName":"Submissions","variables":{"offset":0,"limit":20,"lastKey":null,"questionSlug":"jewels-and-stones"},"query":"query Submissions($offset: Int\u0021, $limit: Int\u0021, $lastKey: String, $questionSlug: String\u0021) {\\n  submissionList(offset: $offset, limit: $limit, lastKey: $lastKey, questionSlug: $questionSlug) {\\n    lastKey\\n    hasNext\\n    submissions {\\n      id\\n      statusDisplay\\n      lang\\n      runtime\\n      timestamp\\n      url\\n      isPending\\n      memory\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n"}' \
  --compressed
```


## Response

```python
{'title': 'Two Sum', 'titleSlug': 'two-sum', 'translatedTitle': None, 'questionId': '1', 'questionFrontendId': '1', 'status': None, 'difficulty': 'Easy', 'isPaidOnly': False, '__typename': 'QuestionNode', 'topicTags': [{'name': 'Array', 'slug': 'array', 'translatedName': None, '__typename': 'TopicTagNode'}, {'name': 'Hash Table', 'slug': 'hash-table', 'translatedName': None, '__typename': 'TopicTagNode'}]}
```