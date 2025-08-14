import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  scenarios: {
    load_test: {
      executor: 'ramping-arrival-rate',
      startRate: 1,
      timeUnit: '1s',
      preAllocatedVUs: 50,
      maxVUs: 100,
      stages: [
        { target: 1, duration: '1m' },   
        { target: 10, duration: '5m' },   
        { target: 10, duration: '2m' },   
        { target: 0, duration: '1m' },    
      ],
    },
  },
  thresholds: {
    http_req_duration: ['p(95)<500'], 
  },
};

export default function () {
  let res;

  res = http.get(`https://api.anytimeagile.com/compute`, {
    headers: { 
      'Cache-Control': 'no-cache, no-store, must-revalidate' 
    }
  });
  check(res, { 'GET /compute status 200': (r) => r.status === 200 });

  sleep(1);
}
