import names
import requests
import hashlib

url = 'http://localhost:8888/api/v1'
ws_url = 'ws://localhost:8888/api/v1'


def gen_random_valid_register_data():
    username = names.get_full_name().replace(' ', '').lower()
    email = username + '@mail.ru'
    register_data = {
        'Email': email,
        'Password': 'password',
        'Username': username,
        'BirthDate': '2003-01-12',
    }
    return register_data


def gen_random_valid_update_data():
    username = names.get_full_name().replace(' ', '').lower()
    email = username + '@mail.ru'
    register_data = {
        'Email': email,
        'Username': username,
        'BirthDate': '2003-01-12',
    }
    return register_data


def get_csrf_headers_and_cookies():
    #response = requests.get(url + '/get_csrf')
    cookies = {
        #'X-Csrf-Token': response.cookies['X-Csrf-Token']
    }
    headers = {
        #'X-Csrf-Token': response.headers['X-Csrf-Token']
    }
    return headers, cookies


def init_authorized_user_headers_and_cookies():
    data = gen_random_valid_register_data()
    headers, cookies = get_csrf_headers_and_cookies()
    response = requests.post(url + "/sign_up", headers=headers, json=data, cookies=cookies)
    cookies['JSESSIONID'] = response.cookies['JSESSIONID']
    return headers, cookies
