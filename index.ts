console.log("Hello via Bun!");
import { Redis } from "ioredis";
const redis = new Redis();
redis.subscribe("random-channel-1", (err, count) => {
  if (err) {
    console.log("error occurred in the reading message ");
  } else {
    console.log(`subscribed successfully `, count);
  }
});

redis.subscribe("random-channel-2", (err, count) => {
  if (err) {
    console.log("error occured in the reading message");
  } else {
    console.log("subscribed another random-channel-2", count);
  }
});

redis.on("message", (channel, message) => {
  console.log(`Received ${message} from ${channel}`);
});
