from django.shortcuts import render
from django.http import HttpResponseRedirect, HttpResponseBadRequest

def home(request):
    return render(request, 'home.html')