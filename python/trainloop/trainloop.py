from typing import List, Dict, Union, Literal
from requests import Session
from requests.adapters import HTTPAdapter
from requests.exceptions import RequestException


class SampleFeedbackType:
    """
    Constants for sample feedback type.
    """

    GOOD = "good"
    BAD = "bad"


class Client:
    """
    A client for sending message data to TrainLoop.
    """

    def __init__(self, api_key: str, base_url: str = "https://app.trainloop.ai"):
        """
        Initialize a new TrainLoop client with an API key

        :param api_key: Authentication API key for TrainLoop
        """
        if not api_key:
            raise ValueError("API key is required")

        self.api_key = api_key
        self.base_url = base_url
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {self.api_key}",
        }
        # Create a session with detailed configuration to match Go's HTTP client
        self.session = Session()
        self.session.mount(
            "https://",
            HTTPAdapter(
                pool_connections=100,
                pool_maxsize=100,
                max_retries=0,
            ),
        )
        self.session.headers.update(self.headers)
        self.timeout = 10

    def send_data(
        self,
        messages: List[Dict[str, str]],
        sample_feedback: Literal[SampleFeedbackType.BAD, SampleFeedbackType.GOOD],
        dataset_id: str,
    ) -> bool:
        """
        Sends messages and sample feedback to the TrainLoop API.

        :param messages: A list of dicts, e.g.:
                [
                    {"role": "system", "content": "..."},
                    {"role": "user", "content": "..."}
                ]
        :param sample_feedback: A feedback type string, either SampleFeedbackType.GOOD or SampleFeedbackType.BAD
        :param dataset_id: The ID of the dataset to send the data to
        :return: True if the request succeeded (status_code == 200), else False
        """
        url = f"{self.base_url}/api/datasets/collect"
        payload = {
            "messages": messages,
            "sample_feedback": sample_feedback,
            "dataset_id": dataset_id,
        }

        try:
            # Send the request using the configured session
            response = self.session.post(url, json=payload)

            # Check for successful status code
            if response.status_code != 200:
                # Read the response body for more detailed error message
                print(
                    f"Request returned non-200 status code: {response.status_code}, body: {response.text}"
                )
                return False

            return True
        except RequestException as e:
            print(f"RequestException: {e}")
            return False
