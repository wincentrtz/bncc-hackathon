from rest_framework import serializers
from rest_framework import viewsets
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions
from django.http import Http404
from .models import OTP
from .totp import TOTPVerification

class OTPSerializer(serializers.ModelSerializer):
    class Meta:
        model = OTP
        fields = ('__all__')

class OTPGenerator(APIView):
    def post(self, request, format=None):
        serializer = OTPSerializer(data=request.data)
        if serializer.is_valid():
            totp = TOTPVerification()
            token = totp.generate_token()
            
            otp = OTP(token=token, user_id=request.data['user_id'])
            serializer = OTPSerializer(otp)
            serializer.save()
            
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

class OTPVerifier(APIView):
    def get_object(self, pk):
        try:
            return OTP.objects.get(pk=pk)
        except OTP.DoesNotExist:
            raise Http404

    def post(self, request, format=None):
        otp = self.get_object(request['token'])

        status = {
            'verified' : False
        }

        if not otp.user_id == request.data['user_id']:
            return Response(status)

        status['verified'] = True

        return Response(status)
