### Step-by-Step Guide to Expose API Using **ngrok**:

1. **Install ngrok**  
   Download and install ngrok from [https://ngrok.com/download](https://ngrok.com/download).  
   - On Linux/Mac:  
     ```sh
     sudo apt install ngrok
     ```  
     **OR**
     ```sh
     sudo snap install ngrok

     ngrok config add-authtoken YOUR_AUTH_TOKEN
     ```
   - On Windows:  
     Download the `.exe` file and add it to your PATH.

2. **Start Your Go API Locally**  
   Run your Go API on a specific port, e.g., `localhost:8080`:  
   ```sh
   go run main.go
   ```

3. **Expose Your API Using ngrok**  
   Start ngrok with the same port:  
   ```sh
   ngrok http 8080
   ```

4. **Get Public URL**  
   After running the command, ngrok will display a public URL, like:  
   ```
   Forwarding   https://random-subdomain.ngrok.io -> http://localhost:8080
   ```

5. **Access Your API Publicly**  
   Use the `https://random-subdomain.ngrok.io` URL in a browser, Postman, or other clients to access your local API.

---

### Advantages of ngrok:  
- **Simple Setup**: No SSH setup required.  
- **Custom Subdomains** (with paid plan): You can set a memorable domain like `https://myapi.ngrok.io`.  
- **HTTPS Support**: Secure requests via HTTPS.  

### Important Notes:  
- Keep the terminal open for ngrok to work.  
- Free accounts have session time limits (e.g., 8 hours).

To stop **ngrok**, simply:

1. **Press `Ctrl+C`** in the terminal where ngrok is running.
2. Alternatively, use this command to kill any ngrok processes:  
   ```sh
   pkill ngrok
   ```  
   *(For Windows, you can stop it from the Task Manager or use `taskkill /IM ngrok.exe /F` in Command Prompt.)*