from enum import Enum
from typing import List, Dict, Union
import requests


class SampleFeedbackType(Enum):
    """
    Enum for sample feedback type.
    """

    GOOD = "good"
    BAD = "bad"


class Trainloop:
    """
    A simple TrainLoop client for sending message data.
    """

    def __init__(self, api_key: str):
        self.api_key = api_key
        self.base_url = "https://app.trainloop.ai"
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {self.api_key}",
        }

    def send_data(
        self,
        messages: List[Dict[str, Union[str, dict]]],
        dataset_id: str,
        sample_feedback: SampleFeedbackType = SampleFeedbackType.GOOD,
    ) -> bool:
        """
        Sends messages and sample feedback to the TrainLoop API.

        :param messages: A list of dicts, e.g.:
                [
                    {"role": "system", "content": "..."},
                    {"role": "user", "content": "..."}
                ]
        :param sample_feedback: A SampleFeedbackType indicating the user feedback, e.g. GOOD or BAD.
        :param dataset_id: The ID of the dataset to send the data to.
        :return: True if the request succeeded (status_code == 200), else False.
        """

        url = f"{self.base_url}/api/datasets/collect"
        payload = {
            "messages": messages,
            "sample_feedback": sample_feedback.value,  # Use .value to get the string representation
            "dataset_id": dataset_id,
        }

        try:
            response = requests.post(
                url, json=payload, headers=self.headers, timeout=10
            )
            if response.status_code == 200:
                return True

            # Read the response body for more detailed error message
            print(f"Request returned status {response.status_code}: {response.text}")
            return False
        except requests.RequestException as e:
            print(f"RequestException: {e}")
            return False
