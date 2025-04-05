import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

/**
 * Constants for sample feedback type.
 */
export enum SampleFeedbackType {
  GOOD = 'good',
  BAD = 'bad'
}

/**
 * A client for sending message data to TrainLoop.
 */
export class Client {
  private apiKey: string;
  private baseUrl: string;
  private httpClient: AxiosInstance;

  /**
   * Initialize a new TrainLoop client with an API key
   * 
   * @param apiKey Authentication API key for TrainLoop
   */
  constructor(apiKey: string, baseUrl: string = 'https://app.trainloop.ai') {
    this.apiKey = apiKey;
    this.baseUrl = baseUrl;

    // Create an axios instance with detailed configuration to match other SDKs
    const config: AxiosRequestConfig = {
      timeout: 10000, // 10 seconds timeout
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.apiKey}`
      },
      maxRedirects: 5
    };

    this.httpClient = axios.create(config);
  }

  /**
   * Sends messages and sample feedback to the TrainLoop API.
   * 
   * @param messages A list of objects containing
   * @param sampleFeedback A feedback type string, either SampleFeedbackType.GOOD or SampleFeedbackType.BAD
   * @param datasetId The ID of the dataset to send the data to
   * @returns A Promise that resolves to true if successful, false otherwise
   */
  async sendData(
    messages: Record<string, string>[],
    sampleFeedback: SampleFeedbackType,
    datasetId: string
  ): Promise<boolean> {
    const url = `${this.baseUrl}/api/datasets/collect`;
    const payload = {
      messages,
      sample_feedback: sampleFeedback,
      dataset_id: datasetId
    };

    try {
      const response = await this.httpClient.post(url, payload);
      return response.status === 200;
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        console.error(`Request returned non-200 status code: ${error.response.status}, body: ${JSON.stringify(error.response.data)}`);
      } else {
        console.error(`Request error: ${error}`);
      }
      return false;
    }
  }
}
