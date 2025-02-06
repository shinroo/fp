FROM infologistix/docker-selenium-python:alpine

# Install Python dependencies
RUN mkdir -p /home/seluser/.local/lib && chmod -R 777 /home/seluser/.local
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy common dependencies
RUN mkdir -p /app
COPY /base /app/base