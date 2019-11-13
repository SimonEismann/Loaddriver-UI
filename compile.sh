docker build -t simoneismann/loaddriver-master:latest ./loaddriver-master &
docker build -t simoneismann/loaddriver-slave:latest ./loaddriver-slave &
docker build -t simoneismann/loaddriver-frontend:latest ./loaddriver-frontend &
wait