from .common import SimilarityAlert
from .database import Database

class SimilarityAlertsRepository:
    def __init__(self):
        self.db = Database()

    def mark_actioned(self, alert_id):
        update_query = """
        UPDATE similarity_alert
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
        similarity_alert.id AS id,
        similarity_alert.similarity_threshold AS threshold,
        account.email AS email,
        account.id AS account_id
        FROM similarity_alert
        LEFT JOIN account ON similarity_alert.account_id = account.id
        WHERE similarity_alert.actioned = FALSE;
        """
        try:
            conn = self.db.get_conn()
            cursor = conn.cursor()
            cursor.execute(select_query)
            rows = cursor.fetchall()
            for row in rows:
                unactioned_alerts.append(
                    SimilarityAlert(
                        row[0],
                        row[1],
                        row[2],
                        row[3]
                    )
                )

        except Exception as e:
            print(f"Failed to select data: {e}")
            raise
        return unactioned_alerts