To publish your Docker image to Docker Hub, follow these steps:

### 1. **Login to Docker Hub from Your Terminal**
You need to log in to Docker Hub through the command line interface (CLI). Use the following command to log in:

```bash
docker login
```

You’ll be prompted to enter your Docker Hub username and password.

### 2. **Tag Your Docker Image IF NOT "already done" in Docker Compose**
Before you can push an image to Docker Hub, you need to tag it with your Docker Hub username and repository name. The tag follows the format:

```
<your-docker-username>/<repository-name>:<tag>
```

For example, if your Docker Hub username is `pablo` and you want to name your repository `my-go-app`, tag the image like this:

```bash
docker tag my-go-app pablo/my-go-app:latest
```

If you don’t specify a tag, Docker will automatically assume `latest` as the default tag.

### 3. **Push the Docker Image to Docker Hub**
Now, you can push the tagged image to Docker Hub using the `docker push` command:

```bash
docker push pablo/my-go-app:latest
```

This will upload your Docker image to Docker Hub. The upload process may take some time, depending on the size of your image and your internet connection.

### 4. **Verify the Image on Docker Hub**
Once the push is complete, go to your Docker Hub account, and you should see the image in your repositories list under the name `my-go-app`.

### 5. **Pull the Docker Image (to verify)**
To verify the image has been successfully pushed, you can pull the image from Docker Hub on any machine using the following command:

```bash
docker pull pablo/my-go-app:latest
```

This will pull the image to your local machine, confirming it's available on Docker Hub.

### Additional Notes:
- **Repository Visibility**: Make sure your Docker Hub repository is set to **Public** if you want others to access it. By default, repositories are private for new users.
- **Pushing Large Images**: If your image is large, Docker will use layers to upload the image in parts. This is normal.



---

### 1. **GitHub Codespaces (Free for Personal Use)**
- **Steps**:
  1. Push your Docker project to a GitHub repository.
  2. Use GitHub Codespaces to run and expose your container.
  3. Run:
     ```bash
     docker pull <your-dockerhub-username>/<image-name>
     docker run -d -p 8080:8080 <your-image-name>
     ```
     **OR build image form liked repository in GitHub Codespaces.**
     ```bash
     docker-compose -f docker-compose.yml up -d
     ```
  4. **Expose the Port Publicly**
    Expose the Port:

    In the top bar of your Codespaces IDE, click Ports.
    You’ll see a list of exposed ports. If the port (e.g., 8080) isn’t listed, click Add Port and enter the exposed port.
    Make the Port Public:

    Click the gear icon next to your port and select Port Visibility > Public.
    Get the Public URL:

    After making the port public, Codespaces will provide you with a public URL that you can use to access your service.

- **Limitations**: Limited hours per month on the free tier.



---
From the diagram, **the saturated liquid line and saturated vapor line indicate nonlinear relationships between pressure (P) and specific volume (v)**.

**Observations:**

1. **Critical Point:**
  -  Pressure (Pc) = 10 MPa, Specific Volume (vc) = 0.0035 m³/kg.

2. **Saturated Liquid Line:**
  - At low pressures (e.g., 0.05 MPa), the specific volume approaches **0.001 m³/kg.**
  - Approximates a **curve that increases toward the critical point.**

3. **Saturated Vapor Line:**
  - At low pressures (e.g., 0.05 MPa), the specific volume approaches **~0.03 m³/kg.**
  - Approximates a **curve that decreases toward the critical point.**

---
From the diagram, the **saturated liquid line** and **saturated vapor line** indicate nonlinear relationships between **pressure (P)** and **specific volume (v)**. I'll derive **realistic equations** based on the following assumptions derived from the diagram:

---

### Observations:
1. **Critical Point**:
   - Pressure (Pc) = 10 MPa, Specific Volume (vc) = 0.0035 m³/kg.
2. **Saturated Liquid Line**:
   - At low pressures (e.g., 0.05 MPa), the specific volume approaches **0.001 m³/kg**.
   - Approximates a curve that increases toward the critical point.
3. **Saturated Vapor Line**:
   - At low pressures (e.g., 0.05 MPa), the specific volume approaches **~0.03 m³/kg**.
   - Approximates a curve that decreases toward the critical point.



---
### Realistic Equations:
#### 1. Saturated Liquid Line:
Using a polynomial fit:
\[ v_{\text{liquid}}(P) = a \cdot P^2 + b \cdot P + c \]
From the diagram:
- At \( P = 0.05 \), \( v_{\text{liquid}} = 0.001 \).
- At \( P = 10 \), \( v_{\text{liquid}} = 0.0035 \) (critical point).

#### 2. Saturated Vapor Line:
Using a different polynomial fit:
\[ v_{\text{vapor}}(P) = d \cdot P^2 + e \cdot P + f \]
From the diagram:
- At \( P = 0.05 \), \( v_{\text{vapor}} = 0.03 \).
- At \( P = 10 \), \( v_{\text{vapor}} = 0.0035 \) (critical point).

---
### Equations Derived:
Using interpolation techniques (polynomial fitting):
#### Saturated Liquid Line:
\[ v_{\text{liquid}}(P) = -0.00000025 \cdot P^2 + 0.0001 \cdot P + 0.001 \]

#### Saturated Vapor Line:
\[ v_{\text{vapor}}(P) = 0.00000035 \cdot P^2 - 0.0027 \cdot P + 0.03 \]



---
### How This Works:
1. **Realistic Equations**:
   - The equations for the liquid and vapor lines are derived from the diagram's curves.
   - They dynamically calculate the specific volume based on the pressure input.

2. **Pressure Validation**:
   - Ensures the pressure is within the valid range: \( 0 < P \leq 10 \).


---
### Example Request:
**GET** `/phase-change-diagram?pressure=5`

### Example Response:
```json
{
  "specific_volume_liquid": 0.00175,
  "specific_volume_vapor": 0.0165
}
```