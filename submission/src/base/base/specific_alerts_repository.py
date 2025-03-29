from .common import SpecificAlert
from .database import Database

class SpecificAlertsRepository:
    def __init__(self):
        self.db = Database()

    def mark_actioned(self, alert_id):
        update_query = """
        UPDATE specific_alert
        SET actioned = TRUE
        WHERE id = %s;
        """
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(update_query, (alert_id,))
            conn.commit()
            print("Data updated successfully!")
        except Exception as e:
            print(f"Failed to update data: {e}")
            raise

    def get_unactioned(self):
        unactioned_alerts = []
        select_query = """
        SELECT
        specific_alert.id AS id,
        specific_alert.breed AS breed,
        specific_alert.gender AS gender,
        specific_alert.life_stage AS life_stage,
        account.email AS email
        FROM specific_alert
        LEFT JOIN account ON specific_alert.account_id = account.id
        WHERE specific_alert.actioned = FALSE;
        """
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(select_query)
            rows = cursor.fetchall()
            for row in rows:
                unactioned_alerts.append(
                    SpecificAlert(
                        row[0],
                        row[1],
                        row[2],
                        row[3],
                        row[4],
                    )
                )

        except Exception as e:
            print(f"Failed to select data: {e}")
            raise
        return unactioned_alerts