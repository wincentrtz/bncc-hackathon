from django.urls import path

import generator.api as api

urlpatterns = [
    path('token/', api.OTPGenerator.as_view()),
    path('verify/', api.OTPVerifier.as_view()),
]