Both in Go and in C, we get different values each time, and almost never zero (as it actually should be).
It seems like both function gets access to the same varibale at the same time. Which results in race condtions. 
So even when we put (in C) that this function should be done before the other function start, we still get race condtions 
because they still access the value at the same time. 