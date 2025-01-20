import os
from trainloop.client import Trainloop, SampleFeedbackType


def main():
    # Initialize the TrainLoop client with your api key
    client = Trainloop(api_key=os.getenv("TRAINLOOP_API_KEY"))

    # Example messages
    messages = [
        {"role": "system", "content": "System message here"},
        {"role": "user", "content": "Hello from the user!"},
    ]

    # Send data
    success = client.send_data(
        messages, sample_feedback=SampleFeedbackType.GOOD, dataset_id="test-dataset"
    )
    if success:
        print("Data sent successfully!")
    else:
        print("Data was not sent. Check logs for details.")


if __name__ == "__main__":
    main()
