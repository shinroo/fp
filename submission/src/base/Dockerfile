FROM infologistix/docker-selenium-python:alpine

# Install required packages
RUN apk add --no-cache \
    postgresql-dev \
    python3-dev \
    gcc \
    musl-dev

# Install Python dependencies
RUN mkdir -p /home/seluser/.local/lib && chmod -R 777 /home/seluser/.local
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy common dependencies
RUN mkdir -p /app
COPY /base /app/base