# Distributed Task Queue

High-performance task queue with Redis backend.

## Features
- At-least-once delivery
- Dead letter queues
- Worker auto-scaling
- Retry with exponential backoff
- Prometheus metrics

## Architecture
```
Producer → Redis Queue → Worker Pool → Result Store
```
