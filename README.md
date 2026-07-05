# event ticketing system

a microservices-based event ticketing platform where users can register, browse events, and book tickets. the goal is to build something that handles the full ticketing flow, from user auth all the way to booking and seat management, while keeping each concern in its own service.

---

## tech stack

- user service: spring boot (java) + postgresql
- ticket service: go + postgresql + redis
- booking service: spring boot (java) + postgresql
- activity service (coming up): go + mongodb


**infrastructure:** docker + docker compose, rabbitmq (async messaging) (not yet here), redis (seat locking)

---

## progress

**what's cooking:** working on user service auth issues. writing tests.

### v0.6
- added a new state to booking which is: 'abandoned'. abandoned bookings are the ones that user didn't complete the payment for and didn't make an attempt on it for the last 5 minutes. 
- added new fields for booking entity for this change, with changes to service and controller side.
- also added a call to ticket service that which sets the tickets assigned to ticket from 'reserved' to 'available' again if the booking is now abandoned.
- tested both services to make sure they are working in rhythm together.


### v0.5
- changed the implementation of locking and releasing tickets, with special changes on dto and method names ('seat' -> 'ticket')
- added a new endpoint for ticket service which checks if tickets are reserved for user.
- added a logic which sets ticket statuses as sold while making a booking
- added a confirmation logic for bookings, so that after payment the booking can be fully complete.
- fixed booking flow, now it reaches ticket service and gets booking item information there

### v0.4
- added restclient and its plumbing so that services can talk to each other (in booking service)
- removed event_seat dto from ticket as it feels like "over-engineering"
- removed reserving tickets for users in ticket service as it is now fully booking service's job.

### v0.3 
- booking service up and running with **lots** of bugs to fix, later
- ticket service is now fully on docker

### v0.2 
- ticket service: both event and ticket crud completed
- booking service: in progress

### v0.1 
- user service: basic crud done
- ticket inventory: events crud in progress
- services are running independently via docker compose
- each service has its own database
