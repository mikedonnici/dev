# Dates and Times 

```python
from datetime import date

mike_bday = date(1970, 11, 3) # create a date type, Mike's birthday
mike_bday.weekday()           # numeric day of week 0 = Monday
# 1 (Tuesday) 
mike_bday.day                 # day attribute of the date
# 3 
mike_bday.year
# 1970
```

## Working with dates

### Maths

- Python does basic arithmetic with dates.
- The result of adding / subtracting two dates is an object of `timedelta` type
- The attributes of a `timedelta` contain the differences between the two dates

```python
from datetime import date
d1 = date(1970, 11, 3) # Mike's birthday 
d2 = date(2012, 11, 2) # Leo's birthday
min([d2, d1])
# datetime.date(1970, 11, 3)
delta = d2 - d1
type(delta)
# datetime.timedelta
delta.days
# 15340
```

- `timedelta` can also be used to add or subtract time from a date

```python
from datetime import date, timedelta
d = date(2020, 1, 1)
td = timedelta(days=29)
d + td
# datetime.date(2020, 1, 30) 
```

### String formatting

- Default is `YYYY-MM-DD` ISO 8601 format (referred to as 'ISO')
- to turn to a string use `.isoformat()`

```python
print(date(2020, 4, 23))                                                                                                                                                                                            
# 2020-04-23
d = date(2020, 4, 23).isoformat())
type(d)
# str
```

- Dates formatted as ISO 8601 strings will sort correctly

```python
dates = ["2012-11-02", "1970-11-03", "2010-03-09", "1975-07-31"]
sorted(dates)
# ['1970-11-03', '1975-07-31', '2010-03-09', '2012-11-02']
```

- Other date format strings can be created with `.strftime()`
- `strftime()` will accept any string with placeholders

```python
d = date(2013, 3, 23)
print(d.strftime("%Y"))
# 2013
print(d.strftime("%Y-%D"))
# 2013-23
print(d.strftime("The month was %B"))
# The month was March
``` 

### Dateimes 

```python
from datetime import datetime

datetime(2020, 3, 25, 15, 18, 37) 
# datetime.datetime(2020, 3, 25, 15, 18, 37)
datetime(2020, 3, 25, 15, 18, 37, 500000) # includes microseconds
```

- `.replace()` method will replace the components of a datetime with the specified values

```python
dt = datetime(2020, 3, 25, 15, 18, 37)
dt2 = dt.replace(minute=0, second=0)   # singular
print(dt2)
# 2020-03-25 15:00:00
```

- `isoformat()` for a date tiem looks like: `2020-04-15T10:46:20`
- `strftime()` (string format time) formats `datetime` objects into strings 
- `strptime()` (string parse time) parses a time string and converts it to a `datetime` object
- `.fromtimestamp()` converts a unix timestamp to a datetime object

#### Durations

Addition and subtraction of `datetime` objects also results in a `timedelta`.

```python
from datetime import datetime
start = datetime(2020, 1, 1, 10, 30)
end = datetime(2020, 2, 2, 16, 45)
td = end - start
type(td)
# datetime.timedelta
td.total_seconds()
# 2787300.0
```

- Similarly, can use `timedelta` objects to creates datatime values adjusted by a specified duration. 

```python
start = datetime(2020, 1, 1, 10, 30)
elapsed = timedelta(hours=175)
end = start + elapsed
end
# datetime.datetime(2020, 1, 8, 17, 30)
```

- Can also use negative `timedelta` values

### Timezones

- Naive `datetime` objects don't know anything about their timezone
- Timezome-award `datetimes` require the use of the `timezone` package

```python
from datetime import datetime, timedelta, timezone

# AU Eastern time
tz = timezone(timedelta(hours=10))
dt = datetime(1970, 11, 3, 7, 46, 0, tzinfo=tz)
print(dt)
# 1970-11-03 07:46:00+10:00
IST = timezone(timedelta(hours=5, minutes=30)) # Indian Standard Time UTC + 5:30
print(dt.astimezone(IST))
# 1970-11-03 03:16:00+05:30
```

- Important to distinguish between changing the `tzinfo` for a `datetime` and _adjusting_ the `datedate` to be in 
a different timezone

```python
# 1st Jan 2020 at 10am, eastern australian time
dt = datetime(2020, 1, 1, 10, 0, tzinfo=timezone(timedelta(hours=10)))
dt.isoformat())
# '2020-01-01T10:00:00+10:00'
dt.replace(tzinfo=timezone.utc).isoformat() # same clock time, different time zone
# '2020-01-01T10:00:00+00:00'
dt.astimezone(timezone.utc).isoformat()     # same moment, different clock time
# '2020-01-01T00:00:00+00:00'                  
```

#### TZ Timezone database

- Timezone information changes frequently
- A global timezone database (tz) is used by many programming languages
- Timezone are keyed with 'Continent/City'
- TZ is available in Python via the `dateutils` module
- `tz` automatically adjusts for daylight saving



```python
from datetime import datetime
from dateutil import tz

et = tz.gettz("America/New_York")
dt = datetime(2020, 4, 14, 14, 30, tzinfo=et)
dt.isoformat()
# '2020-04-14T14:30:00-04:00'
```

- When daylight savings ends clocks are turned back 
- This means an hour of _clock_ times will appear twice
- Thus, it is possible for a `datetime` to be ambiguous
- `tz.datetime_ambiguous(dt)` returns True if this is the case
- `tz.enfold(dt)` specifies that `dt` represents the _second_ time the wall clock showed this time 
- subsequent conversion to UTC will then ensure durations are correct 

### datetimes with pandas

See [pandas/#working-with-datetimes](../pandas/#working-with-datetimes)




  


