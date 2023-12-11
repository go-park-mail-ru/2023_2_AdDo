import websocket
import utils


def on_message(ws, message):
    print(f"Получено сообщение: {message}")


def on_error(ws, error):
    print(f"Ошибка: {error}")


def on_close(ws):
    print("Соединение закрыто")


def on_open(ws):
    message = input("Введите сообщение: ")
    ws.send(message)


if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp(utils.ws_url + "/ws",
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    ws.on_open = on_open
    ws.run_forever()
