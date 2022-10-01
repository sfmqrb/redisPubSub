import json
from random import random
import time
import redis
from data import get_data

red = redis.StrictRedis('localhost',
                        7979,
                        charset="utf-8",
                        decode_responses=True)


def publish_notif(data: dict):
    n = red.publish(f'{data["channel"]}', "wake_up")
    print("num of subscribers: ", n)

while True:
    data = get_data()
    json_object = json.dumps(data, indent=4)
    red.lpush(f'{data["channel"]}', json_object)
    publish_notif(data)
    time.sleep(1 - random())
