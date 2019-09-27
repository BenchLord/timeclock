# Timeclock

## expected sesh
```sh
[~] ➜ clock new
'Name: ' # string from stdin
'Wage: ' # double from stdin

# save user to config file or something...

[~] ➜ clock in
'clocked in at 9:00 am'
[~] ➜ clock in
'already clocked in. clocked in at 9:00 am'
[~] ➜ clock out
'clocked out at 5:00 pm. worked for 8 hrs 0 mins';
[~] ➜ clock out
'you are not clocked in'
[~] ➜ clock hours
'Total hours: 17.32'
'Estimated net: $9001'

# should be able to get hours by month, year, and period
```

needs to have commands for logging in/out
needs to use a config file