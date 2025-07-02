# rickslow
Slow HTTP responses

# Demo

```bash
# Choose the correct OS and Architecture at https://github.com/mikhasd/rickslow/releases
curl -L rickslow.tar.gz -o https://github.com/mikhasd/rickslow/releases/download/v0.0.6-beta/rickslow-linux-amd64.tar.gz
tar -xzf rickslow.tar.gz
./rickslow 2> /dev/null & rickpid=$!
curl http://localhost:10888

# Use to kill the process runnin in background
kill $rickpid
```
