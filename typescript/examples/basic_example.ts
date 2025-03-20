import { Client, SampleFeedbackType } from '../src';
import * as dotenv from 'dotenv';

// Load environment variables
dotenv.config();

async function main() {

  // Initialize the TrainLoop client with your api key
  const client = new Client(process.env.TRAINLOOP_API_KEY || '');

  // Example messages
  const messages = [
    { role: 'system', content: 'System message here' },
    { role: 'user', content: 'Hello from the user!' },
  ];

  // Send data
  const success = await client.sendData(
    messages,
    SampleFeedbackType.GOOD,
    'test-dataset'
  );

  if (success) {
    console.log('Data sent successfully!');
  } else {
    console.log('Data was not sent. Check logs for details.');
  }
}

// Execute the main function
main().catch(error => {
  console.error('Unexpected error:', error);
});
