import os
import smtplib
from dotenv import load_dotenv
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

def send_email(sender_email, sender_password, recipient_email, subject, body):
    # Set up the MIME
    message = MIMEMultipart()
    message['From'] = sender_email
    message['To'] = recipient_email
    message['Subject'] = subject
    
    # Attach the body text
    message.attach(MIMEText(body, 'plain'))
    
    try:
        # Create SMTP session
        with smtplib.SMTP('smtp.gmail.com', 587) as server:
            server.starttls()  # Enable security
            server.login(sender_email, sender_password)  # Login
            text = message.as_string()
            server.sendmail(sender_email, recipient_email, text)
        print("Email sent successfully!")
    except Exception as e:
        print(f"Error sending email: {e}")

if __name__ == "__main__":
    load_dotenv('.env')

    sender_email = "robert.focke96@gmail.com"
    sender_password = os.getenv("GOOGLE_APP_PASSWORD")
    recipient_email = "robert_focke@yahoo.com"
    subject = "Test Email from Python"
    body = "Hello, this is a test email sent from Python!"
    
    send_email(sender_email, sender_password, recipient_email, subject, body)