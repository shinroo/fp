FROM infologistix/docker-selenium-python:alpine

# Install Python dependencies
RUN mkdir -p /home/seluser/.local/lib && chmod -R 777 /home/seluser/.local
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY main.py ./

# Run scraper
CMD ["python3", "main.py"]