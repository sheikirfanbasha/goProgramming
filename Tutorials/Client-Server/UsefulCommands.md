To know which process is using particular port on Mac :

sudo lsof -nP -iTCP:$PORT -sTCP:LISTEN


To kill a process with PID on Mac:

sudo kill -9 PID