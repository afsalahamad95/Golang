import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 50, // virtual users
  duration: "45s", // test duration
};
function generator() {
  let max = 85;
  let min = 30;
  let value = Math.random() * (max - min) + min;
  return value;
}
export default function () {
  const url = "http://localhost:4000/";
  const payload = JSON.stringify({
    Lat: generator(),
    Lon: generator(),
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(url, payload, params);
  check(res, {
    "status is 200": (r) => r.status === 200,
    "response time < 900ms": (r) => r.timings.duration < 900,
  });
  sleep(1);
}
