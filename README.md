# Final Project: SPCA Pet Matchmaker
![](logo.jpeg)

Final Project for University of London B.Sc. Computer Science (CM3070)

This project implements an end to end pet matchmaker for all the SPCAs in South Africa which publish pets which are up for adoption on their websites.

The overall architecture is as follows:
- scrapers using selenium and beautiful soup
- airflow for scheduling and workflow management
- postgres database
- golang REST API
- golang SSR WebApp

## Dev Sources
1. Base image for scrapers: https://hub.docker.com/r/infologistix/docker-selenium-python
2. Similarity metric: https://en.wikipedia.org/wiki/Jaccard_index

## Data Sources
1. American kennel club dataset: https://github.com/kkakey/dog_traits_AKC/blob/main/data/breed_traits.csv
2. Africanis data: https://wagwalking.com/breed/africanis