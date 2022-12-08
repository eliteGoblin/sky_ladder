
csdn session


```
访问page: https://leetcode.com/problemset/all/
```

```
curl 'https://leetcode.com/api/problems/all/' \
  -H 'authority: leetcode.com' \
  -H 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
  -H 'x-newrelic-id: UAQDVFVRGwEAXVlbBAg=' \
  -H 'dnt: 1' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36' \
  -H 'content-type: application/json' \
  -H 'accept: application/json, text/javascript, */*; q=0.01' \
  -H 'x-requested-with: XMLHttpRequest' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://leetcode.com/problemset/all/' \
  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6' \
  -H 'cookie: __cfduid=d77f603eea62f80c53cc89d4633edc07e1617862697; _ga=GA1.2.816940264.1617862699; __stripe_mid=a3bfd9ed-1fa9-4f68-91bb-087c09ff7028889b8e; _gid=GA1.2.532937487.1619642789; csrftoken=Xfi9z4m8yHjKHTMVOTv6F6dsV9kf3OuIVfrjAE3zC28CeOMAQupgEaPyGGqPf0mM; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMTU0NzY1MiIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNDg2OThiODFjOTU5YmRlMTY3Y2ExN2JhMjM4YmJjZmJkNzg0Njg4MSIsImlkIjoxNTQ3NjUyLCJlbWFpbCI6ImZyYW5rLnN1bi5yZWdpc3RlckBnbWFpbC5jb20iLCJ1c2VybmFtZSI6Im1hb3J1aV93YXlfdXAiLCJ1c2VyX3NsdWciOiJtYW9ydWlfd2F5X3VwIiwiYXZhdGFyIjoiaHR0cHM6Ly93d3cuZ3JhdmF0YXIuY29tL2F2YXRhci83OTg5YmFmMDM4NDM3ZmNmMDEzNDAwMDIzZmVmMjMzOC5wbmc_cz0yMDAiLCJyZWZyZXNoZWRfYXQiOjE2MTk2NTM1ODQsImlwIjoiNDkuMjU1LjUuMjQyIiwiaWRlbnRpdHkiOiJiMjE2OTM2Y2ZjODgyOGNhODE5NzdkYzA5M2ZiODU5YyIsInNlc3Npb25faWQiOjg0MTc4ODksIl9zZXNzaW9uX2V4cGlyeSI6MTIwOTYwMH0.meaq4basnSFAnqjs3Rj40mzoytTKHRVKVawaMa8je1E; c_a_u="bWFvcnVpX3dheV91cA==:1lbzKa:RY1aJ2vwcfXUl0U5Q8FVqNycBG4"; __cf_bm=440ed9cefabbab4b716d03a25f2a6786947a7f79-1619674521-1800-AT+Nl3SIH/GSnemARDe+XkAbHTdXZ/Msyx/u5jNuxswV2i1EMZirP9tW6WqNcQmtQGiKWLWMvCMvhIkYPCnY+0w=; _gat=1' \
  --compressed | jq '. | select (.status != null)'
```


## Amazon session

```
curl 'https://leetcode.com/api/problems/all/' \
  -H 'authority: leetcode.com' \
  -H 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
  -H 'x-newrelic-id: UAQDVFVRGwEAXVlbBAg=' \
  -H 'dnt: 1' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36' \
  -H 'content-type: application/json' \
  -H 'accept: application/json, text/javascript, */*; q=0.01' \
  -H 'x-requested-with: XMLHttpRequest' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://leetcode.com/problemset/all/' \
  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6' \
  -H 'cookie: __cfduid=d77f603eea62f80c53cc89d4633edc07e1617862697; _ga=GA1.2.816940264.1617862699; __stripe_mid=a3bfd9ed-1fa9-4f68-91bb-087c09ff7028889b8e; _gid=GA1.2.532937487.1619642789; csrftoken=Xfi9z4m8yHjKHTMVOTv6F6dsV9kf3OuIVfrjAE3zC28CeOMAQupgEaPyGGqPf0mM; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMTU0NzY1MiIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNDg2OThiODFjOTU5YmRlMTY3Y2ExN2JhMjM4YmJjZmJkNzg0Njg4MSIsImlkIjoxNTQ3NjUyLCJlbWFpbCI6ImZyYW5rLnN1bi5yZWdpc3RlckBnbWFpbC5jb20iLCJ1c2VybmFtZSI6Im1hb3J1aV93YXlfdXAiLCJ1c2VyX3NsdWciOiJtYW9ydWlfd2F5X3VwIiwiYXZhdGFyIjoiaHR0cHM6Ly93d3cuZ3JhdmF0YXIuY29tL2F2YXRhci83OTg5YmFmMDM4NDM3ZmNmMDEzNDAwMDIzZmVmMjMzOC5wbmc_cz0yMDAiLCJyZWZyZXNoZWRfYXQiOjE2MTk2NTM1ODQsImlwIjoiNDkuMjU1LjUuMjQyIiwiaWRlbnRpdHkiOiJiMjE2OTM2Y2ZjODgyOGNhODE5NzdkYzA5M2ZiODU5YyIsInNlc3Npb25faWQiOjg0MTc4ODksIl9zZXNzaW9uX2V4cGlyeSI6MTIwOTYwMH0.meaq4basnSFAnqjs3Rj40mzoytTKHRVKVawaMa8je1E; c_a_u="bWFvcnVpX3dheV91cA==:1lbzKa:RY1aJ2vwcfXUl0U5Q8FVqNycBG4"; __cf_bm=8b85c6246e657068f68e17ec0c40c37e0cb9e618-1619676341-1800-Ab4E/EJ3K95QdE7DjyXL5/XyQY7NyWUaHswKpmFN1xlKs2pB1t9eh2MRyWTv7PEShiSn7kkjauYzVttuef4JCkM=; _gat=1' \
  --compressed
```