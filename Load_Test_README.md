# Load Testing with K6

This document describes how to perform load testing on the Person Data API using [k6](https://k6.io/), an open-source load testing tool.

## Prerequisites

1. Install k6 by following the [official installation guide](https://k6.io/docs/getting-started/installation/)
2. Make sure the Person Data API server is running on `localhost:8080` (or update the `baseUrl` in the load test script)

## Running the Load Test

1. Start the API server.
   ```bash
   go run generate_json.go --server --port=:8080
   ```

2. In a separate terminal, run the load test.
   ```bash
   k6 run load_test.js
   ```

## Load Test Scenarios

The load test script includes three different scenarios.
See https://grafana.com/docs/k6/latest/using-k6/scenarios/

1. **Warmup**: A constant number of 5 virtual users for 30 seconds to warm up the system
2. **Ramp-up**: A gradual increase from 0 to 20 virtual users over 30 seconds, then maintaining 20 users for 1 minute, and finally ramping down to 0 over 30 seconds
3. **Stress**: A constant rate of 50 requests per second for 1 minute to stress test the system

## Endpoints Tested

The load test script testCases the following endpoints.

1. `GET /health`: Health check endpoint
2. `GET /api/persons`: Person data endpoint with default parameter (n=10)
3. `GET /api/persons?n=5`: Person data endpoint with a small number of records
4. `GET /api/persons?n=50`: Person data endpoint with a larger number of records
5. `GET /api/persons?n=abc`: Person data endpoint with an invalid parameter (should return 400)
6. `GET /api/persons?n=-5`: Person data endpoint with a negative parameter (should return 400)

## Performance Thresholds

The load test script defines the following performance thresholds.
See https://grafana.com/docs/k6/latest/using-k6/thresholds/
and https://github.com/grafana/k6-learn/blob/main/Modules/II-k6-Foundations/07-Setting-test-criteria-with-thresholds.md

1. 95% of requests should complete within 500ms. You can adjust this to P99 too. 
2. Less than 1% of requests should fail. Because I have limited the number of Virtual Users, this is highly unlikely.
3. Less than 1% of requests should have errors. I have setup this example in case you increase the number of records the API creates.

## Interpreting Results

After running the load test, k6 will output a summary of the results.
See this https://github.com/grafana/k6-learn/blob/main/Modules/II-k6-Foundations/03-Understanding-k6-results.md
- Request metrics (total requests, request rate, etc.)
- Response time metrics (min, max, average, percentiles)
- Error rate metrics
- Checks (assertions) results

## Customizing the Load Test

You can customize the load test by modifying the `options` object in the `load_test.js` file.
See https://grafana.com/docs/k6/latest/using-k6/k6-options/how-to/

- Adjust the number of virtual users or request rates. Do Load vs. Stress tests. 
- Change the duration of each scenario
- Modify the performance thresholds
- Add or remove endpoints to test

## Troubleshooting

If you encounter any issues while running the load test:

1. Make sure the API server is running and accessible (I ran into this a few times)
2. Check that k6 is installed correctly or update. I used the Homebrew. 
3. Verify that the `baseUrl` in the load test script matches the URL of your API server. Most of the time if you used different port.
