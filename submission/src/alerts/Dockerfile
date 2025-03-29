FROM scraper_base:latest

# Copy worker files
COPY main.py /app
COPY .env /app

WORKDIR /app

# Run scraper
CMD ["python3", "main.py"]