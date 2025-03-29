import os
import smtplib
from dotenv import load_dotenv
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from base.similarity_alerts_repository import SimilarityAlertsRepository
from base.specific_alerts_repository import SpecificAlertsRepository
from base.dog_repository import DogRepository

def send_email(sender_email, sender_password, recipient_email, subject, body, image_link):
    # Set up the MIME
    message = MIMEMultipart()
    message['From'] = sender_email
    message['To'] = recipient_email
    message['Subject'] = subject
    
    # Attach the body text
    message.attach(MIMEText(body, 'plain'))
    message.attach(MIMEText(f'<img src="{image_link}" width=300px />', 'html'))
    
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
    subject = "SPCA Pet Matchmaker Alert"
    body = "Hello, we found a matching dog for you!"
    
    try:
        similarity_alerts_repository = SimilarityAlertsRepository()
        specific_alerts_repository = SpecificAlertsRepository()
        unactioned_similarity_alerts = similarity_alerts_repository.get_unactioned()
        unactioned_specific_alerts = specific_alerts_repository.get_unactioned()
        if unactioned_similarity_alerts:
            for alert in unactioned_similarity_alerts:
                print(f"Similarity Alert: {alert}")
        if unactioned_specific_alerts:
            for alert in unactioned_specific_alerts:
                print(f"Specific Alert: {alert}")
    except Exception as e:
        print(f"Failed to get unactioned alerts: {e}")
        raise

    try:
        dog_repository = DogRepository()
    except:
        print("Failed to create DogRepository")
        raise

    try:
        for alert in unactioned_similarity_alerts:
            matching_dogs = dog_repository.get_similar(alert.account_id, alert.threshold)
            if matching_dogs:
                for dog in matching_dogs:
                    body = f"Hello, we found a dog which matched your similarity alert! {dog.name} which is at {dog.spca_name}"
                    print(f"Sending similarity alert email to {alert.email}: {body}, {dog.image_url}")
                    send_email(sender_email, sender_password, alert.email, subject, body, dog.image_url)
                    similarity_alerts_repository.mark_actioned(alert.id)
    except Exception as e:
        print(f"Failed to process similarity alerts: {e}")
        raise

    try:
        for alert in unactioned_specific_alerts:
            matching_dogs = dog_repository.get_specific(alert.breed, alert.gender, alert.life_stage)
            if matching_dogs:
                for dog in matching_dogs:
                    body = f"Hello, we found a dog which matched your specific alert! {dog.name} which is at {dog.spca_name}"
                    print(f"Sending specific alert email to {alert.email}: {body}, {dog.image_url}")
                    send_email(sender_email, sender_password, alert.email, subject, body, dog.image_url)
                    specific_alerts_repository.mark_actioned(alert.id)
    except Exception as e:
        print(f"Failed to process specific alerts: {e}")
        raise