#!/usr/bin/env python3
import requests, re
from getpass import getpass

def main():
    # login
    login_url = 'https://leetcode.com/accounts/login/'
    # print('Login to leetcode...')
    # username = input('username: ')
    # password = getpass('password: ')
    username = "maorui_way_up"
    password = "goulin0009"

    session = requests.Session()

    response = session.request('GET', login_url)
    if response.status_code != 200 or not session.cookies['csrftoken']:
        print('Error in get login page')
        print(response.status_code)
        print(response.text)
        return
    cookie = session.cookies.get_dict()
    print(cookie)

    csrftoken = cookie["csrftoken"]
    __cfduid = cookie["__cfduid"]
    
    headers = {
            'Referer': login_url,
            "Cookie": "__cfduid=" + __cfduid + ";" + "csrftoken=" + csrftoken,
            'x-csrftoken': csrftoken,
            'user-agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36',
            '_ga': 'GA1.2.816940264.1617862699;'
            '__stripe_mid': 'a3bfd9ed-1fa9-4f68-91bb-087c09ff7028889b8e',
            '_gid': 'GA1.2.532937487.1619642789',
            }

    query = {
            "operationName": "allQuestions",
            "variables": {},
            "query": "query allQuestions {\n  allQuestions {\n    ...questionSummaryFields\n    __typename\n  }\n}\n\nfragment questionSummaryFields on QuestionNode {\n  title\n  titleSlug\n  translatedTitle\n  questionId\n  questionFrontendId\n  status\n  difficulty\n  isPaidOnly\n  __typename\n topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n}\n"
            }
    response = session.request('POST', 'https://leetcode.com/graphql', headers = headers, json = query)
    if response.status_code == 200:
        print('Got problem list')
    else:
        print('Error get problem list')
        print(response.status_code)
        print(response.text)
        return
    problems = response.json()['data']['allQuestions']

    print(type(problems))
    print(len(problems), problems[0], type(problems[0]))
    for problem in problems:
        if problem["status"] != None:
            print(problem)

if __name__ == '__main__':
    main()