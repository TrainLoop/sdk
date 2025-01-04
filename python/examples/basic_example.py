from trainloop.client import Trainloop
from openai import OpenAI


def main():
    # Initialize the TrainLoop client with your api key
    client = Trainloop(api_key="tl_...")

    # Example messages
    messages = [
        {"role": "system", "content": "System message here"},
        {"role": "user", "content": "Hello from the user!"},
    ]
    openai_client = OpenAI()

    # Make an openai request
    openai_response = openai_client.chat.completions.create(
        model="gpt-4o",
        messages=messages,
    )

    chat_thread = openai_response.choices[0].message
    print(chat_thread)

    # Send data
    success = client.send_data(messages, sample_feedback="good")
    if success:
        print("Data sent successfully!")
    else:
        print("Data was not sent. Check logs for details.")


if __name__ == "__main__":
    main()
