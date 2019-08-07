import string
from secrets import choice


class Robot(object):

    def __init__(self):
        self.name = ""
        self.__newname()

    def reset(self):
        self.__newname()

    def __newname(self):
        letters = ''.join(choice(string.ascii_letters) for i in range(2))
        numbers = ''.join(choice(string.digits) for i in range(3))
        self.name =  letters.upper() + numbers
