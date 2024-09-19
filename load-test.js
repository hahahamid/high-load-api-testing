import http from 'k6/http';
import { check } from 'k6';

export let options = {
    scenarios: {
        high_load_test: {
            executor: 'constant-arrival-rate',
            rate: 833, // 833 requests per second
            timeUnit: '1s', 
            duration: '1m', 
            preAllocatedVUs: 1000, 
            maxVUs: 1500, 
        },
    },
    thresholds: {
        'http_req_duration': ['p(95)<500'], 
        'http_req_failed': ['rate<0.01'],  
    },
};

export default function () {
    let res = http.post('http://localhost:8080/echo', JSON.stringify({ message: 'Load test for 50K RPM' }), {
        headers: { 'Content-Type': 'application/json' },
    });
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time < 500ms': (r) => r.timings.duration < 500,
    });
}
