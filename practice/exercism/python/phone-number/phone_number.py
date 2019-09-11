class Phone():

    def __init__(self, phone_number):
        self.area_code = None
        self.exchange_code = None
        self.subscriber_number = None
        self.number = None
        self.__cleanNumber(phone_number)

    def __cleanNumber(self, num):
        cleanNumber = ""

        # strip non-numeric chars
        for c in num:
            if c.isdigit():
                cleanNumber += c

        # too many digits
        if len(cleanNumber) > 11:
            raise ValueError("Too many digits")

        # country code other than '1'
        if len(cleanNumber) == 11 and cleanNumber[0] != "1":
            raise ValueError("Country code should be '1'")

        # remove '1' country code
        if len(cleanNumber) == 11 and cleanNumber[0] == "1":
            cleanNumber = cleanNumber[1:]

        # components
        self.area_code = cleanNumber[0:3]
        self.exchange_code = cleanNumber[3:6]
        self.subscriber_number = cleanNumber[6:]

        # check area code and exchange code do NOT start with 0 or 1
        if len(self.area_code) and int(self.area_code[0]) < 2:
            raise ValueError("Area code cannot start with 0 or 1")
        if len(self.exchange_code) and int(self.exchange_code[0]) < 2:
            raise ValueError("Exchange code cannot start with 0 or 1")

        self.number = self.area_code + self.exchange_code + self.subscriber_number

    def pretty(self):
        return f"({self.area_code}) {self.exchange_code}-{self.subscriber_number}"
