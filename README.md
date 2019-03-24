# pipe

a minimalist CI/CD workflow, it exposes a webhook that recieves a `POST` request on the `/push` endpoint then proceeds to run your shell script.

- Download the pipe binary

```bash
curl -s "https://github.com/pariola/pipe/releases/download/v1/main" -L -o main
```

- Make binary executable

```bash
chmod +x ./main
```

- Create your entrypoint script `push.sh`

```bash
#!/bin/sh

# cleanup
echo "Cleaning up!"
rm -rf "$(pwd)/example"


git clone https://blessingpariola@bitbucket.org/blessingpariola/pipe-test.git "$(pwd)/example"
echo "Cloning complete"
```

- Make script executable

```bash
chmod +x ./push.sh
```

- Run

```bash
./main -script="./push.sh" -port=3000
```

- Create a webhook with your fav version control provider {Github, Bitbucket, etc..} then point it to your URL/IP with your specified port
  - https://confluence.atlassian.com/bitbucket/manage-webhooks-735643732.html
  - https://developer.github.com/webhooks/


## NOTE
- You can aswell add this line to the end of your crontab file in `/etc/crontab` so the command is automatically started on system reboot
```bash
@reboot root cd ~ && ./main -script="./push.sh" -port=3000
```