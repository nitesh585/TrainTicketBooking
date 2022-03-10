# Train-ticket booking systems

## Functional Requirements:

### user management
- user can register, login and logout
- maintain history
    - booked , canceled, failed tickets
    - ticket refunds
- update profile
    - password and basic personal details
    - add aadhaar card with KYC


### search engine
- check for the train availability on the basis of source, destination stations and date
- if direct route not found then show different available connected path from source to   
  destination
- add filters of classes, seat priority, etc.
- user should get list of available trains with further details:
- seats available, RAC & waiting list count
- arrival time and departure time
- estimated journey time
- train schedule
- pricing
- user can update the search query


### booking engine
- book ticket for selected train
- user must logged in before booking
- get the details of passengers (age, name and seat priority) and give a review page to user
- after reviewing, redirect to the payment method.
- cancel ticket after booking
    - if cancellation is after chart preparation or as per law, then refund after deduction otherwise not
    - user can cancel whole ticket or cancel ticket for one of the passenger
- give invoice or ticket in the end of successful booking on email or sms as per the subscribed service


### notification module
- user should get alerts about tickets on subscribed services eg. email or sms
- should give alert to user about his/her ticket status
    - alert if the ticket got confirmed from RAC or WL
payment method
- integrate one of payment gateways
- user must logged in and selected at least one seat
- pay for the tickets
- cancel pay anytime before

## Concepts/Topic covered:
- gin-gonic/gin framework
- JWT integration
- Introducing middlewares
- Added Unit test (table-driven)
- MondoDb integration

## Supported APIs:
- /user/signup
- /user/login
- /train/searchRoute

## Working
- integrated payment feature 
- integration of docker build
- logout 