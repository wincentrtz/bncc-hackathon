from django.db import models

# Create your models here.
class OTP(models.Model):
    token = models.CharField(max_length=100)
    user_id = models.IntegerField()