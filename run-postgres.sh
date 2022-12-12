#!/bin/bash
# service postgresql status
# service postgresql start
# service postgresql stop
psql -h localhost -p 5432 -U postgres -W bookings
