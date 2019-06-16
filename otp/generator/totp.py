
from django_otp.oath import TOTP
from django_otp.util import random_hex
import time


class TOTPVerification:
    def __init__(self):
        self.key = random_hex(20)
        self.last_verified_counter = -1
        self.verified = False
        self.number_of_digits = 6
        self.token_validity_period = 35

    def totp_obj(self):
        totp = TOTP(key=self.key,
                    step=self.token_validity_period,
                    digits=self.number_of_digits)
        totp.time = time.time()
        return totp

    def generate_token(self):
        totp = self.totp_obj()
        token = str(totp.token()).zfill(6)
        return token

