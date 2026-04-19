import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 10 },
   // { duration: '10s', target: 10 },
   // { duration: '10s', target: 30 },
   // { duration: '10s', target: 30 },
   // { duration: '10s', target: 0 },
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],
    http_req_failed: ['rate<0.01'],
  },
};

const BASE_URL = 'http://app:8000';


export function setup() {
  const registerRes = http.post(
    `${BASE_URL}/register`,
    JSON.stringify({ username: 'logintest_user', password: 'test1234' }),
    { headers: { 'Content-Type': 'application/json' } }
  );
  console.log('Register status:', registerRes.status);
}

export default function () {
  const loginRes = http.post(
    `${BASE_URL}/login`,
    JSON.stringify({ username: 'logintest_user', password: 'test1234' }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(loginRes, {
    'login status 200': (r) => r.status === 200,
    'login has token': (r) => JSON.parse(r.body).token !== undefined,
  });

  sleep(0.5);


  const failRes = http.post(
    `${BASE_URL}/login`,
    JSON.stringify({ username: 'logintest_user', password: 'wrong_password' }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(failRes, {
    'wrong password returns 401': (r) => r.status === 401,
  });

  sleep(0.5);
}