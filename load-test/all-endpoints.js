import http from 'k6/http';
import { check } from 'k6';

export let options = {
  scenarios: {
    http_traffic: {
      executor: 'ramping-arrival-rate',
      startRate: 1,      
      timeUnit: '1s',
      preAllocatedVUs: 50,
      maxVUs: 100,
      stages: [
        { target: 5, duration: '2m' },   
        { target: 5, duration: '3m' },   
        { target: 0, duration: '1m' },   
      ],
      exec: 'httpRequests',
    },

    cpu_burn: {
      executor: 'constant-vus',
      vus: 2,             
      duration: '6m',     
      exec: 'burnCpu',
    },
  },

  thresholds: {
    'http_req_duration{scenario:http_traffic}': ['p(95)<500'],
  },
};

export function httpRequests() {
  let res;

  res = http.get(`${__ENV.API_URL}/`, {
    headers: { 
      'Cache-Control': 'no-cache, no-store, must-revalidate' 
    }
  });
  check(res, { 'GET / 200': r => r.status === 200 });

  res = http.get(`${__ENV.API_URL}/users`, {
    headers: { 
      'Cache-Control': 'no-cache, no-store, must-revalidate' 
    }
  });
  check(res, { 'GET /users 200': r => r.status === 200 });

  const newUser = {
    id: Math.floor(Math.random() * 1e6),
    name: `User${Math.random().toString(36).substr(2, 5)}`,
    email: `user${Date.now()}@example.com`,
  };
  res = http.post(`${__ENV.API_URL}/users`, JSON.stringify(newUser), {
    headers: { 'Content-Type': 'application/json', 'Cache-Control': 'no-cache, no-store, must-revalidate'  },
  });
  check(res, { 'POST /users 201': r => r.status === 201 });

  const input = {
    x: Math.floor(Math.random() * 100),
    y: Math.floor(Math.random() * 100)
  }
  res = http.post(
    `${__ENV.API_URL}/compute`,
    JSON.stringify(input),
    { headers: { 'Content-Type': 'application/json','Cache-Control': 'no-cache, no-store, must-revalidate'} }
  );
  check(res, { 'POST /compute status 200': (r) => r.status === 200 });
}

export function burnCpu() {
  const end = Date.now() + 1000; 
  while (Date.now() < end) {
    Math.sqrt(12345.6789);
  }
}
