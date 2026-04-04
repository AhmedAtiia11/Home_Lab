# Send 500 rapid requests
for i in $(seq 1 500); do
  curl -s -H "Host: backend.local" http://192.168.1.15:8080/health 
done
wait