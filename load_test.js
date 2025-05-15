import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';

// Base URL for the API
const baseUrl = 'http://localhost:8080';

const errorRate = new Rate('errors');

// Configration for load test

export const options = {
  // Load Testing Scenarios https://grafana.com/docs/k6/latest/using-k6/scenarios/
  // See https://grafana.com/docs/k6/latest/using-k6/scenarios/executors/ for executor options.
  scenarios: {
    // Warm-up with a constant number of virtual users
    warmup: {
      executor: 'constant-vus',
      vus: 5,
      duration: '30s',
      gracefulStop: '5s',
      tags: { scenario: 'warmup' },
    },
    // Ramp-up test to simulate increasing load
    rampup: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '30s', target: 20 },
        { duration: '1m', target: 20 },
        { duration: '30s', target: 0 },
      ],
      gracefulStop: '5s',
      startTime: '30s', // Start after the warmup scenario
      tags: { scenario: 'rampup' },
    },
    // Stress test with a high number of requests
    stress: {
      executor: 'constant-arrival-rate',
      rate: 50,
      timeUnit: '1s',
      duration: '1m',
      preAllocatedVUs: 50,
      maxVUs: 100,
      startTime: '2m', // Start after the rampup scenario
      tags: { scenario: 'stress' },
    },
  },
  thresholds: {
    // Performance thresholds.
    // See https://grafana.com/docs/k6/latest/examples/get-timings-for-an-http-metric/
    // and https://grafana.com/docs/k6/latest/using-k6/metrics/reference/
    http_req_duration: ['p(95)<500'], // 95% of requests should complete within 500ms Add
    http_req_failed: ['rate<0.01'],   // This example is just to show that we can also test for failed request.
    'errors': ['rate<0.01'],          // Less than 1% of requests should have errors
  },
};

function checkResponse(response, expectedStatus) {
  const success = check(response, {
    'status is correct': (r) => r.status === expectedStatus,
    'response body is not empty': (r) => r.body.length > 0,
    'response time < 500ms': (r) => r.timings.duration < 500,
  });
  
  // If any check fails, increment the error rate
  errorRate.add(!success);
  
  return success;
}

// I have only used the default function. K6 gives you setup and teardown functions.
// See https://grafana.com/docs/k6/latest/using-k6/test-lifecycle/
export default function() {
  // Test the health endpoint
  const healthResponse = http.get(`${baseUrl}/health`);
  checkResponse(healthResponse, 200);
  
  // Test the persons endpoint with default parameter (n=10)
  const personsDefaultResponse = http.get(`${baseUrl}/api/persons`);
  checkResponse(personsDefaultResponse, 200);
  
  // Test the persons endpoint with a small number of records (n=5)
  const personsSmallResponse = http.get(`${baseUrl}/api/persons?n=5`);
  checkResponse(personsSmallResponse, 200);
  
  // Test the persons endpoint with a larger number of records (n=50)
  const personsLargeResponse = http.get(`${baseUrl}/api/persons?n=50`);
  checkResponse(personsLargeResponse, 200);
  
  // Test the persons endpoint with an invalid parameter (n=abc)
  const personsInvalidResponse = http.get(`${baseUrl}/api/persons?n=abc`);
  checkResponse(personsInvalidResponse, 400);
  
  // Test the persons endpoint with a negative parameter (n=-5)
  const personsNegativeResponse = http.get(`${baseUrl}/api/persons?n=-5`);
  checkResponse(personsNegativeResponse, 400);
  
  // Sleep between iterations to avoid overwhelming the server
  sleep(1);
}

export function teardown() {
  console.log('Load test completed');
}