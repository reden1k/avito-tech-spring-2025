import http from 'k6/http';
import { check } from 'k6';
import { Trend } from 'k6/metrics';

export let loginResponseTime = new Trend('login_response_time');

export let options = {
  vus: 40,       
  duration: '30s', 
};

export default function () {
  const payload = JSON.stringify({
    email: 'user@example.com',
    password: 'password123',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let res = http.post('http://localhost:8080/login', payload, params);

  check(res, {
    'is status 200': (r) => r.status === 200,
  });

  loginResponseTime.add(res.timings.duration);
}
