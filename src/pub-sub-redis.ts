import { Redis } from "ioredis";

export const redis = new Redis();

if (redis.status !== "ready") {
  console.log("connecting to redis........");
}

console.log("status of redis", redis.status);

// setInterval(() => {
//   const message = { number: Math.random() };
//   const channel = `random-channel-${1 + Math.round(Math.random())}`;

//   redis.publish(channel, JSON.stringify(message));

//   console.log(`message published %s to %s`, JSON.stringify(message), channel);
// }, 2000);
