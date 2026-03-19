# event ticketing system

a microservices-based event ticketing platform where users can register, browse events, and book tickets. the goal is to build something that handles the full ticketing flow — from user auth all the way to booking and seat management — while keeping each concern in its own service.

---

## tech stack

- user service: spring boot (java) + postgresql
- ticket service: go + postgresql + redis
- booking service (coming up): spring boot (java) + postgresql
- activity service (cpming up): go + mongodb


**infrastructure:** docker + docker compose, rabbitmq (async messaging), redis (seat locking)

---

## progress

### v0.1 — week 1
- user service: basic crud done
- ticket inventory: events crud in progress
- services are running independently via docker compose
- each service has its own database