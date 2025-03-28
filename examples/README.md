# TrainLoop SDK Examples

This directory contains examples for using the TrainLoop SDK in different languages:
- Go
- Python
- TypeScript

## Prerequisites

To run these data colletion examples, you'll need:
- Node.js and npm (for TypeScript examples)
- Python 3.7+ (for Python examples)
- Go 1.18+ (for Go examples)
- A TrainLoop API key

## Quickstart

Run `bash setup.sh` to setup the environment and install dependencies. You can then run `npm run ts:basic`, `npm run py:basic`, or `npm run go:basic` to run the examples.

## Manual Setup

1. Clone the repository
2. Create a `.env` file in the root of this directory with your TrainLoop API key:
   ```
   TRAINLOOP_API_KEY=your_api_key_here
   ```

3. Install dependencies for all languages:
   ```
   npm run install:all
   ```

   Or install for a specific language:
   ```
   npm run install:ts    # TypeScript
   npm run install:py    # Python
   npm run install:go    # Go
   ```

## Running Examples

### TypeScript
```
npm run ts:basic
```

### Python
```
npm run py:basic
```

### Go
```
npm run go:basic
```

## Examples

Each language directory contains the following examples:
- Basic usage: Shows how to initialize the client and send data to TrainLoop
