import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 10 },
   // { duration: '10s', target: 10 },
   // { duration: '10s', target: 50 },
   // { duration: '10s', target: 50 },
   // { duration: '10s', target: 0 },
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],
    http_req_failed: ['rate<0.01'],
  },
};

const BASE_URL = 'http://app:8000';
const BATCH_SIZE = 20;

export default function () {
  
  const customerIds = [];

  for (let i = 0; i < BATCH_SIZE; i++) {
    const floors = [1, 2, 3];
    const genders = ['male', 'female'];
    const ageGroups = ['adult', 'child'];

    const floor = floors[Math.floor(Math.random() * floors.length)];
    const gender = genders[Math.floor(Math.random() * genders.length)];
    const ageGroup = ageGroups[Math.floor(Math.random() * ageGroups.length)];

    const enterRes = http.post(
      `${BASE_URL}/?action=enter&Floor=${floor}&Gender=${gender}&AgeGroup=${ageGroup}`
    );

    check(enterRes, {
      'enter status 200': (r) => r.status === 200,
      'enter has id': (r) => JSON.parse(r.body).id !== undefined,
    });

    customerIds.push(JSON.parse(enterRes.body).id);
    sleep(0.1);
  }

  sleep(2);

  
  for (let i = 0; i < customerIds.length; i++) {
    const payment = Math.floor(Math.random() * 200) + 50;

    const exitRes = http.post(
      `${BASE_URL}/?action=exit&id=${customerIds[i]}&Payment=${payment}`
    );

    check(exitRes, {
      'exit status 200': (r) => r.status === 200,
    });

    sleep(0.1);
  }

  sleep(1);
}