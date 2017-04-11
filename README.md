# pubsub

A publish/subscribe mechanism for use intra-program, to decouple producers of messages from the consumers. The subscribers can unregister themselves by simpling returning false from their callback. That's useful when a closed socket or channel has been detected.

Example uses:
 * Logging to multiple end points
 * Pushing state changes to open WebSocket clients

## Problems needing solutions/FAQ

* What if we have a very slow callback in a publish cycle? Should it block the other callbacks?
  * We require that all callbacks be fast and well behaved - if it needs to do slow stuff, defer it to a goroutine

* What happens if another goroutine adds a subscriber during an active publish loop?
  * That subscription operation will fail and return ErrBrokerSaturated
  * TODO: write a test demo'ing this problem and consider fixing it by assigning the new channel early, before the old one is closed
