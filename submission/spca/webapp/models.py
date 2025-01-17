from django.db import models

# Auth Models

# Models

class Province(models.TextChoices):
    GAUTENG = 'GP', 'Gauteng'
    WESTERN_CAPE = 'WC', 'Western Cape'
    EASTERN_CAPE = 'EC', 'Eastern Cape'
    KWAZULU_NATAL = 'KZN', 'KwaZulu-Natal'
    FREE_STATE = 'FS', 'Free State'
    LIMPOPO = 'LP', 'Limpopo'
    MPUMALANGA = 'MP', 'Mpumalanga'
    NORTHERN_CAPE = 'NC', 'Northern Cape'
    NORTH_WEST = 'NW', 'North West'

class SPCABranch(models.Model):
    id = models.UUIDField(primary_key=True)
    name = models.CharField(max_length=200)
    address = models.TextField()
    telephone = models.CharField(max_length=20)
    emergency_telephone = models.CharField(max_length=20)
    latitude = models.FloatField()
    longitude = models.FloatField()
    province = models.CharField(max_length=3, choices=Province.choices)
    website = models.URLField()

class Dog(models.Model):
    id = models.UUIDField(primary_key=True)
    name = models.CharField(max_length=200)
    spca = models.ForeignKey(SPCABranch, on_delete=models.CASCADE)