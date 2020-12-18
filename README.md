# GoYoutube-dl

Simple web server application to kick off youtube-dl. I created this app since there were no low profile, web facing application to run youtube dl

My current setup has this web server and Open Media Vault running in a Raspberry Pi 4 2GB. Youtube-dl configuration outputs to my NAS which is accessible through KODI.

## Usage:
````
Step 1: run nohup go run main.go &
Step 2: Open up a browser and go to http://ipaddress:8080
Step 3: Copy paste the URL and click submit
````

#### youtube-dl config under ~/.config/youtube-dl/config

````
-o /file_path_to_NAS_mount/%(title)s.%(ext)s -f bestvideo+bestaudio/best

````
## License
[MIT](https://choosealicense.com/licenses/mit/)
