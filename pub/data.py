import os
from random import random
import time


def get_data():
    five = int(random() * 5 // 1) + 1
    return {
        "current_time": time.time(),
        "user": os.getlogin(),
        "pid": os.getpid(),
        "ppid": os.getppid(),
        "channel": five # five channels to push
    }
    
if __name__ == "__main__":
    from pprint import pprint
    pprint(get_data())
