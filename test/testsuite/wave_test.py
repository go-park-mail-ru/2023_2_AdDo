import websocket
import utils


def on_message(ws, message):
    print(message)


def on_error(ws, error):
    print(f"Ошибка: {error}")


def on_close(ws, qw, qe):
    print("Соединение закрыто")


def on_open(ws):
    print("соединение открыто")
    while True:
        ws.send(str(1))


if __name__ == "__main__":
    headers, cookies = utils.init_authorized_user_headers_and_cookies()
    header = ''
    cookie = ''

    for key in headers:
        header += key + '=' + headers[key] + ';'
    for key in cookies:
        cookie += key + '=' + cookies[key] + ';'
    print(header)
    print(cookie)

    websocket.enableTrace(True)
    ws = websocket.WebSocketApp(utils.ws_url + "/wave",
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close, header=headers, cookie=cookie)
    ws.on_open = on_open
    ws.run_forever()
